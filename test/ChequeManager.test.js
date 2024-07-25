/**
 * @dev Cheques
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
        //it("Should initialize the DOMAIN_SEPARATOR correctly", async function () {});
        it("Should initialize the DOMAIN_SEPARATOR correctly", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            const chainId = await ethers.provider.getNetwork().then((n) => n.chainId);
            const calculatedDomainSeparator = calculateDomainSeparatorForChain(chainId);

            const cmAccountDomainSeparator = await cmAccount.DOMAIN_SEPARATOR();
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
                timestamp: 1721777321,
            };

            const calculatedHash = calculateMessengerChequeHash(cheque);

            const hashFromContract = await cmAccount.hashMessengerCheque(cheque);

            expect(hashFromContract).to.be.equal(calculatedHash);
        });

        /**
         * Test hashing TypedData
         * @dev
         */
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
                timestamp: 1721777321,
            };

            // Calculate domain separator
            const chainId = await ethers.provider.getNetwork().then((n) => n.chainId);
            const calculatedDomainSeparator = calculateDomainSeparatorForChain(chainId);

            // Calculate typedDataHash
            const calculatedTypedDataHash = calculateTypedDataHash(cheque, calculatedDomainSeparator);

            // Get typedDataHash from contract
            const typedDataHashFromContract = await cmAccount.hashTypedDataV4(cheque);

            // Assert that the calculated typedDataHash is equal to the typedDataHash from contract
            expect(typedDataHashFromContract).to.be.equal(calculatedTypedDataHash);
        });

        it("REMOVE Should hash MessengerCheque correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);
            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: signers.otherAccount1.address,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("1"),
                timestamp: 1721777321,
            };
            const hash = await cmAccount.hashMessengerCheque(cheque);
            // console.log("Cheque Hash From Contract:", hash);
            expect(hash).to.not.be.reverted;

            const typedDataFromContract = await cmAccount.hashTypedDataV4(cheque);
            // console.log("Typed Data Hash From Contract:", typedDataFromContract);
            expect(typedDataFromContract).to.be.not.reverted;

            // console.log(cheque);

            // Sign Cheque
            const signature = await _signMessengerCheque(
                cheque.fromCMAccount,
                cheque.toCMAccount,
                cheque.toBot,
                cheque.counter,
                cheque.amount,
                cheque.timestamp,
                signers.otherAccount1,
            );
            // console.log("Signature From Utils:", signature);

            // Give cheque operator role to signer
            expect(
                await cmAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), signers.otherAccount1.address),
            ).to.not.be.reverted;

            // Chain ID
            const chainId = await ethers.provider.getNetwork().then((n) => n.chainId);

            // Domain separator
            const domainSeparator = calculateDomainSeparatorForChain(chainId);

            // Cheque hash
            const chequeHash = calculateMessengerChequeHash(cheque);
            // console.log("Cheque Hash From Utils:", chequeHash);

            // Typed Data Hash
            const typedDataHash = calculateTypedDataHash(cheque, domainSeparator);
            // console.log("Typed Data Hash From Utils:", typedDataHash);

            // Verify Cheque
            const verifyResult = cmAccount.verifyCheque(cheque, signature);

            await expect(await verifyResult).to.be.emit(cmAccount, "ChequeVerified");
            // .withArgs(cheque, signers.otherAccount1.address, ethers.parseEther("1"));

            const tx = await cmAccount.verifyCheque(cheque, signature);

            const receipt = await tx.wait();

            // Parse event to get the CMAccount address (this is the UUPS proxy address)
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccount.interface.parseLog(log).name === "ChequeVerified";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = cmAccount.interface.parseLog(event);
            const chequeFromEvent = parsedEvent.args.cheque;
            const fromBot = parsedEvent.args.signer;
            const paymentAmount = parsedEvent.args.paymentAmount;

            // console.log(parsedEvent);

            //const { 0: fromBot, 1: paymentAmount } = verifyResult;
            // console.log("Signer From Contract:", fromBot);
            // console.log("Payment Amount From Contract:", paymentAmount);
        });
    });
    describe("Cheque Operations", function () {
        it("Should verify a cheque with a valid signature", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: signers.chequeOperator.address,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("1"),
                timestamp: 1721777321,
            };

            // Grant CHEQUE_OPERATOR_ROLE
            await cmAccount
                .connect(signers.cmAccountAdmin)
                .grantRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), signers.chequeOperator.address);

            // Sign Cheque
            const signature = await signMessengerCheque(cheque, signers.chequeOperator);

            // Verify cheque
            const verifyResponse = await cmAccount.verifyCheque(cheque, signature);

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
                    cheque.amount,
                );

            // Sanity checks: balances should not change
            await expect(await verifyResponse).to.changeEtherBalance(signers.chequeOperator, 0, { includeFee: true });
            await expect(await verifyResponse).to.changeEtherBalance(cmAccount, 0, { includeFee: true });
        });

        it("Should not verify a cheque with an invalid signature", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: signers.chequeOperator.address,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("1"),
                timestamp: 1721777321,
            };

            // Grant CHEQUE_OPERATOR_ROLE
            await cmAccount
                .connect(signers.cmAccountAdmin)
                .grantRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), signers.chequeOperator.address);

            // Sign invalid cheque. Function below changes the chainId to invalidate the cheque.
            const signature = await signInvalidMessengerCheque(cheque, signers.chequeOperator);

            // Verify cheque, should revert and have the wrong address in the event
            await expect(cmAccount.verifyCheque(cheque, signature))
                .to.be.revertedWithCustomError(cmAccount, "NotAllowedToSignCheques")
                .withArgs((addr) => addr !== signers.chequeOperator.address);
        });
        it("Should cash-in multiple cheques correctly", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            const cheque = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: signers.chequeOperator.address,
                toBot: signers.otherAccount2.address,
                counter: 1,
                amount: ethers.parseEther("0.1"),
                timestamp: 1721777321,
            };

            // Grant CHEQUE_OPERATOR_ROLE
            await cmAccount
                .connect(signers.cmAccountAdmin)
                .grantRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), signers.chequeOperator.address);

            // Sign Cheque
            const signature = await signMessengerCheque(cheque, signers.chequeOperator);

            const developerFee = await cmAccount.calculateDeveloperFee(cheque.amount);

            // Cash-in cheque
            const cashInResponse = await cmAccount.cashInCheque(cheque, signature);

            // Check balances
            // CMAccount balance descrease by cheque amount + developerFee
            await expect(await cashInResponse).to.changeEtherBalance(cmAccount, -cheque.amount - developerFee);
            // toCMAccount balance increase by cheque amount, we are using regular wallet here instead of another CMAccount
            // TODO: Use a real CMAccount as a receiver (CMAccount CAM receive not implemented yet)
            await expect(await cashInResponse).to.changeEtherBalance(signers.chequeOperator, cheque.amount);
            // DeveloperWallet balance increase by developerFee
            await expect(await cashInResponse).to.changeEtherBalance(signers.developerWallet, developerFee);

            // Should emit event with correct data
            await expect(await cashInResponse)
                .to.emit(cmAccount, "ChequeCashedIn")
                .withArgs(signers.chequeOperator.address, cheque.toBot, cheque.counter, cheque.amount, developerFee);

            // Sanity checks: should set lastCashIns
            const lastCashIn = await cmAccount.getLastCashIn(signers.chequeOperator, cheque.toBot);
            expect(lastCashIn).to.be.deep.equal([cheque.counter, cheque.amount]);

            /**
             * Second cheque
             */

            // New cheque with a higher counter and amount
            const cheque2 = {
                fromCMAccount: await cmAccount.getAddress(),
                toCMAccount: signers.chequeOperator.address,
                toBot: signers.otherAccount2.address,
                counter: 100,
                amount: ethers.parseEther("0.234"),
                timestamp: 1721777322,
            };
            234 - 100;
            // Sign Cheque
            const signature2 = await signMessengerCheque(cheque2, signers.chequeOperator);

            const developerFee2 = await cmAccount.calculateDeveloperFee(cheque2.amount - cheque.amount);

            // Cash-in cheque
            const cashInResponse2 = await cmAccount.cashInCheque(cheque2, signature2);

            // Check balances
            // CMAccount balance descrease by (cheque2 amount - cheque amount) + developerFee
            await expect(await cashInResponse2).to.changeEtherBalance(
                cmAccount,
                -cheque2.amount + cheque.amount - developerFee2, // Weird calculation but it works
            );
            // toCMAccount balance increase by (cheque2 amount - cheque amount)
            await expect(await cashInResponse2).to.changeEtherBalance(
                signers.chequeOperator,
                cheque2.amount - cheque.amount, // new cheque amount minus the lastCashIn amount
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
                    cheque2.amount,
                    developerFee2,
                );
        });
    });
});
