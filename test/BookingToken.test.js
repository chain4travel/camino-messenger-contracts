const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers } = require("hardhat");

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
                    ethers.ZeroAddress, // zero address
                    true,
                ),
            )
                .to.be.revertedWithCustomError(bookingToken, "NotCMAccount") // Caller is not a CMAccount
                .withArgs(signers.btAdmin.address);
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
                    ethers.ZeroAddress, // zero address
                    true,
                ),
            )
                .to.be.revertedWithCustomError(bookingToken, "NotCMAccount")
                .withArgs(signers.otherAccount1.address); // reservedFor address
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
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            // Check cancellable flag
            expect(await bookingToken.isCancellable(0n)).to.equal(true);

            // Mint again to make sure the token id is incremented
            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // zero address
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    1n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // zero address
                    true,
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
                    true,
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
                    true,
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
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                    true,
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
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check balances
            await expect(buyTx).to.changeEtherBalances([supplierCMAccount, distributorCMAccount], [price, -price]);

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(3); // Bought == 3
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
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                    true,
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
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check balances
            await expect(buyTx).to.changeEtherBalances([supplierCMAccount, distributorCMAccount], [price, -price]);

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());
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
                    nullUSD.getAddress(), // nullUSD address
                    true,
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
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

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
                    nullUSD.getAddress(), // nullUSD address
                    true,
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
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

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
                    ethers.ZeroAddress, // zero address
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    supplierCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // zero address
                    true,
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
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx)
                .to.be.revertedWithCustomError(bookingToken, "ReservationMismatch")
                .withArgs(supplierCMAccount.getAddress(), distributorCMAccount.getAddress());
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
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                    true,
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
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                    true,
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
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

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
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                    true,
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

            // Expire the token
            await expect(supplierCMAccount.connect(signers.btAdmin).recordExpiration(0n))
                .to.emit(bookingToken, "TokenExpired")
                .withArgs(0n);

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(2); // Expired == 2
        });
        it("should revert recording a as expired if it's bought already", async function () {
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
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                    true,
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
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

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
        it("should set and get cancellable flag correctly", async function () {
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
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            // Get cancellable flag
            expect(await bookingToken.isCancellable(0n)).to.equal(true);

            // Set cancellable flag
            await expect(supplierCMAccount.connect(signers.btAdmin).setCancellable(0n, false))
                .to.emit(bookingToken, "TokenCancellableUpdated")
                .withArgs(0n, false);

            // Check cancellable flag
            expect(await bookingToken.isCancellable(0n)).to.equal(false);

            // Try to set cancellable with unauthorized caller, should revert with
            // NotAuthorizedToSetCancellable
            await expect(distributorCMAccount.connect(signers.otherAccount3).setCancellable(0n, true))
                .to.revertedWithCustomError(distributorCMAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount3.address, BOOKING_OPERATOR_ROLE);

            // Try to set cancellable with unauthorized cm account, should revert with
            // NotAuthorizedToSetCancellable

            // Grant BOOKING_OPERATOR_ROLE for other account on distributor cm account
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.otherAccount3.address),
            ).to.not.reverted;

            await expect(distributorCMAccount.connect(signers.otherAccount3).setCancellable(0n, true))
                .to.revertedWithCustomError(bookingToken, "NotAuthorizedToSetCancellable")
                .withArgs(0n, await distributorCMAccount.getAddress());
        });
        it("should initiate cancellation of a booking token correctly", async function () {
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
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                    true,
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

            // Buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Try to cancel the token
            const token_id = 0n;
            const proposer = await supplierCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");

            await expect(supplierCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount))
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                1n, // Pending == 1
            ]);
        });
        it("should revert initiating a proposal if token state is reserved or expired", async function () {
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
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            // Try to cancel the token
            const token_id = 0n;
            const proposer = await supplierCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");

            await expect(supplierCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount))
                .to.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(0n, 1n); // tokenID == 0, Reserved == 1

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                0n,
                ethers.ZeroAddress,
                0n, // NoProposal == 0
            ]);

            // Expire the token

            // Advance time by 24 hours, token should can be expired after
            await network.provider.send("evm_increaseTime", [24 * 60 * 60]);
            await network.provider.send("evm_mine");

            // Expire the token
            await expect(supplierCMAccount.connect(signers.btAdmin).recordExpiration(0n))
                .to.emit(bookingToken, "TokenExpired")
                .withArgs(0n);

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(2); // Expired == 2

            // Try to cancel the token
            await expect(supplierCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount))
                .to.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(0n, 2n); // tokenID == 0, Expired == 2
        });
        it("Native: should accept a cancellation correctly and pay the refund", async function () {
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
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                    true,
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

            // Buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            await expect(buyTx)
                .to.be.emit(bookingToken, "TokenBought")
                .withArgs(0n, await distributorCMAccount.getAddress());

            // Try to cancel the token
            const token_id = 0n;
            const proposer = await distributorCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");

            await expect(distributorCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount))
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                1n, // Pending == 1
            ]);

            // Accept the cancellation, this should send the refund from supplier to distributor
            const acceptTx = supplierCMAccount.connect(signers.btAdmin).acceptCancellationProposal(token_id);

            // Sanity check
            await expect(acceptTx)
                .to.emit(bookingToken, "CancellationAccepted")
                .withArgs(token_id, await supplierCMAccount.getAddress(), refundAmount);

            // Check balances
            await expect(acceptTx).to.changeEtherBalances(
                [distributorCMAccount, supplierCMAccount, bookingToken],
                [refundAmount, -refundAmount, 0n],
            );
        });
        it("Native: should revert accepting if the caller is not authorized", async function () {
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
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                    true,
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

            // Buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            await expect(buyTx).to.not.be.reverted;

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(3); // Bought == 1

            // Try to cancel the token
            const token_id = 0n;
            const proposer = await supplierCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");

            // Initiate the cancellation by the supplier
            await expect(supplierCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount))
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer, // supplier
                1n, // Pending == 1
            ]);

            // Try to accept the cancellation proposed by the supplier with the
            // supplier cm account, this should revert due to being not authorized
            await expect(supplierCMAccount.connect(signers.btAdmin).acceptCancellationProposal(token_id))
                .to.be.revertedWithCustomError(bookingToken, "NotAuthorizedToAcceptCancellation")
                .withArgs(0n, await supplierCMAccount.getAddress());

            // Try to accept the cancellation proposed by the supplier with another account
            await expect(bookingToken.connect(signers.otherAccount3).acceptCancellationProposal(token_id))
                .to.be.revertedWithCustomError(bookingToken, "NotAuthorizedToAcceptCancellation")
                .withArgs(0n, signers.otherAccount3.address);

            // Accept the supplier's proposal with the distributor cm account
            await expect(distributorCMAccount.connect(signers.btAdmin).acceptCancellationProposal(token_id))
                .to.emit(bookingToken, "CancellationProposalAcceptedByTheOwner")
                .withArgs(token_id, await distributorCMAccount.getAddress(), refundAmount);

            // Now the supplier's proposal is accepted by the owner/distributor. Try
            // to accept it again with the distributor. This should revert with the
            // same error
            await expect(distributorCMAccount.connect(signers.btAdmin).acceptCancellationProposal(token_id))
                .to.be.revertedWithCustomError(bookingToken, "NotAuthorizedToAcceptCancellation")
                .withArgs(0n, await distributorCMAccount.getAddress());

            // Sanity check: Try to accept the cancellation proposed by the owner with another account
            await expect(bookingToken.connect(signers.otherAccount2).acceptCancellationProposal(token_id))
                .to.be.revertedWithCustomError(bookingToken, "NotAuthorizedToAcceptCancellation")
                .withArgs(0n, signers.otherAccount2.address);

            // Finally, accept the "owner accepted proposal" with the supplier cm account
            const acceptTx = await supplierCMAccount.connect(signers.btAdmin).acceptCancellationProposal(token_id);

            // Check balances
            await expect(acceptTx).to.changeEtherBalances(
                [distributorCMAccount, supplierCMAccount, bookingToken],
                [refundAmount, -refundAmount, 0n],
            );

            // Check events
            await expect(acceptTx)
                .to.emit(bookingToken, "CancellationAccepted")
                .withArgs(token_id, await supplierCMAccount.getAddress(), refundAmount);
        });
        it("ERC20: should accept a cancellation correctly and pay the refund", async function () {
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
                    nullUSD.getAddress(), // nullUSD address
                    true,
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
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());

            // Try to cancel the token
            const token_id = 0n;
            const proposer = await distributorCMAccount.getAddress();
            const refundAmount = ethers.parseEther("450");

            await expect(distributorCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount))
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                1n, // Pending == 1
            ]);

            // Accept the cancellation, this should send the refund from supplier to distributor
            const acceptTx = supplierCMAccount.connect(signers.btAdmin).acceptCancellationProposal(token_id);

            // Sanity check
            await expect(acceptTx)
                .to.emit(bookingToken, "CancellationAccepted")
                .withArgs(token_id, await supplierCMAccount.getAddress(), refundAmount);

            // Check balances
            await expect(acceptTx).to.changeEtherBalances(
                [distributorCMAccount, supplierCMAccount, bookingToken],
                [0n, 0n, 0n],
            );

            await expect(acceptTx).to.changeTokenBalances(
                nullUSD,
                [distributorCMAccount, supplierCMAccount],
                [refundAmount, -refundAmount],
            );
        });
        it("should revert transferring if there is an cancellation is active or cancelled", async function () {
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
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                    true,
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

            // Buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            await expect(buyTx)
                .to.be.emit(bookingToken, "TokenBought")
                .withArgs(0n, await distributorCMAccount.getAddress());

            // Initiate the proposal
            const token_id = 0n;
            const proposer = await distributorCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");

            await expect(distributorCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount))
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                1n, // Pending == 1
            ]);

            // THERE IS AN ACTIVE CANCELLATION PROPOSAL

            // Set WITHDRAWER_ROLE
            const WITHDRAWER_ROLE = await distributorCMAccount.WITHDRAWER_ROLE();
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(WITHDRAWER_ROLE, signers.withdrawer.address),
            ).to.not.reverted;

            // Try to transfer the token, should revert.
            await expect(
                distributorCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(await bookingToken.getAddress(), signers.otherAccount2.address, 0n),
            )
                .to.be.revertedWithCustomError(bookingToken, "TokenHasActiveCancellationProposalOrCancelled")
                .withArgs(0n);

            // Sanity check
            await expect(
                bookingToken
                    .connect(signers.withdrawer)
                    .safeTransferFrom(await distributorCMAccount.getAddress(), signers.otherAccount2.address, 0n),
            ).to.be.revertedWithCustomError(bookingToken, "TokenHasActiveCancellationProposalOrCancelled");

            await expect(
                bookingToken
                    .connect(signers.withdrawer)
                    .transferFrom(await distributorCMAccount.getAddress(), signers.otherAccount2.address, 0n),
            ).to.be.revertedWithCustomError(bookingToken, "TokenHasActiveCancellationProposalOrCancelled");

            // Accept the cancellation
            await expect(supplierCMAccount.connect(signers.btAdmin).acceptCancellationProposal(0n)).to.be.not.reverted;

            // Sanity check
            // Check token ownership, should revert
            await expect(bookingToken.ownerOf(0n))
                .to.be.revertedWithCustomError(bookingToken, "ERC721NonexistentToken")
                .withArgs(0n);

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(4); // Cancelled == 4

            // Check the proposal
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                4n, // Accepted == 1
            ]);

            // Try to transfer the token, should revert with InvalidTokenStatus
            // because token is burned when cancellation is accepted
            await expect(
                distributorCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(await bookingToken.getAddress(), signers.otherAccount2.address, 0n),
            ).to.be.revertedWithCustomError(bookingToken, "InvalidTokenStatus");
        });
        // it("should allow transfer if cancellation is rejected", async function () {
        //     const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
        //         await loadFixture(deployBookingTokenFixture);

        //     const tokenURI =
        //         "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

        //     const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

        //     const price = ethers.parseEther("0.05");

        //     /***************************************************
        //      *                   SUPPLIER                      *
        //      ***************************************************/

        //     // Grant BOOKING_OPERATOR_ROLE
        //     const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
        //     await expect(
        //         supplierCMAccount
        //             .connect(signers.cmAccountAdmin)
        //             .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
        //     ).to.not.reverted;

        //     await expect(
        //         await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
        //             distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
        //             tokenURI, // tokenURI
        //             expirationTimestamp, // expiration
        //             price, // price
        //             ethers.ZeroAddress, // zero address
        //             true,
        //         ),
        //     )
        //         .to.be.emit(bookingToken, "TokenReserved")
        //         .withArgs(
        //             0n,
        //             distributorCMAccount.getAddress(),
        //             supplierCMAccount.getAddress(),
        //             expirationTimestamp,
        //             price,
        //             ethers.ZeroAddress, // zero address
        //             true,
        //         );

        //     // Check token ownership
        //     expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

        //     // Check token booking status
        //     expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

        //     /***************************************************
        //      *                  DISTRIBUTOR                    *
        //      ***************************************************/

        //     // Grant BOOKING_OPERATOR_ROLE
        //     await expect(
        //         distributorCMAccount
        //             .connect(signers.cmAccountAdmin)
        //             .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
        //     ).to.not.reverted;

        //     // Buy the token
        //     const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

        //     await expect(buyTx)
        //         .to.be.emit(bookingToken, "TokenBought")
        //         .withArgs(0n, await distributorCMAccount.getAddress());

        //     // Try to cancel the token
        //     const token_id = 0n;
        //     const proposer = await distributorCMAccount.getAddress();
        //     const refundAmount = ethers.parseEther("0.045");

        //     await expect(distributorCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount))
        //         .to.emit(bookingToken, "CancellationPending")
        //         .withArgs(token_id, proposer, refundAmount);

        //     // Sanity check
        //     expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
        //         refundAmount,
        //         proposer,
        //         1n, // Pending == 1
        //     ]);

        //     // THERE IS AN ACTIVE CANCELLATION PROPOSAL

        //     // Set WITHDRAWER_ROLE
        //     const WITHDRAWER_ROLE = await distributorCMAccount.WITHDRAWER_ROLE();
        //     await expect(
        //         distributorCMAccount
        //             .connect(signers.cmAccountAdmin)
        //             .grantRole(WITHDRAWER_ROLE, signers.withdrawer.address),
        //     ).to.not.reverted;

        //     // Try to transfer the token, should revert.
        //     await expect(
        //         distributorCMAccount
        //             .connect(signers.withdrawer)
        //             .transferERC721(await bookingToken.getAddress(), signers.otherAccount2.address, 0n),
        //     )
        //         .to.be.revertedWithCustomError(bookingToken, "TokenHasActiveCancellationProposalOrCancelled")
        //         .withArgs(0n);

        //     // Sanity check
        //     await expect(
        //         bookingToken
        //             .connect(signers.withdrawer)
        //             .safeTransferFrom(await distributorCMAccount.getAddress(), signers.otherAccount2.address, 0n),
        //     ).to.be.revertedWithCustomError(bookingToken, "TokenHasActiveCancellationProposalOrCancelled");

        //     await expect(
        //         bookingToken
        //             .connect(signers.withdrawer)
        //             .transferFrom(await distributorCMAccount.getAddress(), signers.otherAccount2.address, 0n),
        //     ).to.be.revertedWithCustomError(bookingToken, "TokenHasActiveCancellationProposalOrCancelled");
        // });
        // FIXME: add tests for other cases
    });
});
