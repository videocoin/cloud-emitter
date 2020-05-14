package contract

import (
	"encoding/base64"
	"fmt"
	"strings"

	consulapi "github.com/hashicorp/consul/api"
)

func LoadKSFromConsul(addr string, env string) error {
	cfg := consulapi.DefaultConfig()
	cfg.Address = addr

	consul, err := consulapi.NewClient(cfg)
	if err != nil {
		return fmt.Errorf("consul: failed to create client: %s", err)
	}

	_, err = consul.Agent().Self()
	if err != nil {
		return fmt.Errorf("consul: failed to connect agent: %s", err)
	}

	pairs, _, err := consul.KV().List(fmt.Sprintf("config/%s/services/emitter/managers", env), nil)
	if err != nil {
		return err
	}

	for _, pair := range pairs {
		path := strings.Split(pair.Key, "/")

		secret, err := base64.RawStdEncoding.DecodeString(path[len(path)-1])
		if err != nil {
			return err
		}

		item := &KSItem{
			Key:    string(pair.Value),
			Secret: string(secret),
		}

		managerKS = append(managerKS, item)
	}

	pairs, _, err = consul.KV().List(fmt.Sprintf("config/%s/services/emitter/validators", env), nil)
	if err != nil {
		return err
	}

	for _, pair := range pairs {
		path := strings.Split(pair.Key, "/")

		secret, err := base64.RawStdEncoding.DecodeString(path[len(path)-1])
		if err != nil {
			return err
		}

		item := &KSItem{
			Key:    string(pair.Value),
			Secret: string(secret),
		}

		validatorKS = append(validatorKS, item)
	}

	managerKS = shuffle(managerKS)
	validatorKS = shuffle(validatorKS)
	managerKSPool, _ = NewRoundRobin(managerKS)
	validatorKSPool, _ = NewRoundRobin(validatorKS)

	return nil
}
