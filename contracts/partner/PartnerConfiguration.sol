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
        string[] _capabilities;
    }

    struct PaymentInfo {
        bool _supportsOffChainPayment; // Supports off chain payments if true
        EnumerableSet.AddressSet _supportedTokens; // Supported on-chain token for payment
    }

    /// @custom:storage-location erc7201:camino.messenger.storage.PartnerConfiguration
    struct PartnerConfigurationStorage {
        EnumerableSet.Bytes32Set _servicesHashSet;
        mapping(bytes32 _serviceHash => Service _service) _supportedServices;
        PaymentInfo _paymentInfo;
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

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    event ServiceAdded(bytes32 serviceHash);
    event ServiceRemoved(bytes32 serviceHash);

    event ServiceFeeUpdated(bytes32 serviceHash, uint256 fee);
    event ServiceCapabilityAdded(bytes32 serviceHash, string capability);
    event ServiceCapabilityRemoved(bytes32 serviceHash, string capability);

    /***************************************************
     *                 INITIALIZATION                  *
     ***************************************************/

    function __PartnerConfiguration_init() internal onlyInitializing {}

    function __PartnerConfiguration_init_unchained() internal onlyInitializing {}

    /***************************************************
     *                    FUNCS                        *
     ***************************************************/

    /**
     * @dev Adds a Service object for a given hash.
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
     * @dev Removes a Service object for a given hash.
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
     * @dev Returns the Service object for a given hash.
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

    // TODO: Add getter for "all services"

    function getServiceCapabilities(bytes32 serviceHash) public view virtual returns (string[] memory capabilities) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        if (!$._servicesHashSet.contains(serviceHash)) {
            revert ServiceDoesNotExist(serviceHash);
        }

        return $._supportedServices[serviceHash]._capabilities;
    }

    function getServiceFee(bytes32 serviceHash) public view virtual returns (uint256 fee) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();

        // Check if the service exists
        if (!$._servicesHashSet.contains(serviceHash)) {
            revert ServiceDoesNotExist(serviceHash);
        }

        return $._supportedServices[serviceHash]._fee;
    }
}
