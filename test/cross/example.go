// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cross

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

// CrossExampleMetaData contains all meta data concerning the CrossExample contract.
var CrossExampleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"token\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"HASH\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"delta\",\"type\":\"uint64\"}],\"name\":\"minus\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"res\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"delta\",\"type\":\"uint64\"}],\"name\":\"plus\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"res\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"saveHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"res\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610b34380380610b34833981810160405281019061003291906100a6565b80600160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550506100d3565b600080fd5b600067ffffffffffffffff82169050919050565b61008381610066565b811461008e57600080fd5b50565b6000815190506100a08161007a565b92915050565b6000602082840312156100bc576100bb610061565b5b60006100ca84828501610091565b91505092915050565b610a52806100e26000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630322ed081461005c5780630a16310e1461007a5780631c01b9bf146100aa57806382bfefc8146100da5780638a8d6eb4146100f8575b600080fd5b610064610128565b60405161007191906103e7565b60405180910390f35b610094600480360381019061008f9190610453565b6101b6565b6040516100a1919061048f565b60405180910390f35b6100c460048036038101906100bf919061050f565b610221565b6040516100d191906103e7565b60405180910390f35b6100e261025f565b6040516100ef919061048f565b60405180910390f35b610112600480360381019061010d9190610453565b610279565b60405161011f919061048f565b60405180910390f35b600080546101359061058b565b80601f01602080910402602001604051908101604052809291908181526020018280546101619061058b565b80156101ae5780601f10610183576101008083540402835291602001916101ae565b820191906000526020600020905b81548152906001019060200180831161019157829003601f168201915b505050505081565b600081600160008282829054906101000a900467ffffffffffffffff166101dd91906105eb565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600160009054906101000a900467ffffffffffffffff169050919050565b6060828260009182610234929190610817565b50828260405160200161024892919061094c565b604051602081830303815290604052905092915050565b600160009054906101000a900467ffffffffffffffff1681565b6000600160009054906101000a900467ffffffffffffffff1682600081836102a19190610974565b67ffffffffffffffff1610156102ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e3906109fc565b60405180910390fd5b83600160008282829054906101000a900467ffffffffffffffff166103119190610974565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600160009054906101000a900467ffffffffffffffff1692505050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610391578082015181840152602081019050610376565b60008484015250505050565b6000601f19601f8301169050919050565b60006103b982610357565b6103c38185610362565b93506103d3818560208601610373565b6103dc8161039d565b840191505092915050565b6000602082019050818103600083015261040181846103ae565b905092915050565b600080fd5b600080fd5b600067ffffffffffffffff82169050919050565b61043081610413565b811461043b57600080fd5b50565b60008135905061044d81610427565b92915050565b60006020828403121561046957610468610409565b5b60006104778482850161043e565b91505092915050565b61048981610413565b82525050565b60006020820190506104a46000830184610480565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126104cf576104ce6104aa565b5b8235905067ffffffffffffffff8111156104ec576104eb6104af565b5b602083019150836001820283011115610508576105076104b4565b5b9250929050565b6000806020838503121561052657610525610409565b5b600083013567ffffffffffffffff8111156105445761054361040e565b5b610550858286016104b9565b92509250509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806105a357607f821691505b6020821081036105b6576105b561055c565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006105f682610413565b915061060183610413565b9250828201905067ffffffffffffffff811115610621576106206105bc565b5b92915050565b600082905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026106c37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610686565b6106cd8683610686565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600061071461070f61070a846106e5565b6106ef565b6106e5565b9050919050565b6000819050919050565b61072e836106f9565b61074261073a8261071b565b848454610693565b825550505050565b600090565b61075761074a565b610762818484610725565b505050565b5b818110156107865761077b60008261074f565b600181019050610768565b5050565b601f8211156107cb5761079c81610661565b6107a584610676565b810160208510156107b4578190505b6107c86107c085610676565b830182610767565b50505b505050565b600082821c905092915050565b60006107ee600019846008026107d0565b1980831691505092915050565b600061080783836107dd565b9150826002028217905092915050565b6108218383610627565b67ffffffffffffffff81111561083a57610839610632565b5b610844825461058b565b61084f82828561078a565b6000601f83116001811461087e576000841561086c578287013590505b61087685826107fb565b8655506108de565b601f19841661088c86610661565b60005b828110156108b45784890135825560018201915060208501945060208101905061088f565b868310156108d157848901356108cd601f8916826107dd565b8355505b6001600288020188555050505b50505050505050565b7f736176653a200000000000000000000000000000000000000000000000000000815250565b600081905092915050565b82818337600083830152505050565b6000610933838561090d565b9350610940838584610918565b82840190509392505050565b6000610957826108e7565b600682019150610968828486610927565b91508190509392505050565b600061097f82610413565b915061098a83610413565b9250828203905067ffffffffffffffff8111156109aa576109a96105bc565b5b92915050565b7f6e6f7420656e6f756768206d696e757300000000000000000000000000000000600082015250565b60006109e6601083610362565b91506109f1826109b0565b602082019050919050565b60006020820190508181036000830152610a15816109d9565b905091905056fea2646970667358221220d2b3b9b35230b48efaf1b6821c5ba4feccad795375010a7e80502bdd1c2e6a6e64736f6c63430008120033",
}

// CrossExampleABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossExampleMetaData.ABI instead.
var CrossExampleABI = CrossExampleMetaData.ABI

// CrossExampleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossExampleMetaData.Bin instead.
var CrossExampleBin = CrossExampleMetaData.Bin

// DeployCrossExample deploys a new Ethereum contract, binding an instance of CrossExample to it.
func DeployCrossExample(auth *bind.TransactOpts, backend bind.ContractBackend, token uint64) (common.Address, *types.Transaction, *CrossExample, error) {
	parsed, err := CrossExampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossExampleBin), backend, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossExample{CrossExampleCaller: CrossExampleCaller{contract: contract}, CrossExampleTransactor: CrossExampleTransactor{contract: contract}, CrossExampleFilterer: CrossExampleFilterer{contract: contract}}, nil
}

// CrossExample is an auto generated Go binding around an Ethereum contract.
type CrossExample struct {
	CrossExampleCaller     // Read-only binding to the contract
	CrossExampleTransactor // Write-only binding to the contract
	CrossExampleFilterer   // Log filterer for contract events
}

// CrossExampleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossExampleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossExampleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossExampleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossExampleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossExampleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossExampleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossExampleSession struct {
	Contract     *CrossExample     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossExampleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossExampleCallerSession struct {
	Contract *CrossExampleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CrossExampleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossExampleTransactorSession struct {
	Contract     *CrossExampleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CrossExampleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossExampleRaw struct {
	Contract *CrossExample // Generic contract binding to access the raw methods on
}

// CrossExampleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossExampleCallerRaw struct {
	Contract *CrossExampleCaller // Generic read-only contract binding to access the raw methods on
}

// CrossExampleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossExampleTransactorRaw struct {
	Contract *CrossExampleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossExample creates a new instance of CrossExample, bound to a specific deployed contract.
func NewCrossExample(address common.Address, backend bind.ContractBackend) (*CrossExample, error) {
	contract, err := bindCrossExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossExample{CrossExampleCaller: CrossExampleCaller{contract: contract}, CrossExampleTransactor: CrossExampleTransactor{contract: contract}, CrossExampleFilterer: CrossExampleFilterer{contract: contract}}, nil
}

// NewCrossExampleCaller creates a new read-only instance of CrossExample, bound to a specific deployed contract.
func NewCrossExampleCaller(address common.Address, caller bind.ContractCaller) (*CrossExampleCaller, error) {
	contract, err := bindCrossExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossExampleCaller{contract: contract}, nil
}

// NewCrossExampleTransactor creates a new write-only instance of CrossExample, bound to a specific deployed contract.
func NewCrossExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossExampleTransactor, error) {
	contract, err := bindCrossExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossExampleTransactor{contract: contract}, nil
}

// NewCrossExampleFilterer creates a new log filterer instance of CrossExample, bound to a specific deployed contract.
func NewCrossExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossExampleFilterer, error) {
	contract, err := bindCrossExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossExampleFilterer{contract: contract}, nil
}

// bindCrossExample binds a generic wrapper to an already deployed contract.
func bindCrossExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossExampleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossExample *CrossExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossExample.Contract.CrossExampleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossExample *CrossExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossExample.Contract.CrossExampleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossExample *CrossExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossExample.Contract.CrossExampleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossExample *CrossExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossExample.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossExample *CrossExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossExample.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossExample *CrossExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossExample.Contract.contract.Transact(opts, method, params...)
}

