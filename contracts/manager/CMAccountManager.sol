// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Account Manager

pragma solidity ^0.8.24;

// UUPS Proxy
import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import { UUPSUpgradeable } from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

// Access
import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import { AccessControlEnumerableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlEnumerableUpgradeable.sol";
import { ReentrancyGuardUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";

// ABI of the CMAccount implementation contract
import { ICMAccount } from "../account/ICMAccount.sol";

// Utils
import { Address } from "@openzeppelin/contracts/utils/Address.sol";

// Service Registry
import { ServiceRegistry } from "../partner/ServiceRegistry.sol";

/**
 * @title Camino Messenger Account Manager
 * @notice This contract manages the creation of the Camino Messenger accounts by
 * deploying {ERC1967Proxy} proxies that point to the{CMAccount} implementation
 * address.
 *
 * Create CM Account: Users who want to create an account should call
 * `createCMAccount(address admin, address upgrader)` function with addresses of
 * the accounts admin and upgrader roles and also send the pre fund amount,
 * which is currently set as 100 CAMs. When the manager contract is paused,
 * account creation is stopped.
 *
 * Developer Fee: This contracts also keeps the info about the developer wallet
 * and fee basis points. Which are used during the cheque cash in to pay for the
 * developer fee.
 *
 * Service Registry: {CMAccountManager} also acts as a registry for the services
 * that {CMAccount} contracts add as a supported or wanted service. Registry
 * works by hashing (keccak256) the service name (string) and creating a mapping
 * as keccak256(serviceName) => serviceName. And provides functions that
 * {CMAccount} function uses to register services. The {CMAccount} only keeps
 * the hashes (byte32) of the registered services.
 * @custom:security-contact https://r.xyz/program/camino-network
 */
contract CMAccountManager is
    Initializable,
    PausableUpgradeable,
    AccessControlEnumerableUpgradeable,
    UUPSUpgradeable,
    ReentrancyGuardUpgradeable,
    ServiceRegistry
{
    using Address for address payable;

    /**
     * @notice Pauser role can pause the contract. Currently this only affects the
     * creation of CM Accounts. When paused, account creation is stopped.
     */
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    /**
     * @notice Upgrader role can upgrade the contract to a new implementation.
     */
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

    /**
     * @notice Versioner role can set new {CMAccount} implementation address. When a
     * new implementation address is set, it is used for the new {CMAccount}
     * creations.
     *
     * The old {CMAccount} contracts are not affected by this. Owners of those
     * should do the upgrade manually by calling the `upgradeToAndCall(address)`
     * function on the account.
     */
    bytes32 public constant VERSIONER_ROLE = keccak256("VERSIONER_ROLE");

    /**
     * @notice Fee admin role can set the developer fee basis points which used for
     * calculating the developer fee that is cut from the cheque payments.
     */
    bytes32 public constant FEE_ADMIN_ROLE = keccak256("FEE_ADMIN_ROLE");

    /**
     * @notice Developer wallet admin role can set the developer wallet address
     * which is used to receive the developer fee.
     */
    bytes32 public constant DEVELOPER_WALLET_ADMIN_ROLE = keccak256("DEVELOPER_WALLET_ADMIN_ROLE");

    /**
     * @notice Prefund admin role can set the mandatory prefund amount for {CMAccount}
     * contracts.
     */
    bytes32 public constant PREFUND_ADMIN_ROLE = keccak256("PREFUND_ADMIN_ROLE");

    /**
     * @notice Service registry admin role can add and remove services to the service
     * registry mapping. Implemented by {ServiceRegistry} contract.
     */
    bytes32 public constant SERVICE_REGISTRY_ADMIN_ROLE = keccak256("SERVICE_REGISTRY_ADMIN_ROLE");

    /**
     * @notice This role is granted to the created CM Accounts. It is used to keep
     * an enumerable list of CM Accounts.
     */
    bytes32 public constant CMACCOUNT_ROLE = keccak256("CMACCOUNT_ROLE");

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    /**
     * @notice CMAccount info struct, to keep track of created CM Accounts and their
     * creators.
     */
    struct CMAccountInfo {
        bool isCMAccount;
        address creator;
    }

    /// @custom:storage-location erc7201:camino.messenger.storage.CMAccountManager
    struct CMAccountManagerStorage {
        /**
         * @dev CM Account implementation address to be used by the CMAccount contract to restrict
         * the implementation address for the UUPS proxies.
         */
        address _latestAccountImplementation;
        /**
         * @dev Prefund amount.
         */
        uint256 _prefundAmount;
        /**
         * @dev Developer wallet address. CMAccount sends the developer fee to this address.
         */
        address _developerWallet;
        /**
         * @dev Developer fee basis points.
         *
         * A basis point (bp) is one hundredth of 1 percentage point.
         *
         * 1 bp = 0.01%, 1/10,000⁠, or 0.0001.
         * 10 bp = 0.1%, 1/1,000⁠, or 0.001.
         * 100 bp = 1%, ⁠1/100⁠, or 0.01.
         */
        uint256 _developerFeeBp;
        /**
         * @dev BookingToken address.
         */
        address _bookingToken;
        /**
         * @dev CMAccount info mapping to track if an address is a CMAccount and initial creators.
         */
        mapping(address account => CMAccountInfo) _cmAccountInfo;
    }

    // keccak256(abi.encode(uint256(keccak256("camino.messenger.storage.CMAccountManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant CMAccountManagerStorageLocation =
        0x2b421af391835920c41e77b6810f6e715f5b713c17bc590f55de6a7d3912e800;

    function _getCMAccountManagerStorage() private pure returns (CMAccountManagerStorage storage $) {
        assembly {
            $.slot := CMAccountManagerStorageLocation
        }
    }

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    /**
     * @notice CM Account created event.
     * @param account The address of the new CMAccount
     */
    event CMAccountCreated(address indexed account);

    /**
     * @notice CM Account implementation address updated event.
     * @param oldImplementation The old implementation address
     * @param newImplementation The new implementation address
     */
    event CMAccountImplementationUpdated(address indexed oldImplementation, address indexed newImplementation);

    /**
     * @notice Developer wallet address updated event.
     * @param oldDeveloperWallet The old developer wallet address
     * @param newDeveloperWallet The new developer wallet address
     */
    event DeveloperWalletUpdated(address indexed oldDeveloperWallet, address indexed newDeveloperWallet);

    /**
     * @notice Developer fee basis points updated event.
     * @param oldDeveloperFeeBp The old developer fee basis points
     * @param newDeveloperFeeBp The new developer fee basis points
     */
    event DeveloperFeeBpUpdated(uint256 indexed oldDeveloperFeeBp, uint256 indexed newDeveloperFeeBp);

    /**
     * @notice Booking token address updated event.
     * @param oldBookingToken The old booking token address
     * @param newBookingToken The new booking token address
     */
    event BookingTokenAddressUpdated(address indexed oldBookingToken, address indexed newBookingToken);

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    /**
     * @notice The implementation of the CMAccount is invalid.
     * @param implementation The implementation address of the CMAccount
     */
    error CMAccountInvalidImplementation(address implementation);

    /**
     * @notice The admin address is invalid.
     * @param admin The admin address
     */
    error CMAccountInvalidAdmin(address admin);

    /**
     * @notice Invalid developer address.
     * @param developerWallet The developer wallet address
     */
    error InvalidDeveloperWallet(address developerWallet);

    /**
     * @notice Invalid booking token address.
     * @param bookingToken The booking token address
     */
    error InvalidBookingTokenAddress(address bookingToken);

    /**
     * @notice Incorrect pre fund amount.
     * @param expected The expected pre fund amount
     */
    error IncorrectPrefundAmount(uint256 expected, uint256 sended);

    /***************************************************
     *                    FUNCS                        *
     ***************************************************/

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(
        address defaultAdmin, // can grant roles
        address pauser, // can pause the manager
        address upgrader, // can upgrade the manager (this contract)
        address versioner, // can set CMAccount implementation address
        address developerWallet, // developer wallet used to receive the developer fee
        uint256 developerFeeBp // developer fee basis points
    ) public initializer {
        __Pausable_init();
        __AccessControl_init();
        __UUPSUpgradeable_init();

        _grantRole(DEFAULT_ADMIN_ROLE, defaultAdmin);
        _grantRole(PAUSER_ROLE, pauser);
        _grantRole(UPGRADER_ROLE, upgrader);
        _grantRole(VERSIONER_ROLE, versioner);

        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();

        $._developerWallet = developerWallet;
        $._developerFeeBp = developerFeeBp;

        // Set initial prefund amount to 100 CAM
        $._prefundAmount = 100 ether;
    }

    /**
     * @notice Pauses the CMAccountManager contract. Currently this only affects the
     * creation of CMAccount. When paused, account creation is stopped.
     */
    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @notice Unpauses the CMAccountManager contract.
     */
    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /**
     * @notice Authorization for the CMAccountManager contract upgrade.
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(UPGRADER_ROLE) {}

    /***************************************************
     *                    ACCOUNT                      *
     ***************************************************/

    /**
     * @notice Creates CMAccount by deploying a ERC1967Proxy with the CMAccount
     * implementation from the manager.
     *
     * Because this function is deploying a contract, it reverts if the caller is
     * not KYC or KYB verified. (For EOAs only)
     *
     * Caller must send the pre-fund amount with the transaction.
     *
     * @dev Emits a {CMAccountCreated} event.
     */
    function createCMAccount(
        address admin,
        address upgrader
    ) external payable nonReentrant whenNotPaused returns (address) {
        return _createCMAccount(admin, upgrader);
    }

    /**
     * @notice Private function to create a `CMAccount`.
     */
    function _createCMAccount(address admin, address upgrader) private returns (address) {
        // Checks
        if (admin == address(0)) {
            revert CMAccountInvalidAdmin(admin);
        }

        uint256 prefundAmount = getPrefundAmount();

        // Check pre-fund amount
        if (msg.value < prefundAmount) {
            revert IncorrectPrefundAmount(prefundAmount, msg.value);
        }

        address latestAccountImplementation = getAccountImplementation();
        if (latestAccountImplementation.code.length == 0) {
            revert CMAccountInvalidImplementation(latestAccountImplementation);
        }

        address bookingToken = getBookingTokenAddress();
        if (bookingToken.code.length == 0) {
            revert InvalidBookingTokenAddress(bookingToken);
        }

        // Create CMAccount Proxy and set the implementation address
        ERC1967Proxy cmAccountProxy = new ERC1967Proxy(latestAccountImplementation, "");

        // Initialize the CMAccount
        ICMAccount(address(cmAccountProxy)).initialize(address(this), bookingToken, prefundAmount, admin, upgrader);

        // Set the isCMAccount and creator
        _setCMAccountInfo(address(cmAccountProxy), CMAccountInfo({ isCMAccount: true, creator: msg.sender }));

        // Grant CMACCOUNT_ROLE
        _grantRole(CMACCOUNT_ROLE, address(cmAccountProxy));

        // Send the pre fund to the CMAccount
        payable(cmAccountProxy).sendValue(msg.value);

        emit CMAccountCreated(address(cmAccountProxy));

        return address(cmAccountProxy);
    }

    function _setCMAccountInfo(address account, CMAccountInfo memory info) internal {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        $._cmAccountInfo[account] = info;
    }

    /**
     * @notice Returns the given account's creator.
     * @param account The account address
     */
    function getCMAccountCreator(address account) public view returns (address) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        return $._cmAccountInfo[account].creator;
    }

    /**
     * @notice Check if an address is CMAccount created by the manager.
     * @param account The account address to check
     */
    function isCMAccount(address account) public view returns (bool) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        return $._cmAccountInfo[account].isCMAccount;
    }

    /***************************************************
     *             ACCOUNT IMPLEMENTATION              *
     ***************************************************/

    /**
     * @notice Returns the CMAccount implementation address.
     */
    function getAccountImplementation() public view returns (address) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        return $._latestAccountImplementation;
    }

    /**
     * @notice Set a new CMAccount implementation address.
     * @param newImplementation The new implementation address
     */
    function setAccountImplementation(address newImplementation) public onlyRole(VERSIONER_ROLE) {
        if (newImplementation.code.length == 0) {
            revert CMAccountInvalidImplementation(newImplementation);
        }

        address oldImplementation = getAccountImplementation();
        _setAccountImplementation(newImplementation);
        emit CMAccountImplementationUpdated(oldImplementation, newImplementation);
    }

    function _setAccountImplementation(address newImplementation) internal {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        $._latestAccountImplementation = newImplementation;
    }

    /***************************************************
     *                    PREFUND                      *
     ***************************************************/

    /**
     * @notice Returns the prefund amount.
     */
    function getPrefundAmount() public view returns (uint256) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        return $._prefundAmount;
    }

    /**
     * @notice Sets the prefund amount.
     */
    function setPrefundAmount(uint256 newPrefundAmount) public onlyRole(PREFUND_ADMIN_ROLE) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        $._prefundAmount = newPrefundAmount;
    }

    /***************************************************
     *                  BOOKING TOKEN                  *
     ***************************************************/

    /**
     * @notice Returns the booking token address.
     */
    function getBookingTokenAddress() public view returns (address) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        return $._bookingToken;
    }

    /**
     * @notice Sets booking token address.
     */
    function setBookingTokenAddress(address token) public onlyRole(VERSIONER_ROLE) {
        if (token.code.length == 0) {
            revert InvalidBookingTokenAddress(token);
        }

        address oldToken = getBookingTokenAddress();
        _setBookingTokenAddress(token);
        emit BookingTokenAddressUpdated(oldToken, token);
    }

    function _setBookingTokenAddress(address token) internal {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        $._bookingToken = token;
    }

    /***************************************************
     *            DEVELOPER WALLET & FEE               *
     ***************************************************/

    /**
     * @notice Returns developer wallet address.
     */
    function getDeveloperWallet() public view returns (address developerWallet) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        return $._developerWallet;
    }

    /**
     * @notice Sets developer wallet address.
     */
    function setDeveloperWallet(address developerWallet) public onlyRole(DEVELOPER_WALLET_ADMIN_ROLE) {
        if (developerWallet == address(0)) {
            revert InvalidDeveloperWallet(developerWallet);
        }

        address oldDeveloperWallet = getDeveloperWallet();
        _setDeveloperWallet(developerWallet);
        emit DeveloperWalletUpdated(oldDeveloperWallet, developerWallet);
    }

    function _setDeveloperWallet(address developerWallet) private {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        $._developerWallet = developerWallet;
    }

    /**
     * @notice Returns developer fee in basis points.
     */
    function getDeveloperFeeBp() public view returns (uint256 developerFeeBp) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        return $._developerFeeBp;
    }

    /**
     * @notice Sets developer fee in basis points.
     *
     * A basis point (bp) is one hundredth of 1 percentage point.
     *
     * 1 bp = 0.01%, 1/10,000⁠, or 0.0001.
     * 10 bp = 0.1%, 1/1,000⁠, or 0.001.
     * 100 bp = 1%, ⁠1/100⁠, or 0.01.
     */
    function setDeveloperFeeBp(uint256 bp) public onlyRole(FEE_ADMIN_ROLE) {
        uint256 oldBp = getDeveloperFeeBp();
        _setDeveloperFeeBp(bp);
        emit DeveloperFeeBpUpdated(oldBp, bp);
    }

    function _setDeveloperFeeBp(uint256 bp) private {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        $._developerFeeBp = bp;
    }

    /***************************************************
     *               SERVICE REGISTRY                  *
     ***************************************************/

    /**
     * @notice Registers a given service name. CM Accounts can only register services
     * if they are also registered in the service registry on the manager contract.
     *
     * @param serviceName Name of the service
     */
    function registerService(string memory serviceName) public onlyRole(SERVICE_REGISTRY_ADMIN_ROLE) {
        _registerServiceName(serviceName);
    }

    /**
     * @notice Unregisters a given service name. CM Accounts will not be able to register
     * the service anymore.
     *
     * @param serviceName Name of the service
     */
    function unregisterService(string memory serviceName) public onlyRole(SERVICE_REGISTRY_ADMIN_ROLE) {
        _unregisterServiceName(serviceName);
    }
}
