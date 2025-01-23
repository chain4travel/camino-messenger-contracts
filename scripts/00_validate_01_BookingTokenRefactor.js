// Test for storage layout incompatibility
//
// First force import the network file for the OZ's upgrades:
//   git checkout a53b8fc03973025f4ed10edb46aa4d2f0c3e76ee
//   yarn hardhat run scripts/00_validate_01_force_import.js --network columbus
//
// That will create a file ./openzeppelin/unknown-501.json
//
// Then run the test:
//   git checkout 1b5d14fca94a9129f517d8dd3e3d6eef47195d96
//   DEBUG=@openzeppelin:* yarn hardhat test scripts/00_validate_01_BookingTokenRefactor --network columbus

const { ethers, upgrades } = require("hardhat");

function getAddressesForNetwork(hre) {
    let addresses;

    if (hre.network.name === "columbus") {
        console.log("Running on columbus");
        addresses = require("../ignition/deployments/chain-501/deployed_addresses.json");
    } else if (hre.network.name === "camino") {
        console.log("Running on camino");
        addresses = require("../ignition/deployments/chain-500/deployed_addresses.json");
    } else if (hre.network.name === "localhost") {
        console.log("Running on localhost");
        addresses = require("../ignition/deployments/chain-31337/deployed_addresses.json");
    } else {
        throw new Error(`Unsupported network: ${hre.network.name}`);
    }

    return addresses;
}

describe("Upgrade Check", function () {
    it("Should validate storage compatibility", async function () {
        const addresses = getAddressesForNetwork(hre);

        // Address at run time on Columbus: 0xe55E387F5474a012D1b048155E25ea78C7DBfBBC
        const NewBookingToken = await ethers.getContractFactory("BookingToken");

        await upgrades.validateUpgrade(addresses["CaminoMessengerModule#BookingTokenProxy"], NewBookingToken);
    });
});

// [Note@2025-01-23]
//
// Output for BookingTokenV2 contract that was used before the refactoring at
// https://github.com/chain4travel/camino-messenger-contracts/pull/55
//
// The output below compares the deployed BookingTokenV2 contract storage layout
// with the new BookingToken contract storage layout after the refactoring.
//
// It is showing the incompatible changes between the two as expected only related
// to cancellation.
//
// There are also some ABI changes. These are only acceptable because the contract
// is not deployed to mainnet yet and the production usage of the contract is
// minimal.
//
// ----------8<----------
// Upgrade Check Running on columbus
// @openzeppelin:upgrades:core manifest file: .openzeppelin/unknown-501.json
// fallback file: .openzeppelin/unknown-501.json +0ms @openzeppelin:upgrades:core
// manifest file: .openzeppelin/unknown-501.json fallback file:
// .openzeppelin/unknown-501.json +240ms 1) Should validate storage compatibility
//
//
//   0 passing (1s) 1 failing
//
//   1) Upgrade Check Should validate storage compatibility: Error: New storage
//        layout is incompatible
//
// BookingTokenV2: Deleted `_cancellationProposals`
//   > Keep the variable even if unused
//
// contracts/booking-token/BookingTokenCancellable.sol:39: Replaced `_isCancellable`
// with `_proposals` of incompatible type
//
// contracts/booking-token/BookingToken.sol:138: Upgraded `_bookingStatus` to an
// incompatible type
//   - In mapping(uint256 => enum BookingToken.BookingStatus)
//     - Bad upgrade to enum BookingToken.BookingStatus
//   - In enum BookingToken.BookingStatus
//     - Replaced `Unspecified` with `UNSPECIFIED`
//     - Replaced `Reserved` with `RESERVED`
//     - Replaced `Expired` with `RESERVATION_EXPIRED`
//     - Replaced `Bought` with `BOUGHT`
//     - Replaced `Cancelled` with `CANCELLED` at assertStorageUpgradeSafe
//       (node_modules/@openzeppelin/upgrades-core/src/storage/index.ts:35:11) at
//       validateImpl
//       (node_modules/@openzeppelin/hardhat-upgrades/src/utils/validate-impl.ts:48:31)
//       at Proxy.validateUpgrade
//       (node_modules/@openzeppelin/hardhat-upgrades/src/validate-upgrade.ts:51:9)
//       at Context.<anonymous> (scripts/validate_01_BookingTokenRefactor.js:32:9)
// ---------->8----------
