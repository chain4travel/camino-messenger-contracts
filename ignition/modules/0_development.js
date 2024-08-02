const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");
const { vars } = require("hardhat/config");

const managerProxyModule = buildModule("ManagerProxyModule", (m) => {
    const cmAccountManager = m.contract("CMAccountManager");

    const ManagerProxy = m.contract("ERC1967Proxy", [cmAccountManager, "0x"]);

    return { ManagerProxy };
});

const bookingTokenProxyModule = buildModule("BookingTokenProxyModule", (m) => {
    const bookingToken = m.contract("BookingToken");

    const BookingTokenProxy = m.contract("ERC1967Proxy", [bookingToken, "0x"]);

    return { BookingTokenProxy };
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
    const { ManagerProxy } = m.useModule(managerProxyModule);

    // Create instance of the proxy contract with the CMAccountManager ABI
    const managerProxy = m.contractAt("CMAccountManager", ManagerProxy);

    // Initialize the manager
    m.call(managerProxy, "initialize", [admin, pauser, upgrader, versioner, developerWallet, developerFeeBp]);

    // Deploy CMAccount implementation
    const CMAccountImpl = m.contract("CMAccount");

    // Set the CMAccount implementation
    m.call(managerProxy, "setAccountImplementation", [CMAccountImpl]);

    // BookingToken
    const { BookingTokenProxy } = m.useModule(bookingTokenProxyModule);

    // Create instance of the proxy contract with the BookingToken ABI
    const bookingTokenProxy = m.contractAt("BookingToken", BookingTokenProxy);

    // Initialize the booking token
    m.call(bookingTokenProxy, "initialize", [managerProxy.address, admin, upgrader]);

    // Set the booking token address in the manager
    m.call(managerProxy, "setBookingToken", [bookingTokenProxy.address]);

    return { managerProxy };
});

module.exports = CMAccountManagerModule;
