const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");
const { vars } = require("hardhat/config");

const proxyModule = buildModule("ProxyModule", (m) => {
    const cmAccountManager = m.contract("CMAccountManager");

    const proxy = m.contract("ERC1967Proxy", [cmAccountManager, "0x"], { id: "CMAccountManagerProxy" });

    return { proxy };
});

const CMAccountManagerModule = buildModule("CMAccountManagerModule", (m) => {
    // Use the first account as the admin. TODO: Should be configurable
    const admin = m.getAccount(0);

    // Use admin for all the roles. TODO: Should be configurable
    const pauser = admin;
    const upgrader = admin;
    const versioner = admin;

    // We need a developer wallet to initialize the CMAccountManager
    const developerWallet = m.getParameter("developerWallet", admin);
    const developerFeeBp = m.getParameter("developerFeeBp", 100);

    // Take the proxy contract from the proxy module
    const { proxy } = m.useModule(proxyModule);

    // Create instance of the proxy contract with the CMAccountManager ABI
    const managerProxy = m.contractAt("CMAccountManager", proxy, { id: "CMAccountManagerProxy" });

    // Initialize the manager
    m.call(managerProxy, "initialize", [admin, pauser, upgrader, versioner, developerWallet, developerFeeBp]);

    // Deploy CMAccount implementation
    const CMAccountImpl = m.contract("CMAccount");

    // Set the CMAccount implementation
    m.call(managerProxy, "setAccountImplementation", [CMAccountImpl]);

    return { managerProxy };
});

module.exports = CMAccountManagerModule;
