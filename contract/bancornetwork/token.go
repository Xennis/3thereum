// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bancornetwork

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
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_onlyOwnerCanUpdateRegistry\",\"type\":\"bool\"}],\"name\":\"restrictRegistryUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_register\",\"type\":\"bool\"}],\"name\":\"registerEtherToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getReturnByPath\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"claimAndConvertFor2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"onlyOwnerCanUpdateRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"convert2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAffiliateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prevRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"rateByPath\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"etherTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_bancorX\",\"type\":\"address\"},{\"name\":\"_conversionId\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"completeXConversion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"convertFor2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"claimAndConvertFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restoreRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"convertByPath\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_targetBlockchain\",\"type\":\"bytes32\"},{\"name\":\"_targetAccount\",\"type\":\"bytes32\"},{\"name\":\"_conversionId\",\"type\":\"uint256\"}],\"name\":\"xConvert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"claimAndConvert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"convertFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_targetBlockchain\",\"type\":\"bytes32\"},{\"name\":\"_targetAccount\",\"type\":\"bytes32\"},{\"name\":\"_conversionId\",\"type\":\"uint256\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"xConvert2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sourceToken\",\"type\":\"address\"},{\"name\":\"_targetToken\",\"type\":\"address\"}],\"name\":\"conversionPath\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"claimAndConvert2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"convert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxAffiliateFee\",\"type\":\"uint256\"}],\"name\":\"setMaxAffiliateFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_smartToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_fromToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_toToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"Conversion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]",
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

// ConversionPath is a free data retrieval call binding the contract method 0xd734fa19.
//
// Solidity: function conversionPath(address _sourceToken, address _targetToken) view returns(address[])
func (_Token *TokenCaller) ConversionPath(opts *bind.CallOpts, _sourceToken common.Address, _targetToken common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "conversionPath", _sourceToken, _targetToken)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ConversionPath is a free data retrieval call binding the contract method 0xd734fa19.
//
// Solidity: function conversionPath(address _sourceToken, address _targetToken) view returns(address[])
func (_Token *TokenSession) ConversionPath(_sourceToken common.Address, _targetToken common.Address) ([]common.Address, error) {
	return _Token.Contract.ConversionPath(&_Token.CallOpts, _sourceToken, _targetToken)
}

// ConversionPath is a free data retrieval call binding the contract method 0xd734fa19.
//
// Solidity: function conversionPath(address _sourceToken, address _targetToken) view returns(address[])
func (_Token *TokenCallerSession) ConversionPath(_sourceToken common.Address, _targetToken common.Address) ([]common.Address, error) {
	return _Token.Contract.ConversionPath(&_Token.CallOpts, _sourceToken, _targetToken)
}

// EtherTokens is a free data retrieval call binding the contract method 0x8077ccf7.
//
// Solidity: function etherTokens(address ) view returns(bool)
func (_Token *TokenCaller) EtherTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "etherTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EtherTokens is a free data retrieval call binding the contract method 0x8077ccf7.
//
// Solidity: function etherTokens(address ) view returns(bool)
func (_Token *TokenSession) EtherTokens(arg0 common.Address) (bool, error) {
	return _Token.Contract.EtherTokens(&_Token.CallOpts, arg0)
}

// EtherTokens is a free data retrieval call binding the contract method 0x8077ccf7.
//
// Solidity: function etherTokens(address ) view returns(bool)
func (_Token *TokenCallerSession) EtherTokens(arg0 common.Address) (bool, error) {
	return _Token.Contract.EtherTokens(&_Token.CallOpts, arg0)
}

