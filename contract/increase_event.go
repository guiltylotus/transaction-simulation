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

// IncreaseEventMetaData contains all meta data concerning the IncreaseEvent contract.
var IncreaseEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"NewID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"OldID\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610205806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636d4ce63c1461003b578063e8927fbc14610059575b600080fd5b610043610063565b6040516100509190610114565b60405180910390f35b61006161006c565b005b60008054905090565b7fec5ec5474b58f382ce897fbea318f1a94104f0962132b46972742df337dd14874260005460405161009f92919061012f565b60405180910390a16000808154809291906100b990610187565b91905055507f5d03ce126a434a5db3b41c75b2ade4154cf6d904110c5a409ca175f611ebad08426000546040516100f192919061012f565b60405180910390a1565b6000819050919050565b61010e816100fb565b82525050565b60006020820190506101296000830184610105565b92915050565b60006040820190506101446000830185610105565b6101516020830184610105565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610192826100fb565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036101c4576101c3610158565b5b60018201905091905056fea264697066735822122081677d967f84140ff7682bae2c427a8f324af4dedb3dc7193ceb9e12b9f25df864736f6c634300080f0033",
}

// IncreaseEventABI is the input ABI used to generate the binding from.
// Deprecated: Use IncreaseEventMetaData.ABI instead.
var IncreaseEventABI = IncreaseEventMetaData.ABI

// IncreaseEventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IncreaseEventMetaData.Bin instead.
var IncreaseEventBin = IncreaseEventMetaData.Bin

// DeployIncreaseEvent deploys a new Ethereum contract, binding an instance of IncreaseEvent to it.
func DeployIncreaseEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IncreaseEvent, error) {
	parsed, err := IncreaseEventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IncreaseEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IncreaseEvent{IncreaseEventCaller: IncreaseEventCaller{contract: contract}, IncreaseEventTransactor: IncreaseEventTransactor{contract: contract}, IncreaseEventFilterer: IncreaseEventFilterer{contract: contract}}, nil
}

// IncreaseEvent is an auto generated Go binding around an Ethereum contract.
type IncreaseEvent struct {
	IncreaseEventCaller     // Read-only binding to the contract
	IncreaseEventTransactor // Write-only binding to the contract
	IncreaseEventFilterer   // Log filterer for contract events
}

// IncreaseEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncreaseEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncreaseEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncreaseEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncreaseEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncreaseEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncreaseEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncreaseEventSession struct {
	Contract     *IncreaseEvent    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IncreaseEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncreaseEventCallerSession struct {
	Contract *IncreaseEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IncreaseEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncreaseEventTransactorSession struct {
	Contract     *IncreaseEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IncreaseEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncreaseEventRaw struct {
	Contract *IncreaseEvent // Generic contract binding to access the raw methods on
}

// IncreaseEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncreaseEventCallerRaw struct {
	Contract *IncreaseEventCaller // Generic read-only contract binding to access the raw methods on
}

// IncreaseEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncreaseEventTransactorRaw struct {
	Contract *IncreaseEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncreaseEvent creates a new instance of IncreaseEvent, bound to a specific deployed contract.
func NewIncreaseEvent(address common.Address, backend bind.ContractBackend) (*IncreaseEvent, error) {
	contract, err := bindIncreaseEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IncreaseEvent{IncreaseEventCaller: IncreaseEventCaller{contract: contract}, IncreaseEventTransactor: IncreaseEventTransactor{contract: contract}, IncreaseEventFilterer: IncreaseEventFilterer{contract: contract}}, nil
}

// NewIncreaseEventCaller creates a new read-only instance of IncreaseEvent, bound to a specific deployed contract.
func NewIncreaseEventCaller(address common.Address, caller bind.ContractCaller) (*IncreaseEventCaller, error) {
	contract, err := bindIncreaseEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncreaseEventCaller{contract: contract}, nil
}

// NewIncreaseEventTransactor creates a new write-only instance of IncreaseEvent, bound to a specific deployed contract.
func NewIncreaseEventTransactor(address common.Address, transactor bind.ContractTransactor) (*IncreaseEventTransactor, error) {
	contract, err := bindIncreaseEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncreaseEventTransactor{contract: contract}, nil
}

// NewIncreaseEventFilterer creates a new log filterer instance of IncreaseEvent, bound to a specific deployed contract.
func NewIncreaseEventFilterer(address common.Address, filterer bind.ContractFilterer) (*IncreaseEventFilterer, error) {
	contract, err := bindIncreaseEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncreaseEventFilterer{contract: contract}, nil
}

// bindIncreaseEvent binds a generic wrapper to an already deployed contract.
func bindIncreaseEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncreaseEventABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncreaseEvent *IncreaseEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IncreaseEvent.Contract.IncreaseEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncreaseEvent *IncreaseEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncreaseEvent.Contract.IncreaseEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncreaseEvent *IncreaseEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncreaseEvent.Contract.IncreaseEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncreaseEvent *IncreaseEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IncreaseEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncreaseEvent *IncreaseEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncreaseEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncreaseEvent *IncreaseEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncreaseEvent.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_IncreaseEvent *IncreaseEventCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IncreaseEvent.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_IncreaseEvent *IncreaseEventSession) Get() (*big.Int, error) {
	return _IncreaseEvent.Contract.Get(&_IncreaseEvent.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_IncreaseEvent *IncreaseEventCallerSession) Get() (*big.Int, error) {
	return _IncreaseEvent.Contract.Get(&_IncreaseEvent.CallOpts)
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() returns()
func (_IncreaseEvent *IncreaseEventTransactor) Increase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncreaseEvent.contract.Transact(opts, "increase")
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() returns()
func (_IncreaseEvent *IncreaseEventSession) Increase() (*types.Transaction, error) {
	return _IncreaseEvent.Contract.Increase(&_IncreaseEvent.TransactOpts)
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() returns()
func (_IncreaseEvent *IncreaseEventTransactorSession) Increase() (*types.Transaction, error) {
	return _IncreaseEvent.Contract.Increase(&_IncreaseEvent.TransactOpts)
}

// IncreaseEventNewIDIterator is returned from FilterNewID and is used to iterate over the raw logs and unpacked data for NewID events raised by the IncreaseEvent contract.
type IncreaseEventNewIDIterator struct {
	Event *IncreaseEventNewID // Event containing the contract specifics and raw log

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
func (it *IncreaseEventNewIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncreaseEventNewID)
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
		it.Event = new(IncreaseEventNewID)
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
func (it *IncreaseEventNewIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncreaseEventNewIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncreaseEventNewID represents a NewID event raised by the IncreaseEvent contract.
type IncreaseEventNewID struct {
	Date *big.Int
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewID is a free log retrieval operation binding the contract event 0x5d03ce126a434a5db3b41c75b2ade4154cf6d904110c5a409ca175f611ebad08.
//
// Solidity: event NewID(uint256 date, uint256 id)
func (_IncreaseEvent *IncreaseEventFilterer) FilterNewID(opts *bind.FilterOpts) (*IncreaseEventNewIDIterator, error) {

	logs, sub, err := _IncreaseEvent.contract.FilterLogs(opts, "NewID")
	if err != nil {
		return nil, err
	}
	return &IncreaseEventNewIDIterator{contract: _IncreaseEvent.contract, event: "NewID", logs: logs, sub: sub}, nil
}

// WatchNewID is a free log subscription operation binding the contract event 0x5d03ce126a434a5db3b41c75b2ade4154cf6d904110c5a409ca175f611ebad08.
//
// Solidity: event NewID(uint256 date, uint256 id)
func (_IncreaseEvent *IncreaseEventFilterer) WatchNewID(opts *bind.WatchOpts, sink chan<- *IncreaseEventNewID) (event.Subscription, error) {

	logs, sub, err := _IncreaseEvent.contract.WatchLogs(opts, "NewID")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncreaseEventNewID)
				if err := _IncreaseEvent.contract.UnpackLog(event, "NewID", log); err != nil {
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

// ParseNewID is a log parse operation binding the contract event 0x5d03ce126a434a5db3b41c75b2ade4154cf6d904110c5a409ca175f611ebad08.
//
// Solidity: event NewID(uint256 date, uint256 id)
func (_IncreaseEvent *IncreaseEventFilterer) ParseNewID(log types.Log) (*IncreaseEventNewID, error) {
	event := new(IncreaseEventNewID)
	if err := _IncreaseEvent.contract.UnpackLog(event, "NewID", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncreaseEventOldIDIterator is returned from FilterOldID and is used to iterate over the raw logs and unpacked data for OldID events raised by the IncreaseEvent contract.
type IncreaseEventOldIDIterator struct {
	Event *IncreaseEventOldID // Event containing the contract specifics and raw log

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
func (it *IncreaseEventOldIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncreaseEventOldID)
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
		it.Event = new(IncreaseEventOldID)
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
func (it *IncreaseEventOldIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncreaseEventOldIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncreaseEventOldID represents a OldID event raised by the IncreaseEvent contract.
type IncreaseEventOldID struct {
	Date *big.Int
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOldID is a free log retrieval operation binding the contract event 0xec5ec5474b58f382ce897fbea318f1a94104f0962132b46972742df337dd1487.
//
// Solidity: event OldID(uint256 date, uint256 id)
func (_IncreaseEvent *IncreaseEventFilterer) FilterOldID(opts *bind.FilterOpts) (*IncreaseEventOldIDIterator, error) {

	logs, sub, err := _IncreaseEvent.contract.FilterLogs(opts, "OldID")
	if err != nil {
		return nil, err
	}
	return &IncreaseEventOldIDIterator{contract: _IncreaseEvent.contract, event: "OldID", logs: logs, sub: sub}, nil
}

// WatchOldID is a free log subscription operation binding the contract event 0xec5ec5474b58f382ce897fbea318f1a94104f0962132b46972742df337dd1487.
//
// Solidity: event OldID(uint256 date, uint256 id)
func (_IncreaseEvent *IncreaseEventFilterer) WatchOldID(opts *bind.WatchOpts, sink chan<- *IncreaseEventOldID) (event.Subscription, error) {

	logs, sub, err := _IncreaseEvent.contract.WatchLogs(opts, "OldID")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncreaseEventOldID)
				if err := _IncreaseEvent.contract.UnpackLog(event, "OldID", log); err != nil {
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

// ParseOldID is a log parse operation binding the contract event 0xec5ec5474b58f382ce897fbea318f1a94104f0962132b46972742df337dd1487.
//
// Solidity: event OldID(uint256 date, uint256 id)
func (_IncreaseEvent *IncreaseEventFilterer) ParseOldID(log types.Log) (*IncreaseEventOldID, error) {
	event := new(IncreaseEventOldID)
	if err := _IncreaseEvent.contract.UnpackLog(event, "OldID", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
