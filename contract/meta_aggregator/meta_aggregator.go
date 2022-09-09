// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package metaaggregator

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

// MetaAggregationRouterSwapDescription is an auto generated low-level Go binding around an user-defined struct.
type MetaAggregationRouterSwapDescription struct {
	SrcToken        common.Address
	DstToken        common.Address
	SrcReceivers    []common.Address
	SrcAmounts      []*big.Int
	DstReceiver     common.Address
	Amount          *big.Int
	MinReturnAmount *big.Int
	Flags           *big.Int
	Permit          []byte
}

// SwapDescriptionExecutor1Inch is an auto generated low-level Go binding around an user-defined struct.
type SwapDescriptionExecutor1Inch struct {
	SrcToken         common.Address
	DstToken         common.Address
	SrcReceiver1Inch common.Address
	DstReceiver      common.Address
	SrcReceivers     []common.Address
	SrcAmounts       []*big.Int
	Amount           *big.Int
	MinReturnAmount  *big.Int
	Flags            *big.Int
	Permit           []byte
}

// MetaAggregatorMetaData contains all meta data concerning the MetaAggregator contract.
var MetaAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"ClientData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Error\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"}],\"name\":\"Exchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"executorData\",\"type\":\"bytes\"}],\"name\":\"MyCallBytesDebug\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"executorData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor1Inch\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"srcReceiver1Inch\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structSwapDescriptionExecutor1Inch\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"executor1InchData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swapExecutor1Inch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router1Inch\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"router1InchData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swapRouter1Inch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"executorData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swapSimpleMode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MetaAggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use MetaAggregatorMetaData.ABI instead.
var MetaAggregatorABI = MetaAggregatorMetaData.ABI

// MetaAggregator is an auto generated Go binding around an Ethereum contract.
type MetaAggregator struct {
	MetaAggregatorCaller     // Read-only binding to the contract
	MetaAggregatorTransactor // Write-only binding to the contract
	MetaAggregatorFilterer   // Log filterer for contract events
}

// MetaAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetaAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaAggregatorSession struct {
	Contract     *MetaAggregator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaAggregatorCallerSession struct {
	Contract *MetaAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MetaAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaAggregatorTransactorSession struct {
	Contract     *MetaAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MetaAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaAggregatorRaw struct {
	Contract *MetaAggregator // Generic contract binding to access the raw methods on
}

// MetaAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaAggregatorCallerRaw struct {
	Contract *MetaAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// MetaAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaAggregatorTransactorRaw struct {
	Contract *MetaAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaAggregator creates a new instance of MetaAggregator, bound to a specific deployed contract.
func NewMetaAggregator(address common.Address, backend bind.ContractBackend) (*MetaAggregator, error) {
	contract, err := bindMetaAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaAggregator{MetaAggregatorCaller: MetaAggregatorCaller{contract: contract}, MetaAggregatorTransactor: MetaAggregatorTransactor{contract: contract}, MetaAggregatorFilterer: MetaAggregatorFilterer{contract: contract}}, nil
}

// NewMetaAggregatorCaller creates a new read-only instance of MetaAggregator, bound to a specific deployed contract.
func NewMetaAggregatorCaller(address common.Address, caller bind.ContractCaller) (*MetaAggregatorCaller, error) {
	contract, err := bindMetaAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaAggregatorCaller{contract: contract}, nil
}

// NewMetaAggregatorTransactor creates a new write-only instance of MetaAggregator, bound to a specific deployed contract.
func NewMetaAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaAggregatorTransactor, error) {
	contract, err := bindMetaAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaAggregatorTransactor{contract: contract}, nil
}

// NewMetaAggregatorFilterer creates a new log filterer instance of MetaAggregator, bound to a specific deployed contract.
func NewMetaAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*MetaAggregatorFilterer, error) {
	contract, err := bindMetaAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaAggregatorFilterer{contract: contract}, nil
}

