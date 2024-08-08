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

        it("should return all registered services correctly", async function () {
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

            const serviceName1 = "cmp.service.accommodation.v1.AccommodationSearchService";
            const serviceHash1 = ethers.keccak256(ethers.toUtf8Bytes(serviceName1));

            const serviceName2 = "cmp.service.accommodation.v2.AccommodationSearchService";
            const serviceHash2 = ethers.keccak256(ethers.toUtf8Bytes(serviceName2));

            const serviceName3 = "cmp.service.accommodation.v3.AccommodationSearchService";
            const serviceHash3 = ethers.keccak256(ethers.toUtf8Bytes(serviceName3));

            await expect(cmAccountManager.connect(signers.otherAccount1).registerService(serviceName1))
                .to.emit(cmAccountManager, "ServiceRegistered")
                .withArgs(serviceName1, serviceHash1);

            await expect(cmAccountManager.connect(signers.otherAccount1).registerService(serviceName2))
                .to.emit(cmAccountManager, "ServiceRegistered")
                .withArgs(serviceName2, serviceHash2);

            await expect(cmAccountManager.connect(signers.otherAccount1).registerService(serviceName3))
                .to.emit(cmAccountManager, "ServiceRegistered")
                .withArgs(serviceName3, serviceHash3);

            // Check all registered service names
            const registeredServices = await cmAccountManager.getRegisteredServices();
            expect(registeredServices).to.be.deep.equal([serviceName1, serviceName2, serviceName3]);

            // Check all registered service hashes
            const registeredServiceHashes = await cmAccountManager.getRegisteredServiceHashes();
            expect(registeredServiceHashes).to.be.deep.equal([serviceHash1, serviceHash2, serviceHash3]);
        });
    });
});
