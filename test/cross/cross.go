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

// CrossTransactionMetaData contains all meta data concerning the CrossTransaction contract.
var CrossTransactionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chainId_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"crossId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"payload\",\"type\":\"string\"}],\"name\":\"CrossStart\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"crossId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"response\",\"type\":\"string\"}],\"name\":\"SetCrossStatusResponse\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"res\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"crossId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"payload\",\"type\":\"string\"}],\"name\":\"StartCross\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"res\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"crossId2Payoad\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"crossId2Response\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"crossId2Status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200139a3803806200139a8339818101604052810190620000379190620001e3565b80600090816200004891906200047f565b505062000566565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620000b9826200006e565b810181811067ffffffffffffffff82111715620000db57620000da6200007f565b5b80604052505050565b6000620000f062000050565b9050620000fe8282620000ae565b919050565b600067ffffffffffffffff8211156200012157620001206200007f565b5b6200012c826200006e565b9050602081019050919050565b60005b83811015620001595780820151818401526020810190506200013c565b60008484015250505050565b60006200017c620001768462000103565b620000e4565b9050828152602081018484840111156200019b576200019a62000069565b5b620001a884828562000139565b509392505050565b600082601f830112620001c857620001c762000064565b5b8151620001da84826020860162000165565b91505092915050565b600060208284031215620001fc57620001fb6200005a565b5b600082015167ffffffffffffffff8111156200021d576200021c6200005f565b5b6200022b84828501620001b0565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200028757607f821691505b6020821081036200029d576200029c6200023f565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003077fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620002c8565b620003138683620002c8565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003606200035a62000354846200032b565b62000335565b6200032b565b9050919050565b6000819050919050565b6200037c836200033f565b620003946200038b8262000367565b848454620002d5565b825550505050565b600090565b620003ab6200039c565b620003b881848462000371565b505050565b5b81811015620003e057620003d4600082620003a1565b600181019050620003be565b5050565b601f8211156200042f57620003f981620002a3565b6200040484620002b8565b8101602085101562000414578190505b6200042c6200042385620002b8565b830182620003bd565b50505b505050565b600082821c905092915050565b6000620004546000198460080262000434565b1980831691505092915050565b60006200046f838362000441565b9150826002028217905092915050565b6200048a8262000234565b67ffffffffffffffff811115620004a657620004a56200007f565b5b620004b282546200026e565b620004bf828285620003e4565b600060209050601f831160018114620004f75760008415620004e2578287015190505b620004ee858262000461565b8655506200055e565b601f1984166200050786620002a3565b60005b8281101562000531578489015182556001820191506020850194506020810190506200050a565b868310156200055157848901516200054d601f89168262000441565b8355505b6001600288020188555050505b505050505050565b610e2480620005766000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630a4dbf7d146100675780631d0060ea1461009757806362cb1fd4146100c757806391670a82146100f85780639a8a059214610129578063b69cfa3e14610147575b600080fd5b610081600480360381019061007c91906106d7565b610177565b60405161008e919061079f565b60405180910390f35b6100b160048036038101906100ac91906106d7565b61022d565b6040516100be919061079f565b60405180910390f35b6100e160048036038101906100dc9190610821565b6102e3565b6040516100ef9291906108cf565b60405180910390f35b610112600480360381019061010d91906108f3565b6103c5565b6040516101209291906108cf565b60405180910390f35b610131610439565b60405161013e919061079f565b60405180910390f35b610161600480360381019061015c91906106d7565b6104c7565b60405161016e919061079f565b60405180910390f35b60018180516020810182018051848252602083016020850120818352809550505050505060009150905080546101ac906109d6565b80601f01602080910402602001604051908101604052809291908181526020018280546101d8906109d6565b80156102255780601f106101fa57610100808354040283529160200191610225565b820191906000526020600020905b81548152906001019060200180831161020857829003601f168201915b505050505081565b6003818051602081018201805184825260208301602085012081835280955050505050506000915090508054610262906109d6565b80601f016020809104026020016040519081016040528092919081815260200182805461028e906109d6565b80156102db5780601f106102b0576101008083540402835291602001916102db565b820191906000526020600020905b8154815290600101906020018083116102be57829003601f168201915b505050505081565b3660008383600188886040516102fa929190610a37565b90815260200160405180910390209182610315929190610c11565b506040518060400160405280600581526020017f73746172740000000000000000000000000000000000000000000000000000008152506002878760405161035e929190610a37565b908152602001604051809103902090816103789190610ce1565b507fe828aac2d3ca53282ba314b9d2a259f177bc1a8b751e292d899d294fbe5230a3868686866040516103ae9493929190610db3565b60405180910390a185859150915094509492505050565b366000858560028a8a6040516103dc929190610a37565b908152602001604051809103902091826103f7929190610c11565b50838360038a8a60405161040c929190610a37565b90815260200160405180910390209182610427929190610c11565b50878791509150965096945050505050565b60008054610446906109d6565b80601f0160208091040260200160405190810160405280929190818152602001828054610472906109d6565b80156104bf5780601f10610494576101008083540402835291602001916104bf565b820191906000526020600020905b8154815290600101906020018083116104a257829003601f168201915b505050505081565b60028180516020810182018051848252602083016020850120818352809550505050505060009150905080546104fc906109d6565b80601f0160208091040260200160405190810160405280929190818152602001828054610528906109d6565b80156105755780601f1061054a57610100808354040283529160200191610575565b820191906000526020600020905b81548152906001019060200180831161055857829003601f168201915b505050505081565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6105e48261059b565b810181811067ffffffffffffffff82111715610603576106026105ac565b5b80604052505050565b600061061661057d565b905061062282826105db565b919050565b600067ffffffffffffffff821115610642576106416105ac565b5b61064b8261059b565b9050602081019050919050565b82818337600083830152505050565b600061067a61067584610627565b61060c565b90508281526020810184848401111561069657610695610596565b5b6106a1848285610658565b509392505050565b600082601f8301126106be576106bd610591565b5b81356106ce848260208601610667565b91505092915050565b6000602082840312156106ed576106ec610587565b5b600082013567ffffffffffffffff81111561070b5761070a61058c565b5b610717848285016106a9565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561075a57808201518184015260208101905061073f565b60008484015250505050565b600061077182610720565b61077b818561072b565b935061078b81856020860161073c565b6107948161059b565b840191505092915050565b600060208201905081810360008301526107b98184610766565b905092915050565b600080fd5b600080fd5b60008083601f8401126107e1576107e0610591565b5b8235905067ffffffffffffffff8111156107fe576107fd6107c1565b5b60208301915083600182028301111561081a576108196107c6565b5b9250929050565b6000806000806040858703121561083b5761083a610587565b5b600085013567ffffffffffffffff8111156108595761085861058c565b5b610865878288016107cb565b9450945050602085013567ffffffffffffffff8111156108885761088761058c565b5b610894878288016107cb565b925092505092959194509250565b60006108ae838561072b565b93506108bb838584610658565b6108c48361059b565b840190509392505050565b600060208201905081810360008301526108ea8184866108a2565b90509392505050565b600080600080600080606087890312156109105761090f610587565b5b600087013567ffffffffffffffff81111561092e5761092d61058c565b5b61093a89828a016107cb565b9650965050602087013567ffffffffffffffff81111561095d5761095c61058c565b5b61096989828a016107cb565b9450945050604087013567ffffffffffffffff81111561098c5761098b61058c565b5b61099889828a016107cb565b92509250509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806109ee57607f821691505b602082108103610a0157610a006109a7565b5b50919050565b600081905092915050565b6000610a1e8385610a07565b9350610a2b838584610658565b82840190509392505050565b6000610a44828486610a12565b91508190509392505050565b600082905092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610abd7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a80565b610ac78683610a80565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000610b0e610b09610b0484610adf565b610ae9565b610adf565b9050919050565b6000819050919050565b610b2883610af3565b610b3c610b3482610b15565b848454610a8d565b825550505050565b600090565b610b51610b44565b610b5c818484610b1f565b505050565b5b81811015610b8057610b75600082610b49565b600181019050610b62565b5050565b601f821115610bc557610b9681610a5b565b610b9f84610a70565b81016020851015610bae578190505b610bc2610bba85610a70565b830182610b61565b50505b505050565b600082821c905092915050565b6000610be860001984600802610bca565b1980831691505092915050565b6000610c018383610bd7565b9150826002028217905092915050565b610c1b8383610a50565b67ffffffffffffffff811115610c3457610c336105ac565b5b610c3e82546109d6565b610c49828285610b84565b6000601f831160018114610c785760008415610c66578287013590505b610c708582610bf5565b865550610cd8565b601f198416610c8686610a5b565b60005b82811015610cae57848901358255600182019150602085019450602081019050610c89565b86831015610ccb5784890135610cc7601f891682610bd7565b8355505b6001600288020188555050505b50505050505050565b610cea82610720565b67ffffffffffffffff811115610d0357610d026105ac565b5b610d0d82546109d6565b610d18828285610b84565b600060209050601f831160018114610d4b5760008415610d39578287015190505b610d438582610bf5565b865550610dab565b601f198416610d5986610a5b565b60005b82811015610d8157848901518255600182019150602085019450602081019050610d5c565b86831015610d9e5784890151610d9a601f891682610bd7565b8355505b6001600288020188555050505b505050505050565b60006040820190508181036000830152610dce8186886108a2565b90508181036020830152610de38184866108a2565b90509594505050505056fea2646970667358221220fe2cffc2a014517e0d1b24e832a713cd8e6d746f4a7e95f8957d7ab19f35503164736f6c63430008120033",
}

// CrossTransactionABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossTransactionMetaData.ABI instead.
var CrossTransactionABI = CrossTransactionMetaData.ABI

// CrossTransactionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossTransactionMetaData.Bin instead.
var CrossTransactionBin = CrossTransactionMetaData.Bin

// DeployCrossTransaction deploys a new Ethereum contract, binding an instance of CrossTransaction to it.
func DeployCrossTransaction(auth *bind.TransactOpts, backend bind.ContractBackend, chainId_ string) (common.Address, *types.Transaction, *CrossTransaction, error) {
	parsed, err := CrossTransactionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossTransactionBin), backend, chainId_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossTransaction{CrossTransactionCaller: CrossTransactionCaller{contract: contract}, CrossTransactionTransactor: CrossTransactionTransactor{contract: contract}, CrossTransactionFilterer: CrossTransactionFilterer{contract: contract}}, nil
}

// CrossTransaction is an auto generated Go binding around an Ethereum contract.
type CrossTransaction struct {
	CrossTransactionCaller     // Read-only binding to the contract
	CrossTransactionTransactor // Write-only binding to the contract
	CrossTransactionFilterer   // Log filterer for contract events
}

// CrossTransactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossTransactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossTransactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossTransactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossTransactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossTransactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossTransactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossTransactionSession struct {
	Contract     *CrossTransaction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossTransactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossTransactionCallerSession struct {
	Contract *CrossTransactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CrossTransactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossTransactionTransactorSession struct {
	Contract     *CrossTransactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CrossTransactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossTransactionRaw struct {
	Contract *CrossTransaction // Generic contract binding to access the raw methods on
}

// CrossTransactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossTransactionCallerRaw struct {
	Contract *CrossTransactionCaller // Generic read-only contract binding to access the raw methods on
}

// CrossTransactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossTransactionTransactorRaw struct {
	Contract *CrossTransactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossTransaction creates a new instance of CrossTransaction, bound to a specific deployed contract.
func NewCrossTransaction(address common.Address, backend bind.ContractBackend) (*CrossTransaction, error) {
	contract, err := bindCrossTransaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossTransaction{CrossTransactionCaller: CrossTransactionCaller{contract: contract}, CrossTransactionTransactor: CrossTransactionTransactor{contract: contract}, CrossTransactionFilterer: CrossTransactionFilterer{contract: contract}}, nil
}

// NewCrossTransactionCaller creates a new read-only instance of CrossTransaction, bound to a specific deployed contract.
func NewCrossTransactionCaller(address common.Address, caller bind.ContractCaller) (*CrossTransactionCaller, error) {
	contract, err := bindCrossTransaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossTransactionCaller{contract: contract}, nil
}

// NewCrossTransactionTransactor creates a new write-only instance of CrossTransaction, bound to a specific deployed contract.
func NewCrossTransactionTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossTransactionTransactor, error) {
	contract, err := bindCrossTransaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossTransactionTransactor{contract: contract}, nil
}

