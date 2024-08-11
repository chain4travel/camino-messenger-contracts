// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Account Implementation

pragma solidity ^0.8.24;

// UUPS Proxy
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

// Access
import "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlEnumerableUpgradeable.sol";

// Cheques
import "./ChequeManager.sol";

// Booking Token
import "../booking-token/BookingTokenOperator.sol";
import { IERC721Receiver } from "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

// Partner Config
import "../partner/PartnerConfiguration.sol";

interface ICMAccountManager {
    function getAccountImplementation() external view returns (address);

    function getDeveloperFeeBp() external view returns (uint256);

    function getDeveloperWallet() external view returns (address);

    function isCMAccount(address account) external view returns (bool);

    function getRegisteredServiceHashByName(string memory serviceName) external view returns (bytes32 serviceHash);

    function getRegisteredServiceNameByHash(bytes32 serviceHash) external view returns (string memory serviceName);
}

/**
 * @dev CM Account manages multiple bots for distributors and suppliers on Camino Messenger.
 *
 * This account holds funds that will be paid to the cheque beneficiaries.
 */
contract CMAccount is
    Initializable,
    AccessControlEnumerableUpgradeable,
    UUPSUpgradeable,
    IERC721Receiver,
    ChequeManager,
    BookingTokenOperator,
    PartnerConfiguration
{
    using Address for address payable;

    /***************************************************
     *                    ROLES                        *
     ***************************************************/

    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    bytes32 public constant CHEQUE_OPERATOR_ROLE = keccak256("CHEQUE_OPERATOR_ROLE");
    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
    bytes32 public constant BOOKING_OPERATOR_ROLE = keccak256("BOOKING_OPERATOR_ROLE");
    bytes32 public constant SERVICE_ADMIN_ROLE = keccak256("SERVICE_ADMIN_ROLE");

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

    /**
     * @dev Prefund amount
     */
    uint256 private _prefundAmount;

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

    /**
     * @dev Error to revert with if the prefund is not spent yet
     */
    error PrefundNotSpentYet(uint256 withdrawableAmount, uint256 prefundLeft, uint256 amount);

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
        uint256 prefundAmount,
        address defaultAdmin,
        address upgrader
    ) public initializer {
        __AccessControl_init();
        __UUPSUpgradeable_init();
        __ChequeManager_init();

        _grantRole(DEFAULT_ADMIN_ROLE, defaultAdmin);
        _grantRole(SERVICE_ADMIN_ROLE, defaultAdmin);
        _grantRole(UPGRADER_ROLE, upgrader);

        _manager = manager;
        _bookingToken = bookingToken;
        _prefundAmount = prefundAmount;
    }

    receive() external payable {}

    function getManagerAddress() public view returns (address) {
        return _manager;
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
     * @dev Verifies if the amount is withdrawable by checking if prefund is spent
     */
    function checkPrefundSpent(uint256 amount) public view {
        uint256 prefundAmount = _prefundAmount;
        uint256 totalChequePayments = getTotalChequePayments();

        // Check if prefund is spent. If total cheque payments is bigger or equal to
        // prefund amount it's ok to withdraw any amount
        if (totalChequePayments < prefundAmount) {
            // Balance should be bigger or equal to the { prefundLeft } because the
            // total sum of prefund is not yet spent. So, we substact that
            // (prefundLeft) from the balance to find the withdrawable amount.
            uint256 prefundLeft = prefundAmount - totalChequePayments;
            uint256 withdrawableAmount = address(this).balance - prefundLeft;

            // If amount is bigger than withdrawable amount, revert.
            // Otherwise, it's ok to withdraw the amount.
            if (amount > withdrawableAmount) {
                revert PrefundNotSpentYet(withdrawableAmount, prefundLeft, amount);
            }
        }
    }

    /**
     * @dev Withdraw CAM from the CMAccount
     *
     * This function reverts if the amount is bigger then the prefund left to spend. This is to prevent
     * spam by forcing user to spend the full prefund for cheques, so they can not just create an account
     * and withdraw the prefund.
     */
    function withdraw(address payable recipient, uint256 amount) public onlyRole(WITHDRAWER_ROLE) {
        // Check if amount is withdrawable according to the prefund spent amount
        checkPrefundSpent(amount);

        recipient.sendValue(amount);
        emit Withdraw(recipient, amount);
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

    /***************************************************
     *                PARTNER CONFIG                   *
     ***************************************************/

    /**
     * @dev Add a service to the account
     *
     * {serviceName} is defined as pkg + service name in protobuf. For example:
     *
     *  ┌────────────── pkg ─────────────┐ ┌───── service name ─────┐
     * "cmp.services.accommodation.v1alpha.AccommodationSearchService")
     *
     * @param serviceName Service name to add to the account as a supported service
     * @param fee Fee of the service in aCAM (wei in ETH terminology)
     * @param capabilities Capabilities of the service (if any, optional)
     */
    function addService(
        string memory serviceName,
        uint256 fee,
        string[] memory capabilities
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        // Check if the service is registered. This function reverts if the service is not registered
        bytes32 serviceHash = ICMAccountManager(_manager).getRegisteredServiceHashByName(serviceName);

        // Create the service object
        Service memory service = Service({ _fee: fee, _capabilities: capabilities });

        _addService(serviceHash, service);
    }

    /**
     * @dev Remove a service from the account by its hash
     */
    function removeService(bytes32 serviceHash) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removeService(serviceHash);
    }

    /**
     * @dev Remove a service from the account by its name
     */
    function removeService(string memory serviceName) public onlyRole(SERVICE_ADMIN_ROLE) {
        bytes32 serviceHash = ICMAccountManager(_manager).getRegisteredServiceHashByName(serviceName);
        _removeService(serviceHash);
    }

    // FEE

    /**
     * @dev Set the fee of a service by hash
     */
    function setServiceFee(bytes32 serviceHash, uint256 fee) public onlyRole(SERVICE_ADMIN_ROLE) {
        _setServiceFee(serviceHash, fee);
    }

    /**
     * @dev Set the fee of a service by name
     */
    function setServiceFee(string memory serviceName, uint256 fee) public onlyRole(SERVICE_ADMIN_ROLE) {
        bytes32 serviceHash = ICMAccountManager(_manager).getRegisteredServiceHashByName(serviceName);
        _setServiceFee(serviceHash, fee);
    }

    // ALL CAPABILITIES

    /**
     * @dev Set all capabilities for a service by hash
     */
    function setServiceCapabilities(
        bytes32 serviceHash,
        string[] memory capabilities
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        _setServiceCapabilities(serviceHash, capabilities);
    }

    /**
     * @dev Set all capabilities for a service by name
     */
    function setServiceCapabilities(
        string memory serviceName,
        string[] memory capabilities
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        bytes32 serviceHash = ICMAccountManager(_manager).getRegisteredServiceHashByName(serviceName);
        _setServiceCapabilities(serviceHash, capabilities);
    }

    // SINGLE CAPABILITY

    /**
     * @dev Add a single capability to the service by hash
     */
    function addServiceCapability(bytes32 serviceHash, string memory capability) public onlyRole(SERVICE_ADMIN_ROLE) {
        _addServiceCapability(serviceHash, capability);
    }

    /**
     * @dev Add a single capability to the service by name
     */
    function addServiceCapability(
        string memory serviceName,
        string memory capability
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        bytes32 serviceHash = ICMAccountManager(_manager).getRegisteredServiceHashByName(serviceName);
        _addServiceCapability(serviceHash, capability);
    }

    /**
     * @dev Remove a single capability from the service by hash
     */
    function removeServiceCapability(
        bytes32 serviceHash,
        string memory capability
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removeServiceCapability(serviceHash, capability);
    }

    /**
     * @dev Remove a single capability from the service by name
     */
    function removeServiceCapability(
        string memory serviceName,
        string memory capability
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        bytes32 serviceHash = ICMAccountManager(_manager).getRegisteredServiceHashByName(serviceName);
        _removeServiceCapability(serviceHash, capability);
    }

    // SERVICES WITH RESOLVED NAMES

    /**
     * @dev Get all supported services. Return a list of service names and a list of service objects.
     */
    function getSupportedServices() public view returns (string[] memory serviceNames, Service[] memory services) {
        // Get all hashes and create a list with predefined length
        bytes32[] memory _serviceHashes = getAllServiceHashes();
        string[] memory _serviceNames = new string[](_serviceHashes.length);
        Service[] memory _allSupportedServicesList = new Service[](_serviceHashes.length);

        for (uint256 i = 0; i < _serviceHashes.length; i++) {
            _serviceNames[i] = ICMAccountManager(_manager).getRegisteredServiceNameByHash(_serviceHashes[i]);
            _allSupportedServicesList[i] = getService(_serviceHashes[i]);
        }

        return (_serviceNames, _allSupportedServicesList);
    }

    /**
     * @dev Get service fee by name. Overloading the getServiceFee function.
     */
    function getServiceFeeByName(string memory serviceName) public view returns (uint256 fee) {
        bytes32 serviceHash = ICMAccountManager(_manager).getRegisteredServiceHashByName(serviceName);
        return getServiceFee(serviceHash);
    }

    /**
     * @dev Get service capabilities by name. Overloading the getServiceCapabilities function.
     */
    function getServiceCapabilitiesByName(
        string memory serviceName
    ) public view returns (string[] memory capabilities) {
        bytes32 serviceHash = ICMAccountManager(_manager).getRegisteredServiceHashByName(serviceName);
        return getServiceCapabilities(serviceHash);
    }

    /***************************************************
     *                   PAYMENT                       *
     ***************************************************/

    function setOffChainPaymentSupported(bool _isSupported) public onlyRole(SERVICE_ADMIN_ROLE) {
        _setOffChainPaymentSupported(_isSupported);
    }

    function addSupportedToken(address _supportedToken) public onlyRole(SERVICE_ADMIN_ROLE) {
        _addSupportedToken(_supportedToken);
    }

    function removeSupportedToken(address _supportedToken) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removeSupportedToken(_supportedToken);
    }

    /***************************************************
     *                  PUBLIC KEY                     *
     ***************************************************/

    /**
     * @dev Add public key with address
     *
     * These public keys are intended to be used with for off-chain encryption of private booking data.
     */
    function addPublicKey(address pubKeyAddress, bytes memory publicKey) public onlyRole(SERVICE_ADMIN_ROLE) {
        _addPublicKey(pubKeyAddress, publicKey);
    }

    /**
     * @dev Remove public key by address
     */
    function removePublicKey(address pubKeyAddress) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removePublicKey(pubKeyAddress);
    }
}
