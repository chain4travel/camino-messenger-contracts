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

        it("should add and return all supported services correctly + setter/getters test", async function () {
            const { cmAccountManager, cmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            const fee1 = 1000n;
            const fee2 = 2000n;
            const fee3 = 3000n;

            const restrictedRate1 = false;
            const restrictedRate2 = true;
            const restrictedRate3 = false;

            const capabilities1 = ["test capability 1"];
            const capabilities2 = ["test capability 2"];
            const capabilities3 = ["test capability 3"];

            // Add services to CM account
            expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    .addService(services.serviceName1, fee1, restrictedRate1, capabilities1),
            )
                .to.emit(cmAccount, "ServiceAdded")
                .withArgs(services.serviceHash1);

            expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    .addService(services.serviceName2, fee2, restrictedRate2, capabilities2),
            )
                .to.emit(cmAccount, "ServiceAdded")
                .withArgs(services.serviceHash2);

            expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    .addService(services.serviceName3, fee3, restrictedRate3, capabilities3),
            )
                .to.emit(cmAccount, "ServiceAdded")
                .withArgs(services.serviceHash3);

            // Get all services
            const servicesFromCMAccount = await cmAccount.getSupportedServices();
            expect(servicesFromCMAccount).to.be.deep.equal([
                [services.serviceName1, services.serviceName2, services.serviceName3],
                [
                    [fee1, restrictedRate1, capabilities1],
                    [fee2, restrictedRate2, capabilities2],
                    [fee3, restrictedRate3, capabilities3],
                ],
            ]);

            // The contract calls below uses
            // contract["functionName(string)"](string) to get the correct
            // overloaded function

            // Get specific fee for a service name
            expect(await cmAccount["getServiceFee(string)"](services.serviceName1)).to.be.equal(fee1);
            expect(await cmAccount["getServiceFee(string)"](services.serviceName2)).to.be.equal(fee2);
            expect(await cmAccount["getServiceFee(string)"](services.serviceName3)).to.be.equal(fee3);

            // Get specific restricted rate for a service name
            expect(await cmAccount["getServiceRestrictedRate(string)"](services.serviceName1)).to.be.equal(
                restrictedRate1,
            );
            expect(await cmAccount["getServiceRestrictedRate(string)"](services.serviceName2)).to.be.equal(
                restrictedRate2,
            );
            expect(await cmAccount["getServiceRestrictedRate(string)"](services.serviceName3)).to.be.equal(
                restrictedRate3,
            );

            // Get specific capabilities for a service name
            expect(await cmAccount["getServiceCapabilities(string)"](services.serviceName1)).to.be.deep.equal(
                capabilities1,
            );
            expect(await cmAccount["getServiceCapabilities(string)"](services.serviceName2)).to.be.deep.equal(
                capabilities2,
            );
            expect(await cmAccount["getServiceCapabilities(string)"](services.serviceName3)).to.be.deep.equal(
                capabilities3,
            );

            // TEST SETTERS
            // with new values for each service field

            const newFee1 = 4000n;
            const newFee2 = 5000n;
            const newFee3 = 6000n;

            const newRestrictedRate1 = true;
            const newRestrictedRate2 = false;
            const newRestrictedRate3 = true;

            const newCapabilities1 = ["test capability 4"];
            const newCapabilities2 = ["test capability 5"];
            const newCapabilities3 = ["test capability 6"];

            // Fee Setter
            await expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    ["setServiceFee(string,uint256)"](services.serviceName1, newFee1),
            )
                .to.emit(cmAccount, "ServiceFeeUpdated")
                .withArgs(services.serviceHash1, newFee1);

            await expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    ["setServiceFee(string,uint256)"](services.serviceName2, newFee2),
            )
                .to.emit(cmAccount, "ServiceFeeUpdated")
                .withArgs(services.serviceHash2, newFee2);

            await expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    ["setServiceFee(string,uint256)"](services.serviceName3, newFee3),
            )
                .to.emit(cmAccount, "ServiceFeeUpdated")
                .withArgs(services.serviceHash3, newFee3);

            // Restricted Rate Setter
            await expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    ["setServiceRestrictedRate(string,bool)"](services.serviceName1, newRestrictedRate1),
            )
                .to.emit(cmAccount, "ServiceRestrictedRateUpdated")
                .withArgs(services.serviceHash1, newRestrictedRate1);

            await expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    ["setServiceRestrictedRate(string,bool)"](services.serviceName2, newRestrictedRate2),
            )
                .to.emit(cmAccount, "ServiceRestrictedRateUpdated")
                .withArgs(services.serviceHash2, newRestrictedRate2);

            await expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    ["setServiceRestrictedRate(string,bool)"](services.serviceName3, newRestrictedRate3),
            ).to.emit(cmAccount, "ServiceRestrictedRateUpdated");

            // Capabilities Setter
            await expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    ["setServiceCapabilities(string,string[])"](services.serviceName1, newCapabilities1),
            )
                .to.emit(cmAccount, "ServiceCapabilitiesUpdated")
                .withArgs(services.serviceHash1);

            await expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    ["setServiceCapabilities(string,string[])"](services.serviceName2, newCapabilities2),
            )
                .to.emit(cmAccount, "ServiceCapabilitiesUpdated")
                .withArgs(services.serviceHash2);

            await expect(
                await cmAccount
                    .connect(signers.cmServiceAdmin)
                    ["setServiceCapabilities(string,string[])"](services.serviceName3, newCapabilities3),
            )
                .to.emit(cmAccount, "ServiceCapabilitiesUpdated")
                .withArgs(services.serviceHash3);

            // TEST GETTERS with hashes

            // Get specific fee for a service name
            expect(await cmAccount["getServiceFee(bytes32)"](services.serviceHash1)).to.be.equal(newFee1);
            expect(await cmAccount["getServiceFee(bytes32)"](services.serviceHash2)).to.be.equal(newFee2);
            expect(await cmAccount["getServiceFee(bytes32)"](services.serviceHash3)).to.be.equal(newFee3);

            // Get specific restricted rate for a service name
            expect(await cmAccount["getServiceRestrictedRate(bytes32)"](services.serviceHash1)).to.be.equal(
                newRestrictedRate1,
            );
            expect(await cmAccount["getServiceRestrictedRate(bytes32)"](services.serviceHash2)).to.be.equal(
                newRestrictedRate2,
            );
            expect(await cmAccount["getServiceRestrictedRate(bytes32)"](services.serviceHash3)).to.be.equal(
                newRestrictedRate3,
            );

            // Get specific capabilities for a service name
            expect(await cmAccount["getServiceCapabilities(bytes32)"](services.serviceHash1)).to.be.deep.equal(
                newCapabilities1,
            );
            expect(await cmAccount["getServiceCapabilities(bytes32)"](services.serviceHash2)).to.be.deep.equal(
                newCapabilities2,
            );
            expect(await cmAccount["getServiceCapabilities(bytes32)"](services.serviceHash3)).to.be.deep.equal(
                newCapabilities3,
            );
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

    describe("Wanted Services", function () {
        it("should add a wanted service correctly", async function () {
            const { cmAccountManager, cmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(cmAccount.connect(signers.cmServiceAdmin).addWantedServices([services.serviceName1]))
                .to.emit(cmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash1);
        });

        it("should add multiple (6) wanted services correctly", async function () {
            const { cmAccountManager, cmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(
                cmAccount
                    .connect(signers.cmServiceAdmin)
                    .addWantedServices([
                        services.serviceName1,
                        services.serviceName2,
                        services.serviceName3,
                        services.serviceName4,
                        services.serviceName5,
                        services.serviceName6,
                    ]),
            )
                .to.emit(cmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash1)
                .to.emit(cmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash2)
                .to.emit(cmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash3)
                .to.emit(cmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash4)
                .to.emit(cmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash5)
                .to.emit(cmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash6);
        });

        it("should revert if a wanted service is already added", async function () {
            const { cmAccountManager, cmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(cmAccount.connect(signers.cmServiceAdmin).addWantedServices([services.serviceName1]))
                .to.emit(cmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash1);

            await expect(
                cmAccount.connect(signers.cmServiceAdmin).addWantedServices([services.serviceName1]),
            ).to.be.revertedWithCustomError(cmAccount, "WantedServiceAlreadyExists");
        });

        it("should revert if a wanted service is not registered", async function () {
            const { cmAccountManager, cmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(
                cmAccount
                    .connect(signers.cmServiceAdmin)
                    .addWantedServices(["cmp.service.test.v0.NonRegisteredService"]),
            ).to.be.revertedWithCustomError(cmAccountManager, "ServiceNotRegistered");
        });

        it("should remove a wanted service correctly", async function () {
            const { cmAccountManager, cmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(cmAccount.connect(signers.cmServiceAdmin).addWantedServices([services.serviceName1]))
                .to.emit(cmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash1);

            await expect(cmAccount.connect(signers.cmServiceAdmin).removeWantedServices([services.serviceName1]))
                .to.emit(cmAccount, "WantedServiceRemoved")
                .withArgs(services.serviceHash1);
        });

        it("should remove multiple wanted services correctly", async function () {
            const { cmAccountManager, cmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(
                cmAccount
                    .connect(signers.cmServiceAdmin)
                    .addWantedServices([services.serviceName1, services.serviceName2]),
            )
                .to.emit(cmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash1)
                .to.emit(cmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash2);

            await expect(
                cmAccount
                    .connect(signers.cmServiceAdmin)
                    .removeWantedServices([services.serviceName1, services.serviceName2]),
            )
                .to.emit(cmAccount, "WantedServiceRemoved")
                .withArgs(services.serviceHash1)
                .to.emit(cmAccount, "WantedServiceRemoved")
                .withArgs(services.serviceHash2);
        });

        it("should revert removal if a wanted service does not exist", async function () {
            const { cmAccountManager, cmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(
                cmAccount.connect(signers.cmServiceAdmin).removeWantedServices([services.serviceName1]),
            ).to.be.revertedWithCustomError(cmAccount, "WantedServiceDoesNotExist");
        });

        it("should add and get multiple wanted services correctly", async function () {
            const { cmAccountManager, cmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(cmAccount.connect(signers.cmServiceAdmin).addWantedServices([services.serviceName1]))
                .to.emit(cmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash1);

            await expect(
                cmAccount
                    .connect(signers.cmServiceAdmin)
                    .addWantedServices([services.serviceName2, services.serviceName3]),
            )
                .to.emit(cmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash2)
                .to.emit(cmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash3);

            // Get wanted services
            expect(await cmAccount.getWantedServices()).to.be.deep.equal([
                services.serviceName1,
                services.serviceName2,
                services.serviceName3,
            ]);

            // Get wanted service by hash
            expect(await cmAccount.getWantedServiceHashes()).to.be.deep.equal([
                services.serviceHash1,
                services.serviceHash2,
                services.serviceHash3,
            ]);
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
                .withArgs(addr);

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
                .withArgs(addr);

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
                .withArgs(addr1);

            // Pubkey 2
            const pubkey2 =
                "0x0407960fdb1ac968edc84eefe2aa4c5edc5b37ea0886eb4efecfd81c5993f9b00c77fc97dd94dc258fcf3c420f8a0601a8cb76030f2ffce68d104e7d83888083e5";
            // Address of the public key
            const addr2 = ethers.computeAddress(pubkey2);

            await expect(cmAccount.connect(signers.cmServiceAdmin).addPublicKey(addr2, pubkey2, publicKeyUse))
                .to.emit(cmAccount, "PublicKeyAdded")
                .withArgs(addr2);

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
                .withArgs(addr);

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
