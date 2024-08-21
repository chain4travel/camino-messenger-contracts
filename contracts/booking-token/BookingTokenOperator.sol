// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./IBookingToken.sol";

// ERC-20 Utils
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

library BookingTokenOperator {
    using SafeERC20 for IERC20;

    error TokenApprovalFailed(address token, address spender, uint256 amount);

    function _mintBookingToken(
        address bookingToken,
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken
    ) public {
        IBookingToken(bookingToken).safeMintWithReservation(reservedFor, uri, expirationTimestamp, price, paymentToken);
    }

    /**
     * @dev Buy a booking token with the specified price
     */
    function _buyBookingToken(address bookingToken, uint256 tokenId) public {
        // Get the price from the booking token contract
        (uint256 price, IERC20 paymentToken) = IBookingToken(bookingToken).getReservationPrice(tokenId);

        // Check if payment is in native currency or in ERC20
        if (address(paymentToken) == address(0)) {
            // Payment is in native currency. Buy the token sending the payment in
            // native currency to the BookingToken contract.
            IBookingToken(bookingToken).buyReservedToken{ value: price }(tokenId);
        } else {
            // Payment is in ERC20. Approve the BookingToken contract for the
            // reservation price. BookingToken should do the transfer to the
            // supplier.
            bool approval = paymentToken.approve(bookingToken, price);

            if (!approval) {
                revert TokenApprovalFailed(bookingToken, address(paymentToken), price);
            }

            // Buy the token
            IBookingToken(bookingToken).buyReservedToken(tokenId);
        }
    }
}
