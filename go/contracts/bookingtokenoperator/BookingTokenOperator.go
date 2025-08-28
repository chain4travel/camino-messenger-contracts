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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"actualPaymentToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"expectedPaymentToken\",\"type\":\"address\"}],\"name\":\"UnexpectedPaymentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedPrice\",\"type\":\"uint256\"}],\"name\":\"UnexpectedPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NATIVE_PAYMENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OFFCHAIN_PAYMENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610e1d61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a35760003560e01c806307e47316146100a857806321b87f3a146100ca578063348e06dd146100ea578063793dddac1461010a5780637adf63b71461012a578063a0f07c741461014a578063b54e72d814610168578063bfb26c0614610188578063c7bffa9614610190578063e4c22569146101b0578063fd13a43e146101d0575b600080fd5b8180156100b457600080fd5b506100c86100c33660046109e7565b6101f0565b005b8180156100d657600080fd5b506100c86100e5366004610a36565b610257565b8180156100f657600080fd5b506100c8610105366004610a36565b6103f9565b81801561011657600080fd5b506100c86101253660046109e7565b61045e565b81801561013657600080fd5b506100c8610145366004610a76565b61048e565b610152600081565b60405161015f9190610ac0565b60405180910390f35b81801561017457600080fd5b506100c8610183366004610ad4565b6106d6565b610152600181565b81801561019c57600080fd5b506100c86101ab366004610b2d565b610741565b8180156101bc57600080fd5b506100c86101cb366004610b7f565b610797565b8180156101dc57600080fd5b506100c86101eb366004610ad4565b61080b565b6040516254232760e71b81526001600160a01b03851690632a1193809061021f90869086908690600401610c93565b600060405180830381600087803b15801561023957600080fd5b505af115801561024d573d6000803e3d6000fd5b5050505050505050565b6040516358c8e84960e11b8152600481018390526000906001600160a01b0385169063b191d09290602401602060405180830381865afa15801561029f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c39190610cad565b90506001600160a01b03811661033a57604051631c54f0f760e01b81526001600160a01b03851690631c54f0f79084906103039087908390600401610cd1565b6000604051808303818588803b15801561031c57600080fd5b505af1158015610330573d6000803e3d6000fd5b50505050506103f3565b6000196001600160a01b038216016103b157604051631c54f0f760e01b81526001600160a01b03851690631c54f0f79061037a9086908690600401610cd1565b600060405180830381600087803b15801561039457600080fd5b505af11580156103a8573d6000803e3d6000fd5b505050506103f3565b6103c56001600160a01b038216858461083d565b604051631c54f0f760e01b81526001600160a01b03851690631c54f0f79061021f9086908690600401610cd1565b50505050565b6040516317ccce3160e31b81526001600160a01b0384169063be667188906104279085908590600401610cd1565b600060405180830381600087803b15801561044157600080fd5b505af1158015610455573d6000803e3d6000fd5b50505050505050565b6040516374fe60e960e01b81526001600160a01b038516906374fe60e99061021f90869086908690600401610c93565b6040516213f74f60e21b81526004810184905260009081906001600160a01b03871690624fdd3c906024016040805180830381865afa1580156104d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f99190610cdf565b91509150838214610533576040516330cb9c4160e01b81526004810186905260248101839052604481018590526064015b60405180910390fd5b826001600160a01b0316816001600160a01b03161461057f5760405163107f801560e21b8152600481018690526001600160a01b0380831660248301528416604482015260640161052a565b6001600160a01b0381166105ee576040516396591edd60e01b8152600481018690526001600160a01b038716906396591edd9084906024016000604051808303818588803b1580156105d057600080fd5b505af11580156105e4573d6000803e3d6000fd5b50505050506106ce565b6000196001600160a01b0382160161065f576040516396591edd60e01b8152600481018690526001600160a01b038716906396591edd90602401600060405180830381600087803b15801561064257600080fd5b505af1158015610656573d6000803e3d6000fd5b505050506106ce565b6106736001600160a01b038216878461083d565b6040516396591edd60e01b8152600481018690526001600160a01b038716906396591edd90602401600060405180830381600087803b1580156106b557600080fd5b505af11580156106c9573d6000803e3d6000fd5b505050505b505050505050565b60405163f7e45f0960e01b81526001600160a01b0386169063f7e45f0990610708908790879087908790600401610d0f565b600060405180830381600087803b15801561072257600080fd5b505af1158015610736573d6000803e3d6000fd5b505050505050505050565b6040516339699c9760e21b8152600481018290526001600160a01b0383169063e5a6725c90602401600060405180830381600087803b15801561078357600080fd5b505af11580156106ce573d6000803e3d6000fd5b604051636d95934160e11b81526001600160a01b0389169063db2b2682906107cf908a908a908a908a908a908a908a90600401610d31565b600060405180830381600087803b1580156107e957600080fd5b505af11580156107fd573d6000803e3d6000fd5b505050505050505050505050565b604051630e95440960e31b81526001600160a01b038616906374aa204890610708908790879087908790600401610d0f565b6000836001600160a01b031663095ea7b38484604051602401610861929190610dce565b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050905061089a8482610901565b6103f3576108f784856001600160a01b031663095ea7b38660006040516024016108c5929190610dce565b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050610950565b6103f38482610950565b6000806000806020600086516020880160008a5af192503d91506000519050828015610946575081156109375780600114610946565b6000866001600160a01b03163b115b9695505050505050565b600080602060008451602086016000885af180610973576040513d6000823e3d81fd5b50506000513d9150811561098b578060011415610998565b6001600160a01b0384163b155b156103f35783604051635274afe760e01b815260040161052a9190610ac0565b6001600160a01b03811681146109cd57600080fd5b50565b803561ffff811681146109e257600080fd5b919050565b600080600080608085870312156109fd57600080fd5b8435610a08816109b8565b935060208501359250610a1d604086016109d0565b9150610a2b606086016109d0565b905092959194509250565b600080600060608486031215610a4b57600080fd5b8335610a56816109b8565b95602085013595506040909401359392505050565b80356109e2816109b8565b60008060008060808587031215610a8c57600080fd5b8435610a97816109b8565b935060208501359250604085013591506060850135610ab5816109b8565b939692955090935050565b6001600160a01b0391909116815260200190565b600080600080600060a08688031215610aec57600080fd5b8535610af7816109b8565b94506020860135935060408601359250610b13606087016109d0565b9150610b21608087016109d0565b90509295509295909350565b60008060408385031215610b4057600080fd5b8235610b4b816109b8565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b803580151581146109e257600080fd5b600080600080600080600080610100898b031215610b9c57600080fd5b8835610ba7816109b8565b97506020890135610bb7816109b8565b965060408901356001600160401b0380821115610bd357600080fd5b818b0191508b601f830112610be757600080fd5b813581811115610bf957610bf9610b59565b604051601f8201601f19908116603f01168101908382118183101715610c2157610c21610b59565b816040528281528e6020848701011115610c3a57600080fd5b82602086016020830137600060208483010152809a5050505050506060890135945060808901359350610c6f60a08a01610a6b565b925060c08901359150610c8460e08a01610b6f565b90509295985092959890939650565b92835261ffff918216602084015216604082015260600190565b600060208284031215610cbf57600080fd5b8151610cca816109b8565b9392505050565b918252602082015260400190565b60008060408385031215610cf257600080fd5b825191506020830151610d04816109b8565b809150509250929050565b938452602084019290925261ffff908116604084015216606082015260800190565b60018060a01b03881681526000602060e0602084015288518060e085015260005b81811015610d6f578a810183015185820161010001528201610d52565b506101009150600082828601015281601f19601f83011685010192505050866040830152856060830152610dae60808301866001600160a01b03169052565b8360a0830152610dc260c083018415159052565b98975050505050505050565b6001600160a01b0392909216825260208201526040019056fea2646970667358221220474533ef269b61616783f6dda315cbfe3084edd3179ebd40357be4bcebc980ae64736f6c63430008180033",
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

