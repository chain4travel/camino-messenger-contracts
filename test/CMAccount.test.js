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
    deployBookingTokenWithNullUSDFixture,
} = require("./utils/fixtures");

describe("CMAccount", function () {
    describe("Upgrade", function () {
        it("should upgrade to new implementation address", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            // Old implementation
            const oldImplementationAddress = await cmAccountManager.getAccountImplementation();

            // Create a new implementation for CMAccount
            const BookingTokenOperator = await ethers.getContractFactory("BookingTokenOperator");
            const bookingTokenOperator = await BookingTokenOperator.deploy();
            const CMAccountImplV2 = await ethers.getContractFactory("CMAccount", {
                libraries: { BookingTokenOperator: await bookingTokenOperator.getAddress() },
            });
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
            const BookingTokenOperator = await ethers.getContractFactory("BookingTokenOperator");
            const bookingTokenOperator = await BookingTokenOperator.deploy();
            const CMAccountImplV2 = await ethers.getContractFactory("CMAccount", {
                libraries: { BookingTokenOperator: await bookingTokenOperator.getAddress() },
            });
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

        it("should revert if prefund not spent yet", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            const withdrawer = signers.withdrawer;
            // Try to withdraw more than prefundLeft. cmAccount has 100 CAM prefund and 1 CAM deposit initially
            const withdrawAmount = ethers.parseEther("50");

            // Try withdraw
            // PrefundNotSpentYet(withdrawableAmount, prefundLeft, amount);
            const withdrawTx = cmAccount.connect(withdrawer).withdraw(withdrawer.address, withdrawAmount);
            await expect(withdrawTx)
                .to.be.revertedWithCustomError(cmAccount, "PrefundNotSpentYet")
                .withArgs(ethers.parseEther("1"), ethers.parseEther("100"), ethers.parseEther("50"));
        });

        it("should withdraw amount if it's not violating the prefund spent", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            const withdrawer = signers.withdrawer;
            // Try to withdraw allowed amount. cmAccount has 100 CAM prefund and 1
            // CAM deposit initially. So max 1 cam is withdrawable.
            const withdrawAmount = ethers.parseEther("1");

            // Try withdraw
            const withdrawTx = cmAccount.connect(withdrawer).withdraw(withdrawer.address, withdrawAmount);
            await expect(withdrawTx).to.be.not.reverted;

            // Check balances
            await expect(withdrawTx).to.changeEtherBalances([cmAccount, withdrawer], [-withdrawAmount, withdrawAmount]);
        });
    });

    describe("Enumerable", function () {
        it("should get role counts correctly", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const DEFAULT_ADMIN_ROLE = await cmAccount.DEFAULT_ADMIN_ROLE();
            const UPGRADER_ROLE = await cmAccount.UPGRADER_ROLE();
            const BOOKING_OPERATOR_ROLE = await cmAccount.BOOKING_OPERATOR_ROLE();

            expect(await cmAccount.getRoleMemberCount(DEFAULT_ADMIN_ROLE)).to.be.equal(1);
            expect(await cmAccount.getRoleMemberCount(UPGRADER_ROLE)).to.be.equal(1);

            // Booking operator role is not granted by default
            expect(await cmAccount.getRoleMemberCount(BOOKING_OPERATOR_ROLE)).to.be.equal(0);

            // Grant booking operator role
            await expect(
                cmAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(await cmAccount.BOOKING_OPERATOR_ROLE(), signers.otherAccount1.address),
            ).to.not.reverted;

            // Booking operator role is granted, count should be 1
            expect(await cmAccount.getRoleMemberCount(BOOKING_OPERATOR_ROLE)).to.be.equal(1);
        });
    });

    describe("Messenger Bot", function () {
        it("should add messenger bot correctly", async function () {
            const { cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const bot = signers.otherAccount1;

            // Register bot
            await expect(cmAccount.connect(signers.cmAccountAdmin).addMessengerBot(bot.address, 0n))
                .to.emit(cmAccount, "MessengerBotAdded")
                .withArgs(bot.address);

            // Check if bot is allowed
            expect(await cmAccount.isBotAllowed(bot.address)).to.be.true;

            // Check roles
            expect(await cmAccount.hasRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), bot.address)).to.be.true;
            expect(await cmAccount.hasRole(await cmAccount.BOOKING_OPERATOR_ROLE(), bot.address)).to.be.true;
            expect(await cmAccount.hasRole(await cmAccount.GAS_WITHDRAWER_ROLE(), bot.address)).to.be.true;
        });

        it("should remove messenger bot correctly", async function () {
            const { cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const bot = signers.otherAccount1;

            // Register bot
            await expect(cmAccount.connect(signers.cmAccountAdmin).addMessengerBot(bot.address, 0n))
                .to.emit(cmAccount, "MessengerBotAdded")
                .withArgs(bot.address);

            // Remove bot
            await expect(cmAccount.connect(signers.cmAccountAdmin).removeMessengerBot(bot.address))
                .to.emit(cmAccount, "MessengerBotRemoved")
                .withArgs(bot.address);
        });

        it("should add messenger bot with gas money withdrawal correctly", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            const bot = signers.otherAccount1;

            const withdrawAmount = ethers.parseEther("0.1"); // Small amount, fixture has +1 CAM surplus over the prefund

            // Register bot
            const withdrawTx = cmAccount
                .connect(signers.cmAccountAdmin)
                ["addMessengerBot(address,uint256)"](bot.address, withdrawAmount);

            await expect(withdrawTx).to.changeEtherBalances([cmAccount, bot], [-withdrawAmount, withdrawAmount]);
            await expect(withdrawTx).to.emit(cmAccount, "MessengerBotAdded").withArgs(bot.address);
        });

        it("should revert adding messenger bot with gas money withdrawal if prefund not spent", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            const bot = signers.otherAccount1;

            const withdrawAmount = ethers.parseEther("10"); // Fixture has +1 CAM surplus over the prefund, 10 CAM should revert

            // Register bot
            await expect(
                cmAccount
                    .connect(signers.cmAccountAdmin)
                    ["addMessengerBot(address,uint256)"](bot.address, withdrawAmount),
            )
                .to.revertedWithCustomError(cmAccount, "PrefundNotSpentYet")
                .withArgs(ethers.parseEther("1"), ethers.parseEther("100"), ethers.parseEther("10"));
        });
    });

    describe("Transfer ERC20 & ERC721", function () {
        it("should transfer ERC20 correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken, nullUSD } =
                await loadFixture(deployBookingTokenWithNullUSDFixture);

            // Supplier and distributor CM accounts has 10k NullUSD from the fixture
            const amount = ethers.parseEther("100");

            // Transfer
            await expect(
                await supplierCMAccount
                    .connect(signers.withdrawer)
                    .transferERC20(await nullUSD.getAddress(), signers.otherAccount1.address, amount),
            ).to.changeTokenBalances(
                nullUSD,
                [await supplierCMAccount.getAddress(), signers.otherAccount1],
                [-amount, amount],
            );

            // Check balance
            expect(await nullUSD.balanceOf(signers.otherAccount1.address)).to.be.equal(amount);
        });

        it("should transfer ERC721 correctly after it expires", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken, nullUSD } =
                await loadFixture(deployBookingTokenWithNullUSDFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // zero address
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // zero address
                );

            // Advance time to after the expiration
            await network.provider.send("evm_setNextBlockTimestamp", [expirationTimestamp + 1]);
            await network.provider.send("evm_mine");

            // Try to transfer the token with the supplier CMAccount
            await expect(
                supplierCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(await bookingToken.getAddress(), signers.otherAccount1.address, 0n),
            )
                .to.emit(bookingToken, "Transfer")
                .withArgs(await supplierCMAccount.getAddress(), signers.otherAccount1.address, 0n);

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await signers.otherAccount1.getAddress());
        });
    });
});