// GetReturnByPath is a free data retrieval call binding the contract method 0x0c8496cc.
//
// Solidity: function getReturnByPath(address[] _path, uint256 _amount) view returns(uint256, uint256)
func (_Token *TokenCaller) GetReturnByPath(opts *bind.CallOpts, _path []common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getReturnByPath", _path, _amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetReturnByPath is a free data retrieval call binding the contract method 0x0c8496cc.
//
// Solidity: function getReturnByPath(address[] _path, uint256 _amount) view returns(uint256, uint256)
func (_Token *TokenSession) GetReturnByPath(_path []common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _Token.Contract.GetReturnByPath(&_Token.CallOpts, _path, _amount)
}

// GetReturnByPath is a free data retrieval call binding the contract method 0x0c8496cc.
//
// Solidity: function getReturnByPath(address[] _path, uint256 _amount) view returns(uint256, uint256)
func (_Token *TokenCallerSession) GetReturnByPath(_path []common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _Token.Contract.GetReturnByPath(&_Token.CallOpts, _path, _amount)
}

// MaxAffiliateFee is a free data retrieval call binding the contract method 0x5d732ff2.
//
// Solidity: function maxAffiliateFee() view returns(uint256)
func (_Token *TokenCaller) MaxAffiliateFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "maxAffiliateFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAffiliateFee is a free data retrieval call binding the contract method 0x5d732ff2.
//
// Solidity: function maxAffiliateFee() view returns(uint256)
func (_Token *TokenSession) MaxAffiliateFee() (*big.Int, error) {
	return _Token.Contract.MaxAffiliateFee(&_Token.CallOpts)
}

// MaxAffiliateFee is a free data retrieval call binding the contract method 0x5d732ff2.
//
// Solidity: function maxAffiliateFee() view returns(uint256)
func (_Token *TokenCallerSession) MaxAffiliateFee() (*big.Int, error) {
	return _Token.Contract.MaxAffiliateFee(&_Token.CallOpts)
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

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_Token *TokenCaller) OnlyOwnerCanUpdateRegistry(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "onlyOwnerCanUpdateRegistry")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_Token *TokenSession) OnlyOwnerCanUpdateRegistry() (bool, error) {
	return _Token.Contract.OnlyOwnerCanUpdateRegistry(&_Token.CallOpts)
}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_Token *TokenCallerSession) OnlyOwnerCanUpdateRegistry() (bool, error) {
	return _Token.Contract.OnlyOwnerCanUpdateRegistry(&_Token.CallOpts)
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

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_Token *TokenCaller) PrevRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "prevRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_Token *TokenSession) PrevRegistry() (common.Address, error) {
	return _Token.Contract.PrevRegistry(&_Token.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_Token *TokenCallerSession) PrevRegistry() (common.Address, error) {
	return _Token.Contract.PrevRegistry(&_Token.CallOpts)
}

// RateByPath is a free data retrieval call binding the contract method 0x7f9c0ecd.
//
// Solidity: function rateByPath(address[] _path, uint256 _amount) view returns(uint256)
func (_Token *TokenCaller) RateByPath(opts *bind.CallOpts, _path []common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "rateByPath", _path, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RateByPath is a free data retrieval call binding the contract method 0x7f9c0ecd.
//
// Solidity: function rateByPath(address[] _path, uint256 _amount) view returns(uint256)
func (_Token *TokenSession) RateByPath(_path []common.Address, _amount *big.Int) (*big.Int, error) {
	return _Token.Contract.RateByPath(&_Token.CallOpts, _path, _amount)
}

// RateByPath is a free data retrieval call binding the contract method 0x7f9c0ecd.
//
// Solidity: function rateByPath(address[] _path, uint256 _amount) view returns(uint256)
func (_Token *TokenCallerSession) RateByPath(_path []common.Address, _amount *big.Int) (*big.Int, error) {
	return _Token.Contract.RateByPath(&_Token.CallOpts, _path, _amount)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Token *TokenCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Token *TokenSession) Registry() (common.Address, error) {
	return _Token.Contract.Registry(&_Token.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Token *TokenCallerSession) Registry() (common.Address, error) {
	return _Token.Contract.Registry(&_Token.CallOpts)
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

// ClaimAndConvert is a paid mutator transaction binding the contract method 0xc7ba24bc.
//
// Solidity: function claimAndConvert(address[] _path, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_Token *TokenTransactor) ClaimAndConvert(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimAndConvert", _path, _amount, _minReturn)
}

// ClaimAndConvert is a paid mutator transaction binding the contract method 0xc7ba24bc.
//
// Solidity: function claimAndConvert(address[] _path, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_Token *TokenSession) ClaimAndConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ClaimAndConvert(&_Token.TransactOpts, _path, _amount, _minReturn)
}

// ClaimAndConvert is a paid mutator transaction binding the contract method 0xc7ba24bc.
//
// Solidity: function claimAndConvert(address[] _path, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_Token *TokenTransactorSession) ClaimAndConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ClaimAndConvert(&_Token.TransactOpts, _path, _amount, _minReturn)
}

