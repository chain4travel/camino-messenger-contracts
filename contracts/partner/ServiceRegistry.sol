// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Service Registry for Partner Configuration

pragma solidity ^0.8.24;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @title ServiceRegistry
 * @dev Service registry is used by the {CMAccountManager} contract to register
 * services by hashing (keccak256) the service name (string) and creating a mapping
 * as keccak256(serviceName) => serviceName.
 */
abstract contract ServiceRegistry is Initializable {
    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    using EnumerableSet for EnumerableSet.Bytes32Set;

    /// @custom:storage-location erc7201:camino.messenger.storage.ServiceRegistry
    struct ServiceRegistryStorage {
        EnumerableSet.Bytes32Set _servicesHashSet; // set of service hashes
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

    event ServiceRegistered(string serviceName, bytes32 serviceHash);
    event ServiceUnregistered(string serviceName, bytes32 serviceHash);

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    error ServiceAlreadyRegistered(string serviceName);
    error ServiceNotRegistered();

    /***************************************************
     *                 INITIALIZATION                  *
     ***************************************************/

    function __ServiceRegistry_init() internal onlyInitializing {}

    function __ServiceRegistry_init_unchained() internal onlyInitializing {}

    /***************************************************
     *                    FUNCS                        *
     ***************************************************/

    /**
     * @dev Adds a new service by its name. This function calculates the hash of the
     * service name and adds it to the registry
     *
     * {serviceName} is the pkg + service name as:
     *
     *  ┌────────────── pkg ─────────────┐ ┌───── service name ─────┐
     * "cmp.services.accommodation.v1alpha.AccommodationSearchService"
     *
     * @param serviceName Name of the service
     */
    function _registerServiceName(string memory serviceName) internal virtual {
        bytes32 serviceHash = keccak256(abi.encodePacked(serviceName));

        ServiceRegistryStorage storage $ = _getServiceRegistryStorage();

        // Try to add the hash to the service set
        bool added = $._servicesHashSet.add(serviceHash);

        if (!added) {
            revert ServiceAlreadyRegistered(serviceName);
        }

        $._serviceNameByHash[serviceHash] = serviceName;
        $._hashByServiceName[serviceName] = serviceHash;

        emit ServiceRegistered(serviceName, serviceHash);
    }

    /**
     * @dev Removes a service by its name. This function calculates the hash of the
     * service name and removes it from the registry.
     *
     * @param serviceName Name of the service
     */
    function _unregisterServiceName(string memory serviceName) internal virtual {
        bytes32 serviceHash = keccak256(abi.encodePacked(serviceName));

        ServiceRegistryStorage storage $ = _getServiceRegistryStorage();

        // Try to remove the hash to the service set
        bool removed = $._servicesHashSet.remove(serviceHash);

        if (!removed) {
            revert ServiceNotRegistered();
        }

        delete $._serviceNameByHash[serviceHash];
        delete $._hashByServiceName[serviceName];

        emit ServiceUnregistered(serviceName, serviceHash);
    }

    /**
     * @dev Returns the name of a service by its hash.
     *
     * @param serviceHash Hash of the service
     */
    function getRegisteredServiceNameByHash(bytes32 serviceHash) public view returns (string memory serviceName) {
        ServiceRegistryStorage storage $ = _getServiceRegistryStorage();

        // Check if the service is registered
        if (!$._servicesHashSet.contains(serviceHash)) {
            revert ServiceNotRegistered();
        }
        return $._serviceNameByHash[serviceHash];
    }

    /**
     * @dev Returns the hash of a service by its name.
     *
     * @param serviceName Name of the service
     */
    function getRegisteredServiceHashByName(string memory serviceName) public view returns (bytes32 serviceHash) {
        ServiceRegistryStorage storage $ = _getServiceRegistryStorage();

        // Check if the service is registered
        if (!$._servicesHashSet.contains(keccak256(abi.encodePacked(serviceName)))) {
            revert ServiceNotRegistered();
        }

        return $._hashByServiceName[serviceName];
    }

    /**
     * @dev Returns all registered service hashes.
     */
    function getAllRegisteredServiceHashes() public view returns (bytes32[] memory services) {
        ServiceRegistryStorage storage $ = _getServiceRegistryStorage();
        return $._servicesHashSet.values();
    }

    /**
     * @dev Returns all registered service names.
     */
    function getAllRegisteredServiceNames() public view returns (string[] memory services) {
        ServiceRegistryStorage storage $ = _getServiceRegistryStorage();

        // Get all hashes and create a list with predefined length
        bytes32[] memory serviceHashes = $._servicesHashSet.values();
        string[] memory serviceNames = new string[](serviceHashes.length);

        // Get all names for the hashes
        for (uint256 i = 0; i < serviceHashes.length; i++) {
            serviceNames[i] = $._serviceNameByHash[serviceHashes[i]];
        }

        return serviceNames;
    }
}
