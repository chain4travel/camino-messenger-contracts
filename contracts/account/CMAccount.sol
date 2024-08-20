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
    BookingTokenOperator,
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

    function getManagerAddress() public view returns (address) {
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
     * @dev Returns true if an address is authorized bot
     */
    function isBotAllowed(address bot) public view override returns (bool) {
        return hasRole(CHEQUE_OPERATOR_ROLE, bot);
    }

    /**
     * @dev Return true if address is a registered CMAccount on the CMAccountManager
     */
    function isCMAccount(address account) internal view override returns (bool) {
        return ICMAccountManager(getManagerAddress()).isCMAccount(account);
    }

    /**
     * @dev Return developer wallet address
     */
    function getDeveloperWallet() public view override returns (address) {
        address developerWallet = ICMAccountManager(getManagerAddress()).getDeveloperWallet();
        return developerWallet;
    }

    /**
     * @dev Return developer fee in basis points
     */
    function getDeveloperFeeBp() public view override returns (uint256) {
        uint256 developerFeeBp = ICMAccountManager(getManagerAddress()).getDeveloperFeeBp();
        return developerFeeBp;
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
    ) external override onlyRole(BOOKING_OPERATOR_ROLE) {
        // Mint the token
        _mintBookingToken(getBookingTokenAddress(), reservedFor, uri, expirationTimestamp, price, paymentToken);
    }

    /**
     * @dev Buy booking token
     */
    function buyBookingToken(uint256 tokenId) external override onlyRole(BOOKING_OPERATOR_ROLE) {
        _buyBookingToken(getBookingTokenAddress(), tokenId);
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
        // Check if the service is registered. This function reverts if the service is not registered
        bytes32 serviceHash = ICMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(serviceName);

        // Create the service object
        Service memory service = Service({ _fee: fee, _capabilities: capabilities, _restrictedRate: restrictedRate });

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
        bytes32 serviceHash = ICMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(serviceName);
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
        bytes32 serviceHash = ICMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(serviceName);
        _setServiceFee(serviceHash, fee);
    }

    // RESTRICTED RATE

    /**
     * @dev Set the restricted rate of a service by hash
     */
    function setServiceRestrictedRate(bytes32 serviceHash, bool restrictedRate) public onlyRole(SERVICE_ADMIN_ROLE) {
        _setServiceRestrictedRate(serviceHash, restrictedRate);
    }

    /**
     * @dev Set the restricted rate of a service by name
     */
    function setServiceRestrictedRate(
        string memory serviceName,
        bool restrictedRate
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        bytes32 serviceHash = ICMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(serviceName);
        _setServiceRestrictedRate(serviceHash, restrictedRate);
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
        bytes32 serviceHash = ICMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(serviceName);
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
        bytes32 serviceHash = ICMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(serviceName);
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
        bytes32 serviceHash = ICMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(serviceName);
        _removeServiceCapability(serviceHash, capability);
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
        bytes32 serviceHash = ICMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(serviceName);
        return getServiceFee(serviceHash);
    }

    /**
     * @dev Get service restricted rate by name. Overloading the getServiceRestrictedRate function.
     */
    function getServiceRestrictedRate(string memory serviceName) public view returns (bool restrictedRate) {
        bytes32 serviceHash = ICMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(serviceName);
        return getServiceRestrictedRate(serviceHash);
    }

    /**
     * @dev Get service capabilities by name. Overloading the getServiceCapabilities function.
     */
    function getServiceCapabilities(string memory serviceName) public view returns (string[] memory capabilities) {
        bytes32 serviceHash = ICMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(serviceName);
        return getServiceCapabilities(serviceHash);
    }

    /***************************************************
     *                WANTED SERVICES                  *
     ***************************************************/

    function addWantedServices(string[] memory serviceNames) public onlyRole(SERVICE_ADMIN_ROLE) {
        for (uint256 i = 0; i < serviceNames.length; i++) {
            bytes32 serviceHash = ICMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(
                serviceNames[i]
            );
            _addWantedService(serviceHash);
        }
    }

    function removeWantedServices(string[] memory serviceNames) public onlyRole(SERVICE_ADMIN_ROLE) {
        for (uint256 i = 0; i < serviceNames.length; i++) {
            bytes32 serviceHash = ICMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(
                serviceNames[i]
            );
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
     * @param use type of public key, enum is defined in PartnerConfiguration contract
     */
    function addPublicKey(address pubKeyAddress, bytes memory data, uint8 use) public onlyRole(SERVICE_ADMIN_ROLE) {
        _addPublicKey(pubKeyAddress, data, use);
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

    // FIXME: Should we allow all bots to be able to mint booking tokens?
    // TODO: Create tests for this

    /**
     * @dev Add messenger bot
     */
    function addMessengerBot(address bot) public onlyRole(BOT_ADMIN_ROLE) {
        // Grant roles to bot
        _grantRole(CHEQUE_OPERATOR_ROLE, bot);
        _grantRole(BOOKING_OPERATOR_ROLE, bot);
        _grantRole(GAS_WITHDRAWER_ROLE, bot);

        emit MessengerBotAdded(bot);
    }

    function addMessengerBot(address bot, uint256 gasMoney) public onlyRole(BOT_ADMIN_ROLE) {
        // Check if we can spend the gasMoney to send it to the bot
        _checkPrefundSpent(gasMoney);

        // Grant roles to bot
        addMessengerBot(bot);

        // Send gasMoney to bot
        payable(bot).sendValue(gasMoney);
    }

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
        _withdrawGasMoney(amount);
    }

    function setGasMoneyWithdrawalLimit(uint256 limit) public onlyRole(BOT_ADMIN_ROLE) {
        _setGasMoneyWithdrawalLimit(limit);
    }

    function setGasMoneyWithdrawalPeriod(uint256 period) public onlyRole(BOT_ADMIN_ROLE) {
        _setGasMoneyWithdrawalPeriod(period);
    }
}
