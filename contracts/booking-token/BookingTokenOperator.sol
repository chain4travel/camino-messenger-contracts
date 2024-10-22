// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "./IBookingToken.sol";

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
        IERC20 paymentToken,
        bool _isCancellable
    ) public {
        IBookingToken(bookingToken).safeMintWithReservation(
            reservedFor,
            uri,
            expirationTimestamp,
            price,
            paymentToken,
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
            // Payment is in native currency. Buy the token by sending the payment
            // in native currency to the BookingToken contract.
            IBookingToken(bookingToken).buyReservedToken{ value: price }(tokenId);
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
    function initiateCancellationProposal(address bookingToken, uint256 tokenId, uint256 refundAmount) public {
        IBookingToken(bookingToken).initiateCancellationProposal(tokenId, refundAmount);
    }

    /**
     * @notice Accepts a cancellation proposal.
     *
     * @param bookingToken booking token contract address
     * @param tokenId token id
     */
    function acceptCancellationProposal(address bookingToken, uint256 tokenId) public {
        // Get paymentToken and refundAmount
        IERC20 paymentToken = IBookingToken(bookingToken).getReservationPaymentToken(tokenId);
        uint256 refundAmount = IBookingToken(bookingToken).getCancellationProposalRefundAmount(tokenId);

        // Check if payment is in native currency or in ERC20
        if (address(paymentToken) != address(0) && refundAmount > 0) {
            // Payment is in ERC20. Approve the BookingToken contract for the
            // refund amount. BookingToken should do the transfer to the
            // supplier.
            bool approval = paymentToken.approve(bookingToken, refundAmount);

            if (!approval) {
                revert TokenApprovalFailed(bookingToken, address(paymentToken), refundAmount);
            }

            // Accept the cancellation
            IBookingToken(bookingToken).acceptCancellationProposal(tokenId);
        } else {
            // Payment is in native currency. Accept the cancellation by sending the
            // payment in native currency to the BookingToken contract.
            IBookingToken(bookingToken).acceptCancellationProposal{ value: refundAmount }(tokenId);
        }
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
    function acceptCounteredCancellationProposal(address bookingToken, uint256 tokenId) external {
        IBookingToken(bookingToken).acceptCounteredCancellationProposal(tokenId);
    }

    /**
     * @notice Cancels a cancellation proposal.
     *
     * @param bookingToken booking token contract address
     * @param tokenId token id
     */
    function cancelCancellationProposal(address bookingToken, uint256 tokenId) public {
        IBookingToken(bookingToken).cancelCancellationProposal(tokenId);
    }

    /**
     * @dev Gets the status of a cancellation proposal.
     *
     * @param bookingToken booking token contract address
     * @param tokenId token id
     */
    function getCancellationProposalStatus(
        address bookingToken,
        uint256 tokenId
    ) public view returns (uint256 refundAmount, address proposedBy, uint status) {
        return IBookingToken(bookingToken).getCancellationProposalStatus(tokenId);
    }
}
