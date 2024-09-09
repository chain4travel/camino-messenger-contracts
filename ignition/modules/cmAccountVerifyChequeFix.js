const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

// Deploys a new CMAccount with the verifyCheque function as a view function
const CMACcountVerifyChequeFixModule = buildModule("CMAccountVerifyChequeFixModule", (m) => {
    // Get the BookingTokenOperator library on Columbus testnet
    bookingTokenOperator = m.contractAt("BookingTokenOperator", "0x10133935503b4f958f6dFF783b628ba25aC010E3");

    // Deploy CMAccount implementation with the BookingTokenOperator library
    const CMAccountImpl = m.contract("CMAccount", [], {
        libraries: { BookingTokenOperator: bookingTokenOperator },
    });

    return { CMAccountImpl };
});

module.exports = CMACcountVerifyChequeFixModule;
