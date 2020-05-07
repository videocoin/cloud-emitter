package service

import (
	"github.com/sirupsen/logrus"
)

// Config primary config for emitter
type Config struct {
	Name    string        `envconfig:"-"`
	Version string        `envconfig:"-"`
	Logger  *logrus.Entry `envconfig:"-"`

	RPCAddr         string `envconfig:"RPC_ADDR" default:"0.0.0.0:5003"`
	AccountsRPCAddr string `envconfig:"ACCOUNTS_RPC_ADDR" default:"0.0.0.0:5001"`
	MQURI           string `envconfig:"MQURI" default:"amqp://guest:guest@127.0.0.1:5672"`

	SymphonyAddr          string `envconfig:"SYMPHONY_ADDR" required:"true"`
	SymphonyOauthClientID string `envconfig:"SYMPHONY_OAUTH_CLIENT_ID" required:"true"`
	SymphonyRPCKey        string `envconfig:"SYMPHONY_RPC_KEY" required:"true"`
	SymphonyFaucetKey     string `envconfig:"SYMPHONY_FAUCET_KEY" required:"true"`

	StreamManagerContractAddr  string `envconfig:"STREAM_MANAGER_CONTRACT_ADDR" required:"true"`
	StakingManagerContractAddr string `envconfig:"STAKING_MANAGER_CONTRACT_ADDR" required:"true"`

	ClientSecret string `envconfig:"CLIENT_SECRET" required:"true"`
}
