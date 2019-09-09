package service

import (
	accountsv1 "github.com/videocoin/cloud-api/accounts/v1"
	streamsv1 "github.com/videocoin/cloud-api/streams/v1"
	"github.com/videocoin/cloud-emitter/contract"
	"github.com/videocoin/cloud-emitter/rpc"
	"github.com/videocoin/cloud-pkg/grpcutil"
	"google.golang.org/grpc"
)

type Service struct {
	cfg *Config
	rpc *rpc.RpcServer
}

func NewService(cfg *Config) (*Service, error) {
	alogger := cfg.Logger.WithField("system", "accountcli")
	aGrpcDialOpts := grpcutil.ClientDialOptsWithRetry(alogger)
	accountsConn, err := grpc.Dial(cfg.AccountsRPCAddr, aGrpcDialOpts...)
	if err != nil {
		return nil, err
	}
	accounts := accountsv1.NewAccountServiceClient(accountsConn)

	slogger := cfg.Logger.WithField("system", "streamscli")
	sGrpcDialOpts := grpcutil.ClientDialOptsWithRetry(slogger)
	streamsConn, err := grpc.Dial(cfg.StreamsRPCAddr, sGrpcDialOpts...)
	if err != nil {
		return nil, err
	}
	streams := streamsv1.NewStreamServiceClient(streamsConn)

	contractOpts := &contract.ContractClientOpts{
		RPCNodeHTTPAddr: cfg.RPCNodeHTTPAddr,
		ContractAddr:    cfg.StreamManagerContractAddr,
		Accounts:        accounts,
		ClientSecret:    cfg.ClientSecret,
		ManagerKey:      cfg.ManagerKey,
		ManagerSecret:   cfg.ManagerSecret,
		Logger:          cfg.Logger.WithField("system", "contract"),
	}

	contract, err := contract.NewContractClient(contractOpts)
	if err != nil {
		return nil, err
	}

	rpcConfig := &rpc.RpcServerOpts{
		Addr:     cfg.RPCAddr,
		Streams:  streams,
		Contract: contract,
		Logger:   cfg.Logger.WithField("system", "rpc"),
	}

	rpc, err := rpc.NewRpcServer(rpcConfig)
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
	go s.rpc.Stop()
	return nil
}
