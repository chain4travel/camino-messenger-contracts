# Camino Messenger Contracts

[![CAMINO NETWORK](https://img.shields.io/badge/CAMINO-NETWORK-b440fc?style=for-the-badge&logoColor=white&labelColor=0085ff)](https://camino.network/) [![CHAT WITH US](https://img.shields.io/badge/DISCORD-5865F2?style=for-the-badge&logo=discord&logoColor=white)](https://discord.com/channels/949247897688494150/1182680860797960253)

[![CI](https://github.com/chain4travel/camino-messenger-contracts/actions/workflows/ci.yaml/badge.svg)](https://github.com/chain4travel/camino-messenger-contracts/actions/workflows/ci.yaml)

This repository contains the smart contracts for the [Camino
Messenger](https://camino.network/camino-messenger-sets-the-global-standard-in-travel-data-management-and-distribution/).

> [!WARNING]
> These contracts are currently in the development phase. The ABI is
> subject to change frequently until they are released into production.

## Columbus Deployed Contracts

Below is a table of deployed contracts and their addresses on Columbus testnet.

| Contract                   | Address                                                                                                                                       |
| -------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| CMAccountManager           | [0xE5B2f76C778D082b07BDd7D51FFe83E3E055B47F](https://columbus.caminoscan.com/address/0xE5B2f76C778D082b07BDd7D51FFe83E3E055B47F?tab=contract) |
| BookingToken               | [0xe55E387F5474a012D1b048155E25ea78C7DBfBBC](https://columbus.caminoscan.com/address/0xe55E387F5474a012D1b048155E25ea78C7DBfBBC?tab=contract) |
| CMAccount (Implementation) | [0xCB43DabB6d579d5a6067aB06575601D3D77882CA](https://columbus.caminoscan.com/address/0xCB43DabB6d579d5a6067aB06575601D3D77882CA?tab=contract) |

## Quickstart

### Clone the repo and change directory into

```sh
git clone git@github.com:chain4travel/camino-messenger-contracts.git
cd camino-messenger-contracts
```

### Install packages

```sh
yarn install
```

### Run tests. This will compile the contracts and run the tests:

```sh
yarn test
```

## Contracts

### CMAccount

The `CMAccount` contract represents a Camino Messenger account. Currently, it
includes functionalities for the management of bots. More features will be
introduced in the future.

This contract works closely with the `CMAccountManager` to handle accounts.

### CMAccountManager

The `CMAccountManager` contract acts as a manager for `CMAccount` contracts. It
handles the creation, registration, verification, and management of accounts. It
also keeps records for the developer wallet, fees, and `CMAccount` implementation
address. Accounts can only be upgraded to the implementation address that the
manager holds.

### ChequeManager

The `ChequeManager` contract handles the processing of cheques. It verifies the
signatures, checks the validity of the cheques, and transfers the funds between
accounts. It also calculates the developer fee and transfers it to the developer's
wallet.

This is a base contract that is inherited by the `CMAccount` contract.

### PartnerConfiguration

The `PartnerConfiguration` contract is used by the `CMAccount` and implements
features to register supported (supplier) and wanted (distributor) services,
register public keys that would be used to encrypt private data, off-chain payment
support, and on-chain supported payment token addresses.

### ServiceRegistry

The `ServiceRegistry` contract is used by the `CMAccountManager` contract and
implements a registry that is used to hash service names to keccak256 hashes and
store them in a mapping as `hash => service name`. `CMAccount` use these to resolve
hashes to service names and service names to hashes.

### BookingToken

The `BookingToken` contract is an ERC-721 NFT contract that is used by the partners
to mint and buy Booking Tokens. A Booking Token represents a booking done on the
Camino Messenger ecosystem.

Only the `CMAccount` contracts are allowed to mint and buy the tokens.

### KYCUtils

The `KYCUtils` contract provides utility functions for KYC (Know Your Customer).

### Proxies

For `CMAccountManager` and `CMAccount` contracts, an `ERC1967Proxy` (UUPS) is used.

The **`hardhat-ignition`** module deploys the `CMAccountManager` contract and then
deploys an `ERC1967Proxy` proxy, setting the implementation address to the
`CMAccountManager`'s address. We will call this proxy **managerProxy** or simply
**manager** in this document.

Then a `CMAccount` contract is deployed, and its address is set by calling
`managerProxy.setAccountImplementation(CMAccount.getAddress())`. After that, the
manager is ready to create CM accounts.

Calling `managerProxy.createCMAccount(...)` with the necessary arguments creates an
`ERC1967Proxy` and sets the implementation address to the recorded account
implementation address in the manager. After it is deployed, it is immediately (same
transaction) initialized with the given arguments.

## Deploy Contracts Locally

### Run local hardhat node

```
yarn hardhat node
```

### Deploy contracts using the ignition module

```
yarn hardhat ignition deploy ignition/modules/0_development.js --network localhost
```

### Output should be similar to this

```
yarn run v1.22.19
$ /hgst/work/github.com/chain4travel/camino-messenger-contracts/node_modules/.bin/hardhat ignition deploy ignition/modules/0_development.js --network localhost
Hardhat Ignition 🚀

Deploying [ CMAccountManagerModule ]

Batch #1
  Executed BookingTokenProxyModule#BookingToken
  Executed CMAccountManagerModule#BookingTokenOperator
  Executed ManagerProxyModule#CMAccountManager

Batch #2
  Executed BookingTokenProxyModule#ERC1967Proxy
  Executed CMAccountManagerModule#CMAccount
  Executed ManagerProxyModule#ERC1967Proxy

Batch #3
  Executed CMAccountManagerModule#BookingToken
  Executed CMAccountManagerModule#CMAccountManager

Batch #4
  Executed CMAccountManagerModule#BookingToken.initialize
  Executed CMAccountManagerModule#CMAccountManager.initialize
  Executed CMAccountManagerModule#CMAccountManager.setAccountImplementation
  Executed CMAccountManagerModule#CMAccountManager.setBookingTokenAddress

[ CMAccountManagerModule ] successfully deployed 🚀

Deployed Addresses

BookingTokenProxyModule#BookingToken - 0x5FbDB2315678afecb367f032d93F642f64180aa3
CMAccountManagerModule#BookingTokenOperator - 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
ManagerProxyModule#CMAccountManager - 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
BookingTokenProxyModule#ERC1967Proxy - 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9
CMAccountManagerModule#CMAccount - 0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9
ManagerProxyModule#ERC1967Proxy - 0x5FC8d32690cc91D4c39d9d3abcBD16989F875707
CMAccountManagerModule#BookingToken - 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9
CMAccountManagerModule#CMAccountManager - 0x5FC8d32690cc91D4c39d9d3abcBD16989F875707
Done in 3.78s.
```

You can also see your deployed contract addresses in the
`ignition/deployments/<chainid>/deployed_addresses.json` file.

### Visualize the deployment

```
yarn hardhat ignition visualize ignition/modules/0_development.js
```

This will open a browswer tab with the deployment flow visualized.

## Cheques: Create & Sign

This document provides a detailed guide on creating, signing, and verifying cheques
using Camino Messenger smart contracts. The guide includes information on the cheque
structure, type hashes, domain separator, and EIP-712 standard.

### EIP-712 Overview

EIP-712 is a standard for hashing and signing of typed structured data. It is used
to improve the usability of off-chain message signing for on-chain verification. The
standard defines a structured format for messages, allowing them to be easily parsed
and verified. For more info see: https://eips.ethereum.org/EIPS/eip-712

#### Cheque Typehash, Domain Typehash, and Domain Separator

-   **Cheque Typehash:** This is the `keccak256` hash of the MessengerCheque struct type.

    ```js
    function calculateMessengerChequeTypeHash() {
        const typeHash = ethers.keccak256(
            ethers.toUtf8Bytes(
                "MessengerCheque(address fromCMAccount,address toCMAccount,address toBot,uint256 counter,uint256 amount,uint256 createdAt,uint256 expiresAt)",
            ),
        );
        return typeHash;
    }
    ```

-   **Domain Typehash:** This is the `keccak256` hash of the EIP-712 domain type.

    ```js
    function calculateDomainTypeHash() {
        const domainTypeHash = ethers.keccak256(
            ethers.toUtf8Bytes("EIP712Domain(string name,string version,uint256 chainId)"),
        );
        return domainTypeHash;
    }
    ```

-   **Domain Separator:** This is a hash that combines the domain typehash with the
    name, version, and chain ID of the domain.

    ```js
    function calculateDomainSeparator(domainName, domainVersion, chainId) {
        const coder = AbiCoder.defaultAbiCoder();
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
    ```

### Creating Cheques

#### Define a cheque structure

```js
const cheque = {
    fromCMAccount: "0x...", // Address of the CM Account of the sending bot (signer)
    toCMAccount: "0x...", // Address of the CM Account of the receiving bot
    toBot: "0x...", // Address of the bot receiving the cheque
    counter: 123, // Counter, needs to be incremented for each cheque
    amount: ethers.parseUnits("1.0", "ether"), // 1 ETH, amount to pay (after substracting the last paid amount)
    createdAt: Math.floor(Date.now() / 1000), // Current Unix timestamp, as an example
    expiresAt: Math.floor(Date.now() / 1000) + 86400, // Expiration timestamp
};
```

**Example:**

```js
const cheque = {
    fromCMAccount: "0x90F79bf6EB2c4f870365E785982E1f101E93b906",
    toCMAccount: "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
    toBot: "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65",
    counter: 123n,
    amount: 1000000000000000000n,
    createdAt: 1722282175n,
    expiresAt: 1722652561n,
};
```

#### Calculate Messenger Cheque Hash

**Messenger Cheque Typehash:**

```js
function calculateMessengerChequeHash(cheque) {
    const chequeTypeHash = calculateMessengerChequeTypeHash();

    const coder = ethers.AbiCoder.defaultAbiCoder();
    const encodedCheque = coder.encode(
        ["bytes32", "address", "address", "address", "uint256", "uint256", "uint256", "uint256"],
        [
            chequeTypeHash,
            cheque.fromCMAccount,
            cheque.toCMAccount,
            cheque.toBot,
            cheque.counter,
            cheque.amount,
            cheque.createdAt,
            cheque.expiresAt,
        ],
    );
    return ethers.keccak256(encodedCheque);
```

#### Sign cheque

The function below uses **`ethers.signTypedData`** to sign the cheque, which is
calculating the type hashes and domain separator from the provided data according to
EIP-712 specification. So, the functions above are for when you want to calculate
the hashes separately.

```js
async function signMessengerCheque(cheque, signer) {
    const chainId = await signer.provider.getNetwork().then((n) => n.chainId);

    const types = {
        MessengerCheque: [
            { name: "fromCMAccount", type: "address" },
            { name: "toCMAccount", type: "address" },
            { name: "toBot", type: "address" },
            { name: "counter", type: "uint256" },
            { name: "amount", type: "uint256" },
            { name: "createdAt", type: "uint256" },
            { name: "expiresAt", type: "uint256" },
        ],
    };

    const DOMAIN_NAME = "CaminoMessenger";
    const DOMAIN_VERSION = "1";

    const domain = {
        name: DOMAIN_NAME,
        version: DOMAIN_VERSION,
        chainId: chainId,
    };

    const signature = await signer.signTypedData(domain, types, cheque);
    return signature;
}
```

If you would like to calculate the typed data hash yourself and sign it, check out
the function that is generating the hash on the `CMAccount` contract:

```solidity
function hashTypedDataV4(MessengerCheque memory cheque) public view returns (bytes32) {
    return keccak256(abi.encodePacked("\x19\x01", DOMAIN_SEPARATOR, hashMessengerCheque(cheque)));
}
```

And also in [`utils/cheques.js`](utils/cheques.js) file (Javascript):

```js
function calculateTypedDataHash(cheque, domainSeparator) {
    const chequeHash = calculateMessengerChequeHash(cheque);
    return ethers.keccak256(ethers.concat([ethers.toUtf8Bytes("\x19\x01"), domainSeparator, chequeHash]));
}
```

> [!TIP]
> All the functions mentioned above can be seen from [`utils/cheques.js`](utils/cheques.js) file.

## Cheques: Verify

Cheque verification is normally done on-chain by the `verifyCheque` function on the
CM Account contract of the cheque's drawer (the bot who signed the cheque).

Signature of the function is like this:

```solidity
function verifyCheque(
    MessengerCheque memory cheque,
    bytes memory signature
) public returns (address signer, uint256 paymentAmount) {}
```

This function does not only verify that the signer of the cheque is a registered bot
on the CM Account, but also other verifications like:

-   If the `fromCMAccount` is the contract itself
-   If the address of `toCMAccount` is a registered CM Account on the manager
-   If `expiresAt` timestamp is bigger then `block.timestamp`
-   Last counter and last amount recorded on the contract are lower then the cheque's
-   If the `toBot` address has the required role (`CHEQUE_OPERATOR_ROLE`)

So, to only verify if cheque's signature is valid, without doing the verifications
above (which can only be done on-chain), you can use the examples below.

### Verify Cheque Signature Off-Chain

#### JavaScript

For example code about how to verify signatures off-chain using JavaScript and
Ethers.js, check out the [`test/ChequeManager.test.js`](test/ChequeManager.test.js) test file.

> [!NOTE]
>
> **Signing without `signTypedData`:**
>
> If you would like to learn how you can sign a cheque without relying `ethers.js`'s
> `signTypedData` (EIP712), there is an example script at
> [`examples/sign_primitive.js`](examples/sign_primitive.js).

#### Go

**TODO:** WIP

#### Python

**TODO:** WIP
