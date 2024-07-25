const { keccak256, toUtf8Bytes, AbiCoder, getAddress, signTypedData } = require("ethers");

const DOMAIN_NAME = "CaminoMessenger";
const DOMAIN_VERSION = "1";

function calculateMessengerChequeTypeHash() {
    const typeHash = keccak256(
        toUtf8Bytes(
            "MessengerCheque(address fromCMAccount,address toCMAccount,address toBot,uint256 counter,uint256 amount,uint256 timestamp)",
        ),
    );
    return typeHash;
}

function calculateDomainTypeHash() {
    const domainTypeHash = keccak256(toUtf8Bytes("EIP712Domain(string name,string version,uint256 chainId)"));
    return domainTypeHash;
}

function calculateDomainSeparator(domainName, domainVersion, chainId) {
    const coder = AbiCoder.defaultAbiCoder();
    const domainSeparator = keccak256(
        coder.encode(
            ["bytes32", "bytes32", "bytes32", "uint256"],
            [
                calculateDomainTypeHash(),
                keccak256(toUtf8Bytes(domainName)),
                keccak256(toUtf8Bytes(domainVersion)),
                ethers.toBigInt(chainId),
            ],
        ),
    );
    return domainSeparator;
}

function calculateDomainSeparatorForChain(_chainId) {
    const domainName = "CaminoMessenger";
    const domainVersion = "1";
    const chainId = _chainId;
    return calculateDomainSeparator(domainName, domainVersion, chainId);
}

function calculateDomainSeparatorCamino() {
    return calculateDomainSeparatorForChain(500);
}

function calculateDomainSeparatorColumbus() {
    return calculateDomainSeparatorForChain(501);
}

function calculateDomainSeparatorKopernikus() {
    return calculateDomainSeparatorForChain(502);
}

function calculateTypedDataHash(cheque, domainSeparator) {
    const chequeHash = calculateMessengerChequeHash(cheque);
    return keccak256(ethers.concat([ethers.toUtf8Bytes("\x19\x01"), domainSeparator, chequeHash]));
}

function calculateMessengerChequeHash(cheque) {
    const chequeTypeHash = calculateMessengerChequeTypeHash();

    const coder = AbiCoder.defaultAbiCoder();
    const encodedCheque = coder.encode(
        ["bytes32", "address", "address", "address", "uint256", "uint256", "uint256"],
        [
            chequeTypeHash,
            getAddress(cheque.fromCMAccount),
            getAddress(cheque.toCMAccount),
            getAddress(cheque.toBot),
            cheque.counter,
            cheque.amount,
            cheque.timestamp,
        ],
    );
    return keccak256(encodedCheque);
}

async function _signMessengerCheque(fromCMAccount, toCMAccount, toBot, counter, amount, timestamp, signer) {
    const chainId = await signer.provider.getNetwork().then((n) => n.chainId);

    const cheque = {
        fromCMAccount: getAddress(fromCMAccount),
        toCMAccount: getAddress(toCMAccount),
        toBot: getAddress(toBot),
        counter: counter,
        amount: amount,
        timestamp: timestamp,
    };

    const signature = await signMessengerCheque(cheque, signer);

    return signature;
}

async function signMessengerCheque(cheque, signer) {
    const chainId = await signer.provider.getNetwork().then((n) => n.chainId);

    const types = {
        MessengerCheque: [
            { name: "fromCMAccount", type: "address" },
            { name: "toCMAccount", type: "address" },
            { name: "toBot", type: "address" },
            { name: "counter", type: "uint256" },
            { name: "amount", type: "uint256" },
            { name: "timestamp", type: "uint256" },
        ],
    };

    const domain = {
        name: DOMAIN_NAME,
        version: DOMAIN_VERSION,
        chainId: chainId,
    };

    const signature = await signer.signTypedData(domain, types, cheque);
    return signature;
}

async function signInvalidMessengerCheque(cheque, signer) {
    const chainId = await signer.provider.getNetwork().then((n) => n.chainId);

    const types = {
        MessengerCheque: [
            { name: "fromCMAccount", type: "address" },
            { name: "toCMAccount", type: "address" },
            { name: "toBot", type: "address" },
            { name: "counter", type: "uint256" },
            { name: "amount", type: "uint256" },
            { name: "timestamp", type: "uint256" },
        ],
    };

    const domain = {
        name: DOMAIN_NAME,
        version: DOMAIN_VERSION,
        chainId: chainId + 2n, // Invalid chainId
    };

    const signature = await signer.signTypedData(domain, types, cheque);
    return signature;
}

module.exports = {
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
};

//{generateMessengerChequeTypeHash, generateDomainTypeHash, _generateDomainSeparator, getDomainSeparatorCamino, getDomainSeparatorColumbus, getDomainSeparatorKopernikus, getDomainSeparatorForChain, signMessengerCheque,

// // Get receipt
// const tx = await verifyResponse;
// const receipt = await tx.wait();

// // Parse event
// const event = receipt.logs.find((log) => {
//     try {
//         return cmAccount.interface.parseLog(log).name === "ChequeVerified";
//     } catch (e) {
//         return false;
//     }
// });

// const parsedEvent = cmAccount.interface.parseLog(event);

// const fromCMAccount = parsedEvent.args.fromCMAccount;
// const toCMAccount = parsedEvent.args.toCMAccount;
// const fromBot = parsedEvent.args.fromBot;
// const toBot = parsedEvent.args.toBot;
// const counter = parsedEvent.args.counter;
// const amount = parsedEvent.args.amount;
// const payment = parsedEvent.args.payment;

// console.log("fromCMAccount:", fromCMAccount);
// console.log("toCMAccount:", toCMAccount);
// console.log("fromBot:", fromBot);
// console.log("toBot:", toBot);
// console.log("counter:", counter);
// console.log("amount:", amount);
// console.log("payment:", payment);
