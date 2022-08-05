package main

import (
	"encoding/hex"
	"fmt"
	"geth/contract"
	"geth/contract/dai"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
	"strconv"
	"strings"
	"time"
)

type NewID struct {
	Date *big.Int
	Id   *big.Int
}

type OldID struct {
	Date *big.Int
	Id   *big.Int
}

type EthTransactionArgs struct {
	From                 *common.Address `json:"from"`
	To                   *common.Address `json:"to"`
	Gas                  *hexutil.Uint64 `json:"gas"`
	GasPrice             *hexutil.Big    `json:"gasPrice"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas"`
	Value                *hexutil.Big    `json:"value"`
	Nonce                *hexutil.Uint64 `json:"nonce"`

	// We accept "data" and "input" for backwards-compatibility reasons.
	// "input" is the newer name and should be preferred by clients.
	// Issue detail: https://github.com/ethereum/go-ethereum/issues/15628
	Data  *hexutil.Bytes `json:"data"`
	Input *hexutil.Bytes `json:"input"`

	// Introduced by AccessListTxType transaction.
	AccessList *types.AccessList `json:"accessList,omitempty"`
	ChainID    *hexutil.Big      `json:"chainId,omitempty"`
}

type TransactionArgs struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Input    string `json:"input"`
	Value    string `json:"value"`
}

type OverrideAccount struct {
	//Nonce     *hexutil.Uint64              `json:"nonce"`
	//Code      *hexutil.Bytes               `json:"code"`
	Balance string `json:"balance"`
	//State     *map[common.Hash]common.Hash `json:"state"`
	StateDiff map[string]string `json:"stateDiff"`
}

type StateOverride map[string]OverrideAccount

type TraceConfig struct {
	DisableStorage   bool          `json:"disableStorage"`
	DisableStack     bool          `json:"disableStack"`
	EnableMemory     bool          `json:"enableMemory"`
	EnableReturnData bool          `json:"enableReturnData"`
	Tracer           string        `json:"tracer"`
	Timeout          string        `json:"timeout"`
	StateOverrides   StateOverride `json:"stateOverrides"`
}

type StructLog struct {
	Pc            uint64                      `json:"pc"`
	Op            string                      `json:"op"`
	Gas           uint64                      `json:"gas"`
	GasCost       uint64                      `json:"gasCost"`
	Memory        []string                    `json:"memory,omitempty"`
	MemorySize    int                         `json:"memSize"`
	Stack         []string                    `json:"stack"`
	ReturnData    []byte                      `json:"returnData,omitempty"`
	Storage       map[common.Hash]common.Hash `json:"-"`
	Depth         int                         `json:"depth"`
	RefundCounter uint64                      `json:"refund"`
	Err           error                       `json:"-"`
}

type DebugTraceCallResponse struct {
	Failed      bool        `json:"failed"`
	Gas         uint64      `json:"gas"`
	ReturnValue string      `json:"returnValue"`
	StructLogs  []StructLog `json:"structLogs"`
}

type RpcClient struct {
	Client *rpc.Client
}

var (
	daiContract      = "0x6b175474e89094c44da98b954eedeac495271d0f"
	daiBalanceOfSlot = "2"
)

func main() {
	// Create an IPC based RPC connection to a remote node
	// NOTE update the path to the ipc file!

	startTime := time.Now()
	client := NewRPCClient("/Users/nguyenducminh/ethdata/geth.ipc")
	//client := NewRPCClient("/Users/nguyenducminh/Library/Ethereum/goerli/geth.ipc")
	structLogs := client.GetStructLogs()
	client.GetEtherKyberSwapLosgs(structLogs)
	//client.GetGoerliLogs(structLogs)
	fmt.Println("Execution time: ", time.Now().Sub(startTime))
}

func NewRPCClient(address string) *RpcClient {
	var err error
	rpcClient := &RpcClient{}
	rpcClient.Client, err = rpc.Dial(address)
	if err != nil {
		panic(err)
	}
	return rpcClient
}

