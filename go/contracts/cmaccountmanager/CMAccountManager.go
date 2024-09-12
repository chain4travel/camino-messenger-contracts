// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cmaccountmanager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CmaccountmanagerMetaData contains all meta data concerning the Cmaccountmanager contract.
var CmaccountmanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"CMAccountInvalidAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"CMAccountInvalidImplementation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sended\",\"type\":\"uint256\"}],\"name\":\"IncorrectPrefundAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bookingToken\",\"type\":\"address\"}],\"name\":\"InvalidBookingTokenAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"developerWallet\",\"type\":\"address\"}],\"name\":\"InvalidDeveloperWallet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"ServiceAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ServiceNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldBookingToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newBookingToken\",\"type\":\"address\"}],\"name\":\"BookingTokenAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"CMAccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"CMAccountImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldDeveloperFeeBp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newDeveloperFeeBp\",\"type\":\"uint256\"}],\"name\":\"DeveloperFeeBpUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldDeveloperWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDeveloperWallet\",\"type\":\"address\"}],\"name\":\"DeveloperWalletUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CMACCOUNT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEVELOPER_WALLET_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PREFUND_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SERVICE_REGISTRY_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSIONER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgrader\",\"type\":\"address\"}],\"name\":\"createCMAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllRegisteredServiceHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"services\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllRegisteredServiceNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"services\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBookingTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCMAccountCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeveloperFeeBp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"developerFeeBp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeveloperWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"developerWallet\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrefundAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"getRegisteredServiceHashByName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getRegisteredServiceNameByHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgrader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"versioner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"developerWallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"developerFeeBp\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isCMAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"registerService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"setAccountImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"setBookingTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bp\",\"type\":\"uint256\"}],\"name\":\"setDeveloperFeeBp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"developerWallet\",\"type\":\"address\"}],\"name\":\"setDeveloperWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrefundAmount\",\"type\":\"uint256\"}],\"name\":\"setPrefundAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"unregisterService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// CmaccountmanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use CmaccountmanagerMetaData.ABI instead.
var CmaccountmanagerABI = CmaccountmanagerMetaData.ABI

// Cmaccountmanager is an auto generated Go binding around an Ethereum contract.
type Cmaccountmanager struct {
	CmaccountmanagerCaller     // Read-only binding to the contract
	CmaccountmanagerTransactor // Write-only binding to the contract
	CmaccountmanagerFilterer   // Log filterer for contract events
}

// CmaccountmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CmaccountmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmaccountmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CmaccountmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmaccountmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CmaccountmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmaccountmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CmaccountmanagerSession struct {
	Contract     *Cmaccountmanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CmaccountmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CmaccountmanagerCallerSession struct {
	Contract *CmaccountmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CmaccountmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CmaccountmanagerTransactorSession struct {
	Contract     *CmaccountmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CmaccountmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CmaccountmanagerRaw struct {
	Contract *Cmaccountmanager // Generic contract binding to access the raw methods on
}

// CmaccountmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CmaccountmanagerCallerRaw struct {
	Contract *CmaccountmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// CmaccountmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CmaccountmanagerTransactorRaw struct {
	Contract *CmaccountmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCmaccountmanager creates a new instance of Cmaccountmanager, bound to a specific deployed contract.
func NewCmaccountmanager(address common.Address, backend bind.ContractBackend) (*Cmaccountmanager, error) {
	contract, err := bindCmaccountmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cmaccountmanager{CmaccountmanagerCaller: CmaccountmanagerCaller{contract: contract}, CmaccountmanagerTransactor: CmaccountmanagerTransactor{contract: contract}, CmaccountmanagerFilterer: CmaccountmanagerFilterer{contract: contract}}, nil
}

// NewCmaccountmanagerCaller creates a new read-only instance of Cmaccountmanager, bound to a specific deployed contract.
func NewCmaccountmanagerCaller(address common.Address, caller bind.ContractCaller) (*CmaccountmanagerCaller, error) {
	contract, err := bindCmaccountmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerCaller{contract: contract}, nil
}

// NewCmaccountmanagerTransactor creates a new write-only instance of Cmaccountmanager, bound to a specific deployed contract.
func NewCmaccountmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*CmaccountmanagerTransactor, error) {
	contract, err := bindCmaccountmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerTransactor{contract: contract}, nil
}

// NewCmaccountmanagerFilterer creates a new log filterer instance of Cmaccountmanager, bound to a specific deployed contract.
func NewCmaccountmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*CmaccountmanagerFilterer, error) {
	contract, err := bindCmaccountmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerFilterer{contract: contract}, nil
}

// bindCmaccountmanager binds a generic wrapper to an already deployed contract.
func bindCmaccountmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CmaccountmanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cmaccountmanager *CmaccountmanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cmaccountmanager.Contract.CmaccountmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cmaccountmanager *CmaccountmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.CmaccountmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cmaccountmanager *CmaccountmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.CmaccountmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cmaccountmanager *CmaccountmanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cmaccountmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cmaccountmanager *CmaccountmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cmaccountmanager *CmaccountmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.contract.Transact(opts, method, params...)
}

// CMACCOUNTROLE is a free data retrieval call binding the contract method 0x7fa34657.
//
// Solidity: function CMACCOUNT_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCaller) CMACCOUNTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "CMACCOUNT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CMACCOUNTROLE is a free data retrieval call binding the contract method 0x7fa34657.
//
// Solidity: function CMACCOUNT_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerSession) CMACCOUNTROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.CMACCOUNTROLE(&_Cmaccountmanager.CallOpts)
}

// CMACCOUNTROLE is a free data retrieval call binding the contract method 0x7fa34657.
//
// Solidity: function CMACCOUNT_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) CMACCOUNTROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.CMACCOUNTROLE(&_Cmaccountmanager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.DEFAULTADMINROLE(&_Cmaccountmanager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.DEFAULTADMINROLE(&_Cmaccountmanager.CallOpts)
}

// DEVELOPERWALLETADMINROLE is a free data retrieval call binding the contract method 0x40df8d84.
//
// Solidity: function DEVELOPER_WALLET_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCaller) DEVELOPERWALLETADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "DEVELOPER_WALLET_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEVELOPERWALLETADMINROLE is a free data retrieval call binding the contract method 0x40df8d84.
//
// Solidity: function DEVELOPER_WALLET_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerSession) DEVELOPERWALLETADMINROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.DEVELOPERWALLETADMINROLE(&_Cmaccountmanager.CallOpts)
}

// DEVELOPERWALLETADMINROLE is a free data retrieval call binding the contract method 0x40df8d84.
//
// Solidity: function DEVELOPER_WALLET_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) DEVELOPERWALLETADMINROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.DEVELOPERWALLETADMINROLE(&_Cmaccountmanager.CallOpts)
}

// FEEADMINROLE is a free data retrieval call binding the contract method 0x4cba593a.
//
// Solidity: function FEE_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCaller) FEEADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "FEE_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FEEADMINROLE is a free data retrieval call binding the contract method 0x4cba593a.
//
// Solidity: function FEE_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerSession) FEEADMINROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.FEEADMINROLE(&_Cmaccountmanager.CallOpts)
}

// FEEADMINROLE is a free data retrieval call binding the contract method 0x4cba593a.
//
// Solidity: function FEE_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) FEEADMINROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.FEEADMINROLE(&_Cmaccountmanager.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerSession) PAUSERROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.PAUSERROLE(&_Cmaccountmanager.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.PAUSERROLE(&_Cmaccountmanager.CallOpts)
}

// PREFUNDADMINROLE is a free data retrieval call binding the contract method 0x01601d2f.
//
// Solidity: function PREFUND_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCaller) PREFUNDADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "PREFUND_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PREFUNDADMINROLE is a free data retrieval call binding the contract method 0x01601d2f.
//
// Solidity: function PREFUND_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerSession) PREFUNDADMINROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.PREFUNDADMINROLE(&_Cmaccountmanager.CallOpts)
}

// PREFUNDADMINROLE is a free data retrieval call binding the contract method 0x01601d2f.
//
// Solidity: function PREFUND_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) PREFUNDADMINROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.PREFUNDADMINROLE(&_Cmaccountmanager.CallOpts)
}

