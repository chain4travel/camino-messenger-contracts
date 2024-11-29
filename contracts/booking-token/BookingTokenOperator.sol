// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.24;

import { IBookingToken, IERC20, CancellationProposalStatus } from "./IBookingToken.sol";

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
     *                   ERRORS                        *
     ***************************************************/

    /**
     * @dev Token approval for the BookingToken address failed.
     *
     * @param token token address
     * @param spender spender address (the BookingToken contract address)
     * @param amount amount of tokens to approve
     */
    error TokenApprovalFailed(address token, address spender, uint256 amount);

    /***************************************************
     *                   FUNCS                         *
     ***************************************************/

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
        IERC20 paymentToken,
        uint256 offchainPaymentCurrency,
        bool _isCancellable
    ) public {
        IBookingToken(bookingToken).safeMintWithReservation(
            reservedFor,
            uri,
            expirationTimestamp,
            price,
            paymentToken,
            offchainPaymentCurrency,
            _isCancellable
        );
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
            bool approval = paymentToken.approve(bookingToken, price);

            if (!approval) {
                revert TokenApprovalFailed(bookingToken, address(paymentToken), price);
            }

            // Buy the token
            IBookingToken(bookingToken).buyReservedToken(tokenId);
        }
    }

    /**
     * @notice Sets the cancellable flag for a token. This can only be called by the
     * supplier of the token.
     * @param tokenId The token id
     * @param _isCancellable The new cancellable flag
     */
    function setCancellable(address bookingToken, uint256 tokenId, bool _isCancellable) external {
        IBookingToken(bookingToken).setCancellable(tokenId, _isCancellable);
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

    /**
     * @notice Initiates a cancellation proposal for a bought token.
     *
     * @param bookingToken booking token contract address
     * @param tokenId token id
     * @param refundAmount proposed refund amount
     */
    function initiateCancellationProposal(
        address bookingToken,
        uint256 tokenId,
        uint256 refundAmount,
        uint16 cancellationReason,
        uint16 cancellationReasonVersion
    ) public {
        IBookingToken(bookingToken).initiateCancellationProposal(
            tokenId,
            refundAmount,
            cancellationReason,
            cancellationReasonVersion
        );
    }

    /**
     * @notice Accepts a cancellation proposal.
     *
     * @param bookingToken booking token contract address
     * @param tokenId token id
     */
    function acceptCancellationProposal(address bookingToken, uint256 tokenId, uint256 checkRefundAmount) public {
        // Get paymentToken and refundAmount
        IERC20 paymentToken = IBookingToken(bookingToken).getReservationPaymentToken(tokenId);
        uint256 refundAmount = IBookingToken(bookingToken).getCancellationProposalRefundAmount(tokenId);

        // Check if payment is in native currency or in ERC20
        if (address(paymentToken) == NATIVE_PAYMENT) {
            // Payment is in native currency. Accept the cancellation by sending the
            // payment in native currency to the BookingToken contract.
            IBookingToken(bookingToken).acceptCancellationProposal{ value: refundAmount }(tokenId, checkRefundAmount);
        } else if (address(paymentToken) == OFFCHAIN_PAYMENT) {
            // Off-chain payment - no on-chain transfer needed
            IBookingToken(bookingToken).acceptCancellationProposal(tokenId, checkRefundAmount);
        } else {
            // Payment is in ERC20. Approve the BookingToken contract for the
            // refund amount. BookingToken should do the transfer to the
            // supplier.
            bool approval = paymentToken.approve(bookingToken, refundAmount);

            if (!approval) {
                revert TokenApprovalFailed(bookingToken, address(paymentToken), refundAmount);
            }

            // Accept the cancellation
            IBookingToken(bookingToken).acceptCancellationProposal(tokenId, checkRefundAmount);
        }
    }

    /**
     * @notice Reject a cancellation proposal for a bought token.
     *
     * @param bookingToken booking token contract address
     * @param tokenId The token id to reject the cancellation for
     * @param rejectionReason The reason for rejecting the cancellation
     * @param rejectionReasonVersion Version of the rejection reason enum from the CMP
     */
    function rejectCancellationProposal(
        address bookingToken,
        uint256 tokenId,
        uint16 rejectionReason,
        uint16 rejectionReasonVersion
    ) external {
        IBookingToken(bookingToken).rejectCancellationProposal(tokenId, rejectionReason, rejectionReasonVersion);
    }

    /**
     * @notice Counters a cancellation proposal.
     *
     * @param bookingToken booking token contract address
     * @param tokenId token id
     * @param refundAmount proposed refund amount
     */
    function counterCancellationProposal(address bookingToken, uint256 tokenId, uint256 refundAmount) public {
        IBookingToken(bookingToken).counterCancellationProposal(tokenId, refundAmount);
    }

    /**
     * @notice Accepts a countered cancellation proposal.
     *
     * @param bookingToken booking token contract address
     * @param tokenId token id
     */
    function acceptCounteredCancellationProposal(
        address bookingToken,
        uint256 tokenId,
        uint256 checkRefundAmount
    ) external {
        IBookingToken(bookingToken).acceptCounteredCancellationProposal(tokenId, checkRefundAmount);
    }

    /**
     * @notice Withdraws a cancellation proposal.
     *
     * @param bookingToken booking token contract address
     * @param tokenId token id for which to withdraw the proposal
     * @param reason The reason for withdrawing the proposal
     * @param reasonVersion The version of the withdrawal reason from the CMP
     */
    function withdrawCancellationProposal(
        address bookingToken,
        uint256 tokenId,
        uint16 reason,
        uint16 reasonVersion
    ) public {
        IBookingToken(bookingToken).withdrawCancellationProposal(tokenId, reason, reasonVersion);
    }
}
