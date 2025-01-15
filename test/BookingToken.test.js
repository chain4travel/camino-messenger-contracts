const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");

const { ethers } = require("hardhat");

const helpers = require("@nomicfoundation/hardhat-network-helpers");

// Fixtures
const {
    setupSigners,
    developerFeeBp,
    deployCMAccountManagerFixture,
    deployCMAccountImplFixture,
    deployCMAccountManagerWithCMAccountImplFixture,
    deployAndConfigureAllFixture,
    deployCMAccountWithDepositFixture,
    deployBookingTokenFixture,
    deployBookingTokenWithNullUSDFixture,
    deployCancellationSupportFixture,
} = require("./utils/fixtures");

describe("BookingToken", function () {
    describe("Main", function () {
        it("should deploy correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            expect(await bookingToken.hasRole(await bookingToken.DEFAULT_ADMIN_ROLE(), signers.btAdmin.address)).to.be
                .true;
            expect(await bookingToken.hasRole(await bookingToken.UPGRADER_ROLE(), signers.btUpgrader.address)).to.be
                .true;
            expect(await bookingToken.isCMAccount(supplierCMAccount.getAddress())).to.be.true;
        });

        it("should return version correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            expect(await bookingToken.version()).to.deep.equal([1, 0, 0]);
        });

        it("should reinitialize correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const currentName = expect(await bookingToken.name()).to.be.equal("BookingToken");
            const currentSymbol = expect(await bookingToken.symbol()).to.be.equal("TRIP");

            const newName = "New Name";
            const newSymbol = "NEW";

            // Try to re-init with unauthorized caller
            await expect(
                bookingToken.connect(signers.otherAccount1).reinitializeV2(newName, newSymbol),
            ).to.be.revertedWithCustomError(bookingToken, "AccessControlUnauthorizedAccount");

            // Reinitialize
            await expect(bookingToken.connect(signers.btAdmin).reinitializeV2(newName, newSymbol)).to.not.reverted;

            // Check new name and symbol
            expect(await bookingToken.name()).to.be.equal(newName);
            expect(await bookingToken.symbol()).to.be.equal(newSymbol);

            // Try to re-init again, should revert
            await expect(
                bookingToken.connect(signers.btAdmin).reinitializeV2("New Name 2", "NEW2"),
            ).to.be.revertedWithCustomError(bookingToken, "InvalidInitialization");
        });

        it("should set/get manager address correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            // Try to set manager address with unauthorized caller
            await expect(
                bookingToken.connect(signers.otherAccount1).setManagerAddress(signers.otherAccount1.address),
            ).to.be.revertedWithCustomError(bookingToken, "AccessControlUnauthorizedAccount");

            // Set manager address
            expect(await bookingToken.connect(signers.btAdmin).setManagerAddress(signers.otherAccount1.address)).to.be
                .not.reverted;

            // Check manager address
            expect(await bookingToken.getManagerAddress()).to.be.equal(signers.otherAccount1.address);
        });

        it("should set/get min expiration timestamp diff correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const newMinExpirationTimestampDiff = 120;

            // Try to set min expiration timestamp diff with unauthorized caller
            await expect(
                bookingToken
                    .connect(signers.otherAccount1)
                    .setMinExpirationTimestampDiff(newMinExpirationTimestampDiff),
            ).to.be.revertedWithCustomError(bookingToken, "AccessControlUnauthorizedAccount");

            // Grant MIN_EXPIRATION_ADMIN_ROLE
            const MIN_EXPIRATION_ADMIN_ROLE = await bookingToken.MIN_EXPIRATION_ADMIN_ROLE();
            await bookingToken.connect(signers.btAdmin).grantRole(MIN_EXPIRATION_ADMIN_ROLE, signers.btAdmin.address);

            // Set min expiration timestamp diff
            expect(
                await bookingToken
                    .connect(signers.btAdmin)
                    .setMinExpirationTimestampDiff(newMinExpirationTimestampDiff),
            ).to.be.not.reverted;

            // Check min expiration timestamp diff
            expect(await bookingToken.getMinExpirationTimestampDiff()).to.be.equal(newMinExpirationTimestampDiff);
        });

        it("should support ERC165", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const _INTERFACE_ID_IERC165 = "0x01ffc9a7";
            const _INTERFACE_ID_IERC721 = "0x80ac58cd";
            const _INTERFACE_ID_IERC721METADATA = "0x5b5e139f";
            const _INTERFACE_ID_IERC721ENUMERABLE = "0x780e9d63";

            expect(await bookingToken.supportsInterface(_INTERFACE_ID_IERC165)).to.be.true;
            expect(await bookingToken.supportsInterface(_INTERFACE_ID_IERC721)).to.be.true;
            expect(await bookingToken.supportsInterface(_INTERFACE_ID_IERC721METADATA)).to.be.true;
            expect(await bookingToken.supportsInterface(_INTERFACE_ID_IERC721ENUMERABLE)).to.be.true;

            expect(await bookingToken.supportsInterface("0xaaaaaaaa")).to.be.false;
        });

        it("should upgrade correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const BookingTokenTest = await ethers.getContractFactory("BookingToken");
            const bookingTokenTest = await BookingTokenTest.deploy();

            // Try to upgrade with unauthorized caller
            await expect(
                bookingToken.connect(signers.otherAccount1).upgradeToAndCall(await bookingTokenTest.getAddress(), "0x"),
            ).to.be.revertedWithCustomError(bookingToken, "AccessControlUnauthorizedAccount");

            // Try to upgrade to unsupported implementation
            const DummyContract = await ethers.getContractFactory("Dummy");
            const dummyContract = await DummyContract.deploy();

            // Check dummy contract
            expect(await dummyContract.getVersion()).to.be.equal("DUMMY");

            await expect(
                bookingToken.connect(signers.btUpgrader).upgradeToAndCall(await dummyContract.getAddress(), "0x"),
            ).to.be.revertedWithCustomError(bookingToken, "ERC1967InvalidImplementation");

            // Upgrade to new implementation
            await expect(
                bookingToken.connect(signers.btUpgrader).upgradeToAndCall(await bookingTokenTest.getAddress(), "0x"),
            ).to.be.not.reverted;
        });
    });

    describe("Mint", function () {
        it("Native: should revert if not called from a CMAccount", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            await expect(
                bookingToken.connect(signers.btAdmin).safeMintWithReservation(
                    distributorCMAccount.getAddress(), // reservedFor
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                ),
            )
                .to.be.revertedWithCustomError(bookingToken, "NotCMAccount") // Caller is not a CMAccount
                .withArgs(signers.btAdmin.address);
        });

        it("Native: should revert invalid min expiration", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const minExpirationTimestamp = await bookingToken.getMinExpirationTimestampDiff();

            // get block time from the chain and mine
            await network.provider.send("evm_mine");
            const block = await ethers.provider.getBlock("latest");

            const invalidExpirationTimestamp = BigInt(block.timestamp) + minExpirationTimestamp - 1n;

            const price = ethers.parseEther("0.05");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Mint the booking token
            await expect(
                supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // reservedFor
                    tokenURI, // tokenURI
                    invalidExpirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                ),
            )
                .to.be.revertedWithCustomError(bookingToken, "ExpirationTimestampTooSoon") // Caller is not a CMAccount
                .withArgs(invalidExpirationTimestamp, minExpirationTimestamp);
        });

        it("Native: should revert if reservedFor is not a CMAccount", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    signers.otherAccount1.address, // set reservedFor to a non-CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                ),
            )
                .to.be.revertedWithCustomError(bookingToken, "NotCMAccount")
                .withArgs(signers.otherAccount1.address); // reservedFor address
        });

        it("Native: should revert off chain payment currency mismatch", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Off chain payment mismatch: Use paymentToken address of address(0) with non-zero off chain payment currency

            // Dummy off chain payment currency
            const dummyOffChainPaymentCurrency = 99n;

            await expect(
                supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // reservedFor
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    dummyOffChainPaymentCurrency, // Provide a dummy off chain payment currency, should revert
                    false,
                ),
            )
                .to.be.revertedWithCustomError(bookingToken, "UnexpectedOffchainPaymentCurrency")
                .withArgs(dummyOffChainPaymentCurrency);
        });

        it("Native: should mint a booking token correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token URI
            expect(await bookingToken.tokenURI(0n)).to.equal(tokenURI);

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            // Check cancellable flag
            expect(await bookingToken.isCancellable(0n)).to.equal(false);

            // Mint again to make sure the token id is incremented
            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    1n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(1n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(1n)).to.equal(1); // Reserved == 1
        });

        it("ERC20: should mint a booking token correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken, nullUSD } =
                await loadFixture(deployBookingTokenWithNullUSDFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("120");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    nullUSD.getAddress(), // nullUSD address
                    0,
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n, // token id
                    distributorCMAccount.getAddress(), // reservedFor
                    supplierCMAccount.getAddress(), // supplier
                    expirationTimestamp,
                    price,
                    nullUSD.getAddress(), // nullUSD address
                    0,
                    false,
                );

            // Sanity check
            expect(await bookingToken.getReservationPrice(0n)).to.be.deep.equal([price, await nullUSD.getAddress()]);
        });
    });

    describe("Buy", function () {
        it("Native: should buy a booking token correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Try to mint with non-auth address
            await expect(
                supplierCMAccount.connect(signers.otherAccount3).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                ),
            ).to.be.revertedWithCustomError(supplierCMAccount, "AccessControlUnauthorizedAccount");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Reverts: try to buy with invalid price
            const BookingTokenOperator = await ethers.getContractFactory("BookingTokenOperator");
            const invalidPrice = price + 1n;
            await expect(
                distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n, invalidPrice, ethers.ZeroAddress),
            )
                .to.revertedWithCustomError(BookingTokenOperator, "UnexpectedPrice")
                .withArgs(0n, price, invalidPrice);

            // Reverts: try to buy with invalid payment token
            const invalidPaymentToken = ethers.getAddress("0x1230000000000000000000000000000000000001");
            await expect(distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n, price, invalidPaymentToken))
                .to.revertedWithCustomError(BookingTokenOperator, "UnexpectedPaymentToken")
                .withArgs(0n, ethers.ZeroAddress, invalidPaymentToken);

            // Reverts: try to buy with non-auth address
            await expect(
                distributorCMAccount.connect(signers.otherAccount3).buyBookingToken(0n, price, ethers.ZeroAddress),
            ).to.be.revertedWithCustomError(distributorCMAccount, "AccessControlUnauthorizedAccount");

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n, price, ethers.ZeroAddress);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check balances
            await expect(buyTx).to.changeEtherBalances([supplierCMAccount, distributorCMAccount], [price, -price]);

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(3); // Bought == 3

            // Try to expire the token, should revert with InvalidTokenStatus
            await expect(distributorCMAccount.connect(signers.btAdmin).recordExpiration(0n))
                .to.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(0n, 3); // Bought == 3
        });

        it("Native: should buy a booking token with zero price correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n, price, ethers.ZeroAddress);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check balances
            await expect(buyTx).to.changeEtherBalances([supplierCMAccount, distributorCMAccount], [price, -price]);

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());
        });

        it("Off-chain: should buy a booking token with off-chain payment correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseUnits("559.99", 5);

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Off-chain payment marker
            const offChainPaymentMarker = await bookingToken.OFFCHAIN_PAYMENT();

            // address(1)
            const OneAddress = ethers.getAddress("0x0000000000000000000000000000000000000001");

            expect(offChainPaymentMarker).to.equal(OneAddress);

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    offChainPaymentMarker, // off-chain payment marker, address(1)
                    6, // off chain payment currency, 6 == Euro
                    true, // cancellable
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    offChainPaymentMarker, // off-chain payment marker, address(1)
                    6, // off chain payment currency, 6 == Euro
                    true, // cancellable
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n, price, OneAddress);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check balances, balances should not change as it is off-chain payment
            await expect(buyTx).to.changeEtherBalances([supplierCMAccount, distributorCMAccount], [0n, 0n]);

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());
        });

        it("Off-chain: should buy a booking token with off-chain payment with zero price correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseUnits("0", 5);

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Off-chain payment marker
            const offChainPaymentMarker = await bookingToken.OFFCHAIN_PAYMENT();

            // address(1)
            const OneAddress = ethers.getAddress("0x0000000000000000000000000000000000000001");

            expect(offChainPaymentMarker).to.equal(OneAddress);

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    offChainPaymentMarker, // off-chain payment marker, address(1)
                    6, // off chain payment currency, 6 == Euro
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    offChainPaymentMarker, // off-chain payment marker, address(1)
                    6, // off chain payment currency, 6 == Euro
                    true, // cancellable
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n, price, OneAddress);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check balances, balances should not change as it is off-chain payment
            await expect(buyTx).to.changeEtherBalances([supplierCMAccount, distributorCMAccount], [0n, 0n]);

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());
        });

        it("Off-chain: should revert with off-chain payment and msg.value > 0", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.5");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Off-chain payment marker
            const offChainPaymentMarker = await bookingToken.OFFCHAIN_PAYMENT();

            // address(1)
            const OneAddress = ethers.getAddress("0x0000000000000000000000000000000000000001");

            expect(offChainPaymentMarker).to.equal(OneAddress);

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    offChainPaymentMarker, // off-chain payment marker, address(1)
                    6, // off chain payment currency, 6 == Euro
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    offChainPaymentMarker, // off-chain payment marker, address(1)
                    6, // off chain payment currency, 6 == Euro
                    true, // cancellable
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/
            // Impersonate the CMAccount contract
            await network.provider.request({
                method: "hardhat_impersonateAccount",
                params: [await distributorCMAccount.getAddress()],
            });

            // Give it some CAM balance
            await network.provider.send("hardhat_setBalance", [
                await distributorCMAccount.getAddress(),
                ethers.toBeHex(price + ethers.parseEther("100")),
            ]);

            // Get the impersonated signer
            const impersonatedSigner = await ethers.getSigner(await distributorCMAccount.getAddress());

            // Try to buy the token with CAM - should revert
            await expect(bookingToken.connect(impersonatedSigner).buyReservedToken(0n, { value: price }))
                .to.be.revertedWithCustomError(bookingToken, "UnexpectedNativePayment")
                .withArgs(price);
        });

        it("ERC20: should buy a booking token correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken, nullUSD } =
                await loadFixture(deployBookingTokenWithNullUSDFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("500");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    nullUSD.getAddress(), // nullUSD address
                    0,
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    nullUSD.getAddress(), // nullUSD address
                    0,
                    false,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount
                .connect(signers.btAdmin)
                .buyBookingToken(0n, price, await nullUSD.getAddress());

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check balances
            // CAM
            await expect(buyTx).to.changeEtherBalances([supplierCMAccount, distributorCMAccount], [0, 0]);
            // Token: NullUSD
            await expect(buyTx).to.changeTokenBalances(
                nullUSD,
                [supplierCMAccount, distributorCMAccount],
                [price, -price],
            );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());
        });

        it("ERC20: should revert if msg.value > 0", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken, nullUSD } =
                await loadFixture(deployBookingTokenWithNullUSDFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("500");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    nullUSD.getAddress(), // nullUSD address
                    0,
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    nullUSD.getAddress(), // nullUSD address
                    0,
                    false,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Impersonate the CMAccount contract
            await network.provider.request({
                method: "hardhat_impersonateAccount",
                params: [await distributorCMAccount.getAddress()],
            });

            // Give it some CAM balance
            await network.provider.send("hardhat_setBalance", [
                await distributorCMAccount.getAddress(),
                ethers.toBeHex(price + ethers.parseEther("100")),
            ]);

            // Get the impersonated signer
            const impersonatedSigner = await ethers.getSigner(await distributorCMAccount.getAddress());

            // Try to buy the token with CAM - should revert
            await expect(bookingToken.connect(impersonatedSigner).buyReservedToken(0n, { value: price }))
                .to.be.revertedWithCustomError(bookingToken, "UnexpectedNativePayment")
                .withArgs(price);
        });

        it("ERC20: should buy a booking token with zero price correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken, nullUSD } =
                await loadFixture(deployBookingTokenWithNullUSDFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    nullUSD.getAddress(), // nullUSD address
                    0,
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    nullUSD.getAddress(), // nullUSD address
                    0,
                    false,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount
                .connect(signers.btAdmin)
                .buyBookingToken(0n, price, await nullUSD.getAddress());

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check balances
            // CAM
            await expect(buyTx).to.changeEtherBalances([supplierCMAccount, distributorCMAccount], [0, 0]);
            // Token: NullUSD
            await expect(buyTx).to.changeTokenBalances(
                nullUSD,
                [supplierCMAccount, distributorCMAccount],
                [price, -price],
            );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());
        });

        it("Native: should revert when trying to buy a booking token reserved for another CMAccount", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    supplierCMAccount.getAddress(), // set reservedFor address to NOT distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    supplierCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                );

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n, price, ethers.ZeroAddress);

            // Check emitted events
            await expect(buyTx)
                .to.be.revertedWithCustomError(bookingToken, "ReservationMismatch")
                .withArgs(supplierCMAccount.getAddress(), distributorCMAccount.getAddress());
        });

        it("Native: should revert if token reservation is expired", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(),
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                );

            // Move time forward and mine a block
            await helpers.time.increaseTo(expirationTimestamp + 1);

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n, price, ethers.ZeroAddress);

            // Check emitted events
            await expect(buyTx)
                .to.be.revertedWithCustomError(bookingToken, "ReservationExpired")
                .withArgs(0n, expirationTimestamp);
        });

        it("Native: should revert if caller is not CMAccount", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            // Try with non-CMAccount address
            await expect(
                bookingToken.connect(signers.otherAccount3).buyReservedToken(0n),
            ).to.be.revertedWithCustomError(bookingToken, "NotCMAccount");
        });
    });

    describe("Transfer", function () {
        it("should revert transfer a booking token if the token is reserved", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            // Try to transfer the token
            await expect(
                supplierCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(await bookingToken.getAddress(), signers.otherAccount1.address, 0n),
            )
                .to.be.revertedWithCustomError(bookingToken, "TokenIsReserved")
                .withArgs(0n, await distributorCMAccount.getAddress());
        });

        it("should transfer a booking token after token is bought", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n, price, ethers.ZeroAddress);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(3); // Bought == 1

            // Set WITHDRAWER_ROLE
            const WITHDRAWER_ROLE = await distributorCMAccount.WITHDRAWER_ROLE();
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(WITHDRAWER_ROLE, signers.withdrawer.address),
            ).to.not.reverted;

            // Try to transfer the token, should not revert
            await expect(
                distributorCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(await bookingToken.getAddress(), signers.otherAccount1.address, 0n),
            )
                .to.emit(bookingToken, "Transfer")
                .withArgs(distributorCMAccount.getAddress(), signers.otherAccount1.address, 0n);

            // BookingToken transferFrom and safeTransferFrom funcs

            // Use booking token to transfer with the new owner
            await expect(
                bookingToken
                    .connect(signers.otherAccount1)
                    .transferFrom(signers.otherAccount1.address, signers.otherAccount2.address, 0n),
            )
                .to.emit(bookingToken, "Transfer")
                .withArgs(signers.otherAccount1.address, signers.otherAccount2.address, 0n);

            await expect(
                bookingToken
                    .connect(signers.otherAccount2)
                    .safeTransferFrom(signers.otherAccount2.address, signers.otherAccount3.address, 0n),
            )
                .to.emit(bookingToken, "Transfer")
                .withArgs(signers.otherAccount2.address, signers.otherAccount3.address, 0n);

            // Try to transfer status unspecified token (checkTransferable func)
            // (this is only possible for non-existing tokens)
            await expect(
                bookingToken
                    .connect(signers.otherAccount3)
                    .safeTransferFrom(signers.otherAccount3.address, signers.otherAccount2.address, 99n),
            ).to.be.revertedWithCustomError(bookingToken, "ERC721NonexistentToken");
        });

        it("should revert transfer if the token is cancelled", async function () {
            const {
                supplierCMAccount,
                distributorCMAccount,
                bookingToken,
                nullUSD,
                tokenWithNativePayment,
                tokenWithNullUSDPayment,
                supplierBookingOperator,
                distributorBookingOperator,
                otherCMAccount,
                otherBookingOperator,
                tokenWithoutBuying,
                tokenWithPassedExpiration,
                offChainPaymentToken,
                offChainPaymentCurrency,
                tokenWithOffChainPayment,
            } = await loadFixture(deployCancellationSupportFixture);

            const supplier = await supplierCMAccount.getAddress();
            const distributor = await distributorCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");
            const paymentToken = ethers.ZeroAddress;
            const cancellationReason = 42;
            const cancellationReasonVersion = 1;

            // INITIATE CANCELLATION PROPOSAL

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            ).to.be.not.reverted;

            // Accept and finalize cancellation proposal
            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .finalizeCancellation(tokenWithNativePayment, refundAmount),
            ).to.be.not.reverted;

            // Check booking token status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(4); // Cancelled == 4

            // Grant WITHDRAWER_ROLE
            const WITHDRAWER_ROLE = await distributorCMAccount.WITHDRAWER_ROLE();
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(WITHDRAWER_ROLE, signers.withdrawer.address),
            ).to.not.reverted;

            // Try to transfer the token, should revert
            await expect(
                distributorCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(
                        await bookingToken.getAddress(),
                        signers.otherAccount1.address,
                        tokenWithNativePayment,
                    ),
            )
                .to.be.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(tokenWithNativePayment, 4); // Cancelled == 4
        });

        it("should withdraw/reject cancellation proposals during transfer", async function () {
            const {
                supplierCMAccount,
                distributorCMAccount,
                bookingToken,
                nullUSD,
                tokenWithNativePayment,
                tokenWithNullUSDPayment,
                supplierBookingOperator,
                distributorBookingOperator,
                otherCMAccount,
                otherBookingOperator,
                tokenWithoutBuying,
                tokenWithPassedExpiration,
                offChainPaymentToken,
                offChainPaymentCurrency,
                tokenWithOffChainPayment,
            } = await loadFixture(deployCancellationSupportFixture);

            const supplier = await supplierCMAccount.getAddress();
            const distributor = await distributorCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");
            const paymentToken = ethers.ZeroAddress;
            const cancellationReason = 42;
            const cancellationReasonVersion = 1;

            // INITIATE CANCELLATION PROPOSAL

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            ).to.be.not.reverted;

            // TRANSFER THE TOKEN

            // Grant withdrawal role
            const WITHDRAWER_ROLE = await distributorCMAccount.WITHDRAWER_ROLE();
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(WITHDRAWER_ROLE, signers.withdrawer.address),
            ).to.not.reverted;

            // Check cancellation proposal status
            expect(await bookingToken.getCancellationProposal(tokenWithNativePayment)).to.deep.equal([
                1n, // Pending == 1
                refundAmount,
                supplier, // initial proposer
                supplier, // current proposer
                false, // ownerAccepted
                true, // supplierAccepted
                0n, // timesCountered
                0n, // timesRejected
            ]);

            await expect(
                distributorCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(
                        await bookingToken.getAddress(),
                        signers.otherAccount1.address,
                        tokenWithNativePayment,
                    ),
            )
                .to.emit(bookingToken, "Transfer")
                .withArgs(distributor, signers.otherAccount1.address, tokenWithNativePayment);

            // Check token ownership
            expect(await bookingToken.ownerOf(tokenWithNativePayment)).to.equal(signers.otherAccount1.address);

            // Check cancellation proposal status
            // Proposal should be rejected
            expect(await bookingToken.getCancellationProposal(tokenWithNativePayment)).to.deep.equal([
                2n, // Rejected == 2
                refundAmount,
                supplier, // initial proposer
                supplier, // current proposer
                false, // ownerAccepted
                true, // supplierAccepted
                0n, // timesCountered
                1n, // timesRejected
            ]);

            // INITIATE CANCELLATION PROPOSAL with Distributor

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .initiateCancellation(
                        tokenWithNullUSDPayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            ).to.be.not.reverted;

            // Check cancellation proposal status
            expect(await bookingToken.getCancellationProposal(tokenWithNullUSDPayment)).to.deep.equal([
                1n, // Pending == 1
                refundAmount,
                distributor, // initial proposer
                distributor, // current proposer
                true, // ownerAccepted
                false, // supplierAccepted
                0n, // timesCountered
                0n, // timesRejected
            ]);

            // TRANSFER THE TOKEN

            await expect(
                distributorCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(
                        await bookingToken.getAddress(),
                        signers.otherAccount1.address,
                        tokenWithNullUSDPayment,
                    ),
            )
                .to.emit(bookingToken, "Transfer")
                .withArgs(distributor, signers.otherAccount1.address, tokenWithNullUSDPayment);

            // Check token ownership
            expect(await bookingToken.ownerOf(tokenWithNullUSDPayment)).to.equal(signers.otherAccount1.address);

            // Check cancellation proposal status
            // Proposal should be rejected
            expect(await bookingToken.getCancellationProposal(tokenWithNullUSDPayment)).to.deep.equal([
                3n, // Withdrawn == 3
                refundAmount,
                distributor, // initial proposer
                distributor, // current proposer
                true, // ownerAccepted
                false, // supplierAccepted
                0n, // timesCountered
                0n, // timesRejected
            ]);
        });
    });

    describe("Expiration", function () {
        it("should record a booking token as expired correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            // Try to expire the token before the expiration timestamp
            await expect(supplierCMAccount.connect(signers.btAdmin).recordExpiration(0n))
                .to.be.revertedWithCustomError(bookingToken, "TokenIsReserved")
                .withArgs(0n, await distributorCMAccount.getAddress());

            // Advance time by 24 hours, token should can be expired after
            await network.provider.send("evm_increaseTime", [24 * 60 * 60]);
            await network.provider.send("evm_mine");

            // Try to expire with non-auth address
            await expect(
                supplierCMAccount.connect(signers.otherAccount1).recordExpiration(0n),
            ).to.be.revertedWithCustomError(bookingToken, "AccessControlUnauthorizedAccount");

            // Expire the token
            await expect(supplierCMAccount.connect(signers.btAdmin).recordExpiration(0n))
                .to.emit(bookingToken, "TokenReservationExpired")
                .withArgs(0n);

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(2); // Expired == 2

            // Try to expire the token again
            await expect(supplierCMAccount.connect(signers.btAdmin).recordExpiration(0n))
                .to.be.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(0n, 2); // RESERVATION_EXPIRED == 2

            // Try to transfer the token, should not revert
            await supplierCMAccount
                .connect(signers.cmAccountAdmin)
                .grantRole(await supplierCMAccount.WITHDRAWER_ROLE(), signers.otherAccount1.address);
            await expect(
                supplierCMAccount
                    .connect(signers.otherAccount1)
                    .transferERC721(
                        bookingToken.getAddress(),
                        signers.otherAccount1.address,
                        0n,
                        distributorCMAccount.getAddress(),
                    ),
            )
                .to.emit(bookingToken, "Transfer")
                .withArgs(supplierCMAccount.getAddress(), signers.otherAccount1.address, 0n);
        });
        it("should revert recording as expired if it's bought already", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0,
                    false,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n, price, ethers.ZeroAddress);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(3); // Bought == 3

            // Try to expire the token, should revert with InvalidTokenStatus
            await expect(distributorCMAccount.connect(signers.btAdmin).recordExpiration(0n))
                .to.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(0n, 3); // Bought == 3
        });
    });
    describe("Cancellation", function () {
        it("should get cancellable flag correctly", async function () {
            const {
                supplierCMAccount,
                distributorCMAccount,
                bookingToken,
                nullUSD,
                tokenWithNativePayment,
                tokenWithNullUSDPayment,
                supplierBookingOperator,
                distributorBookingOperator,
            } = await loadFixture(deployCancellationSupportFixture);

            // Get cancellable flags
            expect(await bookingToken.isCancellable(tokenWithNativePayment)).to.equal(true);
            expect(await bookingToken.isCancellable(tokenWithNullUSDPayment)).to.equal(true);
        });

        it("should revert if not owner or supplier", async function () {
            const {
                supplierCMAccount,
                distributorCMAccount,
                bookingToken,
                nullUSD,
                tokenWithNativePayment,
                tokenWithNullUSDPayment,
                supplierBookingOperator,
                distributorBookingOperator,
                otherCMAccount,
                otherBookingOperator,
            } = await loadFixture(deployCancellationSupportFixture);

            const token_id = tokenWithNativePayment;
            const refundAmount = ethers.parseEther("0.045");
            const reason = 42;
            const reasonVersion = 1;

            await expect(
                otherCMAccount
                    .connect(otherBookingOperator)
                    .initiateCancellation(token_id, refundAmount, reason, reasonVersion),
            ).to.revertedWithCustomError(bookingToken, "NotOwnerOrSupplier");

            await expect(
                otherCMAccount.connect(otherBookingOperator).acceptCancellation(token_id, refundAmount),
            ).to.revertedWithCustomError(bookingToken, "NotOwnerOrSupplier");

            await expect(
                otherCMAccount
                    .connect(otherBookingOperator)
                    .counterCancellation(token_id, refundAmount, reason, reasonVersion),
            ).to.revertedWithCustomError(bookingToken, "NotOwnerOrSupplier");

            await expect(
                otherCMAccount.connect(otherBookingOperator).withdrawCancellation(token_id, reason, reasonVersion),
            ).to.revertedWithCustomError(bookingToken, "NotOwnerOrSupplier");

            await expect(
                otherCMAccount.connect(otherBookingOperator).rejectCancellation(token_id, reason, reasonVersion),
            ).to.revertedWithCustomError(bookingToken, "NotOwnerOrSupplier");

            // Special case for finalize
            await expect(
                otherCMAccount.connect(otherBookingOperator).finalizeCancellation(token_id, refundAmount),
            ).to.revertedWithCustomError(bookingToken, "OnlySupplierCanFinalizeCancellation");
        });

        it("should revert if caller is not CMAccount", async function () {
            const {
                supplierCMAccount,
                distributorCMAccount,
                bookingToken,
                nullUSD,
                tokenWithNativePayment,
                tokenWithNullUSDPayment,
                supplierBookingOperator,
                distributorBookingOperator,
                otherCMAccount,
                otherBookingOperator,
            } = await loadFixture(deployCancellationSupportFixture);

            const token_id = tokenWithNativePayment;
            const refundAmount = ethers.parseEther("0.045");
            const reason = 42;
            const reasonVersion = 1;

            await expect(
                bookingToken
                    .connect(signers.otherAccount3)
                    .initiateCancellation(token_id, refundAmount, reason, reasonVersion),
            ).to.revertedWithCustomError(bookingToken, "NotCMAccount");

            await expect(
                bookingToken.connect(signers.otherAccount3).acceptCancellation(token_id, refundAmount),
            ).to.revertedWithCustomError(bookingToken, "NotCMAccount");

            await expect(
                bookingToken
                    .connect(signers.otherAccount3)
                    .counterCancellation(token_id, refundAmount, reason, reasonVersion),
            ).to.revertedWithCustomError(bookingToken, "NotCMAccount");

            await expect(
                bookingToken.connect(signers.otherAccount3).withdrawCancellation(token_id, reason, reasonVersion),
            ).to.revertedWithCustomError(bookingToken, "NotCMAccount");

            await expect(
                bookingToken.connect(signers.otherAccount3).rejectCancellation(token_id, reason, reasonVersion),
            ).to.revertedWithCustomError(bookingToken, "NotCMAccount");

            // Special case for finalize
            await expect(
                bookingToken.connect(signers.otherAccount3).finalizeCancellation(token_id, refundAmount),
            ).to.revertedWithCustomError(bookingToken, "NotCMAccount");
        });

        it("should initiate a cancellation proposal correctly", async function () {
            const {
                supplierCMAccount,
                distributorCMAccount,
                bookingToken,
                nullUSD,
                tokenWithNativePayment,
                tokenWithNullUSDPayment,
                supplierBookingOperator,
                distributorBookingOperator,
                otherCMAccount,
                otherBookingOperator,
                tokenWithoutBuying,
                tokenWithPassedExpiration,
            } = await loadFixture(deployCancellationSupportFixture);

            // Try to cancel the token
            const proposer = await distributorCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");
            const cancellationReason = 42;
            const cancellationReasonVersion = 1;

            // INITIATE :: OWNER

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNativePayment,
                    proposer, // initial proposer
                    proposer, // current proposer
                    refundAmount,
                    true, // ownerAccepted
                    false, // supplierAccepted
                    0, // timesCountered
                    0, // timesRejected
                )
                .to.emit(bookingToken, "CancellationReasons")
                .withArgs(
                    tokenWithNativePayment,
                    cancellationReason, //
                    cancellationReasonVersion, //
                    0, // proposal.rejectionReason,
                    0, // proposal.rejectionVersion,
                    0, // proposal.counterReason,
                    0, // proposal.counterVersion,
                    0, // proposal.withdrawalReason,
                    0, // proposal.withdrawalVersion
                );

            // Check cancellation proposal state
            expect(await bookingToken.getCancellationProposal(tokenWithNativePayment)).to.deep.equal([
                1n, // Pending == 1
                refundAmount,
                proposer, // initial proposer
                proposer, // current proposer
                true, // ownerAccepted
                false, // supplierAccepted
                0n, // timesCountered
                0n, // timesRejected
            ]);

            // Check cancellation proposal reasons
            expect(await bookingToken.getCancellationReasons(tokenWithNativePayment)).to.deep.equal([
                cancellationReason,
                cancellationReasonVersion,
                0n, // proposal.rejectionReason,
                0n, // proposal.rejectionVersion,
                0n, // proposal.counterReason,
                0n, // proposal.counterVersion,
                0n, // proposal.withdrawalReason,
                0n, // proposal.withdrawalVersion
            ]);

            // INITIATE :: SUPPLIER

            const supplier = await supplierCMAccount.getAddress();

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .initiateCancellation(
                        tokenWithNullUSDPayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNullUSDPayment,
                    supplier, // initial proposer
                    supplier, // current proposer
                    refundAmount,
                    false, // ownerAccepted
                    true, // supplierAccepted
                    0, // timesCountered
                    0, // timesRejected
                )
                .to.emit(bookingToken, "CancellationReasons")
                .withArgs(
                    tokenWithNullUSDPayment,
                    cancellationReason, //
                    cancellationReasonVersion, //
                    0, // proposal.rejectionReason,
                    0, // proposal.rejectionVersion,
                    0, // proposal.counterReason,
                    0, // proposal.counterVersion,
                    0, // proposal.withdrawalReason,
                    0, // proposal.withdrawalVersion
                );

            // Check cancellation proposal state
            expect(await bookingToken.getCancellationProposal(tokenWithNullUSDPayment)).to.deep.equal([
                1n, // Pending == 1
                refundAmount,
                supplier, // initial proposer
                supplier, // current proposer
                false, // ownerAccepted
                true, // supplierAccepted
                0n, // timesCountered
                0n, // timesRejected
            ]);

            // Check cancellation proposal reasons
            expect(await bookingToken.getCancellationReasons(tokenWithNullUSDPayment)).to.deep.equal([
                cancellationReason,
                cancellationReasonVersion,
                0n, // proposal.rejectionReason,
                0n, // proposal.rejectionVersion,
                0n, // proposal.counterReason,
                0n, // proposal.counterVersion,
                0n, // proposal.withdrawalReason,
                0n, // proposal.withdrawalVersion
            ]);

            // REVERTS: TRY TO INIT EXISTING

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidCancellationProposalStatus")
                .withArgs(tokenWithNativePayment, 1n); // PROPOSAL: PENDING: 1

            // REVERTS: TRY TO INIT WITH NON-BOUGHT TOKEN

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .initiateCancellation(
                        tokenWithoutBuying,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(tokenWithoutBuying, 1n); // RESERVED: 1

            // REVERTS: TRY TO INIT WITH NON CM ACCOUNT

            await expect(
                bookingToken
                    .connect(otherBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.revertedWithCustomError(bookingToken, "NotCMAccount")
                .withArgs(otherBookingOperator.address);
        });

        it("should accept a cancellation proposal correctly", async function () {
            const {
                supplierCMAccount,
                distributorCMAccount,
                bookingToken,
                nullUSD,
                tokenWithNativePayment,
                tokenWithNullUSDPayment,
                supplierBookingOperator,
                distributorBookingOperator,
                otherCMAccount,
                otherBookingOperator,
                tokenWithoutBuying,
                tokenWithPassedExpiration,
            } = await loadFixture(deployCancellationSupportFixture);

            // INIT CANCELLATION PROPOSAL

            // Try to cancel the token
            const proposer = await supplierCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");
            const cancellationReason = 42;
            const cancellationReasonVersion = 1;

            // INITIATE

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNativePayment,
                    proposer, // initial proposer
                    proposer, // current proposer
                    refundAmount,
                    false, // ownerAccepted
                    true, // supplierAccepted
                    0, // timesCountered
                    0, // timesRejected
                );

            // REVERTS: TRY ACCEPT WITH DIFFERENT REFUND AMOUNT

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .acceptCancellation(tokenWithNativePayment, refundAmount + 1n),
            )
                .to.revertedWithCustomError(bookingToken, "IncorrectRefundAmount")
                .withArgs(tokenWithNativePayment, refundAmount, refundAmount + 1n);

            // ACCEPT

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .acceptCancellation(tokenWithNativePayment, refundAmount),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNativePayment,
                    proposer, // initial proposer
                    proposer, // current proposer
                    refundAmount,
                    true, // ownerAccepted
                    true, // supplierAccepted
                    0, // timesCountered
                    0, // timesRejected
                );

            // Check cancellation proposal state
            expect(await bookingToken.getCancellationProposal(tokenWithNativePayment)).to.deep.equal([
                1n, // Pending == 1
                refundAmount,
                proposer, // initial proposer
                proposer, // current proposer
                true, // ownerAccepted
                true, // supplierAccepted
                0n, // timesCountered
                0n, // timesRejected
            ]);

            // REVERTS: TRY TO ACCEPT WITH NON-BOUGHT TOKEN
            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .acceptCancellation(tokenWithoutBuying, refundAmount),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(tokenWithoutBuying, 1n); // TOKEN: RESERVED: 1

            // REVERTS: TRY TO ACCEPT WITH NON-INITIATED TOKEN

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .acceptCancellation(tokenWithNullUSDPayment, refundAmount),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidCancellationProposalStatus")
                .withArgs(tokenWithNullUSDPayment, 0n); // PROPOSAL: NO_PROPOSAL: 0

            // ACCEPT WITH SUPPLIER

            // Initiate with distributor first
            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .initiateCancellation(
                        tokenWithNullUSDPayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNullUSDPayment,
                    await distributorCMAccount.getAddress(), // initial proposer
                    await distributorCMAccount.getAddress(), // current proposer
                    refundAmount,
                    true, // ownerAccepted
                    false, // supplierAccepted
                    0, // timesCountered
                    0, // timesRejected
                );

            // Accept with supplier
            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .acceptCancellation(tokenWithNullUSDPayment, refundAmount),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNullUSDPayment,
                    await distributorCMAccount.getAddress(), // initial proposer
                    await distributorCMAccount.getAddress(), // current proposer
                    refundAmount,
                    true, // ownerAccepted
                    true, // supplierAccepted
                    0, // timesCountered
                    0, // timesRejected
                );
        });

        it("should counter a cancellation proposal correctly", async function () {
            const {
                supplierCMAccount,
                distributorCMAccount,
                bookingToken,
                nullUSD,
                tokenWithNativePayment,
                tokenWithNullUSDPayment,
                supplierBookingOperator,
                distributorBookingOperator,
                otherCMAccount,
                otherBookingOperator,
                tokenWithoutBuying,
                tokenWithPassedExpiration,
            } = await loadFixture(deployCancellationSupportFixture);

            // INIT CANCELLATION PROPOSAL

            // Try to cancel the token
            const proposer = await supplierCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");
            const cancellationReason = 42;
            const cancellationReasonVersion = 1;

            const counterCancellationReason = 43;
            const counterCancellationReasonVersion = 2;
            const counterRefundAmount = ethers.parseEther("0.05");

            // REVERTS: TRY WITH NON-BOUGHT TOKEN

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .counterCancellation(
                        tokenWithoutBuying,
                        counterRefundAmount,
                        counterCancellationReason,
                        counterCancellationReasonVersion,
                    ),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(tokenWithoutBuying, 1n); // TOKEN: RESERVED: 1

            // REVERTS: TRY TO COUNTER NON-INITIATED PROPOSAL

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .counterCancellation(
                        tokenWithNativePayment,
                        counterRefundAmount,
                        counterCancellationReason,
                        counterCancellationReasonVersion,
                    ),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidCancellationProposalStatus")
                .withArgs(tokenWithNativePayment, 0n); // PROPOSAL: NO_PROPOSAL: 0

            // INITIATE

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNativePayment,
                    proposer, // initial proposer
                    proposer, // current proposer
                    refundAmount,
                    false, // ownerAccepted
                    true, // supplierAccepted
                    0, // timesCountered
                    0, // timesRejected
                );

            // COUNTER

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .counterCancellation(
                        tokenWithNativePayment,
                        counterRefundAmount,
                        counterCancellationReason,
                        counterCancellationReasonVersion,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNativePayment,
                    proposer, // initial proposer
                    await distributorCMAccount.getAddress(), // new current proposer is distributor CMAccount
                    counterRefundAmount,
                    true, // ownerAccepted
                    false, // supplierAccepted
                    1, // timesCountered
                    0, // timesRejected
                )
                .to.emit(bookingToken, "CancellationReasons")
                .withArgs(
                    tokenWithNativePayment,
                    cancellationReason, // proposal.counterReason,
                    cancellationReasonVersion, // proposal.counterVersion,
                    0, // proposal.rejectionReason,
                    0, // proposal.rejectionVersion,
                    counterCancellationReason, //
                    counterCancellationReasonVersion, //
                    0, // proposal.withdrawalReason,
                    0, // proposal.withdrawalVersion
                );
        });

        it("should withdraw a cancellation proposal correctly", async function () {
            const {
                supplierCMAccount,
                distributorCMAccount,
                bookingToken,
                nullUSD,
                tokenWithNativePayment,
                tokenWithNullUSDPayment,
                supplierBookingOperator,
                distributorBookingOperator,
                otherCMAccount,
                otherBookingOperator,
                tokenWithoutBuying,
                tokenWithPassedExpiration,
            } = await loadFixture(deployCancellationSupportFixture);

            const proposer = await supplierCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");
            const cancellationReason = 42;
            const cancellationReasonVersion = 1;

            const withdrawalReason = 44;
            const withdrawalReasonVersion = 3;

            // REVERTS: TRY TO WITHDRAW NON-INITIATED PROPOSAL

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .withdrawCancellation(tokenWithNativePayment, withdrawalReason, withdrawalReasonVersion),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidCancellationProposalStatus")
                .withArgs(tokenWithNativePayment, 0n); // PROPOSAL: NO_PROPOSAL: 0

            // INITIATE

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNativePayment,
                    proposer, // initial proposer
                    proposer, // current proposer
                    refundAmount,
                    false, // ownerAccepted
                    true, // supplierAccepted
                    0, // timesCountered
                    0, // timesRejected
                );

            // REVERTS: TRY WITH NON-BOUGHT TOKEN

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .withdrawCancellation(tokenWithoutBuying, withdrawalReason, withdrawalReasonVersion),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(tokenWithoutBuying, 1n); // TOKEN: RESERVED: 1

            // REVERTS: TRY TO WITHDRAW WITH NON-CURRENT PROPOSER

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .withdrawCancellation(tokenWithNativePayment, withdrawalReason, withdrawalReasonVersion),
            )
                .to.revertedWithCustomError(bookingToken, "OnlyCurrentProposerCanWithdrawCancellation")
                .withArgs(tokenWithNativePayment);

            // WITHDRAW

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .withdrawCancellation(tokenWithNativePayment, withdrawalReason, withdrawalReasonVersion),
            )
                .to.emit(bookingToken, "CancellationWithdrawn")
                .withArgs(tokenWithNativePayment, withdrawalReason, withdrawalReasonVersion);

            // Check cancellation proposal state
            expect(await bookingToken.getCancellationProposal(tokenWithNativePayment)).to.deep.equal([
                3n, // Withdrawn == 3
                refundAmount,
                proposer, // initial proposer
                proposer, // current proposer
                false, // ownerAccepted
                true, // supplierAccepted
                0n, // timesCountered
                0n, // timesRejected
            ]);

            // Check cancellation proposal reasons
            expect(await bookingToken.getCancellationReasons(tokenWithNativePayment)).to.deep.equal([
                cancellationReason,
                cancellationReasonVersion,
                0n, // proposal.rejectionReason,
                0n, // proposal.rejectionVersion,
                0n, // proposal.counterReason,
                0n, // proposal.counterVersion,
                withdrawalReason, // proposal.withdrawalReason,
                withdrawalReasonVersion, // proposal.withdrawalVersion
            ]);
        });

        it("should reject a cancellation proposal correctly", async function () {
            const {
                supplierCMAccount,
                distributorCMAccount,
                bookingToken,
                nullUSD,
                tokenWithNativePayment,
                tokenWithNullUSDPayment,
                supplierBookingOperator,
                distributorBookingOperator,
                otherCMAccount,
                otherBookingOperator,
                tokenWithoutBuying,
                tokenWithPassedExpiration,
            } = await loadFixture(deployCancellationSupportFixture);

            const proposer = await supplierCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");
            const cancellationReason = 42;
            const cancellationReasonVersion = 1;

            const rejectionReason = 44;
            const rejectionReasonVersion = 3;

            // REVERTS: TRY TO REJECT NON-INITIATED PROPOSAL

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .rejectCancellation(tokenWithNativePayment, rejectionReason, rejectionReasonVersion),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidCancellationProposalStatus")
                .withArgs(tokenWithNativePayment, 0n); // PROPOSAL: NO_PROPOSAL: 0

            // INITIATE

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNativePayment,
                    proposer, // initial proposer
                    proposer, // current proposer
                    refundAmount,
                    false, // ownerAccepted
                    true, // supplierAccepted
                    0, // timesCountered
                    0, // timesRejected
                );

            // REVERTS: TRY WITH NON-BOUGHT TOKEN

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .rejectCancellation(tokenWithoutBuying, rejectionReason, rejectionReasonVersion),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(tokenWithoutBuying, 1n); // TOKEN: RESERVED: 1

            // REVERTS: TRY TO REJECT WITH PROPOSER

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .rejectCancellation(tokenWithNativePayment, rejectionReason, rejectionReasonVersion),
            )
                .to.revertedWithCustomError(bookingToken, "ProposerCanNotRejectCancellation")
                .withArgs(tokenWithNativePayment);

            // REJECT

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .rejectCancellation(tokenWithNativePayment, rejectionReason, rejectionReasonVersion),
            )
                .to.emit(bookingToken, "CancellationRejected")
                .withArgs(tokenWithNativePayment, rejectionReason, rejectionReasonVersion);

            // Check cancellation proposal state
            expect(await bookingToken.getCancellationProposal(tokenWithNativePayment)).to.deep.equal([
                2n, // Rejected == 2
                refundAmount,
                proposer, // initial proposer
                proposer, // current proposer
                false, // ownerAccepted
                true, // supplierAccepted
                0n, // timesCountered
                1n, // timesRejected
            ]);

            // Check cancellation proposal reasons
            expect(await bookingToken.getCancellationReasons(tokenWithNativePayment)).to.deep.equal([
                cancellationReason,
                cancellationReasonVersion,
                rejectionReason, // proposal.rejectionReason,
                rejectionReasonVersion, // proposal.rejectionVersion,
                0n, // proposal.counterReason,
                0n, // proposal.counterVersion,
                0n, // proposal.withdrawalReason,
                0n, // proposal.withdrawalVersion
            ]);
        });

        it("should reinitialize a cancellation proposal correctly", async function () {
            const {
                supplierCMAccount,
                distributorCMAccount,
                bookingToken,
                nullUSD,
                tokenWithNativePayment,
                tokenWithNullUSDPayment,
                supplierBookingOperator,
                distributorBookingOperator,
                otherCMAccount,
                otherBookingOperator,
                tokenWithoutBuying,
                tokenWithPassedExpiration,
            } = await loadFixture(deployCancellationSupportFixture);

            const supplier = await supplierCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");
            const cancellationReason = 42;
            const cancellationReasonVersion = 1;

            const rejectionReason = 44;
            const rejectionReasonVersion = 3;

            // INITIATE

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNativePayment,
                    supplier, // initial proposer
                    supplier, // current proposer
                    refundAmount,
                    false, // ownerAccepted
                    true, // supplierAccepted
                    0, // timesCountered
                    0, // timesRejected
                );

            // REJECT (So we can re-init later)

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .rejectCancellation(tokenWithNativePayment, rejectionReason, rejectionReasonVersion),
            )
                .to.emit(bookingToken, "CancellationRejected")
                .withArgs(tokenWithNativePayment, rejectionReason, rejectionReasonVersion);

            // Check cancellation proposal state
            expect(await bookingToken.getCancellationProposal(tokenWithNativePayment)).to.deep.equal([
                2n, // Rejected == 2
                refundAmount,
                supplier, // initial proposer
                supplier, // current proposer
                false, // ownerAccepted
                true, // supplierAccepted
                0n, // timesCountered
                1n, // timesRejected
            ]);

            // Check cancellation proposal reasons
            expect(await bookingToken.getCancellationReasons(tokenWithNativePayment)).to.deep.equal([
                cancellationReason,
                cancellationReasonVersion,
                rejectionReason, // proposal.rejectionReason,
                rejectionReasonVersion, // proposal.rejectionVersion,
                0n, // proposal.counterReason,
                0n, // proposal.counterVersion,
                0n, // proposal.withdrawalReason,
                0n, // proposal.withdrawalVersion
            ]);

            // RE-INIT with DISTRIBUTOR

            const newRefundAmount = ethers.parseEther("0.05");
            const newCancellationReason = cancellationReason + 1;
            const newCancellationReasonVersion = cancellationReasonVersion + 1;
            const distributor = await distributorCMAccount.getAddress();

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        newRefundAmount,
                        newCancellationReason,
                        newCancellationReasonVersion,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNativePayment,
                    supplier, // initial proposer
                    distributor, // current proposer
                    newRefundAmount,
                    true, // ownerAccepted
                    false, // supplierAccepted
                    0, // timesCountered
                    1n, // timesRejected
                );

            // Check cancellation proposal state
            expect(await bookingToken.getCancellationProposal(tokenWithNativePayment)).to.deep.equal([
                1n, // Pending == 1
                newRefundAmount,
                supplier, // initial proposer
                distributor, // current proposer
                true, // ownerAccepted
                false, // supplierAccepted
                0n, // timesCountered
                1n, // timesRejected
            ]);

            // Check cancellation proposal reasons
            expect(await bookingToken.getCancellationReasons(tokenWithNativePayment)).to.deep.equal([
                newCancellationReason,
                newCancellationReasonVersion,
                0n, // proposal.rejectionReason,
                0n, // proposal.rejectionVersion,
                0n, // proposal.counterReason,
                0n, // proposal.counterVersion,
                0n, // proposal.withdrawalReason,
                0n, // proposal.withdrawalVersion
            ]);

            // WITHDRAW (So we can re-init later)

            const withdrawalReason = 43;
            const withdrawalReasonVersion = 2;

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .withdrawCancellation(tokenWithNativePayment, withdrawalReason, withdrawalReasonVersion),
            )
                .to.emit(bookingToken, "CancellationWithdrawn")
                .withArgs(tokenWithNativePayment, withdrawalReason, withdrawalReasonVersion);

            // RE-INIT with SUPPLIER from a WITHDRAWN PROPOSAL

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        newRefundAmount + 10n,
                        newCancellationReason + 10,
                        newCancellationReasonVersion + 10,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNativePayment,
                    supplier, // initial proposer
                    supplier, // current proposer
                    newRefundAmount + 10n,
                    false, // ownerAccepted
                    true, // supplierAccepted
                    0, // timesCountered
                    1n, // timesRejected
                );

            // Check cancellation proposal state
            expect(await bookingToken.getCancellationProposal(tokenWithNativePayment)).to.deep.equal([
                1n, // Pending == 1
                newRefundAmount + 10n,
                supplier, // initial proposer
                supplier, // current proposer
                false, // ownerAccepted
                true, // supplierAccepted
                0n, // timesCountered
                1n, // timesRejected
            ]);

            // Check cancellation proposal reasons
            expect(await bookingToken.getCancellationReasons(tokenWithNativePayment)).to.deep.equal([
                newCancellationReason + 10,
                newCancellationReasonVersion + 10,
                0n, // proposal.rejectionReason,
                0n, // proposal.rejectionVersion,
                0n, // proposal.counterReason,
                0n, // proposal.counterVersion,
                0n, // proposal.withdrawalReason,
                0n, // proposal.withdrawalVersion
            ]);

            // COUNTER Proposal so we test timesCountered

            const counterReason = 43;
            const counterReasonVersion = 2;

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .counterCancellation(tokenWithNativePayment, refundAmount, counterReason, counterReasonVersion),
            ).to.not.be.reverted;

            // Check cancellation proposal state
            expect(await bookingToken.getCancellationProposal(tokenWithNativePayment)).to.deep.equal([
                1n, // Pending == 1
                refundAmount,
                supplier, // initial proposer
                distributor, // current proposer
                true, // ownerAccepted
                false, // supplierAccepted
                1n, // timesCountered
                1n, // timesRejected
            ]);

            // COUNTER again with the supplier

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .counterCancellation(
                        tokenWithNativePayment,
                        refundAmount + 11n,
                        counterReason,
                        counterReasonVersion,
                    ),
            ).to.not.be.reverted;

            // Check cancellation proposal state
            expect(await bookingToken.getCancellationProposal(tokenWithNativePayment)).to.deep.equal([
                1n, // Pending == 1
                refundAmount + 11n,
                supplier, // initial proposer
                supplier, // current proposer
                false, // ownerAccepted
                true, // supplierAccepted
                2n, // timesCountered
                1n, // timesRejected
            ]);

            // REJECT Proposal so we test timesRejected

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .rejectCancellation(tokenWithNativePayment, rejectionReason, rejectionReasonVersion),
            ).to.not.be.reverted;

            // Check cancellation proposal state
            expect(await bookingToken.getCancellationProposal(tokenWithNativePayment)).to.deep.equal([
                2n, // Rejected == 2
                refundAmount + 11n,
                supplier, // initial proposer
                supplier, // current proposer
                false, // ownerAccepted
                true, // supplierAccepted
                2n, // timesCountered
                2n, // timesRejected
            ]);

            // Check cancellation proposal reasons
            expect(await bookingToken.getCancellationReasons(tokenWithNativePayment)).to.deep.equal([
                newCancellationReason + 10,
                newCancellationReasonVersion + 10,
                rejectionReason, // proposal.rejectionReason,
                rejectionReasonVersion, // proposal.rejectionVersion,
                counterReason, // proposal.counterReason,
                counterReasonVersion, // proposal.counterVersion,
                0n, // proposal.withdrawalReason,
                0n, // proposal.withdrawalVersion
            ]);

            // RE-INIT to check timesCountered and timesRejected

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        newRefundAmount,
                        newCancellationReason,
                        newCancellationReasonVersion,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNativePayment,
                    supplier, // initial proposer
                    supplier, // current proposer
                    newRefundAmount,
                    false, // ownerAccepted
                    true, // supplierAccepted
                    2n, // timesCountered
                    2n, // timesRejected
                );

            // Check cancellation proposal state
            expect(await bookingToken.getCancellationProposal(tokenWithNativePayment)).to.deep.equal([
                1n, // Pending == 1
                newRefundAmount,
                supplier, // initial proposer
                supplier, // current proposer
                false, // ownerAccepted
                true, // supplierAccepted
                2n, // timesCountered
                2n, // timesRejected
            ]);

            // Check cancellation proposal reasons
            expect(await bookingToken.getCancellationReasons(tokenWithNativePayment)).to.deep.equal([
                newCancellationReason,
                newCancellationReasonVersion,
                0n, // proposal.rejectionReason,
                0n, // proposal.rejectionVersion,
                0n, // proposal.counterReason,
                0n, // proposal.counterVersion,
                0n, // proposal.withdrawalReason,
                0n, // proposal.withdrawalVersion
            ]);
        });

        it("should finalize a cancellation proposal correctly", async function () {
            const {
                supplierCMAccount,
                distributorCMAccount,
                bookingToken,
                nullUSD,
                tokenWithNativePayment,
                tokenWithNullUSDPayment,
                supplierBookingOperator,
                distributorBookingOperator,
                otherCMAccount,
                otherBookingOperator,
                tokenWithoutBuying,
                tokenWithPassedExpiration,
                offChainPaymentToken,
                offChainPaymentCurrency,
                tokenWithOffChainPayment,
            } = await loadFixture(deployCancellationSupportFixture);

            const supplier = await supplierCMAccount.getAddress();
            const distributor = await distributorCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");
            const paymentToken = ethers.ZeroAddress;
            const cancellationReason = 42;
            const cancellationReasonVersion = 1;

            // REVERTS: TRY WITH NON-BOUGHT TOKEN

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .finalizeCancellation(tokenWithoutBuying, refundAmount),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(tokenWithoutBuying, 1n); // RESERVED: 1

            // REVERTS: TRY TO FINALIZE NON-INITIATED PROPOSAL

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .finalizeCancellation(tokenWithNativePayment, refundAmount),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidCancellationProposalStatus")
                .withArgs(tokenWithNativePayment, 0n); // PROPOSAL: NO_PROPOSAL: 0

            // REVERTS: TRY TO FINALIZE with NON-SUPPLIER

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .finalizeCancellation(tokenWithNativePayment, refundAmount),
            )
                .to.revertedWithCustomError(bookingToken, "OnlySupplierCanFinalizeCancellation")
                .withArgs(tokenWithNativePayment);

            // INITIATE

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNativePayment,
                    supplier, // initial proposer
                    supplier, // current proposer
                    refundAmount,
                    false, // ownerAccepted
                    true, // supplierAccepted
                    0, // timesCountered
                    0, // timesRejected
                );

            // REVERTS: TRY TO FINALIZE without OWNER ACCEPTED

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .finalizeCancellation(tokenWithNativePayment, refundAmount),
            )
                .to.revertedWithCustomError(bookingToken, "OwnerNotAcceptedCancellation")
                .withArgs(tokenWithNativePayment);

            // ACCEPT PROPOSAL

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .acceptCancellation(tokenWithNativePayment, refundAmount),
            ).to.not.be.reverted;

            // REVERTS: TRY TO FINALIZE with INCORRECT REFUND AMOUNT

            const incorrectRefundAmount = refundAmount + 1n;

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .finalizeCancellation(tokenWithNativePayment, incorrectRefundAmount),
            )
                .to.revertedWithCustomError(bookingToken, "IncorrectRefundAmount")
                .withArgs(tokenWithNativePayment, refundAmount, incorrectRefundAmount);

            // REVERTS: TRY TO FINALIZE with NON-SUPPLIER (new state PENDING)

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .finalizeCancellation(tokenWithNativePayment, refundAmount),
            )
                .to.revertedWithCustomError(bookingToken, "OnlySupplierCanFinalizeCancellation")
                .withArgs(tokenWithNativePayment);

            // FINALIZE

            const finalizeTx = supplierCMAccount
                .connect(supplierBookingOperator)
                .finalizeCancellation(tokenWithNativePayment, refundAmount);

            await expect(finalizeTx).to.emit(bookingToken, "CancellationFinalized").withArgs(tokenWithNativePayment);

            await expect(finalizeTx).to.changeEtherBalances(
                [supplier, distributor, bookingToken, supplierBookingOperator, distributorBookingOperator],
                [-refundAmount, refundAmount, 0n, 0n, 0n],
            );

            // TRY TO INIT A FINALIZED PROPOSAL

            await expect(
                supplierCMAccount
                    .connect(supplierBookingOperator)
                    .initiateCancellation(
                        tokenWithNativePayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(tokenWithNativePayment, 4n); // PROPOSAL: FINALIZED: 4

            // INIT with DISTRIBUTOR with ERC20

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .initiateCancellation(
                        tokenWithNullUSDPayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithNullUSDPayment,
                    distributor, // initial proposer
                    distributor, // current proposer
                    refundAmount,
                    true, // ownerAccepted
                    false, // supplierAccepted
                    0, // timesCountered
                    0, // timesRejected
                );

            // FINALIZE with SUPPLIER

            const nullUSDPaymentToken = await nullUSD.getAddress();

            const nullUSDFinalizeTx = supplierCMAccount
                .connect(supplierBookingOperator)
                .finalizeCancellation(tokenWithNullUSDPayment, refundAmount);

            await expect(nullUSDFinalizeTx)
                .to.emit(bookingToken, "CancellationFinalized")
                .withArgs(tokenWithNullUSDPayment);

            await expect(nullUSDFinalizeTx).to.changeTokenBalances(
                nullUSD,
                [supplier, distributor, bookingToken, supplierBookingOperator, distributorBookingOperator],
                [-refundAmount, refundAmount, 0n, 0n, 0n],
            );

            // TEST Transfer

            // Grant withdrawal role
            const WITHDRAWER_ROLE = await distributorCMAccount.WITHDRAWER_ROLE();
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(WITHDRAWER_ROLE, signers.withdrawer.address),
            ).to.not.reverted;

            await expect(
                distributorCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(
                        await bookingToken.getAddress(),
                        signers.otherAccount1.address,
                        tokenWithNullUSDPayment,
                    ),
            )
                .to.be.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(tokenWithNullUSDPayment, 4); // CANCELLED: 4

            // Test record expiration
            await expect(
                supplierCMAccount.connect(supplierBookingOperator).recordExpiration(tokenWithNullUSDPayment),
            ).to.be.revertedWithCustomError(bookingToken, "InvalidTokenStatus");

            // INIT with DISTRIBUTOR with OFFCHAIN PAYMENT

            await expect(
                distributorCMAccount
                    .connect(distributorBookingOperator)
                    .initiateCancellation(
                        tokenWithOffChainPayment,
                        refundAmount,
                        cancellationReason,
                        cancellationReasonVersion,
                    ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(
                    tokenWithOffChainPayment,
                    distributor, // initial proposer
                    distributor, // current proposer
                    refundAmount,
                    true, // ownerAccepted
                    false, // supplierAccepted
                    0, // timesCountered
                    0, // timesRejected
                );

            // FINALIZE with SUPPLIER

            const offChainFinalizeTx = supplierCMAccount
                .connect(supplierBookingOperator)
                .finalizeCancellation(tokenWithOffChainPayment, refundAmount);

            await expect(offChainFinalizeTx)
                .to.emit(bookingToken, "CancellationFinalized")
                .withArgs(tokenWithOffChainPayment);

            await expect(offChainFinalizeTx).to.changeEtherBalances(
                [supplier, distributor, bookingToken, supplierBookingOperator, distributorBookingOperator],
                [0n, 0n, 0n, 0n, 0n], // Off chain payment, should not change any balances
            );
        });
    });
});