// SERVICEREGISTRYADMINROLE is a free data retrieval call binding the contract method 0x92cf833f.
//
// Solidity: function SERVICE_REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCaller) SERVICEREGISTRYADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "SERVICE_REGISTRY_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SERVICEREGISTRYADMINROLE is a free data retrieval call binding the contract method 0x92cf833f.
//
// Solidity: function SERVICE_REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerSession) SERVICEREGISTRYADMINROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.SERVICEREGISTRYADMINROLE(&_Cmaccountmanager.CallOpts)
}

// SERVICEREGISTRYADMINROLE is a free data retrieval call binding the contract method 0x92cf833f.
//
// Solidity: function SERVICE_REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) SERVICEREGISTRYADMINROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.SERVICEREGISTRYADMINROLE(&_Cmaccountmanager.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerSession) UPGRADERROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.UPGRADERROLE(&_Cmaccountmanager.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.UPGRADERROLE(&_Cmaccountmanager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Cmaccountmanager *CmaccountmanagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Cmaccountmanager *CmaccountmanagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Cmaccountmanager.Contract.UPGRADEINTERFACEVERSION(&_Cmaccountmanager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Cmaccountmanager.Contract.UPGRADEINTERFACEVERSION(&_Cmaccountmanager.CallOpts)
}

// VERSIONERROLE is a free data retrieval call binding the contract method 0xb289819c.
//
// Solidity: function VERSIONER_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCaller) VERSIONERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "VERSIONER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERSIONERROLE is a free data retrieval call binding the contract method 0xb289819c.
//
// Solidity: function VERSIONER_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerSession) VERSIONERROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.VERSIONERROLE(&_Cmaccountmanager.CallOpts)
}

// VERSIONERROLE is a free data retrieval call binding the contract method 0xb289819c.
//
// Solidity: function VERSIONER_ROLE() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) VERSIONERROLE() ([32]byte, error) {
	return _Cmaccountmanager.Contract.VERSIONERROLE(&_Cmaccountmanager.CallOpts)
}

// GetAccountImplementation is a free data retrieval call binding the contract method 0x9d825bc5.
//
// Solidity: function getAccountImplementation() view returns(address)
func (_Cmaccountmanager *CmaccountmanagerCaller) GetAccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "getAccountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountImplementation is a free data retrieval call binding the contract method 0x9d825bc5.
//
// Solidity: function getAccountImplementation() view returns(address)
func (_Cmaccountmanager *CmaccountmanagerSession) GetAccountImplementation() (common.Address, error) {
	return _Cmaccountmanager.Contract.GetAccountImplementation(&_Cmaccountmanager.CallOpts)
}

// GetAccountImplementation is a free data retrieval call binding the contract method 0x9d825bc5.
//
// Solidity: function getAccountImplementation() view returns(address)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) GetAccountImplementation() (common.Address, error) {
	return _Cmaccountmanager.Contract.GetAccountImplementation(&_Cmaccountmanager.CallOpts)
}

// GetAllRegisteredServiceHashes is a free data retrieval call binding the contract method 0x8127f465.
//
// Solidity: function getAllRegisteredServiceHashes() view returns(bytes32[] services)
func (_Cmaccountmanager *CmaccountmanagerCaller) GetAllRegisteredServiceHashes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "getAllRegisteredServiceHashes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllRegisteredServiceHashes is a free data retrieval call binding the contract method 0x8127f465.
//
// Solidity: function getAllRegisteredServiceHashes() view returns(bytes32[] services)
func (_Cmaccountmanager *CmaccountmanagerSession) GetAllRegisteredServiceHashes() ([][32]byte, error) {
	return _Cmaccountmanager.Contract.GetAllRegisteredServiceHashes(&_Cmaccountmanager.CallOpts)
}

// GetAllRegisteredServiceHashes is a free data retrieval call binding the contract method 0x8127f465.
//
// Solidity: function getAllRegisteredServiceHashes() view returns(bytes32[] services)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) GetAllRegisteredServiceHashes() ([][32]byte, error) {
	return _Cmaccountmanager.Contract.GetAllRegisteredServiceHashes(&_Cmaccountmanager.CallOpts)
}

// GetAllRegisteredServiceNames is a free data retrieval call binding the contract method 0xc2a1db78.
//
// Solidity: function getAllRegisteredServiceNames() view returns(string[] services)
func (_Cmaccountmanager *CmaccountmanagerCaller) GetAllRegisteredServiceNames(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "getAllRegisteredServiceNames")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllRegisteredServiceNames is a free data retrieval call binding the contract method 0xc2a1db78.
//
// Solidity: function getAllRegisteredServiceNames() view returns(string[] services)
func (_Cmaccountmanager *CmaccountmanagerSession) GetAllRegisteredServiceNames() ([]string, error) {
	return _Cmaccountmanager.Contract.GetAllRegisteredServiceNames(&_Cmaccountmanager.CallOpts)
}

// GetAllRegisteredServiceNames is a free data retrieval call binding the contract method 0xc2a1db78.
//
// Solidity: function getAllRegisteredServiceNames() view returns(string[] services)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) GetAllRegisteredServiceNames() ([]string, error) {
	return _Cmaccountmanager.Contract.GetAllRegisteredServiceNames(&_Cmaccountmanager.CallOpts)
}

// GetBookingTokenAddress is a free data retrieval call binding the contract method 0x4f3f4639.
//
// Solidity: function getBookingTokenAddress() view returns(address)
func (_Cmaccountmanager *CmaccountmanagerCaller) GetBookingTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "getBookingTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBookingTokenAddress is a free data retrieval call binding the contract method 0x4f3f4639.
//
// Solidity: function getBookingTokenAddress() view returns(address)
func (_Cmaccountmanager *CmaccountmanagerSession) GetBookingTokenAddress() (common.Address, error) {
	return _Cmaccountmanager.Contract.GetBookingTokenAddress(&_Cmaccountmanager.CallOpts)
}

// GetBookingTokenAddress is a free data retrieval call binding the contract method 0x4f3f4639.
//
// Solidity: function getBookingTokenAddress() view returns(address)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) GetBookingTokenAddress() (common.Address, error) {
	return _Cmaccountmanager.Contract.GetBookingTokenAddress(&_Cmaccountmanager.CallOpts)
}

// GetCMAccountCreator is a free data retrieval call binding the contract method 0x2cec1a06.
//
// Solidity: function getCMAccountCreator(address account) view returns(address)
func (_Cmaccountmanager *CmaccountmanagerCaller) GetCMAccountCreator(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "getCMAccountCreator", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCMAccountCreator is a free data retrieval call binding the contract method 0x2cec1a06.
//
// Solidity: function getCMAccountCreator(address account) view returns(address)
func (_Cmaccountmanager *CmaccountmanagerSession) GetCMAccountCreator(account common.Address) (common.Address, error) {
	return _Cmaccountmanager.Contract.GetCMAccountCreator(&_Cmaccountmanager.CallOpts, account)
}

// GetCMAccountCreator is a free data retrieval call binding the contract method 0x2cec1a06.
//
// Solidity: function getCMAccountCreator(address account) view returns(address)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) GetCMAccountCreator(account common.Address) (common.Address, error) {
	return _Cmaccountmanager.Contract.GetCMAccountCreator(&_Cmaccountmanager.CallOpts, account)
}

// GetDeveloperFeeBp is a free data retrieval call binding the contract method 0x3c555938.
//
// Solidity: function getDeveloperFeeBp() view returns(uint256 developerFeeBp)
func (_Cmaccountmanager *CmaccountmanagerCaller) GetDeveloperFeeBp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "getDeveloperFeeBp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeveloperFeeBp is a free data retrieval call binding the contract method 0x3c555938.
//
// Solidity: function getDeveloperFeeBp() view returns(uint256 developerFeeBp)
func (_Cmaccountmanager *CmaccountmanagerSession) GetDeveloperFeeBp() (*big.Int, error) {
	return _Cmaccountmanager.Contract.GetDeveloperFeeBp(&_Cmaccountmanager.CallOpts)
}

