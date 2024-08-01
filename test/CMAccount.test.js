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

describe("CMAccount", function () {
    describe("Upgrade", function () {
        it("should upgrade to new implementation address", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            // Old implementation
            const oldImplementationAddress = await cmAccountManager.getAccountImplementation();

            // Create a new implementation for CMAccount
            const CMAccountImplV2 = await ethers.getContractFactory("CMAccount");
            const cmAccountImplV2 = await CMAccountImplV2.deploy();
            await cmAccountImplV2.waitForDeployment();
            const newImplementationAddress = await cmAccountImplV2.getAddress();

            // Set new implementation on the manager
            await cmAccountManager.connect(signers.managerVersioner).setAccountImplementation(newImplementationAddress);
            await expect(await cmAccountManager.getAccountImplementation()).to.be.equal(newImplementationAddress);

            // Upgrade the account
            await expect(cmAccount.connect(signers.cmAccountUpgrader).upgradeToAndCall(newImplementationAddress, "0x"))
                .to.emit(cmAccount, "CMAccountUpgraded")
                .withArgs(oldImplementationAddress, newImplementationAddress);
        });

        it("should revert upgrade if implementation address does not match", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            // Old implementation
            const oldImplementationAddress = await cmAccountManager.getAccountImplementation();

            // Create a new implementation for CMAccount
            const CMAccountImplV2 = await ethers.getContractFactory("CMAccount");
            const cmAccountImplV2 = await CMAccountImplV2.deploy();
            await cmAccountImplV2.waitForDeployment();
            const newImplementationAddress = await cmAccountImplV2.getAddress();

            // SKIP: DO NOT set new implementation on the manager here

            // Try to upgrade the account
            await expect(cmAccount.connect(signers.cmAccountUpgrader).upgradeToAndCall(newImplementationAddress, "0x"))
                .to.be.revertedWithCustomError(cmAccount, "CMAccountImplementationMismatch")
                .withArgs(oldImplementationAddress, newImplementationAddress);
        });

        it("should revert upgrade if address is not uups upgradeable", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            // Old implementation
            const oldImplementationAddress = await cmAccountManager.getAccountImplementation();

            // Create a new implementation for CMAccount
            const dummyAccountImpl = await ethers.getContractFactory("Dummy");
            const dummyAccountImplV2 = await dummyAccountImpl.deploy();
            await dummyAccountImplV2.waitForDeployment();
            const newImplementationAddress = await dummyAccountImplV2.getAddress();

            // Set new implementation on the manager
            await cmAccountManager.connect(signers.managerVersioner).setAccountImplementation(newImplementationAddress);
            await expect(await cmAccountManager.getAccountImplementation()).to.be.equal(newImplementationAddress);

            // Upgrade the account
            await expect(cmAccount.connect(signers.cmAccountUpgrader).upgradeToAndCall(newImplementationAddress, "0x"))
                .to.be.revertedWithCustomError(cmAccount, "ERC1967InvalidImplementation")
                .withArgs(newImplementationAddress);
        });

        it("should revert upgrade if address is same with the current one", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            // Old implementation
            const oldImplementationAddress = await cmAccountManager.getAccountImplementation();

            // Upgrade the account
            await expect(cmAccount.connect(signers.cmAccountUpgrader).upgradeToAndCall(oldImplementationAddress, "0x"))
                .to.be.revertedWithCustomError(cmAccount, "CMAccountNoUpgradeNeeded")
                .withArgs(oldImplementationAddress, oldImplementationAddress);
        });
    });

    describe("Registering Bots", function () {
        it("should register bots correctly", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const CHEQUE_OPERATOR_ROLE = await cmAccount.CHEQUE_OPERATOR_ROLE();
            const botAddr = signers.chequeOperator.address;

            // Grant CHEQUE_OPERATOR_ROLE
            await expect(cmAccount.connect(signers.cmAccountAdmin).grantRole(CHEQUE_OPERATOR_ROLE, botAddr))
                .to.emit(cmAccount, "RoleGranted")
                .withArgs(CHEQUE_OPERATOR_ROLE, botAddr, signers.cmAccountAdmin.address);

            await expect(await cmAccount.isBotAllowed(botAddr)).to.be.true;
        });
    });

    describe("Deposit", function () {
        it("should allow anyone to send funds", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const anyone = signers.otherAccount1;

            const anyoneInitialBalance = await ethers.provider.getBalance(anyone.address);
            const cmAccountInitialBalance = await ethers.provider.getBalance(cmAccount.getAddress());

            const depositAmount = ethers.parseEther("1");

            // Sender
            const depositTx = {
                to: cmAccount.getAddress(),
                value: depositAmount,
            };

            await expect(await anyone.sendTransaction(depositTx)).to.not.be.reverted;

            // Check balances
            // Sender balance should be lower than the difference between their initial balance and the deposit
            expect(await ethers.provider.getBalance(anyone.address)).to.be.lt(anyoneInitialBalance - depositAmount);

            // CMAccount balance should be equal to the sum of the initial balance and the deposit
            expect(await ethers.provider.getBalance(cmAccount.getAddress())).to.be.equal(
                cmAccountInitialBalance + depositAmount,
            );
        });
    });

    describe("Withdraw", function () {
        it("should allow withdrawer role to withdraw", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            const withdrawer = signers.withdrawer;
            const withdrawAmount = ethers.parseEther("0.5");

            // Withdraw
            const withdrawTx = cmAccount.connect(withdrawer).withdraw(withdrawer.address, withdrawAmount);
            await expect(withdrawTx).to.changeEtherBalances([cmAccount, withdrawer], [-withdrawAmount, withdrawAmount]);
            await expect(withdrawTx).to.emit(cmAccount, "Withdraw").withArgs(withdrawer.address, withdrawAmount);
        });

        it("should revert if not withdrawer role", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            const withdrawer = signers.otherAccount1;
            const withdrawAmount = ethers.parseEther("0.5");

            const WITHDRAWER_ROLE = await cmAccount.WITHDRAWER_ROLE();

            // Withdraw
            const withdrawTx = cmAccount.connect(withdrawer).withdraw(withdrawer.address, withdrawAmount);
            await expect(withdrawTx)
                .to.be.revertedWithCustomError(cmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(withdrawer.address, WITHDRAWER_ROLE);
        });
    });

    describe("Developer Fee", function () {
        it("should get the correct developer fee", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            // Set new fee basis points
            const newFeeBp = 1337;
            await cmAccountManager.connect(signers.feeAdmin).setDeveloperFeeBp(newFeeBp);

            // Get fee basis points from manager
            const managerFeeBp = await cmAccountManager.getDeveloperFeeBp();
            expect(managerFeeBp).to.equal(newFeeBp);

            // Get fee basis points from account, should be same as manager fee basis points
            expect(await cmAccount.getDeveloperFeeBp()).to.equal(managerFeeBp);
        });
    });
});
