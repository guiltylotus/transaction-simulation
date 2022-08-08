package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"geth/contract"
	"geth/contract/dai"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

type EthCallResponse struct {
	data    hexutil.Bytes // The return data from the call
	gasUsed uint64        // The amount of gas used
	status  uint64        // The return status of the call - 0 for failure or 1 for success.
}

type RpcClient struct {
	Client    *rpc.Client
	EthClient *ethclient.Client
}

var (
	daiContract      = "0x6b175474e89094c44da98b954eedeac495271d0f"
	daiBalanceOfSlot = "2"
	daiAllowanceSlot = "3"

	kncContract      = "0xdeFA4e8a7bcBA345F687a2f1456F5Edd9CE97202"
	kncAllowanceSlot = "102"

	router = "0x00555513acf282b42882420e5e5ba87b44d8fa6e"
	wallet = "0xef09879057a9ad798438f3ba561bcdd293d72fc7"

	encodedSwapData = "0xabcffc2600000000000000000000000041684b361557e9282e0373ca51260d9331e518c90000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000000000000000008400000000000000000000000006b175474e89094c44da98b954eedeac495271d0f000000000000000000000000defa4e8a7bcba345f687a2f1456f5edd9ce9720200000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000160000000000000000000000000198c08797dd4341f738ec18fcd05d64f645b822800000000000000000000000000000000000000000000003635c9adc5dea0000000000000000000000000000000000000000000000000001fb304b3f45441997b000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000001a0000000000000000000000000000000000000000000000000000000000000000100000000000000000000000041684b361557e9282e0373ca51260d9331e518c9000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000003635c9adc5dea00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005e0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000006b175474e89094c44da98b954eedeac495271d0f000000000000000000000000defa4e8a7bcba345f687a2f1456f5edd9ce9720200000000000000000000000000000000000000000000001fb304b3f45441997b000000000000000000000000198c08797dd4341f738ec18fcd05d64f645b82280000000000000000000000000000000000000000000000000000000062f73ec800000000000000000000000000000000000000000000000000000000000005a0000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000001e000000000000000000000000000000000000000000000000000000000000003400000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000040100000000000000000000000000000000000000000000000000000000000001200000000000000000000000001116898dda4015ed8ddefb84b6e8bc24528af2d80000000000000000000000006b175474e89094c44da98b954eedeac495271d0f000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec70000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000003635c9adc5dea00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030000000000000000000000001116898dda4015ed8ddefb84b6e8bc24528af2d8000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000005010000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000011b815efb8f581194ae79006d24e0d814b7697f6000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000000000000000000000000000000000003b9388fb0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000061639d6ec06c13a96b5eb9560b359d7c648c7759000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000defa4e8a7bcba345f687a2f1456f5edd9ce97202000000000000000000000000198c08797dd4341f738ec18fcd05d64f645b8228000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000"
)

func main() {
	// Create an IPC based RPC connection to a remote node
	// NOTE update the path to the ipc file!

	startTime := time.Now()
	client := NewRPCClient("/Users/nguyenducminh/ethdata/geth.ipc")
	//client := NewRPCClient("/Users/nguyenducminh/Library/Ethereum/goerli/geth.ipc")
	//client.GetTokenBalanceOf()
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
	rpcClient.EthClient = ethclient.NewClient(rpcClient.Client)
	if err != nil {
		panic(err)
	}
	return rpcClient
}

func (rc *RpcClient) GetTokenBalanceOf() {
	indexDaiBalanceOf := rc.GetIndexBalanceOf(wallet, daiBalanceOfSlot)
	fmt.Println("indexDaiBalanceOf", indexDaiBalanceOf)
	rc.GetBalanceOf(daiContract, indexDaiBalanceOf)

	indexDaiAllowance := rc.GetIndexAllowance(wallet, router, daiAllowanceSlot)
	fmt.Println("indexDaiAllowance", indexDaiAllowance)

	fakeBalance := "0x" + toHashString("0x130EE8E7179044400000")
	fmt.Println("fakeBalance", fakeBalance)

	fakeAllowance := "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
	fmt.Println("fakeAllowance", fakeAllowance)
	//{"params":[{"from":"0x0000000000000000000000000000000000000000",
	//	"data":"0x70a08231000000000000000000000000ef09879057a9ad798438f3ba561bcdd293d72fc7",
	//	"to":"0x6b175474e89094c44da98b954eedeac495271d0f"}]}
	objectEthCall := TransactionArgs{
		From:     "0x0000000000000000000000000000000000000000",
		To:       daiContract,
		GasPrice: "0x9502F9000",
		Gas:      "0x7A1200",
		//Value: "0x8AC7230489E80000",
		Input: "0x70a08231000000000000000000000000ef09879057a9ad798438f3ba561bcdd293d72fc7",
	}

	configEthCall := StateOverride{
		"0xef09879057a9ad798438f3ba561bcdd293d72fc7": OverrideAccount{
			Balance: "0x56BC75E2D63100000",
		},
		daiContract: OverrideAccount{
			StateDiff: map[string]string{
				indexDaiBalanceOf.String(): fakeBalance,
				indexDaiAllowance.String(): fakeAllowance,
			},
		},
	}

	var responseEthCall *string

	if err := rc.Client.Call(
		responseEthCall,
		"eth_call",
		objectEthCall,
		"latest", configEthCall); err != nil {
		panic(err)
	}

	fmt.Println(responseEthCall)
}