// GetDeveloperFeeBp is a free data retrieval call binding the contract method 0x3c555938.
//
// Solidity: function getDeveloperFeeBp() view returns(uint256 developerFeeBp)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) GetDeveloperFeeBp() (*big.Int, error) {
	return _Cmaccountmanager.Contract.GetDeveloperFeeBp(&_Cmaccountmanager.CallOpts)
}

// GetDeveloperWallet is a free data retrieval call binding the contract method 0x0470d3ac.
//
// Solidity: function getDeveloperWallet() view returns(address developerWallet)
func (_Cmaccountmanager *CmaccountmanagerCaller) GetDeveloperWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "getDeveloperWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDeveloperWallet is a free data retrieval call binding the contract method 0x0470d3ac.
//
// Solidity: function getDeveloperWallet() view returns(address developerWallet)
func (_Cmaccountmanager *CmaccountmanagerSession) GetDeveloperWallet() (common.Address, error) {
	return _Cmaccountmanager.Contract.GetDeveloperWallet(&_Cmaccountmanager.CallOpts)
}

// GetDeveloperWallet is a free data retrieval call binding the contract method 0x0470d3ac.
//
// Solidity: function getDeveloperWallet() view returns(address developerWallet)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) GetDeveloperWallet() (common.Address, error) {
	return _Cmaccountmanager.Contract.GetDeveloperWallet(&_Cmaccountmanager.CallOpts)
}

// GetPrefundAmount is a free data retrieval call binding the contract method 0xc39409e1.
//
// Solidity: function getPrefundAmount() view returns(uint256)
func (_Cmaccountmanager *CmaccountmanagerCaller) GetPrefundAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "getPrefundAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrefundAmount is a free data retrieval call binding the contract method 0xc39409e1.
//
// Solidity: function getPrefundAmount() view returns(uint256)
func (_Cmaccountmanager *CmaccountmanagerSession) GetPrefundAmount() (*big.Int, error) {
	return _Cmaccountmanager.Contract.GetPrefundAmount(&_Cmaccountmanager.CallOpts)
}

// GetPrefundAmount is a free data retrieval call binding the contract method 0xc39409e1.
//
// Solidity: function getPrefundAmount() view returns(uint256)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) GetPrefundAmount() (*big.Int, error) {
	return _Cmaccountmanager.Contract.GetPrefundAmount(&_Cmaccountmanager.CallOpts)
}

// GetRegisteredServiceHashByName is a free data retrieval call binding the contract method 0x352af39a.
//
// Solidity: function getRegisteredServiceHashByName(string serviceName) view returns(bytes32 serviceHash)
func (_Cmaccountmanager *CmaccountmanagerCaller) GetRegisteredServiceHashByName(opts *bind.CallOpts, serviceName string) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "getRegisteredServiceHashByName", serviceName)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRegisteredServiceHashByName is a free data retrieval call binding the contract method 0x352af39a.
//
// Solidity: function getRegisteredServiceHashByName(string serviceName) view returns(bytes32 serviceHash)
func (_Cmaccountmanager *CmaccountmanagerSession) GetRegisteredServiceHashByName(serviceName string) ([32]byte, error) {
	return _Cmaccountmanager.Contract.GetRegisteredServiceHashByName(&_Cmaccountmanager.CallOpts, serviceName)
}

// GetRegisteredServiceHashByName is a free data retrieval call binding the contract method 0x352af39a.
//
// Solidity: function getRegisteredServiceHashByName(string serviceName) view returns(bytes32 serviceHash)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) GetRegisteredServiceHashByName(serviceName string) ([32]byte, error) {
	return _Cmaccountmanager.Contract.GetRegisteredServiceHashByName(&_Cmaccountmanager.CallOpts, serviceName)
}

// GetRegisteredServiceNameByHash is a free data retrieval call binding the contract method 0x5a81a626.
//
// Solidity: function getRegisteredServiceNameByHash(bytes32 serviceHash) view returns(string serviceName)
func (_Cmaccountmanager *CmaccountmanagerCaller) GetRegisteredServiceNameByHash(opts *bind.CallOpts, serviceHash [32]byte) (string, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "getRegisteredServiceNameByHash", serviceHash)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetRegisteredServiceNameByHash is a free data retrieval call binding the contract method 0x5a81a626.
//
// Solidity: function getRegisteredServiceNameByHash(bytes32 serviceHash) view returns(string serviceName)
func (_Cmaccountmanager *CmaccountmanagerSession) GetRegisteredServiceNameByHash(serviceHash [32]byte) (string, error) {
	return _Cmaccountmanager.Contract.GetRegisteredServiceNameByHash(&_Cmaccountmanager.CallOpts, serviceHash)
}

