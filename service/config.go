package service

import (
	"github.com/sirupsen/logrus"
)

// Config primary config for emitter
type Config struct {
	Name    string        `envconfig:"-"`
	Version string        `envconfig:"-"`
	Logger  *logrus.Entry `envconfig:"-"`

	RPCAddr                    string `envconfig:"RPC_ADDR" default:"0.0.0.0:5003"`
	AccountsRPCAddr            string `envconfig:"ACCOUNTS_RPC_ADDR" default:"0.0.0.0:5001"`
	MQURI                      string `envconfig:"MQURI" default:"amqp://guest:guest@127.0.0.1:5672"`
	FaucetAddr                 string `envconfig:"FAUCET_ADDR" required:"true"`
	SymphonyAddr               string `envconfig:"SYMPHONY_ADDR" required:"true"`
	SymphonyOauthClientID      string `envconfig:"SYMPHONY_OAUTH_CLIENT_ID" required:"true"`
	SymphonyFaucetKey          string `envconfig:"SYMPHONY_FAUCET_KEY" required:"true"`
	PaymentManagerHost         string `envconfig:"PAYMENT_MANAGER_HOST" required:"true"`
	StreamManagerContractAddr  string `envconfig:"STREAM_MANAGER_CONTRACT_ADDR" required:"true"`
	StakingManagerContractAddr string `envconfig:"STAKING_MANAGER_CONTRACT_ADDR" required:"true"`
	ClientSecret               string `envconfig:"CLIENT_SECRET" required:"true"`
	ManagersKSPath             string `envconfig:"MANAGERS_KS_PATH" default:"/vault/secrets/sa-managers"`
	ValidatorsKSPath           string `envconfig:"VALIDATORS_KS_PATH" default:"/vault/secrets/sa-validators"`
}
