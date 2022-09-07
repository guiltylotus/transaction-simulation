// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paraswap

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

// UtilsAdapter is an auto generated low-level Go binding around an user-defined struct.
type UtilsAdapter struct {
	Adapter    common.Address
	Percent    *big.Int
	NetworkFee *big.Int
	Route      []UtilsRoute
}

// UtilsBuyData is an auto generated low-level Go binding around an user-defined struct.
type UtilsBuyData struct {
	Adapter        common.Address
	FromToken      common.Address
	ToToken        common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	Beneficiary    common.Address
	Route          []UtilsRoute
	Partner        common.Address
	FeePercent     *big.Int
	Permit         []byte
	Deadline       *big.Int
	Uuid           [16]byte
}

// UtilsMegaSwapPath is an auto generated low-level Go binding around an user-defined struct.
type UtilsMegaSwapPath struct {
	FromAmountPercent *big.Int
	Path              []UtilsPath
}

// UtilsMegaSwapSellData is an auto generated low-level Go binding around an user-defined struct.
type UtilsMegaSwapSellData struct {
	FromToken      common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	Beneficiary    common.Address
	Path           []UtilsMegaSwapPath
	Partner        common.Address
	FeePercent     *big.Int
	Permit         []byte
	Deadline       *big.Int
	Uuid           [16]byte
}

// UtilsPath is an auto generated low-level Go binding around an user-defined struct.
type UtilsPath struct {
	To              common.Address
	TotalNetworkFee *big.Int
	Adapters        []UtilsAdapter
}

// UtilsRoute is an auto generated low-level Go binding around an user-defined struct.
type UtilsRoute struct {
	Index          *big.Int
	TargetExchange common.Address
	Percent        *big.Int
	Payload        []byte
	NetworkFee     *big.Int
}

// UtilsSellData is an auto generated low-level Go binding around an user-defined struct.
type UtilsSellData struct {
	FromToken      common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	Beneficiary    common.Address
	Path           []UtilsPath
	Partner        common.Address
	FeePercent     *big.Int
	Permit         []byte
	Deadline       *big.Int
	Uuid           [16]byte
}

// ParaswapRouterMetaData contains all meta data concerning the ParaswapRouter contract.
var ParaswapRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_partnerSharePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxFeePercent\",\"type\":\"uint256\"},{\"internalType\":\"contractIFeeClaimer\",\"name\":\"_feeClaimer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"}],\"name\":\"BoughtV3\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MyDebug\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"}],\"name\":\"SwappedV3\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ROUTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WHITELISTED_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"targetExchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"}],\"internalType\":\"structUtils.Route[]\",\"name\":\"route\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.BuyData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeClaimer\",\"outputs\":[{\"internalType\":\"contractIFeeClaimer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxFeePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromAmountPercent\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalNetworkFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"adapter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"targetExchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"}],\"internalType\":\"structUtils.Route[]\",\"name\":\"route\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.Adapter[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.Path[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.MegaSwapPath[]\",\"name\":\"path\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.MegaSwapSellData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"megaSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalNetworkFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"adapter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"targetExchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"}],\"internalType\":\"structUtils.Route[]\",\"name\":\"route\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.Adapter[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.Path[]\",\"name\":\"path\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.SellData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"multiSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partnerSharePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ParaswapRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use ParaswapRouterMetaData.ABI instead.
var ParaswapRouterABI = ParaswapRouterMetaData.ABI

// ParaswapRouter is an auto generated Go binding around an Ethereum contract.
type ParaswapRouter struct {
	ParaswapRouterCaller     // Read-only binding to the contract
	ParaswapRouterTransactor // Write-only binding to the contract
	ParaswapRouterFilterer   // Log filterer for contract events
}

// ParaswapRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParaswapRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParaswapRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParaswapRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParaswapRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParaswapRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParaswapRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParaswapRouterSession struct {
	Contract     *ParaswapRouter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParaswapRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParaswapRouterCallerSession struct {
	Contract *ParaswapRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ParaswapRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParaswapRouterTransactorSession struct {
	Contract     *ParaswapRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ParaswapRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParaswapRouterRaw struct {
	Contract *ParaswapRouter // Generic contract binding to access the raw methods on
}

// ParaswapRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParaswapRouterCallerRaw struct {
	Contract *ParaswapRouterCaller // Generic read-only contract binding to access the raw methods on
}

// ParaswapRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParaswapRouterTransactorRaw struct {
	Contract *ParaswapRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParaswapRouter creates a new instance of ParaswapRouter, bound to a specific deployed contract.
func NewParaswapRouter(address common.Address, backend bind.ContractBackend) (*ParaswapRouter, error) {
	contract, err := bindParaswapRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ParaswapRouter{ParaswapRouterCaller: ParaswapRouterCaller{contract: contract}, ParaswapRouterTransactor: ParaswapRouterTransactor{contract: contract}, ParaswapRouterFilterer: ParaswapRouterFilterer{contract: contract}}, nil
}

// NewParaswapRouterCaller creates a new read-only instance of ParaswapRouter, bound to a specific deployed contract.
func NewParaswapRouterCaller(address common.Address, caller bind.ContractCaller) (*ParaswapRouterCaller, error) {
	contract, err := bindParaswapRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParaswapRouterCaller{contract: contract}, nil
}

// NewParaswapRouterTransactor creates a new write-only instance of ParaswapRouter, bound to a specific deployed contract.
func NewParaswapRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*ParaswapRouterTransactor, error) {
	contract, err := bindParaswapRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParaswapRouterTransactor{contract: contract}, nil
}

// NewParaswapRouterFilterer creates a new log filterer instance of ParaswapRouter, bound to a specific deployed contract.
func NewParaswapRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*ParaswapRouterFilterer, error) {
	contract, err := bindParaswapRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParaswapRouterFilterer{contract: contract}, nil
}

// bindParaswapRouter binds a generic wrapper to an already deployed contract.
func bindParaswapRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParaswapRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParaswapRouter *ParaswapRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParaswapRouter.Contract.ParaswapRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParaswapRouter *ParaswapRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParaswapRouter.Contract.ParaswapRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParaswapRouter *ParaswapRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParaswapRouter.Contract.ParaswapRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParaswapRouter *ParaswapRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParaswapRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParaswapRouter *ParaswapRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParaswapRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParaswapRouter *ParaswapRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParaswapRouter.Contract.contract.Transact(opts, method, params...)
}

