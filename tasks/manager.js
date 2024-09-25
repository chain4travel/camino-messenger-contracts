require("@nomicfoundation/hardhat-toolbox");

const MANAGER_SCOPE = scope("manager", "CM Account Manager Tasks");

// TODO: Handle transaction failures

function getAddressesForNetwork(hre) {
    let addresses;

    if (hre.network.name === "columbus") {
        console.log("Running on columbus");
        addresses = require("../ignition/deployments/chain-501/deployed_addresses.json");
    } else if (hre.network.name === "camino") {
        console.log("Running on camino");
        addresses = require("../ignition/deployments/chain-500/deployed_addresses.json");
    } else if (hre.network.name === "localhost") {
        console.log("Running on localhost");
        addresses = require("../ignition/deployments/chain-31337/deployed_addresses.json");
    } else {
        throw new Error(`Unsupported network: ${hre.network.name}`);
    }

    return addresses;
}

async function getManager(hre) {
    const addresses = getAddressesForNetwork(hre);
    return await ethers.getContractAt("CMAccountManager", addresses["CaminoMessengerModule#ManagerProxy"]);
}

async function handleRoles(taskArgs, hre, action) {
    const manager = await getManager(hre);

    console.log(
        `${action === "grantRole" ? "Granting" : "Revoking"} role ${taskArgs.role} for address ${taskArgs.address}...`,
    );

    const role = await manager[taskArgs.role]();
    const tx = await manager[action](role, taskArgs.address);
    const txReceipt = await tx.wait();
    console.log("Tx:", txReceipt.hash);
}

function handleTransactionError(error, contract) {
    console.error("❌ Transaction failed!");

    if (error.data.data && contract) {
        const decodedError = contract.interface.parseError(error.data.data);
        console.error("Message:", error.message);
        console.error(`Reason: ${decodedError?.name} (${decodedError?.args})`);
    } else if (error.data?.message) {
        console.error(`Reason: ${error.data.message}`);
    } else if (error.message?.includes("[taskArgs.role] is not a function")) {
        console.error("Reason: CMAccount does not have this role.");
    } else if (error.message) {
        console.error("Message:", error.message);
    } else {
        // General error logging
        console.error("An unexpected error occurred.");
        console.error("Error:", error);
    }
}

async function handleServices(taskArgs, hre, action) {
    const manager = await getManager(hre);

    console.log(`${action === "register" ? "Registering" : "Unregistering"} services...`);

    // Iterate over the services from the services file and perform the action
    const services = require(taskArgs.json);
    for (const service of services) {
        console.log(`⏳ ${action === "register" ? "Registering" : "Unregistering"} Service:`, service);
        try {
            const tx = await manager[`${action}Service`](service);
            const txReceipt = await tx.wait();
            console.log("✅ Service:", service, "Tx:", txReceipt.hash);
        } catch (error) {
            handleTransactionError(error, manager);
        }
        console.log("-----------------------------------------------------------");
    }
}

MANAGER_SCOPE.task("status", "Print status of deployed contracts").setAction(async (taskArgs, hre) => {
    const addresses = getAddressesForNetwork(hre);

    const managerImplementation = await ethers.getContractAt(
        "CMAccountManager",
        addresses["CaminoMessengerModule#CMAccountManager"],
    );

    const manager = await ethers.getContractAt("CMAccountManager", addresses["CaminoMessengerModule#ManagerProxy"]);

    const cmAccount = await ethers.getContractAt("CMAccount", addresses["CaminoMessengerModule#CMAccount"]);

    const bookingTokenImplementation = await ethers.getContractAt(
        "BookingToken",
        addresses["CaminoMessengerModule#BookingToken"],
    );

    const bookingToken = await ethers.getContractAt(
        "BookingToken",
        addresses["CaminoMessengerModule#BookingTokenProxy"],
    );

    console.log("========================= MANAGER =========================");
    console.log(`Proxy: ${await manager.getAddress()}`);
    console.log(`Implementation: ${await managerImplementation.getAddress()}`);

    console.log();
    console.log("======================== CM ACCOUNT ========================");
    console.log(`Implementation: ${await cmAccount.getAddress()}`);

    console.log();
    console.log("====================== BOOKING TOKEN ======================");
    console.log(`Proxy: ${await bookingToken.getAddress()}`);
    console.log(`Implementation: ${await bookingTokenImplementation.getAddress()}`);

    console.log();
    console.log("====================== CONFIGURATION ======================");
    console.log(`CM Account Impl: ${await manager.getAccountImplementation()}`);
    console.log(`Developer Wallet: ${await manager.getDeveloperWallet()}`);
    const feeBasisPoints = await manager.getDeveloperFeeBp();
    const feePercentage = (Number(feeBasisPoints) / 10000) * 100;
    console.log(`Developer Fee: ${feeBasisPoints}bp (${feePercentage}%)`);
    console.log(`Prefund Amount: ${ethers.formatEther(await manager.getPrefundAmount())} CAM`);
});

