// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

enum CancellationProposalStatus {
    NoProposal, // 0, default
    Pending, // 1
    Rejected, // 2
    Countered, // 3
    Accepted // 4
}

enum CancellationRejectionReason {
    Unspecified, // 0, default
    TechnicalError, // 1
    InvalidServiceOrBookingReference, // 2
    BookingIsAlreadyCancelled, // 3
    ServiceHasStartedOrHasBeenDelivered, // 4
    CancellationWindowExpired, // 5
    ServiceCannotBeCancelledOnline, // 6
    RateOrFareCannotBeCancelled, // 7
    EntirePackageMustBeCancelled, // 8, service forms part of a package, the entire package must be cancelled
    RefundCurrencyNotSupported // 9
}

interface IBookingToken {
    function safeMintWithReservation(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken,
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
     */
    function initiateCancellationProposal(uint256 tokenId, uint256 refundAmount) external;

    /**
     * @notice Accepts a cancellation proposal for a bought token.
     *
     * @param tokenId The token id to accept the cancellation for
     */
    function acceptCancellationProposal(uint256 tokenId, uint256 checkRefundAmount) external payable;

    /**
     * @notice Reject a cancellation proposal for a bought token.
     *
     * @param tokenId The token id to reject the cancellation for
     * @param reason The reason for rejecting the cancellation
     */
    function rejectCancellationProposal(uint256 tokenId, CancellationRejectionReason reason) external;

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
     * @param checkRefundAmount The refund amount to check against the proposal
     */
    function acceptCounteredCancellationProposal(uint256 tokenId, uint256 checkRefundAmount) external;

    /**
     * @notice Cancels an active cancellation proposal. Only the initiator can cancel.
     *
     * @param tokenId The token id for which to cancel the proposal
     */
    function cancelCancellationProposal(uint256 tokenId) external;

    // /**
    //  * @notice Retrieves the current cancellation proposal status for a given token.
    //  *
    //  * @param tokenId The token id to check the proposal status for
    //  * @return refundAmount The proposed refund amount
    //  * @return initiatedBy The address that initiated the cancellation
    //  * @return status The status of the cancellation proposal
    //  * @return rejectionReason The reason for rejecting the cancellation
    //  */
    // function getCancellationProposalStatus(
    //     uint256 tokenId
    // )
    //     external
    //     view
    //     returns (
    //         uint256 refundAmount,
    //         address initiatedBy,
    //         CancellationProposalStatus status,
    //         CancellationRejectionReason rejectionReason
    //     );
}
