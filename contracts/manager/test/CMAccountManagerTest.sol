// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Account Manager V2 for Testing Upgrades

/**
 * TESTING ONLY - NOT FOR PRODUCTION
 */

pragma solidity 0.8.24;

import { CMAccountManager } from "../CMAccountManager.sol";

contract CMAccountManagerTest is CMAccountManager {
    function getVersion() public pure returns (string memory) {
        return "TESTING";
    }

    // function setCMAccountInfo(address account, CMAccountInfo memory info) public onlyRole(DEFAULT_ADMIN_ROLE) {
    //     _setCMAccountInfo(account, info);
    // }
}
