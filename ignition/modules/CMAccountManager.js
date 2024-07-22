const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");
const { vars } = require("hardhat/config");

const proxyModule = buildModule("ProxyModule", (m) => {
    const cmAccountManager = m.contract("CMAccountManager");

    const proxy = m.contract("ERC1967Proxy", [cmAccountManager, "0x"]);

    return { cmAccountManager, proxy };
});

const CMAccountManagerModule = buildModule("CMAccountManagerModule", (m) => {
    const admin = m.getAccount(0);

    const { proxy } = m.useModule(proxyModule);

    const cmAccountManager = m.contractAt("CMAccountManager", proxy);

    m.call(cmAccountManager, "initialize", [admin, admin, admin]);

    return { cmAccountManager, proxy };
});

module.exports = CMAccountManagerModule;
