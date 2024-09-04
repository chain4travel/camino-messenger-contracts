const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");
const { vars } = require("hardhat/config");
const CaminoMessengerModule = require("./messenger");

// Simple upgrade test

const upgradeModule = buildModule("UpgradeModule", (m) => {
    const { managerProxy } = m.useModule(CaminoMessengerModule);

    const cmAccountManagerV2 = m.contract("CMAccountManagerV2");

    m.call(managerProxy, "upgradeToAndCall", [cmAccountManagerV2, "0x"]);

    return { managerProxy };
});

module.exports = upgradeModule;
