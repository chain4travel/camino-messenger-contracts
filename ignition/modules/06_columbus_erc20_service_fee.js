const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");
const hre = require("hardhat");

// Notes:
// - After this module is deployed:
//   - Give the caller SERVICE_FEE_TOKEN_ADMIN_ROLE role and ...
//   - call .setServiceFeeToken(serviceFeeTokenAddress) on the manager using that caller

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

const ERC20ServiceFeeModule = buildModule("ERC20ServiceFeeModule", (m) => {
    // Use addresses from the deployment to get the managerProxy and bookingTokenProxy contracts
    const addresses = getAddressesForNetwork(hre);

    /*********************************************
     *        Deploy New CMAccount Impl          *
     *********************************************/

    // Get deployed BookingTokenOperator
    const bookingTokenOperator = m.contractAt(
        "BookingTokenOperator",
        addresses["RefactorCancellationModule#BookingTokenOperator"],
    );

    // Deploy new CMAccount implementation
    const newCMAccountImpl = m.contract("CMAccount", [], {
        libraries: { BookingTokenOperator: bookingTokenOperator },
    });

    /*********************************************
     *      Deploy New CMAccountManager Impl     *
     *********************************************/

    // Get the manager proxy contract
    const managerProxy = m.contractAt("CMAccountManager", addresses["CaminoMessengerModule#ManagerProxy"], {
        id: "ManagerProxy",
    });

    // Deploy new CMAccountManager implementation
    const newCMAccountManagerImpl = m.contract("CMAccountManager");

    // No need to initialize as CMAccountManager is already deployed and we don't have reinitializers.

    // Update the CMAccountManager
    m.call(managerProxy, "upgradeToAndCall", [newCMAccountManagerImpl, "0x"]);

    // Set the new CMAccount implementation in the manager
    m.call(managerProxy, "setAccountImplementation", [newCMAccountImpl]);

    return { newCMAccountImpl, newCMAccountManagerImpl };
});

module.exports = ERC20ServiceFeeModule;
