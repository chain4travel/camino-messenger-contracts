const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers } = require("hardhat");

// Fixtures
const {
    setupSigners,
    developerFeeBp,
    deployCMAccountManagerFixture,
    deployCMAccountImplFixture,
    deployCMAccountManagerWithCMAccountImplFixture,
    deployAndConfigureAllFixture,
    deployCMAccountWithDepositFixture,
} = require("./utils/fixtures");

// Cheque utils
const {
    calculateMessengerChequeTypeHash,
    calculateTypedDataHash,
    calculateMessengerChequeHash,
    calculateDomainTypeHash,
    calculateDomainSeparator,
    calculateDomainSeparatorCamino,
    calculateDomainSeparatorColumbus,
    calculateDomainSeparatorKopernikus,
    calculateDomainSeparatorForChain,
    signMessengerCheque,
    signInvalidMessengerCheque,
    _signMessengerCheque,
} = require("../utils/cheques.js");
const { create } = require("domain");

describe("ChequeManager", function () {
    describe("Main", function () {
        it("Should return the correct MESSENGER_CHEQUE_TYPEHASH", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            const calculatedMessengerChequeTypeHash = calculateMessengerChequeTypeHash();

            const cmAccountMessengerChequeTypeHash = await cmAccount.MESSENGER_CHEQUE_TYPEHASH();
            expect(cmAccountMessengerChequeTypeHash).to.be.equal(calculatedMessengerChequeTypeHash);
        });

        it("Should return the correct DOMAIN_TYPEHASH", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            const calculatedDomainTypeHash = calculateDomainTypeHash();

            const cmAccountDomainTypeHash = await cmAccount.DOMAIN_TYPEHASH();
            expect(cmAccountDomainTypeHash).to.be.equal(calculatedDomainTypeHash);
        });

        it("Should initialize the DOMAIN_SEPARATOR correctly", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            const chainId = await ethers.provider.getNetwork().then((n) => n.chainId);
            const calculatedDomainSeparator = calculateDomainSeparatorForChain(chainId);

            const cmAccountDomainSeparator = await cmAccount.getDomainSeparator();
            expect(cmAccountDomainSeparator).to.be.equal(calculatedDomainSeparator);
        });

        it("Should hash the messenger cheque correctly", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);
            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: signers.chequeOperator.address,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("1"),
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
            };

            const calculatedHash = calculateMessengerChequeHash(cheque);

            const hashFromContract = await cmAccount.hashMessengerCheque(
                cheque.fromCMAccount,
                cheque.toCMAccount,
                cheque.toBot,
                cheque.counter,
                cheque.amount,
                cheque.createdAt,
                cheque.expiresAt,
            );

            expect(hashFromContract).to.be.equal(calculatedHash);
        });

        it("Should hash TypedData correctly", async function () {
            // Set up signers and contract instance
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            // Create a MessengerCheque object
            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: signers.chequeOperator.address,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("1"),
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
            };

            // Calculate domain separator
            const chainId = await ethers.provider.getNetwork().then((n) => n.chainId);
            const calculatedDomainSeparator = calculateDomainSeparatorForChain(chainId);

            // Calculate typedDataHash
            const calculatedTypedDataHash = calculateTypedDataHash(cheque, calculatedDomainSeparator);

            // Get typedDataHash from contract
            const typedDataHashFromContract = await cmAccount.hashTypedDataV4(
                cheque.fromCMAccount,
                cheque.toCMAccount,
                cheque.toBot,
                cheque.counter,
                cheque.amount,
                cheque.createdAt,
                cheque.expiresAt,
            );

            // Assert that the calculated typedDataHash is equal to the typedDataHash from contract
            expect(typedDataHashFromContract).to.be.equal(calculatedTypedDataHash);
        });
    });

    describe("Cheque Operations", function () {
        it("Should verify a cheque with a valid signature", async function () {
            const { cmAccount, cmAccountManager, prefundAmount } = await loadFixture(deployCMAccountWithDepositFixture);

            // Create receiving account (toCMAccount)
            const tx = await cmAccountManager.createCMAccount(
                signers.cmAccountAdmin.address,
                signers.cmAccountUpgrader.address,
                { value: prefundAmount },
            );

            const receipt = await tx.wait();

            // Parse event to get the CMAccount address
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = cmAccountManager.interface.parseLog(event);
            const toCMAccountAddress = parsedEvent.args.account;

            // Define cheque
            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: toCMAccountAddress,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("1"),
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
            };

            // Grant CHEQUE_OPERATOR_ROLE
            await cmAccount
                .connect(signers.cmAccountAdmin)
                .grantRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), signers.chequeOperator.address);

            // Sign Cheque
            const signature = await signMessengerCheque(cheque, signers.chequeOperator);

            // Verify cheque
            const verifyResponse = await cmAccount.verifyCheque(
                cheque.fromCMAccount,
                cheque.toCMAccount,
                cheque.toBot,
                cheque.counter,
                cheque.amount,
                cheque.createdAt,
                cheque.expiresAt,
                signature,
            );

            // Should emit ChequeVerified event with correct data
            await expect(await verifyResponse)
                .to.emit(cmAccount, "ChequeVerified")
                .withArgs(
                    cheque.fromCMAccount,
                    cheque.toCMAccount,
                    signers.chequeOperator,
                    cheque.toBot,
                    cheque.counter,
                    cheque.amount,
                    cheque.amount, // payment
                );

            // Sanity checks: balances should not change
            await expect(await verifyResponse).to.changeEtherBalance(signers.chequeOperator, 0, { includeFee: true });
            await expect(await verifyResponse).to.changeEtherBalance(cmAccount, 0, { includeFee: true });
        });

        it("Should not verify a cheque with an invalid signature", async function () {
            const { cmAccount, cmAccountManager, prefundAmount } = await loadFixture(deployCMAccountWithDepositFixture);

            // Create receiving account (toCMAccount)
            const tx = await cmAccountManager.createCMAccount(
                signers.cmAccountAdmin.address,
                signers.cmAccountUpgrader.address,
                { value: prefundAmount },
            );

            const receipt = await tx.wait();

            // Parse event to get the CMAccount address
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = cmAccountManager.interface.parseLog(event);
            const toCMAccountAddress = parsedEvent.args.account;

            // Define cheque
            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: toCMAccountAddress,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("1"),
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
            };

            // Grant CHEQUE_OPERATOR_ROLE
            await cmAccount
                .connect(signers.cmAccountAdmin)
                .grantRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), signers.chequeOperator.address);

            // Sign invalid cheque. Function below changes the chainId to invalidate the cheque.
            const signature = await signInvalidMessengerCheque(cheque, signers.chequeOperator);

            // Verify cheque, should revert and have the wrong address in the event
            // Because invalid signatures return a different address, we used a predicate to verify that
            // it's not the expected signer.
            await expect(
                cmAccount.verifyCheque(
                    cheque.fromCMAccount,
                    cheque.toCMAccount,
                    cheque.toBot,
                    cheque.counter,
                    cheque.amount,
                    cheque.createdAt,
                    cheque.expiresAt,
                    signature,
                ),
            )
                .to.be.revertedWithCustomError(cmAccount, "NotAllowedToSignCheques")
                .withArgs((addr) => addr !== signers.chequeOperator.address);
        });

        it("Should not verify a cheque with non-allowed signer", async function () {
            const { cmAccount, cmAccountManager, prefundAmount } = await loadFixture(deployCMAccountWithDepositFixture);

            // Create receiving account (toCMAccount)
            const tx = await cmAccountManager.createCMAccount(
                signers.cmAccountAdmin.address,
                signers.cmAccountUpgrader.address,
                { value: prefundAmount },
            );

            const receipt = await tx.wait();

            // Parse event to get the CMAccount address
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = cmAccountManager.interface.parseLog(event);
            const toCMAccountAddress = parsedEvent.args.account;

            // Define cheque
            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: toCMAccountAddress,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("1"),
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
            };

            // Be sure that the signer does not have the CHEQUE_OPERATOR_ROLE role
            const CHEQUE_OPERATOR_ROLE = await cmAccount.CHEQUE_OPERATOR_ROLE();
            expect(await cmAccountManager.hasRole(CHEQUE_OPERATOR_ROLE, signers.chequeOperator.address)).to.be.false;

            // Sign the cheque. Signature is valid but the signer is not allowed to sign on the `fromCMAccount`
            const signature = await signMessengerCheque(cheque, signers.chequeOperator);

            // Verify cheque, should revert
            await expect(
                cmAccount.verifyCheque(
                    cheque.fromCMAccount,
                    cheque.toCMAccount,
                    cheque.toBot,
                    cheque.counter,
                    cheque.amount,
                    cheque.createdAt,
                    cheque.expiresAt,
                    signature,
                ),
            )
                .to.be.revertedWithCustomError(cmAccount, "NotAllowedToSignCheques")
                .withArgs(signers.chequeOperator.address);
        });

        it("Should not verify an expired cheque", async function () {
            const { cmAccount, cmAccountManager, prefundAmount } = await loadFixture(deployCMAccountWithDepositFixture);

            // Create receiving account (toCMAccount)
            const tx = await cmAccountManager.createCMAccount(
                signers.cmAccountAdmin.address,
                signers.cmAccountUpgrader.address,
                { value: prefundAmount },
            );

            const receipt = await tx.wait();

            // Parse event to get the CMAccount address
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = cmAccountManager.interface.parseLog(event);
            const toCMAccountAddress = parsedEvent.args.account;

            const createdAt = ethers.toBigInt(Math.floor(Date.now() / 1000)) - 10000n; // Go back in time
            const expiresAt = createdAt + 120n; // Expiration in 2 minutes, but still in the past

            // Define cheque
            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: toCMAccountAddress,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("1"),
                createdAt: createdAt,
                expiresAt: expiresAt,
            };

            // Grant CHEQUE_OPERATOR_ROLE
            await cmAccount
                .connect(signers.cmAccountAdmin)
                .grantRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), signers.chequeOperator.address);

            // Sign the cheque
            const signature = await signMessengerCheque(cheque, signers.chequeOperator);

            // Verify cheque, should revert with ChequeExpired
            await expect(
                cmAccount.verifyCheque(
                    cheque.fromCMAccount,
                    cheque.toCMAccount,
                    cheque.toBot,
                    cheque.counter,
                    cheque.amount,
                    cheque.createdAt,
                    cheque.expiresAt,
                    signature,
                ),
            )
                .to.be.revertedWithCustomError(cmAccount, "ChequeExpired")
                .withArgs(expiresAt);
        });

        it("Should cash-in multiple cheques correctly", async function () {
            const { cmAccount, cmAccountManager, prefundAmount } = await loadFixture(deployCMAccountWithDepositFixture);

            // Create receiving account (toCMAccount)
            const tx = await cmAccountManager.createCMAccount(
                signers.cmAccountAdmin.address,
                signers.cmAccountUpgrader.address,
                { value: prefundAmount },
            );

            const receipt = await tx.wait();

            // Parse event to get the CMAccount address
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = cmAccountManager.interface.parseLog(event);
            const toCMAccountAddress = parsedEvent.args.account;

            const createdAt = ethers.toBigInt(Math.floor(Date.now() / 1000));

            // Define cheque
            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: toCMAccountAddress,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("0.1"),
                createdAt: createdAt,
                expiresAt: createdAt + 300n,
            };

            // Grant CHEQUE_OPERATOR_ROLE
            await cmAccount
                .connect(signers.cmAccountAdmin)
                .grantRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), signers.chequeOperator.address);

            // Sign Cheque
            const signature = await signMessengerCheque(cheque, signers.chequeOperator);

            // Calculate developer fee
            const developerFeeBp = await cmAccount.getDeveloperFeeBp();
            const developerFee = (cheque.amount * developerFeeBp) / 10000n;

            // Cash-in cheque
            const cashInResponse = await cmAccount.cashInCheque(
                cheque.fromCMAccount,
                cheque.toCMAccount,
                cheque.toBot,
                cheque.counter,
                cheque.amount,
                cheque.createdAt,
                cheque.expiresAt,
                signature,
            );

            // CMAccount balance should decrease by cheque amount (developer fee cut is taken from the cheque amount)
            await expect(await cashInResponse).to.changeEtherBalance(cmAccount, -cheque.amount);

            // toCMAccount balance should increase by cheque amount - developerFee
            await expect(await cashInResponse).to.changeEtherBalance(toCMAccountAddress, cheque.amount - developerFee);

            // DeveloperWallet balance should increase by developerFee
            await expect(await cashInResponse).to.changeEtherBalance(signers.developerWallet, developerFee);

            // Should emit event with correct data
            await expect(await cashInResponse)
                .to.emit(cmAccount, "ChequeCashedIn")
                .withArgs(
                    signers.chequeOperator.address,
                    cheque.toBot,
                    cheque.counter,
                    cheque.amount - developerFee, // paid amount
                    developerFee, // developer cut
                );

            // Sanity checks: should set lastCashIns
            const lastCashIn = await cmAccount.getLastCashIn(signers.chequeOperator, cheque.toBot);
            expect(lastCashIn).to.be.deep.equal([cheque.counter, cheque.amount, createdAt, createdAt + 300n]);
            // Check total cheque payments
            // Total cheque payments should be equal to the last cheque amount
            // because we use same from/to CM accounts
            expect(await cmAccount.getTotalChequePayments()).to.be.equal(cheque.amount);

            /**
             * Second cheque
             */

            const createdAt2 = ethers.toBigInt(Math.floor(Date.now() / 1000));

            // New cheque with a higher counter and amount
            const cheque2 = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: toCMAccountAddress,
                toBot: signers.otherAccount2.address,
                counter: 100,
                amount: ethers.parseEther("0.234"),
                createdAt: createdAt2,
                expiresAt: createdAt2 + 300n,
            };

            // Sign Cheque
            const signature2 = await signMessengerCheque(cheque2, signers.chequeOperator);

            // Calculate developer fee
            const developerFee2 = ((cheque2.amount - cheque.amount) * developerFeeBp) / 10000n;

            // Cash-in cheque
            const cashInResponse2 = await cmAccount.cashInCheque(
                cheque2.fromCMAccount,
                cheque2.toCMAccount,
                cheque2.toBot,
                cheque2.counter,
                cheque2.amount,
                cheque2.createdAt,
                cheque2.expiresAt,
                signature2,
            );

            // CMAccount balance descrease by (cheque2 amount - cheque amount)
            await expect(await cashInResponse2).to.changeEtherBalance(
                cmAccount,
                -cheque2.amount + cheque.amount, // Weird calculation but it works
            );

            // toCMAccount balance increase by (cheque2 amount - cheque amount) - developerFee2
            await expect(await cashInResponse2).to.changeEtherBalance(
                toCMAccountAddress,
                cheque2.amount - cheque.amount - developerFee2, // new cheque amount minus the lastCashIn amount
            );

            // DeveloperWallet balance increase by developerFee
            await expect(await cashInResponse2).to.changeEtherBalance(signers.developerWallet, developerFee2);

            // Should emit event with correct data
            await expect(await cashInResponse2)
                .to.emit(cmAccount, "ChequeCashedIn")
                .withArgs(
                    signers.chequeOperator.address,
                    cheque2.toBot,
                    cheque2.counter,
                    cheque2.amount - cheque.amount - developerFee2, // paid amount for this cheque
                    developerFee2,
                );

            // Sanity checks: should set lastCashIns
            expect(await cmAccount.getLastCashIn(signers.chequeOperator, cheque.toBot)).to.be.deep.equal([
                cheque2.counter,
                cheque2.amount,
                createdAt2,
                createdAt2 + 300n,
            ]);
            // Check total cheque payments
            // Total cheque payments should be equal to the last cheque amount
            // because we use same from/to CM account pairs for cheques above
            expect(await cmAccount.getTotalChequePayments()).to.be.equal(cheque2.amount);
        });

        it("Should not update total cheque payments for same account", async function () {
            const { cmAccount, cmAccountManager, prefundAmount } = await loadFixture(deployCMAccountWithDepositFixture);

            // Define cheque
            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: await cmAccount.getAddress(),
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("0.1"),
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
            };

            // Grant CHEQUE_OPERATOR_ROLE
            await cmAccount
                .connect(signers.cmAccountAdmin)
                .grantRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), signers.chequeOperator.address);

            // Sign Cheque
            const signature = await signMessengerCheque(cheque, signers.chequeOperator);

            // Initial total cheque payments should be zero
            expect(await cmAccount.getTotalChequePayments()).to.be.equal(0n);

            // Cash-in cheque
            const cashInResponse = await cmAccount.cashInCheque(
                cheque.fromCMAccount,
                cheque.toCMAccount,
                cheque.toBot,
                cheque.counter,
                cheque.amount,
                cheque.createdAt,
                cheque.expiresAt,
                signature,
            );
            await expect(cashInResponse).to.be.not.reverted;

            // After cash-in total cheque payments should still be zero because the
            // cheque is from the same account (fromCMAccount === toCMAccount)
            expect(await cmAccount.getTotalChequePayments()).to.be.equal(0n);
        });
    });
});
