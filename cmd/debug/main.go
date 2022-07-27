package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type NewID struct {
	Date *big.Int
	Id   *big.Int
}

type TransactionArgs struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Input    string `json:"input"`
	ChainID  string `json:"chainId,omitempty"`
}

type TraceConfig struct {
	DisableStorage   bool `json:"disableStorage"`
	DisableStack     bool `json:"disableStack"`
	EnableMemory     bool `json:"enableMemory"`
	EnableReturnData bool `json:"enableReturnData"`
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

func main() {
	// Create an IPC based RPC connection to a remote node
	// NOTE update the path to the ipc file!

	cl, err := rpc.Dial("/Users/nguyenducminh/Library/Ethereum/goerli/geth.ipc")
	if err != nil {
		panic(err)
	}
	defer cl.Close()

	GoerliDebug(cl)
}

func GoerliDebug(cl *rpc.Client) {
	object := TransactionArgs{
		From:     "0x7ca04051b273a8ce59ebcc260bb2c10da93d2059",
		To:       "0xead507e4fa6674ab296f6b8194F31bf24E4ac3C6",
		GasPrice: "0xffffffff",
		Gas:      "0xffff",
		Input:    "0xe8927fbc",
	}
	config := TraceConfig{
		DisableStorage:   false,
		DisableStack:     false,
		EnableMemory:     true,
		EnableReturnData: true,
	}

	response := &DebugTraceCallResponse{}
	if err := cl.Call(response, "debug_traceCall", object, "latest", config); err != nil {
		panic(err)
	}
	fmt.Println(response.Failed)
	fmt.Println(response.ReturnValue)
	fmt.Println(response.Gas)
	for _, structLog := range response.StructLogs {
		fmt.Println("Op:", structLog.Op)
		fmt.Println("Gas:", structLog.Gas)
		fmt.Println("Memory:", structLog.Memory)
		fmt.Println("Stack:", structLog.Stack)
		fmt.Println("-------------------")
	}
}
