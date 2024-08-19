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
    event GasMoneyWithdrawalLimitUpdated(uint256 limit);
    event GasMoneyWithdrawalPeriodUpdated(uint256 period);

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

    function _withdrawGasMoney(uint256 amount) internal {
        GasMoneyStorage storage $ = _getGasMoneyStorage();

        uint256 withdrawalLimit = $._withdrawalLimit;
        uint256 withdrawalPeriod = $._withdrawalPeriod;

        // Ensure the withdrawal does not exceed the allowed limit
        if (amount > withdrawalLimit) {
            revert WithdrawalLimitExceeded(withdrawalLimit, amount);
        }

        // Get timestamps
        uint256 currentTime = block.timestamp;
        uint256 withdrawalPeriodStart = $._withdrawalPeriodStart[msg.sender];

        // Reset the withdrawn amount if a new period has started. If more time then
        // the withdrawal period has passed, it is allowed to withdraw full amount.
        if (currentTime > withdrawalPeriodStart + withdrawalPeriod) {
            $._withdrawnAmount[msg.sender] = 0;
            $._withdrawalPeriodStart[msg.sender] = currentTime;
        }

        // Ensure the withdrawal does not exceed the allowed limit for the period
        if ($._withdrawnAmount[msg.sender] + amount > withdrawalLimit) {
            revert WithdrawalLimitExceededForPeriod(withdrawalLimit, amount);
        }

        // Update the withdrawn amount
        // FIXME: Not likely but still, check overflow (safe math libs?)
        $._withdrawnAmount[msg.sender] += amount;

        // Transfer the gas money
        payable(msg.sender).sendValue(amount);

        emit GasMoneyWithdrawal(msg.sender, amount);
    }

    function _setGasMoneyWithdrawalLimit(uint256 limit) internal {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        $._withdrawalLimit = limit;

        emit GasMoneyWithdrawalLimitUpdated(limit);
    }

    function _setGasMoneyWithdrawalPeriod(uint256 period) internal {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        $._withdrawalPeriod = period;

        emit GasMoneyWithdrawalPeriodUpdated(period);
    }

    function getGasMoneyWithdrawalLimit() public view returns (uint256) {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        return $._withdrawalLimit;
    }

    function getGasMoneyWithdrawalPeriod() public view returns (uint256) {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        return $._withdrawalPeriod;
    }

    function getGasMoneyWithdrawalPeriodStart(address account) public view returns (uint256) {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        return $._withdrawalPeriodStart[account];
    }

    function getGasMoneyWithdrawnAmount(address account) public view returns (uint256) {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        return $._withdrawnAmount[account];
    }
}
