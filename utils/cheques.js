const { keccak256, toUtf8Bytes, AbiCoder, getAddress, signTypedData } = require("ethers");

const DOMAIN_NAME = "CaminoMessenger";
const DOMAIN_VERSION = "1";

function generateMessengerChequeTypeHash() {
    const typeHash = keccak256(
        toUtf8Bytes(
            "MessengerCheque(address fromCMAccount,address toCMAccount,address toBot,uint256 counter,uint256 amount,uint256 timestamp)",
        ),
    );
    return typeHash;
}

function generateDomainTypeHash() {
    const domainTypeHash = keccak256(toUtf8Bytes("EIP712Domain(string name,string version,uint256 chainId)"));
    return domainTypeHash;
}

function _generateDomainSeparator(domainName, domainVersion, chainId) {
    const coder = AbiCoder.defaultAbiCoder();
    const domainSeparator = keccak256(
        coder.encode(
            ["bytes32", "bytes32", "bytes32", "uint256"],
            [
                generateDomainTypeHash(),
                keccak256(toUtf8Bytes(domainName)),
                keccak256(toUtf8Bytes(domainVersion)),
                ethers.toBigInt(chainId),
            ],
        ),
    );
    return domainSeparator;
}

function getDomainSeparatorForChain(_chainId) {
    const domainName = "CaminoMessenger";
    const domainVersion = "1";
    const chainId = _chainId;
    return _generateDomainSeparator(domainName, domainVersion, chainId);
}

function getDomainSeparatorCamino() {
    return getDomainSeparatorForChain(500);
}

function getDomainSeparatorColumbus() {
    return getDomainSeparatorForChain(501);
}

function getDomainSeparatorKopernikus() {
    return getDomainSeparatorForChain(502);
}

function generateTypedDataHash(cheque, domainSeparator) {
    const chequeHash = generateMessengerChequeHash(cheque);
    return keccak256(ethers.concat([ethers.toUtf8Bytes("\x19\x01"), domainSeparator, chequeHash]));
}

function generateMessengerChequeHash(cheque) {
    const chequeTypeHash = generateMessengerChequeTypeHash();

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

async function signMessengerCheque(fromCMAccount, toCMAccount, toBot, counter, amount, timestamp, signer) {
    const chainId = await signer.provider.getNetwork().then((n) => n.chainId);
    const domainSeparator = getDomainSeparatorForChain(chainId);

    const cheque = {
        fromCMAccount: getAddress(fromCMAccount),
        toCMAccount: getAddress(toCMAccount),
        toBot: getAddress(toBot),
        counter: counter,
        amount: amount,
        timestamp: timestamp,
    };

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

module.exports = {
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
};

//{generateMessengerChequeTypeHash, generateDomainTypeHash, _generateDomainSeparator, getDomainSeparatorCamino, getDomainSeparatorColumbus, getDomainSeparatorKopernikus, getDomainSeparatorForChain, signMessengerCheque,
