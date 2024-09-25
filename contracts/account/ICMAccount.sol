// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.24;

interface ICMAccount {
    function initialize(
        address manager,
        address bookingToken,
        uint256 prefundAmount,
        address owner,
        address upgrader
    ) external;
}
