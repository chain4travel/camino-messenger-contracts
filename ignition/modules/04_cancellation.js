const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");
const CaminoMessengerModule = require("./01_messenger");
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
const CancellationModule = buildModule("CancellationModule", (m) => {
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
    const newCMAccountImpl = m.contract("CMAccount", [], {
        libraries: { BookingTokenOperator: bookingTokenOperator },
    });

    // Set the new CMAccount implementation in the manager
    m.call(managerProxy, "setAccountImplementation", [newCMAccountImpl], {
        id: "NewCMAccountImpl",
    });

    /*********************************************
     *        Upgrade BookingToken to V2         *
     *********************************************/

    // Deploy the BookingTokenV2 contract
    const bookingTokenV2 = m.contract("BookingTokenV2", [], {
        id: "BookingTokenV2Impl",
    });

    // Encode the Reinitialize function call for BookingTokenV2
    const reinitializeV2 = m.encodeFunctionCall(bookingTokenV2, "reinitializeV2", ["BookingToken", "BToken"]);

    // Upgrade the BookingToken contract to V2
    m.call(bookingTokenProxy, "upgradeToAndCall", [bookingTokenV2, reinitializeV2]);

    return { managerProxy, bookingTokenProxy, bookingTokenOperator, newCMAccountImpl };
});

module.exports = CancellationModule;
