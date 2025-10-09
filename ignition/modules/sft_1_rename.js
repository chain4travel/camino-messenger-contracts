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

const SFTRenameModule = buildModule("SFTRenameModule", (m) => {
    const addresses = getAddressesForNetwork(hre);

    // Fail if network is camino
    if (hre.network.name === "camino") {
        console.log("⚠️  Do not deploy on Camino mainnet!");
        throw new Error(`Unsupported network: ${hre.network.name}`);
    }

    /***************************************************
     *                  SET ACCOUNTS                   *
     ***************************************************/

    // Use the first account as the admin. For local node this is the first account
    // from hardhat's example accounts. For Camino (mainnet) and Columbus (testnet)
    // this is the vars from hardhat's config vars CAMINO_DEPLOYER_PRIVATE_KEY and
    // COLUMBUS_DEPLOYER_PRIVATE_KEY (defined in hardhat.config.js).
    const admin = m.getAccount(0);

    /***************************************************
     *                SERVICE FEE TOKEN                *
     ***************************************************/

    // We will deploy a new implementation for the SFT contract to be able to change
    // the name. Normally, the name and the symbol are immutable in Openzeppelin
    // ERC20 contracts. Because of that, we need to reinitialize the SFT contract.
    // And the reinitialize function is added with an update.

    // Deploy new ServiceFeeToken implementation contract
    const newServiceFeeToken = m.contract("ServiceFeeToken");

    // Encode the reinitializeV2 function call for the name change
    const reinitializeV2 = m.encodeFunctionCall(newServiceFeeToken, "reinitializeV2", ["USD Test Token", "USD.test"]);

    // Get the SFT proxy address
    const sftProxy = m.contractAt("ServiceFeeToken", addresses["ServiceFeeTokenModule#ServiceFeeTokenProxy"], {
        id: "ServiceFeeTokenProxy",
    });

    // Upgrade the ServiceFeeToken contract with the new name and symbol
    m.call(sftProxy, "upgradeToAndCall", [newServiceFeeToken, reinitializeV2]);

    return { sftProxy };
});

module.exports = SFTRenameModule;
