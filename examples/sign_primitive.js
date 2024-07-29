const { getBytes } = require("ethers");
const { ethers } = require("hardhat");

/**
 * You need to run a local node to use this example.
 *
 * Run a local hardhat node:
 *  yarn hardhat node
 *
 * In another terminal, deploy the contracts:
 *  yarn hardhat ignition deploy ignition/modules/0_development.js --network localhost
 *
 * The command above should procude output like this:
 *
 *  CMAccountManagerModule#CMAccount - 0x5FbDB2315678afecb367f032d93F642f64180aa3
 *  ProxyModule#CMAccountManager - 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
 *  ProxyModule#CMAccountManagerProxy - 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
 *  CMAccountManagerModule#CMAccountManagerProxy - 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
 *
 * No you are ready to run this script:
 *  yarn hardhat run examples/sign.js --network localhost
 */

async function main() {
    console.log("\n----------------------- creating cheque and signatures (off-chain) -----------------------");

    // First signer from hardhat local:
    // address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
    const [wallet] = await ethers.getSigners();

    console.log("Signer:", wallet.address);

    // Create a sample cheque, addresses used here are just dummy account addresses
    // from hardhat local node
    const cheque = {
        fromCMAccount: "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65",
        toCMAccount: "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
        toBot: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
        counter: 1,
        amount: ethers.parseUnits("1.0", "ether"), // 1 ETH
        timestamp: Math.floor(Date.now() / 1000), // Current Unix timestamp
    };

    console.log("Cheque:", cheque);

    // Calculate the typehash
    const MESSENGER_CHEQUE_TYPEHASH = ethers.keccak256(
        ethers.toUtf8Bytes(
            "MessengerCheque(address fromCMAccount,address toCMAccount,address toBot,uint256 counter,uint256 amount,uint256 timestamp)",
        ),
    );

    // Calculate domain typehash
    const DOMAIN_TYPEHASH = ethers.keccak256(
        ethers.toUtf8Bytes("EIP712Domain(string name,string version,uint256 chainId)"),
    );

    // Get the chain id from the default provider
    const chainId = (await ethers.provider.getNetwork()).chainId;

    // Get the abi encoder
    const coder = ethers.AbiCoder.defaultAbiCoder();

    // Calculate the domain separator using domain typehash and correct values for
    // name, version, and chain id. (should be same with the CMAccount)
    const DOMAIN_SEPARATOR = ethers.keccak256(
        coder.encode(
            ["bytes32", "bytes32", "bytes32", "uint256"],
            [
                DOMAIN_TYPEHASH,
                ethers.keccak256(ethers.toUtf8Bytes("CaminoMessenger")),
                ethers.keccak256(ethers.toUtf8Bytes("1")),
                chainId,
            ],
        ),
    );

    // Print out what we've calculated so far
    console.log("MESSENGER_CHEQUE_TYPEHASH:", MESSENGER_CHEQUE_TYPEHASH);
    console.log("DOMAIN_TYPEHASH:", DOMAIN_TYPEHASH);
    console.log("DOMAIN_SEPARATOR:", DOMAIN_SEPARATOR);

    // Create cheque hash using the cheque's typehash
    const chequeHash = ethers.keccak256(
        coder.encode(
            ["bytes32", "address", "address", "address", "uint256", "uint256", "uint256"],
            [
                MESSENGER_CHEQUE_TYPEHASH,
                cheque.fromCMAccount,
                cheque.toCMAccount,
                cheque.toBot,
                cheque.counter,
                cheque.amount,
                cheque.timestamp,
            ],
        ),
    );

    // Calculate the typed data hash using domain separator and cheque hash
    const typedDataHash = ethers.keccak256(
        ethers.concat([ethers.toUtf8Bytes("\x19\x01"), DOMAIN_SEPARATOR, chequeHash]),
    );

    // Print them out to the console
    console.log("Cheque Hash:", chequeHash);
    console.log("Typed Data Hash:", typedDataHash);

    // We are not using the `.signMessage` and `.verifyMessage` as these are
    // prefixing the message with "\x19Ethereum Signed Message:\n". See:
    // [EIP-191](https://eips.ethereum.org/EIPS/eip-191)
    //
    //const signature = await wallet.signMessage(ethers.getBytes(typedDataHash));
    //const recoveredAddress = ethers.verifyMessage(ethers.getBytes(typedDataHash), signature);

    // First signer from hardhat local:
    // address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
    // private key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

    // Get a signing key from ethers
    const signingKey = new ethers.SigningKey("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80");
    // Signed the typed data hash
    const signature = signingKey.sign(ethers.getBytes(typedDataHash));

    // Print the serialized signature
    console.log("Signature:", signature.serialized);

    // Recover the address from the digest and signature
    const recoveredAddress = ethers.recoverAddress(getBytes(typedDataHash), signature);
    console.log("Recovered Address:", recoveredAddress, "(Should be same as the 'Signer')");

    // Try to recover the same address from the CMAccount's `recoverSigner(cheque,
    // signature)` function.
    console.log("\n----------------------- trying with a CMAccount (on-chain) -----------------------");

    // First we need to create a CMAccount. So, we get the contract at the address
    // below. Which is always the same address on a fresh hardhat local network
    // because the transaction creating it is the same. (same nonce, same owner
    // etc..)
    console.log("* Getting CMAccountManager and creating a CMAccount...");
    const cmAccountManagerAddress = "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0";
    const manager = await ethers.getContractAt("CMAccountManager", cmAccountManagerAddress);

    // Create a CMAccount
    const tx = await manager.createCMAccount(wallet.address, wallet.address, wallet.address);

    // Get the tx receipt
    const receipt = await tx.wait();

    // Parse event to get the CMAccount address
    const event = receipt.logs.find((log) => {
        try {
            return manager.interface.parseLog(log).name === "CMAccountCreated";
        } catch (e) {
            return false;
        }
    });

    const parsedEvent = manager.interface.parseLog(event);
    const cmAccountAddress = parsedEvent.args.account;

    console.log("* Created CMAccount at address:", cmAccountAddress);

    // Get newly created CM Account contract using the parsed address
    const cmAccount = await ethers.getContractAt("CMAccount", cmAccountAddress);

    console.log("* Calling CMAccount.recoverSigner(cheque, signature)...");

    // Call `recoverSigner` to verify if the recovered signer is the same as our
    // original signer's address
    const res = await cmAccount.recoverSigner(cheque, signature.serialized);
    console.log("Recovered (on-chain):", res, "(Should be same as the 'Signer')");
}

main().catch(console.error);