// GetRegisteredServiceNameByHash is a free data retrieval call binding the contract method 0x5a81a626.
//
// Solidity: function getRegisteredServiceNameByHash(bytes32 serviceHash) view returns(string serviceName)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) GetRegisteredServiceNameByHash(serviceHash [32]byte) (string, error) {
	return _Cmaccountmanager.Contract.GetRegisteredServiceNameByHash(&_Cmaccountmanager.CallOpts, serviceHash)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Cmaccountmanager.Contract.GetRoleAdmin(&_Cmaccountmanager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Cmaccountmanager.Contract.GetRoleAdmin(&_Cmaccountmanager.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Cmaccountmanager *CmaccountmanagerCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Cmaccountmanager *CmaccountmanagerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Cmaccountmanager.Contract.GetRoleMember(&_Cmaccountmanager.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Cmaccountmanager.Contract.GetRoleMember(&_Cmaccountmanager.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Cmaccountmanager *CmaccountmanagerCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Cmaccountmanager *CmaccountmanagerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Cmaccountmanager.Contract.GetRoleMemberCount(&_Cmaccountmanager.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Cmaccountmanager.Contract.GetRoleMemberCount(&_Cmaccountmanager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Cmaccountmanager *CmaccountmanagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Cmaccountmanager *CmaccountmanagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Cmaccountmanager.Contract.HasRole(&_Cmaccountmanager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Cmaccountmanager.Contract.HasRole(&_Cmaccountmanager.CallOpts, role, account)
}

// IsCMAccount is a free data retrieval call binding the contract method 0x12b357b5.
//
// Solidity: function isCMAccount(address account) view returns(bool)
func (_Cmaccountmanager *CmaccountmanagerCaller) IsCMAccount(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "isCMAccount", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCMAccount is a free data retrieval call binding the contract method 0x12b357b5.
//
// Solidity: function isCMAccount(address account) view returns(bool)
func (_Cmaccountmanager *CmaccountmanagerSession) IsCMAccount(account common.Address) (bool, error) {
	return _Cmaccountmanager.Contract.IsCMAccount(&_Cmaccountmanager.CallOpts, account)
}

// IsCMAccount is a free data retrieval call binding the contract method 0x12b357b5.
//
// Solidity: function isCMAccount(address account) view returns(bool)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) IsCMAccount(account common.Address) (bool, error) {
	return _Cmaccountmanager.Contract.IsCMAccount(&_Cmaccountmanager.CallOpts, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Cmaccountmanager *CmaccountmanagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Cmaccountmanager *CmaccountmanagerSession) Paused() (bool, error) {
	return _Cmaccountmanager.Contract.Paused(&_Cmaccountmanager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) Paused() (bool, error) {
	return _Cmaccountmanager.Contract.Paused(&_Cmaccountmanager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerSession) ProxiableUUID() ([32]byte, error) {
	return _Cmaccountmanager.Contract.ProxiableUUID(&_Cmaccountmanager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Cmaccountmanager.Contract.ProxiableUUID(&_Cmaccountmanager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cmaccountmanager *CmaccountmanagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Cmaccountmanager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cmaccountmanager *CmaccountmanagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cmaccountmanager.Contract.SupportsInterface(&_Cmaccountmanager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cmaccountmanager *CmaccountmanagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cmaccountmanager.Contract.SupportsInterface(&_Cmaccountmanager.CallOpts, interfaceId)
}

// CreateCMAccount is a paid mutator transaction binding the contract method 0xfd2fbd52.
//
// Solidity: function createCMAccount(address admin, address upgrader) payable returns(address)
func (_Cmaccountmanager *CmaccountmanagerTransactor) CreateCMAccount(opts *bind.TransactOpts, admin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "createCMAccount", admin, upgrader)
}

// CreateCMAccount is a paid mutator transaction binding the contract method 0xfd2fbd52.
//
// Solidity: function createCMAccount(address admin, address upgrader) payable returns(address)
func (_Cmaccountmanager *CmaccountmanagerSession) CreateCMAccount(admin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.CreateCMAccount(&_Cmaccountmanager.TransactOpts, admin, upgrader)
}

// CreateCMAccount is a paid mutator transaction binding the contract method 0xfd2fbd52.
//
// Solidity: function createCMAccount(address admin, address upgrader) payable returns(address)
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) CreateCMAccount(admin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.CreateCMAccount(&_Cmaccountmanager.TransactOpts, admin, upgrader)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Cmaccountmanager *CmaccountmanagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.GrantRole(&_Cmaccountmanager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.GrantRole(&_Cmaccountmanager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x95b6ef0c.
//
// Solidity: function initialize(address defaultAdmin, address pauser, address upgrader, address versioner, address developerWallet, uint256 developerFeeBp) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactor) Initialize(opts *bind.TransactOpts, defaultAdmin common.Address, pauser common.Address, upgrader common.Address, versioner common.Address, developerWallet common.Address, developerFeeBp *big.Int) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "initialize", defaultAdmin, pauser, upgrader, versioner, developerWallet, developerFeeBp)
}

// Initialize is a paid mutator transaction binding the contract method 0x95b6ef0c.
//
// Solidity: function initialize(address defaultAdmin, address pauser, address upgrader, address versioner, address developerWallet, uint256 developerFeeBp) returns()
func (_Cmaccountmanager *CmaccountmanagerSession) Initialize(defaultAdmin common.Address, pauser common.Address, upgrader common.Address, versioner common.Address, developerWallet common.Address, developerFeeBp *big.Int) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.Initialize(&_Cmaccountmanager.TransactOpts, defaultAdmin, pauser, upgrader, versioner, developerWallet, developerFeeBp)
}

// Initialize is a paid mutator transaction binding the contract method 0x95b6ef0c.
//
// Solidity: function initialize(address defaultAdmin, address pauser, address upgrader, address versioner, address developerWallet, uint256 developerFeeBp) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) Initialize(defaultAdmin common.Address, pauser common.Address, upgrader common.Address, versioner common.Address, developerWallet common.Address, developerFeeBp *big.Int) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.Initialize(&_Cmaccountmanager.TransactOpts, defaultAdmin, pauser, upgrader, versioner, developerWallet, developerFeeBp)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Cmaccountmanager *CmaccountmanagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Cmaccountmanager *CmaccountmanagerSession) Pause() (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.Pause(&_Cmaccountmanager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) Pause() (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.Pause(&_Cmaccountmanager.TransactOpts)
}

// RegisterService is a paid mutator transaction binding the contract method 0x0d6115d0.
//
// Solidity: function registerService(string serviceName) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactor) RegisterService(opts *bind.TransactOpts, serviceName string) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "registerService", serviceName)
}

// RegisterService is a paid mutator transaction binding the contract method 0x0d6115d0.
//
// Solidity: function registerService(string serviceName) returns()
func (_Cmaccountmanager *CmaccountmanagerSession) RegisterService(serviceName string) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.RegisterService(&_Cmaccountmanager.TransactOpts, serviceName)
}

// RegisterService is a paid mutator transaction binding the contract method 0x0d6115d0.
//
// Solidity: function registerService(string serviceName) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) RegisterService(serviceName string) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.RegisterService(&_Cmaccountmanager.TransactOpts, serviceName)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Cmaccountmanager *CmaccountmanagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.RenounceRole(&_Cmaccountmanager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.RenounceRole(&_Cmaccountmanager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Cmaccountmanager *CmaccountmanagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.RevokeRole(&_Cmaccountmanager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.RevokeRole(&_Cmaccountmanager.TransactOpts, role, account)
}

// SetAccountImplementation is a paid mutator transaction binding the contract method 0x09766da2.
//
// Solidity: function setAccountImplementation(address newImplementation) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactor) SetAccountImplementation(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "setAccountImplementation", newImplementation)
}

// SetAccountImplementation is a paid mutator transaction binding the contract method 0x09766da2.
//
// Solidity: function setAccountImplementation(address newImplementation) returns()
func (_Cmaccountmanager *CmaccountmanagerSession) SetAccountImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.SetAccountImplementation(&_Cmaccountmanager.TransactOpts, newImplementation)
}

// SetAccountImplementation is a paid mutator transaction binding the contract method 0x09766da2.
//
// Solidity: function setAccountImplementation(address newImplementation) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) SetAccountImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.SetAccountImplementation(&_Cmaccountmanager.TransactOpts, newImplementation)
}

// SetBookingTokenAddress is a paid mutator transaction binding the contract method 0xc17e30bf.
//
// Solidity: function setBookingTokenAddress(address token) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactor) SetBookingTokenAddress(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "setBookingTokenAddress", token)
}

// SetBookingTokenAddress is a paid mutator transaction binding the contract method 0xc17e30bf.
//
// Solidity: function setBookingTokenAddress(address token) returns()
func (_Cmaccountmanager *CmaccountmanagerSession) SetBookingTokenAddress(token common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.SetBookingTokenAddress(&_Cmaccountmanager.TransactOpts, token)
}

// SetBookingTokenAddress is a paid mutator transaction binding the contract method 0xc17e30bf.
//
// Solidity: function setBookingTokenAddress(address token) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) SetBookingTokenAddress(token common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.SetBookingTokenAddress(&_Cmaccountmanager.TransactOpts, token)
}

// SetDeveloperFeeBp is a paid mutator transaction binding the contract method 0xf85b4a9c.
//
// Solidity: function setDeveloperFeeBp(uint256 bp) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactor) SetDeveloperFeeBp(opts *bind.TransactOpts, bp *big.Int) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "setDeveloperFeeBp", bp)
}

// SetDeveloperFeeBp is a paid mutator transaction binding the contract method 0xf85b4a9c.
//
// Solidity: function setDeveloperFeeBp(uint256 bp) returns()
func (_Cmaccountmanager *CmaccountmanagerSession) SetDeveloperFeeBp(bp *big.Int) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.SetDeveloperFeeBp(&_Cmaccountmanager.TransactOpts, bp)
}

// SetDeveloperFeeBp is a paid mutator transaction binding the contract method 0xf85b4a9c.
//
// Solidity: function setDeveloperFeeBp(uint256 bp) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) SetDeveloperFeeBp(bp *big.Int) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.SetDeveloperFeeBp(&_Cmaccountmanager.TransactOpts, bp)
}

// SetDeveloperWallet is a paid mutator transaction binding the contract method 0x6cd56878.
//
// Solidity: function setDeveloperWallet(address developerWallet) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactor) SetDeveloperWallet(opts *bind.TransactOpts, developerWallet common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "setDeveloperWallet", developerWallet)
}

// SetDeveloperWallet is a paid mutator transaction binding the contract method 0x6cd56878.
//
// Solidity: function setDeveloperWallet(address developerWallet) returns()
func (_Cmaccountmanager *CmaccountmanagerSession) SetDeveloperWallet(developerWallet common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.SetDeveloperWallet(&_Cmaccountmanager.TransactOpts, developerWallet)
}

// SetDeveloperWallet is a paid mutator transaction binding the contract method 0x6cd56878.
//
// Solidity: function setDeveloperWallet(address developerWallet) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) SetDeveloperWallet(developerWallet common.Address) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.SetDeveloperWallet(&_Cmaccountmanager.TransactOpts, developerWallet)
}

// SetPrefundAmount is a paid mutator transaction binding the contract method 0x1c2b3afc.
//
// Solidity: function setPrefundAmount(uint256 newPrefundAmount) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactor) SetPrefundAmount(opts *bind.TransactOpts, newPrefundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "setPrefundAmount", newPrefundAmount)
}

