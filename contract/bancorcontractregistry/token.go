// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bancorcontractregistry

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

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_CONVERTER_UPGRADER\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BNT_TOKEN\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_contractName\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REGISTRY\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractName\",\"type\":\"bytes32\"}],\"name\":\"unregisterAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractNames\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_CONVERTER_FACTORY\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BNT_CONVERTER\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractName\",\"type\":\"bytes32\"},{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"registerAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"itemCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_FORMULA\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEATURES\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_NETWORK\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_GAS_PRICE_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_contractName\",\"type\":\"bytes32\"}],\"name\":\"addressOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_X\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_contractName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"AddressUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// BANCORCONVERTERFACTORY is a free data retrieval call binding the contract method 0x5a46f06c.
//
// Solidity: function BANCOR_CONVERTER_FACTORY() view returns(bytes32)
func (_Token *TokenCaller) BANCORCONVERTERFACTORY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "BANCOR_CONVERTER_FACTORY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANCORCONVERTERFACTORY is a free data retrieval call binding the contract method 0x5a46f06c.
//
// Solidity: function BANCOR_CONVERTER_FACTORY() view returns(bytes32)
func (_Token *TokenSession) BANCORCONVERTERFACTORY() ([32]byte, error) {
	return _Token.Contract.BANCORCONVERTERFACTORY(&_Token.CallOpts)
}

// BANCORCONVERTERFACTORY is a free data retrieval call binding the contract method 0x5a46f06c.
//
// Solidity: function BANCOR_CONVERTER_FACTORY() view returns(bytes32)
func (_Token *TokenCallerSession) BANCORCONVERTERFACTORY() ([32]byte, error) {
	return _Token.Contract.BANCORCONVERTERFACTORY(&_Token.CallOpts)
}

// BANCORCONVERTERUPGRADER is a free data retrieval call binding the contract method 0x0c87355e.
//
// Solidity: function BANCOR_CONVERTER_UPGRADER() view returns(bytes32)
func (_Token *TokenCaller) BANCORCONVERTERUPGRADER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "BANCOR_CONVERTER_UPGRADER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANCORCONVERTERUPGRADER is a free data retrieval call binding the contract method 0x0c87355e.
//
// Solidity: function BANCOR_CONVERTER_UPGRADER() view returns(bytes32)
func (_Token *TokenSession) BANCORCONVERTERUPGRADER() ([32]byte, error) {
	return _Token.Contract.BANCORCONVERTERUPGRADER(&_Token.CallOpts)
}

// BANCORCONVERTERUPGRADER is a free data retrieval call binding the contract method 0x0c87355e.
//
// Solidity: function BANCOR_CONVERTER_UPGRADER() view returns(bytes32)
func (_Token *TokenCallerSession) BANCORCONVERTERUPGRADER() ([32]byte, error) {
	return _Token.Contract.BANCORCONVERTERUPGRADER(&_Token.CallOpts)
}

// BANCORFORMULA is a free data retrieval call binding the contract method 0x6d7bd3fc.
//
// Solidity: function BANCOR_FORMULA() view returns(bytes32)
func (_Token *TokenCaller) BANCORFORMULA(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "BANCOR_FORMULA")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANCORFORMULA is a free data retrieval call binding the contract method 0x6d7bd3fc.
//
// Solidity: function BANCOR_FORMULA() view returns(bytes32)
func (_Token *TokenSession) BANCORFORMULA() ([32]byte, error) {
	return _Token.Contract.BANCORFORMULA(&_Token.CallOpts)
}

// BANCORFORMULA is a free data retrieval call binding the contract method 0x6d7bd3fc.
//
// Solidity: function BANCOR_FORMULA() view returns(bytes32)
func (_Token *TokenCallerSession) BANCORFORMULA() ([32]byte, error) {
	return _Token.Contract.BANCORFORMULA(&_Token.CallOpts)
}

