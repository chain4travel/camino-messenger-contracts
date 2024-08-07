/**
 * @dev ServiceRegistry tests
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

describe("ServiceRegistry", function () {
    describe("Main", function () {
        it("should register a service correctly", async function () {
            const { cmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

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

            await expect(await cmAccountManager.getRegisteredServiceHashByName(serviceName)).to.be.equal(serviceHash);
            await expect(await cmAccountManager.getRegisteredServiceNameByHash(serviceHash)).to.be.equal(serviceName);
        });

        it("should unregister a service correctly", async function () {
            const { cmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

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

            await expect(cmAccountManager.connect(signers.otherAccount1).unregisterService(serviceName))
                .to.emit(cmAccountManager, "ServiceUnregistered")
                .withArgs(serviceName, serviceHash);
        });

        it("should revert if the service is already registered", async function () {
            const { cmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

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

            // Register the service
            await expect(cmAccountManager.connect(signers.otherAccount1).registerService(serviceName)).to.be.not
                .reverted;

            // Register the service again
            await expect(cmAccountManager.connect(signers.otherAccount1).registerService(serviceName))
                .to.be.revertedWithCustomError(cmAccountManager, "ServiceAlreadyRegistered")
                .withArgs(serviceName);
        });

        it("should revert if the caller does not have the SERVICE_REGISTRY_ADMIN_ROLE", async function () {
            const { cmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            const SERVICE_REGISTRY_ADMIN_ROLE = await cmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();

            const serviceName = "cmp.service.accommodation.v1alpha.AccommodationSearchService";

            await expect(cmAccountManager.connect(signers.otherAccount1).registerService(serviceName))
                .to.be.revertedWithCustomError(cmAccountManager, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount1.address, SERVICE_REGISTRY_ADMIN_ROLE);
        });
    });
});
