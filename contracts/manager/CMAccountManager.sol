// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Account Manager

pragma solidity ^0.8.24;

// UUPS Proxy
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

// Access
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlEnumerableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";

// ABI of the CMAccount implementation contract
import { ICMAccount } from "../account/ICMAccount.sol";

// Utils
import "@openzeppelin/contracts/utils/Address.sol";

// Service Registry
import "../partner/ServiceRegistry.sol";

/// @custom:security-contact https://r.xyz/program/camino-network
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
     * @dev Roles
     */
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    bytes32 public constant VERSIONER_ROLE = keccak256("VERSIONER_ROLE");
    bytes32 public constant FEE_ADMIN_ROLE = keccak256("FEE_ADMIN_ROLE");
    bytes32 public constant DEVELOPER_WALLET_ADMIN_ROLE = keccak256("DEVELOPER_WALLET_ADMIN_ROLE");
    bytes32 public constant PREFUND_ADMIN_ROLE = keccak256("PREFUND_ADMIN_ROLE");
    bytes32 public constant SERVICE_REGISTRY_ADMIN_ROLE = keccak256("SERVICE_REGISTRY_ADMIN_ROLE");

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    /**
     * @dev CMAccount info struct
     */
    struct CMAccountInfo {
        bool isCMAccount;
        address creator;
    }

    struct CMAccountManagerStorage {
        /**
         * @dev CM Account implementation address to be used by the CMAccount contract to resctrict
         * the implementation address for the UUPS proxies.
         */
        address _latestAccountImplementation;
        /**
         * @dev Prefund amount
         */
        uint256 _prefundAmount;
        /**
         * @dev Developer wallet address. CMAccount sends the developer fee to this address.
         */
        address _developerWallet;
        /**
         * @dev Developer fee basis points
         *
         * A basis point (bp) is one hundredth of 1 percentage point.
         *
         * 1 bp = 0.01%, 1/10,000⁠, or 0.0001.
         * 10 bp = 0.1%, 1/1,000⁠, or 0.001.
         * 100 bp = 1%, ⁠1/100⁠, or 0.01.
         */
        uint256 _developerFeeBp;
        /**
         * @dev BookingToken address
         */
        address _bookingToken;
        /**
         * @dev CMAccount info mapping to track if an address is a CMAccount and initial creators
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
     * @dev CM Account created
     * @param account The address of the new CMAccount
     */
    event CMAccountCreated(address indexed account);

    /**
     * @dev CM Account implementation address updated
     * @param oldImplementation The old implementation address
     * @param newImplementation The new implementation address
     */
    event CMAccountImplementationUpdated(address indexed oldImplementation, address indexed newImplementation);

    /**
     * @dev Developer wallet address updated
     */
    event DeveloperWalletUpdated(address indexed oldDeveloperWallet, address indexed newDeveloperWallet);

    /**
     * @dev Developer fee basis points updated
     */
    event DeveloperFeeBpUpdated(uint256 indexed oldDeveloperFeeBp, uint256 indexed newDeveloperFeeBp);

    /**
     * @dev Booking token address updated
     */
    event BookingTokenAddressUpdated(address indexed oldBookingToken, address indexed newBookingToken);

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    /**
     * @dev The implementation of the CMAccount is invalid
     * @param implementation The implementation address of the CMAccount
     */
    error CMAccountInvalidImplementation(address implementation);

    /**
     * @dev The admin address is invalid
     * @param admin The admin address
     */
    error CMAccountInvalidAdmin(address admin);

    /**
     * @dev Invalid developer address
     */
    error InvalidDeveloperWallet(address developerWallet);

    /**
     * @dev Invalid booking token address
     */
    error InvalidBookingTokenAddress(address bookingToken);

    /**
     * @dev Incorrect pre fund amount
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
        address defaultAdmin,
        address pauser,
        address upgrader, // upgrade the manager (this contract)
        address versioner, // sets cm account implementation address
        address developerWallet,
        uint256 developerFeeBp
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

    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /**
     * @dev Authorization for the CMAccountManager contract upgrade
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(UPGRADER_ROLE) {}

    /***************************************************
     *                    ACCOUNT                      *
     ***************************************************/

    /**
     * @dev Creates CMAccount by deploying a ERC1967Proxy with the CMAccount implementation from the manager.
     *
     * Because this function is deploying a contract, it reverts if the caller is not KYC or KYB verified.
     *
     * Emits a {CMAccountCreated} event.
     */
    function createCMAccount(
        address admin,
        address upgrader
    ) external payable nonReentrant whenNotPaused returns (address) {
        return _createCMAccount(admin, upgrader);
    }

    /**
     * @dev Private function to create CMAccount
     */
    function _createCMAccount(address admin, address upgrader) private returns (address) {
        // Checks
        address latestAccountImplementation = getAccountImplementation();
        if (latestAccountImplementation.code.length == 0) {
            revert CMAccountInvalidImplementation(latestAccountImplementation);
        }

        address bookingToken = getBookingTokenAddress();
        if (bookingToken.code.length == 0) {
            revert InvalidBookingTokenAddress(bookingToken);
        }

        if (admin == address(0)) {
            revert CMAccountInvalidAdmin(admin);
        }

        uint256 prefundAmount = getPrefundAmount();

        // FIXME: Investigate which checks are more likely to fail frequently and move them up
        // FIXME: Investigate if msg.value < prefund introduces any issues

        // Check pre-fund amount
        if (msg.value < prefundAmount) {
            revert IncorrectPrefundAmount(prefundAmount, msg.value);
        }

        // Create CMAccount Proxy and set the implementation address
        ERC1967Proxy cmAccountProxy = new ERC1967Proxy(latestAccountImplementation, "");

        // Set the isCMAccount and creator
        _setCMAccountInfo(address(cmAccountProxy), CMAccountInfo({ isCMAccount: true, creator: msg.sender }));

        // Initialize the CMAccount
        ICMAccount(address(cmAccountProxy)).initialize(address(this), bookingToken, prefundAmount, admin, upgrader);

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
     * @dev Return account's creator
     */
    function getCMAccountCreator(address account) public view returns (address) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        return $._cmAccountInfo[account].creator;
    }

    /**
     * @dev Check if an address is CMAccount created by the manager
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
     * @dev Get the CMAccount implementation address
     */
    function getAccountImplementation() public view returns (address) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        return $._latestAccountImplementation;
    }

    /**
     * @dev Set a new CMAccount implementation address
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
     * @dev Get prefund amount
     */
    function getPrefundAmount() public view returns (uint256) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        return $._prefundAmount;
    }

    /**
     * @dev Set pre fund amount
     */
    function setPrefundAmount(uint256 newPrefundAmount) public onlyRole(PREFUND_ADMIN_ROLE) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        $._prefundAmount = newPrefundAmount;
    }

    /***************************************************
     *                  BOOKING TOKEN                  *
     ***************************************************/

    /**
     * @dev Get booking token address
     */
    function getBookingTokenAddress() public view returns (address) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        return $._bookingToken;
    }

    /**
     * @dev Set booking token address
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
     * @dev Return developer wallet address
     */
    function getDeveloperWallet() public view returns (address developerWallet) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        return $._developerWallet;
    }

    /**
     * @dev Set developer wallet address
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
     * @dev Return developer fee in basis points
     */
    function getDeveloperFeeBp() public view returns (uint256 developerFeeBp) {
        CMAccountManagerStorage storage $ = _getCMAccountManagerStorage();
        return $._developerFeeBp;
    }

    /**
     * @dev Set developer fee in basis points
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

    function registerService(string memory serviceName) public onlyRole(SERVICE_REGISTRY_ADMIN_ROLE) {
        _registerServiceName(serviceName);
    }

    function unregisterService(string memory serviceName) public onlyRole(SERVICE_REGISTRY_ADMIN_ROLE) {
        _unregisterServiceName(serviceName);
    }
}
