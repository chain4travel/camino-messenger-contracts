/**
 * @dev CMAccount tests
 */
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers } = require("hardhat");

// Fixtures
const {
    setupSigners,
    developerFeeBp,
    deployCMAccountManagerFixture,
    deployCMAccountImplFixture,
    deployCMAccountManagerWithCMAccountImplFixture,
    deployAndConfigureAllFixture,
    deployCMAccountWithDepositFixture,
} = require("./utils/fixtures");

describe("GasMoneyManager", function () {
    describe("Main", function () {
        it("should initialize gas money manager correctly", async function () {
            const { cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const expectedLimit = ethers.parseEther("10"); // 10 CAM
            const expectedPeriod = 24 * 60 * 60; // 24 hours

            expect(await cmAccount.getGasMoneyWithdrawal()).to.be.deep.equal([expectedLimit, expectedPeriod]);
            //expect(await cmAccount.getGasMoneyWithdrawalPeriod()).to.be.equal(expectedPeriod);
        });

        it("should set gas money limit and period correctly", async function () {
            const { cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const expectedLimit = ethers.parseEther("10"); // 10 CAM
            const expectedPeriod = 24 * 60 * 60; // 24 hours

            expect(await cmAccount.getGasMoneyWithdrawal()).to.be.deep.equal([expectedLimit, expectedPeriod]);
            //expect(await cmAccount.getGasMoneyWithdrawalPeriod()).to.be.equal(expectedPeriod);

            const newLimit = ethers.parseEther("20"); // 20 CAM
            const newPeriod = 48 * 60 * 60; // 48 hours

            await expect(cmAccount.connect(signers.cmAccountAdmin).setGasMoneyWithdrawal(newLimit, newPeriod))
                .to.emit(cmAccount, "GasMoneyWithdrawalUpdated")
                .withArgs(newLimit, newPeriod);

            // await expect(cmAccount.connect(signers.cmAccountAdmin).setGasMoneyWithdrawalPeriod(newPeriod))
            //     .to.emit(cmAccount, "GasMoneyWithdrawalPeriodUpdated")
            //     .withArgs(newPeriod);

            expect(await cmAccount.getGasMoneyWithdrawal()).to.be.deep.equal([newLimit, newPeriod]);
            //expect(await cmAccount.getGasMoneyWithdrawalPeriod()).to.be.equal(newPeriod);
        });

        it("should withdraw gas money", async function () {
            const { cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const withdrawer = signers.withdrawer;

            // Register withdrawer as a bot
            await expect(cmAccount.connect(signers.cmAccountAdmin).addMessengerBot(withdrawer.address))
                .to.emit(cmAccount, "MessengerBotAdded")
                .withArgs(withdrawer.address);

            // Withdraw
            const withdrawAmount = ethers.parseEther("1");

            const withdrawTx = cmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount);
            await expect(withdrawTx).to.changeEtherBalances([cmAccount, withdrawer], [-withdrawAmount, withdrawAmount]);
            await expect(withdrawTx)
                .to.emit(cmAccount, "GasMoneyWithdrawal")
                .withArgs(withdrawer.address, withdrawAmount);
        });

        it("should revert if not allowed to withdraw gas money", async function () {
            const { cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const withdrawer = signers.withdrawer;

            // Do not add withdrawer as a bot.

            // Withdraw
            const withdrawAmount = ethers.parseEther("1");

            await expect(cmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount))
                .to.revertedWithCustomError(cmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(withdrawer.address, cmAccount.GAS_WITHDRAWER_ROLE());
        });

        it("should revert if amount is over the limit", async function () {
            const { cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const withdrawer = signers.withdrawer;
            const expectedLimit = ethers.parseEther("10"); // 10 CAM

            // Register withdrawer as a bot
            await expect(cmAccount.connect(signers.cmAccountAdmin).addMessengerBot(withdrawer.address))
                .to.emit(cmAccount, "MessengerBotAdded")
                .withArgs(withdrawer.address);

            // Withdraw
            const withdrawAmount = ethers.parseEther("11"); // 11 CAM, over the limit

            await expect(cmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount))
                .to.revertedWithCustomError(cmAccount, "WithdrawalLimitExceeded")
                .withArgs(expectedLimit, withdrawAmount);
        });

        it("should revert if amount is over the limit for the period", async function () {
            const { cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const withdrawer = signers.withdrawer;
            const expectedLimit = ethers.parseEther("10"); // 10 CAM

            // Register withdrawer as a bot
            await expect(cmAccount.connect(signers.cmAccountAdmin).addMessengerBot(withdrawer.address))
                .to.emit(cmAccount, "MessengerBotAdded")
                .withArgs(withdrawer.address);

            // Withdraw
            const withdrawAmount = ethers.parseEther("1"); // Start with 1 CAM

            const withdrawTx1 = cmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount);
            await expect(withdrawTx1).to.changeEtherBalances(
                [cmAccount, withdrawer],
                [-withdrawAmount, withdrawAmount],
            );
            await expect(withdrawTx1)
                .to.emit(cmAccount, "GasMoneyWithdrawal")
                .withArgs(withdrawer.address, withdrawAmount);

            const withdrawAmount2 = ethers.parseEther("7"); // Withdraw 7 CAM, total 8

            const withdrawTx2 = cmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount2);
            await expect(withdrawTx2).to.changeEtherBalances(
                [cmAccount, withdrawer],
                [-withdrawAmount2, withdrawAmount2],
            );
            await expect(withdrawTx2)
                .to.emit(cmAccount, "GasMoneyWithdrawal")
                .withArgs(withdrawer.address, withdrawAmount2);

            const withdrawAmount3 = ethers.parseEther("3"); // Withdraw 3 CAM, total 11, over the limit

            await expect(cmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount3))
                .to.revertedWithCustomError(cmAccount, "WithdrawalLimitExceededForPeriod")
                .withArgs(expectedLimit, withdrawAmount3);
        });

        it("should allow withdrawal after period resets", async function () {
            const { cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const withdrawer = signers.withdrawer;
            const expectedLimit = ethers.parseEther("10"); // 10 CAM

            // Register withdrawer as a bot
            await expect(cmAccount.connect(signers.cmAccountAdmin).addMessengerBot(withdrawer.address))
                .to.emit(cmAccount, "MessengerBotAdded")
                .withArgs(withdrawer.address);

            // Withdraw
            const withdrawAmount = ethers.parseEther("10"); // withdraw all 10 CAM

            const withdrawTx1 = cmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount);
            await expect(withdrawTx1).to.changeEtherBalances(
                [cmAccount, withdrawer],
                [-withdrawAmount, withdrawAmount],
            );
            await expect(withdrawTx1)
                .to.emit(cmAccount, "GasMoneyWithdrawal")
                .withArgs(withdrawer.address, withdrawAmount);

            const withdrawAmount2 = ethers.parseEther("3"); // Try to withdraw 3 CAM, over the limit

            await expect(cmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount2))
                .to.revertedWithCustomError(cmAccount, "WithdrawalLimitExceededForPeriod")
                .withArgs(expectedLimit, withdrawAmount2);

            // Advance time by 24 hours
            await network.provider.send("evm_increaseTime", [24 * 60 * 60]);
            await network.provider.send("evm_mine");

            // Withdraw again
            const withdrawAmount3 = ethers.parseEther("10"); // Try to withdraw the limit as the period has been reset

            const withdrawTx3 = cmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount3);
            await expect(withdrawTx3).to.changeEtherBalances(
                [cmAccount, withdrawer],
                [-withdrawAmount3, withdrawAmount3],
            );
            await expect(withdrawTx3)
                .to.emit(cmAccount, "GasMoneyWithdrawal")
                .withArgs(withdrawer.address, withdrawAmount3);
        });
    });
});
