// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

// UUPS Proxy
import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { UUPSUpgradeable } from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

// ERC721
import { ERC721Upgradeable, IERC721 } from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import { ERC721URIStorageUpgradeable } from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import { ERC721EnumerableUpgradeable } from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol";

// Access
import { AccessControlUpgradeable } from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

// Manager Interface
import { ICMAccountManager } from "../manager/ICMAccountManager.sol";

// Utils
import { Address } from "@openzeppelin/contracts/utils/Address.sol";
import { SafeERC20, IERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import { ReentrancyGuardUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";

/**
 * @title BookingToken
 * @notice Booking Token contract represents a booking done on the Camino Messenger.
 *
 * Suppliers can mint Booking Tokens and reserve them for a distributor address to
 * buy.
 *
 * Booking Tokens can have zero price, meaning that the payment will be done
 * off-chain.
 *
 * When a token is minted with a reservation, it can not be transferred until the
 * expiration timestamp is reached or the token is bought.
 */
contract BookingToken is
    Initializable,
    ERC721Upgradeable,
    ERC721EnumerableUpgradeable,
    ERC721URIStorageUpgradeable,
    AccessControlUpgradeable,
    ReentrancyGuardUpgradeable,
    UUPSUpgradeable
{
    using Address for address payable;
    using SafeERC20 for IERC20;

    /***************************************************
     *                   CONSTANTS                     *
     ***************************************************/

    /**
     * @notice Upgrader role can upgrade the contract to a new implementation.
     */
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

    /**
     * @notice This role can set the mininum allowed expiration timestamp difference.
     */
    bytes32 public constant MIN_EXPIRATION_ADMIN_ROLE = keccak256("MIN_EXPIRATION_ADMIN_ROLE");

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    enum BookingStatus {
        Unspecified, // 0, default value
        Reserved, // 1
        Expired, // 2
        Bought, // 3
        Cancelled // 4
    }

    // Reservation details
    struct TokenReservation {
        address reservedFor; // CM Account address that can buy the token
        address supplier; // CM Account address that minted the token and created the reservation
        uint256 expirationTimestamp; // Timestamp when the reservation expires
        uint256 price; // Price of the token, only native for now
        IERC20 paymentToken; // Token used to pay for the reserved token
    }

    /// @custom:storage-location erc7201:camino.messenger.storage.BookingToken
    struct BookingTokenStorage {
        // CMAccountManager address
        address _manager;
        // Counter for generating unique token IDs
        uint256 _nextTokenId;
        // Mininum allowed expiration timestamp difference
        uint256 _minExpirationTimestampDiff;
        // Reservation details for each token
        mapping(uint256 tokenId => TokenReservation tokenReservation) _reservations;
        // BookingStatus of each token
        mapping(uint256 tokenId => BookingStatus status) _bookingStatus;
    }

    // keccak256(abi.encode(uint256(keccak256("camino.messenger.storage.BookingToken")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant BookingTokenStorageLocation =
        0x9db9d405bf15683ce835607b1f0b423dc1484d44bb9d5af64a483fa4afd82900;

    function _getBookingTokenStorage() internal pure returns (BookingTokenStorage storage $) {
        assembly {
            $.slot := BookingTokenStorageLocation
        }
    }

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    /**
     * @notice Event emitted when a token is bought.
     *
     * @param tokenId token id
     * @param buyer buyer address
     */
    event TokenBought(uint256 indexed tokenId, address indexed buyer);

    /**
     * @notice Event emitted when a token is expired.
     *
     * @param tokenId token id
     */
    event TokenExpired(uint256 indexed tokenId);

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    /**
     * @notice Error for expiration timestamp too soon. It must be at least
     * `_minExpirationTimestampDiff` seconds in the future.
     */
    error ExpirationTimestampTooSoon(uint256 expirationTimestamp, uint256 minExpirationTimestampDiff);

    /**
     * @notice Address is not a CM Account.
     *
     * @param account account address
     */
    error NotCMAccount(address account);

    /**
     * @notice ReservedFor and buyer mismatch.
     *
     * @param reservedFor reserved for address
     * @param buyer buyer address
     */
    error ReservationMismatch(address reservedFor, address buyer);

    /**
     * @notice Reservation expired.
     *
     * @param tokenId token id
     * @param expirationTimestamp expiration timestamp
     */
    error ReservationExpired(uint256 tokenId, uint256 expirationTimestamp);

    /**
     * @notice Incorrect price.
     *
     * @param price price of the token
     * @param reservationPrice reservation price
     */
    error IncorrectPrice(uint256 price, uint256 reservationPrice);

    /**
     * @notice Supplier is not the owner.
     *
     * @param tokenId token id
     * @param supplier supplier address
     */
    error SupplierIsNotOwner(uint256 tokenId, address supplier);

    /**
     * @notice Token is reserved and can not be transferred.
     *
     * @param tokenId token id
     * @param reservedFor reserved for address
     */
    error TokenIsReserved(uint256 tokenId, address reservedFor);

    /**
     * @notice Insufficient allowance to transfer the ERC20 token to the supplier.
     *
     * @param sender msg.sender
     * @param paymentToken payment token address
     * @param price price of the token
     * @param allowance allowance amount
     */
    error InsufficientAllowance(address sender, IERC20 paymentToken, uint256 price, uint256 allowance);

    /**
     * @notice Invalid token status.
     *
     * @param tokenId token id
     * @param status status
     */
    error InvalidTokenStatus(uint256 tokenId, BookingStatus status);

    /***************************************************
     *                  MODIFIERS                      *
     ***************************************************/

    /**
     * @notice Only CMAccount modifier.
     */
    modifier onlyCMAccount(address account) {
        requireCMAccount(account);
        _;
    }

    /***************************************************
     *                    FUNCS                        *
     ***************************************************/

    function initialize(address manager, address defaultAdmin, address upgrader) public initializer {
        __ERC721_init("BookingToken", "TRIP");
        __ERC721Enumerable_init();
        __ERC721URIStorage_init();
        __AccessControl_init();
        __UUPSUpgradeable_init();

        _grantRole(DEFAULT_ADMIN_ROLE, defaultAdmin);
        _grantRole(UPGRADER_ROLE, upgrader);

        BookingTokenStorage storage $ = _getBookingTokenStorage();

        $._manager = manager;
        $._minExpirationTimestampDiff = 60;
    }

    /**
     * @notice Function to authorize an upgrade for UUPS proxy.
     */

    function _authorizeUpgrade(address newImplementation) internal override onlyRole(UPGRADER_ROLE) {}

    /***************************************************
     *             BOOKING-TOKEN LOGIC                 *
     ***************************************************/

    /**
     * @notice Buys a reserved token. The reservation must be for the message sender.
     *
     * Also the message sender should set allowance for the payment token to this
     * contract to at least the reservation price. (only for ERC20 tokens)
     *
     * For native coin, the message sender should send the exact amount.
     *
     * Only CM Accounts can call this function
     *
     * @param tokenId The token id
     */
    function buyReservedToken(uint256 tokenId) external payable nonReentrant onlyCMAccount(msg.sender) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();

        // Get the reservation for the token
        TokenReservation memory reservation = $._reservations[tokenId];

        // Check if `reservedFor` and `msg.sender` match
        if (reservation.reservedFor != msg.sender) {
            revert ReservationMismatch(reservation.reservedFor, msg.sender);
        }

        // Check expiration timestamp
        if (block.timestamp > reservation.expirationTimestamp) {
            revert ReservationExpired(tokenId, reservation.expirationTimestamp);
        }

        // Check if supplier is still the owner
        address owner = ownerOf(tokenId);
        if (owner != reservation.supplier) {
            revert SupplierIsNotOwner(tokenId, reservation.supplier);
        }

        // Transfer the token. We are using `_transfer` instead of
        // `safeTransferFrom` because this is special transfer without a auth check.
        // Only in this function and only for buying a reserved token
        _transfer(reservation.supplier, msg.sender, tokenId);

        // Do the payment at the end
        if (address(reservation.paymentToken) != address(0) && reservation.price > 0) {
            // Payment is in ERC20.
            //
            // Message sender (buyer of the Booking Token, generally the
            // distributor) must provide enough allowance for this (BookingToken)
            // contract to pay the reservation price for the token to the supplier.
            uint256 allowance = reservation.paymentToken.allowance(msg.sender, address(this));
            if (allowance < reservation.price) {
                revert InsufficientAllowance(msg.sender, reservation.paymentToken, reservation.price, allowance);
            }

            // Transfer the ERC20 tokens from buyer to supplier
            reservation.paymentToken.safeTransferFrom(msg.sender, reservation.supplier, reservation.price);
        } else {
            // Payment is in native currency
            // Check if we receive the right price
            if (msg.value != reservation.price) {
                revert IncorrectPrice(msg.value, reservation.price);
            }

            // Transfer payment to the supplier
            payable(reservation.supplier).sendValue(msg.value);
        }

        // Set the status
        $._bookingStatus[tokenId] = BookingStatus.Bought;

        // Emit event
        emit TokenBought(tokenId, msg.sender);
    }

    /**
     * @notice Return booking status
     *
     * @param tokenId The token id
     * @return The booking status
     */
    function getBookingStatus(uint256 tokenId) public view returns (BookingStatus) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        return $._bookingStatus[tokenId];
    }

    /**
     * @notice Reserve a token for a specific address with an expiration timestamp
     */
    function _reserve(
        uint256 tokenId,
        address reservedFor,
        address supplier,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken
    ) internal {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        $._reservations[tokenId] = TokenReservation(reservedFor, supplier, expirationTimestamp, price, paymentToken);
    }

    /**
     * @notice Check if the token is transferable
     */
    function checkTransferable(uint256 tokenId) internal virtual {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        BookingStatus status = $._bookingStatus[tokenId];

        // If token is bought, expired or cancelled, token is transferable
        if (status == BookingStatus.Bought || status == BookingStatus.Expired) {
            return;
        }

        // Revert if booking status is Cancelled
        if (status == BookingStatus.Cancelled) {
            revert InvalidTokenStatus(tokenId, status);
        }

        // If token is reserved, check if it is expired
        // If expiration time is in the past, token is transferable. Because it can
        // not be bought after expired.
        TokenReservation storage reservation = $._reservations[tokenId];

        if (block.timestamp > reservation.expirationTimestamp) {
            // Token is expired, set status to expired
            $._bookingStatus[tokenId] = BookingStatus.Expired;

            // Emit event
            emit TokenExpired(tokenId);
            return;
        } else {
            // Token is not expired, revert transfer
            revert TokenIsReserved(tokenId, reservation.reservedFor);
        }
    }

    /**
     * @notice Record expiration status if the token is expired
     *
     * @param tokenId The token id
     */
    function recordExpiration(uint256 tokenId) public virtual {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        TokenReservation storage reservation = $._reservations[tokenId];
        BookingStatus status = $._bookingStatus[tokenId];

        // If token is already set as expired, bought or cancelled, revert.
        if (status == BookingStatus.Expired || status == BookingStatus.Bought || status == BookingStatus.Cancelled) {
            revert InvalidTokenStatus(tokenId, status);
        }

        // If expiration time is in the past, set status to expired
        if (block.timestamp > reservation.expirationTimestamp) {
            $._bookingStatus[tokenId] = BookingStatus.Expired;

            // Emit event
            emit TokenExpired(tokenId);
        } else {
            // Token is not expired, revert setting status
            revert TokenIsReserved(tokenId, reservation.reservedFor);
        }
    }

    /**
     * @notice Checks if an address is a CM Account.
     *
     * @param account The address to check
     * @return true if the address is a CM Account
     */
    function isCMAccount(address account) public view returns (bool) {
        return ICMAccountManager(getManagerAddress()).isCMAccount(account);
    }

    /**
     * @notice Checks if the address is a CM Account and reverts if not.
     *
     * @param account The address to check
     */
    function requireCMAccount(address account) internal view {
        if (!isCMAccount(account)) {
            revert NotCMAccount(account);
        }
    }

    /**
     * @notice Sets for the manager address.
     *
     * @param manager The address of the manager
     */
    function setManagerAddress(address manager) public onlyRole(DEFAULT_ADMIN_ROLE) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        $._manager = manager;
    }

    /**
     * @notice Returns for the manager address.
     */
    function getManagerAddress() public view returns (address) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        return $._manager;
    }

    /**
     * @notice Sets minimum expiration timestamp difference in seconds.
     *
     * @param minExpirationTimestampDiff Minimum expiration timestamp difference in seconds
     */
    function setMinExpirationTimestampDiff(
        uint256 minExpirationTimestampDiff
    ) public onlyRole(MIN_EXPIRATION_ADMIN_ROLE) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        $._minExpirationTimestampDiff = minExpirationTimestampDiff;
    }

    /**
     * @notice Returns minimum expiration timestamp difference in seconds.
     */
    function getMinExpirationTimestampDiff() public view returns (uint256) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        return $._minExpirationTimestampDiff;
    }

    /**
     * @notice Returns the token reservation price for a specific token.
     *
     * @param tokenId The token id
     */
    function getReservationPrice(uint256 tokenId) public view returns (uint256 price, IERC20 paymentToken) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        return ($._reservations[tokenId].price, $._reservations[tokenId].paymentToken);
    }

    /***************************************************
     *              TRANSFER OVERRIDES                 *
     ***************************************************/

    /**
     * @notice Override transferFrom to check if token is reserved. It reverts if
     * the token is reserved.
     */
    function transferFrom(address from, address to, uint256 tokenId) public override(ERC721Upgradeable, IERC721) {
        // Verify that the token is transferable (i.e. not reserved)
        checkTransferable(tokenId);
        super.transferFrom(from, to, tokenId);
    }

    /**
     * @notice Override safeTransferFrom to check if token is reserved. It reverts if
     * the token is reserved.
     */
    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId,
        bytes memory data
    ) public override(ERC721Upgradeable, IERC721) {
        // Verify that the token is transferable (i.e. not reserved)
        checkTransferable(tokenId);
        super.safeTransferFrom(from, to, tokenId, data);
    }

    /***************************************************
     *            END BOOKING-TOKEN LOGIC              *
     ***************************************************/

    // Overrides required by Solidity.

    function _update(
        address to,
        uint256 tokenId,
        address auth
    ) internal override(ERC721Upgradeable, ERC721EnumerableUpgradeable) returns (address) {
        return super._update(to, tokenId, auth);
    }

    function _increaseBalance(
        address account,
        uint128 value
    ) internal override(ERC721Upgradeable, ERC721EnumerableUpgradeable) {
        super._increaseBalance(account, value);
    }

    function tokenURI(
        uint256 tokenId
    ) public view override(ERC721Upgradeable, ERC721URIStorageUpgradeable) returns (string memory) {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(
        bytes4 interfaceId
    )
        public
        view
        override(ERC721Upgradeable, ERC721EnumerableUpgradeable, ERC721URIStorageUpgradeable, AccessControlUpgradeable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}