// ClaimAndConvert2 is a paid mutator transaction binding the contract method 0xe57738e5.
//
// Solidity: function claimAndConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_Token *TokenTransactor) ClaimAndConvert2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimAndConvert2", _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvert2 is a paid mutator transaction binding the contract method 0xe57738e5.
//
// Solidity: function claimAndConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_Token *TokenSession) ClaimAndConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ClaimAndConvert2(&_Token.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvert2 is a paid mutator transaction binding the contract method 0xe57738e5.
//
// Solidity: function claimAndConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_Token *TokenTransactorSession) ClaimAndConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ClaimAndConvert2(&_Token.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvertFor is a paid mutator transaction binding the contract method 0xb1e9932b.
//
// Solidity: function claimAndConvertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_Token *TokenTransactor) ClaimAndConvertFor(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimAndConvertFor", _path, _amount, _minReturn, _beneficiary)
}

// ClaimAndConvertFor is a paid mutator transaction binding the contract method 0xb1e9932b.
//
// Solidity: function claimAndConvertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_Token *TokenSession) ClaimAndConvertFor(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Token.Contract.ClaimAndConvertFor(&_Token.TransactOpts, _path, _amount, _minReturn, _beneficiary)
}

// ClaimAndConvertFor is a paid mutator transaction binding the contract method 0xb1e9932b.
//
// Solidity: function claimAndConvertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_Token *TokenTransactorSession) ClaimAndConvertFor(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Token.Contract.ClaimAndConvertFor(&_Token.TransactOpts, _path, _amount, _minReturn, _beneficiary)
}

// ClaimAndConvertFor2 is a paid mutator transaction binding the contract method 0x2978c10e.
//
// Solidity: function claimAndConvertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_Token *TokenTransactor) ClaimAndConvertFor2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimAndConvertFor2", _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvertFor2 is a paid mutator transaction binding the contract method 0x2978c10e.
//
// Solidity: function claimAndConvertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_Token *TokenSession) ClaimAndConvertFor2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ClaimAndConvertFor2(&_Token.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvertFor2 is a paid mutator transaction binding the contract method 0x2978c10e.
//
// Solidity: function claimAndConvertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_Token *TokenTransactorSession) ClaimAndConvertFor2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ClaimAndConvertFor2(&_Token.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x89f9cc61.
//
// Solidity: function completeXConversion(address[] _path, address _bancorX, uint256 _conversionId, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_Token *TokenTransactor) CompleteXConversion(opts *bind.TransactOpts, _path []common.Address, _bancorX common.Address, _conversionId *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "completeXConversion", _path, _bancorX, _conversionId, _minReturn, _beneficiary)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x89f9cc61.
//
// Solidity: function completeXConversion(address[] _path, address _bancorX, uint256 _conversionId, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_Token *TokenSession) CompleteXConversion(_path []common.Address, _bancorX common.Address, _conversionId *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Token.Contract.CompleteXConversion(&_Token.TransactOpts, _path, _bancorX, _conversionId, _minReturn, _beneficiary)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x89f9cc61.
//
// Solidity: function completeXConversion(address[] _path, address _bancorX, uint256 _conversionId, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_Token *TokenTransactorSession) CompleteXConversion(_path []common.Address, _bancorX common.Address, _conversionId *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Token.Contract.CompleteXConversion(&_Token.TransactOpts, _path, _bancorX, _conversionId, _minReturn, _beneficiary)
}

// Convert is a paid mutator transaction binding the contract method 0xf3898a97.
//
// Solidity: function convert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_Token *TokenTransactor) Convert(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "convert", _path, _amount, _minReturn)
}

