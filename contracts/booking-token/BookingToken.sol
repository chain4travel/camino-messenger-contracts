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

contract BookingToken is
    Initializable,
    ERC721Upgradeable,
    ERC721EnumerableUpgradeable,
    ERC721URIStorageUpgradeable,
    AccessControlUpgradeable,
    UUPSUpgradeable
{
    /***************************************************
     *                   CONSTANTS                     *
     ***************************************************/

    /**
     * @dev Roles
     */
    // Define a role identifier for the minter role
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    // Counter for generating unique token IDs
    uint256 private _nextTokenId;

    // Mapping to store supplier names
    mapping(address => string) private _supplierNames;

    // Mapping to store reserved addresses for specific tokens
    mapping(uint256 => address) private _reservedFor;

    // Mapping to store suppliers of tokens
    mapping(uint256 => address) private _tokenSuppliers;

    // Mapping to store expiration timestamps for reservations
    mapping(uint256 => uint256) private _expirationTimestamps;

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    // Events for logging significant actions
    event SupplierRegistered(address indexed supplier, string supplierName);
    event TokenReservation(address indexed reservedFor, uint256 indexed tokenId, uint256 expirationTimestamp);
    event TokenBought(uint256 indexed tokenId, address indexed buyer);

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    /***************************************************
     *                    FUNCS                        *
     ***************************************************/

    function initialize(address defaultAdmin, address minter, address upgrader) public initializer {
        __ERC721_init("BookingToken", "BookingToken");
        __ERC721Enumerable_init();
        __ERC721URIStorage_init();
        __AccessControl_init();
        __UUPSUpgradeable_init();

        _grantRole(DEFAULT_ADMIN_ROLE, defaultAdmin);
        _grantRole(MINTER_ROLE, minter);
        _grantRole(UPGRADER_ROLE, upgrader);
    }

    // Function to authorize an upgrade for UUPS proxy
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(UPGRADER_ROLE) {}

    /***************************************************
     *             BOOKING-TOKEN LOGIC                 *
     ***************************************************/

    // Function to mint a new token with a reservation for a specific address and expiration timestamp
    function safeMint(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp
    ) public onlyRole(MINTER_ROLE) {
        require(
            expirationTimestamp > block.timestamp + 60,
            "BookingToken: Expiration timestamp must be at least 60 seconds in the future"
        );

        uint256 tokenId = _nextTokenId++;
        _safeMint(msg.sender, tokenId);
        _setTokenURI(tokenId, uri);
        _reservedFor[tokenId] = reservedFor;
        _tokenSuppliers[tokenId] = msg.sender;
        _expirationTimestamps[tokenId] = expirationTimestamp;

        emit TokenReservation(reservedFor, tokenId, expirationTimestamp);
    }

    // Function to register a supplier with a name
    function registerSupplier(string memory supplierName) public {
        require(bytes(supplierName).length > 0, "BookingToken: Supplier name cannot be empty");
        _grantRole(MINTER_ROLE, msg.sender);
        _supplierNames[msg.sender] = supplierName;
        emit SupplierRegistered(msg.sender, supplierName);
    }

    // Function to allow a reserved address to buy the token
    function buy(uint256 tokenId) public {
        require(_reservedFor[tokenId] == msg.sender, "BookingToken: You do not have a reservation for this token");
        require(block.timestamp < _expirationTimestamps[tokenId], "BookingToken: Token reservation has expired");
        address owner = ownerOf(tokenId);
        _transfer(owner, msg.sender, tokenId);
        delete _reservedFor[tokenId];
        delete _expirationTimestamps[tokenId];
        emit TokenBought(tokenId, msg.sender);
    }

    // Function to get the supplier name of a given address
    function getSupplierName(address supplier) public view returns (string memory) {
        return _supplierNames[supplier];
    }

    // Function to get the supplier of a given token ID
    function getTokenSupplier(uint256 tokenId) public view returns (address) {
        return _tokenSuppliers[tokenId];
    }

    /***************************************************
     *            END BOOKING-TOKEN LOGIC              *
     ***************************************************/

    // The following functions are overrides required by Solidity.

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
