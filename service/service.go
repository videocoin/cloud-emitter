package service

import (
	accountsv1 "github.com/videocoin/cloud-api/accounts/v1"
	"github.com/videocoin/cloud-pkg/grpcutil"
	"google.golang.org/grpc"
)

type Service struct {
	cfg *Config
	rpc *RpcServer
}

func NewService(cfg *Config) (*Service, error) {
	alogger := cfg.Logger.WithField("system", "accountcli")
	aGrpcDialOpts := grpcutil.ClientDialOptsWithRetry(alogger)
	accountsConn, err := grpc.Dial(cfg.AccountsRPCAddr, aGrpcDialOpts...)
	if err != nil {
		return nil, err
	}

	accounts := accountsv1.NewAccountServiceClient(accountsConn)

	rpcConfig := &RpcServerOptions{
		Addr:            cfg.RPCAddr,
		RPCNodeHTTPAddr: cfg.RPCNodeHTTPAddr,
		ContractAddr:    cfg.StreamManagerContractAddr,
		Logger:          cfg.Logger,
		Accounts:        accounts,
		ClientSecret:    cfg.ClientSecret,
		ManagerKey:      cfg.ManagerKey,
		ManagerSecret:   cfg.ManagerSecret,
	}

	rpc, err := NewRpcServer(rpcConfig)
	if err != nil {
		return nil, err
	}

	svc := &Service{
		cfg: cfg,
		rpc: rpc,
	}

	return svc, nil
}

func (s *Service) Start() error {
	go s.rpc.Start()
	return nil
}

func (s *Service) Stop() error {
	return nil
}
