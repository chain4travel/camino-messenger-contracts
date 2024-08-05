// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Partner Configuration

pragma solidity ^0.8.24;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

abstract contract PartnerConfiguration is Initializable {
    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    struct Capability {
        string _desc; // TODO: We need explanation and examples here
        string _value; // TODO: We need explanation and examples here
    }

    struct Service {
        uint256 _fee;
        Capability[] _capabilities;
    }

    struct Payment {
        bool _offChain; // Supports off chain payments if true
        address[] _onChainPayments; // Supports on chain payments if true
    }

    /// @custom:storage-location erc7201:camino.messenger.storage.PartnerConfiguration
    struct PartnerConfigurationStorage {
        mapping(bytes32 _serviceHash => Service _service) _supportedServices;
    }

    // keccak256(abi.encode(uint256(keccak256("camino.messenger.storage.PartnerConfiguration")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant PartnerConfigurationStorageLocation =
        0x0c7b73796c7cc89b9f849b9056a93200eba741881e57a1b03b9bedb2c0e07100;

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

    function getService(bytes32 serviceHash) public view virtual returns (Service memory service) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._supportedServices[serviceHash];
    }

    function getServiceCapabilities(
        bytes32 serviceHash
    ) public view virtual returns (Capability[] memory capabilities) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._supportedServices[serviceHash]._capabilities;
    }

    function getServiceFee(bytes32 serviceHash) public view virtual returns (uint256 fee) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._supportedServices[serviceHash]._fee;
    }
}
