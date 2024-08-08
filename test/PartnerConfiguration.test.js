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
    deployAndConfigureAllWithRegisteredServicesFixture,
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

        it("should revert if the caller does not have the SERVICE_ADMIN_ROLE", async function () {
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

            // Grant SERVICE_ADMIN_ROLE to otherAccount1
            await expect(
                cmAccount.connect(signers.cmAccountAdmin).grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(cmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.cmAccountAdmin.address);

            const fee = 1000n;
            const capabilities = [];

            // Try to add a service with otherAccount2
            await expect(cmAccount.connect(signers.otherAccount2).addService(serviceName, fee, capabilities))
                .to.be.revertedWithCustomError(cmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount2.address, SERVICE_ADMIN_ROLE);
        });

        it("should add and return all supported services correctly", async function () {
            const { cmAccountManager, cmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            const fee1 = 1000n;
            const fee2 = 2000n;
            const fee3 = 3000n;
            const capabilities1 = ["test capability 1"];
            const capabilities2 = ["test capability 2"];
            const capabilities3 = ["test capability 3"];

            // Add services to CM account
            expect(
                await cmAccount.connect(signers.cmServiceAdmin).addService(services.serviceName1, fee1, capabilities1),
            )
                .to.emit(cmAccount, "ServiceAdded")
                .withArgs(services.serviceHash1);

            expect(
                await cmAccount.connect(signers.cmServiceAdmin).addService(services.serviceName2, fee2, capabilities2),
            )
                .to.emit(cmAccount, "ServiceAdded")
                .withArgs(services.serviceHash2);

            expect(
                await cmAccount.connect(signers.cmServiceAdmin).addService(services.serviceName3, fee3, capabilities3),
            )
                .to.emit(cmAccount, "ServiceAdded")
                .withArgs(services.serviceHash3);

            // Get all services
            const servicesFromCMAccount = await cmAccount.getSupportedServices();
            expect(servicesFromCMAccount).to.be.deep.equal([
                [services.serviceName1, services.serviceName2, services.serviceName3],
                [
                    [fee1, capabilities1],
                    [fee2, capabilities2],
                    [fee3, capabilities3],
                ],
            ]);
        });
    });
});
