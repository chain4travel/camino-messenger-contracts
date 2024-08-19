// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Cheque Manager

pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/Address.sol";

/**
 * @dev ChequeManager manages, verifies, and cashes in messenger cheques.
 *
 * EIP712 Domain name & version:
 *   DOMAIN_NAME = "CaminoMessenger"
 *   DOMAIN_VERSION= "1"
 */
abstract contract ChequeManager is Initializable {
    using ECDSA for bytes32;
    using Address for address payable;

    /***************************************************
     *                   CONSTANTS                     *
     ***************************************************/

    // FIXME: Use pre-computed hash
    // Pre-computed hash of the MessengerCheque struct type
    // > ethers.keccak256(ethers.toUtf8Bytes("MessengerCheque(address fromCMAccount,address toCMAccount,address toBot,uint256 counter,uint256 amount,uint256 createdAt,uint256 expiresAt)"))
    // '0x87b38f131334165ac2b361f08966c9fcff3a953fa7d9d9c2861b7f0b50445bcb'
    bytes32 public constant MESSENGER_CHEQUE_TYPEHASH =
        keccak256(
            "MessengerCheque(address fromCMAccount,address toCMAccount,address toBot,uint256 counter,uint256 amount,uint256 createdAt,uint256 expiresAt)"
        );

    // FIXME: Use pre-computed hash
    // Pre-computed hash of the EIP712Domain type
    // > ethers.keccak256(ethers.toUtf8Bytes("EIP712Domain(string name,string version,uint256 chainId)"))
    // '0xc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e'
    bytes32 public constant DOMAIN_TYPEHASH = keccak256("EIP712Domain(string name,string version,uint256 chainId)");

    /***************************************************
     *                   STRUCTS                       *
     ***************************************************/

    /**
     * @dev Struct representing a Messenger Cheque.
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
     * @dev Struct for tracking the last counter and amount used for the last cash-in operation.
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
         * @dev Mapping to track the cash-in details for each pair of fromBot and toBot addresses.
         */
        mapping(address fromBot => mapping(address toBot => LastCashIn)) _lastCashIns;
        /**
         * @dev Total amount of cheques that have been cashed in.
         */
        uint256 _totalChequePayments;
        /**
         * @dev EIP712 Domain Separator used for signature verification. This variable includes
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
     * @dev Cheque verified event. Emitted when a cheque is verified.
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
     * @dev Cash-in event. Emitted when a cheque is cashed in.
     */
    event ChequeCashedIn(
        address indexed fromBot,
        address indexed toBot,
        uint256 counter,
        uint256 paid,
        uint256 developerFee
    );

    /**
     * @dev Last recorded amount and cheque's amount is the same. There is nothing to pay.
     */
    event NothingToPay();

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    /**
     * @dev Invalid CM Account. Cheque's `fromCMAccount` has to be for `address(this)`.
     */
    error InvalidFromCMAccount(address fromCMAccount);

    /**
     * @dev `toCMAccoun` address is not a registered CMAccount on the manager
     */
    error InvalidToCMAccount(address toCMAccount);

    /**
     * @dev The signer is not allowed to sign cheques
     */
    error NotAllowedToSignCheques(address signer);

    /**
     * @dev Invalid counter for the cheque. The counter on the cheque is not greater then the last
     * recorded counter.
     */
    error InvalidCounter(uint256 chequeCounter, uint256 lastCounter);

    /**
     * @dev Last recorded amount is lower than the cheque's amount.
     */
    error InvalidAmount(uint256 chequeAmount, uint256 lastAmount);

    /**
     * @dev The cheque is expired at the given timestamp
     */
    error ChequeExpired(uint256 expiresAt);

    /***************************************************
     *                    FUNCS                        *
     ***************************************************/

    /**
     * @dev Initializes the contract, setting the domain separator with EIP712 domain type hash and
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

    function getDomainSeparator() public view returns (bytes32) {
        ChequeManagerStorage storage $ = _getChequeManagerStorage();
        return $._domainSeparator;
    }

    /**
     * @dev Returns the hash of the `MessengerCheque` encoded with `MESSENGER_CHEQUE_TYPEHASH`.
     */
    function hashMessengerCheque(MessengerCheque memory cheque) public pure returns (bytes32) {
        return
            keccak256(
                abi.encode(
                    MESSENGER_CHEQUE_TYPEHASH,
                    cheque.fromCMAccount,
                    cheque.toCMAccount,
                    cheque.toBot,
                    cheque.counter,
                    cheque.amount,
                    cheque.createdAt,
                    cheque.expiresAt
                )
            );
    }

    /**
     * @dev Return hash of the typed data (cheque) with prefix and domain separator.
     */
    function hashTypedDataV4(MessengerCheque memory cheque) public view returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", getDomainSeparator(), hashMessengerCheque(cheque)));
    }

    /**
     * @dev Return signer for the given cheque and signature
     */
    function recoverSigner(MessengerCheque memory cheque, bytes memory signature) public view returns (address signer) {
        bytes32 digest = hashTypedDataV4(cheque);
        signer = digest.recover(signature);
        return signer;
    }

    /**
     * @dev Returns signer and payment amount if the signature is valid for the given cheque,
     * the signer is an allowed bot, cheque counter and amounts are valid accrording to last cash ins.
     *
     * Please be aware that `cheque.amount <
     */
    function verifyCheque(
        MessengerCheque memory cheque,
        bytes memory signature
    ) public returns (address signer, uint256 paymentAmount) {
        // Revert if cheque is not for this contract
        if (cheque.fromCMAccount != address(this)) {
            revert InvalidFromCMAccount(cheque.fromCMAccount);
        }

        // Revert if cheque payee is not a CM account
        if (!isCMAccount(cheque.toCMAccount)) {
            revert InvalidToCMAccount(cheque.toCMAccount);
        }

        // Revert if the cheque is expired
        if (block.timestamp >= cheque.expiresAt) {
            revert ChequeExpired(cheque.expiresAt);
        }

        // Recover the signer from the signature. If the signature is invalid, this
        // will recover different signer address.
        bytes32 digest = hashTypedDataV4(cheque);
        signer = digest.recover(signature);

        // Check if the signer is an allowed bot.
        if (!isBotAllowed(signer)) {
            revert NotAllowedToSignCheques(signer);
        }

        // Get the last cash-in details for the signer and `toBot`
        LastCashIn memory lastCashIn = getLastCashIn(signer, cheque.toBot);

        // Revert if the cheque amount is lower then the last recorded amount
        if (cheque.amount < lastCashIn.amount) {
            revert InvalidAmount(cheque.amount, lastCashIn.amount);
        }

        // If cheque amount is same as the last cashed in amount, there is nothing to pay.
        // This might happen if the service is free of charge on Camino Messenger and cheque
        // holder still wants to record the cheque on-chain.
        if (cheque.amount == lastCashIn.amount) {
            // There is nothing to pay but the cheque is still valid, so continue the process.
            emit NothingToPay();
        }

        // Ensure the current cheque's counter is greater than the last recorded one
        if (cheque.counter <= lastCashIn.counter) {
            revert InvalidCounter(cheque.counter, lastCashIn.counter);
        }

        // Everthing is valid. Calculate payment amount.
        paymentAmount = cheque.amount - lastCashIn.amount;

        // Emit event
        emit ChequeVerified(
            cheque.fromCMAccount,
            cheque.toCMAccount,
            signer, // fromBot
            cheque.toBot,
            cheque.counter,
            cheque.amount,
            paymentAmount
        );

        return (signer, paymentAmount);
    }

    /**
     * @dev Cash in a cheque by verifying it and paying the difference between the cheque amount
     * and the last recorded amount for the signer and `toBot` pair.
     *
     * A percentage of the amount is also paid to the developer wallet.
     */
    function cashInCheque(MessengerCheque memory cheque, bytes memory signature) public {
        // Authorize cheque cash in
        _authorizeChequeCashIn(cheque, signature);

        // Verify the cheque and get the signer and payment amount
        (address signer, uint256 paymentAmount) = verifyCheque(cheque, signature);

        // If we didn't revert in the verifyCheque above, the cheque is valid.
        // Update the last cash ins.
        setLastCashIn(signer, cheque.toBot, cheque.counter, cheque.amount, cheque.createdAt, cheque.expiresAt);

        // Calculate developer fee
        uint256 developerFee = calculateDeveloperFee(paymentAmount);

        // Subtract developer fee from payment amount
        uint256 chequePaymentAmount = paymentAmount - developerFee;

        // Transfer developer fee to the developer wallet
        payable(getDeveloperWallet()).sendValue(developerFee);

        // Transfer the cheque payment amount to the `toCMAccount`
        payable(cheque.toCMAccount).sendValue(chequePaymentAmount);

        // Update total cheque payments excluding cheques to the same account
        if (cheque.fromCMAccount != cheque.toCMAccount) {
            addToTotalChequePayments(paymentAmount);
        }

        // Emit cash-in event
        emit ChequeCashedIn(signer, cheque.toBot, cheque.counter, chequePaymentAmount, developerFee);
    }

    /**
     * @dev Returns last cash-in for given `fromBot`, `toBot` pair.
     */
    function getLastCashIn(address fromBot, address toBot) public view returns (LastCashIn memory cashIn) {
        ChequeManagerStorage storage $ = _getChequeManagerStorage();
        return $._lastCashIns[fromBot][toBot];
    }

    /**
     * @dev Sets last cash-in for given `fromBot`, `toBot` pair.
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
     * @dev Calculates the developer fee for a given amount.
     *
     * For amounts lower then fee basis point, the developer fee is 0.
     */
    function calculateDeveloperFee(uint256 amount) public view returns (uint256) {
        return (amount * getDeveloperFeeBp()) / 10000;
    }

    /**
     * @dev Returns total cheque payments
     */
    function getTotalChequePayments() public view returns (uint256) {
        ChequeManagerStorage storage $ = _getChequeManagerStorage();
        return $._totalChequePayments;
    }

    /**
     * @dev Sets total cheque payments
     */
    function setTotalChequePayments(uint256 totalChequePayments) internal {
        ChequeManagerStorage storage $ = _getChequeManagerStorage();
        $._totalChequePayments = totalChequePayments;
    }

    /**
     * @dev Adds to total cheque payments
     */
    function addToTotalChequePayments(uint256 amount) internal {
        ChequeManagerStorage storage $ = _getChequeManagerStorage();
        $._totalChequePayments += amount;
    }

    /***************************************************
     *                   ABSTRACT                      *
     ***************************************************/

    /**
     * @dev Function that should revert when `msg.sender` is not authorized to cash-in the cheque.
     * Called by {cashInCheque}.
     */
    function _authorizeChequeCashIn(MessengerCheque memory cheque, bytes memory signature) internal virtual {}

    /**
     * @dev Abstract function to check if a bot is allowed to sign cheques. This must be implemented
     * by the inheriting contract.
     */
    function isBotAllowed(address bot) public view virtual returns (bool);

    /**
     * @dev Abstract function to check if an address is a Camino Messenger account. This must be
     * implemented by the inheriting contract.
     */
    function isCMAccount(address account) internal view virtual returns (bool);

    /**
     * @dev Abstract function to get the developer wallet. This must be implemented by the inheriting
     * contract.
     */
    function getDeveloperWallet() public view virtual returns (address developerWallet);

    /**
     * @dev Abstract function to get the developer fee in basis points. This must be implemented by
     * the inheriting contract.
     *
     * A basis point (bp) is one hundredth of 1 percentage point.
     *
     * 1 bp = 0.01%, 1/10,000⁠, or 0.0001.
     * 10 bp = 0.1%, 1/1,000⁠, or 0.001.
     * 100 bp = 1%, ⁠1/100⁠, or 0.01.
     */
    function getDeveloperFeeBp() public view virtual returns (uint256 developerFee);
}
