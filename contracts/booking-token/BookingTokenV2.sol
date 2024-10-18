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

    struct CancellationProposal {
        uint256 refundAmount;
        address refundCurrency;
        address initiatedBy;
        bool isActive;
    }
    /// @custom:storage-location erc7201:camino.messenger.storage.BookingTokenV2
    struct BookingTokenV2Storage {
        // Mapping to store the original supplier (minter) of each token
        mapping(uint256 => address) _originalSupplier;
        // Mapping to store the ongoing cancellation proposals for each token
        mapping(uint256 => CancellationProposal) _cancellationProposals;
    }

    // keccak256(abi.encode(uint256(keccak256("camino.messenger.storage.BookingTokenV2")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant BookingTokenV2StorageLocation =
        0x96683f18b5720c3d007c342e56e2dbf25d018d9c7f77c4217427b152a0bd6100;

    /**
     * @notice Retrieves the BookingTokenV2 storage struct from the designated storage slot.
     *
     * @return $ The storage struct reference
     */
    function _getBookingTokenV2Storage() private pure returns (BookingTokenV2Storage storage $) {
        assembly {
            $.slot := BookingTokenV2StorageLocation
        }
    }

    /***************************************************
     *                   EVENTS                        *
     ***************************************************/

    /**
     * @notice Event emitted when a cancellation is initiated.
     *
     * @param tokenId token id
     * @param initiatedBy address that initiated the cancellation
     * @param refundAmount proposed refund amount
     * @param refundCurrency ERC20 token address for the refund
     */
    event CancellationInitiated(
        uint256 indexed tokenId,
        address indexed initiatedBy,
        uint256 refundAmount,
        address refundCurrency
    );

    /**
     * @notice Event emitted when a cancellation is accepted.
     *
     * @param tokenId token id
     * @param acceptedBy address that accepted the cancellation
     * @param refundAmount final refund amount
     * @param refundCurrency ERC20 token address for the refund
     */
    event CancellationAccepted(
        uint256 indexed tokenId,
        address indexed acceptedBy,
        uint256 refundAmount,
        address refundCurrency
    );

    /**
     * @notice Event emitted when a cancellation proposal is countered.
     *
     * @param tokenId token id
     * @param counteredBy address that countered the proposal
     * @param newRefundAmount new proposed refund amount
     * @param newRefundCurrency new proposed ERC20 token address for the refund
     */
    event CancellationCountered(
        uint256 indexed tokenId,
        address indexed counteredBy,
        uint256 newRefundAmount,
        address newRefundCurrency
    );

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
     * @notice Error for when there is no active cancellation proposal.
     *
     * @param tokenId The token id for which there is no active proposal
     */
    error NoActiveCancellationProposal(uint256 tokenId);

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
     * @notice Error for when the caller is not authorized to cancel a proposal.
     *
     * @param caller The address of the caller
     */
    error NotAuthorizedToCancelProposal(address caller);

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
     * @notice Mints a new token with a reservation for a specific address.
     *
     * @param reservedFor The CM Account address that can buy the token
     * @param uri The URI of the token
     * @param expirationTimestamp The expiration timestamp
     * @param price The price of the token
     * @param paymentToken The token used to pay for the reservation. If address(0) then native.
     */
    function safeMintWithReservation(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken
    ) public override onlyCMAccount(msg.sender) {
        super.safeMintWithReservation(reservedFor, uri, expirationTimestamp, price, paymentToken);
        uint256 tokenId = _getBookingTokenStorage()._nextTokenId - 1;
        _getBookingTokenV2Storage()._originalSupplier[tokenId] = msg.sender;
    }

    /**
     * @notice Initiates a cancellation for a bought token.
     *
     * @param tokenId The token id to initiate the cancellation for
     * @param refundAmount The proposed refund amount in wei
     * @param refundCurrency The ERC20 token address for the refund
     */
    function initiateCancellation(uint256 tokenId, uint256 refundAmount, address refundCurrency) external {
        address owner = _requireOwned(tokenId);
        BookingTokenV2Storage storage $ = _getBookingTokenV2Storage();
        address supplier = $._originalSupplier[tokenId];
        if (msg.sender != owner && msg.sender != supplier) {
            revert NotAuthorizedToInitiateCancellation(msg.sender);
        }

        $._cancellationProposals[tokenId] = CancellationProposal({
            refundAmount: refundAmount,
            refundCurrency: refundCurrency,
            initiatedBy: msg.sender,
            isActive: true
        });

        emit CancellationInitiated(tokenId, msg.sender, refundAmount, refundCurrency);
    }

    /**
     * @notice Accepts a cancellation proposal for a bought token.
     *
     * @param tokenId The token id to accept the cancellation for
     */
    function acceptCancellation(uint256 tokenId) external {
        BookingTokenV2Storage storage $ = _getBookingTokenV2Storage();
        address owner = _requireOwned(tokenId);
        CancellationProposal memory proposal = $._cancellationProposals[tokenId];
        if (!proposal.isActive) {
            revert NoActiveCancellationProposal(tokenId);
        }

        address supplier = $._originalSupplier[tokenId];
        if (msg.sender == proposal.initiatedBy) {
            revert NotAuthorizedToAcceptCancellation(msg.sender);
        }
        if (
            (proposal.initiatedBy == owner && msg.sender != supplier) ||
            (proposal.initiatedBy == supplier && msg.sender != owner)
        ) {
            revert NotAuthorizedToAcceptCancellation(msg.sender);
        }

        // Finalize the cancellation
        delete $._cancellationProposals[tokenId];
        _burn(tokenId);

        // Emit the cancellation accepted event
        emit CancellationAccepted(tokenId, msg.sender, proposal.refundAmount, proposal.refundCurrency);
    }

    /**
     * @notice Counters a cancellation proposal with a new proposal.
     *
     * @param tokenId The token id to counter the cancellation for
     * @param newRefundAmount The new proposed refund amount in wei
     * @param newRefundCurrency The new ERC20 token address for the refund
     */
    function counterCancellationProposal(uint256 tokenId, uint256 newRefundAmount, address newRefundCurrency) external {
        BookingTokenV2Storage storage $ = _getBookingTokenV2Storage();
        address owner = _requireOwned(tokenId);
        CancellationProposal storage proposal = $._cancellationProposals[tokenId];
        if (!proposal.isActive) {
            revert NoActiveCancellationProposal(tokenId);
        }

        address supplier = $._originalSupplier[tokenId];
        if (msg.sender == proposal.initiatedBy) {
            revert NotAuthorizedToCounterCancellation(msg.sender);
        }
        if (
            (proposal.initiatedBy == owner && msg.sender != supplier) ||
            (proposal.initiatedBy == supplier && msg.sender != owner)
        ) {
            revert NotAuthorizedToCounterCancellation(msg.sender);
        }

        // Update the proposal with the new values
        proposal.refundAmount = newRefundAmount;
        proposal.refundCurrency = newRefundCurrency;

        // Emit the countered proposal event
        emit CancellationCountered(tokenId, msg.sender, newRefundAmount, newRefundCurrency);
    }

    /**
     * @notice Cancels an active cancellation proposal. Only the initiator can cancel.
     *
     * @param tokenId The token id for which to cancel the proposal
     */
    function cancelCancellationProposal(uint256 tokenId) external {
        BookingTokenV2Storage storage $ = _getBookingTokenV2Storage();
        CancellationProposal storage proposal = $._cancellationProposals[tokenId];
        if (!proposal.isActive) {
            revert NoActiveCancellationProposal(tokenId);
        }
        if (msg.sender != proposal.initiatedBy) {
            revert NotAuthorizedToCancelProposal(msg.sender);
        }

        // Cancel the proposal
        delete $._cancellationProposals[tokenId];

        // Emit the cancellation proposal canceled event
        emit CancellationProposalCanceled(tokenId, msg.sender);
    }

    /**
     * @notice Retrieves the current cancellation proposal status for a given token.
     *
     * @param tokenId The token id to check the proposal status for
     * @return refundAmount The proposed refund amount
     * @return refundCurrency The address of the proposed refund currency
     * @return initiatedBy The address that initiated the cancellation
     * @return isActive The status of the cancellation proposal
     */
    function getCancellationProposalStatus(
        uint256 tokenId
    ) external view returns (uint256 refundAmount, address refundCurrency, address initiatedBy, bool isActive) {
        BookingTokenV2Storage storage $ = _getBookingTokenV2Storage();
        CancellationProposal memory proposal = $._cancellationProposals[tokenId];
        return (proposal.refundAmount, proposal.refundCurrency, proposal.initiatedBy, proposal.isActive);
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
        BookingTokenV2Storage storage $ = _getBookingTokenV2Storage();
        if ($._cancellationProposals[tokenId].isActive) {
            revert TokenHasActiveCancellationProposal(tokenId);
        }
    }
}
