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

// SimSwapMetaData contains all meta data concerning the SimSwap contract.
var SimSwapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"simswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// SimSwapABI is the input ABI used to generate the binding from.
// Deprecated: Use SimSwapMetaData.ABI instead.
var SimSwapABI = SimSwapMetaData.ABI

// SimSwap is an auto generated Go binding around an Ethereum contract.
type SimSwap struct {
	SimSwapCaller     // Read-only binding to the contract
	SimSwapTransactor // Write-only binding to the contract
	SimSwapFilterer   // Log filterer for contract events
}

// SimSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimSwapSession struct {
	Contract     *SimSwap          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimSwapCallerSession struct {
	Contract *SimSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SimSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimSwapTransactorSession struct {
	Contract     *SimSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SimSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimSwapRaw struct {
	Contract *SimSwap // Generic contract binding to access the raw methods on
}

// SimSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimSwapCallerRaw struct {
	Contract *SimSwapCaller // Generic read-only contract binding to access the raw methods on
}

// SimSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimSwapTransactorRaw struct {
	Contract *SimSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimSwap creates a new instance of SimSwap, bound to a specific deployed contract.
func NewSimSwap(address common.Address, backend bind.ContractBackend) (*SimSwap, error) {
	contract, err := bindSimSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimSwap{SimSwapCaller: SimSwapCaller{contract: contract}, SimSwapTransactor: SimSwapTransactor{contract: contract}, SimSwapFilterer: SimSwapFilterer{contract: contract}}, nil
}

// NewSimSwapCaller creates a new read-only instance of SimSwap, bound to a specific deployed contract.
func NewSimSwapCaller(address common.Address, caller bind.ContractCaller) (*SimSwapCaller, error) {
	contract, err := bindSimSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimSwapCaller{contract: contract}, nil
}

// NewSimSwapTransactor creates a new write-only instance of SimSwap, bound to a specific deployed contract.
func NewSimSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*SimSwapTransactor, error) {
	contract, err := bindSimSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimSwapTransactor{contract: contract}, nil
}

// NewSimSwapFilterer creates a new log filterer instance of SimSwap, bound to a specific deployed contract.
func NewSimSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*SimSwapFilterer, error) {
	contract, err := bindSimSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimSwapFilterer{contract: contract}, nil
}

// bindSimSwap binds a generic wrapper to an already deployed contract.
func bindSimSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimSwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimSwap *SimSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimSwap.Contract.SimSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimSwap *SimSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimSwap.Contract.SimSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimSwap *SimSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimSwap.Contract.SimSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimSwap *SimSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimSwap *SimSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimSwap *SimSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimSwap.Contract.contract.Transact(opts, method, params...)
}

// Simswap is a paid mutator transaction binding the contract method 0x96d27420.
//
// Solidity: function simswap(address tokenIn, address tokenOut, address router, bytes data) payable returns(uint256 returnAmount, bytes returnData)
func (_SimSwap *SimSwapTransactor) Simswap(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, router common.Address, data []byte) (*types.Transaction, error) {
	return _SimSwap.contract.Transact(opts, "simswap", tokenIn, tokenOut, router, data)
}

// Simswap is a paid mutator transaction binding the contract method 0x96d27420.
//
// Solidity: function simswap(address tokenIn, address tokenOut, address router, bytes data) payable returns(uint256 returnAmount, bytes returnData)
func (_SimSwap *SimSwapSession) Simswap(tokenIn common.Address, tokenOut common.Address, router common.Address, data []byte) (*types.Transaction, error) {
	return _SimSwap.Contract.Simswap(&_SimSwap.TransactOpts, tokenIn, tokenOut, router, data)
}

// Simswap is a paid mutator transaction binding the contract method 0x96d27420.
//
// Solidity: function simswap(address tokenIn, address tokenOut, address router, bytes data) payable returns(uint256 returnAmount, bytes returnData)
func (_SimSwap *SimSwapTransactorSession) Simswap(tokenIn common.Address, tokenOut common.Address, router common.Address, data []byte) (*types.Transaction, error) {
	return _SimSwap.Contract.Simswap(&_SimSwap.TransactOpts, tokenIn, tokenOut, router, data)
}
