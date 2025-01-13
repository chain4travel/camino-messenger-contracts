// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.24;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

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

    function getReservationPaymentToken(uint256 tokenId) external view returns (IERC20 paymentToken);

    /**
     * @notice Record expiration status if the token is expired
     * @param tokenId The token id to record as expired
     */
    function recordExpiration(uint256 tokenId) external;

    /***************************************************
     *              CANCELLATION LOGIC                 *
     ***************************************************/

    /**
     * @notice Initiates a cancellation for a bought token.
     *
     * @param tokenId The token id to initiate the cancellation for
     * @param refundAmount The proposed refund amount in wei
     * @param cancellationReason The reason for cancellation
     * @param cancellationReasonVersion The version of the cancellation reason
     */
    function initiateCancellation(
        uint256 tokenId,
        uint256 refundAmount,
        uint16 cancellationReason,
        uint16 cancellationReasonVersion
    ) external;

    /**
     * @notice Sets accepted by the owner or supplier flag for a cancellation proposal for a bought token.
     *
     * @param tokenId The token id to accept the cancellation for
     * @param refundAmount The refund amount to check, this is to prevent front-running attacks
     */
    function acceptCancellation(uint256 tokenId, uint256 refundAmount) external;

    /**
     * @notice Reject a cancellation proposal for a bought token.
     *
     * @param tokenId The token id to reject the cancellation for
     * @param rejectionReason The reason for rejection
     * @param rejectionReasonVersion The version of the rejection reason
     */
    function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) external;

    /**
     * @notice Counters a cancellation proposal with a new proposal.
     *
     * @param tokenId The token id to counter the cancellation for
     * @param refundAmount The refund amount to check, this is to prevent front-running attacks
     * @param counterReason The reason for the counter
     * @param counterReasonVersion The version of the counter reason
     */
    function counterCancellation(
        uint256 tokenId,
        uint256 refundAmount,
        uint16 counterReason,
        uint16 counterReasonVersion
    ) external;

    /**
     * @notice Withdraws an active cancellation proposal. Only the current proposer of the proposal can withdraw.
     *
     * @param tokenId The token id for which to withdraw the proposal
     * @param withdrawalReason The reason for withdrawing the proposal
     * @param withdrawalReasonVersion The version of the withdrawal reason
     */
    function withdrawCancellation(uint256 tokenId, uint16 withdrawalReason, uint16 withdrawalReasonVersion) external;

    /**
     * @notice Finalizes a cancellation proposal. Only the supplier of the token can finalize.
     *
     * @param tokenId The token id for which to finalize the proposal
     * @param refundAmount The refund amount to check, this is to prevent front-running attacks
     */
    function finalizeCancellation(uint256 tokenId, uint256 refundAmount) external payable;
}