// NewCrossTransactionFilterer creates a new log filterer instance of CrossTransaction, bound to a specific deployed contract.
func NewCrossTransactionFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossTransactionFilterer, error) {
	contract, err := bindCrossTransaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossTransactionFilterer{contract: contract}, nil
}

// bindCrossTransaction binds a generic wrapper to an already deployed contract.
func bindCrossTransaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossTransactionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossTransaction *CrossTransactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossTransaction.Contract.CrossTransactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossTransaction *CrossTransactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossTransaction.Contract.CrossTransactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossTransaction *CrossTransactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossTransaction.Contract.CrossTransactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossTransaction *CrossTransactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossTransaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossTransaction *CrossTransactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossTransaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossTransaction *CrossTransactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossTransaction.Contract.contract.Transact(opts, method, params...)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(string)
func (_CrossTransaction *CrossTransactionCaller) ChainId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossTransaction.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(string)
func (_CrossTransaction *CrossTransactionSession) ChainId() (string, error) {
	return _CrossTransaction.Contract.ChainId(&_CrossTransaction.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(string)
func (_CrossTransaction *CrossTransactionCallerSession) ChainId() (string, error) {
	return _CrossTransaction.Contract.ChainId(&_CrossTransaction.CallOpts)
}

// CrossId2Payoad is a free data retrieval call binding the contract method 0x0a4dbf7d.
//
// Solidity: function crossId2Payoad(string ) view returns(string)
func (_CrossTransaction *CrossTransactionCaller) CrossId2Payoad(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _CrossTransaction.contract.Call(opts, &out, "crossId2Payoad", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CrossId2Payoad is a free data retrieval call binding the contract method 0x0a4dbf7d.
//
// Solidity: function crossId2Payoad(string ) view returns(string)
func (_CrossTransaction *CrossTransactionSession) CrossId2Payoad(arg0 string) (string, error) {
	return _CrossTransaction.Contract.CrossId2Payoad(&_CrossTransaction.CallOpts, arg0)
}

// CrossId2Payoad is a free data retrieval call binding the contract method 0x0a4dbf7d.
//
// Solidity: function crossId2Payoad(string ) view returns(string)
func (_CrossTransaction *CrossTransactionCallerSession) CrossId2Payoad(arg0 string) (string, error) {
	return _CrossTransaction.Contract.CrossId2Payoad(&_CrossTransaction.CallOpts, arg0)
}

// CrossId2Response is a free data retrieval call binding the contract method 0x1d0060ea.
//
// Solidity: function crossId2Response(string ) view returns(string)
func (_CrossTransaction *CrossTransactionCaller) CrossId2Response(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _CrossTransaction.contract.Call(opts, &out, "crossId2Response", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CrossId2Response is a free data retrieval call binding the contract method 0x1d0060ea.
//
// Solidity: function crossId2Response(string ) view returns(string)
func (_CrossTransaction *CrossTransactionSession) CrossId2Response(arg0 string) (string, error) {
	return _CrossTransaction.Contract.CrossId2Response(&_CrossTransaction.CallOpts, arg0)
}

// CrossId2Response is a free data retrieval call binding the contract method 0x1d0060ea.
//
// Solidity: function crossId2Response(string ) view returns(string)
func (_CrossTransaction *CrossTransactionCallerSession) CrossId2Response(arg0 string) (string, error) {
	return _CrossTransaction.Contract.CrossId2Response(&_CrossTransaction.CallOpts, arg0)
}

// CrossId2Status is a free data retrieval call binding the contract method 0xb69cfa3e.
//
// Solidity: function crossId2Status(string ) view returns(string)
func (_CrossTransaction *CrossTransactionCaller) CrossId2Status(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _CrossTransaction.contract.Call(opts, &out, "crossId2Status", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CrossId2Status is a free data retrieval call binding the contract method 0xb69cfa3e.
//
// Solidity: function crossId2Status(string ) view returns(string)
func (_CrossTransaction *CrossTransactionSession) CrossId2Status(arg0 string) (string, error) {
	return _CrossTransaction.Contract.CrossId2Status(&_CrossTransaction.CallOpts, arg0)
}

// CrossId2Status is a free data retrieval call binding the contract method 0xb69cfa3e.
//
// Solidity: function crossId2Status(string ) view returns(string)
func (_CrossTransaction *CrossTransactionCallerSession) CrossId2Status(arg0 string) (string, error) {
	return _CrossTransaction.Contract.CrossId2Status(&_CrossTransaction.CallOpts, arg0)
}

// SetCrossStatusResponse is a paid mutator transaction binding the contract method 0x91670a82.
//
// Solidity: function SetCrossStatusResponse(string crossId, string status, string response) returns(string res)
func (_CrossTransaction *CrossTransactionTransactor) SetCrossStatusResponse(opts *bind.TransactOpts, crossId string, status string, response string) (*types.Transaction, error) {
	return _CrossTransaction.contract.Transact(opts, "SetCrossStatusResponse", crossId, status, response)
}

// SetCrossStatusResponse is a paid mutator transaction binding the contract method 0x91670a82.
//
// Solidity: function SetCrossStatusResponse(string crossId, string status, string response) returns(string res)
func (_CrossTransaction *CrossTransactionSession) SetCrossStatusResponse(crossId string, status string, response string) (*types.Transaction, error) {
	return _CrossTransaction.Contract.SetCrossStatusResponse(&_CrossTransaction.TransactOpts, crossId, status, response)
}

// SetCrossStatusResponse is a paid mutator transaction binding the contract method 0x91670a82.
//
// Solidity: function SetCrossStatusResponse(string crossId, string status, string response) returns(string res)
func (_CrossTransaction *CrossTransactionTransactorSession) SetCrossStatusResponse(crossId string, status string, response string) (*types.Transaction, error) {
	return _CrossTransaction.Contract.SetCrossStatusResponse(&_CrossTransaction.TransactOpts, crossId, status, response)
}

// StartCross is a paid mutator transaction binding the contract method 0x62cb1fd4.
//
// Solidity: function StartCross(string crossId, string payload) returns(string res)
func (_CrossTransaction *CrossTransactionTransactor) StartCross(opts *bind.TransactOpts, crossId string, payload string) (*types.Transaction, error) {
	return _CrossTransaction.contract.Transact(opts, "StartCross", crossId, payload)
}

// StartCross is a paid mutator transaction binding the contract method 0x62cb1fd4.
//
// Solidity: function StartCross(string crossId, string payload) returns(string res)
func (_CrossTransaction *CrossTransactionSession) StartCross(crossId string, payload string) (*types.Transaction, error) {
	return _CrossTransaction.Contract.StartCross(&_CrossTransaction.TransactOpts, crossId, payload)
}

// StartCross is a paid mutator transaction binding the contract method 0x62cb1fd4.
//
// Solidity: function StartCross(string crossId, string payload) returns(string res)
func (_CrossTransaction *CrossTransactionTransactorSession) StartCross(crossId string, payload string) (*types.Transaction, error) {
	return _CrossTransaction.Contract.StartCross(&_CrossTransaction.TransactOpts, crossId, payload)
}

// CrossTransactionCrossStartIterator is returned from FilterCrossStart and is used to iterate over the raw logs and unpacked data for CrossStart events raised by the CrossTransaction contract.
type CrossTransactionCrossStartIterator struct {
	Event *CrossTransactionCrossStart // Event containing the contract specifics and raw log

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
func (it *CrossTransactionCrossStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossTransactionCrossStart)
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
		it.Event = new(CrossTransactionCrossStart)
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
func (it *CrossTransactionCrossStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossTransactionCrossStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossTransactionCrossStart represents a CrossStart event raised by the CrossTransaction contract.
type CrossTransactionCrossStart struct {
	CrossId string
	Payload string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossStart is a free log retrieval operation binding the contract event 0xe828aac2d3ca53282ba314b9d2a259f177bc1a8b751e292d899d294fbe5230a3.
//
// Solidity: event CrossStart(string crossId, string payload)
func (_CrossTransaction *CrossTransactionFilterer) FilterCrossStart(opts *bind.FilterOpts) (*CrossTransactionCrossStartIterator, error) {

	logs, sub, err := _CrossTransaction.contract.FilterLogs(opts, "CrossStart")
	if err != nil {
		return nil, err
	}
	return &CrossTransactionCrossStartIterator{contract: _CrossTransaction.contract, event: "CrossStart", logs: logs, sub: sub}, nil
}

// WatchCrossStart is a free log subscription operation binding the contract event 0xe828aac2d3ca53282ba314b9d2a259f177bc1a8b751e292d899d294fbe5230a3.
//
// Solidity: event CrossStart(string crossId, string payload)
func (_CrossTransaction *CrossTransactionFilterer) WatchCrossStart(opts *bind.WatchOpts, sink chan<- *CrossTransactionCrossStart) (event.Subscription, error) {

	logs, sub, err := _CrossTransaction.contract.WatchLogs(opts, "CrossStart")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossTransactionCrossStart)
				if err := _CrossTransaction.contract.UnpackLog(event, "CrossStart", log); err != nil {
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

// ParseCrossStart is a log parse operation binding the contract event 0xe828aac2d3ca53282ba314b9d2a259f177bc1a8b751e292d899d294fbe5230a3.
//
// Solidity: event CrossStart(string crossId, string payload)
func (_CrossTransaction *CrossTransactionFilterer) ParseCrossStart(log types.Log) (*CrossTransactionCrossStart, error) {
	event := new(CrossTransactionCrossStart)
	if err := _CrossTransaction.contract.UnpackLog(event, "CrossStart", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
