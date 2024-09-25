// SPDX-License-Identifier: UNLICENSED
//
// Null USD Contract for testing purposes

pragma solidity 0.8.24;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract NullUSD is ERC20 {
    constructor() ERC20("NullUSD", "NUSD") {
        _mint(msg.sender, 1000000 * 10 ** decimals());
    }
}
