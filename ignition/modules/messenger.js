const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

const CaminoMessengerModule = buildModule("CaminoMessengerModule", (m) => {
    /***************************************************
     *                  SET ACCOUNTS                   *
     ***************************************************/

    // Use the first account as the admin. For local node this is the first account
    // from hardhat's example accounts. For Camino (mainnet) and Columbus (testnet)
    // this is the vars from hardhat's config vars CAMINO_DEPLOYER_PRIVATE_KEY and
    // COLUMBUS_DEPLOYER_PRIVATE_KEY (defined in hardhat.config.js).
    const admin = m.getParameter("managerAdmin", m.getAccount(0));

    const pauser = m.getParameter("managerPauser", admin);
    const upgrader = m.getParameter("managerUpgrader", admin);
    const versioner = m.getParameter("managerVersioner", admin);

    /***************************************************
     *                DEVELOPER WALLET                 *
     ***************************************************/

    // We need a developer wallet to initialize the CMAccountManager
    const developerWallet = m.getParameter("developerWallet", admin);

    // Developer fee basis points, 100 bp == 1%
    const developerFeeBp = m.getParameter("developerFeeBp", 100);

    /***************************************************
     *                     MANAGER                     *
     ***************************************************/

    // Deploy CMAccountManager implementation contract
    const cmAccountManager = m.contract("CMAccountManager");

    // Encode the initialize function call for CMAccountManager
    const initializeManagerData = m.encodeFunctionCall(cmAccountManager, "initialize", [
        admin,
        pauser,
        upgrader,
        versioner,
        developerWallet,
        developerFeeBp,
    ]);

    // Deploy the proxy contract for CMAccountManager with the initialize data
    const ManagerProxy = m.contract("ERC1967Proxy", [cmAccountManager, initializeManagerData], {
        id: "ManagerERC1967Proxy",
    });

    // Create instance of the proxy contract with the CMAccountManager ABI
    const managerProxy = m.contractAt("CMAccountManager", ManagerProxy, { id: "ManagerProxy" });

    /***************************************************
     *             CM ACCOUNT IMPLEMENTATION           *
     ***************************************************/

    // BookingTokenOperator library
    const bookingTokenOperator = m.library("BookingTokenOperator");

    // Deploy CMAccount implementation with the BookingTokenOperator library
    const CMAccountImpl = m.contract("CMAccount", [], {
        libraries: { BookingTokenOperator: bookingTokenOperator },
    });

    // Set the CMAccount implementation in the manager
    m.call(managerProxy, "setAccountImplementation", [CMAccountImpl]);

    /***************************************************
     *                  BOOKING TOKEN                  *
     ***************************************************/

    // Booking token admin and upgrader. Set default to the manager admin.
    // Configurable by a parameter json file.
    const bookingAdmin = m.getParameter("bookingAdmin", admin);
    const bookingUpgrader = m.getParameter("bookingUpgrader", admin);

    // Deploy BookingToken implementation contract
    const bookingToken = m.contract("BookingToken");

    // Encode the initialize function call for BookingToken
    const initializeBookingTokenData = m.encodeFunctionCall(bookingToken, "initialize", [
        managerProxy.address,
        bookingAdmin,
        bookingUpgrader,
    ]);

    // Deploy the proxy contract for BookingToken with the initialize data
    const BookingTokenProxy = m.contract("ERC1967Proxy", [bookingToken, initializeBookingTokenData], {
        id: "BookingTokenERC1967Proxy",
    });

    // Create instance of the proxy contract with the BookingToken ABI
    const bookingTokenProxy = m.contractAt("BookingToken", BookingTokenProxy, { id: "BookingTokenProxy" });

    /***************************************************
     *                   POST CONFIG                   *
     ***************************************************/

    // Set the booking token address in the manager
    m.call(managerProxy, "setBookingTokenAddress", [bookingTokenProxy.address]);

    return { managerProxy, bookingTokenProxy, CMAccountImpl };
});

module.exports = CaminoMessengerModule;
