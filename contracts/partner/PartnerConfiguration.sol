// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Partner Configuration

pragma solidity ^0.8.24;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { EnumerableSet } from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @title PartnerConfiguration
 * @notice Partner Configuration is used by the {CMAccount} contract to register
 * supported and wanted services by the partner.
 */
abstract contract PartnerConfiguration is Initializable {
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.AddressSet;

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    /**
     * @notice Struct for storing supported service details for suppliers
     */
    struct Service {
        uint256 _fee;
        /**
         * @dev  If set to true, this means the service is restricted to pre-agreement
         * with the partner. (Not a rack rate)
         */
        bool _restrictedRate;
        string[] _capabilities;
    }

    struct PaymentInfo {
        bool _supportsOffChainPayment; // Supports off chain payments if true
        EnumerableSet.AddressSet _supportedTokens; // Supported on-chain token for payment
    }

    /// @custom:storage-location erc7201:camino.messenger.storage.PartnerConfiguration
    struct PartnerConfigurationStorage {
        // Set of supported service hashes
        EnumerableSet.Bytes32Set _servicesHashSet;
        // Mapping of service hashes to the service details
        mapping(bytes32 _serviceHash => Service _service) _supportedServices;
        // Payment
        PaymentInfo _paymentInfo;
        // Public keys, keep a enumerable list of public key addresses
        EnumerableSet.AddressSet _publicKeyAddressesSet;
        // Public keys for encrypting private data for Booking Token
        mapping(address publicKeyAddress => bytes publicKey) _publicKeys;
        // Services that this distributors want to buy
        EnumerableSet.Bytes32Set _wantedServicesHashSet;
    }

    // keccak256(abi.encode(uint256(keccak256("camino.messenger.storage.PartnerConfiguration")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant PartnerConfigurationStorageLocation =
        0xf2856e5e1b7689dcde1bb551fd115c3cad8d243ea609d47a46b4d22ee58d3000;

    function _getPartnerConfigurationStorage() private pure returns (PartnerConfigurationStorage storage $) {
        assembly {
            $.slot := PartnerConfigurationStorageLocation
        }
    }

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    error ServiceAlreadyExists(bytes32 serviceHash);
    error ServiceDoesNotExist(bytes32 serviceHash);

    error WantedServiceAlreadyExists(bytes32 serviceHash);
    error WantedServiceDoesNotExist(bytes32 serviceHash);

    error PaymentTokenAlreadyExists(address token);
    error PaymentTokenDoesNotExist(address token);

    error PublicKeyAlreadyExists(address pubKeyAddress);
    error PublicKeyDoesNotExist(address pubKeyAddress);
    error InvalidPublicKeyUseType(uint8 use);

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    event PaymentTokenAdded(address indexed token);
    event PaymentTokenRemoved(address indexed token);

    event OffChainPaymentSupportUpdated(bool supportsOffChainPayment);

    event PublicKeyAdded(address indexed pubKeyAddress);
    event PublicKeyRemoved(address indexed pubKeyAddress);

    /***************************************************
     *                 INITIALIZATION                  *
     ***************************************************/

    function __PartnerConfiguration_init() internal onlyInitializing {}

    function __PartnerConfiguration_init_unchained() internal onlyInitializing {}

    /***************************************************
     *                   SERVICE                       *
     ***************************************************/

    /**
     * @notice Adds a supported Service object for a given hash.
     *
     * @param serviceHash Hash of the service
     * @param fee Fee for the service
     * @param capabilities Capabilities for the service
     * @param restrictedRate If the service is restricted to pre-agreement
     */
    function _addService(
        bytes32 serviceHash,
        uint256 fee,
        string[] memory capabilities,
        bool restrictedRate
    ) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Try to add the service to the services hash set
        bool added = $._servicesHashSet.add(serviceHash);
        if (!added) {
            revert ServiceAlreadyExists(serviceHash);
        }
        $._supportedServices[serviceHash] = Service({
            _fee: fee,
            _capabilities: capabilities,
            _restrictedRate: restrictedRate
        });
    }

    /**
     * @notice Removes a supported Service object for a given hash.
     *
     * @param serviceHash Hash of the service
     */
    function _removeService(bytes32 serviceHash) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Try to remove the service
        bool removed = $._servicesHashSet.remove(serviceHash);
        if (!removed) {
            revert ServiceDoesNotExist(serviceHash);
        }

        delete $._supportedServices[serviceHash];
    }

    /**
     * @notice Sets the Service fee for a given hash.
     *
     * @param serviceHash Hash of the service
     * @param fee Fee
     */
    function _setServiceFee(bytes32 serviceHash, uint256 fee) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        _checkServiceExists(serviceHash, $);

        $._supportedServices[serviceHash]._fee = fee;
    }

    /**
     * @notice Sets the Service restricted rate for a given hash.
     *
     * @param serviceHash Hash of the service
     * @param restrictedRate Restricted rate
     */
    function _setServiceRestrictedRate(bytes32 serviceHash, bool restrictedRate) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        _checkServiceExists(serviceHash, $);

        $._supportedServices[serviceHash]._restrictedRate = restrictedRate;
    }

    /**
     * @notice Sets the Service capabilities for a given hash.
     *
     * @param serviceHash Hash of the service
     * @param capabilities Capabilities
     */
    function _setServiceCapabilities(bytes32 serviceHash, string[] memory capabilities) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        _checkServiceExists(serviceHash, $);

        $._supportedServices[serviceHash]._capabilities = capabilities;
    }

    /**
     * @notice Adds a capability to the service.
     *
     * @param serviceHash Hash of the service
     * @param capability Capability
     */
    function _addServiceCapability(bytes32 serviceHash, string memory capability) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        _checkServiceExists(serviceHash, $);

        $._supportedServices[serviceHash]._capabilities.push(capability);
    }

    /**
     * @notice Removes a capability from the service.
     *
     * @param serviceHash Hash of the service
     * @param capability Capability
     */
    function _removeServiceCapability(bytes32 serviceHash, string memory capability) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        _checkServiceExists(serviceHash, $);

        string[] storage capabilities = $._supportedServices[serviceHash]._capabilities;
        for (uint256 i = 0; i < capabilities.length; i++) {
            if (keccak256(abi.encodePacked(capabilities[i])) == keccak256(abi.encodePacked(capability))) {
                capabilities[i] = capabilities[capabilities.length - 1];
                capabilities.pop();
                break;
            }
        }
    }

    /**
     * @notice Returns all supported service hashes.
     */
    function getAllServiceHashes() public view returns (bytes32[] memory serviceHashes) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._servicesHashSet.values();
    }

    /**
     * @notice Returns the Service object for a given hash. Service object contains fee and capabilities.
     *
     * `serviceHash` is keccak256 hash of the pkg + service name as:
     *
     * ```text
     *            ┌────────────── pkg ─────────────┐ ┌───── service name ─────┐
     * keccak256("cmp.services.accommodation.v1alpha.AccommodationSearchService")
     * ```
     * @dev These services are coming from the Camino Messenger Protocol's protobuf
     * definitions.
     *
     * @param serviceHash Hash of the service
     */
    function getService(bytes32 serviceHash) public view virtual returns (Service memory service) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        _checkServiceExists(serviceHash, $);

        return $._supportedServices[serviceHash];
    }

    /**
     * @notice Returns the fee for a given service hash.
     *
     * @param serviceHash Hash of the service
     */
    function getServiceFee(bytes32 serviceHash) public view virtual returns (uint256 fee) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        _checkServiceExists(serviceHash, $);

        return $._supportedServices[serviceHash]._fee;
    }

    /**
     * @notice Returns the restricted rate for a given service hash.
     *
     * @param serviceHash Hash of the service
     */
    function getServiceRestrictedRate(bytes32 serviceHash) public view virtual returns (bool restrictedRate) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        _checkServiceExists(serviceHash, $);

        return $._supportedServices[serviceHash]._restrictedRate;
    }

    /**
     * @notice Returns the capabilities for a given service hash.
     *
     * @param serviceHash Hash of the service
     */
    function getServiceCapabilities(bytes32 serviceHash) public view virtual returns (string[] memory capabilities) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        _checkServiceExists(serviceHash, $);

        return $._supportedServices[serviceHash]._capabilities;
    }

    /**
     * @notice Checks if the service exists.
     *
     * @param serviceHash Hash of the service
     */
    function _checkServiceExists(bytes32 serviceHash, PartnerConfigurationStorage storage $) private view {
        // Check if the service exists
        if (!$._servicesHashSet.contains(serviceHash)) {
            revert ServiceDoesNotExist(serviceHash);
        }
    }

    /***************************************************
     *               WANTED SERVICES                   *
     ***************************************************/

    /**
     * @notice Adds a wanted service hash to the wanted services set.
     *
     * Reverts if the service already exists.
     *
     * @param serviceHash Hash of the service
     */
    function _addWantedService(bytes32 serviceHash) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        bool added = $._wantedServicesHashSet.add(serviceHash);

        if (!added) {
            revert WantedServiceAlreadyExists(serviceHash);
        }
    }

    /**
     * @notice Removes a wanted service hash from the wanted services set.
     *
     * Reverts if the service does not exist.
     *
     * @param serviceHash Hash of the service
     */
    function _removeWantedService(bytes32 serviceHash) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        bool removed = $._wantedServicesHashSet.remove(serviceHash);

        if (!removed) {
            revert WantedServiceDoesNotExist(serviceHash);
        }
    }

    /**
     * @notice Returns all wanted service hashes.
     *
     * @return serviceHashes Wanted service hashes
     */
    function getWantedServiceHashes() public view virtual returns (bytes32[] memory serviceHashes) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._wantedServicesHashSet.values();
    }

    /***************************************************
     *                   PAYMENT                       *
     ***************************************************/

    // PAYMENT INFO: SUPPORTED TOKENS

    /**
     * @notice Adds a supported payment token.
     *
     * @param _token Payment token address to be added
     */
    function _addSupportedToken(address _token) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        bool added = $._paymentInfo._supportedTokens.add(_token);

        if (!added) {
            revert PaymentTokenAlreadyExists(_token);
        }

        emit PaymentTokenAdded(_token);
    }

    /**
     * @notice Removes a supported payment token.
     *
     * @param _token Payment token address to be removed
     */
    function _removeSupportedToken(address _token) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        bool removed = $._paymentInfo._supportedTokens.remove(_token);

        if (!removed) {
            revert PaymentTokenDoesNotExist(_token);
        }

        emit PaymentTokenRemoved(_token);
    }

    /**
     * @notice Returns supported token addresses.
     *
     * @return tokens Supported token addresses
     */
    function getSupportedTokens() public view virtual returns (address[] memory tokens) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._paymentInfo._supportedTokens.values();
    }

    // PAYMENT INFO: OFF-CHAIN PAYMENT SUPPORT

    /**
     * @notice Sets the off-chain payment support is supported.
     */
    function _setOffChainPaymentSupported(bool _supportsOffChainPayment) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        $._paymentInfo._supportsOffChainPayment = _supportsOffChainPayment;
        emit OffChainPaymentSupportUpdated(_supportsOffChainPayment);
    }

    /**
     * @notice Returns true if off-chain payment is supported for the given service.
     */
    function offChainPaymentSupported() public view virtual returns (bool) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._paymentInfo._supportsOffChainPayment;
    }

    /***************************************************
     *                 PUBLIC KEYS                     *
     ***************************************************/

    /**
     * @notice Adds public key with an address. Reverts if the public key already
     * exists.
     *
     * Beware: This functions does not check if the public key is actually for the
     * given address.
     */
    function _addPublicKey(address pubKeyAddress, bytes memory publicKeyData) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        bool added = $._publicKeyAddressesSet.add(pubKeyAddress);

        if (!added) {
            revert PublicKeyAlreadyExists(pubKeyAddress);
        }

        $._publicKeys[pubKeyAddress] = publicKeyData;

        emit PublicKeyAdded(pubKeyAddress);
    }

    /**
     * @notice Removes the public key for a given address
     *
     * Reverts if the public key does not exist
     */
    function _removePublicKey(address pubKeyAddress) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        bool removed = $._publicKeyAddressesSet.remove(pubKeyAddress);

        if (!removed) {
            revert PublicKeyDoesNotExist(pubKeyAddress);
        }

        delete $._publicKeys[pubKeyAddress];
        emit PublicKeyRemoved(pubKeyAddress);
    }

    /**
     * @notice Returns the addresses of all public keys. These can then be used to
     * retrieve the public keys the `getPublicKey(address)` function.
     */
    function getPublicKeysAddresses() public view virtual returns (address[] memory pubKeyAddresses) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._publicKeyAddressesSet.values();
    }

    /**
     * @notice Returns the public key for a given address.
     *
     * Reverts if the public key does not exist
     *
     * @param pubKeyAddress Address of the public key
     */
    function getPublicKey(address pubKeyAddress) public view virtual returns (bytes memory data) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        if (!$._publicKeyAddressesSet.contains(pubKeyAddress)) {
            revert PublicKeyDoesNotExist(pubKeyAddress);
        }

        return $._publicKeys[pubKeyAddress];
    }
}