// NATIVEPAYMENT is a free data retrieval call binding the contract method 0xa0f07c74.
//
// Solidity: function NATIVE_PAYMENT() view returns(address)
func (_Bookingtokenoperator *BookingtokenoperatorCaller) NATIVEPAYMENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bookingtokenoperator.contract.Call(opts, &out, "NATIVE_PAYMENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEPAYMENT is a free data retrieval call binding the contract method 0xa0f07c74.
//
// Solidity: function NATIVE_PAYMENT() view returns(address)
func (_Bookingtokenoperator *BookingtokenoperatorSession) NATIVEPAYMENT() (common.Address, error) {
	return _Bookingtokenoperator.Contract.NATIVEPAYMENT(&_Bookingtokenoperator.CallOpts)
}

// NATIVEPAYMENT is a free data retrieval call binding the contract method 0xa0f07c74.
//
// Solidity: function NATIVE_PAYMENT() view returns(address)
func (_Bookingtokenoperator *BookingtokenoperatorCallerSession) NATIVEPAYMENT() (common.Address, error) {
	return _Bookingtokenoperator.Contract.NATIVEPAYMENT(&_Bookingtokenoperator.CallOpts)
}

// OFFCHAINPAYMENT is a free data retrieval call binding the contract method 0xbfb26c06.
//
// Solidity: function OFFCHAIN_PAYMENT() view returns(address)
func (_Bookingtokenoperator *BookingtokenoperatorCaller) OFFCHAINPAYMENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bookingtokenoperator.contract.Call(opts, &out, "OFFCHAIN_PAYMENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OFFCHAINPAYMENT is a free data retrieval call binding the contract method 0xbfb26c06.
//
// Solidity: function OFFCHAIN_PAYMENT() view returns(address)
func (_Bookingtokenoperator *BookingtokenoperatorSession) OFFCHAINPAYMENT() (common.Address, error) {
	return _Bookingtokenoperator.Contract.OFFCHAINPAYMENT(&_Bookingtokenoperator.CallOpts)
}

// OFFCHAINPAYMENT is a free data retrieval call binding the contract method 0xbfb26c06.
//
// Solidity: function OFFCHAIN_PAYMENT() view returns(address)
func (_Bookingtokenoperator *BookingtokenoperatorCallerSession) OFFCHAINPAYMENT() (common.Address, error) {
	return _Bookingtokenoperator.Contract.OFFCHAINPAYMENT(&_Bookingtokenoperator.CallOpts)
}
