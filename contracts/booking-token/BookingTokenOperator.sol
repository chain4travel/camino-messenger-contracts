// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./IBookingToken.sol";

abstract contract BookingTokenOperator {
    function _mintBookingToken(
        address bookingToken,
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price
    ) internal virtual {
        IBookingToken(bookingToken).safeMintWithReservation(reservedFor, uri, expirationTimestamp, price);
    }

    /**
     * @dev Buy a booking token with the specified price
     */
    function _buyBookingToken(address bookingToken, uint256 tokenId) internal virtual {
        // Get the price from the booking token contract
        uint256 price = _getTokenReservationPrice(bookingToken, tokenId);
        // Buy the token
        IBookingToken(bookingToken).buyReservedToken{ value: price }(tokenId);
    }

    /**
     * @dev Get the price of a booking token
     */
    function _getTokenReservationPrice(address bookingToken, uint256 tokenId) public view returns (uint256) {
        return IBookingToken(bookingToken).getReservationPrice(tokenId);
    }

    /**
     * @dev Mint a booking token
     *
     * This function should be overridden by the implementation
     */
    function mintBookingToken(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price
    ) external virtual;

    /**
     * @dev Buy a booking token
     *
     * This function should be overridden by the implementation
     */
    function buyBookingToken(uint256 tokenId) external virtual;
}