func (rc *RpcClient) GetStructLogs() []StructLog {
	object := TransactionArgs{
		From:     "0xef09879057a9ad798438f3ba561bcdd293d72fc7",
		To:       "0x00555513acf282b42882420e5e5ba87b44d8fa6e",
		GasPrice: "0x9502F9000",
		Gas:      "0x7A1200",
		Value:    "0x8AC7230489E80000",
		Input:    "0xabcffc2600000000000000000000000041684b361557e9282e0373ca51260d9331e518c90000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000006a00000000000000000000000001af3f329e8be154074d8769d1ffa4ee058b1dbc3000000000000000000000000fe56d5892bdffc7bf58f2e84be1b2c32d21c308b00000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000140000000000000000000000000ef09879057a9ad798438f3ba561bcdd293d72fc700000000000000000000000000000000000000000000003635c9adc5dea0000000000000000000000000000000000000000000000000001710ab951157199b42000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000480000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000062ecc2bd0000000000000000000000000000000000000000000000000000000000000440000000000000000000000000000000000000000000000000000000000000000100000000000000000000000066fdb2eccfb58cf098eaa419e5efde841368e489000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000003635c9adc5dea000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000002c000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000c000000000000000000000000066fdb2eccfb58cf098eaa419e5efde841368e4890000000000000000000000001af3f329e8be154074d8769d1ffa4ee058b1dbc3000000000000000000000000e9e7cea3dedca5984780bafc599bd69add087d5600000000000000000000000001f1af8439639694593284ecb0ef1e16cef8de43000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000c000000000000000000000000001f1af8439639694593284ecb0ef1e16cef8de43000000000000000000000000e9e7cea3dedca5984780bafc599bd69add087d56000000000000000000000000fe56d5892bdffc7bf58f2e84be1b2c32d21c308b000000000000000000000000ef09879057a9ad798438f3ba561bcdd293d72fc7000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000",
	}
	// Goerli
	//object := TransactionArgs{
	//	From:     "0x7ca04051b273a8ce59ebcc260bb2c10da93d2059",
	//	To:       "0x64e8f498c3Ee44Dd7A6ce3d1E77858413a380627",
	//	GasPrice: "0xffffffff",
	//	Gas:      "0xffff",
	//	Input:    "0xe8927fbc",
	//}

	indexBalanceOf := rc.GetIndexBalanceOf("0xef09879057a9ad798438f3ba561bcdd293d72fc7", daiBalanceOfSlot)
	fakeBalance := "0x" + toHashString("0x8AC7230489E80000")
	fmt.Println("fakeBalance", fakeBalance)

	config := TraceConfig{
		DisableStorage:   false,
		DisableStack:     false,
		EnableMemory:     true,
		EnableReturnData: true,
		Tracer:           "loggetter",
		Timeout:          "20s",
		StateOverrides: StateOverride{
			"0xef09879057a9ad798438f3ba561bcdd293d72fc7": OverrideAccount{
				Balance: "0x56BC75E2D63100000",
			},
			daiContract: OverrideAccount{
				StateDiff: map[string]string{
					indexBalanceOf: fakeBalance,
				},
			},
		},
	}

	structLogs := make([]StructLog, 0, 0)
	response := &DebugTraceCallResponse{}
	if err := rc.Client.Call(response, "debug_traceCall", object, "latest", config); err != nil {
		panic(err)
	}
	fmt.Println(response.Failed)
	fmt.Println(response.ReturnValue)
	fmt.Println(response.Gas)
	for _, structLog := range response.StructLogs {
		if strings.HasPrefix(structLog.Op, "LOG") {
			structLogs = append(structLogs, structLog)
			fmt.Println("Pc:", structLog.Pc)
			fmt.Println("Op:", structLog.Op)
			fmt.Println("Gas:", structLog.Gas)
			//fmt.Println("Memory:", structLog.Memory)
			//fmt.Println("Stack:", structLog.Stack)
			fmt.Println("-------------------")
		}
	}
	return structLogs
}

