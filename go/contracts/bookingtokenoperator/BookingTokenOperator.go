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
)

// BookingtokenoperatorMetaData contains all meta data concerning the Bookingtokenoperator contract.
var BookingtokenoperatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenApprovalFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NATIVE_PAYMENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OFFCHAIN_PAYMENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610e4561003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ae5760003560e01c806317a1da30146100b35780631b29e1ba146100d55780633212361c146100f557806333e238fe146101155780633f825ef8146101355780637767d51614610155578063a0f07c7414610175578063b0160c6814610199578063bfb26c06146101b9578063c7bffa96146101c1578063e4c22569146101e1578063e68b6eaa14610201575b600080fd5b8180156100bf57600080fd5b506100d36100ce3660046109d3565b610221565b005b8180156100e157600080fd5b506100d36100f0366004610a08565b610286565b81801561010157600080fd5b506100d3610110366004610a4b565b6104dc565b81801561012157600080fd5b506100d3610130366004610a4b565b610544565b81801561014157600080fd5b506100d3610150366004610a9a565b610574565b81801561016157600080fd5b506100d36101703660046109d3565b6105ba565b61017d600081565b6040516001600160a01b03909116815260200160405180910390f35b8180156101a557600080fd5b506100d36101b43660046109d3565b610886565b61017d600181565b8180156101cd57600080fd5b506100d36101dc366004610a08565b6108b4565b8180156101ed57600080fd5b506100d36101fc366004610b2d565b610912565b81801561020d57600080fd5b506100d361021c366004610c41565b610986565b60405163078126f960e51b81526001600160a01b0384169063f024df209061024f9085908590600401610c83565b600060405180830381600087803b15801561026957600080fd5b505af115801561027d573d6000803e3d6000fd5b50505050505050565b6040516213f74f60e21b81526004810182905260009081906001600160a01b03851690624fdd3c906024016040805180830381865afa1580156102cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102f19190610c91565b90925090506001600160a01b038116610365576040516396591edd60e01b8152600481018490526001600160a01b038516906396591edd9084906024016000604051808303818588803b15801561034757600080fd5b505af115801561035b573d6000803e3d6000fd5b50505050506104d6565b6000196001600160a01b038216016103d6576040516396591edd60e01b8152600481018490526001600160a01b038516906396591edd90602401600060405180830381600087803b1580156103b957600080fd5b505af11580156103cd573d6000803e3d6000fd5b505050506104d6565b60405163095ea7b360e01b81526000906001600160a01b0383169063095ea7b3906104079088908790600401610cc1565b6020604051808303816000875af1158015610426573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061044a9190610cda565b90508061047957848284604051633337e73160e01b815260040161047093929190610cfe565b60405180910390fd5b6040516396591edd60e01b8152600481018590526001600160a01b038616906396591edd906024015b600060405180830381600087803b1580156104bc57600080fd5b505af11580156104d0573d6000803e3d6000fd5b50505050505b50505050565b604051636fabeaed60e01b81526001600160a01b03851690636fabeaed9061050c90869086908690600401610d22565b600060405180830381600087803b15801561052657600080fd5b505af115801561053a573d6000803e3d6000fd5b5050505050505050565b6040516334395bf360e11b81526001600160a01b03851690636872b7e69061050c90869086908690600401610d22565b60405163b47177a160e01b8152600481018590526024810184905261ffff8084166044830152821660648201526001600160a01b0386169063b47177a1906084016104a2565b6040516358c8e84960e11b8152600481018390526000906001600160a01b0385169063b191d09290602401602060405180830381865afa158015610602573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106269190610d3c565b604051635080e33f60e11b8152600481018590529091506000906001600160a01b0386169063a101c67e90602401602060405180830381865afa158015610671573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106959190610d59565b90506001600160a01b03821661070c5760405163de62fe4d60e01b81526001600160a01b0386169063de62fe4d9083906106d59088908890600401610c83565b6000604051808303818588803b1580156106ee57600080fd5b505af1158015610702573d6000803e3d6000fd5b505050505061087f565b6000196001600160a01b038316016107835760405163de62fe4d60e01b81526001600160a01b0386169063de62fe4d9061074c9087908790600401610c83565b600060405180830381600087803b15801561076657600080fd5b505af115801561077a573d6000803e3d6000fd5b5050505061087f565b60405163095ea7b360e01b81526000906001600160a01b0384169063095ea7b3906107b49089908690600401610cc1565b6020604051808303816000875af11580156107d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107f79190610cda565b90508061081d57858383604051633337e73160e01b815260040161047093929190610cfe565b60405163de62fe4d60e01b81526001600160a01b0387169063de62fe4d9061084b9088908890600401610c83565b600060405180830381600087803b15801561086557600080fd5b505af1158015610879573d6000803e3d6000fd5b50505050505b5050505050565b60405163017012c560e51b81526001600160a01b03841690632e0258a09061024f9085908590600401610c83565b6040516339699c9760e21b8152600481018290526001600160a01b0383169063e5a6725c90602401600060405180830381600087803b1580156108f657600080fd5b505af115801561090a573d6000803e3d6000fd5b505050505050565b604051636d95934160e11b81526001600160a01b0389169063db2b26829061094a908a908a908a908a908a908a908a90600401610d72565b600060405180830381600087803b15801561096457600080fd5b505af1158015610978573d6000803e3d6000fd5b505050505050505050505050565b6040516304a0b8c560e41b81526004810183905281151560248201526001600160a01b03841690634a0b8c509060440161024f565b6001600160a01b03811681146109d057600080fd5b50565b6000806000606084860312156109e857600080fd5b83356109f3816109bb565b95602085013595506040909401359392505050565b60008060408385031215610a1b57600080fd5b8235610a26816109bb565b946020939093013593505050565b803561ffff81168114610a4657600080fd5b919050565b60008060008060808587031215610a6157600080fd5b8435610a6c816109bb565b935060208501359250610a8160408601610a34565b9150610a8f60608601610a34565b905092959194509250565b600080600080600060a08688031215610ab257600080fd5b8535610abd816109bb565b94506020860135935060408601359250610ad960608701610a34565b9150610ae760808701610a34565b90509295509295909350565b634e487b7160e01b600052604160045260246000fd5b8035610a46816109bb565b80151581146109d057600080fd5b8035610a4681610b14565b600080600080600080600080610100898b031215610b4a57600080fd5b8835610b55816109bb565b97506020890135610b65816109bb565b965060408901356001600160401b0380821115610b8157600080fd5b818b0191508b601f830112610b9557600080fd5b813581811115610ba757610ba7610af3565b604051601f8201601f19908116603f01168101908382118183101715610bcf57610bcf610af3565b816040528281528e6020848701011115610be857600080fd5b82602086016020830137600060208483010152809a5050505050506060890135945060808901359350610c1d60a08a01610b09565b925060c08901359150610c3260e08a01610b22565b90509295985092959890939650565b600080600060608486031215610c5657600080fd5b8335610c61816109bb565b9250602084013591506040840135610c7881610b14565b809150509250925092565b918252602082015260400190565b60008060408385031215610ca457600080fd5b825191506020830151610cb6816109bb565b809150509250929050565b6001600160a01b03929092168252602082015260400190565b600060208284031215610cec57600080fd5b8151610cf781610b14565b9392505050565b6001600160a01b039384168152919092166020820152604081019190915260600190565b92835261ffff918216602084015216604082015260600190565b600060208284031215610d4e57600080fd5b8151610cf7816109bb565b600060208284031215610d6b57600080fd5b5051919050565b60018060a01b03881681526000602060e0602084015288518060e085015260005b81811015610db0578a810183015185820161010001528201610d93565b506101009150600082828601015281601f19601f83011685010192505050866040830152856060830152610def60808301866001600160a01b03169052565b8360a0830152610e0360c083018415159052565b9897505050505050505056fea2646970667358221220e7e812a045171a7119fc299175ccd50f71a64e7b2db4f75289566162bb1f280664736f6c63430008180033",
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
	parsed, err := abi.JSON(strings.NewReader(BookingtokenoperatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