// Convert is a paid mutator transaction binding the contract method 0xf3898a97.
//
// Solidity: function convert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_Token *TokenSession) Convert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Convert(&_Token.TransactOpts, _path, _amount, _minReturn)
}

// Convert is a paid mutator transaction binding the contract method 0xf3898a97.
//
// Solidity: function convert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_Token *TokenTransactorSession) Convert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Convert(&_Token.TransactOpts, _path, _amount, _minReturn)
}

// Convert2 is a paid mutator transaction binding the contract method 0x569706eb.
//
// Solidity: function convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Token *TokenTransactor) Convert2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "convert2", _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// Convert2 is a paid mutator transaction binding the contract method 0x569706eb.
//
// Solidity: function convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Token *TokenSession) Convert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Convert2(&_Token.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// Convert2 is a paid mutator transaction binding the contract method 0x569706eb.
//
// Solidity: function convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Token *TokenTransactorSession) Convert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Convert2(&_Token.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// ConvertByPath is a paid mutator transaction binding the contract method 0xb77d239b.
//
// Solidity: function convertByPath(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Token *TokenTransactor) ConvertByPath(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "convertByPath", _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertByPath is a paid mutator transaction binding the contract method 0xb77d239b.
//
// Solidity: function convertByPath(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Token *TokenSession) ConvertByPath(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ConvertByPath(&_Token.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertByPath is a paid mutator transaction binding the contract method 0xb77d239b.
//
// Solidity: function convertByPath(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Token *TokenTransactorSession) ConvertByPath(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ConvertByPath(&_Token.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertFor is a paid mutator transaction binding the contract method 0xc98fefed.
//
// Solidity: function convertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) payable returns(uint256)
func (_Token *TokenTransactor) ConvertFor(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "convertFor", _path, _amount, _minReturn, _beneficiary)
}

// ConvertFor is a paid mutator transaction binding the contract method 0xc98fefed.
//
// Solidity: function convertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) payable returns(uint256)
func (_Token *TokenSession) ConvertFor(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Token.Contract.ConvertFor(&_Token.TransactOpts, _path, _amount, _minReturn, _beneficiary)
}

// ConvertFor is a paid mutator transaction binding the contract method 0xc98fefed.
//
// Solidity: function convertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) payable returns(uint256)
func (_Token *TokenTransactorSession) ConvertFor(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Token.Contract.ConvertFor(&_Token.TransactOpts, _path, _amount, _minReturn, _beneficiary)
}

// ConvertFor2 is a paid mutator transaction binding the contract method 0xab6214ce.
//
// Solidity: function convertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Token *TokenTransactor) ConvertFor2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "convertFor2", _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertFor2 is a paid mutator transaction binding the contract method 0xab6214ce.
//
// Solidity: function convertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Token *TokenSession) ConvertFor2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ConvertFor2(&_Token.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertFor2 is a paid mutator transaction binding the contract method 0xab6214ce.
//
// Solidity: function convertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Token *TokenTransactorSession) ConvertFor2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ConvertFor2(&_Token.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// RegisterEtherToken is a paid mutator transaction binding the contract method 0x02ef521e.
//
// Solidity: function registerEtherToken(address _token, bool _register) returns()
func (_Token *TokenTransactor) RegisterEtherToken(opts *bind.TransactOpts, _token common.Address, _register bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "registerEtherToken", _token, _register)
}

// RegisterEtherToken is a paid mutator transaction binding the contract method 0x02ef521e.
//
// Solidity: function registerEtherToken(address _token, bool _register) returns()
func (_Token *TokenSession) RegisterEtherToken(_token common.Address, _register bool) (*types.Transaction, error) {
	return _Token.Contract.RegisterEtherToken(&_Token.TransactOpts, _token, _register)
}

