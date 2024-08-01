// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Account Implementation

pragma solidity ^0.8.24;

// UUPS Proxy
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

// Access
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

// Cheques
import "./ChequeManager.sol";

// Booking Token
import "../booking-token/BookingTokenOperator.sol";
import { IERC721Receiver } from "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

interface ICMAccountManager {
    function getAccountImplementation() external view returns (address);

    function getDeveloperFeeBp() external view returns (uint256);

    function getDeveloperWallet() external view returns (address);

    function isCMAccount(address account) external view returns (bool);
}

/**
 * @dev CM Account manages multiple bots for distributors and suppliers on Camino Messenger.
 *
 * This account holds funds that will be paid to the cheque beneficiaries.
 */
contract CMAccount is
    Initializable,
    PausableUpgradeable,
    AccessControlUpgradeable,
    UUPSUpgradeable,
    IERC721Receiver,
    ChequeManager,
    BookingTokenOperator
{
    using Address for address payable;

    /***************************************************
     *                    ROLES                        *
     ***************************************************/

    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    bytes32 public constant CHEQUE_OPERATOR_ROLE = keccak256("CHEQUE_OPERATOR_ROLE");
    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
    bytes32 public constant BOOKING_OPERATOR_ROLE = keccak256("BOOKING_OPERATOR_ROLE");

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    /**
     * @dev Address of the CMAccountManager
     */
    address private _manager;

    /**
     * @dev Address of the CMAccountManager
     */
    address private _bookingToken;

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    /**
     * @dev CMAccount upgrade event. Emitted when the CMAccount implementation is upgraded.
     */
    event CMAccountUpgraded(address indexed oldImplementation, address indexed newImplementation);

    /**
     * @dev Deposit event, emitted when there is a new deposit
     */
    event Deposit(address indexed sender, uint256 amount);

    /**
     * @dev Withdraw event, emitted when there is a new withdrawal
     */
    event Withdraw(address indexed receiver, uint256 amount);

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    /**
     * @dev CMAccount implementation address does not match the one in the manager
     */
    error CMAccountImplementationMismatch(address latestImplementation, address newImplementation);

    /**
     * @dev New implementation is the same as the current implementation, no update needed
     */
    error CMAccountNoUpgradeNeeded(address oldImplementation, address newImplementation);

    /**
     * @dev Error to revert with if depositer is not allowed
     */
    error DepositorNotAllowed(address sender);

    /**
     * @dev Error to revert zero value deposits
     */
    error ZeroValueDeposit(address sender);

    /***************************************************
     *         CONSTRUCTOR & INITIALIZATION            *
     ***************************************************/

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(
        address manager,
        address bookingToken,
        address defaultAdmin,
        address pauser,
        address upgrader
    ) public initializer {
        __Pausable_init();
        __AccessControl_init();
        __UUPSUpgradeable_init();
        __ChequeManager_init();

        _grantRole(DEFAULT_ADMIN_ROLE, defaultAdmin);
        _grantRole(PAUSER_ROLE, pauser);
        _grantRole(UPGRADER_ROLE, upgrader);

        _manager = manager;
        _bookingToken = bookingToken;
    }

    receive() external payable {}

    function getManagerAddress() public view returns (address) {
        return _manager;
    }

    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /**
     * @dev Upgrades the CMAccount implementation.
     *
     * Reverts if the new implementation is the same as the old one.
     *
     * Reverts if the new implementation does not match the implementation address in the manager.
     * Only implementations registered at the manager are allowed.
     *
     * Emits a {CMAccountUpgraded} event.
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(UPGRADER_ROLE) {
        // Get the implementation address from the manager
        address managerImplementation = ICMAccountManager(_manager).getAccountImplementation();
        address oldImplementation = ERC1967Utils.getImplementation();

        // Revert if the new implementation is the same as the old one
        if (oldImplementation == newImplementation) {
            revert CMAccountNoUpgradeNeeded(oldImplementation, newImplementation);
        }

        // Check if new implementation matches the implementation address in the manager
        if (newImplementation != managerImplementation) {
            revert CMAccountImplementationMismatch(managerImplementation, newImplementation);
        }

        emit CMAccountUpgraded(oldImplementation, newImplementation);
    }

    /**
     * @dev Returns true if an address is authorized bot
     */
    function isBotAllowed(address bot) public view override returns (bool) {
        return hasRole(CHEQUE_OPERATOR_ROLE, bot);
    }

    /**
     * @dev Return true if address is a registered CMAccount on the CMAccountManager
     */
    function isCMAccount(address account) internal view override returns (bool) {
        return ICMAccountManager(_manager).isCMAccount(account);
    }

    /**
     * @dev Return developer wallet address
     */
    function getDeveloperWallet() public view override returns (address) {
        address developerWallet = ICMAccountManager(_manager).getDeveloperWallet();
        return developerWallet;
    }

    /**
     * @dev Return developer fee in basis points
     */
    function getDeveloperFeeBp() public view override returns (uint256) {
        uint256 developerFeeBp = ICMAccountManager(_manager).getDeveloperFeeBp();
        return developerFeeBp;
    }

    /**
     * @dev Withdraw CAM from the CMAccount
     */
    function withdraw(address payable recipient, uint256 amount) public onlyRole(WITHDRAWER_ROLE) {
        emit Withdraw(recipient, amount);
        recipient.sendValue(amount);
    }

    /***************************************************
     *                 BOOKING TOKEN                   *
     ***************************************************/

    // TODO: Make sure the contract is able to use its booking tokens with
    // {IERC721-safeTransferFrom}, {IERC721-approve} or {IERC721-setApprovalForAll}.

    /**
     * @dev Mint booking token
     */
    function mintBookingToken(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price
    ) public override onlyRole(BOOKING_OPERATOR_ROLE) {
        // Mint the token
        _mintBookingToken(_bookingToken, reservedFor, uri, expirationTimestamp, price);
    }

    /**
     * @dev Buy booking token
     */
    function buyBookingToken(uint256 tokenId) external override onlyRole(BOOKING_OPERATOR_ROLE) {
        _buyBookingToken(_bookingToken, tokenId);
    }

    /**
     * @dev See {IERC721Receiver-onERC721Received}.
     *
     * Always returns `IERC721Receiver.onERC721Received.selector`.
     */
    function onERC721Received(address, address, uint256, bytes memory) public virtual returns (bytes4) {
        return this.onERC721Received.selector;
    }

    /**
     * @dev Get the price of a booking token
     */
    function getTokenReservationPrice(uint256 tokenId) public view returns (uint256) {
        return _getTokenReservationPrice(_bookingToken, tokenId);
    }
}
