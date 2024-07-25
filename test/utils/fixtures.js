/**
 * @dev Fixtures
 */
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { ethers, upgrades } = require("hardhat");

const developerFeeBp = 100;

async function setupSigners() {
    const [
        managerAdmin,
        managerPauser,
        managerUpgrader,
        managerVersioner,
        cmAccountAdmin,
        cmAccountPauser,
        cmAccountUpgrader,
        developerWallet,
        developerWalletAdmin,
        feeAdmin,
        chequeOperator,
        depositor,
        withdrawer,
        otherAccount1,
        otherAccount2,
        otherAccount3,
    ] = await ethers.getSigners();

    signers = {
        managerAdmin,
        managerPauser,
        managerUpgrader,
        managerVersioner,
        cmAccountAdmin,
        cmAccountPauser,
        cmAccountUpgrader,
        developerWallet,
        developerWalletAdmin,
        feeAdmin,
        chequeOperator,
        depositor,
        withdrawer,
        otherAccount1,
        otherAccount2,
        otherAccount3,
    };
}

async function deployCMAccountManagerFixture() {
    // Set up signers
    await setupSigners();

    const CMAccountManager = await ethers.getContractFactory("CMAccountManager");
    const cmAccountManager = await upgrades.deployProxy(
        CMAccountManager,
        [
            signers.managerAdmin.address,
            signers.managerPauser.address,
            signers.managerUpgrader.address,
            signers.managerVersioner.address,
            signers.developerWallet.address,
            developerFeeBp,
        ],
        { kind: "uups" },
    );
    return { cmAccountManager };
}

async function deployCMAccountImplFixture() {
    const CMAccount = await ethers.getContractFactory("CMAccount");
    const cmAccountImpl = await CMAccount.deploy();
    await cmAccountImpl.waitForDeployment();

    return { cmAccountImpl };
}

async function deployCMAccountManagerWithCMAccountImplFixture() {
    // Set up signers
    await setupSigners();

    const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);
    const { cmAccountImpl } = await loadFixture(deployCMAccountImplFixture);

    const cmAccountImplAddress = await cmAccountImpl.getAddress();

    await cmAccountManager.grantRole(await cmAccountManager.VERSIONER_ROLE(), signers.managerVersioner.address);
    await cmAccountManager.connect(signers.managerVersioner).setAccountImplementation(cmAccountImplAddress);

    return { cmAccountManager, cmAccountImplAddress };
}

async function deployAndConfigureAllFixture() {
    // Set up signers
    await setupSigners();

    const { cmAccountManager, cmAccountImplAddress } = await loadFixture(
        deployCMAccountManagerWithCMAccountImplFixture,
    );

    await cmAccountManager.grantRole(
        await cmAccountManager.DEVELOPER_WALLET_ADMIN_ROLE(),
        signers.developerWalletAdmin.address,
    );
    await cmAccountManager.grantRole(await cmAccountManager.FEE_ADMIN_ROLE(), signers.feeAdmin.address);

    const tx = await cmAccountManager.createCMAccount(
        signers.cmAccountAdmin.address,
        signers.cmAccountPauser.address,
        signers.cmAccountUpgrader.address,
        true, // anyOneCanDeposit
    );

    const receipt = await tx.wait();

    // Parse event to get the CMAccount address (this is the UUPS proxy address)
    const event = receipt.logs.find((log) => {
        try {
            return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
        } catch (e) {
            return false;
        }
    });

    const parsedEvent = cmAccountManager.interface.parseLog(event);
    const cmAccountAddress = parsedEvent.args.account;

    // Get the CMAccount instance at the address
    const cmAccount = await ethers.getContractAt("CMAccount", cmAccountAddress);

    return { cmAccountManager, cmAccount };
}

async function deployCMAccountWithDepositFixture() {
    // Set up signers
    await setupSigners();

    const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

    // Grant withdrawer role
    const WITHDRAWER_ROLE = await cmAccount.WITHDRAWER_ROLE();
    await cmAccount.connect(signers.cmAccountAdmin).grantRole(WITHDRAWER_ROLE, signers.withdrawer.address);

    // Grant depositor role
    const DEPOSITOR_ROLE = await cmAccount.DEPOSITOR_ROLE();
    await cmAccount.connect(signers.cmAccountAdmin).grantRole(DEPOSITOR_ROLE, signers.depositor.address);

    const depositAmount = ethers.parseEther("1");

    await cmAccount.connect(signers.cmAccountAdmin).setAnyoneCanDeposit(true);
    await cmAccount.connect(signers.depositor).deposit({ value: depositAmount });

    return { cmAccount, cmAccountManager, WITHDRAWER_ROLE };
}

module.exports = {
    setupSigners,
    developerFeeBp,
    deployCMAccountManagerFixture,
    deployCMAccountImplFixture,
    deployCMAccountManagerWithCMAccountImplFixture,
    deployAndConfigureAllFixture,
    deployCMAccountWithDepositFixture,
};
