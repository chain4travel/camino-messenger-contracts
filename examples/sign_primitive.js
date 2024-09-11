const { ethers } = require("hardhat");

/**
 * You need to run a local node to use this example.
 *
 * Run a local hardhat node:
 *   yarn hardhat node
 *
 * In another terminal, deploy the contracts:
 *   yarn hardhat ignition deploy ignition/modules/messenger.js --network localhost
 *
 * The command above should produce output like this:
 *
 *   Deployed Addresses
 *
 *   CaminoMessengerModule#BookingToken - 0x5FbDB2315678afecb367f032d93F642f64180aa3
 *   CaminoMessengerModule#BookingTokenOperator - 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
 *   CaminoMessengerModule#CMAccountManager - 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
 *   CaminoMessengerModule#CMAccount - 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9
 *   CaminoMessengerModule#ManagerERC1967Proxy - 0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9
 *   CaminoMessengerModule#ManagerProxy - 0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9
 *   CaminoMessengerModule#BookingTokenERC1967Proxy - 0x5FC8d32690cc91D4c39d9d3abcBD16989F875707
 *   CaminoMessengerModule#BookingTokenProxy - 0x5FC8d32690cc91D4c39d9d3abcBD16989F875707
 *
 * Now you are ready to run this script:
 *   yarn hardhat run examples/sign_primitive.js --network localhost
 *
 * Example Output:
 *
 *   ❯ yarn hardhat run examples/sign_primitive.js --network localhost
 *   yarn run v1.22.19
 *   $ /hgst/work/github.com/chain4travel/camino-messenger-contracts/node_modules/.bin/hardhat run examples/sign_primitive.js --network localhost
 *
 *   ----------------------- Prepare CMAccount ------------------------------------------------
 *   🔑 Signer: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
 *   * Getting CMAccountManager and creating a CMAccount...
 *   * Created CMAccount at address: 0x8271373aC5cD66E9e7dC752c0e39da5d12988Ec6
 *   Registering address as a bot (Cheque signer)...
 *   Done!
 *
 *   ----------------------- creating cheque and signatures (off-chain) -----------------------
 *   Cheque: {
 *     fromCMAccount: '0x8271373aC5cD66E9e7dC752c0e39da5d12988Ec6',
 *     toCMAccount: '0x8271373aC5cD66E9e7dC752c0e39da5d12988Ec6',
 *     toBot: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
 *     counter: 1n,
 *     amount: 1000000000000000000n,
 *     createdAt: 1726063332,
 *     expiresAt: 1726063452
 *   }
 *
 *   ----------------------- Calculated Values (off-chain) ------------------------------------
 *   MESSENGER_CHEQUE_TYPEHASH: 0x87b38f131334165ac2b361f08966c9fcff3a953fa7d9d9c2861b7f0b50445bcb
 *   DOMAIN_TYPEHASH: 0xc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e
 *   DOMAIN_SEPARATOR: 0x792acc3adab7297918d2cdaeb59ac5f091943a65aba244c580164ec2ec307451
 *   Cheque Hash: 0xaf8cf3a3b4742c5a8ab6ec5d5d14e48237bde9d0111af466f1babbf79d9fe574
 *   Typed Data Hash: 0x3a6686e6234a0effbdbb5d9a754dee7e57b2f89f7d3a600b9abd9078a0136cd0
 *   Signature: 0xaef61cbfd3dbd39f1184a425113707cd3ac2a732b5e379d19788e9d174aca8c638a56ecad452bbfedd1a2ccd0c79ce0f942d05ecd33c87af25e239abd33d2ad91c
 *   🔑 Recovered Signer: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 <== Should be same as the 'Signer'
 *
 *   ----------------------- Trying with a CMAccount (on-chain) -------------------------------
 *   * Calling CMAccount.verifyCheque...
 *   🔑 Recovered (on-chain): 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 <== Should be same as the off-chain 'Signer' above
 *   Done in 1.94s.
 */

