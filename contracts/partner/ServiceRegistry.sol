// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Service Registry for Partner Configuration

pragma solidity ^0.8.24;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

abstract contract ServiceRegistry is Initializable {
    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    /// @custom:storage-location erc7201:camino.messenger.storage.ServiceRegistry
    struct ServiceRegistryStorage {
        mapping(bytes32 serviceHash => string serviceName) _serviceNameByHash;
        mapping(string serviceName => bytes32 serviceHash) _hashByServiceName;
    }

    // keccak256(abi.encode(uint256(keccak256("camino.messenger.storage.ServiceRegistry")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant ServiceRegistryStorageLocation =
        0x563e037355fff0507705f481e4b362e4c3996a3b57d07307deabfca3d8168600;

    function _getServiceRegistryStorage() private pure returns (ServiceRegistryStorage storage $) {
        assembly {
            $.slot := ServiceRegistryStorageLocation
        }
    }

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    event ServiceNameAdded(string serviceName, bytes32 serviceHash);
    event ServiceNameRemoved(string serviceName, bytes32 serviceHash);

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    error ServiceAlreadyExists(string serviceName);
    error ServiceDoesNotExist(string serviceName);

    /***************************************************
     *                 INITIALIZATION                  *
     ***************************************************/

    function __ServiceRegistry_init() internal onlyInitializing {}

    function __ServiceRegistry_init_unchained() internal onlyInitializing {}

    /***************************************************
     *                    FUNCS                        *
     ***************************************************/

    /**
     * @dev Add a new service by its name. This function calculates the hash of the
     * service name and adds it to the registry
     *
     * @param serviceName Name of the service
     */
    function _addServiceName(string memory serviceName) internal virtual {
        bytes32 serviceHash = keccak256(abi.encodePacked(serviceName));

        ServiceRegistryStorage storage $ = _getServiceRegistryStorage();

        if (bytes($._serviceNameByHash[serviceHash]).length != 0) {
            revert ServiceAlreadyExists(serviceName);
        }

        $._serviceNameByHash[serviceHash] = serviceName;
        $._hashByServiceName[serviceName] = serviceHash;

        emit ServiceNameAdded(serviceName, serviceHash);
    }

    /**
     * @dev Remove a service by its name. This function calculates the hash of the
     * service name and removes it from the registry
     *
     * @param serviceName Name of the service
     */
    function _removeServiceName(string memory serviceName) internal virtual {
        bytes32 serviceHash = keccak256(abi.encodePacked(serviceName));

        ServiceRegistryStorage storage $ = _getServiceRegistryStorage();

        if ($._hashByServiceName[serviceName] == bytes32(0)) {
            revert ServiceDoesNotExist(serviceName);
        }

        delete $._serviceNameByHash[serviceHash];
        delete $._hashByServiceName[serviceName];

        emit ServiceNameRemoved(serviceName, serviceHash);
    }

    /**
     * @dev Get the name of a service by its hash
     *
     * @param serviceHash Hash of the service
     */
    function getServiceName(bytes32 serviceHash) public view returns (string memory) {
        ServiceRegistryStorage storage $ = _getServiceRegistryStorage();
        return $._serviceNameByHash[serviceHash];
    }

    /**
     * @dev Get the hash of a service by its name
     *
     * @param serviceName Name of the service
     */
    function getServiceHash(string memory serviceName) public view returns (bytes32) {
        ServiceRegistryStorage storage $ = _getServiceRegistryStorage();
        return $._hashByServiceName[serviceName];
    }
}
