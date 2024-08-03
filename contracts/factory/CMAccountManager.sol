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
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";

// Utils
import "@openzeppelin/contracts/utils/Address.sol";

// ABI of the CMAccount implementation contract
interface ICMAccount {
    function initialize(
        address manager,
        address bookingToken,
        uint256 prefundAmount,
        address owner,
        address upgrader
    ) external;
}

/// @custom:security-contact https://r.xyz/program/camino-network
contract CMAccountManager is
    Initializable,
    PausableUpgradeable,
    AccessControlUpgradeable,
    UUPSUpgradeable,
    ReentrancyGuardUpgradeable
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

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    /**
     * @dev CM Account implementation address to be used by the CMAccount contract to resctrict
     * the implementation address for the UUPS proxies.
     */
    address internal _latestAccountImplementation;

    /**
     * @dev Prefund amount
     */
    uint256 private _prefundAmount;

    /**
     * @dev Developer wallet address. CMAccount sends the developer fee to this address.
     */
    address private _developerWallet;

    /**
     * @dev Developer fee basis points
     *
     * A basis point (bp) is one hundredth of 1 percentage point.
     *
     * 1 bp = 0.01%, 1/10,000⁠, or 0.0001.
     * 10 bp = 0.1%, 1/1,000⁠, or 0.001.
     * 100 bp = 1%, ⁠1/100⁠, or 0.01.
     */
    uint256 private _developerFeeBp;

    /**
     * @dev BookingToken address
     */
    address private _bookingToken;

    /**
     * @dev CMAccount info
     */
    struct CMAccountInfo {
        bool isCMAccount;
        address creator;
    }

    /**
     * @dev CMAccount info mapping to track if an address is a CMAccount and initial creators
     */
    mapping(address account => CMAccountInfo) internal cmAccountInfo;

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
        address upgrader,
        address versioner,
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

        _developerWallet = developerWallet;
        _developerFeeBp = developerFeeBp;

        // Set initial prefund amount
        _prefundAmount = 100 ether;
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

    /**
     * @dev Set a new CMAccount implementation address
     */
    function setAccountImplementation(address newImplementation) public onlyRole(VERSIONER_ROLE) {
        if (newImplementation.code.length == 0) {
            revert CMAccountInvalidImplementation(newImplementation);
        }

        emit CMAccountImplementationUpdated(_latestAccountImplementation, newImplementation);
        _latestAccountImplementation = newImplementation;
    }

    /**
     * @dev Get the CMAccount implementation address
     */
    function getAccountImplementation() public view returns (address) {
        return _latestAccountImplementation;
    }

    /**
     * @dev Set booking token address
     */
    function setBookingToken(address token) public onlyRole(VERSIONER_ROLE) {
        if (token.code.length == 0) {
            revert InvalidBookingTokenAddress(token);
        }
        emit BookingTokenAddressUpdated(_bookingToken, token);
        _bookingToken = token;
    }

    /**
     * @dev Get booking token address
     */
    function getBookingToken() public view returns (address) {
        return _bookingToken;
    }

    /**
     * @dev Set pre fund amount
     */
    function setPrefundAmount(uint256 newPrefundAmount) public onlyRole(PREFUND_ADMIN_ROLE) {
        _prefundAmount = newPrefundAmount;
    }

    /**
     * @dev Get pre fund amount
     */
    function getPrefundAmount() public view returns (uint256) {
        return _prefundAmount;
    }

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
        address latestAccountImplementation = _latestAccountImplementation;
        if (latestAccountImplementation.code.length == 0) {
            revert CMAccountInvalidImplementation(latestAccountImplementation);
        }
        if (_bookingToken.code.length == 0) {
            revert InvalidBookingTokenAddress(_bookingToken);
        }
        if (admin == address(0)) {
            revert CMAccountInvalidAdmin(admin);
        }

        uint256 prefundAmount = _prefundAmount;

        // Check pre-fund amount
        if (msg.value != prefundAmount) {
            revert IncorrectPrefundAmount(prefundAmount, msg.value);
        }

        // Create CMAccount Proxy and set the implementation address
        ERC1967Proxy cmAccountProxy = new ERC1967Proxy(latestAccountImplementation, "");

        // Set the isCMAccount and creator
        cmAccountInfo[address(cmAccountProxy)] = CMAccountInfo({ isCMAccount: true, creator: msg.sender });

        // Initialize the CMAccount
        ICMAccount(address(cmAccountProxy)).initialize(address(this), _bookingToken, prefundAmount, admin, upgrader);

        // Send the pre fund to the CMAccount
        payable(cmAccountProxy).sendValue(msg.value);

        emit CMAccountCreated(address(cmAccountProxy));

        return address(cmAccountProxy);
    }

    /**
     * @dev Check if an address is CMAccount created by the manager
     * @param account The account address to check
     */
    function isCMAccount(address account) public view returns (bool) {
        return cmAccountInfo[account].isCMAccount;
    }

    /**
     * @dev Return account's creator
     */
    function getCreator(address account) public view returns (address) {
        return cmAccountInfo[account].creator;
    }

    /**
     * @dev Return developer wallet address
     */
    function getDeveloperWallet() public view returns (address developerWallet) {
        return _developerWallet;
    }

    /**
     * @dev Set developer wallet address
     */
    function setDeveloperWallet(address developerWallet) public onlyRole(DEVELOPER_WALLET_ADMIN_ROLE) {
        if (developerWallet == address(0)) {
            revert InvalidDeveloperWallet(developerWallet);
        }
        emit DeveloperWalletUpdated(_developerWallet, developerWallet);
        _developerWallet = developerWallet;
    }

    /**
     * @dev Return developer fee in basis points
     */
    function getDeveloperFeeBp() public view returns (uint256) {
        return _developerFeeBp;
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
        emit DeveloperFeeBpUpdated(_developerFeeBp, bp);
        _developerFeeBp = bp;
    }
}
