const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");
const { vars } = require("hardhat/config");
const CMAccountManagerModule = require("./CMAccountManager");

const upgradeModule = buildModule("UpgradeModule", (m) => {
    const { cmAccountManager, proxy } = m.useModule(CMAccountManagerModule);

    const cmAccountManagerV2 = m.contract("CMAccountManagerV2");

    m.call(cmAccountManager, "upgradeToAndCall", [cmAccountManagerV2, "0x"]);

    return { cmAccountManager, cmAccountManagerV2 };
});

module.exports = upgradeModule;
