package rpcutils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func SymphonyRPCClient(endpoint string) (*ethclient.Client, error) {
	rpcClient, err := rpc.Dial(endpoint)
	if err != nil {
		return nil, err
	}

	return ethclient.NewClient(rpcClient), nil
}