// BANCORGASPRICELIMIT is a free data retrieval call binding the contract method 0x9249993a.
//
// Solidity: function BANCOR_GAS_PRICE_LIMIT() view returns(bytes32)
func (_Token *TokenCaller) BANCORGASPRICELIMIT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "BANCOR_GAS_PRICE_LIMIT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANCORGASPRICELIMIT is a free data retrieval call binding the contract method 0x9249993a.
//
// Solidity: function BANCOR_GAS_PRICE_LIMIT() view returns(bytes32)
func (_Token *TokenSession) BANCORGASPRICELIMIT() ([32]byte, error) {
	return _Token.Contract.BANCORGASPRICELIMIT(&_Token.CallOpts)
}

// BANCORGASPRICELIMIT is a free data retrieval call binding the contract method 0x9249993a.
//
// Solidity: function BANCOR_GAS_PRICE_LIMIT() view returns(bytes32)
func (_Token *TokenCallerSession) BANCORGASPRICELIMIT() ([32]byte, error) {
	return _Token.Contract.BANCORGASPRICELIMIT(&_Token.CallOpts)
}

// BANCORNETWORK is a free data retrieval call binding the contract method 0x9232494e.
//
// Solidity: function BANCOR_NETWORK() view returns(bytes32)
func (_Token *TokenCaller) BANCORNETWORK(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "BANCOR_NETWORK")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANCORNETWORK is a free data retrieval call binding the contract method 0x9232494e.
//
// Solidity: function BANCOR_NETWORK() view returns(bytes32)
func (_Token *TokenSession) BANCORNETWORK() ([32]byte, error) {
	return _Token.Contract.BANCORNETWORK(&_Token.CallOpts)
}

// BANCORNETWORK is a free data retrieval call binding the contract method 0x9232494e.
//
// Solidity: function BANCOR_NETWORK() view returns(bytes32)
func (_Token *TokenCallerSession) BANCORNETWORK() ([32]byte, error) {
	return _Token.Contract.BANCORNETWORK(&_Token.CallOpts)
}

// BANCORX is a free data retrieval call binding the contract method 0xc4a8598e.
//
// Solidity: function BANCOR_X() view returns(bytes32)
func (_Token *TokenCaller) BANCORX(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "BANCOR_X")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANCORX is a free data retrieval call binding the contract method 0xc4a8598e.
//
// Solidity: function BANCOR_X() view returns(bytes32)
func (_Token *TokenSession) BANCORX() ([32]byte, error) {
	return _Token.Contract.BANCORX(&_Token.CallOpts)
}

// BANCORX is a free data retrieval call binding the contract method 0xc4a8598e.
//
// Solidity: function BANCOR_X() view returns(bytes32)
func (_Token *TokenCallerSession) BANCORX() ([32]byte, error) {
	return _Token.Contract.BANCORX(&_Token.CallOpts)
}

// BNTCONVERTER is a free data retrieval call binding the contract method 0x62614ae6.
//
// Solidity: function BNT_CONVERTER() view returns(bytes32)
func (_Token *TokenCaller) BNTCONVERTER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "BNT_CONVERTER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BNTCONVERTER is a free data retrieval call binding the contract method 0x62614ae6.
//
// Solidity: function BNT_CONVERTER() view returns(bytes32)
func (_Token *TokenSession) BNTCONVERTER() ([32]byte, error) {
	return _Token.Contract.BNTCONVERTER(&_Token.CallOpts)
}

// BNTCONVERTER is a free data retrieval call binding the contract method 0x62614ae6.
//
// Solidity: function BNT_CONVERTER() view returns(bytes32)
func (_Token *TokenCallerSession) BNTCONVERTER() ([32]byte, error) {
	return _Token.Contract.BNTCONVERTER(&_Token.CallOpts)
}

// BNTTOKEN is a free data retrieval call binding the contract method 0x1d000b61.
//
// Solidity: function BNT_TOKEN() view returns(bytes32)
func (_Token *TokenCaller) BNTTOKEN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "BNT_TOKEN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BNTTOKEN is a free data retrieval call binding the contract method 0x1d000b61.
//
// Solidity: function BNT_TOKEN() view returns(bytes32)
func (_Token *TokenSession) BNTTOKEN() ([32]byte, error) {
	return _Token.Contract.BNTTOKEN(&_Token.CallOpts)
}