func (rc *RpcClient) GetEtherKyberSwapLosgs(structLogs []StructLog) {
	fmt.Println("GetEtherKyberSwapLosgs---")
	// Event
	type Swapped struct {
		Sender       common.Address
		SrcToken     common.Address
		DstToken     common.Address
		DstReceiver  common.Address
		SpentAmount  *big.Int
		ReturnAmount *big.Int
	}

	type Exchange struct {
		Pair      common.Address
		AmountOut *big.Int
		Output    common.Address
	}

	type Approval struct {
		Owner   common.Address
		Spender common.Address
		Value   *big.Int
	}

	type Transfer struct {
		From  common.Address
		To    common.Address
		Value *big.Int
	}

	// Approval
	logApprovalSig := []byte("Approval(address,address,uint256)")
	logApprovalSigHash := crypto.Keccak256Hash(logApprovalSig)
	logApprovalEvent := common.HexToHash(logApprovalSigHash.String())
	//fmt.Println("---logApprovalEvent: ", logApprovalEvent.Hex())

	// Transfer
	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logTransferEvent := common.HexToHash(logTransferSigHash.String())
	//fmt.Println("---logTransferEvent: ", logTransferEvent.Hex())

	// Exchange
	logExchangeSig := []byte("Exchange(address,uint256,address)")
	logExchangeSigHash := crypto.Keccak256Hash(logExchangeSig)
	logExchangeEvent := common.HexToHash(logExchangeSigHash.String())
	//fmt.Println("---logExchangeEvent: ", logExchangeEvent.Hex())

	// Swapped
	logSwappedSig := []byte("Swapped(address,address,address,address,uint256,uint256)")
	logSwappedSigHash := crypto.Keccak256Hash(logSwappedSig)
	logSwappedEvent := common.HexToHash(logSwappedSigHash.String())
	//fmt.Println("---logSwappedEvent: ", logSwappedEvent.Hex())

	// contractABI
	contractAbi, err := abi.JSON(strings.NewReader(dai.ContractMetaData.ABI))
	if err != nil {
		panic(err)
	}

	for _, log := range structLogs {
		fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("Opcode: ", log.Op)
		topics, memoryHexString := rc.GetTopicAndData(log)
		vLogData, err := hex.DecodeString(memoryHexString)
		if err != nil {
			panic(err)
		}
		switch topics[0] {
		case logSwappedEvent.Hex():
			var event Swapped
			err = contractAbi.UnpackIntoInterface(&event, "Swapped", vLogData)
			if err != nil {
				fmt.Println("ERR logSwappedEvent", err)
				continue
				//panic(err)
			}
			fmt.Println(
				"DecodeEvent",
				"sender", event.Sender,
				"srcToken", event.SrcToken,
				"dstToken", event.DstToken,
				"dstReceive", event.DstReceiver,
				"spentAmount", event.SpentAmount,
				"returnAmount", event.ReturnAmount,
			)
		case logExchangeEvent.Hex():
			var event Exchange
			err = contractAbi.UnpackIntoInterface(&event, "Exchange", vLogData)
			if err != nil {
				fmt.Println("ERR logExchangeEvent", err)
				continue
				//panic(err)
			}
			fmt.Println(
				"DecodeEvent",
				"Pair", event.Pair,
				"AmountOut", event.AmountOut,
				"Output", event.Output,
			)
		case logApprovalEvent.Hex():
			fmt.Println("This is pproval event")
			continue
			//var event Approval
			//err = contractAbi.UnpackIntoInterface(&event, "Approval", vLogData)
			//if err != nil {
			//	fmt.Println("ERR logApprovalEvent", err)
			//	continue
			//	//panic(err)
			//}
			//fmt.Println(
			//	"DecodeEvent",
			//	"Owner", event.Owner,
			//	"Spender", event.Spender,
			//	"Value", event.Value,
			//)
		case logTransferEvent.Hex():
			fmt.Println("This is Transfer event")
			continue
			//var event Transfer
			//err = contractAbi.UnpackIntoInterface(&event, "Transfer", vLogData)
			//if err != nil {
			//	fmt.Println("ERR logApprovalEvent", err)
			//	continue
			//	//panic(err)
			//}
			//fmt.Println(
			//	"DecodeEvent",
			//	"From", event.From,
			//	"Spender", event.To,
			//	"Value", event.Value,
			//)
		default:
			fmt.Println("Not match any event")
		}
	}
}

