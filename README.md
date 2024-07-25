# Camino Messenger Contracts

[![CAMINO NETWORK](https://img.shields.io/badge/CAMINO-NETWORK-b440fc?style=for-the-badge&logoColor=white&labelColor=0085ff)](https://camino.network/) [![CHAT WITH US](https://img.shields.io/badge/DISCORD-5865F2?style=for-the-badge&logo=discord&logoColor=white)](https://discord.com/channels/949247897688494150/1182680860797960253)

[![CI](https://github.com/chain4travel/camino-messenger-contracts/actions/workflows/ci.yaml/badge.svg)](https://github.com/chain4travel/camino-messenger-contracts/actions/workflows/ci.yaml)

This repository contains the smart contracts for the [Camino
Messenger](https://camino.network/camino-messenger-sets-the-global-standard-in-travel-data-management-and-distribution/).

> [!WARNING]
> These contracts are currently in the development phase. The ABI is
> subject to change frequently until they are released into production.

## Quickstart

Clone the repo and change directory into:

```sh
git clone git@github.com:chain4travel/camino-messenger-contracts.git
cd camino-messenger-contracts
```

Install packages:

```sh
yarn install
```

Run tests:

```sh
yarn test
```

## Contracts

### `CMAccount`

The `CMAccount` contract represents a Camino Messenger account. Currently, it
includes functionalities for the management of bots. More features will be
introduced in future.

This contract works closely with the CMAccountManager to handle accounts.

### `CMAccountManager`

The `CMAccountManager` contract acts as a manager for CMAccount contracts. It
handles the creation, registration, verification, and management of accounts. It
also keeps records for developer wallet, fees, and `CMAccount` implementation
address. Accounts can only be upgraded to this implementation address that the
manager holds.

### `ChequeManager`

The `ChequeManager` contract handles the processing of cheques. It verifies the
signatures, checks the validity of the cheques, and transfers the funds between
accounts. It also calculates the developer fee and transfers it to the developer's
wallet.

This is a base contract that is inherited by the `CMAccount` contract.

### `KYCUtils`

The `KYCUtils` contract provides utility functions for KYC (Know Your Customer).
