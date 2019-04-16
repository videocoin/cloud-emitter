package service

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	Name    string `envconfig:"-"`
	Version string `envconfig:"-"`

	RPCAddr         string `default:"0.0.0.0:5000"`
	AccountsRPCAddr string `default:"0.0.0.0:5001"`

	KeyManager       string
	SecretKeyManager string
	SecretKeyClient  string

	MQURI string `default:"amqp://guest:guest@127.0.0.1:5672"`

	Secret string `default:"secret"`

	Logger *logrus.Entry `envconfig:"-"`
}
