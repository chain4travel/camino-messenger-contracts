// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Partner Configuration

pragma solidity ^0.8.24;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

abstract contract PartnerConfiguration is Initializable {
    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    struct Service {
        uint256 _fee;
        string[] _capabilities;
    }

    struct PaymentInfo {
        bool _supportsOffChainPayment; // Supports off chain payments if true
        address[] _supportedTokens; // Supported on-chain token for payment
    }

    /// @custom:storage-location erc7201:camino.messenger.storage.PartnerConfiguration
    struct PartnerConfigurationStorage {
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

    // TODO: Errors here

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    // TODO: Events here

    /***************************************************
     *                 INITIALIZATION                  *
     ***************************************************/

    function __PartnerConfiguration_init() internal onlyInitializing {}

    function __PartnerConfiguration_init_unchained() internal onlyInitializing {}

    /***************************************************
     *                    FUNCS                        *
     ***************************************************/

    /**
     * @dev Set the Service object for a given hash.
     *
     * @param serviceHash Hash of the service
     * @param service Service object
     */
    function _setService(bytes32 serviceHash, Service memory service) internal {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        $._supportedServices[serviceHash] = service;
    }

    /**
     * @dev Set the Service fee for a given hash. Note that this does not check if
     * the service exists. So it will create an empty service if it does not exist.
     *
     * @param serviceHash Hash of the service
     * @param fee Fee
     */
    function _setServiceFee(bytes32 serviceHash, uint256 fee) internal {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        $._supportedServices[serviceHash]._fee = fee;
    }

    /**
     * @dev Set the Service capabilities for a given hash. Note that this does not check if
     * the service exists. So it will create an empty service if it does not exist.
     *
     * @param serviceHash Hash of the service
     * @param capabilities Capabilities
     */
    function _setServiceCapabilities(bytes32 serviceHash, string[] memory capabilities) internal {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        $._supportedServices[serviceHash]._capabilities = capabilities;
    }

    /**
     * @dev Add a capability to the service. Note that this does not check if
     * the service exists. So it will create an empty service if it does not exist.
     *
     * @param serviceHash Hash of the service
     * @param capability Capability
     */
    function _addServiceCapability(bytes32 serviceHash, string memory capability) internal {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        $._supportedServices[serviceHash]._capabilities.push(capability);
    }

    /**
     * @dev Remove a capability from the service.
     *
     * @param serviceHash Hash of the service
     * @param capability Capability
     */
    function _removeServiceCapability(bytes32 serviceHash, string memory capability) internal {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
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
     * @dev Returns the Service object for a given hash.
     *
     * {serviceHash} is keccak256 hash of the pkg + service name as:
     *
     *            |-------------- pkg -------------| |----- service name -----|
     * keccak256("cmp.services.accommodation.v1alpha.AccommodationSearchService")
     */
    function getService(bytes32 serviceHash) public view virtual returns (Service memory service) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._supportedServices[serviceHash];
    }

    function getServiceCapabilities(bytes32 serviceHash) public view virtual returns (string[] memory capabilities) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._supportedServices[serviceHash]._capabilities;
    }

    function getServiceFee(bytes32 serviceHash) public view virtual returns (uint256 fee) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._supportedServices[serviceHash]._fee;
    }
}
