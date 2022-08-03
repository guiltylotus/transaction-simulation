// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// SimCallMetaData contains all meta data concerning the SimCall contract.
var SimCallMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requestAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"simswap\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"results\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// SimCallABI is the input ABI used to generate the binding from.
// Deprecated: Use SimCallMetaData.ABI instead.
var SimCallABI = SimCallMetaData.ABI

// SimCall is an auto generated Go binding around an Ethereum contract.
type SimCall struct {
	SimCallCaller     // Read-only binding to the contract
	SimCallTransactor // Write-only binding to the contract
	SimCallFilterer   // Log filterer for contract events
}

// SimCallCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimCallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimCallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimCallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimCallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimCallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimCallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimCallSession struct {
	Contract     *SimCall          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimCallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimCallCallerSession struct {
	Contract *SimCallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SimCallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimCallTransactorSession struct {
	Contract     *SimCallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SimCallRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimCallRaw struct {
	Contract *SimCall // Generic contract binding to access the raw methods on
}

// SimCallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimCallCallerRaw struct {
	Contract *SimCallCaller // Generic read-only contract binding to access the raw methods on
}

// SimCallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimCallTransactorRaw struct {
	Contract *SimCallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimCall creates a new instance of SimCall, bound to a specific deployed contract.
func NewSimCall(address common.Address, backend bind.ContractBackend) (*SimCall, error) {
	contract, err := bindSimCall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimCall{SimCallCaller: SimCallCaller{contract: contract}, SimCallTransactor: SimCallTransactor{contract: contract}, SimCallFilterer: SimCallFilterer{contract: contract}}, nil
}

// NewSimCallCaller creates a new read-only instance of SimCall, bound to a specific deployed contract.
func NewSimCallCaller(address common.Address, caller bind.ContractCaller) (*SimCallCaller, error) {
	contract, err := bindSimCall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimCallCaller{contract: contract}, nil
}

// NewSimCallTransactor creates a new write-only instance of SimCall, bound to a specific deployed contract.
func NewSimCallTransactor(address common.Address, transactor bind.ContractTransactor) (*SimCallTransactor, error) {
	contract, err := bindSimCall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimCallTransactor{contract: contract}, nil
}

// NewSimCallFilterer creates a new log filterer instance of SimCall, bound to a specific deployed contract.
func NewSimCallFilterer(address common.Address, filterer bind.ContractFilterer) (*SimCallFilterer, error) {
	contract, err := bindSimCall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimCallFilterer{contract: contract}, nil
}

// bindSimCall binds a generic wrapper to an already deployed contract.
func bindSimCall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimCallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimCall *SimCallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimCall.Contract.SimCallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimCall *SimCallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimCall.Contract.SimCallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimCall *SimCallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimCall.Contract.SimCallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimCall *SimCallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimCall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimCall *SimCallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimCall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimCall *SimCallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimCall.Contract.contract.Transact(opts, method, params...)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address token, address spender) returns(uint256)
func (_SimCall *SimCallTransactor) Approve(opts *bind.TransactOpts, token common.Address, spender common.Address) (*types.Transaction, error) {
	return _SimCall.contract.Transact(opts, "approve", token, spender)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address token, address spender) returns(uint256)
func (_SimCall *SimCallSession) Approve(token common.Address, spender common.Address) (*types.Transaction, error) {
	return _SimCall.Contract.Approve(&_SimCall.TransactOpts, token, spender)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address token, address spender) returns(uint256)
func (_SimCall *SimCallTransactorSession) Approve(token common.Address, spender common.Address) (*types.Transaction, error) {
	return _SimCall.Contract.Approve(&_SimCall.TransactOpts, token, spender)
}

// Simswap is a paid mutator transaction binding the contract method 0xdb31da58.
//
// Solidity: function simswap(address tokenAddress, uint256 requestAmount, address router, bytes data) payable returns(bytes results)
func (_SimCall *SimCallTransactor) Simswap(opts *bind.TransactOpts, tokenAddress common.Address, requestAmount *big.Int, router common.Address, data []byte) (*types.Transaction, error) {
	return _SimCall.contract.Transact(opts, "simswap", tokenAddress, requestAmount, router, data)
}

// Simswap is a paid mutator transaction binding the contract method 0xdb31da58.
//
// Solidity: function simswap(address tokenAddress, uint256 requestAmount, address router, bytes data) payable returns(bytes results)
func (_SimCall *SimCallSession) Simswap(tokenAddress common.Address, requestAmount *big.Int, router common.Address, data []byte) (*types.Transaction, error) {
	return _SimCall.Contract.Simswap(&_SimCall.TransactOpts, tokenAddress, requestAmount, router, data)
}

// Simswap is a paid mutator transaction binding the contract method 0xdb31da58.
//
// Solidity: function simswap(address tokenAddress, uint256 requestAmount, address router, bytes data) payable returns(bytes results)
func (_SimCall *SimCallTransactorSession) Simswap(tokenAddress common.Address, requestAmount *big.Int, router common.Address, data []byte) (*types.Transaction, error) {
	return _SimCall.Contract.Simswap(&_SimCall.TransactOpts, tokenAddress, requestAmount, router, data)
}