// RegisterEtherToken is a paid mutator transaction binding the contract method 0x02ef521e.
//
// Solidity: function registerEtherToken(address _token, bool _register) returns()
func (_Token *TokenTransactorSession) RegisterEtherToken(_token common.Address, _register bool) (*types.Transaction, error) {
	return _Token.Contract.RegisterEtherToken(&_Token.TransactOpts, _token, _register)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_Token *TokenTransactor) RestoreRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "restoreRegistry")
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_Token *TokenSession) RestoreRegistry() (*types.Transaction, error) {
	return _Token.Contract.RestoreRegistry(&_Token.TransactOpts)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_Token *TokenTransactorSession) RestoreRegistry() (*types.Transaction, error) {
	return _Token.Contract.RestoreRegistry(&_Token.TransactOpts)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_Token *TokenTransactor) RestrictRegistryUpdate(opts *bind.TransactOpts, _onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "restrictRegistryUpdate", _onlyOwnerCanUpdateRegistry)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_Token *TokenSession) RestrictRegistryUpdate(_onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _Token.Contract.RestrictRegistryUpdate(&_Token.TransactOpts, _onlyOwnerCanUpdateRegistry)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_Token *TokenTransactorSession) RestrictRegistryUpdate(_onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _Token.Contract.RestrictRegistryUpdate(&_Token.TransactOpts, _onlyOwnerCanUpdateRegistry)
}

// SetMaxAffiliateFee is a paid mutator transaction binding the contract method 0xf3bc7d2a.
//
// Solidity: function setMaxAffiliateFee(uint256 _maxAffiliateFee) returns()
func (_Token *TokenTransactor) SetMaxAffiliateFee(opts *bind.TransactOpts, _maxAffiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setMaxAffiliateFee", _maxAffiliateFee)
}

// SetMaxAffiliateFee is a paid mutator transaction binding the contract method 0xf3bc7d2a.
//
// Solidity: function setMaxAffiliateFee(uint256 _maxAffiliateFee) returns()
func (_Token *TokenSession) SetMaxAffiliateFee(_maxAffiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetMaxAffiliateFee(&_Token.TransactOpts, _maxAffiliateFee)
}

// SetMaxAffiliateFee is a paid mutator transaction binding the contract method 0xf3bc7d2a.
//
// Solidity: function setMaxAffiliateFee(uint256 _maxAffiliateFee) returns()
func (_Token *TokenTransactorSession) SetMaxAffiliateFee(_maxAffiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetMaxAffiliateFee(&_Token.TransactOpts, _maxAffiliateFee)
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

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_Token *TokenTransactor) UpdateRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateRegistry")
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_Token *TokenSession) UpdateRegistry() (*types.Transaction, error) {
	return _Token.Contract.UpdateRegistry(&_Token.TransactOpts)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_Token *TokenTransactorSession) UpdateRegistry() (*types.Transaction, error) {
	return _Token.Contract.UpdateRegistry(&_Token.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_Token *TokenTransactor) WithdrawTokens(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdrawTokens", _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_Token *TokenSession) WithdrawTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.WithdrawTokens(&_Token.TransactOpts, _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_Token *TokenTransactorSession) WithdrawTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.WithdrawTokens(&_Token.TransactOpts, _token, _to, _amount)
}

// XConvert is a paid mutator transaction binding the contract method 0xc52173de.
//
// Solidity: function xConvert(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId) payable returns(uint256)
func (_Token *TokenTransactor) XConvert(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "xConvert", _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId)
}

// XConvert is a paid mutator transaction binding the contract method 0xc52173de.
//
// Solidity: function xConvert(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId) payable returns(uint256)
func (_Token *TokenSession) XConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.XConvert(&_Token.TransactOpts, _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId)
}

// XConvert is a paid mutator transaction binding the contract method 0xc52173de.
//
// Solidity: function xConvert(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId) payable returns(uint256)
func (_Token *TokenTransactorSession) XConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.XConvert(&_Token.TransactOpts, _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId)
}

// XConvert2 is a paid mutator transaction binding the contract method 0xcb32564e.
//
// Solidity: function xConvert2(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Token *TokenTransactor) XConvert2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "xConvert2", _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId, _affiliateAccount, _affiliateFee)
}

