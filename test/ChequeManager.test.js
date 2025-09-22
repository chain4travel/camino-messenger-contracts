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
    deployNullUSDFixture,
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
            const { cmAccount, nullUSD } = await loadFixture(deployCMAccountWithDepositFixture);

            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: signers.chequeOperator.address,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("1"),
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
                paymentToken: await nullUSD.getAddress(),
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
                cheque.paymentToken,
            );

            expect(hashFromContract).to.be.equal(calculatedHash);
        });

        it("Should hash TypedData correctly", async function () {
            // Set up signers and contract instance
            const { cmAccount, nullUSD } = await loadFixture(deployCMAccountWithDepositFixture);

            // Create a MessengerCheque object
            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: signers.chequeOperator.address,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("1"),
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
                paymentToken: await nullUSD.getAddress(),
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
                cheque.paymentToken,
            );

            // Assert that the calculated typedDataHash is equal to the typedDataHash from contract
            expect(typedDataHashFromContract).to.be.equal(calculatedTypedDataHash);
        });
    });

    describe("Cheque Operations", function () {
        it("Should verify a cheque with a valid signature", async function () {
            const { cmAccount, cmAccountManager, prefundAmount, nullUSD, nullUSDDecimals } = await loadFixture(
                deployCMAccountWithDepositFixture,
            );

            // Approve service fee
            await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);

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
                amount: ethers.parseUnits("1", nullUSDDecimals),
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
                paymentToken: await nullUSD.getAddress(),
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
                cheque.paymentToken,
                signature,
            );

            expect(verifyResponse).to.be.deep.equal([
                signers.chequeOperator.address,
                ethers.parseUnits("1", nullUSDDecimals),
            ]);
        });

        it("Should not verify a cheque with an invalid signature", async function () {
            const { cmAccount, cmAccountManager, prefundAmount, nullUSD, nullUSDDecimals } = await loadFixture(
                deployCMAccountWithDepositFixture,
            );

            // Approve service fee
            await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);

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
                paymentToken: await nullUSD.getAddress(),
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
                    cheque.paymentToken,
                    signature,
                ),
            )
                .to.be.revertedWithCustomError(cmAccount, "NotAllowedToSignCheques")
                .withArgs((addr) => addr !== signers.chequeOperator.address);
        });

        it("Should not verify a cheque with non-allowed signer", async function () {
            const { cmAccount, cmAccountManager, prefundAmount, nullUSD, nullUSDDecimals } = await loadFixture(
                deployCMAccountWithDepositFixture,
            );

            // Approve service fee
            await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);

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
                paymentToken: await nullUSD.getAddress(),
            };

            // Be sure that the signer does not have the CHEQUE_OPERATOR_ROLE role
            const CHEQUE_OPERATOR_ROLE = await cmAccount.CHEQUE_OPERATOR_ROLE();
            expect(await cmAccount.hasRole(CHEQUE_OPERATOR_ROLE, signers.chequeOperator.address)).to.be.false;

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
                    cheque.paymentToken,
                    signature,
                ),
            )
                .to.be.revertedWithCustomError(cmAccount, "NotAllowedToSignCheques")
                .withArgs(signers.chequeOperator.address);
        });

        it("Should not verify a cheque if from/to is not CMAccount", async function () {
            const { cmAccount, cmAccountManager, prefundAmount, nullUSD, nullUSDDecimals } = await loadFixture(
                deployCMAccountWithDepositFixture,
            );

            // Approve service fee
            await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);

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

            // Define cheques

            const chequeWithInvalidFrom = {
                fromCMAccount: signers.otherAccount3.address,
                toCMAccount: toCMAccountAddress,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("1"),
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
                paymentToken: await nullUSD.getAddress(),
            };

            const chequeWithInvalidTo = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: signers.otherAccount3.address,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("1"),
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
                paymentToken: await nullUSD.getAddress(),
            };

            // Grant CHEQUE_OPERATOR_ROLE
            await cmAccount
                .connect(signers.cmAccountAdmin)
                .grantRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), signers.chequeOperator.address);

            const signatureFrom = await signMessengerCheque(chequeWithInvalidFrom, signers.chequeOperator);
            const signatureTo = await signMessengerCheque(chequeWithInvalidTo, signers.chequeOperator);

            // Verify cheques, should revert
            await expect(
                cmAccount.verifyCheque(
                    chequeWithInvalidFrom.fromCMAccount,
                    chequeWithInvalidFrom.toCMAccount,
                    chequeWithInvalidFrom.toBot,
                    chequeWithInvalidFrom.counter,
                    chequeWithInvalidFrom.amount,
                    chequeWithInvalidFrom.createdAt,
                    chequeWithInvalidFrom.expiresAt,
                    chequeWithInvalidFrom.paymentToken,
                    signatureFrom,
                ),
            )
                .to.be.revertedWithCustomError(cmAccount, "InvalidFromCMAccount")
                .withArgs(signers.otherAccount3.address);

            await expect(
                cmAccount.verifyCheque(
                    chequeWithInvalidTo.fromCMAccount,
                    chequeWithInvalidTo.toCMAccount,
                    chequeWithInvalidTo.toBot,
                    chequeWithInvalidTo.counter,
                    chequeWithInvalidTo.amount,
                    chequeWithInvalidTo.createdAt,
                    chequeWithInvalidTo.expiresAt,
                    chequeWithInvalidTo.paymentToken,
                    signatureTo,
                ),
            )
                .to.be.revertedWithCustomError(cmAccount, "InvalidToCMAccount")
                .withArgs(signers.otherAccount3.address);
        });

        it("Should not verify an expired cheque", async function () {
            const { cmAccount, cmAccountManager, prefundAmount, nullUSD, nullUSDDecimals } = await loadFixture(
                deployCMAccountWithDepositFixture,
            );

            // Approve service fee
            await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);

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
                paymentToken: await nullUSD.getAddress(),
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
                    cheque.paymentToken,
                    signature,
                ),
            )
                .to.be.revertedWithCustomError(cmAccount, "ChequeExpired")
                .withArgs(expiresAt);
        });

        it("Should not verify/cash in a cheque with an invalid payment token", async function () {
            const { cmAccount, cmAccountManager, prefundAmount, nullUSD, nullUSDDecimals } = await loadFixture(
                deployCMAccountWithDepositFixture,
            );

            // Approve service fee
            await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);

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
                paymentToken: "0x0000000000000000000000000000000000000001", // Invalid payment token
            };

            // Grant CHEQUE_OPERATOR_ROLE
            await cmAccount
                .connect(signers.cmAccountAdmin)
                .grantRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), signers.chequeOperator.address);

            // Sign the cheque
            const signature = await signMessengerCheque(cheque, signers.chequeOperator);

            // Verify cheque, should revert with InvalidPaymentToken
            await expect(
                cmAccount.verifyCheque(
                    cheque.fromCMAccount,
                    cheque.toCMAccount,
                    cheque.toBot,
                    cheque.counter,
                    cheque.amount,
                    cheque.createdAt,
                    cheque.expiresAt,
                    cheque.paymentToken,
                    signature,
                ),
            )
                .to.be.revertedWithCustomError(cmAccount, "InvalidPaymentToken")
                .withArgs(cheque.paymentToken, await nullUSD.getAddress());

            // Try to cash-in the cheque, should revert with InvalidPaymentToken
            await expect(
                cmAccount.cashInCheque(
                    cheque.fromCMAccount,
                    cheque.toCMAccount,
                    cheque.toBot,
                    cheque.counter,
                    cheque.amount,
                    cheque.createdAt,
                    cheque.expiresAt,
                    cheque.paymentToken,
                    signature,
                ),
            )
                .to.be.revertedWithCustomError(cmAccount, "InvalidPaymentToken")
                .withArgs(cheque.paymentToken, await nullUSD.getAddress());
        });

        it("Should cash-in multiple cheques correctly", async function () {
            const { cmAccount, cmAccountManager, prefundAmount, nullUSD, nullUSDDecimals } = await loadFixture(
                deployCMAccountWithDepositFixture,
            );

            // Approve service fee
            await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);

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
                paymentToken: await nullUSD.getAddress(),
            };

            // Grant CHEQUE_OPERATOR_ROLE
            await cmAccount
                .connect(signers.cmAccountAdmin)
                .grantRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), signers.chequeOperator.address);

            // Sign Cheque
            const signature = await signMessengerCheque(cheque, signers.chequeOperator);

            // Calculate developer fee
            const developerFeeBp = await cmAccountManager.getDeveloperFeeBp();
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
                cheque.paymentToken,
                signature,
            );

            // CMAccount balance should decrease by cheque amount (developer fee cut is taken from the cheque amount)
            await expect(cashInResponse).to.changeTokenBalance(nullUSD, cmAccount, -cheque.amount);

            // toCMAccount balance should increase by cheque amount - developerFee
            await expect(cashInResponse).to.changeTokenBalance(
                nullUSD,
                toCMAccountAddress,
                cheque.amount - developerFee,
            );

            // DeveloperWallet balance should increase by developerFee
            await expect(cashInResponse).to.changeTokenBalance(nullUSD, signers.developerWallet, developerFee);

            // Should emit event with correct data
            await expect(cashInResponse)
                .to.emit(cmAccount, "ChequeCashedIn")
                .withArgs(
                    cheque.fromCMAccount,
                    cheque.toCMAccount,
                    signers.chequeOperator.address, // fromBot
                    cheque.toBot,
                    cheque.counter,
                    cheque.amount,
                    cheque.amount - developerFee, // paid amount
                    developerFee, // developer cut
                    cheque.paymentToken,
                );

            // Sanity checks: should set lastCashIns
            const lastCashIn = await cmAccount.getLastCashIn(
                signers.chequeOperator.address,
                cheque.toBot,
                cheque.paymentToken,
            );
            expect(lastCashIn).to.be.deep.equal([cheque.counter, cheque.amount, createdAt, createdAt + 300n]);

            // Check total cheque payments
            // Total cheque payments should be equal to the last cheque amount
            // because we use same from/to CM accounts
            expect(await cmAccount.getTotalChequePaymentsPerToken(cheque.paymentToken)).to.be.equal(cheque.amount);

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
                paymentToken: await nullUSD.getAddress(),
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
                cheque2.paymentToken,
                signature2,
            );

            // CMAccount balance decrease by (cheque2 amount - cheque amount)
            await expect(cashInResponse2).to.changeTokenBalance(
                nullUSD,
                cmAccount,
                -cheque2.amount + cheque.amount, // Weird calculation but it works
            );

            // toCMAccount balance increase by (cheque2 amount - cheque amount) - developerFee2
            await expect(cashInResponse2).to.changeTokenBalance(
                nullUSD,
                toCMAccountAddress,
                cheque2.amount - cheque.amount - developerFee2, // new cheque amount minus the lastCashIn amount
            );

            // DeveloperWallet balance increase by developerFee
            await expect(cashInResponse2).to.changeTokenBalance(nullUSD, signers.developerWallet, developerFee2);

            // Should emit event with correct data
            await expect(cashInResponse2)
                .to.emit(cmAccount, "ChequeCashedIn")
                .withArgs(
                    cheque2.fromCMAccount,
                    cheque2.toCMAccount,
                    signers.chequeOperator.address, // fromBot
                    cheque2.toBot,
                    cheque2.counter,
                    cheque2.amount,
                    cheque2.amount - cheque.amount - developerFee2, // paid amount for this cheque
                    developerFee2,
                    cheque2.paymentToken,
                );

            // Sanity checks: should set lastCashIns
            expect(
                await cmAccount.getLastCashIn(signers.chequeOperator.address, cheque.toBot, cheque2.paymentToken),
            ).to.be.deep.equal([cheque2.counter, cheque2.amount, createdAt2, createdAt2 + 300n]);

            // Check total cheque payments
            // Total cheque payments should be equal to the last cheque amount
            // because we use same from/to CM account pairs for cheques above
            expect(await cmAccount.getTotalChequePaymentsPerToken(cheque2.paymentToken)).to.be.equal(cheque2.amount);

            // DIFFERENT CM ACCOUNT ----------------------------------------------------------------

            // Approve ERC20 service fee
            await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);

            // Create different CM Account to test total cheque payments
            const diffCMAccount_tx = await cmAccountManager.createCMAccount(
                signers.cmAccountAdmin.address,
                signers.cmAccountUpgrader.address,
                { value: prefundAmount },
            );

            const diffCMAccount_receipt = await diffCMAccount_tx.wait();

            // Parse event to get the CMAccount address
            const diffCMAccount_event = diffCMAccount_receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const diffCMAccount_parsedEvent = cmAccountManager.interface.parseLog(diffCMAccount_event);
            const diffCMAccountAddress = diffCMAccount_parsedEvent.args.account;

            // New cheque with a higher counter and amount
            const diffCMAccount_cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: diffCMAccountAddress,
                toBot: signers.otherAccount3.address, // Use different bot
                counter: 100,
                amount: ethers.parseEther("0.432"),
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
                paymentToken: await nullUSD.getAddress(),
            };

            // Sign Cheque
            const diffCMAccount_signature = await signMessengerCheque(diffCMAccount_cheque, signers.chequeOperator);

            // Cash-in cheque
            const diffCMAccount_cashInResponse = await cmAccount.cashInCheque(
                diffCMAccount_cheque.fromCMAccount,
                diffCMAccount_cheque.toCMAccount,
                diffCMAccount_cheque.toBot,
                diffCMAccount_cheque.counter,
                diffCMAccount_cheque.amount,
                diffCMAccount_cheque.createdAt,
                diffCMAccount_cheque.expiresAt,
                diffCMAccount_cheque.paymentToken,
                diffCMAccount_signature,
            );

            // Calculate developer fee
            const diffCMAccount_developerFee = (diffCMAccount_cheque.amount * developerFeeBp) / 10000n;

            // Should emit event with correct data
            await expect(diffCMAccount_cashInResponse)
                .to.emit(cmAccount, "ChequeCashedIn")
                .withArgs(
                    diffCMAccount_cheque.fromCMAccount,
                    diffCMAccount_cheque.toCMAccount,
                    signers.chequeOperator.address, // fromBot
                    diffCMAccount_cheque.toBot,
                    diffCMAccount_cheque.counter,
                    diffCMAccount_cheque.amount,
                    diffCMAccount_cheque.amount - diffCMAccount_developerFee, // paid amount for this cheque
                    diffCMAccount_developerFee,
                    diffCMAccount_cheque.paymentToken,
                );

            // CMAccount balance decrease by cheque amount
            await expect(diffCMAccount_cashInResponse).to.changeTokenBalance(
                nullUSD,
                cmAccount,
                -diffCMAccount_cheque.amount,
            );

            // diffCMAccount balance increase by cheque amount - developerFee
            await expect(diffCMAccount_cashInResponse).to.changeTokenBalance(
                nullUSD,
                diffCMAccountAddress,
                diffCMAccount_cheque.amount - diffCMAccount_developerFee,
            );

            // DeveloperWallet balance increase by developerFee
            await expect(diffCMAccount_cashInResponse).to.changeTokenBalance(
                nullUSD,
                signers.developerWallet,
                diffCMAccount_developerFee,
            );

            // Check total cheque payments per payment token, it should be equal to the sum of cheque2.amount and diffCMAccount_cheque.amount
            expect(await cmAccount.getTotalChequePaymentsPerToken(diffCMAccount_cheque.paymentToken)).to.be.equal(
                diffCMAccount_cheque.amount + cheque2.amount,
            );

            // CHECK INVALID AMOUNT AND COUNTER ----------------------------------------------------

            // Cheque with invalid amount
            const chequeWithInvalidAmount = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: toCMAccountAddress,
                toBot: signers.otherAccount2.address,
                counter: cheque2.counter + 1,
                amount: cheque2.amount - 1n,
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
                paymentToken: await nullUSD.getAddress(),
            };

            // Sign Cheque
            const signatureWithInvalidAmount = await signMessengerCheque(
                chequeWithInvalidAmount,
                signers.chequeOperator,
            );

            // Try to cash-in cheque with invalid amount
            await expect(
                cmAccount.cashInCheque(
                    chequeWithInvalidAmount.fromCMAccount,
                    chequeWithInvalidAmount.toCMAccount,
                    chequeWithInvalidAmount.toBot,
                    chequeWithInvalidAmount.counter,
                    chequeWithInvalidAmount.amount,
                    chequeWithInvalidAmount.createdAt,
                    chequeWithInvalidAmount.expiresAt,
                    chequeWithInvalidAmount.paymentToken,
                    signatureWithInvalidAmount,
                ),
            )
                .to.be.revertedWithCustomError(cmAccount, "InvalidAmount")
                .withArgs(chequeWithInvalidAmount.amount, cheque2.amount);

            // Cheque with invalid counter
            const chequeWithInvalidCounter = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: toCMAccountAddress,
                toBot: signers.otherAccount2.address,
                counter: cheque2.counter, // Same counter as cheque2
                amount: cheque2.amount, // Same amount is OK (for zero value cheque from zero fee services)
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
                paymentToken: await nullUSD.getAddress(),
            };

            // Sign Cheque
            const signatureWithInvalidCounter = await signMessengerCheque(
                chequeWithInvalidCounter,
                signers.chequeOperator,
            );

            // Try to cash-in cheque with invalid counter
            await expect(
                cmAccount.cashInCheque(
                    chequeWithInvalidCounter.fromCMAccount,
                    chequeWithInvalidCounter.toCMAccount,
                    chequeWithInvalidCounter.toBot,
                    chequeWithInvalidCounter.counter,
                    chequeWithInvalidCounter.amount,
                    chequeWithInvalidCounter.createdAt,
                    chequeWithInvalidCounter.expiresAt,
                    chequeWithInvalidCounter.paymentToken,
                    signatureWithInvalidCounter,
                ),
            )
                .to.be.revertedWithCustomError(cmAccount, "InvalidCounter")
                .withArgs(chequeWithInvalidCounter.counter, cheque2.counter);

            // CHECK INVALID AMOUNT AND COUNTER ----------------------------------------------------

            // Consume all prefund: create cheque3 with 100 CAM
            const cheque3 = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: toCMAccountAddress,
                toBot: signers.otherAccount2.address,
                counter: cheque2.counter + 1,
                amount: ethers.parseEther("100"),
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
                paymentToken: await nullUSD.getAddress(),
            };

            // Sign Cheque
            const signature3 = await signMessengerCheque(cheque3, signers.chequeOperator);

            // Cash-in cheque
            const cashInResponse3 = await cmAccount.cashInCheque(
                cheque3.fromCMAccount,
                cheque3.toCMAccount,
                cheque3.toBot,
                cheque3.counter,
                cheque3.amount,
                cheque3.createdAt,
                cheque3.expiresAt,
                cheque3.paymentToken,
                signature3,
            );

            await expect(cashInResponse3).to.be.not.reverted;

            // Try withdraw
            const withdrawAmount = ethers.parseEther("0.1");
            const withdrawer = signers.withdrawer;
            const withdrawTx = cmAccount.connect(withdrawer).withdraw(withdrawer.address, withdrawAmount);
            await expect(withdrawTx).to.be.not.reverted;

            // Check balances
            await expect(withdrawTx).to.changeEtherBalances([cmAccount, withdrawer], [-withdrawAmount, withdrawAmount]);
        });

        it("Should not update total cheque payments for same account", async function () {
            const { cmAccount, cmAccountManager, prefundAmount, nullUSD, nullUSDDecimals } = await loadFixture(
                deployCMAccountWithDepositFixture,
            );

            // Define cheque
            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: await cmAccount.getAddress(),
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("0.1"),
                createdAt: ethers.toBigInt(Math.floor(Date.now() / 1000)),
                expiresAt: ethers.toBigInt(Math.floor(Date.now() / 1000)) + 300n,
                paymentToken: await nullUSD.getAddress(),
            };

            // Grant CHEQUE_OPERATOR_ROLE
            await cmAccount
                .connect(signers.cmAccountAdmin)
                .grantRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), signers.chequeOperator.address);

            // Sign Cheque
            const signature = await signMessengerCheque(cheque, signers.chequeOperator);

            // Initial total cheque payments should be zero
            expect(await cmAccount.getTotalChequePaymentsPerToken(cheque.paymentToken)).to.be.equal(0n);

            // Cash-in cheque
            const cashInResponse = await cmAccount.cashInCheque(
                cheque.fromCMAccount,
                cheque.toCMAccount,
                cheque.toBot,
                cheque.counter,
                cheque.amount,
                cheque.createdAt,
                cheque.expiresAt,
                cheque.paymentToken,
                signature,
            );
            await expect(cashInResponse).to.be.not.reverted;

            // After cash-in total cheque payments should still be zero because the
            // cheque is from the same account (fromCMAccount === toCMAccount)
            expect(await cmAccount.getTotalChequePaymentsPerToken(cheque.paymentToken)).to.be.equal(0n);

            // Check legacy `getTotalChequePayments` function (for old CAM cheques)
            expect(await cmAccount.getTotalChequePayments()).to.be.equal(0n);
        });
    });
});
