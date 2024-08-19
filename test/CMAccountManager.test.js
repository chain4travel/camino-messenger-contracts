/**
 * @dev CMAccountManager tests
 */
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

// Fixtures
const {
    setupSigners,
    developerFeeBp,
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

        it("should set developer fee bassis points", async function () {
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
        });
    });

    describe("Upgrades", function () {
        it("should upgrade correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            const CMAccountManagerV2 = await ethers.getContractFactory("CMAccountManagerV2", signers.managerUpgrader);
            const cmAccountManagerV2 = await upgrades.upgradeProxy(cmAccountManager, CMAccountManagerV2);

            await expect(await cmAccountManagerV2.getVersion()).to.be.equal("V2");
        });

        it("should not upgrade if the caller does not have the upgrader role", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            const CMAccountManagerV2 = await ethers.getContractFactory("CMAccountManagerV2", signers.managerPauser);

            await expect(upgrades.upgradeProxy(cmAccountManager, CMAccountManagerV2)).to.be.revertedWithCustomError(
                CMAccountManagerV2,
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

        it("should fail if the prefund amount is lower then the minimum", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, prefundAmount } = await loadFixture(deployAndConfigureAllFixture);

            await expect(
                cmAccountManager.createCMAccount(signers.cmAccountAdmin.address, signers.cmAccountUpgrader, {
                    value: prefundAmount - 1n,
                }),
            )
                .to.be.revertedWithCustomError(cmAccountManager, "IncorrectPrefundAmount")
                .withArgs(prefundAmount, prefundAmount - 1n);
        });

        it("should fail if the prefund amount is zero", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, prefundAmount } = await loadFixture(deployAndConfigureAllFixture);

            await expect(cmAccountManager.createCMAccount(signers.cmAccountAdmin.address, signers.cmAccountUpgrader))
                .to.be.revertedWithCustomError(cmAccountManager, "IncorrectPrefundAmount")
                .withArgs(prefundAmount, 0n);
        });

        it("should allow the prefund amount to be higher then the minimum", async function () {
            const { cmAccountManager, prefundAmount } = await loadFixture(deployAndConfigureAllFixture);

            const overPrefund = ethers.parseEther("100");
            const newPrefundAmount = prefundAmount + overPrefund;

            const tx = await cmAccountManager.createCMAccount(
                signers.cmAccountAdmin.address,
                signers.cmAccountUpgrader,
                {
                    value: newPrefundAmount,
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

            const { cmAccountManager, prefundAmount } = await loadFixture(deployAndConfigureAllFixture);

            newPrefundAmount = prefundAmount + ethers.parseEther("100");

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
