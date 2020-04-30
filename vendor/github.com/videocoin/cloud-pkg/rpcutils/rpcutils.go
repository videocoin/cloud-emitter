package rpcutils

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SymphonyRPCClient(symphonyEndpoint, oauthClientID, accessToken string) (*ethclient.Client, error) {
	c, err := google.JWTConfigFromJSON([]byte(accessToken))
	if err != nil {
		return nil, err
	}

	c.UseIDToken = true
	c.PrivateClaims = map[string]interface{}{
		"target_audience": oauthClientID,
	}

	rpcClient, err := rpc.DialHTTPWithClient(
		symphonyEndpoint,
		oauth2.NewClient(context.Background(), c.TokenSource(context.Background())))
	if err != nil {
		return nil, err
	}

	return ethclient.NewClient(rpcClient), nil
}
