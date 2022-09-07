// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paraswaptransfer

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

// ParaswapTransferProxyMetaData contains all meta data concerning the ParaswapTransferProxy contract.
var ParaswapTransferProxyMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MyTransferDebug\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ParaswapTransferProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ParaswapTransferProxyMetaData.ABI instead.
var ParaswapTransferProxyABI = ParaswapTransferProxyMetaData.ABI

// ParaswapTransferProxy is an auto generated Go binding around an Ethereum contract.
type ParaswapTransferProxy struct {
	ParaswapTransferProxyCaller     // Read-only binding to the contract
	ParaswapTransferProxyTransactor // Write-only binding to the contract
	ParaswapTransferProxyFilterer   // Log filterer for contract events
}

// ParaswapTransferProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParaswapTransferProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParaswapTransferProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParaswapTransferProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParaswapTransferProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParaswapTransferProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParaswapTransferProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParaswapTransferProxySession struct {
	Contract     *ParaswapTransferProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ParaswapTransferProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParaswapTransferProxyCallerSession struct {
	Contract *ParaswapTransferProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ParaswapTransferProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParaswapTransferProxyTransactorSession struct {
	Contract     *ParaswapTransferProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ParaswapTransferProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParaswapTransferProxyRaw struct {
	Contract *ParaswapTransferProxy // Generic contract binding to access the raw methods on
}

// ParaswapTransferProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParaswapTransferProxyCallerRaw struct {
	Contract *ParaswapTransferProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ParaswapTransferProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParaswapTransferProxyTransactorRaw struct {
	Contract *ParaswapTransferProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParaswapTransferProxy creates a new instance of ParaswapTransferProxy, bound to a specific deployed contract.
func NewParaswapTransferProxy(address common.Address, backend bind.ContractBackend) (*ParaswapTransferProxy, error) {
	contract, err := bindParaswapTransferProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ParaswapTransferProxy{ParaswapTransferProxyCaller: ParaswapTransferProxyCaller{contract: contract}, ParaswapTransferProxyTransactor: ParaswapTransferProxyTransactor{contract: contract}, ParaswapTransferProxyFilterer: ParaswapTransferProxyFilterer{contract: contract}}, nil
}

// NewParaswapTransferProxyCaller creates a new read-only instance of ParaswapTransferProxy, bound to a specific deployed contract.
func NewParaswapTransferProxyCaller(address common.Address, caller bind.ContractCaller) (*ParaswapTransferProxyCaller, error) {
	contract, err := bindParaswapTransferProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParaswapTransferProxyCaller{contract: contract}, nil
}

// NewParaswapTransferProxyTransactor creates a new write-only instance of ParaswapTransferProxy, bound to a specific deployed contract.
func NewParaswapTransferProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ParaswapTransferProxyTransactor, error) {
	contract, err := bindParaswapTransferProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParaswapTransferProxyTransactor{contract: contract}, nil
}

// NewParaswapTransferProxyFilterer creates a new log filterer instance of ParaswapTransferProxy, bound to a specific deployed contract.
func NewParaswapTransferProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ParaswapTransferProxyFilterer, error) {
	contract, err := bindParaswapTransferProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParaswapTransferProxyFilterer{contract: contract}, nil
}

// bindParaswapTransferProxy binds a generic wrapper to an already deployed contract.
func bindParaswapTransferProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParaswapTransferProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParaswapTransferProxy *ParaswapTransferProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParaswapTransferProxy.Contract.ParaswapTransferProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParaswapTransferProxy *ParaswapTransferProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParaswapTransferProxy.Contract.ParaswapTransferProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParaswapTransferProxy *ParaswapTransferProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParaswapTransferProxy.Contract.ParaswapTransferProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParaswapTransferProxy *ParaswapTransferProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParaswapTransferProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParaswapTransferProxy *ParaswapTransferProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParaswapTransferProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParaswapTransferProxy *ParaswapTransferProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParaswapTransferProxy.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParaswapTransferProxy *ParaswapTransferProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ParaswapTransferProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParaswapTransferProxy *ParaswapTransferProxySession) Owner() (common.Address, error) {
	return _ParaswapTransferProxy.Contract.Owner(&_ParaswapTransferProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParaswapTransferProxy *ParaswapTransferProxyCallerSession) Owner() (common.Address, error) {
	return _ParaswapTransferProxy.Contract.Owner(&_ParaswapTransferProxy.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParaswapTransferProxy *ParaswapTransferProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParaswapTransferProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParaswapTransferProxy *ParaswapTransferProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _ParaswapTransferProxy.Contract.RenounceOwnership(&_ParaswapTransferProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParaswapTransferProxy *ParaswapTransferProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ParaswapTransferProxy.Contract.RenounceOwnership(&_ParaswapTransferProxy.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x15dacbea.
//
// Solidity: function transferFrom(address token, address from, address to, uint256 amount) returns()
func (_ParaswapTransferProxy *ParaswapTransferProxyTransactor) TransferFrom(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ParaswapTransferProxy.contract.Transact(opts, "transferFrom", token, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x15dacbea.
//
// Solidity: function transferFrom(address token, address from, address to, uint256 amount) returns()
func (_ParaswapTransferProxy *ParaswapTransferProxySession) TransferFrom(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ParaswapTransferProxy.Contract.TransferFrom(&_ParaswapTransferProxy.TransactOpts, token, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x15dacbea.
//
// Solidity: function transferFrom(address token, address from, address to, uint256 amount) returns()
func (_ParaswapTransferProxy *ParaswapTransferProxyTransactorSession) TransferFrom(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ParaswapTransferProxy.Contract.TransferFrom(&_ParaswapTransferProxy.TransactOpts, token, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParaswapTransferProxy *ParaswapTransferProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ParaswapTransferProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParaswapTransferProxy *ParaswapTransferProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ParaswapTransferProxy.Contract.TransferOwnership(&_ParaswapTransferProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParaswapTransferProxy *ParaswapTransferProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ParaswapTransferProxy.Contract.TransferOwnership(&_ParaswapTransferProxy.TransactOpts, newOwner)
}

// ParaswapTransferProxyMyTransferDebugIterator is returned from FilterMyTransferDebug and is used to iterate over the raw logs and unpacked data for MyTransferDebug events raised by the ParaswapTransferProxy contract.
type ParaswapTransferProxyMyTransferDebugIterator struct {
	Event *ParaswapTransferProxyMyTransferDebug // Event containing the contract specifics and raw log

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
func (it *ParaswapTransferProxyMyTransferDebugIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParaswapTransferProxyMyTransferDebug)
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
		it.Event = new(ParaswapTransferProxyMyTransferDebug)
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
func (it *ParaswapTransferProxyMyTransferDebugIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParaswapTransferProxyMyTransferDebugIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParaswapTransferProxyMyTransferDebug represents a MyTransferDebug event raised by the ParaswapTransferProxy contract.
type ParaswapTransferProxyMyTransferDebug struct {
	Sender      common.Address
	DstReceiver common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMyTransferDebug is a free log retrieval operation binding the contract event 0xc1209790a4467a38d547e89c25c60977b25e17eacd7a0234750ba270c059fc48.
//
// Solidity: event MyTransferDebug(address sender, address dstReceiver, uint256 amount)
func (_ParaswapTransferProxy *ParaswapTransferProxyFilterer) FilterMyTransferDebug(opts *bind.FilterOpts) (*ParaswapTransferProxyMyTransferDebugIterator, error) {

	logs, sub, err := _ParaswapTransferProxy.contract.FilterLogs(opts, "MyTransferDebug")
	if err != nil {
		return nil, err
	}
	return &ParaswapTransferProxyMyTransferDebugIterator{contract: _ParaswapTransferProxy.contract, event: "MyTransferDebug", logs: logs, sub: sub}, nil
}

// WatchMyTransferDebug is a free log subscription operation binding the contract event 0xc1209790a4467a38d547e89c25c60977b25e17eacd7a0234750ba270c059fc48.
//
// Solidity: event MyTransferDebug(address sender, address dstReceiver, uint256 amount)
func (_ParaswapTransferProxy *ParaswapTransferProxyFilterer) WatchMyTransferDebug(opts *bind.WatchOpts, sink chan<- *ParaswapTransferProxyMyTransferDebug) (event.Subscription, error) {

	logs, sub, err := _ParaswapTransferProxy.contract.WatchLogs(opts, "MyTransferDebug")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParaswapTransferProxyMyTransferDebug)
				if err := _ParaswapTransferProxy.contract.UnpackLog(event, "MyTransferDebug", log); err != nil {
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

// ParseMyTransferDebug is a log parse operation binding the contract event 0xc1209790a4467a38d547e89c25c60977b25e17eacd7a0234750ba270c059fc48.
//
// Solidity: event MyTransferDebug(address sender, address dstReceiver, uint256 amount)
func (_ParaswapTransferProxy *ParaswapTransferProxyFilterer) ParseMyTransferDebug(log types.Log) (*ParaswapTransferProxyMyTransferDebug, error) {
	event := new(ParaswapTransferProxyMyTransferDebug)
	if err := _ParaswapTransferProxy.contract.UnpackLog(event, "MyTransferDebug", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParaswapTransferProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ParaswapTransferProxy contract.
type ParaswapTransferProxyOwnershipTransferredIterator struct {
	Event *ParaswapTransferProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ParaswapTransferProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParaswapTransferProxyOwnershipTransferred)
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
		it.Event = new(ParaswapTransferProxyOwnershipTransferred)
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
func (it *ParaswapTransferProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParaswapTransferProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParaswapTransferProxyOwnershipTransferred represents a OwnershipTransferred event raised by the ParaswapTransferProxy contract.
type ParaswapTransferProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ParaswapTransferProxy *ParaswapTransferProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ParaswapTransferProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ParaswapTransferProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ParaswapTransferProxyOwnershipTransferredIterator{contract: _ParaswapTransferProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ParaswapTransferProxy *ParaswapTransferProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ParaswapTransferProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ParaswapTransferProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParaswapTransferProxyOwnershipTransferred)
				if err := _ParaswapTransferProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ParaswapTransferProxy *ParaswapTransferProxyFilterer) ParseOwnershipTransferred(log types.Log) (*ParaswapTransferProxyOwnershipTransferred, error) {
	event := new(ParaswapTransferProxyOwnershipTransferred)
	if err := _ParaswapTransferProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
