// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paraswaputil

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

// ParaswapUtilsMetaData contains all meta data concerning the ParaswapUtils contract.
var ParaswapUtilsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MyUtilDebug\",\"type\":\"event\"}]",
}

// ParaswapUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use ParaswapUtilsMetaData.ABI instead.
var ParaswapUtilsABI = ParaswapUtilsMetaData.ABI

// ParaswapUtils is an auto generated Go binding around an Ethereum contract.
type ParaswapUtils struct {
	ParaswapUtilsCaller     // Read-only binding to the contract
	ParaswapUtilsTransactor // Write-only binding to the contract
	ParaswapUtilsFilterer   // Log filterer for contract events
}

// ParaswapUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParaswapUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParaswapUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParaswapUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParaswapUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParaswapUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParaswapUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParaswapUtilsSession struct {
	Contract     *ParaswapUtils    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParaswapUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParaswapUtilsCallerSession struct {
	Contract *ParaswapUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ParaswapUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParaswapUtilsTransactorSession struct {
	Contract     *ParaswapUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ParaswapUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParaswapUtilsRaw struct {
	Contract *ParaswapUtils // Generic contract binding to access the raw methods on
}

// ParaswapUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParaswapUtilsCallerRaw struct {
	Contract *ParaswapUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ParaswapUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParaswapUtilsTransactorRaw struct {
	Contract *ParaswapUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParaswapUtils creates a new instance of ParaswapUtils, bound to a specific deployed contract.
func NewParaswapUtils(address common.Address, backend bind.ContractBackend) (*ParaswapUtils, error) {
	contract, err := bindParaswapUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ParaswapUtils{ParaswapUtilsCaller: ParaswapUtilsCaller{contract: contract}, ParaswapUtilsTransactor: ParaswapUtilsTransactor{contract: contract}, ParaswapUtilsFilterer: ParaswapUtilsFilterer{contract: contract}}, nil
}

// NewParaswapUtilsCaller creates a new read-only instance of ParaswapUtils, bound to a specific deployed contract.
func NewParaswapUtilsCaller(address common.Address, caller bind.ContractCaller) (*ParaswapUtilsCaller, error) {
	contract, err := bindParaswapUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParaswapUtilsCaller{contract: contract}, nil
}

// NewParaswapUtilsTransactor creates a new write-only instance of ParaswapUtils, bound to a specific deployed contract.
func NewParaswapUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ParaswapUtilsTransactor, error) {
	contract, err := bindParaswapUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParaswapUtilsTransactor{contract: contract}, nil
}

// NewParaswapUtilsFilterer creates a new log filterer instance of ParaswapUtils, bound to a specific deployed contract.
func NewParaswapUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ParaswapUtilsFilterer, error) {
	contract, err := bindParaswapUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParaswapUtilsFilterer{contract: contract}, nil
}

// bindParaswapUtils binds a generic wrapper to an already deployed contract.
func bindParaswapUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParaswapUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParaswapUtils *ParaswapUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParaswapUtils.Contract.ParaswapUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParaswapUtils *ParaswapUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParaswapUtils.Contract.ParaswapUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParaswapUtils *ParaswapUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParaswapUtils.Contract.ParaswapUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParaswapUtils *ParaswapUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParaswapUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParaswapUtils *ParaswapUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParaswapUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParaswapUtils *ParaswapUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParaswapUtils.Contract.contract.Transact(opts, method, params...)
}

// ParaswapUtilsMyUtilDebugIterator is returned from FilterMyUtilDebug and is used to iterate over the raw logs and unpacked data for MyUtilDebug events raised by the ParaswapUtils contract.
type ParaswapUtilsMyUtilDebugIterator struct {
	Event *ParaswapUtilsMyUtilDebug // Event containing the contract specifics and raw log

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
func (it *ParaswapUtilsMyUtilDebugIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParaswapUtilsMyUtilDebug)
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
		it.Event = new(ParaswapUtilsMyUtilDebug)
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
func (it *ParaswapUtilsMyUtilDebugIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParaswapUtilsMyUtilDebugIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParaswapUtilsMyUtilDebug represents a MyUtilDebug event raised by the ParaswapUtils contract.
type ParaswapUtilsMyUtilDebug struct {
	Src    common.Address
	Dst    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMyUtilDebug is a free log retrieval operation binding the contract event 0x6ae69222768d00c831213d6afd4cec422b4d4f89949133c7c4e26ba5e02a5fa2.
//
// Solidity: event MyUtilDebug(address src, address dst, uint256 amount)
func (_ParaswapUtils *ParaswapUtilsFilterer) FilterMyUtilDebug(opts *bind.FilterOpts) (*ParaswapUtilsMyUtilDebugIterator, error) {

	logs, sub, err := _ParaswapUtils.contract.FilterLogs(opts, "MyUtilDebug")
	if err != nil {
		return nil, err
	}
	return &ParaswapUtilsMyUtilDebugIterator{contract: _ParaswapUtils.contract, event: "MyUtilDebug", logs: logs, sub: sub}, nil
}

// WatchMyUtilDebug is a free log subscription operation binding the contract event 0x6ae69222768d00c831213d6afd4cec422b4d4f89949133c7c4e26ba5e02a5fa2.
//
// Solidity: event MyUtilDebug(address src, address dst, uint256 amount)
func (_ParaswapUtils *ParaswapUtilsFilterer) WatchMyUtilDebug(opts *bind.WatchOpts, sink chan<- *ParaswapUtilsMyUtilDebug) (event.Subscription, error) {

	logs, sub, err := _ParaswapUtils.contract.WatchLogs(opts, "MyUtilDebug")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParaswapUtilsMyUtilDebug)
				if err := _ParaswapUtils.contract.UnpackLog(event, "MyUtilDebug", log); err != nil {
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

// ParseMyUtilDebug is a log parse operation binding the contract event 0x6ae69222768d00c831213d6afd4cec422b4d4f89949133c7c4e26ba5e02a5fa2.
//
// Solidity: event MyUtilDebug(address src, address dst, uint256 amount)
func (_ParaswapUtils *ParaswapUtilsFilterer) ParseMyUtilDebug(log types.Log) (*ParaswapUtilsMyUtilDebug, error) {
	event := new(ParaswapUtilsMyUtilDebug)
	if err := _ParaswapUtils.contract.UnpackLog(event, "MyUtilDebug", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
