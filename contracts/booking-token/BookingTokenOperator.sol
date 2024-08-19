// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./IBookingToken.sol";

// ERC-20 Utils
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

abstract contract BookingTokenOperator {
    using SafeERC20 for IERC20;

    function _mintBookingToken(
        address bookingToken,
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken
    ) internal virtual {
        IBookingToken(bookingToken).safeMintWithReservation(reservedFor, uri, expirationTimestamp, price, paymentToken);
    }

    /**
     * @dev Buy a booking token with the specified price
     */
    function _buyBookingToken(address bookingToken, uint256 tokenId) internal virtual {
        // Get the price from the booking token contract
        (uint256 price, IERC20 paymentToken) = _getTokenReservationPrice(bookingToken, tokenId);

        // Check if payment is in native currency or in ERC20
        if (address(paymentToken) == address(0)) {
            // Payment is in native currency. Buy the token sending the payment in
            // native currency to the BookingToken contract.
            IBookingToken(bookingToken).buyReservedToken{ value: price }(tokenId);
        } else {
            // Payment is in ERC20. Approve the BookingToken contract for the
            // reservation price. BookingToken should do the transfer to the
            // supplier.
            paymentToken.approve(bookingToken, price);

            // Buy the token
            IBookingToken(bookingToken).buyReservedToken(tokenId);
        }
    }

    /**
     * @dev Get the price of a booking token
     */
    function _getTokenReservationPrice(
        address bookingToken,
        uint256 tokenId
    ) public view returns (uint256 price, IERC20 paymentToken) {
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
        uint256 price,
        IERC20 paymentToken
    ) external virtual;

    /**
     * @dev Buy a booking token
     *
     * This function should be overridden by the implementation
     */
    function buyBookingToken(uint256 tokenId) external virtual;
}