func (rc *RpcClient) GetStructLogs() []StructLog {
	object := TransactionArgs{
		From:     wallet,
		To:       router,
		GasPrice: "0x9502F9000",
		Gas:      "0x7A1200",
		//Value:    "0x8AC7230489E80000",
		Input: encodedSwapData,
	}
	// Goerli
	//object := TransactionArgs{
	//	From:     "0x7ca04051b273a8ce59ebcc260bb2c10da93d2059",
	//	To:       "0x64e8f498c3Ee44Dd7A6ce3d1E77858413a380627",
	//	GasPrice: "0xffffffff",
	//	Gas:      "0xffff",
	//	Input:    "0xe8927fbc",
	//}

	indexDaiBalanceOf := rc.GetIndexBalanceOf(wallet, daiBalanceOfSlot)
	fmt.Println("indexDaiBalanceOf", indexDaiBalanceOf)
	//rc.GetBalanceOf(daiContract, indexDaiBalanceOf)

	indexDaiAllowance := rc.GetIndexAllowance(wallet, router, daiAllowanceSlot)
	fmt.Println("indexDaiAllowance", indexDaiAllowance)

	indexKncAllowance := rc.GetIndexAllowance(wallet, router, kncAllowanceSlot)
	fmt.Println("indexKncAllowance", indexKncAllowance)

	fakeBalance := "0x" + toHashString("0x130EE8E7179044400000")
	fmt.Println("fakeBalance", fakeBalance)

	fakeAllowance := "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
	fmt.Println("fakeAllowance", fakeAllowance)

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
					indexDaiBalanceOf.String(): fakeBalance,
					indexDaiAllowance.String(): fakeAllowance,
				},
			},
			kncContract: OverrideAccount{
				StateDiff: map[string]string{
					indexKncAllowance.String(): fakeAllowance,
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
	if response.Failed {
		returnValueString, err := hex.DecodeString(response.ReturnValue)
		if err != nil {
			panic(err)
		}
		fmt.Println("returnValueString", string(returnValueString))
		panic("Call Failed")
	}
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

	type ErrorEvent struct {
		Reason string
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

	// Error
	logErrorSig := []byte("Error(string)")
	logErrorSigHash := crypto.Keccak256Hash(logErrorSig)
	logErrorEvent := common.HexToHash(logErrorSigHash.String())
	//fmt.Println("---logErrorEvent: ", logErrorEvent.Hex())

	// contractABI
	contractAbi, err := abi.JSON(strings.NewReader(contract.ContractMetaData.ABI))
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
		case logErrorEvent.Hex():
			var event ErrorEvent
			err = contractAbi.UnpackIntoInterface(&event, "Error", vLogData)
			if err != nil {
				fmt.Println("ERR ErrorEvent", err)
			}
			fmt.Println("Decode Error Event",
				"Reason", event.Reason,
			)
			continue
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

func (rc *RpcClient) GetBalanceOf(contractAddress string, index common.Hash) {
	state, err := rc.EthClient.StorageAt(
		context.Background(),
		common.HexToAddress(contractAddress),
		index, nil)
	if err != nil {
		panic(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(dai.ContractMetaData.ABI))
	if err != nil {
		panic(err)
	}
	balanceOf, err := contractAbi.Unpack("balanceOf", state)
	if err != nil {
		panic(err)
	}
	fmt.Println("Storage At: balanceOf", balanceOf)

	//realByteState, err := hex.DecodeString("000000000000000000000000000000000000000000000000DC11DF3655EF0D3E")
	//if err != nil {
	//	panic(err)
	//}
	realByteState := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 54, 53, 201, 173, 197, 222, 160, 0, 0}
	checkBalanceOf, err := contractAbi.Unpack("balanceOf", realByteState)
	if err != nil {
		panic(err)
	}
	fmt.Println("Storage At: CheckBalanceOf", checkBalanceOf)

}

func (rc *RpcClient) GetIndexBalanceOf(owner string, slot string) common.Hash {
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
	return common.BytesToHash(index)
}

func (rc *RpcClient) GetIndexAllowance(owner string, spender string, slot string) common.Hash {
	slot = toHashString(slot)
	owner = toHashString(owner)
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
	index := solsha3.SoliditySHA3(
		// types
		[]string{"address", "address"},

		// values
		[]interface{}{
			spender,
			tempStr,
		},
	)
	return common.BytesToHash(index)
}
