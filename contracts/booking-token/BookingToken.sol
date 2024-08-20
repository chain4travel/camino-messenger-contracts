// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// UUPS Proxy
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

// ERC721
import "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol";

// Access
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

// Manager Interface
import { ICMAccountManager } from "../manager/ICMAccountManager.sol";

// Utils
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract BookingToken is
    Initializable,
    ERC721Upgradeable,
    ERC721EnumerableUpgradeable,
    ERC721URIStorageUpgradeable,
    AccessControlUpgradeable,
    UUPSUpgradeable
{
    using Address for address payable;
    using SafeERC20 for IERC20;

    /***************************************************
     *                   CONSTANTS                     *
     ***************************************************/

    /**
     * @dev Roles
     */
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    bytes32 public constant MIN_EXPIRATION_ADMIN_ROLE = keccak256("MIN_EXPIRATION_ADMIN_ROLE");

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    // Reservation details
    struct TokenReservation {
        address reservedFor; // CM Account address that can buy the token
        address supplier; // CM Account address that minted the token and created the reservation
        uint256 expirationTimestamp; // Timestamp when the reservation expires
        uint256 price; // Price of the token, only native for now
        IERC20 paymentToken; // Token used to pay for the reserved token
    }

    struct BookingTokenStorage {
        // CMAccountManager address
        address _manager;
        // Counter for generating unique token IDs
        uint256 _nextTokenId;
        // Mininum allowed expiration timestamp difference
        uint256 _minExpirationTimestampDiff;
        // Reservation details for each token
        mapping(uint256 tokenId => TokenReservation tokenReservation) _reservations;
    }

    // keccak256(abi.encode(uint256(keccak256("camino.messenger.storage.BookingToken")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant BookingTokenStorageLocation =
        0x9db9d405bf15683ce835607b1f0b423dc1484d44bb9d5af64a483fa4afd82900;

    function _getBookingTokenStorage() private pure returns (BookingTokenStorage storage $) {
        assembly {
            $.slot := BookingTokenStorageLocation
        }
    }

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    // Events for logging significant actions
    event TokenReserved(
        uint256 indexed tokenId,
        address indexed reservedFor,
        address indexed supplier,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken
    );

    // Reserved token bought
    event TokenBought(uint256 indexed tokenId, address indexed buyer);

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    /**
     * @dev Error for expiration timestamp too soon. It must be at least
       {_minExpirationTimestampDiff} seconds in the future
     */
    error ExpirationTimestampTooSoon(uint256 expirationTimestamp, uint256 minExpirationTimestampDiff);

    /**
     * @dev Address is not a CM Account
     */
    error NotCMAccount(address account);

    /**
     * @dev ReservedFor and buyer mismatch
     */
    error ReservationMismatch(address reservedFor, address buyer);

    /**
     * @dev Reservation expired
     */
    error ReservationExpired(uint256 tokenId, uint256 expirationTimestamp);

    /**
     * @dev Incorrect price
     */
    error IncorrectPrice(uint256 price, uint256 reservationPrice);

    /**
     * @dev Supplier is not the owner
     */
    error SupplierIsNotOwner(uint256 tokenId, address supplier);

    /**
     * @dev Token is reserved and can not be transferred
     */
    error TokenIsReserved(uint256 tokenId, address reservedFor);

    /**
     * @dev Insufficient allowance to transfer the ERC20 token to the supplier
     */
    error InsufficientAllowance(address sender, IERC20 paymentToken, uint256 price, uint256 allowance);

    /***************************************************
     *                  MODIFIERS                      *
     ***************************************************/

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

    // Function to authorize an upgrade for UUPS proxy
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(UPGRADER_ROLE) {}

    /***************************************************
     *             BOOKING-TOKEN LOGIC                 *
     ***************************************************/

    /**
     * @dev Mints a new token with a reservation for a specific address
     *
     * @param reservedFor The CM Account address that can buy the token
     * @param uri The URI of the token
     * @param expirationTimestamp The expiration timestamp
     * @param price The price of the token
     * @param paymentToken The token used to pay for the reservation. If address(0) then native.
     */
    function safeMintWithReservation(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken
    ) public onlyCMAccount(msg.sender) {
        // Require reservedFor to be a CM Account
        requireCMAccount(reservedFor);

        BookingTokenStorage storage $ = _getBookingTokenStorage();

        // Expiration timestamp should be at least _minExpirationTimestampDiff seconds in the future
        uint256 minExpirationTimestampDiff = $._minExpirationTimestampDiff;
        if (!(expirationTimestamp > (block.timestamp + minExpirationTimestampDiff))) {
            revert ExpirationTimestampTooSoon(expirationTimestamp, minExpirationTimestampDiff);
        }

        // Increment the token id
        uint256 tokenId = $._nextTokenId++;

        // Mint the token for the supplier (the caller)
        _safeMint(msg.sender, tokenId);
        _setTokenURI(tokenId, uri);

        // Store the reservation
        _reserve(tokenId, reservedFor, msg.sender, expirationTimestamp, price, paymentToken);

        emit TokenReserved(tokenId, reservedFor, msg.sender, expirationTimestamp, price, paymentToken);
    }

    function buyReservedToken(uint256 tokenId) external payable onlyCMAccount(msg.sender) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();

        // Get the reservation for the token
        TokenReservation memory reservation = $._reservations[tokenId];

        // Check reservationedFor and msg.sender match
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

        if (address(reservation.paymentToken) == address(0)) {
            // Payment is in native currency
            // Check if we receive the right price
            if (msg.value != reservation.price) {
                revert IncorrectPrice(msg.value, reservation.price);
            }

            // Transfer payment to the supplier
            payable(reservation.supplier).sendValue(msg.value);
        } else {
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
        }

        // Transfer the token. We are using `_transfer` instead of
        // `safeTransferFrom` because this is special transfer without a auth check.
        // Only in this function and only for buying a reserved token
        _transfer(reservation.supplier, msg.sender, tokenId);

        // Delete the reservation
        delete $._reservations[tokenId];

        // Emit event
        emit TokenBought(tokenId, msg.sender);
    }

    /**
     * @dev Reserve a token for a specific address with an expiration timestamp
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
     * @dev Check if the token is transferable
     */
    function checkTransferable(uint256 tokenId) internal {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        TokenReservation memory reservation = $._reservations[tokenId];

        // If expiration time is in the past, token is transferable. Because it can
        // not be bought after expired.
        //
        // This is also true if expirationTimestamp is 0. Which means there is no
        // reservation and token is transferable. No need to check for the
        // reservedFor address.
        if (block.timestamp <= reservation.expirationTimestamp) {
            // Token is not transferable
            revert TokenIsReserved(tokenId, reservation.reservedFor);
        } else if (reservation.reservedFor != address(0)) {
            // Clean up: Token is transferable but has expired reservation
            delete $._reservations[tokenId];
        }
    }

    /**
     * @dev Check if an address is a CM Account
     */
    function isCMAccount(address account) public view returns (bool) {
        return ICMAccountManager(getManagerAddress()).isCMAccount(account);
    }

    /**
     * @dev Check if the address is a CM Account and revert if not
     */
    function requireCMAccount(address account) internal view {
        if (!isCMAccount(account)) {
            revert NotCMAccount(account);
        }
    }

    /**
     * @dev Setter for _manager
     */
    function setManagerAddress(address manager) public onlyRole(DEFAULT_ADMIN_ROLE) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        $._manager = manager;
    }

    /**
     * @dev Getter for _manager
     */
    function getManagerAddress() public view returns (address) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        return $._manager;
    }

    /**
     * @dev Setter for _minExpirationTimestampDiff
     */
    function setMinExpirationTimestampDiff(
        uint256 minExpirationTimestampDiff
    ) public onlyRole(MIN_EXPIRATION_ADMIN_ROLE) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        $._minExpirationTimestampDiff = minExpirationTimestampDiff;
    }

    /**
     * @dev Getter for _minExpirationTimestampDiff
     */
    function getMinExpirationTimestampDiff() public view returns (uint256) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        return $._minExpirationTimestampDiff;
    }

    /**
     * @dev Get token reservation price for a specific token
     */
    function getReservationPrice(uint256 tokenId) public view returns (uint256 price, IERC20 paymentToken) {
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        return ($._reservations[tokenId].price, $._reservations[tokenId].paymentToken);
    }

    /***************************************************
     *              TRANSFER OVERRIDES                 *
     ***************************************************/

    /**
     * @dev Override transferFrom to check if token is reserved
     */
    function transferFrom(address from, address to, uint256 tokenId) public override(ERC721Upgradeable, IERC721) {
        // Verify that the token is transferable (i.e. not reserved)
        checkTransferable(tokenId);
        super.transferFrom(from, to, tokenId);
    }

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
