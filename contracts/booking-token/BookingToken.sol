// SPDX-License-Identifier: UNLICENSED
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

// Cancellable
import { BookingTokenCancellable, CancellationProposalStatus } from "./BookingTokenCancellable.sol";

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
    UUPSUpgradeable,
    BookingTokenCancellable
{
    using Address for address payable;
    using SafeERC20 for IERC20;

    /***************************************************
     *                    VERSION                      *
     ***************************************************/

    uint16 constant VERSION_MAJOR = 1;
    uint16 constant VERSION_MINOR = 0;
    uint16 constant VERSION_PATCH = 0;

    /**
     * @notice Returns the semantic version of the contract.
     *
     * - no version() func: Legacy version without Cancellation support
     * - v1.0.0: Version with Cancellation support
     *
     * @return major Major version (breaking changes)
     * @return minor Minor version (backwards-compatible features)
     * @return patch Patch version (backwards-compatible fixes)
     */
    function version() external pure virtual returns (uint16 major, uint16 minor, uint16 patch) {
        return (VERSION_MAJOR, VERSION_MINOR, VERSION_PATCH);
    }

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

    /**
     * @dev Special address for native payments.
     * @notice Tokens are directly transferred to the recipient.
     */
    address public constant NATIVE_PAYMENT = address(0);

    /**
     * @dev Special address for offchain payments. The enum for this
     * is defined in the Camino Messenger Protocol's
     * cmp.types.<version>.IsoCurrency enum (currency.proto file).
     * @notice A third-party service is used to handle payments.
     */
    address public constant OFFCHAIN_PAYMENT = address(1);

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    enum BookingStatus {
        UNSPECIFIED, // 0, default value
        RESERVED, // 1
        RESERVATION_EXPIRED, // 2
        BOUGHT, // 3
        CANCELLED // 4
    }

    // Reservation details
    struct TokenReservation {
        address reservedFor; // CM Account address that can buy the token
        address supplier; // CM Account address that minted the token and created the reservation
        uint256 expirationTimestamp; // Timestamp when the reservation expires
        uint256 price; // Price of the token, only native for now
        IERC20 paymentToken; // Token used to pay for the reserved token
        uint256 offchainPaymentCurrency; // Offchain payment currency
        bool cancellable; // Is the token (booking) cancellable
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
     * @notice Event emitted when a token is reserved.
     *
     * @param tokenId token id
     * @param reservedFor reserved for address
     * @param supplier supplier address
     * @param expirationTimestamp expiration timestamp
     * @param price price of the token
     * @param paymentToken payment token address
     */
    event TokenReserved(
        uint256 indexed tokenId,
        address indexed reservedFor,
        address indexed supplier,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken,
        uint256 offchainPaymentCurrency,
        bool cancellable
    );

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
    event TokenReservationExpired(uint256 indexed tokenId);

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

    /**
     * @notice Unexpected offchain payment currency. Thrown when offchain payment currency is provided
     * but payment token is not address(1).
     *
     * @param offchainPaymentCurrency offchain payment currency
     */
    error UnexpectedOffchainPaymentCurrency(uint256 offchainPaymentCurrency);

    /**
     * @notice Error for when there is unexpected native payment.
     *
     * @param amount The unexpected amount
     */
    error UnexpectedNativePayment(uint256 amount);

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
     *                   INITIALIZE                    *
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

    /***************************************************
     *                  REINITIALIZE                   *
     ***************************************************/

    /**
     * @notice This function allows reinitializing the contract to update the name and symbol
     * @dev Only callable by DEFAULT_ADMIN_ROLE
     * @param newName New token name
     * @param newSymbol New token symbol
     */
    function reinitializeV2(
        string memory newName,
        string memory newSymbol
    ) public reinitializer(2) onlyRole(DEFAULT_ADMIN_ROLE) {
        __ERC721_init(newName, newSymbol);
    }

    /***************************************************
     *             BOOKING-TOKEN LOGIC                 *
     ***************************************************/

    /**
     * @notice Function to authorize an upgrade for UUPS proxy.
     */
    function _authorizeUpgrade(address newImplementation) internal virtual override onlyRole(UPGRADER_ROLE) {}

    /**
     * @notice Mints a new token with a reservation for a specific address.
     *
     * @param reservedFor The CM Account address that can buy the token
     * @param uri The URI of the token
     * @param expirationTimestamp The expiration timestamp
     * @param price The price of the token
     * @param paymentToken The token used to pay for the reservation. If address(0) then native.
     * @param offchainPaymentCurrency The offchain payment currency
     * @param cancellable The flag that represents whether the booking is cancellable
     */
    function safeMintWithReservation(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken,
        uint256 offchainPaymentCurrency,
        bool cancellable
    ) public virtual onlyCMAccount(msg.sender) {
        // Require reservedFor to be a CM Account
        requireCMAccount(reservedFor);

        BookingTokenStorage storage $ = _getBookingTokenStorage();

        // Expiration timestamp should be at least `_minExpirationTimestampDiff`
        // seconds in the future
        uint256 minExpirationTimestampDiff = $._minExpirationTimestampDiff;
        if (!(expirationTimestamp > (block.timestamp + minExpirationTimestampDiff))) {
            revert ExpirationTimestampTooSoon(expirationTimestamp, minExpirationTimestampDiff);
        }

        // Revert if the off chain payment currency is provided but payment token is not address(1)
        if (offchainPaymentCurrency > 0 && address(paymentToken) != OFFCHAIN_PAYMENT) {
            revert UnexpectedOffchainPaymentCurrency(offchainPaymentCurrency);
        }

        // Increment the token id
        uint256 tokenId = $._nextTokenId++;

        // Mint the token for the supplier (the caller)
        _safeMint(msg.sender, tokenId);
        _setTokenURI(tokenId, uri);

        // Store the reservation
        _reserve(
            tokenId,
            reservedFor,
            msg.sender,
            expirationTimestamp,
            price,
            paymentToken,
            offchainPaymentCurrency,
            cancellable
        );

        // Set the status
        $._bookingStatus[tokenId] = BookingStatus.RESERVED;

        emit TokenReserved(
            tokenId,
            reservedFor,
            msg.sender,
            expirationTimestamp,
            price,
            paymentToken,
            offchainPaymentCurrency,
            cancellable
        );
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
        IERC20 paymentToken,
        uint256 offchainPaymentCurrency,
        bool cancellable
    ) internal virtual {
        BookingTokenStorage storage $ = _getBookingTokenStorage();

        $._reservations[tokenId] = TokenReservation(
            reservedFor,
            supplier,
            expirationTimestamp,
            price,
            paymentToken,
            offchainPaymentCurrency,
            cancellable
        );
    }

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
    function buyReservedToken(uint256 tokenId) public payable virtual nonReentrant onlyCMAccount(msg.sender) {
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
        processPayment(reservation.paymentToken, reservation.price, reservation.supplier);

        // Set the status
        $._bookingStatus[tokenId] = BookingStatus.BOUGHT;

        // Emit event
        emit TokenBought(tokenId, msg.sender);
    }

    function processPayment(IERC20 paymentToken, uint256 paymentAmount, address recipient) internal virtual {
        // Handle the payment based on payment type
        if (address(paymentToken) == NATIVE_PAYMENT) {
            // Payment is in native currency (CAM)
            if (msg.value != paymentAmount) {
                revert IncorrectPrice(msg.value, paymentAmount);
            }

            // Transfer payment to the supplier
            payable(recipient).sendValue(msg.value);
        } else if (address(paymentToken) == OFFCHAIN_PAYMENT) {
            // Off-chain payment - no on-chain transfer needed
            // Ensure no native currency was sent
            if (msg.value > 0) {
                revert UnexpectedNativePayment(msg.value);
            }
        } else {
            // Payment is in ERC20
            // Ensure no native currency was sent
            if (msg.value > 0) {
                revert UnexpectedNativePayment(msg.value);
            }

            // Transfer the ERC20 tokens from distributor to supplier
            IERC20(paymentToken).safeTransferFrom(msg.sender, recipient, paymentAmount);
        }
    }

    /***************************************************
     *                 TOKEN GETTERS                   *
     ***************************************************/

    /**
     * @notice Return booking status
     *
     * @param tokenId The token id
     * @return The booking status
     */
    function getBookingStatus(uint256 tokenId) public view virtual returns (BookingStatus) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        return $._bookingStatus[tokenId];
    }

    /**
     * @notice Returns the token reservation price for a specific token.
     *
     * @param tokenId The token id
     */
    function getReservationPrice(uint256 tokenId) public view virtual returns (uint256 price, IERC20 paymentToken) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        return ($._reservations[tokenId].price, $._reservations[tokenId].paymentToken);
    }

    /**
     * @notice Retrieves the payment token for a given token.
     *
     * @param tokenId The token id to retrieve the payment token for
     * @return paymentToken The payment token
     */
    function getReservationPaymentToken(uint256 tokenId) external view returns (IERC20 paymentToken) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        return $._reservations[tokenId].paymentToken;
    }

    /**
     * @notice Returns if the token is cancellable
     *
     * @param tokenId The token id
     */
    function isCancellable(uint256 tokenId) public view virtual returns (bool) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        return $._reservations[tokenId].cancellable;
    }

    /***************************************************
     *                CONTRACT LOGIC                   *
     ***************************************************/

    /**
     * @notice Check if the token is transferable
     */
    function checkTransferable(uint256 tokenId) internal virtual {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        BookingStatus status = $._bookingStatus[tokenId];

        // Check Cancellation Proposal status

        // Get the current proposer and status
        (
            CancellationProposalStatus cancellationStatus,
            address currentProposer
        ) = _getCancellationProposalStatusAndCurrentProposer(tokenId);

        // If there is a pending cancellation proposal, withdraw or reject it
        // automatically before the transfer.
        if (cancellationStatus == CancellationProposalStatus.PENDING) {
            address owner = _requireOwned(tokenId);
            address supplier = $._reservations[tokenId].supplier;

            // Check if the current proposer is the owner
            if (msg.sender != currentProposer) {
                // FIXME: Define a reason in the  CMP and update this
                _rejectCancellation(owner, supplier, tokenId, 99, 1);
            } else {
                // FIXME: Define a reason in the  CMP and update this
                _withdrawCancellation(owner, supplier, tokenId, 99, 1);
            }
        }

        // If token is UNSPECIFIED, BOUGHT, or EXPIRED, token is transferable, return early.
        if (
            status == BookingStatus.BOUGHT ||
            status == BookingStatus.RESERVATION_EXPIRED ||
            status == BookingStatus.UNSPECIFIED
        ) {
            return;
        }

        // Revert if booking status is CANCELLED
        if (status == BookingStatus.CANCELLED) {
            revert InvalidTokenStatus(tokenId, status);
        }

        // Only RESERVED state is left. If expiration time is in the past, token is
        // transferable even if it is reserved. Because it can not be bought after
        // expired.
        TokenReservation storage reservation = $._reservations[tokenId];

        if (block.timestamp > reservation.expirationTimestamp) {
            // Token is expired, set status to expired
            $._bookingStatus[tokenId] = BookingStatus.RESERVATION_EXPIRED;

            // Emit event
            emit TokenReservationExpired(tokenId);
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
        if (
            status == BookingStatus.RESERVATION_EXPIRED ||
            status == BookingStatus.BOUGHT ||
            status == BookingStatus.CANCELLED
        ) {
            revert InvalidTokenStatus(tokenId, status);
        }

        // If expiration time is in the past, set status to expired
        if (block.timestamp > reservation.expirationTimestamp) {
            $._bookingStatus[tokenId] = BookingStatus.RESERVATION_EXPIRED;

            // Emit event
            emit TokenReservationExpired(tokenId);
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
    function isCMAccount(address account) public view virtual returns (bool) {
        return ICMAccountManager(getManagerAddress()).isCMAccount(account);
    }

    /**
     * @notice Checks if the address is a CM Account and reverts if not.
     *
     * @param account The address to check
     */
    function requireCMAccount(address account) internal view virtual {
        if (!isCMAccount(account)) {
            revert NotCMAccount(account);
        }
    }

    /**
     * @notice Sets for the manager address.
     *
     * @param manager The address of the manager
     */
    function setManagerAddress(address manager) public virtual onlyRole(DEFAULT_ADMIN_ROLE) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        $._manager = manager;
    }

    /**
     * @notice Returns for the manager address.
     */
    function getManagerAddress() public view virtual returns (address) {
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
    ) public virtual onlyRole(MIN_EXPIRATION_ADMIN_ROLE) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        $._minExpirationTimestampDiff = minExpirationTimestampDiff;
    }

    /**
     * @notice Returns minimum expiration timestamp difference in seconds.
     */
    function getMinExpirationTimestampDiff() public view virtual returns (uint256) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        return $._minExpirationTimestampDiff;
    }

    /***************************************************
     *              CANCELLATION LOGIC                 *
     ***************************************************/

    function initiateCancellation(
        uint256 tokenId,
        uint256 refundAmount,
        uint16 cancellationReason,
        uint16 cancellationReasonVersion
    ) external virtual onlyCMAccount(msg.sender) {
        // Revert if token does not exist
        address owner = _requireOwned(tokenId);

        // Get storage
        BookingTokenStorage storage $ = _getBookingTokenStorage();

        // Revert if token is not BOUGHT
        if ($._bookingStatus[tokenId] != BookingStatus.BOUGHT) {
            revert InvalidTokenStatus(tokenId, $._bookingStatus[tokenId]);
        }

        address supplier = $._reservations[tokenId].supplier;

        _initiateCancellation(owner, supplier, tokenId, refundAmount, cancellationReason, cancellationReasonVersion);
    }

    function acceptCancellation(uint256 tokenId, uint256 refundAmount) external virtual onlyCMAccount(msg.sender) {
        // Revert if token does not exist
        address owner = _requireOwned(tokenId);

        // Get storage
        BookingTokenStorage storage $ = _getBookingTokenStorage();

        // Revert if token is not BOUGHT
        if ($._bookingStatus[tokenId] != BookingStatus.BOUGHT) {
            revert InvalidTokenStatus(tokenId, $._bookingStatus[tokenId]);
        }

        address supplier = $._reservations[tokenId].supplier;

        _acceptCancellation(owner, supplier, tokenId, refundAmount);
    }

    function counterCancellation(
        uint256 tokenId,
        uint256 refundAmount,
        uint16 counterReason,
        uint16 counterReasonVersion
    ) external virtual onlyCMAccount(msg.sender) {
        // Revert if token does not exist
        address owner = _requireOwned(tokenId);

        // Get storage
        BookingTokenStorage storage $ = _getBookingTokenStorage();

        // Revert if token is not BOUGHT
        if ($._bookingStatus[tokenId] != BookingStatus.BOUGHT) {
            revert InvalidTokenStatus(tokenId, $._bookingStatus[tokenId]);
        }

        address supplier = $._reservations[tokenId].supplier;

        _counterCancellation(owner, supplier, tokenId, refundAmount, counterReason, counterReasonVersion);
    }

    function withdrawCancellation(
        uint256 tokenId,
        uint16 withdrawalReason,
        uint16 withdrawalReasonVersion
    ) external virtual onlyCMAccount(msg.sender) {
        // Revert if token does not exist
        address owner = _requireOwned(tokenId);

        // Get storage
        BookingTokenStorage storage $ = _getBookingTokenStorage();

        // Revert if token is not BOUGHT
        if ($._bookingStatus[tokenId] != BookingStatus.BOUGHT) {
            revert InvalidTokenStatus(tokenId, $._bookingStatus[tokenId]);
        }

        address supplier = $._reservations[tokenId].supplier;

        _withdrawCancellation(owner, supplier, tokenId, withdrawalReason, withdrawalReasonVersion);
    }

    function rejectCancellation(
        uint256 tokenId,
        uint16 rejectionReason,
        uint16 rejectionReasonVersion
    ) external virtual onlyCMAccount(msg.sender) {
        // Revert if token does not exist
        address owner = _requireOwned(tokenId);

        // Get storage
        BookingTokenStorage storage $ = _getBookingTokenStorage();

        // Revert if token is not BOUGHT
        if ($._bookingStatus[tokenId] != BookingStatus.BOUGHT) {
            revert InvalidTokenStatus(tokenId, $._bookingStatus[tokenId]);
        }

        address supplier = $._reservations[tokenId].supplier;

        _rejectCancellation(owner, supplier, tokenId, rejectionReason, rejectionReasonVersion);
    }

    function finalizeCancellation(
        uint256 tokenId,
        uint256 checkRefundAmount
    ) external payable virtual onlyCMAccount(msg.sender) {
        // Revert if token does not exist
        address owner = _requireOwned(tokenId);

        // Get storage
        BookingTokenStorage storage $ = _getBookingTokenStorage();

        // Revert if token is not BOUGHT
        if ($._bookingStatus[tokenId] != BookingStatus.BOUGHT) {
            revert InvalidTokenStatus(tokenId, $._bookingStatus[tokenId]);
        }

        address supplier = $._reservations[tokenId].supplier;

        uint256 refundAmount = _finalizeCancellation(supplier, tokenId, checkRefundAmount);

        IERC20 paymentToken = $._reservations[tokenId].paymentToken;

        // Update BookingToken status
        $._bookingStatus[tokenId] = BookingStatus.CANCELLED;

        // Process payment
        processPayment(paymentToken, refundAmount, owner);
    }

    /***************************************************
     *              TRANSFER OVERRIDES                 *
     ***************************************************/

    /**
     * @notice Override transferFrom to check if token is reserved. It reverts if
     * the token is reserved.
     */
    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public virtual override(ERC721Upgradeable, IERC721) {
        // Verify that the token is transferable (i.e. not reserved)
        checkTransferable(tokenId);
        super.transferFrom(from, to, tokenId);
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