func (rc *RpcClient) GetGoerliLogs(structLogs []StructLog) {
	// type of log events
	logNewIDSig := []byte("NewID(uint256,uint256)")
	LogOldIDSig := []byte("OldID(uint256,uint256)")

	logNewIDSigHash := crypto.Keccak256Hash(logNewIDSig)
	logOldIDSigHash := crypto.Keccak256Hash(LogOldIDSig)

	logNewIDEvent := common.HexToHash(logNewIDSigHash.String())
	logOldIDEvent := common.HexToHash(logOldIDSigHash.String())

	// contractABI
	contractAbi, err := abi.JSON(strings.NewReader(contract.IncreaseEventMetaData.ABI))
	if err != nil {
		panic(err)
	}

	for _, log := range structLogs {
		if log.Op != "LOG1" {
			continue
		}
		topics, memoryHexString := rc.GetTopicAndData(log)
		vLogData, err := hex.DecodeString(memoryHexString)
		if err != nil {
			panic(err)
		}

		switch topics[0] {
		case logNewIDEvent.Hex():
			var event NewID
			err = contractAbi.UnpackIntoInterface(&event, "NewID", vLogData)
			if err != nil {
				panic(err)
			}
			fmt.Println("DecodeEvent", "date", event.Date, "id", event.Id)
		case logOldIDEvent.Hex():
			var event OldID
			err = contractAbi.UnpackIntoInterface(&event, "OldID", vLogData)
			if err != nil {
				panic(err)
			}
			fmt.Println("DecodeEvent", "date", event.Date, "id", event.Id)
		}
	}
}

func (rc *RpcClient) GetTopicAndData(log StructLog) ([]string, string) {
	offset := hex2int(log.Stack[len(log.Stack)-1])
	length := hex2int(log.Stack[len(log.Stack)-2])

	hexMemory := strings.Join(log.Memory[:], "")
	byteMemory, err := hex.DecodeString(hexMemory)
	if err != nil {
		panic(err)
	}
	byteData := byteMemory[offset : offset+length]
	memory := hex.EncodeToString(byteData)

	topicCount, err := strconv.Atoi(string(log.Op[len(log.Op)-1]))
	if err != nil {
		panic(err)
	}

	//fmt.Println("TOPIC COUNT", topicCount)
	topics := make([]string, 0, topicCount+1)
	for i := len(log.Stack) - 3; i > len(log.Stack)-3-topicCount; i-- {
		topics = append(topics, log.Stack[i])
	}

	fmt.Println("TOPIC", topics)
	fmt.Println("MEMORY", memory)
	return topics, memory
}

func big2hex(b *big.Int) string {
	return fmt.Sprintf("%x", b)
}

func hex2int(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return uint64(result)
}

func toHashString(hexStr string) string {
	str := strings.Replace(hexStr, "0x", "", -1)
	str = strings.ToLower(str)
	return fmt.Sprintf("%064v", str)
}

func (rc *RpcClient) GetIndexBalanceOf(owner string, slot string) string {
	slot = toHashString(slot)
	owner = toHashString(owner)

	index := solsha3.SoliditySHA3(
		// types
		[]string{"address", "uint256"},

		// values
		[]interface{}{
			owner,
			slot,
		},
	)
	return common.BytesToHash(index).String()
}

func (rc *RpcClient) GetAllowanceIndex(owner string, spender string, slot string) {
	slot = toHashString(slot)
	fmt.Println("slot", slot)
	owner = toHashString(owner)
	fmt.Println("owner", owner)
	temp := solsha3.SoliditySHA3(
		// types
		[]string{"address", "uint256"},

		// values
		[]interface{}{
			owner,
			slot,
		},
	)

	tempStr := toHashString(common.BytesToHash(temp).String())
	spender = toHashString(spender)
	fmt.Println("spender", spender)
	index := solsha3.SoliditySHA3(
		// types
		[]string{"address", "address"},

		// values
		[]interface{}{
			spender,
			tempStr,
		},
	)
	fmt.Println("index", index)
}
