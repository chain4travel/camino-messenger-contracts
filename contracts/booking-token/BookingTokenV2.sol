// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import { BookingToken, Address, SafeERC20, IERC20 } from "./BookingToken.sol";
import { CancellationProposalStatus, CancellationRejectionReason } from "./IBookingToken.sol";

//import { CancellationProposalStatus, CancellationRejectionReason } from "./IBookingToken.sol";

/**
 * @title BookingTokenV2
 * @notice This contract extends BookingToken to add additional functionality.
 * Specifically, it introduces a cancellation process after the token is bought.
 */
contract BookingTokenV2 is BookingToken {
    using Address for address payable;
    using SafeERC20 for IERC20;

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    struct CancellationProposal {
        uint256 refundAmount;
        address proposedBy;
        CancellationProposalStatus status;
        CancellationRejectionReason rejectionReason;
    }
    /// @custom:storage-location erc7201:camino.messenger.storage.BookingTokenCancellable
    struct BookingTokenCancellableStorage {
        // Mapping to store the ongoing cancellation proposals for each token
        mapping(uint256 tokenId => CancellationProposal cancellationProposal) _cancellationProposals;
        // Mapping to store the cancellable flag for each token
        mapping(uint256 tokenId => bool cancellable) _isCancellable;
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
     * @notice Event emitted when a token is reserved.
     *
     * @param tokenId token id
     * @param reservedFor reserved for address
     * @param supplier supplier address
     * @param expirationTimestamp expiration timestamp
     * @param price price of the token
     * @param paymentToken payment token address
     */
    event TokenReserved(
        uint256 indexed tokenId,
        address indexed reservedFor,
        address indexed supplier,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken,
        bool isCancellable
    );

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
     * @notice Event emitted when a cancellation proposal is accepted by the owner.
     *
     * @param tokenId token id
     * @param acceptedBy address that accepted the proposal
     * @param refundAmount proposed refund amount
     */
    event CancellationProposalAcceptedByTheOwner(
        uint256 indexed tokenId,
        address indexed acceptedBy,
        uint256 refundAmount
    );

    /**
     * @notice Event emitted when a cancellation proposal is countered.
     *
     * @param tokenId token id
     * @param counteredBy address that countered the proposal
     * @param newRefundAmount new proposed refund amount
     */
    event CancellationCountered(uint256 indexed tokenId, address indexed counteredBy, uint256 newRefundAmount);

    /**
     * @notice Event emitted when a cancellation proposal is cancelled.
     *
     * @param tokenId token id
     * @param cancelledBy address that cancelled the proposal
     */
    event CancellationProposalCancelled(uint256 indexed tokenId, address indexed cancelledBy);

    /**
     * @notice Event emitted when the cancellable flag for a token is updated.
     *
     * @param tokenId token id
     * @param isCancellable new cancellable flag
     */
    event TokenCancellableUpdated(uint256 indexed tokenId, bool isCancellable);

    /**
     * @notice Event emitted when a cancellation proposal is rejected.
     *
     * @param tokenId token id
     * @param rejectedBy address that rejected the proposal
     * @param reason reason for rejection
     */
    event CancellationRejected(uint256 indexed tokenId, address indexed rejectedBy, CancellationRejectionReason reason);

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
     * @param tokenId The token id
     * @param caller The address of the caller
     */
    error NotAuthorizedToAcceptCancellation(uint256 tokenId, address caller);

    /**
     * @notice Error for when the caller is not authorized to reject a cancellation.
     *
     * @param tokenId The token id
     * @param caller The address of the caller
     */
    error NotAuthorizedToRejectCancellation(uint256 tokenId, address caller);

    /**
     * @notice Error for when the caller is not authorized to counter a cancellation.
     *
     * @param tokenId The token id
     * @param caller The address of the caller
     */
    error NotAuthorizedToCounterCancellation(uint256 tokenId, address caller);

    /**
     * @notice Error for when the caller is not authorized to accept a counter proposal.
     *
     * @param tokenId The token id
     * @param caller The address of the caller
     */
    error NotAuthorizedToAcceptCounterProposal(uint256 tokenId, address caller);

    /**
     * @notice Error for when the caller is not authorized to cancel a proposal.
     *
     * @param tokenId The token id
     * @param caller The address of the caller
     */
    error NotAuthorizedToCancelProposal(uint256 tokenId, address caller);

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
    error TokenHasActiveCancellationProposalOrCancelled(uint256 tokenId);

    /**
     * @notice Incorrect amount
     *
     * @param actual The actual amount
     * @param expected The expected amount
     */
    error IncorrectAmount(uint256 actual, uint256 expected);

    /**
     * @notice Error for when the caller is not authorized to set the cancellable flag.
     *
     * @param tokenId The token id
     * @param caller The address of the caller
     */
    error NotAuthorizedToSetCancellable(uint256 tokenId, address caller);

    /***************************************************
     *                   FUNCTIONS                     *
     ***************************************************/
    /**
     * @notice Mints a new token with a reservation for a specific address.
     *
     * @param reservedFor The CM Account address that can buy the token
     * @param uri The URI of the token
     * @param expirationTimestamp The expiration timestamp
     * @param price The price of the token
     * @param paymentToken The token used to pay for the reservation. If address(0) then native.
     * @param _isCancellable The cancellable flag
     */
    function safeMintWithReservation(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken,
        bool _isCancellable
    ) public virtual onlyCMAccount(msg.sender) {
        // Require reservedFor to be a CM Account
        requireCMAccount(reservedFor);

        BookingTokenStorage storage $ = _getBookingTokenStorage();

        // Expiration timestamp should be at least `_minExpirationTimestampDiff`
        // seconds in the future
        uint256 minExpirationTimestampDiff = $._minExpirationTimestampDiff;
        if (!(expirationTimestamp > (block.timestamp + minExpirationTimestampDiff))) {
            revert ExpirationTimestampTooSoon(expirationTimestamp, minExpirationTimestampDiff);
        }

        // Increment the token id
        uint256 tokenId = $._nextTokenId++;

        // Mint the token for the supplier (the caller)
        _safeMint(msg.sender, tokenId);
        _setTokenURI(tokenId, uri);

        // Store the reservation
        _reserve(tokenId, reservedFor, msg.sender, expirationTimestamp, price, paymentToken);

        // Set the status
        $._bookingStatus[tokenId] = BookingStatus.Reserved;

        // Set the cancellable flag
        _getBookingTokenCancellableStorage()._isCancellable[tokenId] = _isCancellable;

        emit TokenReserved(tokenId, reservedFor, msg.sender, expirationTimestamp, price, paymentToken, _isCancellable);
    }

    /**
     * @notice Mints a new token with a reservation for a specific address. Setting
     * the cancellable flag to false. Original function signature from V1.
     *
     * @param reservedFor The CM Account address that can buy the token
     * @param uri The URI of the token
     * @param expirationTimestamp The expiration timestamp
     * @param price The price of the token
     * @param paymentToken The token used to pay for the reservation. If address(0)
       then native.
     */
    function safeMintWithReservation(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken
    ) public {
        safeMintWithReservation(reservedFor, uri, expirationTimestamp, price, paymentToken, false);
    }

    /**
     * @notice Returns true if the token is cancellable and false otherwise.
     * @param tokenId The token id
     */
    function isCancellable(uint256 tokenId) external view returns (bool) {
        return _getBookingTokenCancellableStorage()._isCancellable[tokenId];
    }

    /**
     * @notice Retrieves the refund amount for a given token.
     *
     * @param tokenId The token id to retrieve the refund amount for
     * @return refundAmount The refund amount in wei
     */
    function getCancellationProposalRefundAmount(uint256 tokenId) external view returns (uint256 refundAmount) {
        return _getBookingTokenCancellableStorage()._cancellationProposals[tokenId].refundAmount;
    }

    /**
     * @notice Retrieves the payment token for a given token.
     *
     * @param tokenId The token id to retrieve the payment token for
     * @return paymentToken The payment token
     */
    function getReservationPaymentToken(uint256 tokenId) external view returns (IERC20 paymentToken) {
        return _getBookingTokenStorage()._reservations[tokenId].paymentToken;
    }

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

        // Revert if there is already an active cancellation proposal
        if (cancellableStorage._cancellationProposals[tokenId].status != CancellationProposalStatus.NoProposal) {
            revert TokenHasActiveCancellationProposalOrCancelled(tokenId);
        }

        // Store cancellation proposal
        cancellableStorage._cancellationProposals[tokenId] = CancellationProposal({
            refundAmount: refundAmount,
            proposedBy: msg.sender,
            status: CancellationProposalStatus.Pending,
            rejectionReason: CancellationRejectionReason.Unspecified
        });

        emit CancellationPending(tokenId, msg.sender, refundAmount);
    }

    /**
     * @notice Reject a cancellation proposal for a bought token.
     *
     * @param tokenId The token id to reject the cancellation for
     * @param reason The reason for rejecting the cancellation
     */
    function rejectCancellationProposal(uint256 tokenId, CancellationRejectionReason reason) external {
        BookingTokenCancellableStorage storage cancellableStorage = _getBookingTokenCancellableStorage();
        CancellationProposal memory proposal = cancellableStorage._cancellationProposals[tokenId];

        // Revert if the cancellation proposal status is not "Pending"
        if (proposal.status != CancellationProposalStatus.Pending) {
            revert NoPendingCancellationProposal(tokenId);
        }

        BookingTokenStorage storage $ = _getBookingTokenStorage();
        TokenReservation memory reservation = $._reservations[tokenId];

        address owner = _requireOwned(tokenId);

        // Revert if the caller is not the supplier and the proposer is not the
        // owner. Only the supplier can reject a cancellation proposal.
        if (msg.sender != reservation.supplier || proposal.proposedBy != owner) {
            revert NotAuthorizedToRejectCancellation(tokenId, msg.sender);
        }

        // Reject the cancellation proposal
        cancellableStorage._cancellationProposals[tokenId].status = CancellationProposalStatus.Rejected;
        cancellableStorage._cancellationProposals[tokenId].rejectionReason = reason;

        emit CancellationRejected(tokenId, msg.sender, reason);
    }

    /**
     * @notice Accepts a cancellation proposal for a bought token and finalizes it.
     *
     * @param tokenId The token id to accept the cancellation for
     */
    function acceptCancellationProposal(uint256 tokenId, uint256 checkRefundAmount) external payable {
        BookingTokenCancellableStorage storage cancellableStorage = _getBookingTokenCancellableStorage();
        CancellationProposal memory proposal = cancellableStorage._cancellationProposals[tokenId];

        // Revert if the checkRefundAmount is not equal to the proposal refund amount
        if (checkRefundAmount != proposal.refundAmount) {
            revert IncorrectAmount(checkRefundAmount, proposal.refundAmount);
        }

        // Revert if the cancellation proposal status is not "Pending"
        if (proposal.status != CancellationProposalStatus.Pending) {
            revert NoPendingCancellationProposal(tokenId);
        }

        BookingTokenStorage storage $ = _getBookingTokenStorage();
        TokenReservation memory reservation = $._reservations[tokenId];

        address owner = _requireOwned(tokenId);

        // Set the proposer to the owner if the proposedBy is the supplier and the
        // caller is the owner. This means that the distributor/owner is accepting
        // the proposal. Setting the proposedBy to the owner, which allows the
        // supplier to accept the proposal.
        if (proposal.proposedBy == reservation.supplier && msg.sender == owner) {
            // Set proposedBy to distributor/owner
            cancellableStorage._cancellationProposals[tokenId].proposedBy = owner;

            // Emit event so the supplier can get notified
            emit CancellationProposalAcceptedByTheOwner(tokenId, owner, proposal.refundAmount);

            // We're done here because there is nothing else to do for the owner.
            return;
        }

        // Revert if the caller is not the supplier and the proposer is not the
        // owner. Only the supplier can accept a cancellation proposal after this
        // point.
        if (msg.sender != reservation.supplier || proposal.proposedBy != owner) {
            revert NotAuthorizedToAcceptCancellation(tokenId, msg.sender);
        }

        // Finalize the cancellation

        // Do the payment of the refund
        if (address(reservation.paymentToken) != address(0) && proposal.refundAmount > 0) {
            // Payment is in ERC20.
            //
            // Message sender (supplier of the Booking Token) must provide enough
            // allowance for this (BookingToken) contract to pay the cancellation
            // refund amount to the owner.
            uint256 allowance = reservation.paymentToken.allowance(msg.sender, address(this));
            if (allowance < proposal.refundAmount) {
                revert InsufficientAllowance(msg.sender, reservation.paymentToken, proposal.refundAmount, allowance);
            }

            // Transfer proposal refund amount in the ERC20 tokens from the supplier
            // to the owner
            reservation.paymentToken.safeTransferFrom(msg.sender, owner, proposal.refundAmount);
        } else {
            // Payment is in native currency or refund is zero.
            // Check if we receive the right refund amount
            if (msg.value != proposal.refundAmount) {
                revert IncorrectAmount(msg.value, proposal.refundAmount);
            }

            // Transfer payment to the owner
            payable(owner).sendValue(msg.value);
        }

        // Set token status to "Cancelled"
        $._bookingStatus[tokenId] = BookingStatus.Cancelled;

        // Set cancellation proposal status to "accepted"
        cancellableStorage._cancellationProposals[tokenId].status = CancellationProposalStatus.Accepted;

        // Burn the token
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

        if (
            proposal.status != CancellationProposalStatus.Pending &&
            proposal.status != CancellationProposalStatus.Rejected
        ) {
            revert NoPendingCancellationProposal(tokenId);
        }

        BookingTokenStorage storage $ = _getBookingTokenStorage();
        address supplier = $._reservations[tokenId].supplier;

        address owner = _requireOwned(tokenId);

        // Revert if the caller is not the supplier
        if (msg.sender != supplier || proposal.proposedBy != owner) {
            revert NotAuthorizedToCounterCancellation(tokenId, msg.sender);
        }

        // Update the proposal with the new values
        cancellableStorage._cancellationProposals[tokenId].refundAmount = newRefundAmount;

        // Set cancellation proposal status to "countered"
        cancellableStorage._cancellationProposals[tokenId].status = CancellationProposalStatus.Countered;
        cancellableStorage._cancellationProposals[tokenId].rejectionReason = CancellationRejectionReason.Unspecified;

        // Emit the countered proposal event
        emit CancellationCountered(tokenId, msg.sender, newRefundAmount);
    }

    /**
     * @notice Accept a countered cancellation proposal
     * @param tokenId The token id to accept the countered cancellation proposal for
     */
    function acceptCounteredCancellationProposal(uint256 tokenId, uint256 checkRefundAmount) external {
        address owner = _requireOwned(tokenId);

        // Revert if the caller is not the owner
        if (msg.sender != owner) {
            revert NotAuthorizedToAcceptCounterProposal(tokenId, msg.sender);
        }

        BookingTokenCancellableStorage storage cancellableStorage = _getBookingTokenCancellableStorage();
        CancellationProposal storage proposal = cancellableStorage._cancellationProposals[tokenId];

        // Revert if the checkRefundAmount is not equal to the proposal refund amount
        if (checkRefundAmount != proposal.refundAmount) {
            revert IncorrectAmount(checkRefundAmount, proposal.refundAmount);
        }

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
            revert NotAuthorizedToCancelProposal(tokenId, msg.sender);
        }

        // Revert if the cancellation proposal status is not "Pending" or "Countered"
        if (
            proposal.status != CancellationProposalStatus.Pending &&
            proposal.status != CancellationProposalStatus.Countered
        ) {
            revert NoPendingCancellationProposal(tokenId);
        }

        // Cancel the proposal by deleting it from the storage
        delete cancellableStorage._cancellationProposals[tokenId];

        // Emit the cancellation proposal cancelled event
        emit CancellationProposalCancelled(tokenId, msg.sender);
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
    )
        external
        view
        returns (
            uint256 refundAmount,
            address proposedBy,
            CancellationProposalStatus status,
            CancellationRejectionReason rejectionReason
        )
    {
        BookingTokenCancellableStorage storage cancellableStorage = _getBookingTokenCancellableStorage();
        CancellationProposal memory proposal = cancellableStorage._cancellationProposals[tokenId];
        return (proposal.refundAmount, proposal.proposedBy, proposal.status, proposal.rejectionReason);
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
            revert TokenHasActiveCancellationProposalOrCancelled(tokenId);
        }
    }
}
