// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Service Registry for Partner Configuration

pragma solidity ^0.8.24;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

// TODO: Tidy up. Check CMAccount contracts.

contract ServiceRegistry is Initializable {
    /// @custom:storage-location erc7201:camino.messenger.storage.ServiceRegistry
    struct ServiceRegistryStorage {
        mapping(bytes32 serviceHash => string serviceName) _serviceNameByHash;
        mapping(string serviceName => bytes32 serviceHash) _hashByServiceName;
    }

    event ServiceNameAdded(string serviceName, bytes32 serviceHash);
    event ServiceNameRemoved(string serviceName, bytes32 serviceHash);

    error ServiceAlreadyExists(string serviceName);
    error ServiceDoesNotExist(string serviceName);

    // keccak256(abi.encode(uint256(keccak256("camino.messenger.storage.ServiceRegistry")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant ServiceRegistryStorageLocation =
        0x563e037355fff0507705f481e4b362e4c3996a3b57d07307deabfca3d8168600;

    function _getServiceRegistryStorage() private pure returns (ServiceRegistryStorage storage $) {
        assembly {
            $.slot := ServiceRegistryStorageLocation
        }
    }

    function addServiceName(string memory serviceName) internal virtual {
        bytes32 serviceHash = keccak256(abi.encodePacked(serviceName));

        ServiceRegistryStorage storage $ = _getServiceRegistryStorage();

        if (bytes($._serviceNameByHash[serviceHash]).length != 0) {
            revert ServiceAlreadyExists(serviceName);
        }

        $._serviceNameByHash[serviceHash] = serviceName;
        $._hashByServiceName[serviceName] = serviceHash;

        emit ServiceNameAdded(serviceName, serviceHash);
    }

    function removeServiceName(string memory serviceName) internal virtual {
        bytes32 serviceHash = keccak256(abi.encodePacked(serviceName));

        ServiceRegistryStorage storage $ = _getServiceRegistryStorage();

        if ($._hashByServiceName[serviceName] == bytes32(0)) {
            revert ServiceDoesNotExist(serviceName);
        }

        delete $._serviceNameByHash[serviceHash];
        delete $._hashByServiceName[serviceName];

        emit ServiceNameRemoved(serviceName, serviceHash);
    }

    function getServiceName(bytes32 serviceHash) public view returns (string memory) {
        ServiceRegistryStorage storage $ = _getServiceRegistryStorage();
        return $._serviceNameByHash[serviceHash];
    }

    function getServiceHash(string memory serviceName) public view returns (bytes32) {
        ServiceRegistryStorage storage $ = _getServiceRegistryStorage();
        return $._hashByServiceName[serviceName];
    }
}