MANAGER_SCOPE.task("services:register", "Register services")
    .addParam("json", "Full path to the services json file")
    .setAction(async (taskArgs, hre) => {
        await handleServices(taskArgs, hre, "register");
    });

MANAGER_SCOPE.task("services:unregister", "Unregister services")
    .addParam("json", "Full path to the services json file")
    .setAction(async (taskArgs, hre) => {
        await handleServices(taskArgs, hre, "unregister");
    });

MANAGER_SCOPE.task("services:list", "List registered services").setAction(async (taskArgs, hre) => {
    const addresses = getAddressesForNetwork(hre);
    const manager = await ethers.getContractAt("CMAccountManager", addresses["CaminoMessengerModule#ManagerProxy"]);
    console.log("Getting all registered services...");
    const services = await manager.getAllRegisteredServiceNames();
    console.log(services);
});

MANAGER_SCOPE.task("role:grant", "Grant role")
    .addParam("role", "Role to grant. Ex: SERVICE_REGISTRY_ADMIN_ROLE")
    .addParam("address", "Address to grant role to")
    .setAction(async (taskArgs, hre) => {
        await handleRoles(taskArgs, hre, "grantRole");
    });

MANAGER_SCOPE.task("role:revoke", "Revoke role")
    .addParam("role", "Role to grant. Ex: SERVICE_REGISTRY_ADMIN_ROLE")
    .addParam("address", "Address to revoke role to")
    .setAction(async (taskArgs, hre) => {
        await handleRoles(taskArgs, hre, "revokeRole");
    });

MANAGER_SCOPE.task("role:has", "Check if address has role")
    .addParam("role", "Role to check. Ex: SERVICE_REGISTRY_ADMIN_ROLE")
    .addParam("address", "Address to check")
    .setAction(async (taskArgs, hre) => {
        const manager = await getManager(hre);
        const role = await manager[taskArgs.role]();
        const hasRole = await manager.hasRole(role, taskArgs.address);
        console.log(`${taskArgs.address} ${hasRole ? "has" : "does not have"} role ${taskArgs.role}`);
        console.log(`${hasRole ? "🟢" : "🔴"}`, hasRole);
    });

MANAGER_SCOPE.task("role:members", "List role members")
    .addParam("role", "Role to list. Ex: SERVICE_REGISTRY_ADMIN_ROLE")
    .setAction(async (taskArgs, hre) => {
        const manager = await getManager(hre);
        const role = await manager[taskArgs.role]();
        const count = await manager.getRoleMemberCount(role);
        console.log("Role:", taskArgs.role);
        console.log("Total Members:", count);

        // Iterate over the members of the role
        const members = [];
        for (let i = 0; i < count; i++) {
            const member = await manager.getRoleMember(role, i);
            members.push(member);
        }
        console.log(members);
    });

MANAGER_SCOPE.task("account:list", "List CM Accounts").setAction(async (taskArgs, hre) => {
    await hre.run({ scope: "manager", task: "role:members" }, { role: "CMACCOUNT_ROLE" });
});

MANAGER_SCOPE.task("account:set-implementation", "Set CMAccount implementation address")
    .addParam("address", "Implementation address to set as the new CMAccount impl")
    .setAction(async (taskArgs, hre) => {
        const manager = await getManager(hre);
        const tx = await manager.setAccountImplementation(taskArgs.address);
        const txReceipt = await tx.wait();
        console.log("Tx:", txReceipt.hash);
    });

module.exports = {};
