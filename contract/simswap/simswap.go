// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simswap

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"getallowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"token1Diff\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"}],\"name\":\"getbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"token1Diff\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"simswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"token1Diff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token2Diff\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address token, address spender) returns(uint256)
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, token common.Address, spender common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", token, spender)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address token, address spender) returns(uint256)
func (_Contract *ContractSession) Approve(token common.Address, spender common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, token, spender)
}

// Approve is a paid mutator transaction binding the contract method 0x7e5465ba.
//
// Solidity: function approve(address token, address spender) returns(uint256)
func (_Contract *ContractTransactorSession) Approve(token common.Address, spender common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, token, spender)
}

// Getallowance is a paid mutator transaction binding the contract method 0x21c4f09f.
//
// Solidity: function getallowance(address token, address router) payable returns(uint256 token1Diff)
func (_Contract *ContractTransactor) Getallowance(opts *bind.TransactOpts, token common.Address, router common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "getallowance", token, router)
}

// Getallowance is a paid mutator transaction binding the contract method 0x21c4f09f.
//
// Solidity: function getallowance(address token, address router) payable returns(uint256 token1Diff)
func (_Contract *ContractSession) Getallowance(token common.Address, router common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Getallowance(&_Contract.TransactOpts, token, router)
}

// Getallowance is a paid mutator transaction binding the contract method 0x21c4f09f.
//
// Solidity: function getallowance(address token, address router) payable returns(uint256 token1Diff)
func (_Contract *ContractTransactorSession) Getallowance(token common.Address, router common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Getallowance(&_Contract.TransactOpts, token, router)
}

// Getbalance is a paid mutator transaction binding the contract method 0x68116177.
//
// Solidity: function getbalance(address tokenIn) payable returns(uint256 token1Diff)
func (_Contract *ContractTransactor) Getbalance(opts *bind.TransactOpts, tokenIn common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "getbalance", tokenIn)
}

// Getbalance is a paid mutator transaction binding the contract method 0x68116177.
//
// Solidity: function getbalance(address tokenIn) payable returns(uint256 token1Diff)
func (_Contract *ContractSession) Getbalance(tokenIn common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Getbalance(&_Contract.TransactOpts, tokenIn)
}

// Getbalance is a paid mutator transaction binding the contract method 0x68116177.
//
// Solidity: function getbalance(address tokenIn) payable returns(uint256 token1Diff)
func (_Contract *ContractTransactorSession) Getbalance(tokenIn common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Getbalance(&_Contract.TransactOpts, tokenIn)
}

// Simswap is a paid mutator transaction binding the contract method 0x96d27420.
//
// Solidity: function simswap(address tokenIn, address tokenOut, address router, bytes data) payable returns(uint256 token1Diff, uint256 token2Diff)
func (_Contract *ContractTransactor) Simswap(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, router common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "simswap", tokenIn, tokenOut, router, data)
}

// Simswap is a paid mutator transaction binding the contract method 0x96d27420.
//
// Solidity: function simswap(address tokenIn, address tokenOut, address router, bytes data) payable returns(uint256 token1Diff, uint256 token2Diff)
func (_Contract *ContractSession) Simswap(tokenIn common.Address, tokenOut common.Address, router common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.Simswap(&_Contract.TransactOpts, tokenIn, tokenOut, router, data)
}

// Simswap is a paid mutator transaction binding the contract method 0x96d27420.
//
// Solidity: function simswap(address tokenIn, address tokenOut, address router, bytes data) payable returns(uint256 token1Diff, uint256 token2Diff)
func (_Contract *ContractTransactorSession) Simswap(tokenIn common.Address, tokenOut common.Address, router common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.Simswap(&_Contract.TransactOpts, tokenIn, tokenOut, router, data)
}
