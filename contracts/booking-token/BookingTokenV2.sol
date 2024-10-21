// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import { BookingToken, IERC20 } from "./BookingToken.sol";

/**
 * @title BookingTokenV2
 * @notice This contract extends BookingToken to add additional functionality.
 * Specifically, it introduces a cancellation process after the token is bought.
 */
contract BookingTokenV2 is BookingToken {
    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    enum CancellationProposalStatus {
        NoProposal, // 0, default
        Pending, // 1
        Rejected, // 2
        Countered, // 3
        Accepted // 4
    }

    struct CancellationProposal {
        uint256 refundAmount;
        address proposedBy;
        CancellationProposalStatus status;
    }
    /// @custom:storage-location erc7201:camino.messenger.storage.BookingTokenCancellable
    struct BookingTokenCancellableStorage {
        // Mapping to store the ongoing cancellation proposals for each token
        mapping(uint256 => CancellationProposal) _cancellationProposals;
    }

    // keccak256(abi.encode(uint256(keccak256("camino.messenger.storage.BookingTokenCancellable")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant BookingTokenCancellableStorageLocation =
        0x56ee42015b616c256a09657decaf7aa5d877decbe489a4b4b22f8bb476600500;

    /**
     * @notice Retrieves the BookingTokenCancellable storage struct from the designated storage slot.
     *
     * @return $ The storage struct reference
     */
    function _getBookingTokenCancellableStorage() private pure returns (BookingTokenCancellableStorage storage $) {
        assembly {
            $.slot := BookingTokenCancellableStorageLocation
        }
    }

    /***************************************************
     *                   EVENTS                        *
     ***************************************************/

    /**
     * @notice Event emitted when a cancellation is initiated.
     *
     * @param tokenId token id
     * @param proposedBy address that initiated the cancellation
     * @param refundAmount proposed refund amount
     */
    event CancellationPending(uint256 indexed tokenId, address indexed proposedBy, uint256 refundAmount);

    /**
     * @notice Event emitted when a cancellation is accepted.
     *
     * @param tokenId token id
     * @param acceptedBy address that accepted the cancellation
     * @param refundAmount final refund amount
     */
    event CancellationAccepted(uint256 indexed tokenId, address indexed acceptedBy, uint256 refundAmount);

    /**
     * @notice Event emitted when a cancellation proposal is countered.
     *
     * @param tokenId token id
     * @param counteredBy address that countered the proposal
     * @param newRefundAmount new proposed refund amount
     */
    event CancellationCountered(uint256 indexed tokenId, address indexed counteredBy, uint256 newRefundAmount);

    /**
     * @notice Event emitted when a cancellation proposal is canceled.
     *
     * @param tokenId token id
     * @param canceledBy address that canceled the proposal
     */
    event CancellationProposalCanceled(uint256 indexed tokenId, address indexed canceledBy);

    /***************************************************
     *                   ERRORS                        *
     ***************************************************/

    /**
     * @notice Error for when the caller is not authorized to initiate a cancellation.
     *
     * @param caller The address of the caller
     */
    error NotAuthorizedToInitiateCancellation(address caller);

    /**
     * @notice Error for when there is no pending cancellation proposal.
     *
     * @param tokenId The token id for which there is no pending proposal
     */
    error NoPendingCancellationProposal(uint256 tokenId);

    /**
     * @notice Error for when the caller is not authorized to accept a cancellation.
     *
     * @param caller The address of the caller
     */
    error NotAuthorizedToAcceptCancellation(address caller);

    /**
     * @notice Error for when the caller is not authorized to counter a cancellation.
     *
     * @param caller The address of the caller
     */
    error NotAuthorizedToCounterCancellation(address caller);

    /**
     * @notice Error for when the caller is not authorized to accept a counter proposal.
     *
     * @param caller The address of the caller
     */
    error NotAuthorizedToAcceptCounterProposal(address caller);

    /**
     * @notice Error for when the caller is not authorized to cancel a proposal.
     *
     * @param caller The address of the caller
     */
    error NotAuthorizedToCancelProposal(address caller);

    /**
     * @notice Error for when there is no countered cancellation proposal.
     *
     * @param tokenId The token id for which there is no countered proposal
     */
    error NoCounteredCancellationProposal(uint256 tokenId);

    /**
     * @notice Error for when a token has an active cancellation proposal and cannot be transferred.
     *
     * @param tokenId The token id that has an active cancellation proposal
     */
    error TokenHasActiveCancellationProposal(uint256 tokenId);

    /***************************************************
     *                   FUNCTIONS                     *
     ***************************************************/

    /**
     * @notice Initiates a cancellation for a bought token.
     *
     * @param tokenId The token id to initiate the cancellation for
     * @param refundAmount The proposed refund amount in wei
     */
    function initiateCancellationProposal(uint256 tokenId, uint256 refundAmount) external {
        // Revert if the token status is not "bought"
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        if ($._bookingStatus[tokenId] != BookingStatus.Bought) {
            revert InvalidTokenStatus(tokenId, $._bookingStatus[tokenId]);
        }

        // Get owner and supplier
        address owner = _requireOwned(tokenId);
        address supplier = $._reservations[tokenId].supplier;

        // Revert if the caller is not the owner or supplier
        if (msg.sender != owner && msg.sender != supplier) {
            revert NotAuthorizedToInitiateCancellation(msg.sender);
        }

        BookingTokenCancellableStorage storage cancellableStorage = _getBookingTokenCancellableStorage();

        // FIXME: Should we allow cancellation proposals to be initiated when there
        // is an active cancellation proposal and the status is Rejected?

        // Store cancellation proposal
        cancellableStorage._cancellationProposals[tokenId] = CancellationProposal({
            refundAmount: refundAmount,
            proposedBy: msg.sender,
            status: CancellationProposalStatus.Pending
        });

        emit CancellationPending(tokenId, msg.sender, refundAmount);
    }

    /**
     * @notice Accepts a cancellation proposal for a bought token and finalizes it.
     *
     * @param tokenId The token id to accept the cancellation for
     */
    function acceptCancellationProposal(uint256 tokenId) external {
        BookingTokenCancellableStorage storage cancellableStorage = _getBookingTokenCancellableStorage();
        CancellationProposal memory proposal = cancellableStorage._cancellationProposals[tokenId];

        // Revert if the cancellation proposal status is not "Pending"
        if (proposal.status != CancellationProposalStatus.Pending) {
            revert NoPendingCancellationProposal(tokenId);
        }

        BookingTokenStorage storage $ = _getBookingTokenStorage();
        address supplier = $._reservations[tokenId].supplier;

        // Revert if the caller is not the supplier, only the supplier can accept a
        // cancellation proposal
        if (msg.sender != supplier) {
            revert NotAuthorizedToAcceptCancellation(msg.sender);
        }

        // Set token status to "cancelled"
        $._bookingStatus[tokenId] = BookingStatus.Canceled;

        // Set cancellation proposal status to "accepted"
        cancellableStorage._cancellationProposals[tokenId].status = CancellationProposalStatus.Accepted;

        // Finalize the cancellation
        // TODO: Transfer refund to owner
        _burn(tokenId);

        // Emit the cancellation accepted event
        emit CancellationAccepted(tokenId, msg.sender, proposal.refundAmount);
    }

    /**
     * @notice Counters a cancellation proposal with a new proposal.
     *
     * @param tokenId The token id to counter the cancellation for
     * @param newRefundAmount The new proposed refund amount in wei
     */
    function counterCancellationProposal(uint256 tokenId, uint256 newRefundAmount) external {
        BookingTokenCancellableStorage storage cancellableStorage = _getBookingTokenCancellableStorage();
        CancellationProposal storage proposal = cancellableStorage._cancellationProposals[tokenId];

        if (proposal.status != CancellationProposalStatus.Pending) {
            revert NoPendingCancellationProposal(tokenId);
        }

        BookingTokenStorage storage $ = _getBookingTokenStorage();
        address supplier = $._reservations[tokenId].supplier;

        // Revert if the caller is not the supplier
        if (msg.sender != supplier) {
            revert NotAuthorizedToCounterCancellation(msg.sender);
        }

        // Update the proposal with the new values
        cancellableStorage._cancellationProposals[tokenId].refundAmount = newRefundAmount;

        // Set cancellation proposal status to "countered"
        cancellableStorage._cancellationProposals[tokenId].status = CancellationProposalStatus.Countered;

        // Emit the countered proposal event
        emit CancellationCountered(tokenId, msg.sender, newRefundAmount);
    }

    /**
     * @notice Accept a countered cancellation proposal
     * @param tokenId The token id to accept the countered cancellation proposal for
     */
    function acceptCounteredCancellationProposal(uint256 tokenId) external {
        address owner = _requireOwned(tokenId);

        // Revert if the caller is not the owner
        if (msg.sender != owner) {
            revert NotAuthorizedToAcceptCounterProposal(msg.sender);
        }

        BookingTokenCancellableStorage storage cancellableStorage = _getBookingTokenCancellableStorage();
        CancellationProposal storage proposal = cancellableStorage._cancellationProposals[tokenId];

        // Revert if the cancellation proposal status is not "Countered"
        if (proposal.status != CancellationProposalStatus.Countered) {
            revert NoCounteredCancellationProposal(tokenId);
        }

        // Set status to "Pending"
        cancellableStorage._cancellationProposals[tokenId].status = CancellationProposalStatus.Pending;

        // Emit the cancellation pending event to notify the supplier
        emit CancellationPending(tokenId, proposal.proposedBy, proposal.refundAmount);
    }

    /**
     * @notice Cancels a pending cancellation proposal. Only the proposer can cancel
     * the proposal.
     *
     * @param tokenId The token id for which to cancel the proposal
     */
    function cancelCancellationProposal(uint256 tokenId) external {
        BookingTokenCancellableStorage storage cancellableStorage = _getBookingTokenCancellableStorage();
        CancellationProposal storage proposal = cancellableStorage._cancellationProposals[tokenId];

        // Revert if the caller is not the proposer
        if (msg.sender != proposal.proposedBy) {
            revert NotAuthorizedToCancelProposal(msg.sender);
        }

        // Revert if the cancellation proposal status is not "Pending"
        if (proposal.status != CancellationProposalStatus.Pending) {
            revert NoPendingCancellationProposal(tokenId);
        }

        // Cancel the proposal by deleting it from the storage
        delete cancellableStorage._cancellationProposals[tokenId];

        // Emit the cancellation proposal canceled event
        emit CancellationProposalCanceled(tokenId, msg.sender);
    }

    /**
     * @notice Retrieves the current cancellation proposal status for a given token.
     *
     * @param tokenId The token id to check the proposal status for
     * @return refundAmount The proposed refund amount
     * @return proposedBy The address that initiated the cancellation
     * @return status The status of the cancellation proposal
     */
    function getCancellationProposalStatus(
        uint256 tokenId
    ) external view returns (uint256 refundAmount, address proposedBy, CancellationProposalStatus status) {
        BookingTokenCancellableStorage storage cancellableStorage = _getBookingTokenCancellableStorage();
        CancellationProposal memory proposal = cancellableStorage._cancellationProposals[tokenId];
        return (proposal.refundAmount, proposal.proposedBy, proposal.status);
    }

    /***************************************************
     *             OVERRIDE FUNCTIONS                  *
     ***************************************************/

    /**
     * @notice Checks if the token is transferable, considering reservation and cancellation status.
     *
     * @param tokenId The token id to check
     */
    function checkTransferable(uint256 tokenId) internal override {
        super.checkTransferable(tokenId);
        BookingTokenCancellableStorage storage cancellableStorage = _getBookingTokenCancellableStorage();
        CancellationProposal storage proposal = cancellableStorage._cancellationProposals[tokenId];

        // Allow transfer if cancellation status is "Rejected". Only in this state
        // the token is still usable.
        if (proposal.status == CancellationProposalStatus.Rejected) {
            return;
        }

        // FIXME: Should we prevent transfer if there is an active cancellation proposal?
        if (proposal.status != CancellationProposalStatus.NoProposal) {
            revert TokenHasActiveCancellationProposal(tokenId);
        }
    }
}
