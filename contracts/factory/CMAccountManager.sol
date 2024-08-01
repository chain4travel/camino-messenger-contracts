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

// ABI of the CMAccount implementation contract
interface ICMAccount {
    function initialize(
        address manager,
        address bookingToken,
        address owner,
        address pauser,
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
    /**
     * @dev Roles
     */
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    bytes32 public constant VERSIONER_ROLE = keccak256("VERSIONER_ROLE");
    bytes32 public constant FEE_ADMIN_ROLE = keccak256("FEE_ADMIN_ROLE");
    bytes32 public constant DEVELOPER_WALLET_ADMIN_ROLE = keccak256("DEVELOPER_WALLET_ADMIN_ROLE");

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    /**
     * @dev CM Account implementation address to be used by the CMAccount contract to resctrict
     * the implementation address for the UUPS proxies.
     */
    address internal _latestAccountImplementation;

    address private _developerWallet;
    uint256 private _developerFeeBp;

    address private _bookingToken;

    /**
     * @dev CM Accounts mapping
     */
    mapping(address account => bool isCMAccount) internal cmAccounts;

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
    error BookingTokenAddressInvalid(address bookingToken);

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
            revert BookingTokenAddressInvalid(token);
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
     * @dev Creates CMAccount by deploying a ERC1967Proxy with the CMAccount implementation from the manager.
     *
     * Because this function is deploying a contract, it reverts if the caller is not KYC or KYB verified.
     *
     * Emits a {CMAccountCreated} event.
     */
    function createCMAccount(
        address admin,
        address pauser,
        address upgrader
    ) external nonReentrant whenNotPaused returns (address) {
        return _createCMAccount(admin, pauser, upgrader);
    }

    function _createCMAccount(address admin, address pauser, address upgrader) private returns (address) {
        // Checks
        address latestAccountImplementation = _latestAccountImplementation;
        if (latestAccountImplementation.code.length == 0) {
            revert CMAccountInvalidImplementation(latestAccountImplementation);
        }
        if (_bookingToken.code.length == 0) {
            revert BookingTokenAddressInvalid(_bookingToken);
        }
        if (admin == address(0)) {
            revert CMAccountInvalidAdmin(admin);
        }

        // Create CMAccount Proxy and set the implementation address
        ERC1967Proxy cmAccountProxy = new ERC1967Proxy(latestAccountImplementation, "");
        emit CMAccountCreated(address(cmAccountProxy));

        // Add to the known CMAccounts
        cmAccounts[address(cmAccountProxy)] = true;

        // Initialize the CMAccount
        ICMAccount(address(cmAccountProxy)).initialize(address(this), _bookingToken, admin, pauser, upgrader);

        return address(cmAccountProxy);
    }

    /**
     * @dev Check if an address is CMAccount created by the manager
     * @param account The account address to check
     */
    function isCMAccount(address account) public view returns (bool) {
        return cmAccounts[account];
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
     * 100 bp = 1%, 10−2, ⁠1/100⁠, or 0.01.
     */
    function setDeveloperFeeBp(uint256 bp) public onlyRole(FEE_ADMIN_ROLE) {
        emit DeveloperFeeBpUpdated(_developerFeeBp, bp);
        _developerFeeBp = bp;
    }
}
