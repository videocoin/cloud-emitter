package service

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	Name    string `envconfig:"-"`
	Version string `envconfig:"-"`

	RPCAddr         string `default:"0.0.0.0:5003"`
	AccountsRPCAddr string `default:"0.0.0.0:5001"`

	MQURI string `default:"amqp://guest:guest@127.0.0.1:5672"`

	NodeRPCAddr     string `default:""`
	ContractAddress string `default:"0x8bFEF9CaB41460dB146b1E701B782edBa7883bC9"`

	Secret string `default:""`

	MKey    string `default:""`
	MSecret string `default:""`

	Logger *logrus.Entry `envconfig:"error"`
}
