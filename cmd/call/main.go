package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"geth/contract/simswap"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"
)

const (
	expBase = 10
)

type Account struct {
	Nonce     string            `json:"nonce,omitempty"`
	Balance   string            `json:"balance,omitempty"`
	Code      string            `json:"code,omitempty"`
	State     map[string]string `json:"state,omitempty"`
	StateDiff map[string]string `json:"stateDiff,omitempty"`
}

type OverrideAccounts map[common.Address]Account

type roundTripperExt struct {
	c          *http.Client
	appendData json.RawMessage
}

type reqMessage struct {
	JSONRPC string            `json:"jsonrpc"`
	ID      int               `json:"id"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
}

var (
	SimSwapAddress = common.HexToAddress("0x1111111111111111111111111111111111111100")
	MyWallet       = common.HexToAddress("0x198c08797dd4341f738ec18fcd05d64f645b8228")
	//Router         = common.HexToAddress("0x00555513Acf282B42882420E5e5bA87b44D8fA6E") //kyberswap
	Router = common.HexToAddress("0x1111111254fb6c44bac0bed2854e76f90643097d") // 1inch

	DAIContract = common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")
	KNCContract = common.HexToAddress("0xdeFA4e8a7bcBA345F687a2f1456F5Edd9CE97202")

	DAIBalanceOfSlot = "2"
	DAIAllowanceSlot = "3"
	KNCBalanceSlot   = "101"
	KNCAllowanceSlot = "102"

	InputData = "0x7c025200000000000000000000000000f2f400c138f9fb900576263af0bc7fcde2b1b8a8000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000001800000000000000000000000006b175474e89094c44da98b954eedeac495271d0f000000000000000000000000defa4e8a7bcba345f687a2f1456f5edd9ce9720200000000000000000000000020d6b227f4a5a2a13d520329f01bb1f8f9d2d628000000000000000000000000198c08797dd4341f738ec18fcd05d64f645b8228000000000000000000000000000000000000000000000001158e460913d00000000000000000000000000000000000000000000000000000a87e0b9cef4381720000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000018080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a475d22a2700000000000000000000000020d6b227f4a5a2a13d520329f01bb1f8f9d2d6280000000000000000000000006b175474e89094c44da98b954eedeac495271d0f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000061639d6ec06c13a96b5eb9560b359d7c648c775900000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a475d22a2700000000000000000000000061639d6ec06c13a96b5eb9560b359d7c648c7759000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000defa4e8a7bcba345f687a2f1456f5edd9ce972020000000000000000000000001111111254fb6c44bac0bed2854e76f90643097d000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000cfee7c08"
)

func InitCommonContract() *OverrideAccounts {
	indexDaiBalanceOf := getIndexBalanceOf(MyWallet.String(), DAIBalanceOfSlot)
	fmt.Println("indexDaiBalanceOf", indexDaiBalanceOf)

	indexDaiAllowance := getIndexAllowance(MyWallet.String(), SimSwapAddress.String(), DAIAllowanceSlot)
	fmt.Println("indexDaiAllowance", indexDaiAllowance)

	indexKncAllowance := getIndexAllowance(MyWallet.String(), SimSwapAddress.String(), KNCAllowanceSlot)
	fmt.Println("indexKncAllowance", indexKncAllowance)

	indexKncBalance := getIndexAllowance(MyWallet.String(), SimSwapAddress.String(), KNCBalanceSlot)
	fmt.Println("indexKncBalance", indexKncBalance)

	fakeBalance := "0x" + toHashString("0x3635C9ADC5DEA00000")
	fmt.Println("fakeBalance", fakeBalance)

	fakeAllowance := "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
	fmt.Println("fakeAllowance", fakeAllowance)

	return &OverrideAccounts{
		SimSwapAddress: {
			Nonce: "0x10",
			Code:  "0x60806040526004361061003f5760003560e01c806321c4f09f1461004457806368116177146100745780637e5465ba146100a457806396d27420146100e1575b600080fd5b61005e600480360381019061005991906107cc565b610112565b60405161006b9190610825565b60405180910390f35b61008e60048036038101906100899190610840565b6101a3565b60405161009b9190610825565b60405180910390f35b3480156100b057600080fd5b506100cb60048036038101906100c691906107cc565b610231565b6040516100d89190610825565b60405180910390f35b6100fb60048036038101906100f691906108d2565b6102e1565b60405161010992919061095a565b60405180910390f35b60008083905060008173ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33866040518363ffffffff1660e01b8152600401610155929190610992565b602060405180830381865afa158015610172573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061019691906109e7565b9050809250505092915050565b60008082905060008173ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016101e49190610a14565b602060405180830381865afa158015610201573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061022591906109e7565b90508092505050919050565b6000808390508073ffffffffffffffffffffffffffffffffffffffff1663095ea7b3847f80000000000000000000000000000000000000000000000000000000000000006040518363ffffffff1660e01b8152600401610292929190610a74565b6020604051808303816000875af11580156102b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102d59190610ad5565b50600091505092915050565b6000806000879050600087905060008273ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016103299190610a14565b602060405180830381865afa158015610346573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036a91906109e7565b905060008273ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016103a79190610a14565b602060405180830381865afa1580156103c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e891906109e7565b90503073ffffffffffffffffffffffffffffffffffffffff16637e5465ba8c8b6040518363ffffffff1660e01b8152600401610425929190610992565b6020604051808303816000875af1158015610444573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061046891906109e7565b506104b78989898080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506105e0565b5060008473ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016104f39190610a14565b602060405180830381865afa158015610510573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053491906109e7565b905060008473ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016105719190610a14565b602060405180830381865afa15801561058e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b291906109e7565b905081846105c09190610b31565b975082816105ce9190610b31565b96505050505050509550959350505050565b60606106058383604051806060016040528060278152602001610d116027913961060d565b905092915050565b6060610618846106da565b610657576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064e90610be8565b60405180910390fd5b6000808573ffffffffffffffffffffffffffffffffffffffff168560405161067f9190610c82565b600060405180830381855af49150503d80600081146106ba576040519150601f19603f3d011682016040523d82523d6000602084013e6106bf565b606091505b50915091506106cf8282866106fd565b925050509392505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6060831561070d5782905061075d565b6000835111156107205782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107549190610cee565b60405180910390fd5b9392505050565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006107998261076e565b9050919050565b6107a98161078e565b81146107b457600080fd5b50565b6000813590506107c6816107a0565b92915050565b600080604083850312156107e3576107e2610764565b5b60006107f1858286016107b7565b9250506020610802858286016107b7565b9150509250929050565b6000819050919050565b61081f8161080c565b82525050565b600060208201905061083a6000830184610816565b92915050565b60006020828403121561085657610855610764565b5b6000610864848285016107b7565b91505092915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126108925761089161086d565b5b8235905067ffffffffffffffff8111156108af576108ae610872565b5b6020830191508360018202830111156108cb576108ca610877565b5b9250929050565b6000806000806000608086880312156108ee576108ed610764565b5b60006108fc888289016107b7565b955050602061090d888289016107b7565b945050604061091e888289016107b7565b935050606086013567ffffffffffffffff81111561093f5761093e610769565b5b61094b8882890161087c565b92509250509295509295909350565b600060408201905061096f6000830185610816565b61097c6020830184610816565b9392505050565b61098c8161078e565b82525050565b60006040820190506109a76000830185610983565b6109b46020830184610983565b9392505050565b6109c48161080c565b81146109cf57600080fd5b50565b6000815190506109e1816109bb565b92915050565b6000602082840312156109fd576109fc610764565b5b6000610a0b848285016109d2565b91505092915050565b6000602082019050610a296000830184610983565b92915050565b6000819050919050565b6000819050919050565b6000610a5e610a59610a5484610a2f565b610a39565b61080c565b9050919050565b610a6e81610a43565b82525050565b6000604082019050610a896000830185610983565b610a966020830184610a65565b9392505050565b60008115159050919050565b610ab281610a9d565b8114610abd57600080fd5b50565b600081519050610acf81610aa9565b92915050565b600060208284031215610aeb57610aea610764565b5b6000610af984828501610ac0565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610b3c8261080c565b9150610b478361080c565b925082821015610b5a57610b59610b02565b5b828203905092915050565b600082825260208201905092915050565b7f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60008201527f6e74726163740000000000000000000000000000000000000000000000000000602082015250565b6000610bd2602683610b65565b9150610bdd82610b76565b604082019050919050565b60006020820190508181036000830152610c0181610bc5565b9050919050565b600081519050919050565b600081905092915050565b60005b83811015610c3c578082015181840152602081019050610c21565b83811115610c4b576000848401525b50505050565b6000610c5c82610c08565b610c668185610c13565b9350610c76818560208601610c1e565b80840191505092915050565b6000610c8e8284610c51565b915081905092915050565b600081519050919050565b6000601f19601f8301169050919050565b6000610cc082610c99565b610cca8185610b65565b9350610cda818560208601610c1e565b610ce381610ca4565b840191505092915050565b60006020820190508181036000830152610d088184610cb5565b90509291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212208a670ec4dc4c15570f950427558a542a07e394ffe618b2bae4e15e7d4a2b177864736f6c634300080f0033",
		},
		MyWallet: {
			Balance: "0x8ac7230489e80000",
		},
		DAIContract: {
			StateDiff: map[string]string{
				indexDaiBalanceOf.String(): fakeBalance,
				indexDaiAllowance.String(): fakeAllowance,
			},
		},
		KNCContract: {
			StateDiff: map[string]string{
				indexKncAllowance.String(): fakeAllowance,
				indexKncBalance.String():   fakeBalance,
			},
		},
	}
}

func main() {
	startTime := time.Now()

	commonContract := InitCommonContract()
	// https://etherscan.io/tx/0x606e8c8084855d3fb20cb1c69f520d0a1feae6c35a9d3659a9cda8a1cf53e9e2#eventlog
	//rawurl := "https://proxy.kyberengineering.io/ethereum" // "http://localhost:8545/" //  "https://mainnet.infura.io/v3/3d85e3bded764846bc25e1ca36f73b91" // "https://proxy.kyberengineering.io/ethereum"
	rawurl := "https://mainnet.infura.io/v3/c8a0f577c41240ab90d542d4c1f9f1ba"

	// Generate EncodedSwapData
	simClient, err := NewClient(rawurl, SimSwapAddress, commonContract)
	if err != nil {
		panic(err)
	}
	ab, err := abi.JSON(bytes.NewBufferString(simswap.ContractMetaData.ABI))
	if err != nil {
		panic(err)
	}
	data, err := ab.Pack("simswap", DAIContract, KNCContract, Router, hexutil.MustDecode(InputData))
	if err != nil {
		panic(err)
	}
	msg := ethereum.CallMsg{
		From:      MyWallet,
		To:        &SimSwapAddress,
		Gas:       1000000,
		GasPrice:  FloatToTokenAmount(100, 9),
		GasFeeCap: FloatToTokenAmount(100, 9),
		GasTipCap: FloatToTokenAmount(100, 9),
		Value:     big.NewInt(0),
		Data:      data,
	}
	res, err := simClient.CallContract(context.Background(), msg, nil)
	if err != nil {
		panic(err)
	}

	result, err := ab.Unpack("simswap", res)
	if err != nil {
		panic(err)
	}
	fmt.Println("result'", result)

	fmt.Println("Execution time: ", time.Now().Sub(startTime))
}

func NewClient(rpcURL string, simAddress common.Address, commonContract *OverrideAccounts) (*ethclient.Client, error) {
	httpClient := http.DefaultClient

	sc, err := newSimClient(rpcURL, httpClient, commonContract)
	if err != nil {
		return nil, err
	}

	return sc, nil
}

func newSimClient(url string, client *http.Client, commonContract *OverrideAccounts) (*ethclient.Client, error) {
	round, err := newRoundTripExt(client, commonContract)
	if err != nil {
		return nil, err
	}

	cc := &http.Client{Transport: round}
	r, err := rpc.DialHTTPWithClient(url, cc)
	if err != nil {
		return nil, errors.WithMessage(err, "simclient: dial rpc")
	}

	ethClient := ethclient.NewClient(r)

	return ethClient, nil
}

func newRoundTripExt(c *http.Client, accounts *OverrideAccounts) (http.RoundTripper, error) {
	data, err := json.Marshal(accounts)
	if err != nil {
		return nil, err
	}
	return &roundTripperExt{
		c:          c,
		appendData: data,
	}, nil
}

func (r roundTripperExt) RoundTrip(request *http.Request) (*http.Response, error) {
	// Trick: append Config OrderrideAcount to eth_call
	rt := request.Clone(context.Background())
	body, _ := ioutil.ReadAll(request.Body)
	_ = request.Body.Close()
	if len(body) > 0 {
		rt.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	var req reqMessage
	if err := json.Unmarshal(body, &req); err == nil {
		if req.Method == "eth_call" {
			req.Params = append(req.Params, r.appendData)
		}
		d2, err := json.Marshal(req)
		if err != nil {
			panic(err)
		}
		rt.ContentLength = int64(len(d2))
		rt.Body = ioutil.NopCloser(bytes.NewBuffer(d2))
	}
	return r.c.Do(rt)
}

func FloatToTokenAmount(amount float64, decimals int64) *big.Int {
	weiFloat := big.NewFloat(amount)
	decimalsBigFloat := big.NewFloat(0).SetInt(Exp10(decimals))
	amountBig := new(big.Float).Mul(weiFloat, decimalsBigFloat)
	r, _ := amountBig.Int(nil)

	return r
}

// Exp10 ...
func Exp10(n int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(expBase), big.NewInt(n), nil)
}

func toHashString(hexStr string) string {
	str := strings.Replace(hexStr, "0x", "", -1)
	str = strings.ToLower(str)
	return fmt.Sprintf("%064v", str)
}

func getIndexBalanceOf(owner string, slot string) common.Hash {
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

func getIndexAllowance(owner string, spender string, slot string) common.Hash {
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
