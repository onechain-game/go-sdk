package client

import (
	"github.com/onechain-game/go-sdk/client/basic"
	"github.com/onechain-game/go-sdk/client/query"
	"github.com/onechain-game/go-sdk/client/transaction"
	"github.com/onechain-game/go-sdk/common/types"
	"github.com/onechain-game/go-sdk/keys"
)

// dexClient wrapper
type dexClient struct {
	query.QueryClient
	transaction.TransactionClient
	basic.BasicClient
}

// DexClient methods
type DexClient interface {
	basic.BasicClient
	query.QueryClient
	transaction.TransactionClient
}

func init() {
	resty.DefaultClient.SetRedirectPolicy(resty.FlexibleRedirectPolicy(10))
}

func NewDexClientWithApiKey(baseUrl string, network types.ChainNetwork, keyManager keys.KeyManager, apiKey string) (DexClient, error) {
	types.SetNetwork(network)
	c := basic.NewClient(baseUrl+"/internal", apiKey)
	q := query.NewClient(c)
	n, err := q.GetNodeInfo()
	if err != nil {
		return nil, err
	}
	t := transaction.NewClient(n.NodeInfo.Network, keyManager, q, c)
	return &dexClient{BasicClient: c, QueryClient: q, TransactionClient: t}, nil
}

func NewDexClient(baseUrl string, network types.ChainNetwork, keyManager keys.KeyManager) (DexClient, error) {
	types.SetNetwork(network)
	c := basic.NewClient(baseUrl, "")
	q := query.NewClient(c)
	n, err := q.GetNodeInfo()
	if err != nil {
		return nil, err
	}
	t := transaction.NewClient(n.NodeInfo.Network, keyManager, q, c)
	return &dexClient{BasicClient: c, QueryClient: q, TransactionClient: t}, nil
}
