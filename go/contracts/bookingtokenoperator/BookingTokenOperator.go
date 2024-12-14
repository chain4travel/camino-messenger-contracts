// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bookingtokenoperator

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

// BookingtokenoperatorMetaData contains all meta data concerning the Bookingtokenoperator contract.
var BookingtokenoperatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenApprovalFailed\",\"type\":\"error\"}]",
	Bin: "0x61056e61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c80631b29e1ba1461004557806353ea54c514610067575b600080fd5b81801561005157600080fd5b50610065610060366004610312565b610087565b005b81801561007357600080fd5b50610065610082366004610364565b61028c565b6040516213f74f60e21b81526004810182905260009081906001600160a01b03851690624fdd3c906024016040805180830381865afa1580156100ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100f2919061045c565b90925090506001600160a01b0381161580159061010f5750600082115b156102295760405163095ea7b360e01b81526001600160a01b038581166004830152602482018490526000919083169063095ea7b3906044016020604051808303816000875af1158015610167573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061018b919061048c565b9050806101c957604051633337e73160e01b81526001600160a01b038087166004830152831660248201526044810184905260640160405180910390fd5b6040516396591edd60e01b8152600481018590526001600160a01b038616906396591edd90602401600060405180830381600087803b15801561020b57600080fd5b505af115801561021f573d6000803e3d6000fd5b5050505050610286565b6040516396591edd60e01b8152600481018490526001600160a01b038516906396591edd9084906024016000604051808303818588803b15801561026c57600080fd5b505af1158015610280573d6000803e3d6000fd5b50505050505b50505050565b604051632ea5d6d960e11b81526001600160a01b03871690635d4badb2906102c090889088908890889088906004016104b5565b600060405180830381600087803b1580156102da57600080fd5b505af11580156102ee573d6000803e3d6000fd5b50505050505050505050565b6001600160a01b038116811461030f57600080fd5b50565b6000806040838503121561032557600080fd5b8235610330816102fa565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b803561035f816102fa565b919050565b60008060008060008060c0878903121561037d57600080fd5b8635610388816102fa565b95506020870135610398816102fa565b945060408701356001600160401b03808211156103b457600080fd5b818901915089601f8301126103c857600080fd5b8135818111156103da576103da61033e565b604051601f8201601f19908116603f011681019083821181831017156104025761040261033e565b816040528281528c602084870101111561041b57600080fd5b826020860160208301376000602084830101528098505050505050606087013592506080870135915061045060a08801610354565b90509295509295509295565b6000806040838503121561046f57600080fd5b825191506020830151610481816102fa565b809150509250929050565b60006020828403121561049e57600080fd5b815180151581146104ae57600080fd5b9392505050565b60018060a01b03861681526000602060a0602084015286518060a085015260005b818110156104f25788810183015185820160c0015282016104d6565b50600060c0828601015260c0601f19601f8301168501019250505084604083015283606083015261052e60808301846001600160a01b03169052565b969550505050505056fea26469706673582212204657a5a9ac752819f1d460d5e76e1f84fdb4107e950a89489ba00b8ef64c7e1664736f6c63430008180033",
}

// BookingtokenoperatorABI is the input ABI used to generate the binding from.
// Deprecated: Use BookingtokenoperatorMetaData.ABI instead.
var BookingtokenoperatorABI = BookingtokenoperatorMetaData.ABI

// BookingtokenoperatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BookingtokenoperatorMetaData.Bin instead.
var BookingtokenoperatorBin = BookingtokenoperatorMetaData.Bin

// DeployBookingtokenoperator deploys a new Ethereum contract, binding an instance of Bookingtokenoperator to it.
func DeployBookingtokenoperator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bookingtokenoperator, error) {
	parsed, err := BookingtokenoperatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BookingtokenoperatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bookingtokenoperator{BookingtokenoperatorCaller: BookingtokenoperatorCaller{contract: contract}, BookingtokenoperatorTransactor: BookingtokenoperatorTransactor{contract: contract}, BookingtokenoperatorFilterer: BookingtokenoperatorFilterer{contract: contract}}, nil
}

