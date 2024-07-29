const DOMAIN_NAME = "CaminoMessenger";
const DOMAIN_VERSION = "1";

function calculateMessengerChequeTypeHash() {
    const typeHash = ethers.keccak256(
        ethers.toUtf8Bytes(
            "MessengerCheque(address fromCMAccount,address toCMAccount,address toBot,uint256 counter,uint256 amount,uint256 timestamp)",
        ),
    );
    return typeHash;
}

function calculateDomainTypeHash() {
    const domainTypeHash = ethers.keccak256(
        ethers.toUtf8Bytes("EIP712Domain(string name,string version,uint256 chainId)"),
    );
    return domainTypeHash;
}

function calculateDomainSeparator(domainName, domainVersion, chainId) {
    const coder = ethers.AbiCoder.defaultAbiCoder();
    const domainSeparator = ethers.keccak256(
        coder.encode(
            ["bytes32", "bytes32", "bytes32", "uint256"],
            [
                calculateDomainTypeHash(),
                ethers.keccak256(ethers.toUtf8Bytes(domainName)),
                ethers.keccak256(ethers.toUtf8Bytes(domainVersion)),
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
    return ethers.keccak256(ethers.concat([ethers.toUtf8Bytes("\x19\x01"), domainSeparator, chequeHash]));
}

function calculateMessengerChequeHash(cheque) {
    const chequeTypeHash = calculateMessengerChequeTypeHash();

    const coder = ethers.AbiCoder.defaultAbiCoder();
    const encodedCheque = coder.encode(
        ["bytes32", "address", "address", "address", "uint256", "uint256", "uint256"],
        [
            chequeTypeHash,
            cheque.fromCMAccount,
            cheque.toCMAccount,
            cheque.toBot,
            cheque.counter,
            cheque.amount,
            cheque.timestamp,
        ],
    );
    return ethers.keccak256(encodedCheque);
}

async function _signMessengerCheque(fromCMAccount, toCMAccount, toBot, counter, amount, timestamp, signer) {
    const chainId = await signer.provider.getNetwork().then((n) => n.chainId);

    const cheque = {
        fromCMAccount: fromCMAccount,
        toCMAccount: toCMAccount,
        toBot: toBot,
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
        chainId: chainId + 42n, // Invalid chainId
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
