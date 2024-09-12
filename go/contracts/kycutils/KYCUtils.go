// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kycutils

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

// KycutilsMetaData contains all meta data concerning the Kycutils contract.
var KycutilsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NotKYBVerified\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NotKYCVerified\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NotVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ADMIN_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KYB_VERIFIED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KYC_EXPIRED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KYC_VERIFIED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// KycutilsABI is the input ABI used to generate the binding from.
// Deprecated: Use KycutilsMetaData.ABI instead.
var KycutilsABI = KycutilsMetaData.ABI

// Kycutils is an auto generated Go binding around an Ethereum contract.
type Kycutils struct {
	KycutilsCaller     // Read-only binding to the contract
	KycutilsTransactor // Write-only binding to the contract
	KycutilsFilterer   // Log filterer for contract events
}

// KycutilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type KycutilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KycutilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KycutilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KycutilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KycutilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KycutilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KycutilsSession struct {
	Contract     *Kycutils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KycutilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KycutilsCallerSession struct {
	Contract *KycutilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// KycutilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KycutilsTransactorSession struct {
	Contract     *KycutilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KycutilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type KycutilsRaw struct {
	Contract *Kycutils // Generic contract binding to access the raw methods on
}

// KycutilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KycutilsCallerRaw struct {
	Contract *KycutilsCaller // Generic read-only contract binding to access the raw methods on
}

// KycutilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KycutilsTransactorRaw struct {
	Contract *KycutilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKycutils creates a new instance of Kycutils, bound to a specific deployed contract.
func NewKycutils(address common.Address, backend bind.ContractBackend) (*Kycutils, error) {
	contract, err := bindKycutils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kycutils{KycutilsCaller: KycutilsCaller{contract: contract}, KycutilsTransactor: KycutilsTransactor{contract: contract}, KycutilsFilterer: KycutilsFilterer{contract: contract}}, nil
}

// NewKycutilsCaller creates a new read-only instance of Kycutils, bound to a specific deployed contract.
func NewKycutilsCaller(address common.Address, caller bind.ContractCaller) (*KycutilsCaller, error) {
	contract, err := bindKycutils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KycutilsCaller{contract: contract}, nil
}

// NewKycutilsTransactor creates a new write-only instance of Kycutils, bound to a specific deployed contract.
func NewKycutilsTransactor(address common.Address, transactor bind.ContractTransactor) (*KycutilsTransactor, error) {
	contract, err := bindKycutils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KycutilsTransactor{contract: contract}, nil
}

// NewKycutilsFilterer creates a new log filterer instance of Kycutils, bound to a specific deployed contract.
func NewKycutilsFilterer(address common.Address, filterer bind.ContractFilterer) (*KycutilsFilterer, error) {
	contract, err := bindKycutils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KycutilsFilterer{contract: contract}, nil
}

// bindKycutils binds a generic wrapper to an already deployed contract.
func bindKycutils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KycutilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kycutils *KycutilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kycutils.Contract.KycutilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kycutils *KycutilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kycutils.Contract.KycutilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kycutils *KycutilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kycutils.Contract.KycutilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kycutils *KycutilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kycutils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kycutils *KycutilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kycutils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kycutils *KycutilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kycutils.Contract.contract.Transact(opts, method, params...)
}

// ADMINADDR is a free data retrieval call binding the contract method 0xb11569f5.
//
// Solidity: function ADMIN_ADDR() view returns(address)
func (_Kycutils *KycutilsCaller) ADMINADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kycutils.contract.Call(opts, &out, "ADMIN_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADMINADDR is a free data retrieval call binding the contract method 0xb11569f5.
//
// Solidity: function ADMIN_ADDR() view returns(address)
func (_Kycutils *KycutilsSession) ADMINADDR() (common.Address, error) {
	return _Kycutils.Contract.ADMINADDR(&_Kycutils.CallOpts)
}

// ADMINADDR is a free data retrieval call binding the contract method 0xb11569f5.
//
// Solidity: function ADMIN_ADDR() view returns(address)
func (_Kycutils *KycutilsCallerSession) ADMINADDR() (common.Address, error) {
	return _Kycutils.Contract.ADMINADDR(&_Kycutils.CallOpts)
}

// KYBVERIFIED is a free data retrieval call binding the contract method 0xa58a021e.
//
// Solidity: function KYB_VERIFIED() view returns(uint256)
func (_Kycutils *KycutilsCaller) KYBVERIFIED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kycutils.contract.Call(opts, &out, "KYB_VERIFIED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KYBVERIFIED is a free data retrieval call binding the contract method 0xa58a021e.
//
// Solidity: function KYB_VERIFIED() view returns(uint256)
func (_Kycutils *KycutilsSession) KYBVERIFIED() (*big.Int, error) {
	return _Kycutils.Contract.KYBVERIFIED(&_Kycutils.CallOpts)
}

// KYBVERIFIED is a free data retrieval call binding the contract method 0xa58a021e.
//
// Solidity: function KYB_VERIFIED() view returns(uint256)
func (_Kycutils *KycutilsCallerSession) KYBVERIFIED() (*big.Int, error) {
	return _Kycutils.Contract.KYBVERIFIED(&_Kycutils.CallOpts)
}

// KYCEXPIRED is a free data retrieval call binding the contract method 0x97041685.
//
// Solidity: function KYC_EXPIRED() view returns(uint256)
func (_Kycutils *KycutilsCaller) KYCEXPIRED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kycutils.contract.Call(opts, &out, "KYC_EXPIRED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KYCEXPIRED is a free data retrieval call binding the contract method 0x97041685.
//
// Solidity: function KYC_EXPIRED() view returns(uint256)
func (_Kycutils *KycutilsSession) KYCEXPIRED() (*big.Int, error) {
	return _Kycutils.Contract.KYCEXPIRED(&_Kycutils.CallOpts)
}

// KYCEXPIRED is a free data retrieval call binding the contract method 0x97041685.
//
// Solidity: function KYC_EXPIRED() view returns(uint256)
func (_Kycutils *KycutilsCallerSession) KYCEXPIRED() (*big.Int, error) {
	return _Kycutils.Contract.KYCEXPIRED(&_Kycutils.CallOpts)
}

// KYCVERIFIED is a free data retrieval call binding the contract method 0x517103bf.
//
// Solidity: function KYC_VERIFIED() view returns(uint256)
func (_Kycutils *KycutilsCaller) KYCVERIFIED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kycutils.contract.Call(opts, &out, "KYC_VERIFIED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KYCVERIFIED is a free data retrieval call binding the contract method 0x517103bf.
//
// Solidity: function KYC_VERIFIED() view returns(uint256)
func (_Kycutils *KycutilsSession) KYCVERIFIED() (*big.Int, error) {
	return _Kycutils.Contract.KYCVERIFIED(&_Kycutils.CallOpts)
}

// KYCVERIFIED is a free data retrieval call binding the contract method 0x517103bf.
//
// Solidity: function KYC_VERIFIED() view returns(uint256)
func (_Kycutils *KycutilsCallerSession) KYCVERIFIED() (*big.Int, error) {
	return _Kycutils.Contract.KYCVERIFIED(&_Kycutils.CallOpts)
}
