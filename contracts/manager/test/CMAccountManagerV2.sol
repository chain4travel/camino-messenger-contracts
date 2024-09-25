// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Account Manager V2 for Testing Upgrades

/**
 * TESTING ONLY - NOT FOR PRODUCTION
 */

pragma solidity 0.8.24;

import { CMAccountManager } from "../CMAccountManager.sol";

contract CMAccountManagerV2 is CMAccountManager {
    function getVersion() public pure returns (string memory) {
        return "V2";
    }
}
