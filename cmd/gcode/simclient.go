package bb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type reqMessage struct {
	JSONRPC string            `json:"jsonrpc"`
	ID      int               `json:"id"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
}

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

func newRoundTripExt(c *http.Client, accounts OverrideAccounts) (http.RoundTripper, error) {
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
	rt := request.Clone(context.Background())
	body, _ := ioutil.ReadAll(request.Body)
	_ = request.Body.Close()
	if len(body) > 0 {
		rt.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	var req reqMessage
	if err := json.Unmarshal(body, &req); err == nil {
		if req.Method == "eth_call" /* && (bytes.Contains(req.Params[0], []byte(`0x99f9fbd2`)) || bytes.Contains(req.Params[0], []byte(`0x110bb26c`)))*/ {
			req.Params = append(req.Params, r.appendData)
		}
		d2, err := json.Marshal(req)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(d2))
		rt.ContentLength = int64(len(d2))
		rt.Body = ioutil.NopCloser(bytes.NewBuffer(d2))
	}
	return r.c.Do(rt)
}

func NewClient(url string, client *http.Client, accounts OverrideAccounts) (*ethclient.Client, error) {
	round, err := newRoundTripExt(client, accounts)
	if err != nil {
		return nil, err
	}
	cc := &http.Client{Transport: round}
	r, err := rpc.DialHTTPWithClient(url, cc)
	if err != nil {
		return nil, err
	}
	ethClient := ethclient.NewClient(r)
	return ethClient, err
}

var (
	// SimUtilAddress = common.HexToAddress("0x1111111111111111111111111111111111111100")
	SimSwapAddress   = common.HexToAddress("0x1111111111111111111111111111111111111101")
	CCallerAddress   = common.HexToAddress("0x1111111111111111111111111111111111111102")
	ProxyCallAddress = common.HexToAddress("0x1111111111111111111111111111111111111103")
	MyWallet         = book[WALLET]
	KNCAddress       = book[KNC]
	PRLAddress       = book[PRL]
	CommonContract   = OverrideAccounts{
		/*SimUtilAddress: {
			Nonce: "0x10",
			Code:  "0x608060405234801561001057600080fd5b506004361061002b5760003560e01c8063ef5bfc3714610030575b600080fd5b6100fc6004803603604081101561004657600080fd5b810190808035906020019064010000000081111561006357600080fd5b82018360208201111561007557600080fd5b8035906020019184602083028401116401000000008311171561009757600080fd5b9091929391929390803590602001906401000000008111156100b857600080fd5b8201836020820111156100ca57600080fd5b803590602001918460208302840111640100000000831117156100ec57600080fd5b9091929391929390505050610153565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561013f578082015181840152602081019050610124565b505050509050019250505060405180910390f35b60608083839050868690500267ffffffffffffffff8111801561017557600080fd5b506040519080825280602002602001820160405280156101a45781602001602082028036833780820191505090505b50905060005b868690508110156103c45760005b858590508110156103b65773eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff168686838181106101fa57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156102975787878381811061023f57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16318382888890508502018151811061028657fe5b6020026020010181815250506103a9565b8585828181106102a357fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a082318989858181106102e757fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561034e57600080fd5b505afa158015610362573d6000803e3d6000fd5b505050506040513d602081101561037857600080fd5b81019080805190602001909291905050508382888890508502018151811061039c57fe5b6020026020010181815250505b80806001019150506101b8565b5080806001019150506101aa565b508091505094935050505056fea26469706673582212200b8738e9ce84a5af761886e464d14ce62b0d8bce0316bc15a43ba74267a846ea64736f6c634300060c0033",
		},*/
		SimSwapAddress: {
			Nonce: "0x10",
			Code:  "0x6080604052600436106100295760003560e01c806321bd38fa1461002e5780637e5465ba1461005b575b600080fd5b61004161003c36600461067f565b610089565b604080519283526020830191909152015b60405180910390f35b34801561006757600080fd5b5061007b610076366004610721565b6103e2565b604051908152602001610052565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152336004820152600090819087908790839073ffffffffffffffffffffffffffffffffffffffff8416906370a0823190602401602060405180830381865afa1580156100fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101229190610754565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815233600482015290915060009073ffffffffffffffffffffffffffffffffffffffff8416906370a0823190602401602060405180830381865afa158015610192573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101b69190610754565b6040517f7e5465ba00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808e1660048301528b1660248201529091503090637e5465ba906044016020604051808303816000875af115801561022d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102519190610754565b506102928989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104aa92505050565b506040517f70a0823100000000000000000000000000000000000000000000000000000000815233600482015260009073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa158015610300573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103249190610754565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815233600482015290915060009073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa158015610394573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b89190610754565b90506103c4828561076d565b97506103d0838261076d565b96505050505050509550959350505050565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f80000000000000000000000000000000000000000000000000000000000000006024830152600091849182169063095ea7b3906044016020604051808303816000875af115801561047b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049f91906107ab565b506000949350505050565b60606104cf838360405180606001604052806027815260200161086b602791396104d6565b9392505050565b606073ffffffffffffffffffffffffffffffffffffffff84163b610581576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60448201527f6e7472616374000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6000808573ffffffffffffffffffffffffffffffffffffffff16856040516105a991906107fd565b600060405180830381855af49150503d80600081146105e4576040519150601f19603f3d011682016040523d82523d6000602084013e6105e9565b606091505b50915091506105f9828286610603565b9695505050505050565b606083156106125750816104cf565b8251156106225782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105789190610819565b803573ffffffffffffffffffffffffffffffffffffffff8116811461067a57600080fd5b919050565b60008060008060006080868803121561069757600080fd5b6106a086610656565b94506106ae60208701610656565b93506106bc60408701610656565b9250606086013567ffffffffffffffff808211156106d957600080fd5b818801915088601f8301126106ed57600080fd5b8135818111156106fc57600080fd5b89602082850101111561070e57600080fd5b9699959850939650602001949392505050565b6000806040838503121561073457600080fd5b61073d83610656565b915061074b60208401610656565b90509250929050565b60006020828403121561076657600080fd5b5051919050565b6000828210156107a6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500390565b6000602082840312156107bd57600080fd5b815180151581146104cf57600080fd5b60005b838110156107e85781810151838201526020016107d0565b838111156107f7576000848401525b50505050565b6000825161080f8184602087016107cd565b9190910192915050565b60208152600082518060208401526108388160408501602087016107cd565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122097159c15cc03bdf22b842e2bcd5bbb948cbb636d7f708b8b20600dccc0bb1d3a64736f6c634300080f0033",
			StateDiff: map[string]string{
				"0x6493c8a5d89f617219e546ba18427d9c4c24a8d987a0365e97e8981e6ed7a4e1": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000",
				"0xff718e3a57f7d0b84395bd64981a58ae6c435a3a95fcd7655f14e5d8a715e05f": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000",
			},
		},
		MyWallet: {
			Balance: "0x8ac7230489e80000",
		}, // 0x6b577f75accc83fade72d8d9776565a84cac37c9530c507115a8960eaa62789e
		KNCAddress: {
			StateDiff: map[string]string{
				"0x6b577f75accc83fade72d8d9776565a84cac37c9530c507115a8960eaa62789e": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000", // balance
				"0x9132bb4fbeedc04775ed2ee59ca1e7627bc052ecc563af16df325f35ea54804f": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000", // allownace
				"0x7e7d35177e80a1d99d70c6ee1eefe7128244a19567487d475dffaed4cc559b1b": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000",
			},
		},
	}
	CommonContractBSC = OverrideAccounts{
		SimSwapAddress: {
			Nonce: "0x10",
			Code:  "0x6080604052600436106100295760003560e01c806321bd38fa1461002e5780637e5465ba1461005b575b600080fd5b61004161003c36600461067f565b610089565b604080519283526020830191909152015b60405180910390f35b34801561006757600080fd5b5061007b610076366004610721565b6103e2565b604051908152602001610052565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152336004820152600090819087908790839073ffffffffffffffffffffffffffffffffffffffff8416906370a0823190602401602060405180830381865afa1580156100fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101229190610754565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815233600482015290915060009073ffffffffffffffffffffffffffffffffffffffff8416906370a0823190602401602060405180830381865afa158015610192573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101b69190610754565b6040517f7e5465ba00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808e1660048301528b1660248201529091503090637e5465ba906044016020604051808303816000875af115801561022d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102519190610754565b506102928989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104aa92505050565b506040517f70a0823100000000000000000000000000000000000000000000000000000000815233600482015260009073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa158015610300573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103249190610754565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815233600482015290915060009073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa158015610394573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b89190610754565b90506103c4828561076d565b97506103d0838261076d565b96505050505050509550959350505050565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f80000000000000000000000000000000000000000000000000000000000000006024830152600091849182169063095ea7b3906044016020604051808303816000875af115801561047b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049f91906107ab565b506000949350505050565b60606104cf838360405180606001604052806027815260200161086b602791396104d6565b9392505050565b606073ffffffffffffffffffffffffffffffffffffffff84163b610581576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60448201527f6e7472616374000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6000808573ffffffffffffffffffffffffffffffffffffffff16856040516105a991906107fd565b600060405180830381855af49150503d80600081146105e4576040519150601f19603f3d011682016040523d82523d6000602084013e6105e9565b606091505b50915091506105f9828286610603565b9695505050505050565b606083156106125750816104cf565b8251156106225782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105789190610819565b803573ffffffffffffffffffffffffffffffffffffffff8116811461067a57600080fd5b919050565b60008060008060006080868803121561069757600080fd5b6106a086610656565b94506106ae60208701610656565b93506106bc60408701610656565b9250606086013567ffffffffffffffff808211156106d957600080fd5b818801915088601f8301126106ed57600080fd5b8135818111156106fc57600080fd5b89602082850101111561070e57600080fd5b9699959850939650602001949392505050565b6000806040838503121561073457600080fd5b61073d83610656565b915061074b60208401610656565b90509250929050565b60006020828403121561076657600080fd5b5051919050565b6000828210156107a6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500390565b6000602082840312156107bd57600080fd5b815180151581146104cf57600080fd5b60005b838110156107e85781810151838201526020016107d0565b838111156107f7576000848401525b50505050565b6000825161080f8184602087016107cd565b9190910192915050565b60208152600082518060208401526108388160408501602087016107cd565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122097159c15cc03bdf22b842e2bcd5bbb948cbb636d7f708b8b20600dccc0bb1d3a64736f6c634300080f0033",
		},
		CCallerAddress: {
			Nonce: "0x10",
			Code:  "0x6080604052348015600f57600080fd5b506004361060285760003560e01c8063fe86f4aa14602d575b600080fd5b6040805133815290519081900360200190f3fea2646970667358221220057f2b15859fa57a833714824211c00847818d25b0a461344a8767131bc74b6b64736f6c634300080f0033",
		},
		ProxyCallAddress: {
			Nonce: "0x10",
			Code:  "0x608060405234801561001057600080fd5b506004361061002b5760003560e01c80632a31f6b414610030575b600080fd5b61004361003e366004610245565b610059565b604051610050919061035d565b60405180910390f35b606061009b8484848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506100a592505050565b90505b9392505050565b606061009e838360405180606001604052806027815260200161038d60279139606073ffffffffffffffffffffffffffffffffffffffff84163b610170576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60448201527f6e7472616374000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6000808573ffffffffffffffffffffffffffffffffffffffff16856040516101989190610370565b600060405180830381855af49150503d80600081146101d3576040519150601f19603f3d011682016040523d82523d6000602084013e6101d8565b606091505b50915091506101e88282866101f2565b9695505050505050565b6060831561020157508161009e565b8251156102115782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610167919061035d565b60008060006040848603121561025a57600080fd5b833573ffffffffffffffffffffffffffffffffffffffff8116811461027e57600080fd5b9250602084013567ffffffffffffffff8082111561029b57600080fd5b818601915086601f8301126102af57600080fd5b8135818111156102be57600080fd5b8760208285010111156102d057600080fd5b6020830194508093505050509250925092565b60005b838110156102fe5781810151838201526020016102e6565b8381111561030d576000848401525b50505050565b6000815180845261032b8160208601602086016102e3565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061009e6020830184610313565b600082516103828184602087016102e3565b919091019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220df9fc23c6dc4d21173f0b22e85887c5104650d5b4e58122b5cac93b6aee994fe64736f6c634300080f0033",
		},
		MyWallet: {
			Balance: "0x8ac7230489e80000",
		}, // 0x6b577f75accc83fade72d8d9776565a84cac37c9530c507115a8960eaa62789e
		KNCAddress: {
			StateDiff: map[string]string{
				// "0x6493c8a5d89f617219e546ba18427d9c4c24a8d987a0365e97e8981e6ed7a4e1": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000",
				// "0xff718e3a57f7d0b84395bd64981a58ae6c435a3a95fcd7655f14e5d8a715e05f": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000",

				// "0xa31cdb984800dfe48437c0b4134530cbb78aa3f315acbaa9e3c12e461c077ea9": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000", // balances
				// "0x4c4982cb44f65c21ffdf6e89fe5d2b2b62b528fff53914517c8a1dcda766d14b": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000", // knc on router
				// "0x921b2cc64281b1c695bde5cdae2b887532b87baf35e6c2ae3d3edb7f56d7c9c7": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000", // knc on sim

				"0x3a3fdf7724b95c02a5edd803731f08c8d89aba3eaa11b36a629f6da7093bb414": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000",
				// "0xbe3685b6c3af032e4804e56c7cab23283af73b6fdfcc14c61f229a944cf4b270": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000",
				"0x934bafe733d669269688fb8bf1035bb224f4842ce37f1ef18f8bd76233a30e10": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000",
			},
		},
	}
)