async function main() {
    console.log("\n----------------------- Prepare CMAccount ------------------------------------------------");
    // First signer from hardhat local:
    // address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
    const [wallet] = await ethers.getSigners();

    console.log("🔑 Signer:", wallet.address);
    // First we need to create a CMAccount. So, we get the contract at the address
    // below. Which is always the same address on a fresh hardhat local network
    // because the transaction creating it is the same. (same nonce, same owner
    // etc..)
    console.log("* Getting CMAccountManager and creating a CMAccount...");
    const cmAccountManagerAddress = "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9";
    const manager = await ethers.getContractAt("CMAccountManager", cmAccountManagerAddress);

    // Create a CMAccount
    const tx = await manager.createCMAccount(wallet.address, wallet.address, {
        value: ethers.parseEther("200"),
    });

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

    console.log("Registering address as a bot (Cheque signer)...");

    // Register address as a bot
    const tx2 = await cmAccount.addMessengerBot(wallet.address, 0);
    const receipt2 = await tx2.wait();
    console.log("Done!");

    console.log("\n----------------------- creating cheque and signatures (off-chain) -----------------------");

    // Create a sample cheque, CM account addresses used here are same as above. Bot address is just dummy address.
    const cheque = {
        fromCMAccount: cmAccountAddress,
        toCMAccount: cmAccountAddress,
        toBot: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
        counter: 1n,
        amount: ethers.parseUnits("1.0", "ether"), // 1 ETH
        createdAt: Math.floor(Date.now() / 1000), // Unix timestamp
        expiresAt: Math.floor(Date.now() / 1000) + 120, // Unix timestamp
    };

    console.log("Cheque:", cheque);

    // Calculate the typehash
    const MESSENGER_CHEQUE_TYPEHASH = ethers.keccak256(
        ethers.toUtf8Bytes(
            "MessengerCheque(address fromCMAccount,address toCMAccount,address toBot,uint256 counter,uint256 amount,uint256 createdAt,uint256 expiresAt)",
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
    console.log("\n----------------------- Calculated Values (off-chain) ------------------------------------");
    console.log("MESSENGER_CHEQUE_TYPEHASH:", MESSENGER_CHEQUE_TYPEHASH);
    console.log("DOMAIN_TYPEHASH:", DOMAIN_TYPEHASH);
    console.log("DOMAIN_SEPARATOR:", DOMAIN_SEPARATOR);

    // Create cheque hash using the cheque's typehash
    const chequeHash = ethers.keccak256(
        coder.encode(
            ["bytes32", "address", "address", "address", "uint256", "uint256", "uint256", "uint256"],
            [
                MESSENGER_CHEQUE_TYPEHASH,
                cheque.fromCMAccount,
                cheque.toCMAccount,
                cheque.toBot,
                cheque.counter,
                cheque.amount,
                cheque.createdAt,
                cheque.expiresAt,
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
    const recoveredAddress = ethers.recoverAddress(ethers.getBytes(typedDataHash), signature);
    console.log("🔑 Recovered Signer:", recoveredAddress, "<== Should be same as the 'Signer'");

    // Try to recover the same address from the CMAccount's `verifyCheque(cheque ..., signature)` function.
    console.log("\n----------------------- Trying with a CMAccount (on-chain) -------------------------------");

    console.log("* Calling CMAccount.verifyCheque...");

    // Call `verifyCheque` to verify if the recovered signer is the same as our
    // original signer's address
    const res = await cmAccount.verifyCheque(
        cheque.fromCMAccount,
        cheque.toCMAccount,
        cheque.toBot,
        cheque.counter,
        cheque.amount,
        cheque.createdAt,
        cheque.expiresAt,
        signature.serialized,
    );
    console.log("🔑 Recovered (on-chain):", res[0], "<== Should be same as the off-chain 'Signer' above");
}

main().catch(console.error);
