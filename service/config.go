package service

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	Name    string `envconfig:"-"`
	Version string `envconfig:"-"`

	RPCAddr         string `default:"0.0.0.0:5003"`
	AccountsRPCAddr string `default:"0.0.0.0:5001"`

	MQURI string `default:"amqp://guest:guest@127.0.0.1:5672" envconfig:"MQURI"`

	NodeHTTPAddr string `default:"" envconfig:"NODEHTTPADDR"`
	ContractAddr string `default:"" envconfig:"CONTRACTADDR"`

	MKey    string `default:""`
	MSecret string `default:""`

	Secret string `default:""`

	Logger *logrus.Entry `envconfig:"-"`
}
