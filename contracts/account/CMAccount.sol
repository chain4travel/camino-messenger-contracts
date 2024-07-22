// SPDX-License-Identifier: UNLICENSED
//
// Camino Messenger Account Implementation

pragma solidity ^0.8.24;

// UUPS Proxy
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

// Access
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

// Cashier
import "./MessengerCashier.sol";

interface ICMAccountManager {
    function getAccountImplementation() external view returns (address);

    function getDeveloperFeeBp() external view returns (uint256);

    function getDeveloperWallet() external view returns (address);
}

/**
 * @dev CM Account manages multiple bots for distributors and suppliers on Camino Messenger.
 *
 * This account holds funds that will be paid to the cheque beneficiaries.
 */
contract CMAccount is Initializable, PausableUpgradeable, AccessControlUpgradeable, UUPSUpgradeable, MessengerCashier {
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    bytes32 public constant CHEQUE_OPERATOR_ROLE = keccak256("CHEQUE_OPERATOR_ROLE");

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    address private _manager;

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    /**
     * CMAccount upgrade event. Emitted when the CMAccount implementation is upgraded.
     */
    event CMAccountUpgraded(address indexed oldImplementation, address indexed newImplementation);

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    /**
     * @dev CMAccount implementation address does not match the one in the manager
     */
    error CMAccountImplementationMismatch(address latestImplementation, address newImplementation);

    /**
     * @dev New implementation is the same as the current implementation, no update needed
     */
    error CMAccountNoUpdateNeeded(address oldImplementation, address newImplementation);

    /***************************************************
     *         CONSTRUCTOR & INITIALIZATION            *
     ***************************************************/

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @dev Account should be able to receive CAM
     *
     * FIXME: Check if this is ok legally or we need to introduce checks
     */
    receive() external payable virtual {}

    function initialize(address manager, address defaultAdmin, address pauser, address upgrader) public initializer {
        __Pausable_init();
        __AccessControl_init();
        __UUPSUpgradeable_init();
        __MessengerCashier_init();

        _grantRole(DEFAULT_ADMIN_ROLE, defaultAdmin);
        _grantRole(PAUSER_ROLE, pauser);
        _grantRole(UPGRADER_ROLE, upgrader);

        _manager = manager;
    }

    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /**
     * @dev Upgrades the CMAccount implementation.
     *
     * Reverts if the new implementation is the same as the old one.
     *
     * Reverts if the new implementation does not match the implementation address in the manager.
     * Only implementations registered at the manager are allowed.
     *
     * Emits a {CMAccountUpgraded} event.
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(UPGRADER_ROLE) {
        // Get the implementation address from the manager
        address managerImplementation = ICMAccountManager(_manager).getAccountImplementation();
        address oldImplementation = ERC1967Utils.getImplementation();

        // Revert if the new implementation is the same as the old one
        if (oldImplementation == newImplementation) {
            revert CMAccountNoUpdateNeeded(oldImplementation, newImplementation);
        }

        // Check if new implementation matches the implementation address in the manager
        if (newImplementation != managerImplementation) {
            revert CMAccountImplementationMismatch(managerImplementation, newImplementation);
        }

        emit CMAccountUpgraded(oldImplementation, newImplementation);
    }

    /**
     * @dev Returns true if an address is authorized bot
     */
    function isBotAllowed(address bot) public view override returns (bool) {
        return hasRole(CHEQUE_OPERATOR_ROLE, bot);
    }

    /**
     * @dev Return developer wallet address
     */
    function getDeveloperWallet() public view override returns (address) {
        address developerWallet = ICMAccountManager(_manager).getDeveloperWallet();
        return developerWallet;
    }

    /**
     * @dev Return developer fee in basis points
     */
    function getDeveloperFeeBp() public view override returns (uint256) {
        uint256 developerFeeBp = ICMAccountManager(_manager).getDeveloperFeeBp();
        return developerFeeBp;
    }
}
