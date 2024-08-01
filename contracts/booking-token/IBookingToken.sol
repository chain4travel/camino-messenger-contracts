// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IBookingToken {
    function safeMintWithReservation(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price
    ) external;

    function buyReservedToken(uint256 tokenId) external payable;

    function getReservationPrice(uint256 tokenId) external view returns (uint256);
}
