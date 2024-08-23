// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./IBookingToken.sol";

// ERC-20 Utils
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title BookingTokenOperator
 * @notice Booking token operator contract is used by the {CMAccount} contract to mint
 * and buy booking tokens.
 *
 * We made this a library so that we can use it in the {CMAccount} contract without
 * increasing the size of the contract.
 */
library BookingTokenOperator {
    using SafeERC20 for IERC20;

    /**
     * @dev Token approval for the BookingToken address failed.
     *
     * @param token token address
     * @param spender spender address (the BookingToken contract address)
     * @param amount amount of tokens to approve
     */
    error TokenApprovalFailed(address token, address spender, uint256 amount);

    /**
     * @dev Mints a booking token.
     *
     * @param bookingToken booking token contract address
     * @param reservedFor address of the CM Account that can buy the token
     * (generally the distributor)
     * @param uri URI of the token
     * @param expirationTimestamp expiration timestamp of the token in seconds
     * @param price price of the token
     * @param paymentToken payment token address
     */
    function mintBookingToken(
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
     * @dev Buys a booking token with the specified price and payment token in the
     * reservation.
     *
     * @param bookingToken booking token contract address
     * @param tokenId token id
     */
    function buyBookingToken(address bookingToken, uint256 tokenId) public {
        // Get the price from the booking token contract
        (uint256 price, IERC20 paymentToken) = IBookingToken(bookingToken).getReservationPrice(tokenId);

        // Check if payment is in native currency or in ERC20
        if (address(paymentToken) != address(0) && price > 0) {
            // Payment is in ERC20. Approve the BookingToken contract for the
            // reservation price. BookingToken should do the transfer to the
            // supplier.
            bool approval = paymentToken.approve(bookingToken, price);

            if (!approval) {
                revert TokenApprovalFailed(bookingToken, address(paymentToken), price);
            }

            // Buy the token
            IBookingToken(bookingToken).buyReservedToken(tokenId);
        } else {
            // Payment is in native currency. Buy the token sending the payment in
            // native currency to the BookingToken contract.
            IBookingToken(bookingToken).buyReservedToken{ value: price }(tokenId);
        }
    }
}
