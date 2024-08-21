// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Gas Money Manager

pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/Address.sol";

abstract contract GasMoneyManager is Initializable {
    using Address for address payable;

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    struct GasMoneyStorage {
        mapping(address => uint256) _withdrawalPeriodStart;
        mapping(address => uint256) _withdrawnAmount;
        uint256 _withdrawalLimit;
        uint256 _withdrawalPeriod;
    }

    // keccak256(abi.encode(uint256(keccak256("camino.messenger.storage.GasMoneyManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant GasMoneyStorageLocation =
        0x99a652063088b6badaeb0c7f680676baf720654b4f86f50167944489af637d00;

    function _getGasMoneyStorage() private pure returns (GasMoneyStorage storage $) {
        assembly {
            $.slot := GasMoneyStorageLocation
        }
    }

    /***************************************************
     *                   EVENTS                        *
     ***************************************************/

    event GasMoneyWithdrawal(address indexed withdrawer, uint256 amount);
    event GasMoneyWithdrawalUpdated(uint256 limit, uint256 period);
    //event GasMoneyWithdrawalPeriodUpdated(uint256 period);

    /***************************************************
     *                   ERRORS                        *
     ***************************************************/

    error WithdrawalLimitExceeded(uint256 limit, uint256 amount);
    error WithdrawalLimitExceededForPeriod(uint256 limit, uint256 amount);

    /***************************************************
     *               INITIALIZATION                    *
     ***************************************************/

    function __GasMoneyManager_init(uint256 withdrawalLimit, uint256 withdrawalPeriod) internal onlyInitializing {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        $._withdrawalLimit = withdrawalLimit;
        $._withdrawalPeriod = withdrawalPeriod;
    }

    /***************************************************
     *                   LOGIC                        *
     ***************************************************/

    /**
     * @dev Withdraw gas money
     *
     * This functions is intended to be called by the bot to withdraw gas money.
     * Inheriting contract should restrict who can call this within a public
     * function.
     */
    function _withdrawGasMoney(uint256 amount) internal {
        GasMoneyStorage storage $ = _getGasMoneyStorage();

        // uint256 withdrawalLimit = $._withdrawalLimit;
        // uint256 withdrawalPeriod = $._withdrawalPeriod;

        // Ensure the withdrawal does not exceed the allowed limit
        if (amount > $._withdrawalLimit) {
            revert WithdrawalLimitExceeded($._withdrawalLimit, amount);
        }

        // Get timestamps
        uint256 currentTime = block.timestamp;

        // Reset the withdrawn amount if a new period has started. If more time then
        // the withdrawal period has passed, it is allowed to withdraw full amount.
        if (currentTime > $._withdrawalPeriodStart[msg.sender] + $._withdrawalPeriod) {
            $._withdrawnAmount[msg.sender] = 0;
            $._withdrawalPeriodStart[msg.sender] = currentTime;
        }

        // Ensure the withdrawal does not exceed the allowed limit for the period
        if ($._withdrawnAmount[msg.sender] + amount > $._withdrawalLimit) {
            revert WithdrawalLimitExceededForPeriod($._withdrawalLimit, amount);
        }

        // Update the withdrawn amount
        // FIXME: Not likely but still, check overflow (safe math libs?)
        $._withdrawnAmount[msg.sender] += amount;

        // Transfer the gas money
        payable(msg.sender).sendValue(amount);

        emit GasMoneyWithdrawal(msg.sender, amount);
    }

    /**
     * @dev Set the gas money withdrawal limit and period
     *
     * @param limit the withdrawal limit for the period
     * @param period the withdrawal period in seconds
     */
    function _setGasMoneyWithdrawal(uint256 limit, uint256 period) internal {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        $._withdrawalLimit = limit;
        $._withdrawalPeriod = period;

        emit GasMoneyWithdrawalUpdated(limit, period);
    }

    function getGasMoneyWithdrawal() public view returns (uint256, uint256) {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        return ($._withdrawalLimit, $._withdrawalPeriod);
    }

    /**
     * @dev Get the gas money withdrawal details for an account
     *
     * @param account address of the account
     * @return periodStart timestamp of the withdrawal period start
     * @return withdrawnAmount amount withdrawn within the period
     */
    function getGasMoneyWithdrawalForAccount(
        address account
    ) public view returns (uint256 periodStart, uint256 withdrawnAmount) {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        return ($._withdrawalPeriodStart[account], $._withdrawnAmount[account]);
    }
}