// XConvert2 is a paid mutator transaction binding the contract method 0xcb32564e.
//
// Solidity: function xConvert2(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Token *TokenSession) XConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.XConvert2(&_Token.TransactOpts, _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId, _affiliateAccount, _affiliateFee)
}

// XConvert2 is a paid mutator transaction binding the contract method 0xcb32564e.
//
// Solidity: function xConvert2(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Token *TokenTransactorSession) XConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.XConvert2(&_Token.TransactOpts, _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId, _affiliateAccount, _affiliateFee)
}

// TokenConversionIterator is returned from FilterConversion and is used to iterate over the raw logs and unpacked data for Conversion events raised by the Token contract.
type TokenConversionIterator struct {
	Event *TokenConversion // Event containing the contract specifics and raw log

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
func (it *TokenConversionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenConversion)
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
		it.Event = new(TokenConversion)
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
func (it *TokenConversionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenConversionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenConversion represents a Conversion event raised by the Token contract.
type TokenConversion struct {
	SmartToken common.Address
	FromToken  common.Address
	ToToken    common.Address
	FromAmount *big.Int
	ToAmount   *big.Int
	Trader     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConversion is a free log retrieval operation binding the contract event 0x7154b38b5dd31bb3122436a96d4e09aba5b323ae1fd580025fab55074334c095.
//
// Solidity: event Conversion(address indexed _smartToken, address indexed _fromToken, address indexed _toToken, uint256 _fromAmount, uint256 _toAmount, address _trader)
func (_Token *TokenFilterer) FilterConversion(opts *bind.FilterOpts, _smartToken []common.Address, _fromToken []common.Address, _toToken []common.Address) (*TokenConversionIterator, error) {

	var _smartTokenRule []interface{}
	for _, _smartTokenItem := range _smartToken {
		_smartTokenRule = append(_smartTokenRule, _smartTokenItem)
	}
	var _fromTokenRule []interface{}
	for _, _fromTokenItem := range _fromToken {
		_fromTokenRule = append(_fromTokenRule, _fromTokenItem)
	}
	var _toTokenRule []interface{}
	for _, _toTokenItem := range _toToken {
		_toTokenRule = append(_toTokenRule, _toTokenItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Conversion", _smartTokenRule, _fromTokenRule, _toTokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenConversionIterator{contract: _Token.contract, event: "Conversion", logs: logs, sub: sub}, nil
}

// WatchConversion is a free log subscription operation binding the contract event 0x7154b38b5dd31bb3122436a96d4e09aba5b323ae1fd580025fab55074334c095.
//
// Solidity: event Conversion(address indexed _smartToken, address indexed _fromToken, address indexed _toToken, uint256 _fromAmount, uint256 _toAmount, address _trader)
func (_Token *TokenFilterer) WatchConversion(opts *bind.WatchOpts, sink chan<- *TokenConversion, _smartToken []common.Address, _fromToken []common.Address, _toToken []common.Address) (event.Subscription, error) {

	var _smartTokenRule []interface{}
	for _, _smartTokenItem := range _smartToken {
		_smartTokenRule = append(_smartTokenRule, _smartTokenItem)
	}
	var _fromTokenRule []interface{}
	for _, _fromTokenItem := range _fromToken {
		_fromTokenRule = append(_fromTokenRule, _fromTokenItem)
	}
	var _toTokenRule []interface{}
	for _, _toTokenItem := range _toToken {
		_toTokenRule = append(_toTokenRule, _toTokenItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Conversion", _smartTokenRule, _fromTokenRule, _toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenConversion)
				if err := _Token.contract.UnpackLog(event, "Conversion", log); err != nil {
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

// ParseConversion is a log parse operation binding the contract event 0x7154b38b5dd31bb3122436a96d4e09aba5b323ae1fd580025fab55074334c095.
//
// Solidity: event Conversion(address indexed _smartToken, address indexed _fromToken, address indexed _toToken, uint256 _fromAmount, uint256 _toAmount, address _trader)
func (_Token *TokenFilterer) ParseConversion(log types.Log) (*TokenConversion, error) {
	event := new(TokenConversion)
	if err := _Token.contract.UnpackLog(event, "Conversion", log); err != nil {
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
