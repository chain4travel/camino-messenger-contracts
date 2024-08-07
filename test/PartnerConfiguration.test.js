/**
 * @dev PartnerConfiguration tests
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

describe("PartnerConfiguration", function () {
    describe("Main", function () {
        it("should add a supported service correctly", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const SERVICE_REGISTRY_ADMIN_ROLE = await cmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();

            // Grant SERVICE_REGISTRY_ADMIN_ROLE
            await expect(
                cmAccountManager
                    .connect(signers.managerAdmin)
                    .grantRole(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(cmAccountManager, "RoleGranted")
                .withArgs(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address, signers.managerAdmin.address);

            const serviceName = "cmp.service.accommodation.v1alpha.AccommodationSearchService";
            const serviceHash = ethers.keccak256(ethers.toUtf8Bytes(serviceName));

            await expect(cmAccountManager.connect(signers.otherAccount1).registerService(serviceName))
                .to.emit(cmAccountManager, "ServiceRegistered")
                .withArgs(serviceName, serviceHash);

            // get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await cmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE
            await expect(
                cmAccount.connect(signers.cmAccountAdmin).grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(cmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.cmAccountAdmin.address);

            const fee = 1000n;
            const capabilities = [];

            await expect(cmAccount.connect(signers.otherAccount1).addService(serviceName, fee, capabilities))
                .to.emit(cmAccount, "ServiceAdded")
                .withArgs(serviceHash);
        });
    });
});