// HASH is a free data retrieval call binding the contract method 0x0322ed08.
//
// Solidity: function HASH() view returns(string)
func (_CrossExample *CrossExampleCaller) HASH(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossExample.contract.Call(opts, &out, "HASH")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// HASH is a free data retrieval call binding the contract method 0x0322ed08.
//
// Solidity: function HASH() view returns(string)
func (_CrossExample *CrossExampleSession) HASH() (string, error) {
	return _CrossExample.Contract.HASH(&_CrossExample.CallOpts)
}

// HASH is a free data retrieval call binding the contract method 0x0322ed08.
//
// Solidity: function HASH() view returns(string)
func (_CrossExample *CrossExampleCallerSession) HASH() (string, error) {
	return _CrossExample.Contract.HASH(&_CrossExample.CallOpts)
}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(uint64)
func (_CrossExample *CrossExampleCaller) TOKEN(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _CrossExample.contract.Call(opts, &out, "TOKEN")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(uint64)
func (_CrossExample *CrossExampleSession) TOKEN() (uint64, error) {
	return _CrossExample.Contract.TOKEN(&_CrossExample.CallOpts)
}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(uint64)
func (_CrossExample *CrossExampleCallerSession) TOKEN() (uint64, error) {
	return _CrossExample.Contract.TOKEN(&_CrossExample.CallOpts)
}

// Minus is a paid mutator transaction binding the contract method 0x8a8d6eb4.
//
// Solidity: function minus(uint64 delta) returns(uint64 res)
func (_CrossExample *CrossExampleTransactor) Minus(opts *bind.TransactOpts, delta uint64) (*types.Transaction, error) {
	return _CrossExample.contract.Transact(opts, "minus", delta)
}

// Minus is a paid mutator transaction binding the contract method 0x8a8d6eb4.
//
// Solidity: function minus(uint64 delta) returns(uint64 res)
func (_CrossExample *CrossExampleSession) Minus(delta uint64) (*types.Transaction, error) {
	return _CrossExample.Contract.Minus(&_CrossExample.TransactOpts, delta)
}

// Minus is a paid mutator transaction binding the contract method 0x8a8d6eb4.
//
// Solidity: function minus(uint64 delta) returns(uint64 res)
func (_CrossExample *CrossExampleTransactorSession) Minus(delta uint64) (*types.Transaction, error) {
	return _CrossExample.Contract.Minus(&_CrossExample.TransactOpts, delta)
}

// Plus is a paid mutator transaction binding the contract method 0x0a16310e.
//
// Solidity: function plus(uint64 delta) returns(uint64 res)
func (_CrossExample *CrossExampleTransactor) Plus(opts *bind.TransactOpts, delta uint64) (*types.Transaction, error) {
	return _CrossExample.contract.Transact(opts, "plus", delta)
}

// Plus is a paid mutator transaction binding the contract method 0x0a16310e.
//
// Solidity: function plus(uint64 delta) returns(uint64 res)
func (_CrossExample *CrossExampleSession) Plus(delta uint64) (*types.Transaction, error) {
	return _CrossExample.Contract.Plus(&_CrossExample.TransactOpts, delta)
}

// Plus is a paid mutator transaction binding the contract method 0x0a16310e.
//
// Solidity: function plus(uint64 delta) returns(uint64 res)
func (_CrossExample *CrossExampleTransactorSession) Plus(delta uint64) (*types.Transaction, error) {
	return _CrossExample.Contract.Plus(&_CrossExample.TransactOpts, delta)
}

// SaveHash is a paid mutator transaction binding the contract method 0x1c01b9bf.
//
// Solidity: function saveHash(string hash) returns(string res)
func (_CrossExample *CrossExampleTransactor) SaveHash(opts *bind.TransactOpts, hash string) (*types.Transaction, error) {
	return _CrossExample.contract.Transact(opts, "saveHash", hash)
}

// SaveHash is a paid mutator transaction binding the contract method 0x1c01b9bf.
//
// Solidity: function saveHash(string hash) returns(string res)
func (_CrossExample *CrossExampleSession) SaveHash(hash string) (*types.Transaction, error) {
	return _CrossExample.Contract.SaveHash(&_CrossExample.TransactOpts, hash)
}

// SaveHash is a paid mutator transaction binding the contract method 0x1c01b9bf.
//
// Solidity: function saveHash(string hash) returns(string res)
func (_CrossExample *CrossExampleTransactorSession) SaveHash(hash string) (*types.Transaction, error) {
	return _CrossExample.Contract.SaveHash(&_CrossExample.TransactOpts, hash)
}
