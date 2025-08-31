/**
 * @dev CMAccountManager tests
 */
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

const { setCode } = require("@nomicfoundation/hardhat-network-helpers");

// Fixtures
const {
    setupSigners,
    developerFeeBp,
    deployNullUSDFixture,
    deployCMAccountManagerFixture,
    deployCMAccountImplFixture,
    deployCMAccountManagerWithCMAccountImplFixture,
    deployAndConfigureAllFixture,
} = require("./utils/fixtures");

describe("CMAccountManager", function () {
    describe("Main", function () {
        it("should deploy correctly with the right state", async function () {
            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            const DEFAULT_ADMIN_ROLE = await cmAccountManager.DEFAULT_ADMIN_ROLE();
            const PAUSER_ROLE = await cmAccountManager.PAUSER_ROLE();
            const UPGRADER_ROLE = await cmAccountManager.UPGRADER_ROLE();
            const VERSIONER_ROLE = await cmAccountManager.VERSIONER_ROLE();

            // Check roles
            expect(await cmAccountManager.hasRole(DEFAULT_ADMIN_ROLE, signers.managerAdmin.address)).to.be.true;
            expect(await cmAccountManager.hasRole(PAUSER_ROLE, signers.managerPauser.address)).to.be.true;
            expect(await cmAccountManager.hasRole(UPGRADER_ROLE, signers.managerUpgrader.address)).to.be.true;
            expect(await cmAccountManager.hasRole(VERSIONER_ROLE, signers.managerVersioner.address)).to.be.true;

            // Check state
            expect(await cmAccountManager.getDeveloperWallet()).to.be.equal(signers.developerWallet.address);
            expect(await cmAccountManager.getDeveloperFeeBp()).to.be.equal(developerFeeBp);
            expect(await cmAccountManager.paused()).to.be.false;
            expect(await cmAccountManager.getPrefundAmount()).to.be.equal(ethers.parseEther("100"));
        });

        it("should get role counts correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            const DEFAULT_ADMIN_ROLE = await cmAccountManager.DEFAULT_ADMIN_ROLE();
            const PAUSER_ROLE = await cmAccountManager.PAUSER_ROLE();
            const UPGRADER_ROLE = await cmAccountManager.UPGRADER_ROLE();
            const VERSIONER_ROLE = await cmAccountManager.VERSIONER_ROLE();
            const DEVELOPER_WALLET_ADMIN_ROLE = await cmAccountManager.DEVELOPER_WALLET_ADMIN_ROLE();

            expect(await cmAccountManager.getRoleMemberCount(DEFAULT_ADMIN_ROLE)).to.be.equal(1);
            expect(await cmAccountManager.getRoleMemberCount(PAUSER_ROLE)).to.be.equal(1);
            expect(await cmAccountManager.getRoleMemberCount(UPGRADER_ROLE)).to.be.equal(1);
            expect(await cmAccountManager.getRoleMemberCount(VERSIONER_ROLE)).to.be.equal(1);

            // Developer wallet admin role is not granted by default
            expect(await cmAccountManager.getRoleMemberCount(DEVELOPER_WALLET_ADMIN_ROLE)).to.be.equal(0);

            // Grant developer wallet role
            await expect(
                cmAccountManager
                    .connect(signers.managerAdmin)
                    .grantRole(await cmAccountManager.DEVELOPER_WALLET_ADMIN_ROLE(), signers.otherAccount1.address),
            ).to.not.reverted;

            // Developer wallet admin role is granted, count should be 1
            expect(await cmAccountManager.getRoleMemberCount(DEVELOPER_WALLET_ADMIN_ROLE)).to.be.equal(1);
        });

        it("should set developer wallet and roles correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            await expect(
                cmAccountManager
                    .connect(signers.managerAdmin)
                    .grantRole(
                        await cmAccountManager.DEVELOPER_WALLET_ADMIN_ROLE(),
                        signers.developerWalletAdmin.address,
                    ),
            ).to.not.reverted;

            oldDeveloperWallet = signers.developerWallet.address;
            newDeveloperWallet = signers.otherAccount1.address;

            await expect(cmAccountManager.connect(signers.developerWalletAdmin).setDeveloperWallet(newDeveloperWallet))
                .to.emit(cmAccountManager, "DeveloperWalletUpdated")
                .withArgs(oldDeveloperWallet, newDeveloperWallet);

            expect(await cmAccountManager.getDeveloperWallet()).to.be.equal(newDeveloperWallet);

            // Try to set developer wallet with non-auth address
            await expect(
                cmAccountManager.connect(signers.otherAccount1).setDeveloperWallet(ethers.ZeroAddress),
            ).to.be.revertedWithCustomError(cmAccountManager, "AccessControlUnauthorizedAccount");
        });

        it("should fail to set developer wallet to zero address", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            await expect(
                cmAccountManager
                    .connect(signers.managerAdmin)
                    .grantRole(
                        await cmAccountManager.DEVELOPER_WALLET_ADMIN_ROLE(),
                        signers.developerWalletAdmin.address,
                    ),
            ).to.not.reverted;

            await expect(cmAccountManager.connect(signers.developerWalletAdmin).setDeveloperWallet(ethers.ZeroAddress))
                .to.be.revertedWithCustomError(cmAccountManager, "InvalidDeveloperWallet")
                .withArgs(ethers.ZeroAddress);
        });

        it("should set developer fee basis points", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            const oldFeeBp = await cmAccountManager.getDeveloperFeeBp();
            const newFeeBp = 500;

            await expect(
                await cmAccountManager
                    .connect(signers.managerAdmin)
                    .grantRole(await cmAccountManager.FEE_ADMIN_ROLE(), signers.feeAdmin.address),
            ).to.not.reverted;

            await expect(cmAccountManager.connect(signers.feeAdmin).setDeveloperFeeBp(newFeeBp))
                .to.emit(cmAccountManager, "DeveloperFeeBpUpdated")
                .withArgs(oldFeeBp, newFeeBp);

            await expect(await cmAccountManager.getDeveloperFeeBp()).to.be.equal(newFeeBp);

            // Try to set developer fee with non-auth address
            await expect(
                cmAccountManager.connect(signers.otherAccount1).setDeveloperFeeBp(ethers.ZeroAddress),
            ).to.be.revertedWithCustomError(cmAccountManager, "AccessControlUnauthorizedAccount");
        });

        it("should set and get booking token addr correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, prefundAmount } = await loadFixture(deployAndConfigureAllFixture);

            // Get booking token address
            expect(await cmAccountManager.getBookingTokenAddress()).to.be.not.reverted;

            // Try to set booking token addr with non-auth address
            await expect(
                cmAccountManager.connect(signers.otherAccount3).setBookingTokenAddress(ethers.ZeroAddress),
            ).to.be.revertedWithCustomError(cmAccountManager, "AccessControlUnauthorizedAccount");

            // Try to set booking token to invalid addresses
            await expect(cmAccountManager.connect(signers.managerVersioner).setBookingTokenAddress(ethers.ZeroAddress))
                .to.be.revertedWithCustomError(cmAccountManager, "InvalidBookingTokenAddress")
                .withArgs(ethers.ZeroAddress);

            await expect(
                cmAccountManager
                    .connect(signers.managerVersioner)
                    .setBookingTokenAddress(signers.otherAccount1.address),
            )
                .to.be.revertedWithCustomError(cmAccountManager, "InvalidBookingTokenAddress")
                .withArgs(signers.otherAccount1.address);
        });

        it("should set and get correct prefund amount", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, prefundAmount } = await loadFixture(deployAndConfigureAllFixture);

            newPrefundAmount = prefundAmount + ethers.parseEther("100");

            // Grant the role
            const PREFUND_ADMIN_ROLE = await cmAccountManager.PREFUND_ADMIN_ROLE();
            await cmAccountManager
                .connect(signers.managerAdmin)
                .grantRole(PREFUND_ADMIN_ROLE, signers.otherAccount3.address);

            expect(await cmAccountManager.getPrefundAmount()).to.be.equal(prefundAmount);

            expect(await cmAccountManager.connect(signers.otherAccount3).setPrefundAmount(newPrefundAmount)).to.be.not
                .reverted;

            expect(await cmAccountManager.getPrefundAmount()).to.be.equal(newPrefundAmount);

            // Revoke the role
            await cmAccountManager
                .connect(signers.managerAdmin)
                .revokeRole(PREFUND_ADMIN_ROLE, signers.otherAccount3.address);

            // Try to set prefund amount
            await expect(
                cmAccountManager.connect(signers.otherAccount3).setPrefundAmount(newPrefundAmount),
            ).to.be.revertedWithCustomError(cmAccountManager, "AccessControlUnauthorizedAccount");
        });

        it("should set and get correct service fee token address", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerWithCMAccountImplFixture);
            const { nullUSD } = await loadFixture(deployNullUSDFixture);

            // Get service fee token address - should be zero address
            expect(await cmAccountManager.getServiceFeeToken()).to.be.equal(ethers.ZeroAddress);

            // Try to set service fee token addr with non-auth address
            await expect(
                cmAccountManager.connect(signers.otherAccount3).setServiceFeeToken(ethers.ZeroAddress),
            ).to.be.revertedWithCustomError(cmAccountManager, "AccessControlUnauthorizedAccount");

            // Grant the SERVICE_FEE_TOKEN_ADMIN_ROLE role
            const SERVICE_FEE_TOKEN_ADMIN_ROLE = await cmAccountManager.SERVICE_FEE_TOKEN_ADMIN_ROLE();
            await cmAccountManager
                .connect(signers.managerAdmin)
                .grantRole(SERVICE_FEE_TOKEN_ADMIN_ROLE, signers.otherAccount1.address);

            // Try to set service fee token with auth address
            await expect(cmAccountManager.connect(signers.otherAccount1).setServiceFeeToken(await nullUSD.getAddress()))
                .to.emit(cmAccountManager, "ServiceFeeTokenUpdated")
                .withArgs(ethers.ZeroAddress, await nullUSD.getAddress());

            // Get service fee token address
            expect(await cmAccountManager.getServiceFeeToken()).to.be.equal(await nullUSD.getAddress());
        });
    });

    describe("Upgrades", function () {
        it("should upgrade correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            const CMAccountManagerTest = await ethers.getContractFactory(
                "CMAccountManagerTest",
                signers.managerUpgrader,
            );
            const cmAccountManagerTest = await upgrades.upgradeProxy(cmAccountManager, CMAccountManagerTest);

            await expect(await cmAccountManagerTest.getVersion()).to.be.equal("TESTING");
        });

        it("should not upgrade if the caller does not have the upgrader role", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            const CMAccountManagerTest = await ethers.getContractFactory("CMAccountManagerTest", signers.managerPauser);

            await expect(upgrades.upgradeProxy(cmAccountManager, CMAccountManagerTest)).to.be.revertedWithCustomError(
                CMAccountManagerTest,
                "AccessControlUnauthorizedAccount",
            );
        });
    });

    describe("CMAccount Implementation", function () {
        it("should set CMAccount implementation correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);
            const { cmAccountImpl } = await loadFixture(deployCMAccountImplFixture);

            const cmAccountImplAddress = await cmAccountImpl.getAddress();

            await expect(
                await cmAccountManager.connect(signers.managerVersioner).setAccountImplementation(cmAccountImplAddress),
            )
                .to.emit(cmAccountManager, "CMAccountImplementationUpdated")
                .withArgs(ethers.ZeroAddress, cmAccountImplAddress);
        });

        it("should get CMAccount implementation correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, cmAccountImplAddress } = await loadFixture(
                deployCMAccountManagerWithCMAccountImplFixture,
            );

            await expect(await cmAccountManager.getAccountImplementation()).to.be.equal(cmAccountImplAddress);
        });

        it("should revert if the implementation is zero code length address", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, cmAccountImplAddress } = await loadFixture(
                deployCMAccountManagerWithCMAccountImplFixture,
            );

            await expect(
                cmAccountManager.connect(signers.managerVersioner).setAccountImplementation(ethers.ZeroAddress),
            )
                .to.be.revertedWithCustomError(cmAccountManager, "CMAccountInvalidImplementation")
                .withArgs(ethers.ZeroAddress);
        });

        it("should revert if the caller does not have the versioner role", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, cmAccountImplAddress } = await loadFixture(
                deployCMAccountManagerWithCMAccountImplFixture,
            );

            await expect(
                cmAccountManager.connect(signers.otherAccount1).setAccountImplementation(cmAccountImplAddress),
            ).to.be.revertedWithCustomError(cmAccountManager, "AccessControlUnauthorizedAccount");
        });
    });

    describe("Pausable", function () {
        it("should pause and unpause the contract", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            await cmAccountManager.connect(signers.managerPauser).pause();
            await expect(await cmAccountManager.paused()).to.be.true;

            await cmAccountManager.connect(signers.managerPauser).unpause();
            await expect(await cmAccountManager.paused()).to.be.false;
        });

        it("should not allow non-pauser to pause", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            await expect(cmAccountManager.connect(signers.otherAccount1).pause()).to.be.revertedWithCustomError(
                cmAccountManager,
                "AccessControlUnauthorizedAccount",
            );

            await expect(cmAccountManager.connect(signers.otherAccount1).unpause()).to.be.revertedWithCustomError(
                cmAccountManager,
                "AccessControlUnauthorizedAccount",
            );
        });
    });

    describe("CMAccount", function () {
        it("should create CMAccount correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, cmAccount, prefundAmount } = await loadFixture(deployAndConfigureAllFixture);

            const cmAccountManagerAddress = await cmAccountManager.getAddress();
            const cmAccountAddress = await cmAccount.getAddress();

            expect(await cmAccountManager.isCMAccount(cmAccountAddress)).to.be.true;
            expect(await cmAccountManager.isCMAccount(signers.otherAccount1.address)).to.be.false;
            expect(await cmAccountManager.isCMAccount(ethers.ZeroAddress)).to.be.false;
            expect(await cmAccount.getManagerAddress()).to.be.equal(cmAccountManagerAddress);

            // Check balance for prefund
            expect(await ethers.provider.getBalance(cmAccountAddress)).to.be.equal(prefundAmount);
        });

        it("should revert creating for zero code btoken and impl addr", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, bookingToken, prefundAmount } = await loadFixture(deployAndConfigureAllFixture);

            // Get account impl
            const CMAccountImplAddr = await cmAccountManager.getAccountImplementation();

            // Set booking token code to zero
            await network.provider.send("hardhat_setCode", [await bookingToken.getAddress(), "0x"]);

            // Create CMAccount
            await expect(
                cmAccountManager.createCMAccount(signers.cmAccountAdmin.address, signers.cmAccountUpgrader.address, {
                    value: prefundAmount,
                }),
            ).to.be.revertedWithCustomError(cmAccountManager, "InvalidBookingTokenAddress");

            // Set acct impl code to zero
            await network.provider.send("hardhat_setCode", [CMAccountImplAddr, "0x"]);

            // Create CMAccount
            await expect(
                cmAccountManager.createCMAccount(signers.cmAccountAdmin.address, signers.cmAccountUpgrader.address, {
                    value: prefundAmount,
                }),
            ).to.be.revertedWithCustomError(cmAccountManager, "CMAccountInvalidImplementation");
        });

        it("should fail if admin is zero address", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, prefundAmount } = await loadFixture(deployAndConfigureAllFixture);

            await expect(
                cmAccountManager.createCMAccount(ethers.ZeroAddress, signers.cmAccountUpgrader, {
                    value: prefundAmount,
                }),
            ).to.be.revertedWithCustomError(cmAccountManager, "CMAccountInvalidAdmin");
        });

        it("should fail if the manager is paused", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, prefundAmount } = await loadFixture(deployAndConfigureAllFixture);

            await cmAccountManager.connect(signers.managerPauser).pause();
            await expect(
                cmAccountManager.createCMAccount(signers.cmAccountAdmin.address, signers.cmAccountUpgrader, {
                    value: prefundAmount,
                }),
            ).to.be.revertedWithCustomError(cmAccountManager, "EnforcedPause");
        });

        it("should not fail for any msg.value", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, nullUSD, prefundAmount } = await loadFixture(deployAndConfigureAllFixture);

            // Approve service fee
            await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);

            // Zero msg.value
            await expect(
                cmAccountManager.createCMAccount(signers.cmAccountAdmin.address, signers.cmAccountUpgrader, {
                    value: 0n,
                }),
            ).to.be.not.reverted;

            // Approve service fee
            await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);

            const nonZeroMsgValue = ethers.parseEther("300");

            // Non-zero msg.value
            const nonZeroMsgValueTx = await cmAccountManager.createCMAccount(
                signers.cmAccountAdmin.address,
                signers.cmAccountUpgrader,
                {
                    value: nonZeroMsgValue,
                },
            );

            await expect(nonZeroMsgValueTx).to.be.not.reverted;

            const receipt = await nonZeroMsgValueTx.wait();

            // Parse event to get the CMAccount address (this is the UUPS proxy address)
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = cmAccountManager.interface.parseLog(event);
            const cmAccountAddress = parsedEvent.args.account;

            // Check balances
            await expect(nonZeroMsgValueTx).to.changeEtherBalances(
                [signers.managerAdmin, cmAccountAddress],
                [-nonZeroMsgValue, nonZeroMsgValue],
            );
            await expect(nonZeroMsgValueTx).to.changeTokenBalances(
                nullUSD,
                [signers.managerAdmin, cmAccountAddress],
                [-prefundAmount, prefundAmount],
            );
        });

        it("should fail if service fee token address is zero", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerWithCMAccountImplFixture);

            // Set dummy booking token address
            await cmAccountManager
                .connect(signers.managerVersioner)
                .setBookingTokenAddress(await cmAccountManager.getAddress());

            // Check if service fee token address is zero
            const serviceFeeTokenAddress = await cmAccountManager.getServiceFeeToken();
            expect(serviceFeeTokenAddress).to.equal(ethers.ZeroAddress);

            // Try CM Account creation
            await expect(cmAccountManager.createCMAccount(signers.cmAccountAdmin.address, signers.cmAccountUpgrader))
                .to.be.revertedWithCustomError(cmAccountManager, "InvalidServiceFeeToken")
                .withArgs(ethers.ZeroAddress);

            // Try setting service fee token ----------------------------------------------------------------------

            // Set required role
            const SERVICE_FEE_TOKEN_ADMIN_ROLE = await cmAccountManager.SERVICE_FEE_TOKEN_ADMIN_ROLE();
            await cmAccountManager.grantRole(SERVICE_FEE_TOKEN_ADMIN_ROLE, signers.feeAdmin.address);

            // Try setting service fee token
            await expect(cmAccountManager.connect(signers.feeAdmin).setServiceFeeToken(ethers.ZeroAddress))
                .to.be.revertedWithCustomError(cmAccountManager, "InvalidServiceFeeToken")
                .withArgs(ethers.ZeroAddress);

            // Try with non-zero non-contract address
            await expect(
                cmAccountManager.connect(signers.feeAdmin).setServiceFeeToken(signers.otherAccount1.address),
            ).to.be.revertedWithCustomError(cmAccountManager, "InvalidServiceFeeToken");
        });

        it("should fail if the prefund amount is not approved", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, nullUSD, prefundAmount } = await loadFixture(deployAndConfigureAllFixture);

            await expect(cmAccountManager.createCMAccount(signers.cmAccountAdmin.address, signers.cmAccountUpgrader))
                .to.be.revertedWithCustomError(nullUSD, "ERC20InsufficientAllowance")
                .withArgs(await cmAccountManager.getAddress(), 0n, prefundAmount);
        });

        it("should allow the prefund amount to be higher then the minimum", async function () {
            const { cmAccountManager, prefundAmount, nullUSD, nullUSDDecimals } =
                await loadFixture(deployAndConfigureAllFixture);

            const overPrefund = ethers.parseEther("100");
            const newPrefundAmount = prefundAmount + overPrefund;

            // Approve service fee
            await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);

            const tx = await cmAccountManager.createCMAccount(
                signers.cmAccountAdmin.address,
                signers.cmAccountUpgrader,
                {
                    value: newPrefundAmount, // This is not required anymore
                },
            );

            const receipt = await tx.wait();

            // Parse event to get the CMAccount address (this is the UUPS proxy address)
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = cmAccountManager.interface.parseLog(event);
            const cmAccountAddress = parsedEvent.args.account;

            expect(await ethers.provider.getBalance(cmAccountAddress)).to.be.equal(newPrefundAmount);
        });

        it("should set and get correct account creator", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, prefundAmount, nullUSD } = await loadFixture(deployAndConfigureAllFixture);

            newPrefundAmount = prefundAmount + ethers.parseEther("100");

            // Approve service fee
            await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);

            // Create distributor CMAccount
            // This is called with managerAdmin as the signer
            const tx = await cmAccountManager.createCMAccount(
                signers.cmAccountAdmin.address,
                signers.cmAccountUpgrader.address,
                { value: prefundAmount },
            );

            const receipt = await tx.wait();

            // Parse event to get the CMAccount address (this is the UUPS proxy address)
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = cmAccountManager.interface.parseLog(event);
            const newCMAccountAddress = parsedEvent.args.account;

            expect(await cmAccountManager.getCMAccountCreator(newCMAccountAddress)).to.be.equal(
                signers.managerAdmin.address,
            );
        });
    });
});
