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
    generateMessengerChequeTypeHash,
    generateTypedDataHash,
    generateMessengerChequeHash,
    generateDomainTypeHash,
    _generateDomainSeparator,
    getDomainSeparatorCamino,
    getDomainSeparatorColumbus,
    getDomainSeparatorKopernikus,
    getDomainSeparatorForChain,
    signMessengerCheque,
} = require("../utils/cheques.js");
const { generateDomainSeparator } = require("../utils/cheques.js");

describe("ChequeManager", function () {
    describe("Main", function () {
        it("Should initialize the DOMAIN_SEPARATOR correctly", async function () {
            const { cmAccount } = await loadFixture(deployCMAccountWithDepositFixture);

            const domainSeparator = await cmAccount.DOMAIN_SEPARATOR();
            expect(domainSeparator).to.be.a("string");
        });

        it("Should hash MessengerCheque correctly", async function () {
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
            console.log("Cheque Hash From Contract:", hash);
            expect(hash).to.not.be.reverted;

            const typedDataFromContract = await cmAccount.hashTypedDataV4(cheque);
            console.log("Typed Data Hash From Contract:", typedDataFromContract);
            expect(typedDataFromContract).to.be.not.reverted;

            console.log(cheque);

            // Sign Cheque
            const signature = await signMessengerCheque(
                cheque.fromCMAccount,
                cheque.toCMAccount,
                cheque.toBot,
                cheque.counter,
                cheque.amount,
                cheque.timestamp,
                signers.otherAccount1,
            );
            console.log("Signature From Utils:", signature);

            // Give cheque operator role to signer
            expect(
                await cmAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(await cmAccount.CHEQUE_OPERATOR_ROLE(), signers.otherAccount1.address),
            ).to.not.be.reverted;

            // Chain ID
            const chainId = await ethers.provider.getNetwork().then((n) => n.chainId);

            // Domain separator
            const domainSeparator = getDomainSeparatorForChain(chainId);

            // Cheque hash
            const chequeHash = generateMessengerChequeHash(cheque);
            console.log("Cheque Hash From Utils:", chequeHash);

            // Typed Data Hash
            const typedDataHash = generateTypedDataHash(cheque, domainSeparator);
            console.log("Typed Data Hash From Utils:", typedDataHash);

            // Verify Cheque
            const verifyResult = cmAccount.verifyCheque(cheque, signature);

            await expect(await verifyResult).to.be.emit(cmAccount, "ChequeVerified");
            // .withArgs(cheque, signers.otherAccount1.address, ethers.parseEther("1"));

            const tx = await cmAccount.verifyCheque(cheque, signature);

            const receipt = await tx.wait();
            console.log(receipt);

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

            console.log(parsedEvent);

            //const { 0: fromBot, 1: paymentAmount } = verifyResult;
            // console.log("Signer From Contract:", fromBot);
            // console.log("Payment Amount From Contract:", paymentAmount);
        });

        // it("Should verify a valid cheque", async function () {
        //     const cheque = {
        //         fromCMAccount: cmAccount.address,
        //         toCMAccount: addr1.address,
        //         toBot: addr2.address,
        //         counter: 1,
        //         amount: ethers.utils.parseEther("1"),
        //         timestamp: Math.floor(Date.now() / 1000),
        //     };

        //     const domainSeparator = await cmAccount.DOMAIN_SEPARATOR();
        //     const typedData = {
        //         types: {
        //             EIP712Domain: [
        //                 { name: "name", type: "string" },
        //                 { name: "version", type: "string" },
        //                 { name: "chainId", type: "uint256" },
        //                 { name: "verifyingContract", type: "address" },
        //             ],
        //             MessengerCheque: [
        //                 { name: "fromCMAccount", type: "address" },
        //                 { name: "toCMAccount", type: "address" },
        //                 { name: "toBot", type: "address" },
        //                 { name: "counter", type: "uint256" },
        //                 { name: "amount", type: "uint256" },
        //                 { name: "timestamp", type: "uint256" },
        //             ],
        //         },
        //         primaryType: "MessengerCheque",
        //         domain: {
        //             name: "CaminoMessenger",
        //             version: "1",
        //             chainId: await cmAccount.provider.getNetwork().then((n) => n.chainId),
        //             verifyingContract: cmAccount.address,
        //         },
        //         message: cheque,
        //     };

        //     const signature = await owner._signTypedData(
        //         typedData.domain,
        //         { MessengerCheque: typedData.types.MessengerCheque },
        //         cheque,
        //     );
        //     const [signer, paymentAmount] = await cmAccount.verifyCheque(cheque, signature);

        //     expect(signer).to.equal(owner.address);
        //     expect(paymentAmount).to.equal(ethers.utils.parseEther("1"));
        // });

        // it("Should revert invalid cheque with incorrect counter", async function () {
        //     const cheque = {
        //         fromCMAccount: cmAccount.address,
        //         toCMAccount: addr1.address,
        //         toBot: addr2.address,
        //         counter: 0, // Invalid counter
        //         amount: ethers.utils.parseEther("1"),
        //         timestamp: Math.floor(Date.now() / 1000),
        //     };

        //     const domainSeparator = await cmAccount.DOMAIN_SEPARATOR();
        //     const typedData = {
        //         types: {
        //             EIP712Domain: [
        //                 { name: "name", type: "string" },
        //                 { name: "version", type: "string" },
        //                 { name: "chainId", type: "uint256" },
        //                 { name: "verifyingContract", type: "address" },
        //             ],
        //             MessengerCheque: [
        //                 { name: "fromCMAccount", type: "address" },
        //                 { name: "toCMAccount", type: "address" },
        //                 { name: "toBot", type: "address" },
        //                 { name: "counter", type: "uint256" },
        //                 { name: "amount", type: "uint256" },
        //                 { name: "timestamp", type: "uint256" },
        //             ],
        //         },
        //         primaryType: "MessengerCheque",
        //         domain: {
        //             name: "CaminoMessenger",
        //             version: "1",
        //             chainId: await cmAccount.provider.getNetwork().then((n) => n.chainId),
        //             verifyingContract: cmAccount.address,
        //         },
        //         message: cheque,
        //     };

        //     const signature = await owner._signTypedData(
        //         typedData.domain,
        //         { MessengerCheque: typedData.types.MessengerCheque },
        //         cheque,
        //     );

        //     await expect(cmAccount.verifyCheque(cheque, signature)).to.be.revertedWith("InvalidCounter(0, 0)");
        // });

        // it("Should emit ChequeCashedIn event when cashing in a cheque", async function () {
        //     const cheque = {
        //         fromCMAccount: cmAccount.address,
        //         toCMAccount: addr1.address,
        //         toBot: addr2.address,
        //         counter: 1,
        //         amount: ethers.utils.parseEther("1"),
        //         timestamp: Math.floor(Date.now() / 1000),
        //     };

        //     const domainSeparator = await cmAccount.DOMAIN_SEPARATOR();
        //     const typedData = {
        //         types: {
        //             EIP712Domain: [
        //                 { name: "name", type: "string" },
        //                 { name: "version", type: "string" },
        //                 { name: "chainId", type: "uint256" },
        //                 { name: "verifyingContract", type: "address" },
        //             ],
        //             MessengerCheque: [
        //                 { name: "fromCMAccount", type: "address" },
        //                 { name: "toCMAccount", type: "address" },
        //                 { name: "toBot", type: "address" },
        //                 { name: "counter", type: "uint256" },
        //                 { name: "amount", type: "uint256" },
        //                 { name: "timestamp", type: "uint256" },
        //             ],
        //         },
        //         primaryType: "MessengerCheque",
        //         domain: {
        //             name: "CaminoMessenger",
        //             version: "1",
        //             chainId: await cmAccount.provider.getNetwork().then((n) => n.chainId),
        //             verifyingContract: cmAccount.address,
        //         },
        //         message: cheque,
        //     };

        //     const signature = await owner._signTypedData(
        //         typedData.domain,
        //         { MessengerCheque: typedData.types.MessengerCheque },
        //         cheque,
        //     );

        //     await expect(cmAccount.cashInCheque(cheque, signature))
        //         .to.emit(cmAccount, "ChequeCashedIn")
        //         .withArgs(owner.address, addr2.address, 1, ethers.utils.parseEther("1"), 0);
        // });

        // it("Should record a debt when cheque bounces", async function () {
        //     await cmAccount.recordDebt(addr1.address, ethers.utils.parseEther("1"));
        //     const lockedAmount = await cmAccount.getLockedAmount(addr1.address);
        //     expect(lockedAmount).to.equal(ethers.utils.parseEther("1"));
        // });

        // it("Should clear a debt when paid", async function () {
        //     await cmAccount.recordDebt(addr1.address, ethers.utils.parseEther("1"));
        //     await cmAccount.payDebt(addr1.address, ethers.utils.parseEther("1"));
        //     const lockedAmount = await cmAccount.getLockedAmount(addr1.address);
        //     expect(lockedAmount).to.equal(ethers.utils.parseEther("0"));
        // });

        // it("Should prevent withdrawal of locked amounts", async function () {
        //     const depositAmount = ethers.utils.parseEther("2");

        //     await cmAccount.connect(addr1).deposit({ value: depositAmount });
        //     await cmAccount.recordDebt(addr1.address, ethers.utils.parseEther("1"));

        //     await expect(
        //         cmAccount.connect(addr1).withdraw(addr1.address, ethers.utils.parseEther("2")),
        //     ).to.be.revertedWith("Insufficient unlocked funds");

        //     await cmAccount.connect(addr1).withdraw(addr1.address, ethers.utils.parseEther("1"));
        //     const lockedAmount = await cmAccount.getLockedAmount(addr1.address);
        //     expect(lockedAmount).to.equal(ethers.utils.parseEther("1"));
        // });
    });
});