// ROUTERROLE is a free data retrieval call binding the contract method 0x30d643b5.
//
// Solidity: function ROUTER_ROLE() view returns(bytes32)
func (_ParaswapRouter *ParaswapRouterCaller) ROUTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ParaswapRouter.contract.Call(opts, &out, "ROUTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROUTERROLE is a free data retrieval call binding the contract method 0x30d643b5.
//
// Solidity: function ROUTER_ROLE() view returns(bytes32)
func (_ParaswapRouter *ParaswapRouterSession) ROUTERROLE() ([32]byte, error) {
	return _ParaswapRouter.Contract.ROUTERROLE(&_ParaswapRouter.CallOpts)
}

// ROUTERROLE is a free data retrieval call binding the contract method 0x30d643b5.
//
// Solidity: function ROUTER_ROLE() view returns(bytes32)
func (_ParaswapRouter *ParaswapRouterCallerSession) ROUTERROLE() ([32]byte, error) {
	return _ParaswapRouter.Contract.ROUTERROLE(&_ParaswapRouter.CallOpts)
}

// WHITELISTEDROLE is a free data retrieval call binding the contract method 0x7a3226ec.
//
// Solidity: function WHITELISTED_ROLE() view returns(bytes32)
func (_ParaswapRouter *ParaswapRouterCaller) WHITELISTEDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ParaswapRouter.contract.Call(opts, &out, "WHITELISTED_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WHITELISTEDROLE is a free data retrieval call binding the contract method 0x7a3226ec.
//
// Solidity: function WHITELISTED_ROLE() view returns(bytes32)
func (_ParaswapRouter *ParaswapRouterSession) WHITELISTEDROLE() ([32]byte, error) {
	return _ParaswapRouter.Contract.WHITELISTEDROLE(&_ParaswapRouter.CallOpts)
}

// WHITELISTEDROLE is a free data retrieval call binding the contract method 0x7a3226ec.
//
// Solidity: function WHITELISTED_ROLE() view returns(bytes32)
func (_ParaswapRouter *ParaswapRouterCallerSession) WHITELISTEDROLE() ([32]byte, error) {
	return _ParaswapRouter.Contract.WHITELISTEDROLE(&_ParaswapRouter.CallOpts)
}

// FeeClaimer is a free data retrieval call binding the contract method 0x81cbd3ea.
//
// Solidity: function feeClaimer() view returns(address)
func (_ParaswapRouter *ParaswapRouterCaller) FeeClaimer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ParaswapRouter.contract.Call(opts, &out, "feeClaimer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeClaimer is a free data retrieval call binding the contract method 0x81cbd3ea.
//
// Solidity: function feeClaimer() view returns(address)
func (_ParaswapRouter *ParaswapRouterSession) FeeClaimer() (common.Address, error) {
	return _ParaswapRouter.Contract.FeeClaimer(&_ParaswapRouter.CallOpts)
}

// FeeClaimer is a free data retrieval call binding the contract method 0x81cbd3ea.
//
// Solidity: function feeClaimer() view returns(address)
func (_ParaswapRouter *ParaswapRouterCallerSession) FeeClaimer() (common.Address, error) {
	return _ParaswapRouter.Contract.FeeClaimer(&_ParaswapRouter.CallOpts)
}

// GetKey is a free data retrieval call binding the contract method 0x82678dd6.
//
// Solidity: function getKey() pure returns(bytes32)
func (_ParaswapRouter *ParaswapRouterCaller) GetKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ParaswapRouter.contract.Call(opts, &out, "getKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0x82678dd6.
//
// Solidity: function getKey() pure returns(bytes32)
func (_ParaswapRouter *ParaswapRouterSession) GetKey() ([32]byte, error) {
	return _ParaswapRouter.Contract.GetKey(&_ParaswapRouter.CallOpts)
}

// GetKey is a free data retrieval call binding the contract method 0x82678dd6.
//
// Solidity: function getKey() pure returns(bytes32)
func (_ParaswapRouter *ParaswapRouterCallerSession) GetKey() ([32]byte, error) {
	return _ParaswapRouter.Contract.GetKey(&_ParaswapRouter.CallOpts)
}

// MaxFeePercent is a free data retrieval call binding the contract method 0xd830a05b.
//
// Solidity: function maxFeePercent() view returns(uint256)
func (_ParaswapRouter *ParaswapRouterCaller) MaxFeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ParaswapRouter.contract.Call(opts, &out, "maxFeePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFeePercent is a free data retrieval call binding the contract method 0xd830a05b.
//
// Solidity: function maxFeePercent() view returns(uint256)
func (_ParaswapRouter *ParaswapRouterSession) MaxFeePercent() (*big.Int, error) {
	return _ParaswapRouter.Contract.MaxFeePercent(&_ParaswapRouter.CallOpts)
}

// MaxFeePercent is a free data retrieval call binding the contract method 0xd830a05b.
//
// Solidity: function maxFeePercent() view returns(uint256)
func (_ParaswapRouter *ParaswapRouterCallerSession) MaxFeePercent() (*big.Int, error) {
	return _ParaswapRouter.Contract.MaxFeePercent(&_ParaswapRouter.CallOpts)
}

// PartnerSharePercent is a free data retrieval call binding the contract method 0x12070a41.
//
// Solidity: function partnerSharePercent() view returns(uint256)
func (_ParaswapRouter *ParaswapRouterCaller) PartnerSharePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ParaswapRouter.contract.Call(opts, &out, "partnerSharePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PartnerSharePercent is a free data retrieval call binding the contract method 0x12070a41.
//
// Solidity: function partnerSharePercent() view returns(uint256)
func (_ParaswapRouter *ParaswapRouterSession) PartnerSharePercent() (*big.Int, error) {
	return _ParaswapRouter.Contract.PartnerSharePercent(&_ParaswapRouter.CallOpts)
}

// PartnerSharePercent is a free data retrieval call binding the contract method 0x12070a41.
//
// Solidity: function partnerSharePercent() view returns(uint256)
func (_ParaswapRouter *ParaswapRouterCallerSession) PartnerSharePercent() (*big.Int, error) {
	return _ParaswapRouter.Contract.PartnerSharePercent(&_ParaswapRouter.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0x35326910.
//
// Solidity: function buy((address,address,address,uint256,uint256,uint256,address,(uint256,address,uint256,bytes,uint256)[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_ParaswapRouter *ParaswapRouterTransactor) Buy(opts *bind.TransactOpts, data UtilsBuyData) (*types.Transaction, error) {
	return _ParaswapRouter.contract.Transact(opts, "buy", data)
}

// Buy is a paid mutator transaction binding the contract method 0x35326910.
//
// Solidity: function buy((address,address,address,uint256,uint256,uint256,address,(uint256,address,uint256,bytes,uint256)[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_ParaswapRouter *ParaswapRouterSession) Buy(data UtilsBuyData) (*types.Transaction, error) {
	return _ParaswapRouter.Contract.Buy(&_ParaswapRouter.TransactOpts, data)
}

// Buy is a paid mutator transaction binding the contract method 0x35326910.
//
// Solidity: function buy((address,address,address,uint256,uint256,uint256,address,(uint256,address,uint256,bytes,uint256)[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_ParaswapRouter *ParaswapRouterTransactorSession) Buy(data UtilsBuyData) (*types.Transaction, error) {
	return _ParaswapRouter.Contract.Buy(&_ParaswapRouter.TransactOpts, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes ) returns()
func (_ParaswapRouter *ParaswapRouterTransactor) Initialize(opts *bind.TransactOpts, arg0 []byte) (*types.Transaction, error) {
	return _ParaswapRouter.contract.Transact(opts, "initialize", arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes ) returns()
func (_ParaswapRouter *ParaswapRouterSession) Initialize(arg0 []byte) (*types.Transaction, error) {
	return _ParaswapRouter.Contract.Initialize(&_ParaswapRouter.TransactOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes ) returns()
func (_ParaswapRouter *ParaswapRouterTransactorSession) Initialize(arg0 []byte) (*types.Transaction, error) {
	return _ParaswapRouter.Contract.Initialize(&_ParaswapRouter.TransactOpts, arg0)
}

// MegaSwap is a paid mutator transaction binding the contract method 0x46c67b6d.
//
// Solidity: function megaSwap((address,uint256,uint256,uint256,address,(uint256,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_ParaswapRouter *ParaswapRouterTransactor) MegaSwap(opts *bind.TransactOpts, data UtilsMegaSwapSellData) (*types.Transaction, error) {
	return _ParaswapRouter.contract.Transact(opts, "megaSwap", data)
}

// MegaSwap is a paid mutator transaction binding the contract method 0x46c67b6d.
//
// Solidity: function megaSwap((address,uint256,uint256,uint256,address,(uint256,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_ParaswapRouter *ParaswapRouterSession) MegaSwap(data UtilsMegaSwapSellData) (*types.Transaction, error) {
	return _ParaswapRouter.Contract.MegaSwap(&_ParaswapRouter.TransactOpts, data)
}

// MegaSwap is a paid mutator transaction binding the contract method 0x46c67b6d.
//
// Solidity: function megaSwap((address,uint256,uint256,uint256,address,(uint256,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_ParaswapRouter *ParaswapRouterTransactorSession) MegaSwap(data UtilsMegaSwapSellData) (*types.Transaction, error) {
	return _ParaswapRouter.Contract.MegaSwap(&_ParaswapRouter.TransactOpts, data)
}

// MultiSwap is a paid mutator transaction binding the contract method 0xa94e78ef.
//
// Solidity: function multiSwap((address,uint256,uint256,uint256,address,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_ParaswapRouter *ParaswapRouterTransactor) MultiSwap(opts *bind.TransactOpts, data UtilsSellData) (*types.Transaction, error) {
	return _ParaswapRouter.contract.Transact(opts, "multiSwap", data)
}

// MultiSwap is a paid mutator transaction binding the contract method 0xa94e78ef.
//
// Solidity: function multiSwap((address,uint256,uint256,uint256,address,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_ParaswapRouter *ParaswapRouterSession) MultiSwap(data UtilsSellData) (*types.Transaction, error) {
	return _ParaswapRouter.Contract.MultiSwap(&_ParaswapRouter.TransactOpts, data)
}

// MultiSwap is a paid mutator transaction binding the contract method 0xa94e78ef.
//
// Solidity: function multiSwap((address,uint256,uint256,uint256,address,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_ParaswapRouter *ParaswapRouterTransactorSession) MultiSwap(data UtilsSellData) (*types.Transaction, error) {
	return _ParaswapRouter.Contract.MultiSwap(&_ParaswapRouter.TransactOpts, data)
}

// ParaswapRouterBoughtV3Iterator is returned from FilterBoughtV3 and is used to iterate over the raw logs and unpacked data for BoughtV3 events raised by the ParaswapRouter contract.
type ParaswapRouterBoughtV3Iterator struct {
	Event *ParaswapRouterBoughtV3 // Event containing the contract specifics and raw log

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
func (it *ParaswapRouterBoughtV3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParaswapRouterBoughtV3)
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
		it.Event = new(ParaswapRouterBoughtV3)
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
func (it *ParaswapRouterBoughtV3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParaswapRouterBoughtV3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParaswapRouterBoughtV3 represents a BoughtV3 event raised by the ParaswapRouter contract.
type ParaswapRouterBoughtV3 struct {
	Uuid           [16]byte
	Partner        common.Address
	FeePercent     *big.Int
	Initiator      common.Address
	Beneficiary    common.Address
	SrcToken       common.Address
	DestToken      common.Address
	SrcAmount      *big.Int
	ReceivedAmount *big.Int
	ExpectedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBoughtV3 is a free log retrieval operation binding the contract event 0x4cc7e95e48af62690313a0733e93308ac9a73326bc3c29f1788b1191c376d5b6.
//
// Solidity: event BoughtV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_ParaswapRouter *ParaswapRouterFilterer) FilterBoughtV3(opts *bind.FilterOpts, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (*ParaswapRouterBoughtV3Iterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _ParaswapRouter.contract.FilterLogs(opts, "BoughtV3", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return &ParaswapRouterBoughtV3Iterator{contract: _ParaswapRouter.contract, event: "BoughtV3", logs: logs, sub: sub}, nil
}

// WatchBoughtV3 is a free log subscription operation binding the contract event 0x4cc7e95e48af62690313a0733e93308ac9a73326bc3c29f1788b1191c376d5b6.
//
// Solidity: event BoughtV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_ParaswapRouter *ParaswapRouterFilterer) WatchBoughtV3(opts *bind.WatchOpts, sink chan<- *ParaswapRouterBoughtV3, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _ParaswapRouter.contract.WatchLogs(opts, "BoughtV3", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParaswapRouterBoughtV3)
				if err := _ParaswapRouter.contract.UnpackLog(event, "BoughtV3", log); err != nil {
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

// ParseBoughtV3 is a log parse operation binding the contract event 0x4cc7e95e48af62690313a0733e93308ac9a73326bc3c29f1788b1191c376d5b6.
//
// Solidity: event BoughtV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_ParaswapRouter *ParaswapRouterFilterer) ParseBoughtV3(log types.Log) (*ParaswapRouterBoughtV3, error) {
	event := new(ParaswapRouterBoughtV3)
	if err := _ParaswapRouter.contract.UnpackLog(event, "BoughtV3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParaswapRouterMyDebugIterator is returned from FilterMyDebug and is used to iterate over the raw logs and unpacked data for MyDebug events raised by the ParaswapRouter contract.
type ParaswapRouterMyDebugIterator struct {
	Event *ParaswapRouterMyDebug // Event containing the contract specifics and raw log

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
func (it *ParaswapRouterMyDebugIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParaswapRouterMyDebug)
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
		it.Event = new(ParaswapRouterMyDebug)
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
func (it *ParaswapRouterMyDebugIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParaswapRouterMyDebugIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParaswapRouterMyDebug represents a MyDebug event raised by the ParaswapRouter contract.
type ParaswapRouterMyDebug struct {
	Sender      common.Address
	DstReceiver common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMyDebug is a free log retrieval operation binding the contract event 0x312723079d4bc09a73253d15f0e7c8ee28a716a42bc96e13886815962cf93b41.
//
// Solidity: event MyDebug(address sender, address dstReceiver, uint256 amount)
func (_ParaswapRouter *ParaswapRouterFilterer) FilterMyDebug(opts *bind.FilterOpts) (*ParaswapRouterMyDebugIterator, error) {

	logs, sub, err := _ParaswapRouter.contract.FilterLogs(opts, "MyDebug")
	if err != nil {
		return nil, err
	}
	return &ParaswapRouterMyDebugIterator{contract: _ParaswapRouter.contract, event: "MyDebug", logs: logs, sub: sub}, nil
}

// WatchMyDebug is a free log subscription operation binding the contract event 0x312723079d4bc09a73253d15f0e7c8ee28a716a42bc96e13886815962cf93b41.
//
// Solidity: event MyDebug(address sender, address dstReceiver, uint256 amount)
func (_ParaswapRouter *ParaswapRouterFilterer) WatchMyDebug(opts *bind.WatchOpts, sink chan<- *ParaswapRouterMyDebug) (event.Subscription, error) {

	logs, sub, err := _ParaswapRouter.contract.WatchLogs(opts, "MyDebug")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParaswapRouterMyDebug)
				if err := _ParaswapRouter.contract.UnpackLog(event, "MyDebug", log); err != nil {
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

// ParseMyDebug is a log parse operation binding the contract event 0x312723079d4bc09a73253d15f0e7c8ee28a716a42bc96e13886815962cf93b41.
//
// Solidity: event MyDebug(address sender, address dstReceiver, uint256 amount)
func (_ParaswapRouter *ParaswapRouterFilterer) ParseMyDebug(log types.Log) (*ParaswapRouterMyDebug, error) {
	event := new(ParaswapRouterMyDebug)
	if err := _ParaswapRouter.contract.UnpackLog(event, "MyDebug", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParaswapRouterSwappedV3Iterator is returned from FilterSwappedV3 and is used to iterate over the raw logs and unpacked data for SwappedV3 events raised by the ParaswapRouter contract.
type ParaswapRouterSwappedV3Iterator struct {
	Event *ParaswapRouterSwappedV3 // Event containing the contract specifics and raw log

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
func (it *ParaswapRouterSwappedV3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParaswapRouterSwappedV3)
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
		it.Event = new(ParaswapRouterSwappedV3)
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
func (it *ParaswapRouterSwappedV3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParaswapRouterSwappedV3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParaswapRouterSwappedV3 represents a SwappedV3 event raised by the ParaswapRouter contract.
type ParaswapRouterSwappedV3 struct {
	Uuid           [16]byte
	Partner        common.Address
	FeePercent     *big.Int
	Initiator      common.Address
	Beneficiary    common.Address
	SrcToken       common.Address
	DestToken      common.Address
	SrcAmount      *big.Int
	ReceivedAmount *big.Int
	ExpectedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSwappedV3 is a free log retrieval operation binding the contract event 0xe00361d207b252a464323eb23d45d42583e391f2031acdd2e9fa36efddd43cb0.
//
// Solidity: event SwappedV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_ParaswapRouter *ParaswapRouterFilterer) FilterSwappedV3(opts *bind.FilterOpts, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (*ParaswapRouterSwappedV3Iterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _ParaswapRouter.contract.FilterLogs(opts, "SwappedV3", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return &ParaswapRouterSwappedV3Iterator{contract: _ParaswapRouter.contract, event: "SwappedV3", logs: logs, sub: sub}, nil
}

// WatchSwappedV3 is a free log subscription operation binding the contract event 0xe00361d207b252a464323eb23d45d42583e391f2031acdd2e9fa36efddd43cb0.
//
// Solidity: event SwappedV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_ParaswapRouter *ParaswapRouterFilterer) WatchSwappedV3(opts *bind.WatchOpts, sink chan<- *ParaswapRouterSwappedV3, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _ParaswapRouter.contract.WatchLogs(opts, "SwappedV3", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParaswapRouterSwappedV3)
				if err := _ParaswapRouter.contract.UnpackLog(event, "SwappedV3", log); err != nil {
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

// ParseSwappedV3 is a log parse operation binding the contract event 0xe00361d207b252a464323eb23d45d42583e391f2031acdd2e9fa36efddd43cb0.
//
// Solidity: event SwappedV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_ParaswapRouter *ParaswapRouterFilterer) ParseSwappedV3(log types.Log) (*ParaswapRouterSwappedV3, error) {
	event := new(ParaswapRouterSwappedV3)
	if err := _ParaswapRouter.contract.UnpackLog(event, "SwappedV3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
