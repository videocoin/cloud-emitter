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
	StreamsRPCAddr  string `envconfig:"STREAMS_RPC_ADDR" default:"0.0.0.0:5002"`
	MQURI           string `envconfig:"MQURI" default:"amqp://guest:guest@127.0.0.1:5672"`

	RPCNodeHTTPAddr           string `envconfig:"RPC_NODE_HTTP_ADDR" required:"true"`
	StreamManagerContractAddr string `envconfig:"STREAM_MANAGER_CONTRACT_ADDR" required:"true"`
	FaucetURL                 string `envconfig:"FAUCET_URL" required:"true"`

	ManagerKey      string `envconfig:"MANAGER_KEY" required:"true"`
	ManagerSecret   string `envconfig:"MANAGER_SECRET" required:"true"`
	ClientSecret    string `envconfig:"CLIENT_SECRET" required:"true"`
	ValidatorKey    string `envconfig:"VALIDATOR_KEY" required:"true"`
	ValidatorSecret string `envconfig:"VALIDATOR_SECRET" required:"true"`
}
