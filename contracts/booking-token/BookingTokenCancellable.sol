// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.24;

enum CancellationProposalStatus {
    NO_PROPOSAL, // 0, default
    PENDING, // 1
    REJECTED, // 2
    WITHDRAWN, // 3
    FINALIZED // 4
}

contract BookingTokenCancellable {
    struct Proposal {
        uint256 refundAmount; // Slot n (32 bytes)
        address initialProposer; // Slot n+1 (20 bytes)
        uint32 timesCountered; // Packed above (4 bytes, 24 bytes total)
        bool ownerAccepted; // Packed above (1 byte, 25 bytes total)
        bool supplierAccepted; // Packed above (1 byte, 26 bytes total, 6 bytes remaining)
        address currentProposer; // Slot n+2 (20 bytes)
        uint32 timesRejected; // Packed above (4 bytes, 24 bytes total)
        CancellationProposalStatus status; // Packed above (1 byte, 25 bytes total)
        uint16 cancellationReason; // Packed above (2 bytes, 27 bytes total)
        uint16 cancellationVersion; // Packed above (2 bytes, 29 bytes total)
        uint16 rejectionReason; // Packed above (2 bytes, 31 bytes total, 1 byte remaining)
        uint16 rejectionVersion; // Slot n+3 (2 bytes)
        uint16 counterReason; // Packed above (2 bytes, 4 bytes total)
        uint16 counterVersion; // Packed above (2 bytes, 6 bytes total)
        uint16 withdrawalReason; // Packed above (2 bytes, 8 bytes total)
        uint16 withdrawalVersion; // Packed above (2 bytes, 10 bytes total, 22 bytes remaining)
    }

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    /// @custom:storage-location erc7201:camino.messenger.storage.BookingTokenCancellable
    struct BookingTokenCancellableStorage {
        // Mapping to store the ongoing cancellation proposals for each token
        mapping(uint256 tokenId => Proposal proposal) _proposals;
    }

    // keccak256(abi.encode(uint256(keccak256("camino.messenger.storage.BookingTokenCancellableV2")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant BookingTokenCancellableStorageLocation =
        0x0b8fb32ffc7043fda9e0ee2bcb4236acf95ab448752c73eff6cc7f2640ff8500;

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

    event CancellationPending(
        uint256 indexed tokenId,
        address indexed initialProposer,
        address indexed currentProposer,
        uint256 refundAmount,
        bool ownerAccepted,
        bool supplierAccepted,
        uint32 timesCountered,
        uint32 timesRejected
    );

    event CancellationReasons(
        uint256 indexed tokenId,
        uint16 cancellationReason,
        uint16 cancellationReasonVersion,
        uint16 rejectionReason,
        uint16 rejectionVersion,
        uint16 counterReason,
        uint16 counterVersion,
        uint16 withdrawalReason,
        uint16 withdrawalVersion
    );

    event CancellationWithdrawn(uint256 indexed tokenId, uint16 withdrawalReason, uint16 withdrawalVersion);

    event CancellationRejected(uint256 indexed tokenId, uint16 rejectionReason, uint16 rejectionVersion);

    event CancellationFinalized(uint256 indexed tokenId);

    /***************************************************
     *                   ERRORS                        *
     ***************************************************/

    error NotOwnerOrSupplier();

    error CancellationProposalExists(uint256 tokenId);

    error IncorrectRefundAmount(uint256 tokenId, uint256 existing, uint256 checked);

    error InvalidCancellationProposalStatus(uint256 tokenId, CancellationProposalStatus status);

    error OnlySupplierCanFinalizeCancellation(uint256 tokenId);

    error OwnerNotAcceptedCancellation(uint256 tokenId);

    error ProposerCanNotRejectCancellation(uint256 tokenId);

    error OnlyCurrentProposerCanWithdrawCancellation(uint256 tokenId);

    /***************************************************
     *             CANCELLATION LOGIC                  *
     ***************************************************/

    function requireOwnerOrSupplier(address owner, address supplier) internal view {
        if (msg.sender != owner && msg.sender != supplier) {
            revert NotOwnerOrSupplier();
        }
    }

    modifier onlyOwnerOrSupplier(address owner, address supplier) {
        requireOwnerOrSupplier(owner, supplier);
        _;
    }

    function _getCancellationProposalStatusAndCurrentProposer(
        uint256 tokenId
    ) internal view returns (CancellationProposalStatus status, address currentProposer) {
        return (
            _getBookingTokenCancellableStorage()._proposals[tokenId].status,
            _getBookingTokenCancellableStorage()._proposals[tokenId].currentProposer
        );
    }

    function getCancellationProposal(
        uint256 tokenId
    )
        external
        view
        returns (
            CancellationProposalStatus,
            uint256 refundAmount,
            address initialProposer,
            address currentProposer,
            bool ownerAccepted,
            bool supplierAccepted,
            uint32 timesCountered,
            uint32 timesRejected
        )
    {
        Proposal storage proposal = _getBookingTokenCancellableStorage()._proposals[tokenId];

        return (
            proposal.status,
            proposal.refundAmount,
            proposal.initialProposer,
            proposal.currentProposer,
            proposal.ownerAccepted,
            proposal.supplierAccepted,
            proposal.timesCountered,
            proposal.timesRejected
        );
    }

    function getCancellationReasons(
        uint256 tokenId
    )
        external
        view
        returns (
            uint16 cancellationReason,
            uint16 cancellationVersion,
            uint16 rejectionReason,
            uint16 rejectionVersion,
            uint16 counterReason,
            uint16 counterVersion,
            uint16 withdrawalReason,
            uint16 withdrawalVersion
        )
    {
        Proposal storage proposal = _getBookingTokenCancellableStorage()._proposals[tokenId];

        return (
            proposal.cancellationReason,
            proposal.cancellationVersion,
            proposal.rejectionReason,
            proposal.rejectionVersion,
            proposal.counterReason,
            proposal.counterVersion,
            proposal.withdrawalReason,
            proposal.withdrawalVersion
        );
    }

    function _initiateCancellation(
        address owner,
        address supplier,
        uint256 tokenId,
        uint256 refundAmount,
        uint16 cancellationReason,
        uint16 cancellationReasonVersion
    ) internal virtual onlyOwnerOrSupplier(owner, supplier) {
        Proposal storage proposal = _getBookingTokenCancellableStorage()._proposals[tokenId];

        // Revert if proposal is finalized or pending
        if (
            proposal.status == CancellationProposalStatus.FINALIZED ||
            proposal.status == CancellationProposalStatus.PENDING
        ) {
            revert InvalidCancellationProposalStatus(tokenId, proposal.status);
        }

        // Set initial owner if this is the first proposal
        if (proposal.status == CancellationProposalStatus.NO_PROPOSAL) {
            proposal.initialProposer = msg.sender;
        }

        // REST IS WITHDRAW/REJECTED LOGIC

        // Set the current proposer
        proposal.currentProposer = msg.sender;

        // Set refund amount
        proposal.refundAmount = refundAmount;

        // Set accepted flags
        proposal.ownerAccepted = (msg.sender == owner);
        proposal.supplierAccepted = (msg.sender == supplier);

        // Set new cancellation reason
        proposal.cancellationReason = cancellationReason;
        proposal.cancellationVersion = cancellationReasonVersion;

        // Reset other reasons
        proposal.rejectionReason = 0;
        proposal.rejectionVersion = 0;
        proposal.counterReason = 0;
        proposal.counterVersion = 0;
        proposal.withdrawalReason = 0;
        proposal.withdrawalVersion = 0;

        // Set status to PENDING
        proposal.status = CancellationProposalStatus.PENDING;

        // Emit event
        emit CancellationPending(
            tokenId,
            proposal.initialProposer,
            proposal.currentProposer,
            proposal.refundAmount,
            proposal.ownerAccepted,
            proposal.supplierAccepted,
            proposal.timesCountered,
            proposal.timesRejected
        );

        // Emit reasons event
        emit CancellationReasons(
            tokenId,
            proposal.cancellationReason,
            proposal.cancellationVersion,
            proposal.rejectionReason,
            proposal.rejectionVersion,
            proposal.counterReason,
            proposal.counterVersion,
            proposal.withdrawalReason,
            proposal.withdrawalVersion
        );
    }

    /**
     * @notice Used by the owner or supplier to accept a cancellation proposal that
     * is initiated or countered by the other party
     *
     * @param owner Owner of the token
     * @param supplier Supplier of the token
     * @param tokenId Token ID
     * @param checkRefundAmount Refund amount to check against, to prevent front-running
     */
    function _acceptCancellation(
        address owner,
        address supplier,
        uint256 tokenId,
        uint256 checkRefundAmount
    ) internal virtual onlyOwnerOrSupplier(owner, supplier) {
        Proposal storage proposal = _getBookingTokenCancellableStorage()._proposals[tokenId];

        // Revert if not in PENDING state
        if (proposal.status != CancellationProposalStatus.PENDING) {
            revert InvalidCancellationProposalStatus(tokenId, proposal.status);
        }

        // Revert if refund amount does not match
        if (proposal.refundAmount != checkRefundAmount) {
            revert IncorrectRefundAmount(tokenId, proposal.refundAmount, checkRefundAmount);
        }

        // Accept the cancellation
        if (msg.sender == owner) {
            proposal.ownerAccepted = true;
        } else {
            proposal.supplierAccepted = true;
        }

        // Emit event
        emit CancellationPending(
            tokenId,
            proposal.initialProposer,
            proposal.currentProposer,
            proposal.refundAmount,
            proposal.ownerAccepted,
            proposal.supplierAccepted,
            proposal.timesCountered,
            proposal.timesRejected
        );

        // Emit reasons event
        emit CancellationReasons(
            tokenId,
            proposal.cancellationReason,
            proposal.cancellationVersion,
            proposal.rejectionReason,
            proposal.rejectionVersion,
            proposal.counterReason,
            proposal.counterVersion,
            proposal.withdrawalReason,
            proposal.withdrawalVersion
        );
    }

    function _counterCancellation(
        address owner,
        address supplier,
        uint256 tokenId,
        uint256 refundAmount,
        uint16 counterReason,
        uint16 counterVersion
    ) internal virtual onlyOwnerOrSupplier(owner, supplier) {
        Proposal storage proposal = _getBookingTokenCancellableStorage()._proposals[tokenId];

        // Revert if not in PENDING state
        if (proposal.status != CancellationProposalStatus.PENDING) {
            revert InvalidCancellationProposalStatus(tokenId, proposal.status);
        }

        // Set new refund amount
        proposal.refundAmount = refundAmount;

        // Set the current proposer to the counter proposer
        proposal.currentProposer = msg.sender;

        // Set accepted flags
        proposal.ownerAccepted = (msg.sender == owner);
        proposal.supplierAccepted = (msg.sender == supplier);

        // Set the counter reason
        proposal.counterReason = counterReason;
        proposal.counterVersion = counterVersion;

        // Increment times countered
        proposal.timesCountered++;

        // Emit event
        emit CancellationPending(
            tokenId,
            proposal.initialProposer,
            proposal.currentProposer,
            proposal.refundAmount,
            proposal.ownerAccepted,
            proposal.supplierAccepted,
            proposal.timesCountered,
            proposal.timesRejected
        );

        // Emit reasons event
        emit CancellationReasons(
            tokenId,
            proposal.cancellationReason,
            proposal.cancellationVersion,
            proposal.rejectionReason,
            proposal.rejectionVersion,
            proposal.counterReason,
            proposal.counterVersion,
            proposal.withdrawalReason,
            proposal.withdrawalVersion
        );
    }

    function _withdrawCancellation(
        address owner,
        address supplier,
        uint256 tokenId,
        uint16 withdrawalReason,
        uint16 withdrawalVersion
    ) internal virtual onlyOwnerOrSupplier(owner, supplier) {
        Proposal storage proposal = _getBookingTokenCancellableStorage()._proposals[tokenId];

        // Revert if not in PENDING state
        if (proposal.status != CancellationProposalStatus.PENDING) {
            revert InvalidCancellationProposalStatus(tokenId, proposal.status);
        }

        // Only current proposer can withdraw
        if (msg.sender != proposal.currentProposer) {
            revert OnlyCurrentProposerCanWithdrawCancellation(tokenId);
        }

        // Set withdrawal reason
        proposal.withdrawalReason = withdrawalReason;
        proposal.withdrawalVersion = withdrawalVersion;

        // Set status to WITHDRAWN
        proposal.status = CancellationProposalStatus.WITHDRAWN;

        // Emit event
        emit CancellationWithdrawn(tokenId, withdrawalReason, withdrawalVersion);
    }

    function _rejectCancellation(
        address owner,
        address supplier,
        uint256 tokenId,
        uint16 rejectionReason,
        uint16 rejectionReasonVersion
    ) internal virtual onlyOwnerOrSupplier(owner, supplier) {
        Proposal storage proposal = _getBookingTokenCancellableStorage()._proposals[tokenId];

        // Revert if not in PENDING state
        if (proposal.status != CancellationProposalStatus.PENDING) {
            revert InvalidCancellationProposalStatus(tokenId, proposal.status);
        }

        // Proposer can not reject the cancellation
        if (msg.sender == proposal.currentProposer) {
            revert ProposerCanNotRejectCancellation(tokenId);
        }

        // Set reason
        proposal.rejectionReason = rejectionReason;
        proposal.rejectionVersion = rejectionReasonVersion;

        // Set status to REJECTED
        proposal.status = CancellationProposalStatus.REJECTED;

        // Increment times rejected
        proposal.timesRejected++;

        // Emit event
        emit CancellationRejected(tokenId, rejectionReason, rejectionReasonVersion);
    }

    function _finalizeCancellation(
        address supplier,
        uint256 tokenId,
        uint256 checkRefundAmount
    ) internal virtual returns (uint256 refundAmount) {
        // Only supplier can finalize the cancellation
        if (msg.sender != supplier) {
            revert OnlySupplierCanFinalizeCancellation(tokenId);
        }

        Proposal storage proposal = _getBookingTokenCancellableStorage()._proposals[tokenId];

        // Revert if not in PENDING state
        if (proposal.status != CancellationProposalStatus.PENDING) {
            revert InvalidCancellationProposalStatus(tokenId, proposal.status);
        }

        // Revert if refund amount does not match
        if (proposal.refundAmount != checkRefundAmount) {
            revert IncorrectRefundAmount(tokenId, proposal.refundAmount, checkRefundAmount);
        }

        // Revert if owner has not accepted the cancellation
        if (!proposal.ownerAccepted) {
            revert OwnerNotAcceptedCancellation(tokenId);
        }

        // Set supplier accepted
        if (!proposal.supplierAccepted) {
            proposal.supplierAccepted = true;
        }

        // Set status to FINALIZED
        proposal.status = CancellationProposalStatus.FINALIZED;

        // Emit event. Payment should be handled by the inheriting contract.
        emit CancellationFinalized(tokenId);

        return proposal.refundAmount;
    }
}
