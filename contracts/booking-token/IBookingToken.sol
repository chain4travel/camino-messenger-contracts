// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IBookingToken {
    function safeMintWithReservation(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken
    ) external;

    function buyReservedToken(uint256 tokenId) external payable;

    function getReservationPrice(uint256 tokenId) external view returns (uint256 price, IERC20 paymentToken);

    /**
     * @notice Initiates a cancellation for a bought token.
     *
     * @param tokenId The token id to initiate the cancellation for
     * @param refundAmount The proposed refund amount in wei
     * @param refundCurrency The ERC20 token address for the refund
     */
    function initiateCancellation(uint256 tokenId, uint256 refundAmount, address refundCurrency) external;

    /**
     * @notice Accepts a cancellation proposal for a bought token.
     *
     * @param tokenId The token id to accept the cancellation for
     */
    function acceptCancellation(uint256 tokenId) external;

    /**
     * @notice Counters a cancellation proposal with a new proposal.
     *
     * @param tokenId The token id to counter the cancellation for
     * @param newRefundAmount The new proposed refund amount in wei
     * @param newRefundCurrency The new ERC20 token address for the refund
     */
    function counterCancellationProposal(uint256 tokenId, uint256 newRefundAmount, address newRefundCurrency) external;

    /**
     * @notice Cancels an active cancellation proposal. Only the initiator can cancel.
     *
     * @param tokenId The token id for which to cancel the proposal
     */
    function cancelCancellationProposal(uint256 tokenId) external;

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
    ) external view returns (uint256 refundAmount, address refundCurrency, address initiatedBy, bool isActive);
}
