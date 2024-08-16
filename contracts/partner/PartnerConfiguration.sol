// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Partner Configuration

pragma solidity ^0.8.24;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

abstract contract PartnerConfiguration is Initializable {
    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.AddressSet;

    struct Service {
        uint256 _fee;
        /**
         * @dev  If set to true, this means the service is restricted to pre-aggreement
         * with the partner. (Not a rack rate)
         */
        bool _restrictedRate;
        string[] _capabilities;
    }

    struct PaymentInfo {
        bool _supportsOffChainPayment; // Supports off chain payments if true
        EnumerableSet.AddressSet _supportedTokens; // Supported on-chain token for payment
    }

    /**
     * @dev Purpose of the public key. Currently we only have one.
     */
    enum PublicKeyUseType {
        EncryptPrivateData
    }

    struct PublicKey {
        PublicKeyUseType _use;
        bytes _data;
    }

    /// @custom:storage-location erc7201:camino.messenger.storage.PartnerConfiguration
    struct PartnerConfigurationStorage {
        EnumerableSet.Bytes32Set _servicesHashSet;
        mapping(bytes32 _serviceHash => Service _service) _supportedServices;
        PaymentInfo _paymentInfo;
        EnumerableSet.AddressSet _publicKeyAddressesSet; // Keep a enumerable list of public key addresses
        mapping(address publicKeyAddress => PublicKey publicKey) _publicKeys; // Public keys for ecrypting private data for Booking Token

        // TODO: Add support for distributors defining services they want to buy
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

    error PaymentTokenAlreadyExists(address token);
    error PaymentTokenDoesNotExist(address token);

    error PublicKeyAlreadyExists(address pubKeyAddress);
    error PublicKeyDoesNotExist(address pubKeyAddress);
    error InvalidPublicKeyUseType(uint8 use);

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    event ServiceAdded(bytes32 serviceHash);
    event ServiceRemoved(bytes32 serviceHash);

    event ServiceFeeUpdated(bytes32 serviceHash, uint256 fee);
    event ServiceRestrictedRateUpdated(bytes32 serviceHash, bool restrictedRate);

    event ServiceCapabilitiesUpdated(bytes32 serviceHash);
    event ServiceCapabilityAdded(bytes32 serviceHash, string capability);
    event ServiceCapabilityRemoved(bytes32 serviceHash, string capability);

    event PaymentTokenAdded(address token);
    event PaymentTokenRemoved(address token);

    event OffChainPaymentSupportUpdated(bool supportsOffChainPayment);

    event PublicKeyAdded(address indexed pubKeyAddress, PublicKey publicKey);
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
     * @dev Adds a supported Service object for a given hash.
     *
     * @param serviceHash Hash of the service
     * @param service Service object
     */
    function _addService(bytes32 serviceHash, Service memory service) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Try to add the service to the services hash set
        bool added = $._servicesHashSet.add(serviceHash);
        if (!added) {
            revert ServiceAlreadyExists(serviceHash);
        }
        $._supportedServices[serviceHash] = service;

        emit ServiceAdded(serviceHash);
    }

    /**
     * @dev Removes a supported Service object for a given hash.
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

        emit ServiceRemoved(serviceHash);
    }

    /**
     * @dev Set the Service fee for a given hash.
     *
     * @param serviceHash Hash of the service
     * @param fee Fee
     */
    function _setServiceFee(bytes32 serviceHash, uint256 fee) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        if (!$._servicesHashSet.contains(serviceHash)) {
            revert ServiceDoesNotExist(serviceHash);
        }

        $._supportedServices[serviceHash]._fee = fee;

        emit ServiceFeeUpdated(serviceHash, fee);
    }

    /**
     * @dev Set the Service restricted rate for a given hash.
     *
     * @param serviceHash Hash of the service
     * @param restrictedRate Restricted rate
     */
    function _setServiceRestrictedRate(bytes32 serviceHash, bool restrictedRate) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        if (!$._servicesHashSet.contains(serviceHash)) {
            revert ServiceDoesNotExist(serviceHash);
        }

        $._supportedServices[serviceHash]._restrictedRate = restrictedRate;

        emit ServiceRestrictedRateUpdated(serviceHash, restrictedRate);
    }

    /**
     * @dev Set the Service capabilities for a given hash.
     *
     * @param serviceHash Hash of the service
     * @param capabilities Capabilities
     */
    function _setServiceCapabilities(bytes32 serviceHash, string[] memory capabilities) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        if (!$._servicesHashSet.contains(serviceHash)) {
            revert ServiceDoesNotExist(serviceHash);
        }

        $._supportedServices[serviceHash]._capabilities = capabilities;

        emit ServiceCapabilitiesUpdated(serviceHash);
    }

    /**
     * @dev Add a capability to the service.
     *
     * @param serviceHash Hash of the service
     * @param capability Capability
     */
    function _addServiceCapability(bytes32 serviceHash, string memory capability) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        if (!$._servicesHashSet.contains(serviceHash)) {
            revert ServiceDoesNotExist(serviceHash);
        }

        $._supportedServices[serviceHash]._capabilities.push(capability);

        emit ServiceCapabilityAdded(serviceHash, capability);
    }

    /**
     * @dev Remove a capability from the service.
     *
     * @param serviceHash Hash of the service
     * @param capability Capability
     */
    function _removeServiceCapability(bytes32 serviceHash, string memory capability) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        if (!$._servicesHashSet.contains(serviceHash)) {
            revert ServiceDoesNotExist(serviceHash);
        }

        string[] storage capabilities = $._supportedServices[serviceHash]._capabilities;
        for (uint256 i = 0; i < capabilities.length; i++) {
            if (keccak256(abi.encodePacked(capabilities[i])) == keccak256(abi.encodePacked(capability))) {
                capabilities[i] = capabilities[capabilities.length - 1];
                capabilities.pop();
                emit ServiceCapabilityRemoved(serviceHash, capability);
                break;
            }
        }
    }

    /**
     * @dev Returns all supported service hashes
     */
    function getAllServiceHashes() public view returns (bytes32[] memory serviceHashes) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._servicesHashSet.values();
    }

    /**
     * @dev Returns the Service object for a given hash. Service object contains fee and capabilities.
     *
     * {serviceHash} is keccak256 hash of the pkg + service name as:
     *
     *            ┌────────────── pkg ─────────────┐ ┌───── service name ─────┐
     * keccak256("cmp.services.accommodation.v1alpha.AccommodationSearchService")
     */
    function getService(bytes32 serviceHash) public view virtual returns (Service memory service) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        if (!$._servicesHashSet.contains(serviceHash)) {
            revert ServiceDoesNotExist(serviceHash);
        }

        return $._supportedServices[serviceHash];
    }

    function getServiceFee(bytes32 serviceHash) public view virtual returns (uint256 fee) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        if (!$._servicesHashSet.contains(serviceHash)) {
            revert ServiceDoesNotExist(serviceHash);
        }

        return $._supportedServices[serviceHash]._fee;
    }

    function getServiceRestrictedRate(bytes32 serviceHash) public view virtual returns (bool restrictedRate) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        if (!$._servicesHashSet.contains(serviceHash)) {
            revert ServiceDoesNotExist(serviceHash);
        }

        return $._supportedServices[serviceHash]._restrictedRate;
    }

    function getServiceCapabilities(bytes32 serviceHash) public view virtual returns (string[] memory capabilities) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        if (!$._servicesHashSet.contains(serviceHash)) {
            revert ServiceDoesNotExist(serviceHash);
        }

        return $._supportedServices[serviceHash]._capabilities;
    }

    /***************************************************
     *                   PAYMENT                       *
     ***************************************************/

    // PAYMENT INFO: SUPPORTED TOKENS

    function _addSupportedToken(address _token) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        bool added = $._paymentInfo._supportedTokens.add(_token);

        if (!added) {
            revert PaymentTokenAlreadyExists(_token);
        }

        emit PaymentTokenAdded(_token);
    }

    function _removeSupportedToken(address _token) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        bool removed = $._paymentInfo._supportedTokens.remove(_token);

        if (!removed) {
            revert PaymentTokenDoesNotExist(_token);
        }

        emit PaymentTokenRemoved(_token);
    }

    /**
     * @dev Return supported tokens
     */
    function getSupportedTokens() public view virtual returns (address[] memory tokens) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._paymentInfo._supportedTokens.values();
    }

    // PAYMENT INFO: OFF-CHAIN PAYMENT SUPPORT

    function _setOffChainPaymentSupported(bool _supportsOffChainPayment) internal virtual {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        $._paymentInfo._supportsOffChainPayment = _supportsOffChainPayment;
        emit OffChainPaymentSupportUpdated(_supportsOffChainPayment);
    }

    /**
     * @dev Return true if off-chain payment is supported for the given service
     */
    function offChainPaymentSupported() public view virtual returns (bool) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._paymentInfo._supportsOffChainPayment;
    }

    /***************************************************
     *                 PUBLIC KEYS                     *
     ***************************************************/

    /**
     * @dev Check if valid public key use enum
     *
     * This needs to be updated when new enums are added
     */
    function _isValidPublicKeyUse(uint8 use) internal virtual returns (bool) {
        return use < uint(PublicKeyUseType.EncryptPrivateData) + 1;
    }

    /**
     * @dev Add public key with an address
     */
    function _addPublicKey(address pubKeyAddress, bytes memory publicKeyData, uint8 use) internal virtual {
        // Check if {use} is valid enum and revert early if not
        if (!_isValidPublicKeyUse(use)) {
            revert InvalidPublicKeyUseType(use);
        }

        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        bool added = $._publicKeyAddressesSet.add(pubKeyAddress);

        if (!added) {
            revert PublicKeyAlreadyExists(pubKeyAddress);
        }

        PublicKey memory publicKey = PublicKey(PublicKeyUseType(use), publicKeyData);

        $._publicKeys[pubKeyAddress] = publicKey;

        emit PublicKeyAdded(pubKeyAddress, publicKey);
    }

    /**
     * @dev Remove public key by address
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
     * @dev Return all public keys
     */
    function getPublicKeys()
        public
        view
        virtual
        returns (address[] memory pubKeyAddresses, PublicKey[] memory publicKeys)
    {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        address[] memory _pubKeyAddresses = $._publicKeyAddressesSet.values();

        PublicKey[] memory _publicKeys = new PublicKey[](_pubKeyAddresses.length);
        for (uint256 i = 0; i < _pubKeyAddresses.length; i++) {
            _publicKeys[i] = $._publicKeys[_pubKeyAddresses[i]];
        }

        // return addresses and public keys
        return (_pubKeyAddresses, _publicKeys);
    }
}
