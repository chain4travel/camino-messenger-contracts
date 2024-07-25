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
            await expect(await cmAccountManager.hasRole(DEFAULT_ADMIN_ROLE, signers.managerAdmin.address)).to.be.true;
            await expect(await cmAccountManager.hasRole(PAUSER_ROLE, signers.managerPauser.address)).to.be.true;
            await expect(await cmAccountManager.hasRole(UPGRADER_ROLE, signers.managerUpgrader.address)).to.be.true;
            await expect(await cmAccountManager.hasRole(VERSIONER_ROLE, signers.managerVersioner.address)).to.be.true;

            // Check state
            await expect(await cmAccountManager.getDeveloperWallet()).to.be.equal(signers.developerWallet.address);
            await expect(await cmAccountManager.getDeveloperFeeBp()).to.be.equal(developerFeeBp);
            await expect(await cmAccountManager.paused()).to.be.false;
        });

        it("should set developer wallet and roles correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            oldDeveloperWallet = signers.developerWallet.address;
            newDeveloperWallet = signers.otherAccount1.address;

            await expect(cmAccountManager.connect(signers.developerWalletAdmin).setDeveloperWallet(newDeveloperWallet))
                .to.emit(cmAccountManager, "DeveloperWalletUpdated")
                .withArgs(oldDeveloperWallet, newDeveloperWallet);
        });

        it("should fail to set developer wallet to zero address", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            await expect(cmAccountManager.connect(signers.developerWalletAdmin).setDeveloperWallet(ethers.ZeroAddress))
                .to.be.revertedWithCustomError(cmAccountManager, "InvalidDeveloperWallet")
                .withArgs(ethers.ZeroAddress);
        });

        it("should set developer fee bassis points", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            const oldFeeBp = await cmAccountManager.getDeveloperFeeBp();
            const newFeeBp = 500;

            await expect(cmAccountManager.connect(signers.feeAdmin).setDeveloperFeeBp(newFeeBp))
                .to.emit(cmAccountManager, "DeveloperFeeBpUpdated")
                .withArgs(oldFeeBp, newFeeBp);

            await expect(await cmAccountManager.getDeveloperFeeBp()).to.be.equal(newFeeBp);
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

            const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const cmAccountManagerAddress = await cmAccountManager.getAddress();
            const cmAccountAddress = await cmAccount.getAddress();

            await expect(await cmAccountManager.isCMAccount(cmAccountAddress)).to.be.true;
            await expect(await cmAccountManager.isCMAccount(signers.otherAccount1.address)).to.be.false;
            await expect(await cmAccountManager.isCMAccount(ethers.ZeroAddress)).to.be.false;
            await expect(await cmAccount.getManagerAddress()).to.be.equal(cmAccountManagerAddress);
        });

        it("should fail if admin is zero address", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerWithCMAccountImplFixture);

            await expect(
                cmAccountManager.createCMAccount(
                    ethers.ZeroAddress,
                    signers.cmAccountPauser,
                    signers.cmAccountUpgrader,
                    true,
                ),
            ).to.be.revertedWithCustomError(cmAccountManager, "CMAccountInvalidAdmin");
        });

        it("should fail if the manager is paused", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerWithCMAccountImplFixture);

            await cmAccountManager.connect(signers.managerPauser).pause();
            await expect(
                cmAccountManager.createCMAccount(
                    signers.cmAccountAdmin.address,
                    signers.cmAccountPauser,
                    signers.cmAccountUpgrader,
                    true,
                ),
            ).to.be.revertedWithCustomError(cmAccountManager, "EnforcedPause");
        });
    });
});