// BNTTOKEN is a free data retrieval call binding the contract method 0x1d000b61.
//
// Solidity: function BNT_TOKEN() view returns(bytes32)
func (_Token *TokenCallerSession) BNTTOKEN() ([32]byte, error) {
	return _Token.Contract.BNTTOKEN(&_Token.CallOpts)
}

// CONTRACTFEATURES is a free data retrieval call binding the contract method 0x83315b6e.
//
// Solidity: function CONTRACT_FEATURES() view returns(bytes32)
func (_Token *TokenCaller) CONTRACTFEATURES(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "CONTRACT_FEATURES")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTRACTFEATURES is a free data retrieval call binding the contract method 0x83315b6e.
//
// Solidity: function CONTRACT_FEATURES() view returns(bytes32)
func (_Token *TokenSession) CONTRACTFEATURES() ([32]byte, error) {
	return _Token.Contract.CONTRACTFEATURES(&_Token.CallOpts)
}

// CONTRACTFEATURES is a free data retrieval call binding the contract method 0x83315b6e.
//
// Solidity: function CONTRACT_FEATURES() view returns(bytes32)
func (_Token *TokenCallerSession) CONTRACTFEATURES() ([32]byte, error) {
	return _Token.Contract.CONTRACTFEATURES(&_Token.CallOpts)
}

// CONTRACTREGISTRY is a free data retrieval call binding the contract method 0x25f9bfef.
//
// Solidity: function CONTRACT_REGISTRY() view returns(bytes32)
func (_Token *TokenCaller) CONTRACTREGISTRY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "CONTRACT_REGISTRY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTRACTREGISTRY is a free data retrieval call binding the contract method 0x25f9bfef.
//
// Solidity: function CONTRACT_REGISTRY() view returns(bytes32)
func (_Token *TokenSession) CONTRACTREGISTRY() ([32]byte, error) {
	return _Token.Contract.CONTRACTREGISTRY(&_Token.CallOpts)
}

// CONTRACTREGISTRY is a free data retrieval call binding the contract method 0x25f9bfef.
//
// Solidity: function CONTRACT_REGISTRY() view returns(bytes32)
func (_Token *TokenCallerSession) CONTRACTREGISTRY() ([32]byte, error) {
	return _Token.Contract.CONTRACTREGISTRY(&_Token.CallOpts)
}