// bindMetaAggregator binds a generic wrapper to an already deployed contract.
func bindMetaAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaAggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaAggregator *MetaAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaAggregator.Contract.MetaAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaAggregator *MetaAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaAggregator.Contract.MetaAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaAggregator *MetaAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaAggregator.Contract.MetaAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaAggregator *MetaAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaAggregator *MetaAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaAggregator *MetaAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaAggregator.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_MetaAggregator *MetaAggregatorCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaAggregator.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_MetaAggregator *MetaAggregatorSession) WETH() (common.Address, error) {
	return _MetaAggregator.Contract.WETH(&_MetaAggregator.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_MetaAggregator *MetaAggregatorCallerSession) WETH() (common.Address, error) {
	return _MetaAggregator.Contract.WETH(&_MetaAggregator.CallOpts)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_MetaAggregator *MetaAggregatorCaller) IsWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MetaAggregator.contract.Call(opts, &out, "isWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_MetaAggregator *MetaAggregatorSession) IsWhitelist(arg0 common.Address) (bool, error) {
	return _MetaAggregator.Contract.IsWhitelist(&_MetaAggregator.CallOpts, arg0)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_MetaAggregator *MetaAggregatorCallerSession) IsWhitelist(arg0 common.Address) (bool, error) {
	return _MetaAggregator.Contract.IsWhitelist(&_MetaAggregator.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaAggregator *MetaAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaAggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaAggregator *MetaAggregatorSession) Owner() (common.Address, error) {
	return _MetaAggregator.Contract.Owner(&_MetaAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaAggregator *MetaAggregatorCallerSession) Owner() (common.Address, error) {
	return _MetaAggregator.Contract.Owner(&_MetaAggregator.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetaAggregator *MetaAggregatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaAggregator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetaAggregator *MetaAggregatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MetaAggregator.Contract.RenounceOwnership(&_MetaAggregator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetaAggregator *MetaAggregatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MetaAggregator.Contract.RenounceOwnership(&_MetaAggregator.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_MetaAggregator *MetaAggregatorTransactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaAggregator.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_MetaAggregator *MetaAggregatorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaAggregator.Contract.RescueFunds(&_MetaAggregator.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_MetaAggregator *MetaAggregatorTransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaAggregator.Contract.RescueFunds(&_MetaAggregator.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xabcffc26.
//
// Solidity: function swap(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_MetaAggregator *MetaAggregatorTransactor) Swap(opts *bind.TransactOpts, caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _MetaAggregator.contract.Transact(opts, "swap", caller, desc, executorData, clientData)
}

// Swap is a paid mutator transaction binding the contract method 0xabcffc26.
//
// Solidity: function swap(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_MetaAggregator *MetaAggregatorSession) Swap(caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _MetaAggregator.Contract.Swap(&_MetaAggregator.TransactOpts, caller, desc, executorData, clientData)
}

// Swap is a paid mutator transaction binding the contract method 0xabcffc26.
//
// Solidity: function swap(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_MetaAggregator *MetaAggregatorTransactorSession) Swap(caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _MetaAggregator.Contract.Swap(&_MetaAggregator.TransactOpts, caller, desc, executorData, clientData)
}

// SwapExecutor1Inch is a paid mutator transaction binding the contract method 0x1fb3e235.
//
// Solidity: function swapExecutor1Inch(address caller, (address,address,address,address,address[],uint256[],uint256,uint256,uint256,bytes) desc, bytes executor1InchData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_MetaAggregator *MetaAggregatorTransactor) SwapExecutor1Inch(opts *bind.TransactOpts, caller common.Address, desc SwapDescriptionExecutor1Inch, executor1InchData []byte, clientData []byte) (*types.Transaction, error) {
	return _MetaAggregator.contract.Transact(opts, "swapExecutor1Inch", caller, desc, executor1InchData, clientData)
}

// SwapExecutor1Inch is a paid mutator transaction binding the contract method 0x1fb3e235.
//
// Solidity: function swapExecutor1Inch(address caller, (address,address,address,address,address[],uint256[],uint256,uint256,uint256,bytes) desc, bytes executor1InchData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_MetaAggregator *MetaAggregatorSession) SwapExecutor1Inch(caller common.Address, desc SwapDescriptionExecutor1Inch, executor1InchData []byte, clientData []byte) (*types.Transaction, error) {
	return _MetaAggregator.Contract.SwapExecutor1Inch(&_MetaAggregator.TransactOpts, caller, desc, executor1InchData, clientData)
}

// SwapExecutor1Inch is a paid mutator transaction binding the contract method 0x1fb3e235.
//
// Solidity: function swapExecutor1Inch(address caller, (address,address,address,address,address[],uint256[],uint256,uint256,uint256,bytes) desc, bytes executor1InchData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_MetaAggregator *MetaAggregatorTransactorSession) SwapExecutor1Inch(caller common.Address, desc SwapDescriptionExecutor1Inch, executor1InchData []byte, clientData []byte) (*types.Transaction, error) {
	return _MetaAggregator.Contract.SwapExecutor1Inch(&_MetaAggregator.TransactOpts, caller, desc, executor1InchData, clientData)
}

// SwapRouter1Inch is a paid mutator transaction binding the contract method 0xa72bd2f5.
//
// Solidity: function swapRouter1Inch(address router1Inch, bytes router1InchData, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_MetaAggregator *MetaAggregatorTransactor) SwapRouter1Inch(opts *bind.TransactOpts, router1Inch common.Address, router1InchData []byte, desc MetaAggregationRouterSwapDescription, clientData []byte) (*types.Transaction, error) {
	return _MetaAggregator.contract.Transact(opts, "swapRouter1Inch", router1Inch, router1InchData, desc, clientData)
}

// SwapRouter1Inch is a paid mutator transaction binding the contract method 0xa72bd2f5.
//
// Solidity: function swapRouter1Inch(address router1Inch, bytes router1InchData, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_MetaAggregator *MetaAggregatorSession) SwapRouter1Inch(router1Inch common.Address, router1InchData []byte, desc MetaAggregationRouterSwapDescription, clientData []byte) (*types.Transaction, error) {
	return _MetaAggregator.Contract.SwapRouter1Inch(&_MetaAggregator.TransactOpts, router1Inch, router1InchData, desc, clientData)
}

// SwapRouter1Inch is a paid mutator transaction binding the contract method 0xa72bd2f5.
//
// Solidity: function swapRouter1Inch(address router1Inch, bytes router1InchData, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_MetaAggregator *MetaAggregatorTransactorSession) SwapRouter1Inch(router1Inch common.Address, router1InchData []byte, desc MetaAggregationRouterSwapDescription, clientData []byte) (*types.Transaction, error) {
	return _MetaAggregator.Contract.SwapRouter1Inch(&_MetaAggregator.TransactOpts, router1Inch, router1InchData, desc, clientData)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0xa7f5c104.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_MetaAggregator *MetaAggregatorTransactor) SwapSimpleMode(opts *bind.TransactOpts, caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _MetaAggregator.contract.Transact(opts, "swapSimpleMode", caller, desc, executorData, clientData)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0xa7f5c104.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_MetaAggregator *MetaAggregatorSession) SwapSimpleMode(caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _MetaAggregator.Contract.SwapSimpleMode(&_MetaAggregator.TransactOpts, caller, desc, executorData, clientData)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0xa7f5c104.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_MetaAggregator *MetaAggregatorTransactorSession) SwapSimpleMode(caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _MetaAggregator.Contract.SwapSimpleMode(&_MetaAggregator.TransactOpts, caller, desc, executorData, clientData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetaAggregator *MetaAggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MetaAggregator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetaAggregator *MetaAggregatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MetaAggregator.Contract.TransferOwnership(&_MetaAggregator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetaAggregator *MetaAggregatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MetaAggregator.Contract.TransferOwnership(&_MetaAggregator.TransactOpts, newOwner)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x0d392cd9.
//
// Solidity: function updateWhitelist(address addr, bool value) returns()
func (_MetaAggregator *MetaAggregatorTransactor) UpdateWhitelist(opts *bind.TransactOpts, addr common.Address, value bool) (*types.Transaction, error) {
	return _MetaAggregator.contract.Transact(opts, "updateWhitelist", addr, value)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x0d392cd9.
//
// Solidity: function updateWhitelist(address addr, bool value) returns()
func (_MetaAggregator *MetaAggregatorSession) UpdateWhitelist(addr common.Address, value bool) (*types.Transaction, error) {
	return _MetaAggregator.Contract.UpdateWhitelist(&_MetaAggregator.TransactOpts, addr, value)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x0d392cd9.
//
// Solidity: function updateWhitelist(address addr, bool value) returns()
func (_MetaAggregator *MetaAggregatorTransactorSession) UpdateWhitelist(addr common.Address, value bool) (*types.Transaction, error) {
	return _MetaAggregator.Contract.UpdateWhitelist(&_MetaAggregator.TransactOpts, addr, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MetaAggregator *MetaAggregatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaAggregator.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MetaAggregator *MetaAggregatorSession) Receive() (*types.Transaction, error) {
	return _MetaAggregator.Contract.Receive(&_MetaAggregator.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MetaAggregator *MetaAggregatorTransactorSession) Receive() (*types.Transaction, error) {
	return _MetaAggregator.Contract.Receive(&_MetaAggregator.TransactOpts)
}

// MetaAggregatorClientDataIterator is returned from FilterClientData and is used to iterate over the raw logs and unpacked data for ClientData events raised by the MetaAggregator contract.
type MetaAggregatorClientDataIterator struct {
	Event *MetaAggregatorClientData // Event containing the contract specifics and raw log

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
func (it *MetaAggregatorClientDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaAggregatorClientData)
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
		it.Event = new(MetaAggregatorClientData)
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
func (it *MetaAggregatorClientDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaAggregatorClientDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaAggregatorClientData represents a ClientData event raised by the MetaAggregator contract.
type MetaAggregatorClientData struct {
	ClientData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClientData is a free log retrieval operation binding the contract event 0x095e66fa4dd6a6f7b43fb8444a7bd0edb870508c7abf639bc216efb0bcff9779.
//
// Solidity: event ClientData(bytes clientData)
func (_MetaAggregator *MetaAggregatorFilterer) FilterClientData(opts *bind.FilterOpts) (*MetaAggregatorClientDataIterator, error) {

	logs, sub, err := _MetaAggregator.contract.FilterLogs(opts, "ClientData")
	if err != nil {
		return nil, err
	}
	return &MetaAggregatorClientDataIterator{contract: _MetaAggregator.contract, event: "ClientData", logs: logs, sub: sub}, nil
}

// WatchClientData is a free log subscription operation binding the contract event 0x095e66fa4dd6a6f7b43fb8444a7bd0edb870508c7abf639bc216efb0bcff9779.
//
// Solidity: event ClientData(bytes clientData)
func (_MetaAggregator *MetaAggregatorFilterer) WatchClientData(opts *bind.WatchOpts, sink chan<- *MetaAggregatorClientData) (event.Subscription, error) {

	logs, sub, err := _MetaAggregator.contract.WatchLogs(opts, "ClientData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaAggregatorClientData)
				if err := _MetaAggregator.contract.UnpackLog(event, "ClientData", log); err != nil {
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

// ParseClientData is a log parse operation binding the contract event 0x095e66fa4dd6a6f7b43fb8444a7bd0edb870508c7abf639bc216efb0bcff9779.
//
// Solidity: event ClientData(bytes clientData)
func (_MetaAggregator *MetaAggregatorFilterer) ParseClientData(log types.Log) (*MetaAggregatorClientData, error) {
	event := new(MetaAggregatorClientData)
	if err := _MetaAggregator.contract.UnpackLog(event, "ClientData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaAggregatorErrorIterator is returned from FilterError and is used to iterate over the raw logs and unpacked data for Error events raised by the MetaAggregator contract.
type MetaAggregatorErrorIterator struct {
	Event *MetaAggregatorError // Event containing the contract specifics and raw log

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
func (it *MetaAggregatorErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaAggregatorError)
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
		it.Event = new(MetaAggregatorError)
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
func (it *MetaAggregatorErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaAggregatorErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaAggregatorError represents a Error event raised by the MetaAggregator contract.
type MetaAggregatorError struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterError is a free log retrieval operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_MetaAggregator *MetaAggregatorFilterer) FilterError(opts *bind.FilterOpts) (*MetaAggregatorErrorIterator, error) {

	logs, sub, err := _MetaAggregator.contract.FilterLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return &MetaAggregatorErrorIterator{contract: _MetaAggregator.contract, event: "Error", logs: logs, sub: sub}, nil
}

// WatchError is a free log subscription operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_MetaAggregator *MetaAggregatorFilterer) WatchError(opts *bind.WatchOpts, sink chan<- *MetaAggregatorError) (event.Subscription, error) {

	logs, sub, err := _MetaAggregator.contract.WatchLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaAggregatorError)
				if err := _MetaAggregator.contract.UnpackLog(event, "Error", log); err != nil {
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

// ParseError is a log parse operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_MetaAggregator *MetaAggregatorFilterer) ParseError(log types.Log) (*MetaAggregatorError, error) {
	event := new(MetaAggregatorError)
	if err := _MetaAggregator.contract.UnpackLog(event, "Error", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaAggregatorExchangeIterator is returned from FilterExchange and is used to iterate over the raw logs and unpacked data for Exchange events raised by the MetaAggregator contract.
type MetaAggregatorExchangeIterator struct {
	Event *MetaAggregatorExchange // Event containing the contract specifics and raw log

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
func (it *MetaAggregatorExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaAggregatorExchange)
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
		it.Event = new(MetaAggregatorExchange)
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
func (it *MetaAggregatorExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaAggregatorExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaAggregatorExchange represents a Exchange event raised by the MetaAggregator contract.
type MetaAggregatorExchange struct {
	Pair      common.Address
	AmountOut *big.Int
	Output    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExchange is a free log retrieval operation binding the contract event 0xddac40937f35385a34f721af292e5a83fc5b840f722bff57c2fc71adba708c48.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (_MetaAggregator *MetaAggregatorFilterer) FilterExchange(opts *bind.FilterOpts) (*MetaAggregatorExchangeIterator, error) {

	logs, sub, err := _MetaAggregator.contract.FilterLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return &MetaAggregatorExchangeIterator{contract: _MetaAggregator.contract, event: "Exchange", logs: logs, sub: sub}, nil
}

// WatchExchange is a free log subscription operation binding the contract event 0xddac40937f35385a34f721af292e5a83fc5b840f722bff57c2fc71adba708c48.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (_MetaAggregator *MetaAggregatorFilterer) WatchExchange(opts *bind.WatchOpts, sink chan<- *MetaAggregatorExchange) (event.Subscription, error) {

	logs, sub, err := _MetaAggregator.contract.WatchLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaAggregatorExchange)
				if err := _MetaAggregator.contract.UnpackLog(event, "Exchange", log); err != nil {
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

// ParseExchange is a log parse operation binding the contract event 0xddac40937f35385a34f721af292e5a83fc5b840f722bff57c2fc71adba708c48.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (_MetaAggregator *MetaAggregatorFilterer) ParseExchange(log types.Log) (*MetaAggregatorExchange, error) {
	event := new(MetaAggregatorExchange)
	if err := _MetaAggregator.contract.UnpackLog(event, "Exchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaAggregatorMyCallBytesDebugIterator is returned from FilterMyCallBytesDebug and is used to iterate over the raw logs and unpacked data for MyCallBytesDebug events raised by the MetaAggregator contract.
type MetaAggregatorMyCallBytesDebugIterator struct {
	Event *MetaAggregatorMyCallBytesDebug // Event containing the contract specifics and raw log

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
func (it *MetaAggregatorMyCallBytesDebugIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaAggregatorMyCallBytesDebug)
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
		it.Event = new(MetaAggregatorMyCallBytesDebug)
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
func (it *MetaAggregatorMyCallBytesDebugIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaAggregatorMyCallBytesDebugIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaAggregatorMyCallBytesDebug represents a MyCallBytesDebug event raised by the MetaAggregator contract.
type MetaAggregatorMyCallBytesDebug struct {
	Caller       common.Address
	ExecutorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMyCallBytesDebug is a free log retrieval operation binding the contract event 0x394ae19bc3a97247de8a034ad95a30a8621c3dcdf2b9d6f06e57ddccdcdcb01a.
//
// Solidity: event MyCallBytesDebug(address caller, bytes executorData)
func (_MetaAggregator *MetaAggregatorFilterer) FilterMyCallBytesDebug(opts *bind.FilterOpts) (*MetaAggregatorMyCallBytesDebugIterator, error) {

	logs, sub, err := _MetaAggregator.contract.FilterLogs(opts, "MyCallBytesDebug")
	if err != nil {
		return nil, err
	}
	return &MetaAggregatorMyCallBytesDebugIterator{contract: _MetaAggregator.contract, event: "MyCallBytesDebug", logs: logs, sub: sub}, nil
}

// WatchMyCallBytesDebug is a free log subscription operation binding the contract event 0x394ae19bc3a97247de8a034ad95a30a8621c3dcdf2b9d6f06e57ddccdcdcb01a.
//
// Solidity: event MyCallBytesDebug(address caller, bytes executorData)
func (_MetaAggregator *MetaAggregatorFilterer) WatchMyCallBytesDebug(opts *bind.WatchOpts, sink chan<- *MetaAggregatorMyCallBytesDebug) (event.Subscription, error) {

	logs, sub, err := _MetaAggregator.contract.WatchLogs(opts, "MyCallBytesDebug")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaAggregatorMyCallBytesDebug)
				if err := _MetaAggregator.contract.UnpackLog(event, "MyCallBytesDebug", log); err != nil {
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

// ParseMyCallBytesDebug is a log parse operation binding the contract event 0x394ae19bc3a97247de8a034ad95a30a8621c3dcdf2b9d6f06e57ddccdcdcb01a.
//
// Solidity: event MyCallBytesDebug(address caller, bytes executorData)
func (_MetaAggregator *MetaAggregatorFilterer) ParseMyCallBytesDebug(log types.Log) (*MetaAggregatorMyCallBytesDebug, error) {
	event := new(MetaAggregatorMyCallBytesDebug)
	if err := _MetaAggregator.contract.UnpackLog(event, "MyCallBytesDebug", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaAggregatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MetaAggregator contract.
type MetaAggregatorOwnershipTransferredIterator struct {
	Event *MetaAggregatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MetaAggregatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaAggregatorOwnershipTransferred)
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
		it.Event = new(MetaAggregatorOwnershipTransferred)
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
func (it *MetaAggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaAggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaAggregatorOwnershipTransferred represents a OwnershipTransferred event raised by the MetaAggregator contract.
type MetaAggregatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetaAggregator *MetaAggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MetaAggregatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetaAggregator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MetaAggregatorOwnershipTransferredIterator{contract: _MetaAggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetaAggregator *MetaAggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MetaAggregatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetaAggregator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaAggregatorOwnershipTransferred)
				if err := _MetaAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MetaAggregator *MetaAggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*MetaAggregatorOwnershipTransferred, error) {
	event := new(MetaAggregatorOwnershipTransferred)
	if err := _MetaAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaAggregatorSwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the MetaAggregator contract.
type MetaAggregatorSwappedIterator struct {
	Event *MetaAggregatorSwapped // Event containing the contract specifics and raw log

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
func (it *MetaAggregatorSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaAggregatorSwapped)
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
		it.Event = new(MetaAggregatorSwapped)
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
func (it *MetaAggregatorSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaAggregatorSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaAggregatorSwapped represents a Swapped event raised by the MetaAggregator contract.
type MetaAggregatorSwapped struct {
	Sender       common.Address
	SrcToken     common.Address
	DstToken     common.Address
	DstReceiver  common.Address
	SpentAmount  *big.Int
	ReturnAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_MetaAggregator *MetaAggregatorFilterer) FilterSwapped(opts *bind.FilterOpts) (*MetaAggregatorSwappedIterator, error) {

	logs, sub, err := _MetaAggregator.contract.FilterLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return &MetaAggregatorSwappedIterator{contract: _MetaAggregator.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_MetaAggregator *MetaAggregatorFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *MetaAggregatorSwapped) (event.Subscription, error) {

	logs, sub, err := _MetaAggregator.contract.WatchLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaAggregatorSwapped)
				if err := _MetaAggregator.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_MetaAggregator *MetaAggregatorFilterer) ParseSwapped(log types.Log) (*MetaAggregatorSwapped, error) {
	event := new(MetaAggregatorSwapped)
	if err := _MetaAggregator.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
