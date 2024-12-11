// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.24;

/**
 * @title CancellationUtils
 * @notice Library for packing and unpacking cancellation-related data into a single
 * uint256. Each field uses 16 bits, allowing values up to 65,535.
 *
 * Layout:
 * [rejection_reason_version(16)][rejection_reason(16)][cancellation_reason_version(16)][cancellation_reason(16)]
 * [------------------REJECTION_MASK------------------][-------------------CANCELLATION_MASK--------------------]
 *
 * The versions values are the package versions from the [Camino Messenger Protocol](https://github.com/chain4travel/camino-messenger-protocol)
 *
 * Reason Enums: camino-messenger-protocol/proto/cmp/services/cancellation/<version>/reason.proto
 */
library CancellationUtils {
    // Constants for bit manipulation
    uint256 private constant SEGMENT_MASK = 0xFFFF; // 16 bits mask
    uint256 private constant CANCELLATION_REASON_SHIFT = 0; // First 16 bits
    uint256 private constant CANCELLATION_VERSION_SHIFT = 16; // Next 16 bits
    uint256 private constant REJECTION_REASON_SHIFT = 32; // Next 16 bits
    uint256 private constant REJECTION_VERSION_SHIFT = 48; // Final 16 bits

    // Masks for different sections
    uint256 private constant CANCELLATION_MASK = 0xFFFFFFFF;
    uint256 private constant REJECTION_MASK = 0xFFFFFFFF00000000;

    /**
     * @notice Packs cancellation and rejection data into a single uint256
     * @param cancellationReason The reason for cancellation
     * @param cancellationVersion Version of the cancellation reason enum
     * @param rejectionReason The reason for rejection
     * @param rejectionVersion Version of the rejection reason enum
     * @return packed The packed uint256 containing all data
     */
    function packReasons(
        uint16 cancellationReason,
        uint16 cancellationVersion,
        uint16 rejectionReason,
        uint16 rejectionVersion
    ) internal pure returns (uint256) {
        // Pack all values into a single uint256
        return
            (uint256(rejectionVersion) << REJECTION_VERSION_SHIFT) |
            (uint256(rejectionReason) << REJECTION_REASON_SHIFT) |
            (uint256(cancellationVersion) << CANCELLATION_VERSION_SHIFT) |
            (uint256(cancellationReason) << CANCELLATION_REASON_SHIFT);
    }

    /**
     * @notice Unpacks a uint256 into its component values
     * @param packed The packed uint256 containing all data
     * @return cancellationReason The reason for cancellation
     * @return cancellationVersion Version of the cancellation reason enum
     * @return rejectionReason The reason for rejection
     * @return rejectionVersion Version of the rejection reason enum
     */
    function unpackReasons(
        uint256 packed
    )
        internal
        pure
        returns (uint16 cancellationReason, uint16 cancellationVersion, uint16 rejectionReason, uint16 rejectionVersion)
    {
        cancellationReason = uint16((packed >> CANCELLATION_REASON_SHIFT) & SEGMENT_MASK);
        cancellationVersion = uint16((packed >> CANCELLATION_VERSION_SHIFT) & SEGMENT_MASK);
        rejectionReason = uint16((packed >> REJECTION_REASON_SHIFT) & SEGMENT_MASK);
        rejectionVersion = uint16((packed >> REJECTION_VERSION_SHIFT) & SEGMENT_MASK);
    }

    /**
     * @notice Updates only the cancellation-related fields in the packed value
     * @param packed The existing packed uint256
     * @param cancellationReason New cancellation reason
     * @param cancellationVersion New cancellation version
     * @return newPacked The updated packed value with only cancellation fields modified
     */
    function updateCancellationReason(
        uint256 packed,
        uint16 cancellationReason,
        uint16 cancellationVersion
    ) internal pure returns (uint256 newPacked) {
        // Clear old cancellation values while preserving rejection values
        newPacked = packed & ~CANCELLATION_MASK;

        // Add new cancellation values
        newPacked |=
            (uint256(cancellationVersion) << CANCELLATION_VERSION_SHIFT) |
            (uint256(cancellationReason) << CANCELLATION_REASON_SHIFT);
    }

    /**
     * @notice Updates only the rejection-related fields in the packed value
     * @param packed The existing packed uint256
     * @param rejectionReason New rejection reason
     * @param rejectionVersion New rejection version
     * @return newPacked The updated packed value with only rejection fields modified
     */
    function updateRejectionReason(
        uint256 packed,
        uint16 rejectionReason,
        uint16 rejectionVersion
    ) internal pure returns (uint256 newPacked) {
        // Clear old rejection values while preserving cancellation values
        newPacked = packed & ~REJECTION_MASK;

        // Add new rejection values
        newPacked |=
            (uint256(rejectionVersion) << REJECTION_VERSION_SHIFT) |
            (uint256(rejectionReason) << REJECTION_REASON_SHIFT);
    }

    /**
     * @notice Gets only the cancellation-related values from the packed data
     * @param packed The packed uint256 containing all data
     * @return cancellationReason The reason for cancellation
     * @return cancellationVersion Version of the cancellation reason enum
     */
    function getCancellationReason(
        uint256 packed
    ) internal pure returns (uint16 cancellationReason, uint16 cancellationVersion) {
        cancellationReason = uint16((packed >> CANCELLATION_REASON_SHIFT) & SEGMENT_MASK);
        cancellationVersion = uint16((packed >> CANCELLATION_VERSION_SHIFT) & SEGMENT_MASK);
    }

    /**
     * @notice Gets only the rejection-related values from the packed data
     * @param packed The packed uint256 containing all data
     * @return rejectionReason The reason for rejection
     * @return rejectionVersion Version of the rejection reason enum
     */
    function getRejectionReason(
        uint256 packed
    ) internal pure returns (uint16 rejectionReason, uint16 rejectionVersion) {
        rejectionReason = uint16((packed >> REJECTION_REASON_SHIFT) & SEGMENT_MASK);
        rejectionVersion = uint16((packed >> REJECTION_VERSION_SHIFT) & SEGMENT_MASK);
    }
}