// AddressOf is a free data retrieval call binding the contract method 0xbb34534c.
//
// Solidity: function addressOf(bytes32 _contractName) view returns(address)
func (_Token *TokenCaller) AddressOf(opts *bind.CallOpts, _contractName [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "addressOf", _contractName)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressOf is a free data retrieval call binding the contract method 0xbb34534c.
//
// Solidity: function addressOf(bytes32 _contractName) view returns(address)
func (_Token *TokenSession) AddressOf(_contractName [32]byte) (common.Address, error) {
	return _Token.Contract.AddressOf(&_Token.CallOpts, _contractName)
}

// AddressOf is a free data retrieval call binding the contract method 0xbb34534c.
//
// Solidity: function addressOf(bytes32 _contractName) view returns(address)
func (_Token *TokenCallerSession) AddressOf(_contractName [32]byte) (common.Address, error) {
	return _Token.Contract.AddressOf(&_Token.CallOpts, _contractName)
}

// ContractNames is a free data retrieval call binding the contract method 0x3ca6bb92.
//
// Solidity: function contractNames(uint256 ) view returns(string)
func (_Token *TokenCaller) ContractNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "contractNames", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractNames is a free data retrieval call binding the contract method 0x3ca6bb92.
//
// Solidity: function contractNames(uint256 ) view returns(string)
func (_Token *TokenSession) ContractNames(arg0 *big.Int) (string, error) {
	return _Token.Contract.ContractNames(&_Token.CallOpts, arg0)
}

// ContractNames is a free data retrieval call binding the contract method 0x3ca6bb92.
//
// Solidity: function contractNames(uint256 ) view returns(string)
func (_Token *TokenCallerSession) ContractNames(arg0 *big.Int) (string, error) {
	return _Token.Contract.ContractNames(&_Token.CallOpts, arg0)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _contractName) view returns(address)
func (_Token *TokenCaller) GetAddress(opts *bind.CallOpts, _contractName [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getAddress", _contractName)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _contractName) view returns(address)
func (_Token *TokenSession) GetAddress(_contractName [32]byte) (common.Address, error) {
	return _Token.Contract.GetAddress(&_Token.CallOpts, _contractName)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _contractName) view returns(address)
func (_Token *TokenCallerSession) GetAddress(_contractName [32]byte) (common.Address, error) {
	return _Token.Contract.GetAddress(&_Token.CallOpts, _contractName)
}

// ItemCount is a free data retrieval call binding the contract method 0x6bfb0d01.
//
// Solidity: function itemCount() view returns(uint256)
func (_Token *TokenCaller) ItemCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "itemCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ItemCount is a free data retrieval call binding the contract method 0x6bfb0d01.
//
// Solidity: function itemCount() view returns(uint256)
func (_Token *TokenSession) ItemCount() (*big.Int, error) {
	return _Token.Contract.ItemCount(&_Token.CallOpts)
}

// ItemCount is a free data retrieval call binding the contract method 0x6bfb0d01.
//
// Solidity: function itemCount() view returns(uint256)
func (_Token *TokenCallerSession) ItemCount() (*big.Int, error) {
	return _Token.Contract.ItemCount(&_Token.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Token *TokenCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "newOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Token *TokenSession) NewOwner() (common.Address, error) {
	return _Token.Contract.NewOwner(&_Token.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Token *TokenCallerSession) NewOwner() (common.Address, error) {
	return _Token.Contract.NewOwner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Token *TokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Token *TokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _Token.Contract.AcceptOwnership(&_Token.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Token *TokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Token.Contract.AcceptOwnership(&_Token.TransactOpts)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0x662de379.
//
// Solidity: function registerAddress(bytes32 _contractName, address _contractAddress) returns()
func (_Token *TokenTransactor) RegisterAddress(opts *bind.TransactOpts, _contractName [32]byte, _contractAddress common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "registerAddress", _contractName, _contractAddress)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0x662de379.
//
// Solidity: function registerAddress(bytes32 _contractName, address _contractAddress) returns()
func (_Token *TokenSession) RegisterAddress(_contractName [32]byte, _contractAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.RegisterAddress(&_Token.TransactOpts, _contractName, _contractAddress)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0x662de379.
//
// Solidity: function registerAddress(bytes32 _contractName, address _contractAddress) returns()
func (_Token *TokenTransactorSession) RegisterAddress(_contractName [32]byte, _contractAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.RegisterAddress(&_Token.TransactOpts, _contractName, _contractAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Token *TokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Token *TokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Token *TokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, _newOwner)
}

// UnregisterAddress is a paid mutator transaction binding the contract method 0x2bbd9530.
//
// Solidity: function unregisterAddress(bytes32 _contractName) returns()
func (_Token *TokenTransactor) UnregisterAddress(opts *bind.TransactOpts, _contractName [32]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "unregisterAddress", _contractName)
}

// UnregisterAddress is a paid mutator transaction binding the contract method 0x2bbd9530.
//
// Solidity: function unregisterAddress(bytes32 _contractName) returns()
func (_Token *TokenSession) UnregisterAddress(_contractName [32]byte) (*types.Transaction, error) {
	return _Token.Contract.UnregisterAddress(&_Token.TransactOpts, _contractName)
}

// UnregisterAddress is a paid mutator transaction binding the contract method 0x2bbd9530.
//
// Solidity: function unregisterAddress(bytes32 _contractName) returns()
func (_Token *TokenTransactorSession) UnregisterAddress(_contractName [32]byte) (*types.Transaction, error) {
	return _Token.Contract.UnregisterAddress(&_Token.TransactOpts, _contractName)
}

// TokenAddressUpdateIterator is returned from FilterAddressUpdate and is used to iterate over the raw logs and unpacked data for AddressUpdate events raised by the Token contract.
type TokenAddressUpdateIterator struct {
	Event *TokenAddressUpdate // Event containing the contract specifics and raw log

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
func (it *TokenAddressUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAddressUpdate)
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
		it.Event = new(TokenAddressUpdate)
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
func (it *TokenAddressUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAddressUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAddressUpdate represents a AddressUpdate event raised by the Token contract.
type TokenAddressUpdate struct {
	ContractName    [32]byte
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddressUpdate is a free log retrieval operation binding the contract event 0xfc08d1253c81bcd5444fc7056ef1f5a5df4c9220b6fd70d7449267f1f0f29918.
//
// Solidity: event AddressUpdate(bytes32 indexed _contractName, address _contractAddress)
func (_Token *TokenFilterer) FilterAddressUpdate(opts *bind.FilterOpts, _contractName [][32]byte) (*TokenAddressUpdateIterator, error) {

	var _contractNameRule []interface{}
	for _, _contractNameItem := range _contractName {
		_contractNameRule = append(_contractNameRule, _contractNameItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "AddressUpdate", _contractNameRule)
	if err != nil {
		return nil, err
	}
	return &TokenAddressUpdateIterator{contract: _Token.contract, event: "AddressUpdate", logs: logs, sub: sub}, nil
}

// WatchAddressUpdate is a free log subscription operation binding the contract event 0xfc08d1253c81bcd5444fc7056ef1f5a5df4c9220b6fd70d7449267f1f0f29918.
//
// Solidity: event AddressUpdate(bytes32 indexed _contractName, address _contractAddress)
func (_Token *TokenFilterer) WatchAddressUpdate(opts *bind.WatchOpts, sink chan<- *TokenAddressUpdate, _contractName [][32]byte) (event.Subscription, error) {

	var _contractNameRule []interface{}
	for _, _contractNameItem := range _contractName {
		_contractNameRule = append(_contractNameRule, _contractNameItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "AddressUpdate", _contractNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAddressUpdate)
				if err := _Token.contract.UnpackLog(event, "AddressUpdate", log); err != nil {
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

// ParseAddressUpdate is a log parse operation binding the contract event 0xfc08d1253c81bcd5444fc7056ef1f5a5df4c9220b6fd70d7449267f1f0f29918.
//
// Solidity: event AddressUpdate(bytes32 indexed _contractName, address _contractAddress)
func (_Token *TokenFilterer) ParseAddressUpdate(log types.Log) (*TokenAddressUpdate, error) {
	event := new(TokenAddressUpdate)
	if err := _Token.contract.UnpackLog(event, "AddressUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenOwnerUpdateIterator is returned from FilterOwnerUpdate and is used to iterate over the raw logs and unpacked data for OwnerUpdate events raised by the Token contract.
type TokenOwnerUpdateIterator struct {
	Event *TokenOwnerUpdate // Event containing the contract specifics and raw log

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
func (it *TokenOwnerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnerUpdate)
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
		it.Event = new(TokenOwnerUpdate)
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
func (it *TokenOwnerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnerUpdate represents a OwnerUpdate event raised by the Token contract.
type TokenOwnerUpdate struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdate is a free log retrieval operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_Token *TokenFilterer) FilterOwnerUpdate(opts *bind.FilterOpts, _prevOwner []common.Address, _newOwner []common.Address) (*TokenOwnerUpdateIterator, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenOwnerUpdateIterator{contract: _Token.contract, event: "OwnerUpdate", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdate is a free log subscription operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_Token *TokenFilterer) WatchOwnerUpdate(opts *bind.WatchOpts, sink chan<- *TokenOwnerUpdate, _prevOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnerUpdate)
				if err := _Token.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
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

// ParseOwnerUpdate is a log parse operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_Token *TokenFilterer) ParseOwnerUpdate(log types.Log) (*TokenOwnerUpdate, error) {
	event := new(TokenOwnerUpdate)
	if err := _Token.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