// Bookingtokenoperator is an auto generated Go binding around an Ethereum contract.
type Bookingtokenoperator struct {
	BookingtokenoperatorCaller     // Read-only binding to the contract
	BookingtokenoperatorTransactor // Write-only binding to the contract
	BookingtokenoperatorFilterer   // Log filterer for contract events
}

// BookingtokenoperatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BookingtokenoperatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookingtokenoperatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BookingtokenoperatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookingtokenoperatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BookingtokenoperatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookingtokenoperatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BookingtokenoperatorSession struct {
	Contract     *Bookingtokenoperator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BookingtokenoperatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BookingtokenoperatorCallerSession struct {
	Contract *BookingtokenoperatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BookingtokenoperatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BookingtokenoperatorTransactorSession struct {
	Contract     *BookingtokenoperatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BookingtokenoperatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BookingtokenoperatorRaw struct {
	Contract *Bookingtokenoperator // Generic contract binding to access the raw methods on
}

// BookingtokenoperatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BookingtokenoperatorCallerRaw struct {
	Contract *BookingtokenoperatorCaller // Generic read-only contract binding to access the raw methods on
}

// BookingtokenoperatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BookingtokenoperatorTransactorRaw struct {
	Contract *BookingtokenoperatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBookingtokenoperator creates a new instance of Bookingtokenoperator, bound to a specific deployed contract.
func NewBookingtokenoperator(address common.Address, backend bind.ContractBackend) (*Bookingtokenoperator, error) {
	contract, err := bindBookingtokenoperator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenoperator{BookingtokenoperatorCaller: BookingtokenoperatorCaller{contract: contract}, BookingtokenoperatorTransactor: BookingtokenoperatorTransactor{contract: contract}, BookingtokenoperatorFilterer: BookingtokenoperatorFilterer{contract: contract}}, nil
}

// NewBookingtokenoperatorCaller creates a new read-only instance of Bookingtokenoperator, bound to a specific deployed contract.
func NewBookingtokenoperatorCaller(address common.Address, caller bind.ContractCaller) (*BookingtokenoperatorCaller, error) {
	contract, err := bindBookingtokenoperator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BookingtokenoperatorCaller{contract: contract}, nil
}

// NewBookingtokenoperatorTransactor creates a new write-only instance of Bookingtokenoperator, bound to a specific deployed contract.
func NewBookingtokenoperatorTransactor(address common.Address, transactor bind.ContractTransactor) (*BookingtokenoperatorTransactor, error) {
	contract, err := bindBookingtokenoperator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BookingtokenoperatorTransactor{contract: contract}, nil
}

// NewBookingtokenoperatorFilterer creates a new log filterer instance of Bookingtokenoperator, bound to a specific deployed contract.
func NewBookingtokenoperatorFilterer(address common.Address, filterer bind.ContractFilterer) (*BookingtokenoperatorFilterer, error) {
	contract, err := bindBookingtokenoperator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BookingtokenoperatorFilterer{contract: contract}, nil
}

// bindBookingtokenoperator binds a generic wrapper to an already deployed contract.
func bindBookingtokenoperator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BookingtokenoperatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bookingtokenoperator *BookingtokenoperatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bookingtokenoperator.Contract.BookingtokenoperatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bookingtokenoperator *BookingtokenoperatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookingtokenoperator.Contract.BookingtokenoperatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bookingtokenoperator *BookingtokenoperatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bookingtokenoperator.Contract.BookingtokenoperatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bookingtokenoperator *BookingtokenoperatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bookingtokenoperator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bookingtokenoperator *BookingtokenoperatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookingtokenoperator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bookingtokenoperator *BookingtokenoperatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bookingtokenoperator.Contract.contract.Transact(opts, method, params...)
}
