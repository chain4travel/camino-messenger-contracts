// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Account Implementation

pragma solidity ^0.8.24;

// UUPS Proxy
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

// Access
// Size Impact: +0.411 (Enumerable)
import "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlEnumerableUpgradeable.sol";
//import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

// ERC721
import { IERC721 } from "@openzeppelin/contracts/token/ERC721/IERC721.sol";

// Manager Interface
import { ICMAccountManager } from "../manager/ICMAccountManager.sol";

// Cheques
import "./ChequeManager.sol";

// Booking Token
import "../booking-token/BookingTokenOperator.sol";
import { IERC721Receiver } from "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

// Partner Config
import "../partner/PartnerConfiguration.sol";

// Partner Config
import "./GasMoneyManager.sol";

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
    PartnerConfiguration,
    GasMoneyManager
{
    using Address for address payable;
    using SafeERC20 for IERC20;

    /***************************************************
     *                    ROLES                        *
     ***************************************************/

    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    bytes32 public constant BOT_ADMIN_ROLE = keccak256("BOT_ADMIN_ROLE");
    bytes32 public constant CHEQUE_OPERATOR_ROLE = keccak256("CHEQUE_OPERATOR_ROLE");
    bytes32 public constant GAS_WITHDRAWER_ROLE = keccak256("GAS_WITHDRAWER_ROLE");
    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
    bytes32 public constant BOOKING_OPERATOR_ROLE = keccak256("BOOKING_OPERATOR_ROLE");
    bytes32 public constant SERVICE_ADMIN_ROLE = keccak256("SERVICE_ADMIN_ROLE");

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    struct CMAccountStorage {
        /**
         * @dev Address of the CMAccountManager
         */
        address _manager;
        /**
         * @dev Address of the BookingToken contract
         */
        address _bookingToken;
        /**
         * @dev Prefund amount
         */
        uint256 _prefundAmount;
    }

    // keccak256(abi.encode(uint256(keccak256("camino.messenger.storage.CMAccount")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant CMAccountStorageLocation =
        0x0c7b73796c7cc89b9f849b9056a93200eba741881e57a1b03b9bedb2c0e07100;

    function _getCMAccountStorage() private pure returns (CMAccountStorage storage $) {
        assembly {
            $.slot := CMAccountStorageLocation
        }
    }

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

    /**
     * @dev Messenger bot added
     */
    event MessengerBotAdded(address indexed bot);

    /**
     * @dev Messenger bot removed
     */
    event MessengerBotRemoved(address indexed bot);

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

    /**
     * @dev Error to revert if transfer to zero address
     */
    error TransferToZeroAddress();

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
        _grantRole(BOT_ADMIN_ROLE, defaultAdmin);
        _grantRole(UPGRADER_ROLE, upgrader);

        CMAccountStorage storage $ = _getCMAccountStorage();

        $._manager = manager;
        $._bookingToken = bookingToken;
        $._prefundAmount = prefundAmount;

        // Initialize GasMoneyManager
        uint256 withdrawalLimit = 10 ether; // 10 CAM
        uint256 withdrawalPeriod = 24 hours; // per 24 hours
        __GasMoneyManager_init(withdrawalLimit, withdrawalPeriod);
    }

    receive() external payable {}

    /***************************************************
     *                    Getters                      *
     ***************************************************/

    function getManagerAddress() public view override returns (address) {
        CMAccountStorage storage $ = _getCMAccountStorage();
        return $._manager;
    }

    function getBookingTokenAddress() public view returns (address) {
        CMAccountStorage storage $ = _getCMAccountStorage();
        return $._bookingToken;
    }

    function getPrefundAmount() public view returns (uint256) {
        CMAccountStorage storage $ = _getCMAccountStorage();
        return $._prefundAmount;
    }

    /***************************************************
     *                    Account                      *
     ***************************************************/

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
        address managerImplementation = ICMAccountManager(getManagerAddress()).getAccountImplementation();
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
     * @dev Returns true if an address is authorized to sign cheques
     */
    function isBotAllowed(address bot) public view override returns (bool) {
        return hasRole(CHEQUE_OPERATOR_ROLE, bot);
    }

    /**
     * @dev Verifies if the amount is withdrawable by checking if prefund is spent
     */
    function _checkPrefundSpent(uint256 amount) private view {
        uint256 prefundAmount = getPrefundAmount();
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
    function withdraw(address payable recipient, uint256 amount) external onlyRole(WITHDRAWER_ROLE) {
        // Check if amount is withdrawable according to the prefund spent amount
        _checkPrefundSpent(amount);

        recipient.sendValue(amount);
        emit Withdraw(recipient, amount);
    }

    /***************************************************
     *                 BOOKING TOKEN                   *
     ***************************************************/

    /**
     * @dev Mint booking token
     *
     * @param reservedFor The account to reserve the token for
     * @param uri The URI of the token
     * @param expirationTimestamp The expiration timestamp
     * @param price The price of the token
     * @param paymentToken The payment token, if address(0) then native
     */
    function mintBookingToken(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken
    ) external onlyRole(BOOKING_OPERATOR_ROLE) {
        // Mint the token
        BookingTokenOperator._mintBookingToken(
            getBookingTokenAddress(),
            reservedFor,
            uri,
            expirationTimestamp,
            price,
            paymentToken
        );
    }

    /**
     * @dev Buy booking token
     */
    function buyBookingToken(uint256 tokenId) external onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator._buyBookingToken(getBookingTokenAddress(), tokenId);
    }

    /**
     * @dev See {IERC721Receiver-onERC721Received}.
     *
     * Always returns `IERC721Receiver.onERC721Received.selector`.
     */
    function onERC721Received(address, address, uint256, bytes memory) public virtual returns (bytes4) {
        return this.onERC721Received.selector;
    }

    /***************************************************
     *                ERC20 & ERC721                   *
     ***************************************************/

    function transferERC20(IERC20 token, address to, uint256 amount) external onlyRole(WITHDRAWER_ROLE) {
        if (to == address(0)) {
            revert TransferToZeroAddress();
        }
        token.safeTransfer(to, amount);
    }

    function transferERC721(IERC721 token, address to, uint256 tokenId) external onlyRole(WITHDRAWER_ROLE) {
        if (to == address(0)) {
            revert TransferToZeroAddress();
        }
        token.safeTransferFrom(address(this), to, tokenId);
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
        bool restrictedRate,
        string[] memory capabilities
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        _addService(getServiceHash(serviceName), fee, capabilities, restrictedRate);
    }

    /**
     * @dev Remove a service from the account by its name
     */
    function removeService(string memory serviceName) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removeService(getServiceHash(serviceName));
    }

    // FEE

    /**
     * @dev Set the fee of a service by name
     */
    function setServiceFee(string memory serviceName, uint256 fee) public onlyRole(SERVICE_ADMIN_ROLE) {
        _setServiceFee(getServiceHash(serviceName), fee);
    }

    // RESTRICTED RATE

    /**
     * @dev Set the restricted rate of a service by name
     */
    function setServiceRestrictedRate(
        string memory serviceName,
        bool restrictedRate
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        _setServiceRestrictedRate(getServiceHash(serviceName), restrictedRate);
    }

    // ALL CAPABILITIES

    /**
     * @dev Set all capabilities for a service by name
     */
    function setServiceCapabilities(
        string memory serviceName,
        string[] memory capabilities
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        _setServiceCapabilities(getServiceHash(serviceName), capabilities);
    }

    // SINGLE CAPABILITY

    /**
     * @dev Add a single capability to the service by name
     */
    function addServiceCapability(
        string memory serviceName,
        string memory capability
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        _addServiceCapability(getServiceHash(serviceName), capability);
    }

    /**
     * @dev Remove a single capability from the service by name
     */
    function removeServiceCapability(
        string memory serviceName,
        string memory capability
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removeServiceCapability(getServiceHash(serviceName), capability);
    }

    /**
     * @dev Get service hash by name. Returns the keccak256 hash of the service name
     * from the account manager
     */
    function getServiceHash(string memory serviceName) private view returns (bytes32 serviceHash) {
        return ICMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(serviceName);
    }

    /***************************************************
     *           SERVICES WITH RESOLVED NAMES          *
     ***************************************************/

    /**
     * @dev Get all supported services. Return a list of service names and a list of service objects.
     */
    function getSupportedServices() public view returns (string[] memory serviceNames, Service[] memory services) {
        // Get all hashes and create a list with predefined length
        bytes32[] memory _serviceHashes = getAllServiceHashes();
        string[] memory _serviceNames = new string[](_serviceHashes.length);
        Service[] memory _allSupportedServicesList = new Service[](_serviceHashes.length);

        for (uint256 i = 0; i < _serviceHashes.length; i++) {
            _serviceNames[i] = ICMAccountManager(getManagerAddress()).getRegisteredServiceNameByHash(_serviceHashes[i]);
            _allSupportedServicesList[i] = getService(_serviceHashes[i]);
        }

        return (_serviceNames, _allSupportedServicesList);
    }

    /**
     * @dev Get service fee by name. Overloading the getServiceFee function.
     */
    function getServiceFee(string memory serviceName) public view returns (uint256 fee) {
        return getServiceFee(getServiceHash(serviceName));
    }

    /**
     * @dev Get service restricted rate by name. Overloading the getServiceRestrictedRate function.
     */
    function getServiceRestrictedRate(string memory serviceName) public view returns (bool restrictedRate) {
        return getServiceRestrictedRate(getServiceHash(serviceName));
    }

    /**
     * @dev Get service capabilities by name. Overloading the getServiceCapabilities function.
     */
    function getServiceCapabilities(string memory serviceName) public view returns (string[] memory capabilities) {
        return getServiceCapabilities(getServiceHash(serviceName));
    }

    /***************************************************
     *                WANTED SERVICES                  *
     ***************************************************/

    function addWantedServices(string[] memory serviceNames) public onlyRole(SERVICE_ADMIN_ROLE) {
        for (uint256 i = 0; i < serviceNames.length; i++) {
            bytes32 serviceHash = getServiceHash(serviceNames[i]);
            _addWantedService(serviceHash);
        }
    }

    function removeWantedServices(string[] memory serviceNames) public onlyRole(SERVICE_ADMIN_ROLE) {
        for (uint256 i = 0; i < serviceNames.length; i++) {
            bytes32 serviceHash = getServiceHash(serviceNames[i]);
            _removeWantedService(serviceHash);
        }
    }

    function getWantedServices() public view returns (string[] memory serviceNames) {
        bytes32[] memory _wantedServiceHashes = getWantedServiceHashes();

        string[] memory _wantedServiceNames = new string[](_wantedServiceHashes.length);

        for (uint256 i = 0; i < _wantedServiceHashes.length; i++) {
            _wantedServiceNames[i] = ICMAccountManager(getManagerAddress()).getRegisteredServiceNameByHash(
                _wantedServiceHashes[i]
            );
        }

        return _wantedServiceNames;
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
     *
     * @param pubKeyAddress address of the public key
     * @param data public key data
     */
    function addPublicKey(address pubKeyAddress, bytes memory data) public onlyRole(SERVICE_ADMIN_ROLE) {
        _addPublicKey(pubKeyAddress, data);
    }

    /**
     * @dev Remove public key by address
     */
    function removePublicKey(address pubKeyAddress) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removePublicKey(pubKeyAddress);
    }

    /***************************************************
     *                MESSENGER BOTS                   *
     ***************************************************/

    /**
     * @dev Add messenger bot with initial gas money
     */
    function addMessengerBot(address bot, uint256 gasMoney) public onlyRole(BOT_ADMIN_ROLE) {
        // Check if we can spend the gasMoney to send it to the bot
        _checkPrefundSpent(gasMoney);

        // Grant roles to bot
        _grantRole(CHEQUE_OPERATOR_ROLE, bot);
        _grantRole(BOOKING_OPERATOR_ROLE, bot);
        _grantRole(GAS_WITHDRAWER_ROLE, bot);

        emit MessengerBotAdded(bot);

        // Send gasMoney to bot
        payable(bot).sendValue(gasMoney);
    }

    /**
     * @dev Remove messenger bot
     */
    function removeMessengerBot(address bot) public onlyRole(BOT_ADMIN_ROLE) {
        _revokeRole(CHEQUE_OPERATOR_ROLE, bot);
        _revokeRole(BOOKING_OPERATOR_ROLE, bot);
        _revokeRole(GAS_WITHDRAWER_ROLE, bot);

        emit MessengerBotRemoved(bot);
    }

    /***************************************************
     *              GAS MONEY WITHDRAW                 *
     ***************************************************/

    function withdrawGasMoney(uint256 amount) public onlyRole(GAS_WITHDRAWER_ROLE) {
        _checkPrefundSpent(amount);
        _withdrawGasMoney(amount);
    }

    function setGasMoneyWithdrawal(uint256 limit, uint256 period) public onlyRole(BOT_ADMIN_ROLE) {
        _setGasMoneyWithdrawal(limit, period);
    }
}
