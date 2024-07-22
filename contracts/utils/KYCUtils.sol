// SPDX-License-Identifier: UNLICENSED
//
// Camino KYC Utilities

pragma solidity ^0.8.24;

interface ICaminoAdmin {
    function getKycState(address account) external view returns (uint256);
}

library KYCUtils {
    /**
     * @notice Admin contract address
     */
    address public constant ADMIN_ADDR = 0x010000000000000000000000000000000000000a;

    /**
     * Constants for KYC states
     */
    uint256 public constant KYC_VERIFIED = 1 << 0; // Bit 0
    uint256 public constant KYC_EXPIRED = 1 << 1; // Bit 1
    uint256 public constant KYB_VERIFIED = 1 << 8; // Bit 8

    /**
     * Errors
     */
    error NotKYCVerified(address account);
    error NotKYBVerified(address account);
    error NotVerified(address account);

    /**
     * @dev Returns KYC state from the CaminoAdmin contract
     * @param account address to check the state
     */
    function getKYCState(address account) internal view returns (uint256) {
        return ICaminoAdmin(ADMIN_ADDR).getKycState(account);
    }

    /**
     * @dev Returns true if the address is KYC verified
     * @param account address to check the state
     */
    function isKYCVerified(address account) internal view returns (bool) {
        uint256 kycState = getKYCState(account);
        return kycState & KYC_VERIFIED != 0;
    }

    /**
     * @dev Returns true if the address is KYB verified
     * @param account address to check the state
     */
    function isKYBVerified(address account) internal view returns (bool) {
        uint256 kycState = getKYCState(account);
        return kycState & KYB_VERIFIED != 0;
    }

    /**
     * @dev Returns true if the address is KYC or KYB verified
     * @param account address to check the state
     */
    function isVerified(address account) internal view returns (bool) {
        uint256 kycState = getKYCState(account);
        return (kycState & KYC_VERIFIED != 0) || (kycState & KYB_VERIFIED != 0);
    }

    /**
     * @dev Reverts with `NotKYCVerified(account)` if the account is not KYC verified.
     * @param account address to check the state
     */
    function requireKYCVerified(address account) internal view {
        if (!isKYCVerified(account)) {
            revert NotKYCVerified(account);
        }
    }

    /**
     * @dev Reverts with `NotKYBVerified(account)` if the account is not KYB verified.
     * @param account address to check the state
     */
    function requireKYBVerified(address account) internal view {
        if (!isKYBVerified(account)) {
            revert NotKYBVerified(account);
        }
    }

    /**
     * @dev Reverts with `NotVerified(account)` if the account is not KYC or KYB verified.
     * @param account address to check the state
     */
    function requireVerified(address account) internal view {
        if (!isVerified(account)) {
            revert NotVerified(account);
        }
    }
}
