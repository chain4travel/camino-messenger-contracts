const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");

const { ethers } = require("hardhat");

const helpers = require("@nomicfoundation/hardhat-network-helpers");

// Fixtures
const {
    setupSigners,
    developerFeeBp,
    deployCMAccountManagerFixture,
    deployCMAccountImplFixture,
    deployCMAccountManagerWithCMAccountImplFixture,
    deployAndConfigureAllFixture,
    deployCMAccountWithDepositFixture,
    deployBookingTokenFixture,
    deployBookingTokenWithNullUSDFixture,
    deployCancellationSupportFixture,
    deployServiceFeeTokenFixture,
} = require("./utils/fixtures");

describe("ServiceFeeToken", function () {
    it("Should deploy", async function () {
        const { serviceFeeToken } = await loadFixture(deployServiceFeeTokenFixture);

        expect(await serviceFeeToken.name()).to.equal("USD Service Fee Token");
        expect(await serviceFeeToken.symbol()).to.equal("USD.test");
    });

    it("Should pause/unpause", async function () {
        const { serviceFeeToken } = await loadFixture(deployServiceFeeTokenFixture);

        expect(await serviceFeeToken.paused()).to.be.false;

        await serviceFeeToken.connect(signers.managerPauser).pause();
        expect(await serviceFeeToken.paused()).to.be.true;

        // Try to mint
        const amount = ethers.parseEther("100");
        await expect(
            serviceFeeToken.connect(signers.withdrawer).mint(signers.otherAccount1.address, amount),
        ).to.be.revertedWithCustomError(serviceFeeToken, "EnforcedPause");

        // Try to transfer
        await expect(
            serviceFeeToken.connect(signers.otherAccount1).transfer(signers.otherAccount2.address, amount),
        ).to.be.revertedWithCustomError(serviceFeeToken, "EnforcedPause");

        // Try to burn
        await expect(serviceFeeToken.connect(signers.otherAccount1).burn(amount)).to.be.revertedWithCustomError(
            serviceFeeToken,
            "EnforcedPause",
        );

        // unpause
        await serviceFeeToken.connect(signers.managerPauser).unpause();
        expect(await serviceFeeToken.paused()).to.be.false;

        // Try to mint
        await expect(serviceFeeToken.connect(signers.withdrawer).mint(signers.otherAccount1.address, amount)).to.be.not
            .reverted;
    });

    it("Should mint", async function () {
        const { serviceFeeToken } = await loadFixture(deployServiceFeeTokenFixture);

        const amount = ethers.parseEther("100");

        await expect(serviceFeeToken.connect(signers.withdrawer).mint(signers.otherAccount1.address, amount)).to.be.not
            .reverted;

        expect(await serviceFeeToken.balanceOf(signers.otherAccount1.address)).to.equal(amount);
    });

    it("Should burn", async function () {
        const { serviceFeeToken } = await loadFixture(deployServiceFeeTokenFixture);

        const amount = ethers.parseEther("100");

        await serviceFeeToken.connect(signers.withdrawer).mint(signers.otherAccount1.address, amount);
        expect(await serviceFeeToken.balanceOf(signers.otherAccount1.address)).to.equal(amount);

        const burnAmount = ethers.parseEther("30");

        await expect(serviceFeeToken.connect(signers.otherAccount1).burn(burnAmount)).to.be.not.reverted;

        expect(await serviceFeeToken.balanceOf(signers.otherAccount1.address)).to.equal(amount - burnAmount);
    });

    it("Should transfer", async function () {
        const { serviceFeeToken } = await loadFixture(deployServiceFeeTokenFixture);

        const amount = ethers.parseEther("100");

        await serviceFeeToken.connect(signers.withdrawer).mint(signers.otherAccount1.address, amount);
        expect(await serviceFeeToken.balanceOf(signers.otherAccount1.address)).to.equal(amount);

        await expect(serviceFeeToken.connect(signers.otherAccount1).transfer(signers.otherAccount2.address, amount)).to
            .be.not.reverted;

        expect(await serviceFeeToken.balanceOf(signers.otherAccount1.address)).to.equal(0);
        expect(await serviceFeeToken.balanceOf(signers.otherAccount2.address)).to.equal(amount);
    });

    it("Should not allow unauthorized calls", async function () {
        const { serviceFeeToken } = await loadFixture(deployServiceFeeTokenFixture);

        const amount = ethers.parseEther("100");

        // mint
        await expect(
            serviceFeeToken.connect(signers.otherAccount1).mint(signers.otherAccount2.address, amount),
        ).to.be.revertedWithCustomError(serviceFeeToken, "AccessControlUnauthorizedAccount");

        // pause
        await expect(serviceFeeToken.connect(signers.otherAccount1).pause()).to.be.revertedWithCustomError(
            serviceFeeToken,
            "AccessControlUnauthorizedAccount",
        );

        // unpause
        await expect(serviceFeeToken.connect(signers.otherAccount1).unpause()).to.be.revertedWithCustomError(
            serviceFeeToken,
            "AccessControlUnauthorizedAccount",
        );

        // upgrade to and call
        await expect(
            serviceFeeToken.connect(signers.otherAccount1).upgradeToAndCall(signers.otherAccount3.address, "0x"),
        ).to.be.revertedWithCustomError(serviceFeeToken, "AccessControlUnauthorizedAccount");
    });

    it("Should reinitialize correctly", async function () {
        const { serviceFeeToken } = await loadFixture(deployServiceFeeTokenFixture);

        expect(await serviceFeeToken.name()).to.be.equal("USD Service Fee Token");
        expect(await serviceFeeToken.symbol()).to.be.equal("USD.test");

        const newName = "USD Test Token";
        const newSymbol = "USD.test.new";

        // Try to re-init with unauthorized caller
        await expect(
            serviceFeeToken.connect(signers.otherAccount1).reinitializeV2(newName, newSymbol),
        ).to.be.revertedWithCustomError(serviceFeeToken, "AccessControlUnauthorizedAccount");

        // Reinitialize
        await expect(serviceFeeToken.connect(signers.feeAdmin).reinitializeV2(newName, newSymbol)).to.not.reverted;

        // Check new name and symbol
        expect(await serviceFeeToken.name()).to.be.equal(newName);
        expect(await serviceFeeToken.symbol()).to.be.equal(newSymbol);
    });
});
