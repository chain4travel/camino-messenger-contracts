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
    describe("Services", function () {
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
            const restrictedRate = false;
            const capabilities = [];

            await expect(
                cmAccount.connect(signers.otherAccount1).addService(serviceName, fee, restrictedRate, capabilities),
            )
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
            const restrictedRate = false;
            const capabilities = [];

            // Try to add a service with otherAccount2
            await expect(
                cmAccount.connect(signers.otherAccount2).addService(serviceName, fee, restrictedRate, capabilities),
            )
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

            const restrictedRate = false;

            const capabilities1 = ["test capability 1"];
            const capabilities2 = ["test capability 2"];
            const capabilities3 = ["test capability 3"];

            // Add services to CM account
            expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    .addService(services.serviceName1, fee1, restrictedRate, capabilities1),
            )
                .to.emit(cmAccount, "ServiceAdded")
                .withArgs(services.serviceHash1);

            expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    .addService(services.serviceName2, fee2, restrictedRate, capabilities2),
            )
                .to.emit(cmAccount, "ServiceAdded")
                .withArgs(services.serviceHash2);

            expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    .addService(services.serviceName3, fee3, restrictedRate, capabilities3),
            )
                .to.emit(cmAccount, "ServiceAdded")
                .withArgs(services.serviceHash3);

            // Get all services
            const servicesFromCMAccount = await cmAccount.getSupportedServices();
            expect(servicesFromCMAccount).to.be.deep.equal([
                [services.serviceName1, services.serviceName2, services.serviceName3],
                [
                    [fee1, restrictedRate, capabilities1],
                    [fee2, restrictedRate, capabilities2],
                    [fee3, restrictedRate, capabilities3],
                ],
            ]);

            // Get specific fee from a service name
            expect(await cmAccount.getServiceFeeByName(services.serviceName1)).to.be.equal(fee1);
            expect(await cmAccount.getServiceFeeByName(services.serviceName2)).to.be.equal(fee2);
            expect(await cmAccount.getServiceFeeByName(services.serviceName3)).to.be.equal(fee3);

            // Get specific capabilities from a service name
            expect(await cmAccount.getServiceCapabilitiesByName(services.serviceName1)).to.be.deep.equal(capabilities1);
            expect(await cmAccount.getServiceCapabilitiesByName(services.serviceName2)).to.be.deep.equal(capabilities2);
            expect(await cmAccount.getServiceCapabilitiesByName(services.serviceName3)).to.be.deep.equal(capabilities3);
        });

        it("should revert if the service is not registered", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            // Non registered service
            const serviceName = "cmp.service.accommodation.v0.AccommodationSearchService";

            const fee = 1000n;
            const restrictedRate = false;
            const capabilities = [];

            await expect(
                cmAccount.connect(signers.cmServiceAdmin).addService(serviceName, fee, restrictedRate, capabilities),
            ).to.be.revertedWithCustomError(cmAccountManager, "ServiceNotRegistered");
        });
    });

    describe("Payment", function () {
        it("should set and remove payment info correctly", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            // Get off chain payment supported expecting false
            expect(await cmAccount.offChainPaymentSupported()).to.be.equal(false);

            // Set off chain payment supported
            await expect(cmAccount.connect(signers.cmAccountAdmin).setOffChainPaymentSupported(true))
                .to.emit(cmAccount, "OffChainPaymentSupportUpdated")
                .withArgs(true);

            // Get off chain payment supported expecting true
            expect(await cmAccount.offChainPaymentSupported()).to.be.equal(true);

            // Set supported tokens
            const supportedToken1 = "0x0000000000000000000000000000000000000001";
            const supportedToken2 = "0x0000000000000000000000000000000000000002";

            await expect(cmAccount.connect(signers.cmServiceAdmin).addSupportedToken(supportedToken1))
                .to.emit(cmAccount, "PaymentTokenAdded")
                .withArgs(supportedToken1);

            await expect(cmAccount.connect(signers.cmServiceAdmin).addSupportedToken(supportedToken2))
                .to.emit(cmAccount, "PaymentTokenAdded")
                .withArgs(supportedToken2);

            // Get supported tokens
            const supportedTokens = await cmAccount.getSupportedTokens();
            expect(supportedTokens).to.be.deep.equal([supportedToken1, supportedToken2]);

            // Revert if token is already supported
            await expect(cmAccount.connect(signers.cmServiceAdmin).addSupportedToken(supportedToken1))
                .to.be.revertedWithCustomError(cmAccount, "PaymentTokenAlreadyExists")
                .withArgs(supportedToken1);

            // Remove supported token
            await expect(cmAccount.connect(signers.cmServiceAdmin).removeSupportedToken(supportedToken1))
                .to.emit(cmAccount, "PaymentTokenRemoved")
                .withArgs(supportedToken1);

            // Get supported tokens, should only return supportedToken2
            const supportedTokensAfterRemoval = await cmAccount.getSupportedTokens();
            expect(supportedTokensAfterRemoval).to.be.deep.equal([supportedToken2]);
        });
    });
    describe("PublicKeys", function () {
        it("should add a public keys correctly", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            // Pubkey
            const pubkey =
                "0x04fbe3e51d1e56c8ff935360cd32931f5a13ce4aac17f18ed8265c33f06468532fcb8b84eba84c0fae7ce88f64f97e7b6c7cf847b32b697b9e304de7ad2842e6ab";
            // Address of the public key
            const addr = ethers.computeAddress(pubkey);

            const publicKeyUse = 0;

            await expect(cmAccount.connect(signers.cmServiceAdmin).addPublicKey(addr, pubkey, publicKeyUse))
                .to.emit(cmAccount, "PublicKeyAdded")
                .withArgs(addr, [publicKeyUse, pubkey]);

            // Get public keys and check if they are correct, should include only addr and pubkey
            const publicKeys = await cmAccount.getPublicKeys();
            expect(publicKeys).to.be.deep.equal([[addr], [[publicKeyUse, pubkey]]]);
        });

        it("should remove a public key correctly", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            // Pubkey
            const pubkey =
                "0x04fbe3e51d1e56c8ff935360cd32931f5a13ce4aac17f18ed8265c33f06468532fcb8b84eba84c0fae7ce88f64f97e7b6c7cf847b32b697b9e304de7ad2842e6ab";
            // Address of the public key
            const addr = ethers.computeAddress(pubkey);

            const publicKeyUse = 0;

            await expect(cmAccount.connect(signers.cmServiceAdmin).addPublicKey(addr, pubkey, publicKeyUse))
                .to.emit(cmAccount, "PublicKeyAdded")
                .withArgs(addr, [publicKeyUse, pubkey]);

            await expect(cmAccount.connect(signers.cmServiceAdmin).removePublicKey(addr))
                .to.emit(cmAccount, "PublicKeyRemoved")
                .withArgs(addr);

            // Get public keys, it should be a array of two empty arrays
            const publicKeys = await cmAccount.getPublicKeys();
            expect(publicKeys).to.be.deep.equal([[], []]);
        });

        it("should get public keys correctly", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            const publicKeyUse = 0;

            // Pubkey 1
            const pubkey1 =
                "0x04fbe3e51d1e56c8ff935360cd32931f5a13ce4aac17f18ed8265c33f06468532fcb8b84eba84c0fae7ce88f64f97e7b6c7cf847b32b697b9e304de7ad2842e6ab";
            // Address of the public key
            const addr1 = ethers.computeAddress(pubkey1);

            await expect(cmAccount.connect(signers.cmServiceAdmin).addPublicKey(addr1, pubkey1, publicKeyUse))
                .to.emit(cmAccount, "PublicKeyAdded")
                .withArgs(addr1, [publicKeyUse, pubkey1]);

            // Pubkey 2
            const pubkey2 =
                "0x0407960fdb1ac968edc84eefe2aa4c5edc5b37ea0886eb4efecfd81c5993f9b00c77fc97dd94dc258fcf3c420f8a0601a8cb76030f2ffce68d104e7d83888083e5";
            // Address of the public key
            const addr2 = ethers.computeAddress(pubkey2);

            await expect(cmAccount.connect(signers.cmServiceAdmin).addPublicKey(addr2, pubkey2, publicKeyUse))
                .to.emit(cmAccount, "PublicKeyAdded")
                .withArgs(addr2, [publicKeyUse, pubkey2]);

            // Get public keys
            const publicKeys = await cmAccount.getPublicKeys();
            expect(publicKeys).to.be.deep.equal([
                [addr1, addr2],
                [
                    [publicKeyUse, pubkey1],
                    [publicKeyUse, pubkey2],
                ],
            ]);
        });

        it("should revert when adding the same public key", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            const publicKeyUse = 0;

            // Pubkey
            const pubkey =
                "0x04fbe3e51d1e56c8ff935360cd32931f5a13ce4aac17f18ed8265c33f06468532fcb8b84eba84c0fae7ce88f64f97e7b6c7cf847b32b697b9e304de7ad2842e6ab";
            // Address of the public key
            const addr = ethers.computeAddress(pubkey);

            await expect(cmAccount.connect(signers.cmServiceAdmin).addPublicKey(addr, pubkey, publicKeyUse))
                .to.emit(cmAccount, "PublicKeyAdded")
                .withArgs(addr, [publicKeyUse, pubkey]);

            await expect(cmAccount.connect(signers.cmServiceAdmin).addPublicKey(addr, pubkey, publicKeyUse))
                .to.be.revertedWithCustomError(cmAccount, "PublicKeyAlreadyExists")
                .withArgs(addr);
        });

        it("should revert when adding a public key with an invalid use type", async function () {
            const { cmAccountManager, cmAccount } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            const publicKeyUse = 255;

            // Pubkey
            const pubkey =
                "0x04fbe3e51d1e56c8ff935360cd32931f5a13ce4aac17f18ed8265c33f06468532fcb8b84eba84c0fae7ce88f64f97e7b6c7cf847b32b697b9e304de7ad2842e6ab";
            // Address of the public key
            const addr = ethers.computeAddress(pubkey);

            await expect(cmAccount.connect(signers.cmServiceAdmin).addPublicKey(addr, pubkey, publicKeyUse))
                .to.be.revertedWithCustomError(cmAccount, "InvalidPublicKeyUseType")
                .withArgs(publicKeyUse);
        });
    });
});
