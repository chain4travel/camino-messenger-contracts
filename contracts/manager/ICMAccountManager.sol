// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.24;

interface ICMAccountManager {
    function getAccountImplementation() external view returns (address);

    function getDeveloperFeeBp() external view returns (uint256);

    function getDeveloperWallet() external view returns (address);

    function isCMAccount(address account) external view returns (bool);

    function getRegisteredServiceHashByName(string memory serviceName) external view returns (bytes32 serviceHash);

    function getRegisteredServiceNameByHash(bytes32 serviceHash) external view returns (string memory serviceName);
}
