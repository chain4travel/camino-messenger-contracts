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

// Utils
import "@openzeppelin/contracts/utils/Address.sol";

interface ICMAccountManager {
    function getAccountImplementation() external view returns (address);

    function getDeveloperFeeBp() external view returns (uint256);

    function getDeveloperWallet() external view returns (address);

    function isCMAccount(address account) external view returns (bool);
}

contract BookingToken is
    Initializable,
    ERC721Upgradeable,
    ERC721EnumerableUpgradeable,
    ERC721URIStorageUpgradeable,
    AccessControlUpgradeable,
    UUPSUpgradeable
{
    using Address for address payable;

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

    // CMAccountManager address
    address private _manager;

    // Counter for generating unique token IDs
    uint256 private _nextTokenId;

    // Mininum allowed expiration timestamp difference
    uint256 private _minExpirationTimestampDiff;

    // Reservation details
    struct TokenReservation {
        address reservedFor; // CM Account address that can buy the token
        address supplier; // CM Account address that minted the token and created the reservation
        uint256 expirationTimestamp; // Timestamp when the reservation expires
        uint256 price; // Price of the token, only native for now
    }

    // Reservation details for each token
    mapping(uint256 tokenId => TokenReservation tokenReservation) private _reservations;

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    // Events for logging significant actions
    event TokenReserved(
        uint256 indexed tokenId,
        address indexed reservedFor,
        address indexed supplier,
        uint256 expirationTimestamp,
        uint256 price
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

        _manager = manager;
        _minExpirationTimestampDiff = 60;
    }

    // Function to authorize an upgrade for UUPS proxy
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(UPGRADER_ROLE) {}

    /***************************************************
     *             BOOKING-TOKEN LOGIC                 *
     ***************************************************/

    /**
     * @dev Mints a new token with a reservation for a specific address with an expiration timestamp
     */
    function safeMintWithReservation(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price
    ) public onlyCMAccount(msg.sender) {
        // Require reservedFor to be a CM Account
        requireCMAccount(reservedFor);

        // Expiration timestamp should be at least _minExpirationTimestampDiff seconds in the future
        if (!(expirationTimestamp > (block.timestamp + _minExpirationTimestampDiff))) {
            revert ExpirationTimestampTooSoon(expirationTimestamp, _minExpirationTimestampDiff);
        }

        // Increment the token id
        uint256 tokenId = _nextTokenId++;

        // Mint the token for the supplier (the caller)
        _safeMint(msg.sender, tokenId);
        _setTokenURI(tokenId, uri);

        // Store the reservation
        _reserve(tokenId, reservedFor, msg.sender, expirationTimestamp, price);

        emit TokenReserved(tokenId, reservedFor, msg.sender, expirationTimestamp, price);
    }

    function buyReservedToken(uint256 tokenId) external payable onlyCMAccount(msg.sender) {
        // Get the reservation for the token
        TokenReservation memory reservation = _reservations[tokenId];

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

        // Check if we receive the right price
        if (msg.value != reservation.price) {
            revert IncorrectPrice(msg.value, reservation.price);
        }

        // Transfer payment to the supplier
        payable(reservation.supplier).sendValue(msg.value);

        // Transfer the token
        _transfer(reservation.supplier, msg.sender, tokenId);

        // Delete the reservation
        delete _reservations[tokenId];

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
        uint256 price
    ) internal {
        _reservations[tokenId] = TokenReservation(reservedFor, supplier, expirationTimestamp, price);
    }

    /**
     * @dev Check if the token is transferable
     */
    function checkTransferable(uint256 tokenId) internal {
        TokenReservation memory reservation = _reservations[tokenId];

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
            delete _reservations[tokenId];
        }
    }

    /**
     * @dev Check if an address is a CM Account
     */
    function isCMAccount(address account) public view returns (bool) {
        return ICMAccountManager(_manager).isCMAccount(account);
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
        _manager = manager;
    }

    /**
     * @dev Getter for _manager
     */
    function getManagerAddress() public view returns (address) {
        return _manager;
    }

    /**
     * @dev Setter for _minExpirationTimestampDiff
     */
    function setMinExpirationTimestampDiff(
        uint256 minExpirationTimestampDiff
    ) public onlyRole(MIN_EXPIRATION_ADMIN_ROLE) {
        _minExpirationTimestampDiff = minExpirationTimestampDiff;
    }

    /**
     * @dev Getter for _minExpirationTimestampDiff
     */
    function getMinExpirationTimestampDiff() public view returns (uint256) {
        return _minExpirationTimestampDiff;
    }

    /**
     * @dev Get token reservation price for a specific token
     */
    function getReservationPrice(uint256 tokenId) public view returns (uint256) {
        return _reservations[tokenId].price;
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