// SetPrefundAmount is a paid mutator transaction binding the contract method 0x1c2b3afc.
//
// Solidity: function setPrefundAmount(uint256 newPrefundAmount) returns()
func (_Cmaccountmanager *CmaccountmanagerSession) SetPrefundAmount(newPrefundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.SetPrefundAmount(&_Cmaccountmanager.TransactOpts, newPrefundAmount)
}

// SetPrefundAmount is a paid mutator transaction binding the contract method 0x1c2b3afc.
//
// Solidity: function setPrefundAmount(uint256 newPrefundAmount) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) SetPrefundAmount(newPrefundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.SetPrefundAmount(&_Cmaccountmanager.TransactOpts, newPrefundAmount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Cmaccountmanager *CmaccountmanagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Cmaccountmanager *CmaccountmanagerSession) Unpause() (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.Unpause(&_Cmaccountmanager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.Unpause(&_Cmaccountmanager.TransactOpts)
}

// UnregisterService is a paid mutator transaction binding the contract method 0x5e818619.
//
// Solidity: function unregisterService(string serviceName) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactor) UnregisterService(opts *bind.TransactOpts, serviceName string) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "unregisterService", serviceName)
}

// UnregisterService is a paid mutator transaction binding the contract method 0x5e818619.
//
// Solidity: function unregisterService(string serviceName) returns()
func (_Cmaccountmanager *CmaccountmanagerSession) UnregisterService(serviceName string) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.UnregisterService(&_Cmaccountmanager.TransactOpts, serviceName)
}

// UnregisterService is a paid mutator transaction binding the contract method 0x5e818619.
//
// Solidity: function unregisterService(string serviceName) returns()
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) UnregisterService(serviceName string) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.UnregisterService(&_Cmaccountmanager.TransactOpts, serviceName)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Cmaccountmanager *CmaccountmanagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccountmanager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Cmaccountmanager *CmaccountmanagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.UpgradeToAndCall(&_Cmaccountmanager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Cmaccountmanager *CmaccountmanagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccountmanager.Contract.UpgradeToAndCall(&_Cmaccountmanager.TransactOpts, newImplementation, data)
}

