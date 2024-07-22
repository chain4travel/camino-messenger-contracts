// test/CMAccountManager.js
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

// @dev TODO: Extend and tidy up tests
// @dev TODO: Add ChequeManager tests
// @dev TODO: Use fixtures for CM account creations in various tests

describe("CMAccountManager", function () {
    async function deployCMAccountManagerFixture() {
        const [defaultAdmin, pauser, upgrader, versioner, developerWallet, initialOwner, otherAccount] =
            await ethers.getSigners();

        const CMAccountManager = await ethers.getContractFactory("CMAccountManager");
        const cmAccountManager = await upgrades.deployProxy(
            CMAccountManager,
            [defaultAdmin.address, pauser.address, upgrader.address, versioner.address, developerWallet.address, 100],
            { kind: "uups" },
        );

        const CMAccount = await ethers.getContractFactory("CMAccount");
        const cmAccount = await CMAccount.deploy();
        await cmAccount.waitForDeployment();

        cmAccountAddress = await cmAccount.getAddress();

        await cmAccountManager.grantRole(await cmAccountManager.VERSIONER_ROLE(), versioner.address);
        await cmAccountManager.connect(versioner).setAccountImplementation(cmAccountAddress);

        return {
            cmAccountManager,
            cmAccount,
            defaultAdmin,
            pauser,
            upgrader,
            versioner,
            initialOwner,
            otherAccount,
        };
    }

    describe("Deployment", function () {
        it("should deploy correctly with the right roles", async function () {
            const { cmAccountManager, defaultAdmin, pauser, upgrader } =
                await loadFixture(deployCMAccountManagerFixture);

            expect(await cmAccountManager.hasRole(await cmAccountManager.DEFAULT_ADMIN_ROLE(), defaultAdmin.address)).to
                .be.true;
            expect(await cmAccountManager.hasRole(await cmAccountManager.PAUSER_ROLE(), pauser.address)).to.be.true;
            expect(await cmAccountManager.hasRole(await cmAccountManager.UPGRADER_ROLE(), upgrader.address)).to.be.true;
        });
    });

    describe("CMAccount Implementation", function () {
        it("should revert if the account implementation is invalid", async function () {
            const { cmAccountManager, versioner, initialOwner, pauser, upgrader } =
                await loadFixture(deployCMAccountManagerFixture);

            await expect(
                cmAccountManager.connect(versioner).setAccountImplementation(ethers.ZeroAddress),
            ).to.be.revertedWithCustomError(cmAccountManager, "CMAccountInvalidImplementation");
        });

        it("should set the account implementation correctly", async function () {
            const { cmAccountManager, defaultAdmin, versioner, initialOwner, pauser, upgrader } =
                await loadFixture(deployCMAccountManagerFixture);

            // Create a new implementation for CMAccount
            const CMAccountImplV2 = await ethers.getContractFactory("CMAccount");
            const cmAccountImplV2 = await CMAccountImplV2.deploy();
            await cmAccountImplV2.waitForDeployment();
            const newImplementation = await cmAccountImplV2.getAddress();

            // Get old implementation address from the factory
            const oldImplementation = await cmAccountManager.getAccountImplementation();

            // Set implementation address in the factory
            await expect(cmAccountManager.connect(versioner).setAccountImplementation(newImplementation))
                .to.emit(cmAccountManager, "CMAccountImplementationUpdated")
                .withArgs(oldImplementation, newImplementation);
        });
    });

    describe("CMAccount Creation", function () {
        it("should create a new CMAccount", async function () {
            const { cmAccountManager, cmAccount, initialOwner, pauser, upgrader, otherAccount } =
                await loadFixture(deployCMAccountManagerFixture);

            const anyOneCanDeposit = true;

            await expect(
                await cmAccountManager.createCMAccount(
                    initialOwner.address,
                    pauser.address,
                    upgrader.address,
                    anyOneCanDeposit,
                ),
            ).to.emit(cmAccountManager, "CMAccountCreated");

            // Create CMAccount
            const tx = await cmAccountManager.createCMAccount(
                initialOwner.address,
                pauser.address,
                upgrader.address,
                anyOneCanDeposit,
            );
            const receipt = await tx.wait();

            // Decode the event log
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                    const anyOneCanDeposit = true;
                }
            });

            const parsedEvent = cmAccountManager.interface.parseLog(event);
            expect(parsedEvent.name).to.be.equal("CMAccountCreated");

            const accountAddress = parsedEvent.args.account;

            expect(await cmAccountManager.isCMAccount(accountAddress)).to.be.true;
            expect(await cmAccountManager.isCMAccount(otherAccount)).to.be.false;
        });

        it("should revert if the owner address is invalid", async function () {
            const { cmAccountManager, cmAccount, pauser, upgrader } = await loadFixture(deployCMAccountManagerFixture);

            const anyOneCanDeposit = true;

            await expect(
                cmAccountManager.createCMAccount(
                    ethers.ZeroAddress,
                    pauser.address,
                    upgrader.address,
                    anyOneCanDeposit,
                ),
            ).to.be.revertedWithCustomError(cmAccountManager, "CMAccountInvalidOwner");
        });
    });

    describe("Factory Pausable", function () {
        it("should pause and unpause the contract", async function () {
            const { cmAccountManager, pauser } = await loadFixture(deployCMAccountManagerFixture);

            await cmAccountManager.connect(pauser).pause();
            expect(await cmAccountManager.paused()).to.be.true;

            await cmAccountManager.connect(pauser).unpause();
            expect(await cmAccountManager.paused()).to.be.false;
        });

        it("should not allow non-pauser to pause", async function () {
            const { cmAccountManager, otherAccount } = await loadFixture(deployCMAccountManagerFixture);

            await expect(cmAccountManager.connect(otherAccount).pause()).to.be.reverted;
        });

        it("should not allow to create CMAccount when paused", async function () {
            const { cmAccountManager, pauser, otherAccount } = await loadFixture(deployCMAccountManagerFixture);

            // Pause the factory
            await cmAccountManager.connect(pauser).pause();
            expect(await cmAccountManager.paused()).to.be.true;

            const anyOneCanDeposit = true;

            // Try to create CMAccount
            await expect(
                cmAccountManager
                    .connect(pauser)
                    .createCMAccount(
                        otherAccount.address,
                        otherAccount.address,
                        otherAccount.address,
                        anyOneCanDeposit,
                    ),
            ).to.be.revertedWithCustomError(cmAccountManager, "EnforcedPause");
        });
    });

    describe("CM Account Upgrade", function () {
        it("should upgrade if implementation address matches the factory", async function () {
            const { cmAccountManager, defaultAdmin, versioner, initialOwner, pauser, upgrader } =
                await loadFixture(deployCMAccountManagerFixture);

            // Get old implementation address from the factory
            const oldImplementation = await cmAccountManager.getAccountImplementation();

            const anyOneCanDeposit = true;

            // Create CMAccount with old implementation
            const tx = await cmAccountManager.createCMAccount(
                initialOwner.address,
                pauser.address,
                upgrader.address,
                anyOneCanDeposit,
            );
            const receipt = await tx.wait();

            // Decode the event log
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = cmAccountManager.interface.parseLog(event);
            const accountAddress = parsedEvent.args.account;

            expect(accountAddress).to.not.be.null;
            expect(accountAddress).to.not.equal(ethers.ZeroAddress);

            // Create a new implementation for CMAccount
            const CMAccountImplV2 = await ethers.getContractFactory("CMAccount");
            const cmAccountImplV2 = await CMAccountImplV2.deploy();
            await cmAccountImplV2.waitForDeployment();
            const cmAccountImplV2Address = await cmAccountImplV2.getAddress();

            // Set new implementation address in the factory
            await cmAccountManager.connect(versioner).setAccountImplementation(cmAccountImplV2Address);

            // Get new implementation address from the factory
            const newImplementation = await cmAccountManager.getAccountImplementation();
            expect(newImplementation).to.equal(cmAccountImplV2Address);

            // Get the contract at the account proxy address with the CMAccount interface
            cmAccountProxy = await ethers.getContractAt("CMAccount", accountAddress);

            // Upgrade the account
            await expect(cmAccountProxy.connect(upgrader).upgradeToAndCall(newImplementation, "0x"))
                .to.emit(cmAccountProxy, "CMAccountUpgraded")
                .withArgs(oldImplementation, newImplementation);
        });

        it("should revert if implementation address does not match the factory", async function () {
            const { cmAccountManager, cmAccount, initialOwner, pauser, upgrader, otherAccount } =
                await loadFixture(deployCMAccountManagerFixture);

            const anyOneCanDeposit = true;

            // Create CMAccount
            const tx = await cmAccountManager.createCMAccount(
                initialOwner.address,
                pauser.address,
                upgrader.address,
                anyOneCanDeposit,
            );
            const receipt = await tx.wait();

            // Decode the event log
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = cmAccountManager.interface.parseLog(event);
            const accountAddress = parsedEvent.args.account;

            //console.log(parsedEvent.args.account);

            expect(accountAddress).to.not.be.null;
            expect(accountAddress).to.not.equal(ethers.ZeroAddress);

            // Create a new implementation for CMAccount
            const CMAccountImplV2 = await ethers.getContractFactory("CMAccount");
            const cmAccountImplV2 = await CMAccountImplV2.deploy();
            await cmAccountImplV2.waitForDeployment();
            const cmAccountImplV2Address = await cmAccountImplV2.getAddress();

            // Factory CMAccount implementation
            const factoryImplementation = await cmAccountManager.getAccountImplementation();

            cmAccountProxy = await ethers.getContractAt("CMAccount", accountAddress);

            await expect(cmAccountProxy.connect(upgrader).upgradeToAndCall(cmAccountImplV2Address, "0x"))
                .to.revertedWithCustomError(cmAccountProxy, "CMAccountImplementationMismatch")
                .withArgs(factoryImplementation, cmAccountImplV2Address);
        });

        it("should revert if the new implementation is same as the current implementation", async function () {
            const { cmAccountManager, cmAccount, initialOwner, pauser, upgrader } =
                await loadFixture(deployCMAccountManagerFixture);

            const anyOneCanDeposit = true;

            // Create CMAccount
            const tx = await cmAccountManager.createCMAccount(
                initialOwner.address,
                pauser.address,
                upgrader.address,
                anyOneCanDeposit,
            );
            const receipt = await tx.wait();

            // Decode the event log
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = cmAccountManager.interface.parseLog(event);
            const accountAddress = parsedEvent.args.account;

            expect(accountAddress).to.not.be.null;
            expect(accountAddress).to.not.equal(ethers.ZeroAddress);

            // Factory CMAccount implementation
            const factoryImplementation = await cmAccountManager.getAccountImplementation();

            await expect(cmAccountProxy.connect(upgrader).upgradeToAndCall(factoryImplementation, "0x"))
                .to.revertedWithCustomError(cmAccountProxy, "CMAccountNoUpdateNeeded")
                .withArgs(factoryImplementation, factoryImplementation);
        });
    });
});
