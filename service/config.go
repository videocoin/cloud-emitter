package service

import (
	"github.com/sirupsen/logrus"
)

// Config primary config for emitter
type Config struct {
	Name    string `envconfig:"-"`
	Version string `envconfig:"-"`

	RPCAddr         string `default:"0.0.0.0:5003" envconfig:"RPC_ADDR"`
	AccountsRPCAddr string `default:"0.0.0.0:5001" envconfig:"ACCOUNTS_RPC_ADDR"`

	RPCNodeHTTPAddr           string `default:"" envconfig:"RPC_NODE_HTTP_ADDR"`
	StreamManagerContractAddr string `default:"" envconfig:"STREAM_MANAGER_CONTRACT_ADDR"`

	ManagerKey    string `default:"" envconfig:"MANAGER_KEY"`
	ManagerSecret string `default:"" envconfig:"MANAGER_SECRET"`
	ClientSecret  string `default:"" envconfig:"CLIENT_SECRET"`

	Logger *logrus.Entry `envconfig:"-"`
}
