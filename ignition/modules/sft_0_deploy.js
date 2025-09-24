const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");
const hre = require("hardhat");

const ServiceFeeTokenModule = buildModule("ServiceFeeTokenModule", (m) => {
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

    const pauser = m.getParameter("sftPauser", admin);
    const minter = m.getParameter("sftMinter", admin);
    const upgrader = m.getParameter("sftUpgrader", admin);

    /***************************************************
     *                SERVICE FEE TOKEN                *
     ***************************************************/

    // Deploy ServiceFeeToken implementation contract
    const serviceFeeToken = m.contract("ServiceFeeToken");

    // Encode the initialize function call for ServiceFeeToken
    const initializeServiceFeeTokenData = m.encodeFunctionCall(serviceFeeToken, "initialize", [
        admin,
        pauser,
        minter,
        upgrader,
    ]);

    // Deploy the proxy contract for ServiceFeeToken with the initialize data
    const ServiceFeeTokenProxy = m.contract("ERC1967Proxy", [serviceFeeToken, initializeServiceFeeTokenData], {
        id: "ServiceFeeTokenERC1967Proxy",
    });

    // Create instance of the proxy contract with the ServiceFeeToken ABI
    const serviceFeeTokenProxy = m.contractAt("ServiceFeeToken", ServiceFeeTokenProxy, {
        id: "ServiceFeeTokenProxy",
    });

    return { serviceFeeTokenProxy };
});

module.exports = ServiceFeeTokenModule;
