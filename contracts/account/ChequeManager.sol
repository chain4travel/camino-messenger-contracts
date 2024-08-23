// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Cheque Manager

pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";

// Manager Interface
import { ICMAccountManager } from "../manager/ICMAccountManager.sol";

/**
 * @notice ChequeManager manages, verifies, and cashes in the messenger cheques.
 *
 * EIP712 Domain name & version:
 *   DOMAIN_NAME = "CaminoMessenger"
 *   DOMAIN_VERSION= "1"
 */
abstract contract ChequeManager is Initializable, ReentrancyGuardUpgradeable {
    using ECDSA for bytes32;
    using Address for address payable;

    /***************************************************
     *                   CONSTANTS                     *
     ***************************************************/

    /**
     * @notice Pre-computed hash of the MessengerCheque type
     *
     * ```
     * keccak256(
     *     "MessengerCheque(address fromCMAccount,address toCMAccount,address toBot,uint256 counter,uint256 amount,uint256 createdAt,uint256 expiresAt)"
     * );
     * ```
     */
    bytes32 public constant MESSENGER_CHEQUE_TYPEHASH =
        0x87b38f131334165ac2b361f08966c9fcff3a953fa7d9d9c2861b7f0b50445bcb;

    /**
     * @notice Pre-computed hash of the EIP712Domain type
     *
     * ```
     * keccak256("EIP712Domain(string name,string version,uint256 chainId)");
     * ```
     */
    bytes32 public constant DOMAIN_TYPEHASH = 0xc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e;

    /***************************************************
     *                   STRUCTS                       *
     ***************************************************/

    /**
     * @notice Struct representing a Messenger Cheque.
     */
    struct MessengerCheque {
        address fromCMAccount; // CM Account that will pay the amount
        address toCMAccount; // CM Account that will receive the amount
        address toBot; // The address of the bot that receives the cheque
        uint256 counter; // This should be increased with every cheque
        uint256 amount; // The amount to be transferred
        uint256 createdAt; // Creation timestamp of the cheque
        uint256 expiresAt; // Expiration timestamp of the cheque
    }

    /**
     * @notice Struct for tracking the counter, amount and timestamps used for the last
     * cash-in operation.
     */
    struct LastCashIn {
        uint256 counter;
        uint256 amount;
        uint256 createdAt;
        uint256 expiresAt;
    }

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    /// @custom:storage-location erc7201:camino.messenger.storage.ChequeManager
    struct ChequeManagerStorage {
        /**
         * @notice Mapping to track the cash-in details for each pair of fromBot and toBot addresses.
         */
        mapping(address fromBot => mapping(address toBot => LastCashIn)) _lastCashIns;
        /**
         * @notice Total amount of cheques that have been cashed in.
         */
        uint256 _totalChequePayments;
        /**
         * @notice EIP712 Domain Separator used for signature verification. This variable includes
         * dynamic chain ID, hence it is not a constant.
         */
        bytes32 _domainSeparator;
    }

    // keccak256(abi.encode(uint256(keccak256("camino.messenger.storage.ChequeManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant ChequeManagerStorageLocation =
        0x175f7e400d42af44d9ebd24e9efee8a2c4ed78ddf46a83e51a493ae382c87600;

    function _getChequeManagerStorage() private pure returns (ChequeManagerStorage storage $) {
        assembly {
            $.slot := ChequeManagerStorageLocation
        }
    }

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    /**
     * @notice Cheque verified event. Emitted when a cheque is verified.
     */
    event ChequeVerified(
        address indexed fromCMAccount,
        address indexed toCMAccount,
        address fromBot,
        address toBot,
        uint256 counter,
        uint256 amount,
        uint256 payment
    );

    /**
     * @notice Cash-in event. Emitted when a cheque is cashed in.
     */
    event ChequeCashedIn(
        address indexed fromBot,
        address indexed toBot,
        uint256 counter,
        uint256 paid,
        uint256 developerFee
    );

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    /**
     * @notice Invalid CM Account. Cheque's `fromCMAccount` has to be for `address(this)`.
     */
    error InvalidFromCMAccount(address fromCMAccount);

    /**
     * @notice `toCMAccount` address is not a registered CMAccount on the manager.
     */
    error InvalidToCMAccount(address toCMAccount);

    /**
     * @notice The signer is not allowed to sign cheques
     */
    error NotAllowedToSignCheques(address signer);

    /**
     * @notice Invalid counter for the cheque. The counter on the cheque is not greater then the last
     * recorded counter.
     */
    error InvalidCounter(uint256 chequeCounter, uint256 lastCounter);

    /**
     * @notice Last recorded amount is lower than the cheque's amount.
     */
    error InvalidAmount(uint256 chequeAmount, uint256 lastAmount);

    /**
     * @notice The cheque is expired at the given timestamp.
     */
    error ChequeExpired(uint256 expiresAt);

    /***************************************************
     *                    FUNCS                        *
     ***************************************************/

    /**
     * @notice Initializes the contract, setting the domain separator with EIP712 domain type hash and
     * the domain.
     *
     * EIP712Domain {
     *   string name;
     *   string version;
     *   uint256 chainid;
     * }
     */
    function __ChequeManager_init() internal onlyInitializing {
        ChequeManagerStorage storage $ = _getChequeManagerStorage();

        $._domainSeparator = keccak256(
            abi.encode(DOMAIN_TYPEHASH, keccak256("CaminoMessenger"), keccak256("1"), block.chainid)
        );
    }

    /**
     * @notice Returns the domain separator.
     */
    function getDomainSeparator() public view returns (bytes32) {
        ChequeManagerStorage storage $ = _getChequeManagerStorage();
        return $._domainSeparator;
    }

    /**
     * @notice Returns the hash of the `MessengerCheque` encoded with
     * `MESSENGER_CHEQUE_TYPEHASH`.
     */
    function hashMessengerCheque(
        address fromCMAccount,
        address toCMAccount,
        address toBot,
        uint256 counter,
        uint256 amount,
        uint256 createdAt,
        uint256 expiresAt
    ) public pure returns (bytes32) {
        return
            keccak256(
                abi.encode(
                    MESSENGER_CHEQUE_TYPEHASH,
                    fromCMAccount,
                    toCMAccount,
                    toBot,
                    counter,
                    amount,
                    createdAt,
                    expiresAt
                )
            );
    }

    /**
     * @notice Returns the hash of the typed data (cheque) with prefix and domain
     * separator.
     */
    function hashTypedDataV4(
        address fromCMAccount,
        address toCMAccount,
        address toBot,
        uint256 counter,
        uint256 amount,
        uint256 createdAt,
        uint256 expiresAt
    ) public view returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    "\x19\x01",
                    getDomainSeparator(),
                    hashMessengerCheque(fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt)
                )
            );
    }

    /**
     * @notice Returns the signer for the given cheque and signature. Uses {ECDSA} library to
     * recover the signer.
     */
    function recoverSigner(
        address fromCMAccount,
        address toCMAccount,
        address toBot,
        uint256 counter,
        uint256 amount,
        uint256 createdAt,
        uint256 expiresAt,
        bytes memory signature
    ) internal view returns (address signer) {
        bytes32 digest = hashTypedDataV4(fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt);
        signer = digest.recover(signature);
        return signer;
    }

    /**
     * @notice Returns signer and payment amount if the signature is valid for the
     * given cheque, the signer is an allowed bot, cheque counter and amounts are
     * valid according to last cash ins.
     *
     * Please be aware that `cheque.amount < paymentAmount` for a valid cheque as
     * long as the last amount is lower than the cheque amount. Only the difference
     * between the cheque amount and the last recorded amount is paid.
     */
    function verifyCheque(
        address fromCMAccount,
        address toCMAccount,
        address toBot,
        uint256 counter,
        uint256 amount,
        uint256 createdAt,
        uint256 expiresAt,
        bytes memory signature
    ) public returns (address signer, uint256 paymentAmount) {
        // Revert if cheque is not for this contract
        if (fromCMAccount != address(this)) {
            revert InvalidFromCMAccount(fromCMAccount);
        }

        // Revert if cheque payee is not a CM account
        if (!ICMAccountManager(getManagerAddress()).isCMAccount(toCMAccount)) {
            revert InvalidToCMAccount(toCMAccount);
        }

        // Revert if the cheque is expired
        if (block.timestamp >= expiresAt) {
            revert ChequeExpired(expiresAt);
        }

        // Recover signer
        signer = recoverSigner(fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt, signature);

        // Check if the signer is an allowed bot.
        if (!isBotAllowed(signer)) {
            revert NotAllowedToSignCheques(signer);
        }

        ChequeManagerStorage storage $ = _getChequeManagerStorage();
        LastCashIn storage lastCashIn = $._lastCashIns[signer][toBot];

        // Revert if the cheque amount is lower then the last recorded amount
        if (amount < lastCashIn.amount) {
            revert InvalidAmount(amount, lastCashIn.amount);
        }

        // Ensure the current cheque's counter is greater than the last recorded one
        if (counter <= lastCashIn.counter) {
            revert InvalidCounter(counter, lastCashIn.counter);
        }

        // Everything is valid. Calculate payment amount.
        paymentAmount = amount - lastCashIn.amount;

        // Emit event
        emit ChequeVerified(
            fromCMAccount,
            toCMAccount,
            signer, // fromBot
            toBot,
            counter,
            amount,
            paymentAmount
        );

        return (signer, paymentAmount);
    }

    /**
     * @notice Cash in a cheque by verifying it and paying the difference between the
     * cheque amount and the last recorded amount for the signer and `toBot` pair.
     *
     * A percentage of the amount is also paid to the developer wallet.
     *
     * @param fromCMAccount The CM Account that will pay the amount. This contract.
     * @param toCMAccount The CM Account that will receive the amount.
     * @param toBot The address of the bot that received the cheque.
     * @param counter The counter of the cheque. Should be increased with every
     * cheque.
     * @param amount The amount on the cheque. Should be greater then or equal the
     * last recorded amount.
     * @param createdAt The creation timestamp of the cheque.
     * @param expiresAt The expiration timestamp of the cheque.
     * @param signature The signature of the cheque.
     */
    function cashInCheque(
        address fromCMAccount,
        address toCMAccount,
        address toBot,
        uint256 counter,
        uint256 amount,
        uint256 createdAt,
        uint256 expiresAt,
        bytes memory signature
    ) public nonReentrant {
        // Verify the cheque and get the signer and payment amount
        (address signer, uint256 paymentAmount) = verifyCheque(
            fromCMAccount,
            toCMAccount,
            toBot,
            counter,
            amount,
            createdAt,
            expiresAt,
            signature
        );

        // If we didn't revert in the verifyCheque above, the cheque is valid.
        // Update the last cash ins.
        setLastCashIn(signer, toBot, counter, amount, createdAt, expiresAt);

        // Calculate developer fee
        // For amounts lower then fee basis point, the developer fee is 0.
        uint256 developerFee = (paymentAmount * ICMAccountManager(getManagerAddress()).getDeveloperFeeBp()) / 10000;

        // Subtract developer fee from payment amount
        uint256 chequePaymentAmount = paymentAmount - developerFee;

        // Update total cheque payments excluding cheques to the same account
        if (fromCMAccount != toCMAccount) {
            ChequeManagerStorage storage $ = _getChequeManagerStorage();
            $._totalChequePayments += paymentAmount;
        }

        // Transfer developer fee to the developer wallet
        payable(ICMAccountManager(getManagerAddress()).getDeveloperWallet()).sendValue(developerFee);

        // Transfer the cheque payment amount to the `toCMAccount`
        payable(toCMAccount).sendValue(chequePaymentAmount);

        // Emit cash-in event
        emit ChequeCashedIn(signer, toBot, counter, chequePaymentAmount, developerFee);
    }

    /**
     * @notice Returns last cash-in details for given `fromBot` & `toBot` pair.
     *
     * @param fromBot The address of the bot that sent the cheque.
     * @param toBot The address of the bot that received the cheque.
     *
     * Returns (lastCounter, lastAmount, lastCreatedAt, lastExpiresAt)
     *
     * @return lastCounter The last counter of the cheque.
     * @return lastAmount The last amount of the cheque.
     * @return lastCreatedAt The last creation timestamp of the cheque.
     * @return lastExpiresAt The last expiration timestamp of the cheque.
     */
    function getLastCashIn(
        address fromBot,
        address toBot
    ) public view returns (uint256 lastCounter, uint256 lastAmount, uint256 lastCreatedAt, uint256 lastExpiresAt) {
        ChequeManagerStorage storage $ = _getChequeManagerStorage();
        LastCashIn storage lastCashIn = $._lastCashIns[fromBot][toBot];
        return (lastCashIn.counter, lastCashIn.amount, lastCashIn.createdAt, lastCashIn.expiresAt);
    }

    /**
     * @notice Sets last cash-in for given `fromBot`, `toBot` pair.
     *
     * @param fromBot The address of the bot that sent the cheque.
     * @param toBot The address of the bot that received the cheque.
     * @param counter The counter of the cheque.
     * @param amount The amount of the cheque.
     * @param createdAt The creation timestamp of the cheque.
     * @param expiresAt The expiration timestamp of the cheque.
     */
    function setLastCashIn(
        address fromBot,
        address toBot,
        uint256 counter,
        uint256 amount,
        uint256 createdAt,
        uint256 expiresAt
    ) internal {
        ChequeManagerStorage storage $ = _getChequeManagerStorage();
        $._lastCashIns[fromBot][toBot] = LastCashIn(counter, amount, createdAt, expiresAt);
    }

    /**
     * @notice Returns total cheque payments. This is the sum of all cashed in cheques.
     *
     * @return totalChequePayments The total cheque payments made.
     */
    function getTotalChequePayments() public view returns (uint256) {
        ChequeManagerStorage storage $ = _getChequeManagerStorage();
        return $._totalChequePayments;
    }

    /***************************************************
     *                   ABSTRACT                      *
     ***************************************************/

    /**
     * @notice Abstract function to check if a bot is allowed to sign cheques. This
     * must be implemented by the inheriting contract.
     */
    function isBotAllowed(address bot) public view virtual returns (bool);

    /**
     * @notice Abstract function to get the manager address. This must be implemented
     * by the inheriting contract.
     */
    function getManagerAddress() public view virtual returns (address);
}
