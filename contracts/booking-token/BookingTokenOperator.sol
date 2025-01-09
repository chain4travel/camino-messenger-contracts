// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.24;

import { IBookingToken, IERC20 } from "./IBookingToken.sol";
import { CancellationProposalStatus } from "./BookingTokenCancellable.sol";

// ERC-20 Utils
import { SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

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

    /***************************************************
     *                   CONSTANTS                     *
     ***************************************************/

    /**
     * @dev Special address for native payments.
     * @notice Tokens are directly transferred to the recipient.
     */
    address public constant NATIVE_PAYMENT = address(0);

    /**
     * @dev Special address for offchain payments.
     * @notice A third-party service is used to handle payments.
     */
    address public constant OFFCHAIN_PAYMENT = address(1);

    /***************************************************
     *                   FUNCS                         *
     ***************************************************/

    error UnexpectedPrice(uint256 tokenId, uint256 actualPrice, uint256 expectedPrice);

    error UnexpectedPaymentToken(uint256 tokenId, IERC20 actualPaymentToken, IERC20 expectedPaymentToken);

    /***************************************************
     *                   FUNCS                         *
     ***************************************************/

    /**
     * @dev Mints a booking token with offchain payment currency and cancellable support.
     *
     * @param bookingToken booking token contract address
     * @param reservedFor address of the CM Account that can buy the token
     * (generally the distributor)
     * @param uri URI of the token
     * @param expirationTimestamp expiration timestamp of the token in seconds
     * @param price price of the token
     * @param paymentToken payment token address
     * @param offchainPaymentCurrency payment token address
     * @param cancellable cancellable flag
     */
    function mintBookingToken(
        address bookingToken,
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken,
        uint256 offchainPaymentCurrency,
        bool cancellable
    ) public {
        IBookingToken(bookingToken).safeMintWithReservation(
            reservedFor,
            uri,
            expirationTimestamp,
            price,
            paymentToken,
            offchainPaymentCurrency,
            cancellable
        );
    }

    /**
     * @dev Buys a booking token with the specified price and payment token in the
     * reservation.
     *
     * @param bookingToken booking token contract address
     * @param tokenId token id
     */
    function buyBookingToken(
        address bookingToken,
        uint256 tokenId,
        uint256 expectedPrice,
        IERC20 expectedPaymentToken
    ) public {
        // Get the price from the booking token contract
        (uint256 price, IERC20 paymentToken) = IBookingToken(bookingToken).getReservationPrice(tokenId);

        // Check if the price is correct
        if (price != expectedPrice) {
            revert UnexpectedPrice(tokenId, price, expectedPrice);
        }

        // Check if the payment token is correct
        if (address(paymentToken) != address(expectedPaymentToken)) {
            revert UnexpectedPaymentToken(tokenId, paymentToken, expectedPaymentToken);
        }

        if (address(paymentToken) == NATIVE_PAYMENT) {
            // Payment is in native currency. Buy the token by sending the payment
            // in native currency to the BookingToken contract.
            IBookingToken(bookingToken).buyReservedToken{ value: price }(tokenId);
        } else if (address(paymentToken) == OFFCHAIN_PAYMENT) {
            // Off-chain payment - no on-chain transfer needed
            IBookingToken(bookingToken).buyReservedToken(tokenId);
        } else {
            // Payment is in ERC20. Approve the BookingToken contract for the
            // reservation price. BookingToken should do the transfer to the
            // supplier.
            paymentToken.forceApprove(bookingToken, price);

            // Buy the token
            IBookingToken(bookingToken).buyReservedToken(tokenId);
        }
    }

    /**
     * @notice Record the expiration of a booking token.
     *
     * @param bookingToken booking token contract address
     * @param tokenId token id
     */
    function recordExpiration(address bookingToken, uint256 tokenId) public {
        IBookingToken(bookingToken).recordExpiration(tokenId);
    }

    /***************************************************
     *              CANCELLATION LOGIC                 *
     ***************************************************/

    /**
     * @notice Initiates a cancellation proposal for a bought token.
     *
     * @param bookingToken booking token contract address
     * @param tokenId token id
     * @param refundAmount proposed refund amount
     * @param cancellationReason cancellation reason
     * @param cancellationReasonVersion cancellation reason version
     */
    function initiateCancellation(
        address bookingToken,
        uint256 tokenId,
        uint256 refundAmount,
        uint16 cancellationReason,
        uint16 cancellationReasonVersion
    ) external {
        IBookingToken(bookingToken).initiateCancellation(
            tokenId,
            refundAmount,
            cancellationReason,
            cancellationReasonVersion
        );
    }

    /**
     * @notice Sets accepted by the owner or supplier flag for a cancellation proposal for a bought token.
     *
     * @param tokenId The token id to accept the cancellation for
     * @param refundAmount The refund amount to check, this is to prevent front-running attacks
     */
    function acceptCancellation(address bookingToken, uint256 tokenId, uint256 refundAmount) external {
        IBookingToken(bookingToken).acceptCancellation(tokenId, refundAmount);
    }

    /**
     * @notice Counters a cancellation proposal.
     *
     * @param bookingToken booking token contract address
     * @param tokenId token id
     * @param refundAmount proposed refund amount
     */
    function counterCancellation(
        address bookingToken,
        uint256 tokenId,
        uint256 refundAmount,
        uint16 counterReason,
        uint16 counterReasonVersion
    ) public {
        IBookingToken(bookingToken).counterCancellation(tokenId, refundAmount, counterReason, counterReasonVersion);
    }

    /**
     * @notice Withdraws a cancellation proposal.
     *
     * @param bookingToken booking token contract address
     * @param tokenId token id for which to withdraw the proposal
     * @param reason The reason for withdrawing the proposal
     * @param reasonVersion The version of the withdrawal reason from the CMP
     */
    function withdrawCancellation(address bookingToken, uint256 tokenId, uint16 reason, uint16 reasonVersion) public {
        IBookingToken(bookingToken).withdrawCancellation(tokenId, reason, reasonVersion);
    }

    /**
     * @notice Reject a cancellation proposal for a bought token.
     *
     * @param bookingToken booking token contract address
     * @param tokenId The token id to reject the cancellation for
     * @param rejectionReason The reason for rejecting the cancellation
     * @param rejectionReasonVersion Version of the rejection reason enum from the CMP
     */
    function rejectCancellation(
        address bookingToken,
        uint256 tokenId,
        uint16 rejectionReason,
        uint16 rejectionReasonVersion
    ) external {
        IBookingToken(bookingToken).rejectCancellation(tokenId, rejectionReason, rejectionReasonVersion);
    }

    /**
     * @notice Finalizes a cancellation proposal by transferring the refund amount
     * to the Booking Token contract.
     *
     * @param bookingToken BookingToken contract address
     * @param tokenId The token id for which to finalize the proposal
     * @param refundAmount The refund amount to check, this is to prevent front-running attacks
     */
    function finalizeCancellation(address bookingToken, uint256 tokenId, uint256 refundAmount) public {
        IERC20 paymentToken = IBookingToken(bookingToken).getReservationPaymentToken(tokenId);

        // Check if payment is in native currency or in ERC20
        if (address(paymentToken) == NATIVE_PAYMENT) {
            // Payment is in native currency. Finalize the cancellation by sending
            // the payment in native currency to the BookingToken contract.
            IBookingToken(bookingToken).finalizeCancellation{ value: refundAmount }(tokenId, refundAmount);
        } else if (address(paymentToken) == OFFCHAIN_PAYMENT) {
            // Off-chain payment - no on-chain transfer needed
            IBookingToken(bookingToken).finalizeCancellation(tokenId, refundAmount);
        } else {
            // Payment is in ERC20. Approve the BookingToken contract for the
            // refund amount. BookingToken should do the transfer to the
            // supplier.
            paymentToken.forceApprove(bookingToken, refundAmount);

            // Accept the cancellation
            IBookingToken(bookingToken).finalizeCancellation(tokenId, refundAmount);
        }
    }
}
