const { ethers } = require("hardhat");

/**
 * You need to run a local node to use this example.
 *
 * Run a local hardhat node:
 *   yarn hardhat node
 *
 * In another terminal, deploy the contracts:
 *   yarn hardhat ignition deploy ignition/modules/0_development.js --network localhost
 *
 * The command above should procude output like this:
 *
 *   Deployed Addresses
 *
 *   BookingTokenProxyModule#BookingToken - 0x5FbDB2315678afecb367f032d93F642f64180aa3
 *   CMAccountManagerModule#CMAccount - 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
 *   ManagerProxyModule#CMAccountManager - 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
 *   BookingTokenProxyModule#ERC1967Proxy - 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9
 *   ManagerProxyModule#ERC1967Proxy - 0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9
 *   CMAccountManagerModule#BookingToken - 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9
 *   CMAccountManagerModule#CMAccountManager - 0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9
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
 *   ----------------------- creating cheque and signatures (off-chain) -----------------------
 *   Signer: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
 *   Cheque: {
 *     fromCMAccount: '0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65',
 *     toCMAccount: '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC',
 *     toBot: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
 *     counter: 1,
 *     amount: 1000000000000000000n,
 *     createdAt: 1722565705,
 *     expiresAt: 1722565825
 *   }
 *
 *   ----------------------- Calculated Values -----------------------
 *   MESSENGER_CHEQUE_TYPEHASH: 0x87b38f131334165ac2b361f08966c9fcff3a953fa7d9d9c2861b7f0b50445bcb
 *   DOMAIN_TYPEHASH: 0xc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e
 *   DOMAIN_SEPARATOR: 0x792acc3adab7297918d2cdaeb59ac5f091943a65aba244c580164ec2ec307451
 *   Cheque Hash: 0x91add1c1535d9bc4cc3cd2d7fbef1b30e1b2d8dead0087d7c568c8dca1b63430
 *   Typed Data Hash: 0x208cfb408b5d2498fb4fa978535ab394be394918e3c17fd9ee9f3260b84056cc
 *   Signature: 0x10576abbb5e8c2a813fde2809c371cf70d682a79c8003e4c677d3c25187022410f5587dee3c798a9978adadf31bca336ffd158910b6df12ca8bba4fb30699caf1b
 *   Recovered Address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 (Should be same as the 'Signer')
 *
 *   ----------------------- Trying with a CMAccount (on-chain) -----------------------
 *   * Getting CMAccountManager and creating a CMAccount...
 *   * Created CMAccount at address: 0x856e4424f806D16E8CBC702B3c0F2ede5468eae5
 *   * Calling CMAccount.recoverSigner(cheque, signature)...
 *   Recovered (on-chain): 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 (Should be same as the 'Signer')
 *   Done in 1.66s.
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
    console.log("\n----------------------- Calculated Values (off-chain) -----------------------");
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
    console.log("Recovered Address:", recoveredAddress, "(Should be same as the 'Signer')");

    // Try to recover the same address from the CMAccount's `recoverSigner(cheque,
    // signature)` function.
    console.log("\n----------------------- Trying with a CMAccount (on-chain) -----------------------");

    // First we need to create a CMAccount. So, we get the contract at the address
    // below. Which is always the same address on a fresh hardhat local network
    // because the transaction creating it is the same. (same nonce, same owner
    // etc..)
    console.log("* Getting CMAccountManager and creating a CMAccount...");
    const cmAccountManagerAddress = "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707";
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

    console.log("* Calling CMAccount.recoverSigner...");

    // Call `recoverSigner` to verify if the recovered signer is the same as our
    // original signer's address

    // Update: We changed the recoverSigner function to internal. So it's not
    // visible anymore on the CMAccount contract. If you really want to test it
    // change the visibility to public, redeploy the contracts and enable the
    // following lines below. Then the run this script again.
    console.log("**** Visibility of recoverSigner is switched to internal. Read the comments in the file. ****");

    // const res = await cmAccount.recoverSigner(
    //     cheque.fromCMAccount,
    //     cheque.toCMAccount,
    //     cheque.toBot,
    //     cheque.counter,
    //     cheque.amount,
    //     cheque.createdAt,
    //     cheque.expiresAt,
    //     signature.serialized,
    // );
    // console.log("Recovered (on-chain):", res, "(Should be same as the 'Signer')");
}

main().catch(console.error);