// CmaccountmanagerBookingTokenAddressUpdatedIterator is returned from FilterBookingTokenAddressUpdated and is used to iterate over the raw logs and unpacked data for BookingTokenAddressUpdated events raised by the Cmaccountmanager contract.
type CmaccountmanagerBookingTokenAddressUpdatedIterator struct {
	Event *CmaccountmanagerBookingTokenAddressUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountmanagerBookingTokenAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountmanagerBookingTokenAddressUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountmanagerBookingTokenAddressUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountmanagerBookingTokenAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountmanagerBookingTokenAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountmanagerBookingTokenAddressUpdated represents a BookingTokenAddressUpdated event raised by the Cmaccountmanager contract.
type CmaccountmanagerBookingTokenAddressUpdated struct {
	OldBookingToken common.Address
	NewBookingToken common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBookingTokenAddressUpdated is a free log retrieval operation binding the contract event 0x8752daf55fa0ccd919c149bbc809abaad4433f02ba7aa93ac6d5acfd2f8dc22c.
//
// Solidity: event BookingTokenAddressUpdated(address indexed oldBookingToken, address indexed newBookingToken)
func (_Cmaccountmanager *CmaccountmanagerFilterer) FilterBookingTokenAddressUpdated(opts *bind.FilterOpts, oldBookingToken []common.Address, newBookingToken []common.Address) (*CmaccountmanagerBookingTokenAddressUpdatedIterator, error) {

	var oldBookingTokenRule []interface{}
	for _, oldBookingTokenItem := range oldBookingToken {
		oldBookingTokenRule = append(oldBookingTokenRule, oldBookingTokenItem)
	}
	var newBookingTokenRule []interface{}
	for _, newBookingTokenItem := range newBookingToken {
		newBookingTokenRule = append(newBookingTokenRule, newBookingTokenItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.FilterLogs(opts, "BookingTokenAddressUpdated", oldBookingTokenRule, newBookingTokenRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerBookingTokenAddressUpdatedIterator{contract: _Cmaccountmanager.contract, event: "BookingTokenAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchBookingTokenAddressUpdated is a free log subscription operation binding the contract event 0x8752daf55fa0ccd919c149bbc809abaad4433f02ba7aa93ac6d5acfd2f8dc22c.
//
// Solidity: event BookingTokenAddressUpdated(address indexed oldBookingToken, address indexed newBookingToken)
func (_Cmaccountmanager *CmaccountmanagerFilterer) WatchBookingTokenAddressUpdated(opts *bind.WatchOpts, sink chan<- *CmaccountmanagerBookingTokenAddressUpdated, oldBookingToken []common.Address, newBookingToken []common.Address) (event.Subscription, error) {

	var oldBookingTokenRule []interface{}
	for _, oldBookingTokenItem := range oldBookingToken {
		oldBookingTokenRule = append(oldBookingTokenRule, oldBookingTokenItem)
	}
	var newBookingTokenRule []interface{}
	for _, newBookingTokenItem := range newBookingToken {
		newBookingTokenRule = append(newBookingTokenRule, newBookingTokenItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.WatchLogs(opts, "BookingTokenAddressUpdated", oldBookingTokenRule, newBookingTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountmanagerBookingTokenAddressUpdated)
				if err := _Cmaccountmanager.contract.UnpackLog(event, "BookingTokenAddressUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBookingTokenAddressUpdated is a log parse operation binding the contract event 0x8752daf55fa0ccd919c149bbc809abaad4433f02ba7aa93ac6d5acfd2f8dc22c.
//
// Solidity: event BookingTokenAddressUpdated(address indexed oldBookingToken, address indexed newBookingToken)
func (_Cmaccountmanager *CmaccountmanagerFilterer) ParseBookingTokenAddressUpdated(log types.Log) (*CmaccountmanagerBookingTokenAddressUpdated, error) {
	event := new(CmaccountmanagerBookingTokenAddressUpdated)
	if err := _Cmaccountmanager.contract.UnpackLog(event, "BookingTokenAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountmanagerCMAccountCreatedIterator is returned from FilterCMAccountCreated and is used to iterate over the raw logs and unpacked data for CMAccountCreated events raised by the Cmaccountmanager contract.
type CmaccountmanagerCMAccountCreatedIterator struct {
	Event *CmaccountmanagerCMAccountCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountmanagerCMAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountmanagerCMAccountCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountmanagerCMAccountCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountmanagerCMAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountmanagerCMAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountmanagerCMAccountCreated represents a CMAccountCreated event raised by the Cmaccountmanager contract.
type CmaccountmanagerCMAccountCreated struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCMAccountCreated is a free log retrieval operation binding the contract event 0x22de16ef21e3a33810fbcaf0e737ab9c9e2854fa565d8535041456c789afcd93.
//
// Solidity: event CMAccountCreated(address indexed account)
func (_Cmaccountmanager *CmaccountmanagerFilterer) FilterCMAccountCreated(opts *bind.FilterOpts, account []common.Address) (*CmaccountmanagerCMAccountCreatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.FilterLogs(opts, "CMAccountCreated", accountRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerCMAccountCreatedIterator{contract: _Cmaccountmanager.contract, event: "CMAccountCreated", logs: logs, sub: sub}, nil
}

// WatchCMAccountCreated is a free log subscription operation binding the contract event 0x22de16ef21e3a33810fbcaf0e737ab9c9e2854fa565d8535041456c789afcd93.
//
// Solidity: event CMAccountCreated(address indexed account)
func (_Cmaccountmanager *CmaccountmanagerFilterer) WatchCMAccountCreated(opts *bind.WatchOpts, sink chan<- *CmaccountmanagerCMAccountCreated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.WatchLogs(opts, "CMAccountCreated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountmanagerCMAccountCreated)
				if err := _Cmaccountmanager.contract.UnpackLog(event, "CMAccountCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCMAccountCreated is a log parse operation binding the contract event 0x22de16ef21e3a33810fbcaf0e737ab9c9e2854fa565d8535041456c789afcd93.
//
// Solidity: event CMAccountCreated(address indexed account)
func (_Cmaccountmanager *CmaccountmanagerFilterer) ParseCMAccountCreated(log types.Log) (*CmaccountmanagerCMAccountCreated, error) {
	event := new(CmaccountmanagerCMAccountCreated)
	if err := _Cmaccountmanager.contract.UnpackLog(event, "CMAccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountmanagerCMAccountImplementationUpdatedIterator is returned from FilterCMAccountImplementationUpdated and is used to iterate over the raw logs and unpacked data for CMAccountImplementationUpdated events raised by the Cmaccountmanager contract.
type CmaccountmanagerCMAccountImplementationUpdatedIterator struct {
	Event *CmaccountmanagerCMAccountImplementationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountmanagerCMAccountImplementationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountmanagerCMAccountImplementationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountmanagerCMAccountImplementationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountmanagerCMAccountImplementationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountmanagerCMAccountImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountmanagerCMAccountImplementationUpdated represents a CMAccountImplementationUpdated event raised by the Cmaccountmanager contract.
type CmaccountmanagerCMAccountImplementationUpdated struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCMAccountImplementationUpdated is a free log retrieval operation binding the contract event 0xecd599290ada9c149ae0fb600bc6ada3c103a716eadb877b5341b7d298b62d93.
//
// Solidity: event CMAccountImplementationUpdated(address indexed oldImplementation, address indexed newImplementation)
func (_Cmaccountmanager *CmaccountmanagerFilterer) FilterCMAccountImplementationUpdated(opts *bind.FilterOpts, oldImplementation []common.Address, newImplementation []common.Address) (*CmaccountmanagerCMAccountImplementationUpdatedIterator, error) {

	var oldImplementationRule []interface{}
	for _, oldImplementationItem := range oldImplementation {
		oldImplementationRule = append(oldImplementationRule, oldImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.FilterLogs(opts, "CMAccountImplementationUpdated", oldImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerCMAccountImplementationUpdatedIterator{contract: _Cmaccountmanager.contract, event: "CMAccountImplementationUpdated", logs: logs, sub: sub}, nil
}

// WatchCMAccountImplementationUpdated is a free log subscription operation binding the contract event 0xecd599290ada9c149ae0fb600bc6ada3c103a716eadb877b5341b7d298b62d93.
//
// Solidity: event CMAccountImplementationUpdated(address indexed oldImplementation, address indexed newImplementation)
func (_Cmaccountmanager *CmaccountmanagerFilterer) WatchCMAccountImplementationUpdated(opts *bind.WatchOpts, sink chan<- *CmaccountmanagerCMAccountImplementationUpdated, oldImplementation []common.Address, newImplementation []common.Address) (event.Subscription, error) {

	var oldImplementationRule []interface{}
	for _, oldImplementationItem := range oldImplementation {
		oldImplementationRule = append(oldImplementationRule, oldImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.WatchLogs(opts, "CMAccountImplementationUpdated", oldImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountmanagerCMAccountImplementationUpdated)
				if err := _Cmaccountmanager.contract.UnpackLog(event, "CMAccountImplementationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCMAccountImplementationUpdated is a log parse operation binding the contract event 0xecd599290ada9c149ae0fb600bc6ada3c103a716eadb877b5341b7d298b62d93.
//
// Solidity: event CMAccountImplementationUpdated(address indexed oldImplementation, address indexed newImplementation)
func (_Cmaccountmanager *CmaccountmanagerFilterer) ParseCMAccountImplementationUpdated(log types.Log) (*CmaccountmanagerCMAccountImplementationUpdated, error) {
	event := new(CmaccountmanagerCMAccountImplementationUpdated)
	if err := _Cmaccountmanager.contract.UnpackLog(event, "CMAccountImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountmanagerDeveloperFeeBpUpdatedIterator is returned from FilterDeveloperFeeBpUpdated and is used to iterate over the raw logs and unpacked data for DeveloperFeeBpUpdated events raised by the Cmaccountmanager contract.
type CmaccountmanagerDeveloperFeeBpUpdatedIterator struct {
	Event *CmaccountmanagerDeveloperFeeBpUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountmanagerDeveloperFeeBpUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountmanagerDeveloperFeeBpUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountmanagerDeveloperFeeBpUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountmanagerDeveloperFeeBpUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountmanagerDeveloperFeeBpUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountmanagerDeveloperFeeBpUpdated represents a DeveloperFeeBpUpdated event raised by the Cmaccountmanager contract.
type CmaccountmanagerDeveloperFeeBpUpdated struct {
	OldDeveloperFeeBp *big.Int
	NewDeveloperFeeBp *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDeveloperFeeBpUpdated is a free log retrieval operation binding the contract event 0xc1fec27bbb8c533fedd3d449016c429341bfc0df21a9600e6162605e08f1c78f.
//
// Solidity: event DeveloperFeeBpUpdated(uint256 indexed oldDeveloperFeeBp, uint256 indexed newDeveloperFeeBp)
func (_Cmaccountmanager *CmaccountmanagerFilterer) FilterDeveloperFeeBpUpdated(opts *bind.FilterOpts, oldDeveloperFeeBp []*big.Int, newDeveloperFeeBp []*big.Int) (*CmaccountmanagerDeveloperFeeBpUpdatedIterator, error) {

	var oldDeveloperFeeBpRule []interface{}
	for _, oldDeveloperFeeBpItem := range oldDeveloperFeeBp {
		oldDeveloperFeeBpRule = append(oldDeveloperFeeBpRule, oldDeveloperFeeBpItem)
	}
	var newDeveloperFeeBpRule []interface{}
	for _, newDeveloperFeeBpItem := range newDeveloperFeeBp {
		newDeveloperFeeBpRule = append(newDeveloperFeeBpRule, newDeveloperFeeBpItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.FilterLogs(opts, "DeveloperFeeBpUpdated", oldDeveloperFeeBpRule, newDeveloperFeeBpRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerDeveloperFeeBpUpdatedIterator{contract: _Cmaccountmanager.contract, event: "DeveloperFeeBpUpdated", logs: logs, sub: sub}, nil
}

// WatchDeveloperFeeBpUpdated is a free log subscription operation binding the contract event 0xc1fec27bbb8c533fedd3d449016c429341bfc0df21a9600e6162605e08f1c78f.
//
// Solidity: event DeveloperFeeBpUpdated(uint256 indexed oldDeveloperFeeBp, uint256 indexed newDeveloperFeeBp)
func (_Cmaccountmanager *CmaccountmanagerFilterer) WatchDeveloperFeeBpUpdated(opts *bind.WatchOpts, sink chan<- *CmaccountmanagerDeveloperFeeBpUpdated, oldDeveloperFeeBp []*big.Int, newDeveloperFeeBp []*big.Int) (event.Subscription, error) {

	var oldDeveloperFeeBpRule []interface{}
	for _, oldDeveloperFeeBpItem := range oldDeveloperFeeBp {
		oldDeveloperFeeBpRule = append(oldDeveloperFeeBpRule, oldDeveloperFeeBpItem)
	}
	var newDeveloperFeeBpRule []interface{}
	for _, newDeveloperFeeBpItem := range newDeveloperFeeBp {
		newDeveloperFeeBpRule = append(newDeveloperFeeBpRule, newDeveloperFeeBpItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.WatchLogs(opts, "DeveloperFeeBpUpdated", oldDeveloperFeeBpRule, newDeveloperFeeBpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountmanagerDeveloperFeeBpUpdated)
				if err := _Cmaccountmanager.contract.UnpackLog(event, "DeveloperFeeBpUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeveloperFeeBpUpdated is a log parse operation binding the contract event 0xc1fec27bbb8c533fedd3d449016c429341bfc0df21a9600e6162605e08f1c78f.
//
// Solidity: event DeveloperFeeBpUpdated(uint256 indexed oldDeveloperFeeBp, uint256 indexed newDeveloperFeeBp)
func (_Cmaccountmanager *CmaccountmanagerFilterer) ParseDeveloperFeeBpUpdated(log types.Log) (*CmaccountmanagerDeveloperFeeBpUpdated, error) {
	event := new(CmaccountmanagerDeveloperFeeBpUpdated)
	if err := _Cmaccountmanager.contract.UnpackLog(event, "DeveloperFeeBpUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountmanagerDeveloperWalletUpdatedIterator is returned from FilterDeveloperWalletUpdated and is used to iterate over the raw logs and unpacked data for DeveloperWalletUpdated events raised by the Cmaccountmanager contract.
type CmaccountmanagerDeveloperWalletUpdatedIterator struct {
	Event *CmaccountmanagerDeveloperWalletUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountmanagerDeveloperWalletUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountmanagerDeveloperWalletUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountmanagerDeveloperWalletUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountmanagerDeveloperWalletUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountmanagerDeveloperWalletUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountmanagerDeveloperWalletUpdated represents a DeveloperWalletUpdated event raised by the Cmaccountmanager contract.
type CmaccountmanagerDeveloperWalletUpdated struct {
	OldDeveloperWallet common.Address
	NewDeveloperWallet common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeveloperWalletUpdated is a free log retrieval operation binding the contract event 0x88bf70d4f3c446213b5064d7cfe95ec0ed196748f014c19a833117bac32468fd.
//
// Solidity: event DeveloperWalletUpdated(address indexed oldDeveloperWallet, address indexed newDeveloperWallet)
func (_Cmaccountmanager *CmaccountmanagerFilterer) FilterDeveloperWalletUpdated(opts *bind.FilterOpts, oldDeveloperWallet []common.Address, newDeveloperWallet []common.Address) (*CmaccountmanagerDeveloperWalletUpdatedIterator, error) {

	var oldDeveloperWalletRule []interface{}
	for _, oldDeveloperWalletItem := range oldDeveloperWallet {
		oldDeveloperWalletRule = append(oldDeveloperWalletRule, oldDeveloperWalletItem)
	}
	var newDeveloperWalletRule []interface{}
	for _, newDeveloperWalletItem := range newDeveloperWallet {
		newDeveloperWalletRule = append(newDeveloperWalletRule, newDeveloperWalletItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.FilterLogs(opts, "DeveloperWalletUpdated", oldDeveloperWalletRule, newDeveloperWalletRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerDeveloperWalletUpdatedIterator{contract: _Cmaccountmanager.contract, event: "DeveloperWalletUpdated", logs: logs, sub: sub}, nil
}

// WatchDeveloperWalletUpdated is a free log subscription operation binding the contract event 0x88bf70d4f3c446213b5064d7cfe95ec0ed196748f014c19a833117bac32468fd.
//
// Solidity: event DeveloperWalletUpdated(address indexed oldDeveloperWallet, address indexed newDeveloperWallet)
func (_Cmaccountmanager *CmaccountmanagerFilterer) WatchDeveloperWalletUpdated(opts *bind.WatchOpts, sink chan<- *CmaccountmanagerDeveloperWalletUpdated, oldDeveloperWallet []common.Address, newDeveloperWallet []common.Address) (event.Subscription, error) {

	var oldDeveloperWalletRule []interface{}
	for _, oldDeveloperWalletItem := range oldDeveloperWallet {
		oldDeveloperWalletRule = append(oldDeveloperWalletRule, oldDeveloperWalletItem)
	}
	var newDeveloperWalletRule []interface{}
	for _, newDeveloperWalletItem := range newDeveloperWallet {
		newDeveloperWalletRule = append(newDeveloperWalletRule, newDeveloperWalletItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.WatchLogs(opts, "DeveloperWalletUpdated", oldDeveloperWalletRule, newDeveloperWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountmanagerDeveloperWalletUpdated)
				if err := _Cmaccountmanager.contract.UnpackLog(event, "DeveloperWalletUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeveloperWalletUpdated is a log parse operation binding the contract event 0x88bf70d4f3c446213b5064d7cfe95ec0ed196748f014c19a833117bac32468fd.
//
// Solidity: event DeveloperWalletUpdated(address indexed oldDeveloperWallet, address indexed newDeveloperWallet)
func (_Cmaccountmanager *CmaccountmanagerFilterer) ParseDeveloperWalletUpdated(log types.Log) (*CmaccountmanagerDeveloperWalletUpdated, error) {
	event := new(CmaccountmanagerDeveloperWalletUpdated)
	if err := _Cmaccountmanager.contract.UnpackLog(event, "DeveloperWalletUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountmanagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Cmaccountmanager contract.
type CmaccountmanagerInitializedIterator struct {
	Event *CmaccountmanagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountmanagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountmanagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountmanagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountmanagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountmanagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountmanagerInitialized represents a Initialized event raised by the Cmaccountmanager contract.
type CmaccountmanagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Cmaccountmanager *CmaccountmanagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*CmaccountmanagerInitializedIterator, error) {

	logs, sub, err := _Cmaccountmanager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerInitializedIterator{contract: _Cmaccountmanager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Cmaccountmanager *CmaccountmanagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CmaccountmanagerInitialized) (event.Subscription, error) {

	logs, sub, err := _Cmaccountmanager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountmanagerInitialized)
				if err := _Cmaccountmanager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Cmaccountmanager *CmaccountmanagerFilterer) ParseInitialized(log types.Log) (*CmaccountmanagerInitialized, error) {
	event := new(CmaccountmanagerInitialized)
	if err := _Cmaccountmanager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountmanagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Cmaccountmanager contract.
type CmaccountmanagerPausedIterator struct {
	Event *CmaccountmanagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountmanagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountmanagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountmanagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountmanagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountmanagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountmanagerPaused represents a Paused event raised by the Cmaccountmanager contract.
type CmaccountmanagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Cmaccountmanager *CmaccountmanagerFilterer) FilterPaused(opts *bind.FilterOpts) (*CmaccountmanagerPausedIterator, error) {

	logs, sub, err := _Cmaccountmanager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerPausedIterator{contract: _Cmaccountmanager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Cmaccountmanager *CmaccountmanagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CmaccountmanagerPaused) (event.Subscription, error) {

	logs, sub, err := _Cmaccountmanager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountmanagerPaused)
				if err := _Cmaccountmanager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Cmaccountmanager *CmaccountmanagerFilterer) ParsePaused(log types.Log) (*CmaccountmanagerPaused, error) {
	event := new(CmaccountmanagerPaused)
	if err := _Cmaccountmanager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountmanagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Cmaccountmanager contract.
type CmaccountmanagerRoleAdminChangedIterator struct {
	Event *CmaccountmanagerRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountmanagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountmanagerRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountmanagerRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountmanagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountmanagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountmanagerRoleAdminChanged represents a RoleAdminChanged event raised by the Cmaccountmanager contract.
type CmaccountmanagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Cmaccountmanager *CmaccountmanagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CmaccountmanagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerRoleAdminChangedIterator{contract: _Cmaccountmanager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Cmaccountmanager *CmaccountmanagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CmaccountmanagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountmanagerRoleAdminChanged)
				if err := _Cmaccountmanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Cmaccountmanager *CmaccountmanagerFilterer) ParseRoleAdminChanged(log types.Log) (*CmaccountmanagerRoleAdminChanged, error) {
	event := new(CmaccountmanagerRoleAdminChanged)
	if err := _Cmaccountmanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountmanagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Cmaccountmanager contract.
type CmaccountmanagerRoleGrantedIterator struct {
	Event *CmaccountmanagerRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountmanagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountmanagerRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountmanagerRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountmanagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountmanagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountmanagerRoleGranted represents a RoleGranted event raised by the Cmaccountmanager contract.
type CmaccountmanagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccountmanager *CmaccountmanagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CmaccountmanagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerRoleGrantedIterator{contract: _Cmaccountmanager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccountmanager *CmaccountmanagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CmaccountmanagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountmanagerRoleGranted)
				if err := _Cmaccountmanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccountmanager *CmaccountmanagerFilterer) ParseRoleGranted(log types.Log) (*CmaccountmanagerRoleGranted, error) {
	event := new(CmaccountmanagerRoleGranted)
	if err := _Cmaccountmanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountmanagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Cmaccountmanager contract.
type CmaccountmanagerRoleRevokedIterator struct {
	Event *CmaccountmanagerRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountmanagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountmanagerRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountmanagerRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountmanagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountmanagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountmanagerRoleRevoked represents a RoleRevoked event raised by the Cmaccountmanager contract.
type CmaccountmanagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccountmanager *CmaccountmanagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CmaccountmanagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerRoleRevokedIterator{contract: _Cmaccountmanager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccountmanager *CmaccountmanagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CmaccountmanagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountmanagerRoleRevoked)
				if err := _Cmaccountmanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccountmanager *CmaccountmanagerFilterer) ParseRoleRevoked(log types.Log) (*CmaccountmanagerRoleRevoked, error) {
	event := new(CmaccountmanagerRoleRevoked)
	if err := _Cmaccountmanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountmanagerServiceRegisteredIterator is returned from FilterServiceRegistered and is used to iterate over the raw logs and unpacked data for ServiceRegistered events raised by the Cmaccountmanager contract.
type CmaccountmanagerServiceRegisteredIterator struct {
	Event *CmaccountmanagerServiceRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountmanagerServiceRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountmanagerServiceRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountmanagerServiceRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountmanagerServiceRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountmanagerServiceRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountmanagerServiceRegistered represents a ServiceRegistered event raised by the Cmaccountmanager contract.
type CmaccountmanagerServiceRegistered struct {
	ServiceName string
	ServiceHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceRegistered is a free log retrieval operation binding the contract event 0xbad43a69707ac20dd539f3163b927e83baef6e967f2c95432129b1ded4166458.
//
// Solidity: event ServiceRegistered(string serviceName, bytes32 serviceHash)
func (_Cmaccountmanager *CmaccountmanagerFilterer) FilterServiceRegistered(opts *bind.FilterOpts) (*CmaccountmanagerServiceRegisteredIterator, error) {

	logs, sub, err := _Cmaccountmanager.contract.FilterLogs(opts, "ServiceRegistered")
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerServiceRegisteredIterator{contract: _Cmaccountmanager.contract, event: "ServiceRegistered", logs: logs, sub: sub}, nil
}

// WatchServiceRegistered is a free log subscription operation binding the contract event 0xbad43a69707ac20dd539f3163b927e83baef6e967f2c95432129b1ded4166458.
//
// Solidity: event ServiceRegistered(string serviceName, bytes32 serviceHash)
func (_Cmaccountmanager *CmaccountmanagerFilterer) WatchServiceRegistered(opts *bind.WatchOpts, sink chan<- *CmaccountmanagerServiceRegistered) (event.Subscription, error) {

	logs, sub, err := _Cmaccountmanager.contract.WatchLogs(opts, "ServiceRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountmanagerServiceRegistered)
				if err := _Cmaccountmanager.contract.UnpackLog(event, "ServiceRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceRegistered is a log parse operation binding the contract event 0xbad43a69707ac20dd539f3163b927e83baef6e967f2c95432129b1ded4166458.
//
// Solidity: event ServiceRegistered(string serviceName, bytes32 serviceHash)
func (_Cmaccountmanager *CmaccountmanagerFilterer) ParseServiceRegistered(log types.Log) (*CmaccountmanagerServiceRegistered, error) {
	event := new(CmaccountmanagerServiceRegistered)
	if err := _Cmaccountmanager.contract.UnpackLog(event, "ServiceRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountmanagerServiceUnregisteredIterator is returned from FilterServiceUnregistered and is used to iterate over the raw logs and unpacked data for ServiceUnregistered events raised by the Cmaccountmanager contract.
type CmaccountmanagerServiceUnregisteredIterator struct {
	Event *CmaccountmanagerServiceUnregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountmanagerServiceUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountmanagerServiceUnregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountmanagerServiceUnregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountmanagerServiceUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountmanagerServiceUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountmanagerServiceUnregistered represents a ServiceUnregistered event raised by the Cmaccountmanager contract.
type CmaccountmanagerServiceUnregistered struct {
	ServiceName string
	ServiceHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceUnregistered is a free log retrieval operation binding the contract event 0x57158eaa7e642cefd761327d5cd6c147ddaad706ec257f90f4d8b59b3d365eb7.
//
// Solidity: event ServiceUnregistered(string serviceName, bytes32 serviceHash)
func (_Cmaccountmanager *CmaccountmanagerFilterer) FilterServiceUnregistered(opts *bind.FilterOpts) (*CmaccountmanagerServiceUnregisteredIterator, error) {

	logs, sub, err := _Cmaccountmanager.contract.FilterLogs(opts, "ServiceUnregistered")
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerServiceUnregisteredIterator{contract: _Cmaccountmanager.contract, event: "ServiceUnregistered", logs: logs, sub: sub}, nil
}

// WatchServiceUnregistered is a free log subscription operation binding the contract event 0x57158eaa7e642cefd761327d5cd6c147ddaad706ec257f90f4d8b59b3d365eb7.
//
// Solidity: event ServiceUnregistered(string serviceName, bytes32 serviceHash)
func (_Cmaccountmanager *CmaccountmanagerFilterer) WatchServiceUnregistered(opts *bind.WatchOpts, sink chan<- *CmaccountmanagerServiceUnregistered) (event.Subscription, error) {

	logs, sub, err := _Cmaccountmanager.contract.WatchLogs(opts, "ServiceUnregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountmanagerServiceUnregistered)
				if err := _Cmaccountmanager.contract.UnpackLog(event, "ServiceUnregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceUnregistered is a log parse operation binding the contract event 0x57158eaa7e642cefd761327d5cd6c147ddaad706ec257f90f4d8b59b3d365eb7.
//
// Solidity: event ServiceUnregistered(string serviceName, bytes32 serviceHash)
func (_Cmaccountmanager *CmaccountmanagerFilterer) ParseServiceUnregistered(log types.Log) (*CmaccountmanagerServiceUnregistered, error) {
	event := new(CmaccountmanagerServiceUnregistered)
	if err := _Cmaccountmanager.contract.UnpackLog(event, "ServiceUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountmanagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Cmaccountmanager contract.
type CmaccountmanagerUnpausedIterator struct {
	Event *CmaccountmanagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountmanagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountmanagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountmanagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountmanagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountmanagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountmanagerUnpaused represents a Unpaused event raised by the Cmaccountmanager contract.
type CmaccountmanagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Cmaccountmanager *CmaccountmanagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CmaccountmanagerUnpausedIterator, error) {

	logs, sub, err := _Cmaccountmanager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerUnpausedIterator{contract: _Cmaccountmanager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Cmaccountmanager *CmaccountmanagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CmaccountmanagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _Cmaccountmanager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountmanagerUnpaused)
				if err := _Cmaccountmanager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Cmaccountmanager *CmaccountmanagerFilterer) ParseUnpaused(log types.Log) (*CmaccountmanagerUnpaused, error) {
	event := new(CmaccountmanagerUnpaused)
	if err := _Cmaccountmanager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountmanagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Cmaccountmanager contract.
type CmaccountmanagerUpgradedIterator struct {
	Event *CmaccountmanagerUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountmanagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountmanagerUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountmanagerUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountmanagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountmanagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountmanagerUpgraded represents a Upgraded event raised by the Cmaccountmanager contract.
type CmaccountmanagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Cmaccountmanager *CmaccountmanagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CmaccountmanagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountmanagerUpgradedIterator{contract: _Cmaccountmanager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Cmaccountmanager *CmaccountmanagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CmaccountmanagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Cmaccountmanager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountmanagerUpgraded)
				if err := _Cmaccountmanager.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Cmaccountmanager *CmaccountmanagerFilterer) ParseUpgraded(log types.Log) (*CmaccountmanagerUpgraded, error) {
	event := new(CmaccountmanagerUpgraded)
	if err := _Cmaccountmanager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
