const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");
const hre = require("hardhat");

function getAddressesForNetwork(hre) {
    let addresses;

    if (hre.network.name === "columbus") {
        console.log("Running on columbus");
        addresses = require("../deployments/chain-501/deployed_addresses.json");
    } else if (hre.network.name === "camino") {
        console.log("Running on camino");
        addresses = require("../deployments/chain-500/deployed_addresses.json");
    } else if (hre.network.name === "localhost") {
        console.log("Running on localhost");
        addresses = require("../deployments/chain-31337/deployed_addresses.json");
    } else {
        throw new Error(`Unsupported network: ${hre.network.name}`);
    }

    return addresses;
}

// Upgrades for the BookingToken Cancellation support
const RefactorCancellationModule = buildModule("RefactorCancellationModule", (m) => {
    // Exit if not columbus testnet
    if (hre.network.name === "camino") {
        console.log("ERROR: You shouldn't run this on Camino mainnet!");
        process.exit(1);
    }

    // Use addresses from the deployment to get the managerProxy and bookingTokenProxy contracts
    const addresses = getAddressesForNetwork(hre);

    const managerProxy = m.contractAt("CMAccountManager", addresses["CaminoMessengerModule#ManagerProxy"], {
        id: "ManagerProxy",
    });

    const bookingTokenProxy = m.contractAt("BookingToken", addresses["CaminoMessengerModule#BookingTokenProxy"], {
        id: "BookingTokenProxy",
    });

    // BookingTokenOperator is updated to support Cancellation. We need to deploy a
    // new library for CMAccount implementation.
    const bookingTokenOperator = m.library("BookingTokenOperator");

    // Deploy a new CMAccount implementation with the new BookingTokenOperator
    // library. There are also updates to the CMAccount impl contract to support
    // cancellation functions like initiate/accept/reject cancellation.
    //
    // Note: There is also changes after refactoring that are ABI breaking changes
    // only for the columbus testnet.
    const newCMAccountImpl = m.contract("CMAccount", [], {
        libraries: { BookingTokenOperator: bookingTokenOperator },
    });

    // Set the new CMAccount implementation in the manager
    m.call(managerProxy, "setAccountImplementation", [newCMAccountImpl], {
        id: "NewCMAccountImpl",
    });

    /*********************************************
     *      Upgrade BookingToken to Refactor     *
     *********************************************/

    // Deploy the BookingToken contract
    const bookingToken = m.contract("BookingToken", [], {
        id: "BookingTokenImpl",
    });

    // Disable this line below for columbus testnet as we already did it before.
    //
    // Encode the Reinitialize function call for BookingToken
    //const reinitializeV2 = m.encodeFunctionCall(bookingToken, "reinitializeV2", ["BookingToken", "BToken"]);
    //m.call(bookingTokenProxy, "upgradeToAndCall", [bookingToken, reinitializeV2]);

    // Upgrade the BookingToken contract
    m.call(bookingTokenProxy, "upgradeToAndCall", [bookingToken, "0x"]); // TODO: Use the reinitializeV2 for mainnet!

    return { managerProxy, bookingTokenProxy, bookingTokenOperator, newCMAccountImpl };
});

module.exports = RefactorCancellationModule;
