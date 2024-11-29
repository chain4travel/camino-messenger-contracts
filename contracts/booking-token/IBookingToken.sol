// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.24;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

enum CancellationProposalStatus {
    NoProposal, // 0, default
    Pending, // 1
    Rejected, // 2
    Countered, // 3
    Accepted // 4
}

interface IBookingToken {
    function safeMintWithReservation(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken,
        uint256 offchainPaymentCurrency,
        bool isCancellable
    ) external;

    function buyReservedToken(uint256 tokenId) external payable;

    function getReservationPrice(uint256 tokenId) external view returns (uint256 price, IERC20 paymentToken);

    function getCancellationProposalRefundAmount(uint256 tokenId) external view returns (uint256 refundAmount);

    function getReservationPaymentToken(uint256 tokenId) external view returns (IERC20 paymentToken);

    /**
     * @notice Sets the cancellable flag for a token. This can only be called by the
     * supplier of the token.
     * @param tokenId The token id
     * @param _isCancellable The new cancellable flag
     */
    function setCancellable(uint256 tokenId, bool _isCancellable) external;

    /**
     * @notice Record expiration status if the token is expired
     * @param tokenId The token id to record as expired
     */
    function recordExpiration(uint256 tokenId) external;

    /**
     * @notice Initiates a cancellation for a bought token.
     *
     * @param tokenId The token id to initiate the cancellation for
     * @param refundAmount The proposed refund amount in wei
     * @param cancellationReason The reason for cancellation
     * @param cancellationReasonVersion The version of the cancellation reason
     */
    function initiateCancellationProposal(
        uint256 tokenId,
        uint256 refundAmount,
        uint16 cancellationReason,
        uint16 cancellationReasonVersion
    ) external;

    /**
     * @notice Accepts a cancellation proposal for a bought token.
     *
     * @param tokenId The token id to accept the cancellation for
     * @param checkRefundAmount The refund amount to check, this is to prevent front-running attacks
     */
    function acceptCancellationProposal(uint256 tokenId, uint256 checkRefundAmount) external payable;

    /**
     * @notice Reject a cancellation proposal for a bought token.
     *
     * @param tokenId The token id to reject the cancellation for
     * @param rejectionReason The reason for rejection
     * @param rejectionReasonVersion The version of the rejection reason
     */
    function rejectCancellationProposal(
        uint256 tokenId,
        uint16 rejectionReason,
        uint16 rejectionReasonVersion
    ) external;

    /**
     * @notice Counters a cancellation proposal with a new proposal.
     *
     * @param tokenId The token id to counter the cancellation for
     * @param newRefundAmount The new proposed refund amount in wei
     */
    function counterCancellationProposal(uint256 tokenId, uint256 newRefundAmount) external;

    /**
     * @notice Accept a countered cancellation proposal
     * @param tokenId The token id to accept the countered cancellation proposal for
     * @param checkRefundAmount The refund amount to check, this is to prevent front-running attacks
     */
    function acceptCounteredCancellationProposal(uint256 tokenId, uint256 checkRefundAmount) external;

    /**
     * @notice Withdraws an active cancellation proposal. Only the initiator can withdraw.
     *
     * @param tokenId The token id for which to withdraw the proposal
     * @param reason The reason for withdrawing the proposal
     * @param reasonVersion The version of the withdrawal reason from the CMP
     */
    function withdrawCancellationProposal(uint256 tokenId, uint16 reason, uint16 reasonVersion) external;
}
