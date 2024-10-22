// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Account Implementation

pragma solidity 0.8.24;

import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { UUPSUpgradeable } from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import { AccessControlEnumerableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlEnumerableUpgradeable.sol";
import { Address } from "@openzeppelin/contracts/utils/Address.sol";
import { SafeERC20, IERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { IERC721 } from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import { IERC721Receiver } from "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import { ERC1967Utils } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Utils.sol";

import { ICMAccountManager } from "../manager/ICMAccountManager.sol";
import { ChequeManager } from "./ChequeManager.sol";
import { BookingTokenOperator } from "../booking-token/BookingTokenOperator.sol";
import { PartnerConfiguration } from "../partner/PartnerConfiguration.sol";
import { GasMoneyManager } from "./GasMoneyManager.sol";

/**
 * @title Camino Messenger Account
 * @notice A CM Account manages funds, minting/buying of booking tokens, provided
 * or wanted services, and multiple bots for distributors and suppliers on
 * Camino Messenger ecosystem.
 *
 * Registering bots is done by role based access control. Bot's with
 * `CHEQUE_OPERATOR_ROLE` can issue cheques to paid by the {CMAccount} contract.
 * Bot can also have `GAS_WITHDRAWER_ROLE` and `BOOKING_OPERATOR_ROLE`.
 *
 * `GAS_WITHDRAWER_ROLE` enables a bot to withdraw native coins (CAM) from the
 * contract to be used as gas money. This restricted with a `limit`
 * (wei/aCAM) and `period` (seconds) by the `BOT_ADMIN_ROLE`. Default starting
 * values are 10 CAM per 24 hours.
 *
 * `BOOKING_OPERATOR_ROLE` enables a bot to mint and buy Booking Tokens by
 * calling the corresponding functions on the {BookingToken} contract. The buy
 * operation pays the price of the Booking Token with the funds on the
 * {CMAccount} contract.
 *
 * @dev This contract uses UUPS style upgradeability. The authorization function
 * `_authorizeUpgrade(address)` can be called by the `UPGRADER_ROLE` and is
 * restricted to only upgrade to the implementation address registered on the
 * {CMAccountManager} contract.
 * @custom:security-contact https://r.xyz/program/camino-network
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

    /**
     * @notice Upgrader role can upgrade the contract to a new implementation.
     */
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

    /**
     * @notice Bot admin role can add & remove bots and set gas money withdrawal
     * parameters.
     */
    bytes32 public constant BOT_ADMIN_ROLE = keccak256("BOT_ADMIN_ROLE");

    /**
     * @notice Cheque operator role can issue cheques to be paid by this CMAccount
     * contract.
     */
    bytes32 public constant CHEQUE_OPERATOR_ROLE = keccak256("CHEQUE_OPERATOR_ROLE");

    /**
     * @notice Gas withdrawer role can withdraw gas money from the contract. This is
     * intended to be used by the bots and is granted when `addMessengerBot` is
     * called.
     */
    bytes32 public constant GAS_WITHDRAWER_ROLE = keccak256("GAS_WITHDRAWER_ROLE");

    /**
     * @notice Withdrawer role can withdraw funds from the contract.
     */
    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");

    /**
     * @notice Booking operator role can mint and buy booking tokens using the
     * functions on this contract. This is generally used by the bots. The
     * price for the booking token is paid by this contract.
     */
    bytes32 public constant BOOKING_OPERATOR_ROLE = keccak256("BOOKING_OPERATOR_ROLE");

    /**
     * @notice Service admin role can add & remove supported & wanted services.
     */
    bytes32 public constant SERVICE_ADMIN_ROLE = keccak256("SERVICE_ADMIN_ROLE");

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    /// @custom:storage-location erc7201:camino.messenger.storage.CMAccount
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
     * @notice CMAccount upgrade event. Emitted when the CMAccount implementation is upgraded.
     */
    event CMAccountUpgraded(address indexed oldImplementation, address indexed newImplementation);

    /**
     * @notice Deposit event, emitted when there is a new deposit
     */
    event Deposit(address indexed sender, uint256 amount);

    /**
     * @notice Withdraw event, emitted when there is a new withdrawal
     */
    event Withdraw(address indexed receiver, uint256 amount);

    /**
     * @notice Messenger bot added
     */
    event MessengerBotAdded(address indexed bot);

    /**
     * @notice Messenger bot removed
     */
    event MessengerBotRemoved(address indexed bot);

    // Partner Config Events

    event ServiceAdded(string indexed serviceName);
    event ServiceRemoved(string indexed serviceName);

    event WantedServiceAdded(string indexed serviceName);
    event WantedServiceRemoved(string indexed serviceName);

    event ServiceFeeUpdated(string indexed serviceName, uint256 fee);
    event ServiceRestrictedRateUpdated(string indexed serviceName, bool restrictedRate);

    event ServiceCapabilitiesUpdated(string indexed serviceName);
    event ServiceCapabilityAdded(string indexed serviceName, string capability);
    event ServiceCapabilityRemoved(string indexed serviceName, string capability);

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    /**
     * @notice CMAccount implementation address does not match the one in the manager
     */
    error CMAccountImplementationMismatch(address latestImplementation, address newImplementation);

    /**
     * @notice New implementation is the same as the current implementation, no update needed
     */
    error CMAccountNoUpgradeNeeded(address oldImplementation, address newImplementation);

    /**
     * @notice Error to revert with if depositer is not allowed
     */
    error DepositorNotAllowed(address sender);

    /**
     * @notice Error to revert zero value deposits
     */
    error ZeroValueDeposit(address sender);

    /**
     * @notice Error to revert with if the prefund is not spent yet
     */
    error PrefundNotSpentYet(uint256 withdrawableAmount, uint256 prefundLeft, uint256 amount);

    /**
     * @notice Error to revert if transfer to zero address
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

    /**
     * @notice Returns the CMAccountManager address.
     *
     * @return CMAccountManager address
     */
    function getManagerAddress() public view override returns (address) {
        CMAccountStorage storage $ = _getCMAccountStorage();
        return $._manager;
    }

    /**
     * @notice Returns the booking token address.
     *
     * @return BookingToken address
     */
    function getBookingTokenAddress() public view returns (address) {
        CMAccountStorage storage $ = _getCMAccountStorage();
        return $._bookingToken;
    }

    /**
     * @notice Returns the prefund amount.
     *
     * @return prefund amount
     */
    function getPrefundAmount() public view returns (uint256) {
        CMAccountStorage storage $ = _getCMAccountStorage();
        return $._prefundAmount;
    }

    /***************************************************
     *                    Account                      *
     ***************************************************/

    /**
     * @notice Authorizes the upgrade of the CMAccount.
     *
     * Reverts if the new implementation is the same as the old one.
     *
     * Reverts if the new implementation does not match the implementation address
     * in the manager. Only implementations registered at the manager are allowed.
     *
     * @dev Emits a {CMAccountUpgraded} event.
     *
     * @param newImplementation The new implementation address
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
     * @notice Returns true if an address is authorized to sign cheques
     *
     * @param bot The bot's address
     */
    function isBotAllowed(address bot) public view override returns (bool) {
        return hasRole(CHEQUE_OPERATOR_ROLE, bot);
    }

    /**
     * @notice Verifies if the amount is withdrawable by checking if prefund is spent
     *
     * @param amount The amount to check if it's withdrawable
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
     * @notice Withdraw CAM from the CMAccount
     *
     * This function reverts if the amount is bigger then the prefund left to spend. This is to prevent
     * spam by forcing user to spend the full prefund for cheques, so they can not just create an account
     * and withdraw the prefund.
     */
    function withdraw(address payable recipient, uint256 amount) external nonReentrant onlyRole(WITHDRAWER_ROLE) {
        // Check if amount is withdrawable according to the prefund spent amount
        _checkPrefundSpent(amount);

        recipient.sendValue(amount);
        emit Withdraw(recipient, amount);
    }

    /***************************************************
     *                 BOOKING TOKEN                   *
     ***************************************************/

    /**
     * @notice Mints booking token.
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
        IERC20 paymentToken,
        bool _isCancellable
    ) external onlyRole(BOOKING_OPERATOR_ROLE) {
        // Mint the token
        BookingTokenOperator.mintBookingToken(
            getBookingTokenAddress(),
            reservedFor,
            uri,
            expirationTimestamp,
            price,
            paymentToken,
            _isCancellable
        );
    }

    /**
     * @notice Buys booking token.
     *
     * @param tokenId The token id
     */
    function buyBookingToken(uint256 tokenId) external nonReentrant onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.buyBookingToken(getBookingTokenAddress(), tokenId);
    }

    /**
     * @notice Record expiration status if the token is expired
     */
    function recordExpiration(uint256 tokenId) external onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.recordExpiration(getBookingTokenAddress(), tokenId);
    }

    /**
     * @notice Set cancellable flag for booking token
     * @param tokenId The token id
     * @param cancellable The cancellable flag
     */
    function setCancellable(uint256 tokenId, bool cancellable) external onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.setCancellable(getBookingTokenAddress(), tokenId, cancellable);
    }

    /**
     * @notice Always returns `IERC721Receiver.onERC721Received.selector`.
     *
     * @dev See {IERC721Receiver-onERC721Received}.
     */
    function onERC721Received(address, address, uint256, bytes memory) public virtual returns (bytes4) {
        return this.onERC721Received.selector;
    }

    /***************************************************
     *                ERC20 & ERC721                   *
     ***************************************************/

    /**
     * @notice Transfers ERC20 tokens.
     *
     * This function reverts if `to` is the zero address.
     *
     * @param token The ERC20 token
     * @param to The address to transfer the tokens to
     * @param amount The amount of tokens to transfer
     */
    function transferERC20(IERC20 token, address to, uint256 amount) external onlyRole(WITHDRAWER_ROLE) {
        if (to == address(0)) {
            revert TransferToZeroAddress();
        }
        token.safeTransfer(to, amount);
    }

    /**
     * @notice Transfers ERC721 tokens.
     *
     * This function reverts if `to` is the zero address.
     *
     * @param token The ERC721 token
     * @param to The address to transfer the tokens to
     * @param tokenId The token id of the token
     */
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
     * @notice Adds a service to the account as a supported service.
     *
     * `serviceName` is defined as pkg + service name in protobuf. For example:
     *
     * ```text
     *  ┌────────────── pkg ─────────────┐ ┌───── service name ─────┐
     * "cmp.services.accommodation.v1alpha.AccommodationSearchService")
     * ```
     *
     * @dev These services are coming from the Camino Messenger Protocol's protobuf
     * definitions.
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
        emit ServiceAdded(serviceName);
    }

    /**
     * @notice Remove a service from the account by its name
     */
    function removeService(string memory serviceName) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removeService(getServiceHash(serviceName));
        emit ServiceRemoved(serviceName);
    }

    // FEE

    /**
     * @notice Set the fee of a service by name
     */
    function setServiceFee(string memory serviceName, uint256 fee) public onlyRole(SERVICE_ADMIN_ROLE) {
        _setServiceFee(getServiceHash(serviceName), fee);
        emit ServiceFeeUpdated(serviceName, fee);
    }

    // RESTRICTED RATE

    /**
     * @notice Set the restricted rate of a service by name
     */
    function setServiceRestrictedRate(
        string memory serviceName,
        bool restrictedRate
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        _setServiceRestrictedRate(getServiceHash(serviceName), restrictedRate);
        emit ServiceRestrictedRateUpdated(serviceName, restrictedRate);
    }

    // ALL CAPABILITIES

    /**
     * @notice Set all capabilities for a service by name
     */
    function setServiceCapabilities(
        string memory serviceName,
        string[] memory capabilities
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        _setServiceCapabilities(getServiceHash(serviceName), capabilities);
        emit ServiceCapabilitiesUpdated(serviceName);
    }

    // SINGLE CAPABILITY

    /**
     * @notice Add a single capability to the service by name
     */
    function addServiceCapability(
        string memory serviceName,
        string memory capability
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        _addServiceCapability(getServiceHash(serviceName), capability);
        emit ServiceCapabilityAdded(serviceName, capability);
    }

    /**
     * @notice Remove a single capability from the service by name
     */
    function removeServiceCapability(
        string memory serviceName,
        string memory capability
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removeServiceCapability(getServiceHash(serviceName), capability);
        emit ServiceCapabilityRemoved(serviceName, capability);
    }

    /**
     * @notice Get service hash by name. Returns the keccak256 hash of the service name
     * from the account manager
     */
    function getServiceHash(string memory serviceName) private view returns (bytes32 serviceHash) {
        return ICMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(serviceName);
    }

    /***************************************************
     *           SERVICES WITH RESOLVED NAMES          *
     ***************************************************/

    /**
     * @notice Get all supported services. Return a list of service names and a list of service objects.
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
     * @notice Get service fee by name. Overloading the getServiceFee function.
     */
    function getServiceFee(string memory serviceName) public view returns (uint256 fee) {
        return getServiceFee(getServiceHash(serviceName));
    }

    /**
     * @notice Get service restricted rate by name. Overloading the getServiceRestrictedRate function.
     */
    function getServiceRestrictedRate(string memory serviceName) public view returns (bool restrictedRate) {
        return getServiceRestrictedRate(getServiceHash(serviceName));
    }

    /**
     * @notice Get service capabilities by name. Overloading the getServiceCapabilities function.
     */
    function getServiceCapabilities(string memory serviceName) public view returns (string[] memory capabilities) {
        return getServiceCapabilities(getServiceHash(serviceName));
    }

    /***************************************************
     *                WANTED SERVICES                  *
     ***************************************************/

    /**
     * @notice Adds wanted services.
     *
     * @param serviceNames List of service names
     */
    function addWantedServices(string[] memory serviceNames) public onlyRole(SERVICE_ADMIN_ROLE) {
        for (uint256 i = 0; i < serviceNames.length; i++) {
            bytes32 serviceHash = getServiceHash(serviceNames[i]);
            _addWantedService(serviceHash);
            emit WantedServiceAdded(serviceNames[i]);
        }
    }

    /**
     * @notice Removes wanted services.
     *
     * @param serviceNames List of service names
     */
    function removeWantedServices(string[] memory serviceNames) public onlyRole(SERVICE_ADMIN_ROLE) {
        for (uint256 i = 0; i < serviceNames.length; i++) {
            bytes32 serviceHash = getServiceHash(serviceNames[i]);
            _removeWantedService(serviceHash);
            emit WantedServiceRemoved(serviceNames[i]);
        }
    }

    /**
     * @notice Get all wanted services.
     *
     * @return serviceNames List of service names
     */
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

    /**
     * @notice Sets if off-chain payment is supported.
     *
     * @param _isSupported true if off-chain payment is supported
     */
    function setOffChainPaymentSupported(bool _isSupported) public onlyRole(SERVICE_ADMIN_ROLE) {
        _setOffChainPaymentSupported(_isSupported);
    }

    /**
     * @notice Adds a supported payment token.
     *
     * @param _supportedToken address of the token
     */
    function addSupportedToken(address _supportedToken) public onlyRole(SERVICE_ADMIN_ROLE) {
        _addSupportedToken(_supportedToken);
    }

    /**
     * @notice Removes a supported payment token.
     *
     * @param _supportedToken address of the token
     */
    function removeSupportedToken(address _supportedToken) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removeSupportedToken(_supportedToken);
    }

    /***************************************************
     *                  PUBLIC KEY                     *
     ***************************************************/

    /**
     * @notice Add public key with address
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
     * @notice Remove public key by address
     */
    function removePublicKey(address pubKeyAddress) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removePublicKey(pubKeyAddress);
    }

    /***************************************************
     *                MESSENGER BOTS                   *
     ***************************************************/

    /**
     * @notice Adds messenger bot with initial gas money.
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
     * @notice Removes messenger bot by revoking the roles.
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

    /**
     * @notice Withdraw gas money. Requires the `GAS_WITHDRAWER_ROLE`.
     *
     * @param amount The amount to withdraw in aCAM (wei)
     */
    function withdrawGasMoney(uint256 amount) public nonReentrant onlyRole(GAS_WITHDRAWER_ROLE) {
        _checkPrefundSpent(amount);
        _withdrawGasMoney(amount);
    }

    /**
     * @notice Set gas money withdrawal parameters. Requires the `BOT_ADMIN_ROLE`.
     *
     * @param limit Amount of gas money to withdraw in wei per period
     * @param period Duration of the withdrawal period in seconds
     */
    function setGasMoneyWithdrawal(uint256 limit, uint256 period) public onlyRole(BOT_ADMIN_ROLE) {
        _setGasMoneyWithdrawal(limit, period);
    }

    /***************************************************
     *                 CANCELLATION                    *
     ***************************************************/

    function initiateCancellationProposal(
        uint256 tokenId,
        uint256 refundAmount
    ) public onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.initiateCancellationProposal(getBookingTokenAddress(), tokenId, refundAmount);
    }

    function acceptCancellationProposal(uint256 tokenId) public onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.acceptCancellationProposal(getBookingTokenAddress(), tokenId);
    }

    function counterCancellationProposal(uint256 tokenId, uint256 refundAmount) public onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.counterCancellationProposal(getBookingTokenAddress(), tokenId, refundAmount);
    }

    function acceptCounteredCancellationProposal(uint256 tokenId) public onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.acceptCounteredCancellationProposal(getBookingTokenAddress(), tokenId);
    }

    function cancelCancellationProposal(uint256 tokenId) public onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.cancelCancellationProposal(getBookingTokenAddress(), tokenId);
    }
}
