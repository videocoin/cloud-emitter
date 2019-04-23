package service

import (
	accountsv1 "github.com/VideoCoin/cloud-api/accounts/v1"
	"github.com/VideoCoin/cloud-pkg/grpcutil"
	"github.com/VideoCoin/cloud-pkg/mqmux"
	"google.golang.org/grpc"
)

type Service struct {
	cfg *Config
	rpc *RpcServer
	eb  *EventBus
}

func NewService(cfg *Config) (*Service, error) {
	alogger := cfg.Logger.WithField("system", "accountcli")
	aGrpcDialOpts := grpcutil.ClientDialOptsWithRetry(alogger)
	accountsConn, err := grpc.Dial(cfg.AccountsRPCAddr, aGrpcDialOpts...)
	if err != nil {
		return nil, err
	}

	accounts := accountsv1.NewAccountServiceClient(accountsConn)

	mq, err := mqmux.NewWorkerMux(cfg.MQURI, cfg.Name)
	if err != nil {
		return nil, err
	}
	mq.Logger = cfg.Logger.WithField("system", "mq")

	eblogger := cfg.Logger.WithField("system", "eventbus")
	eb, err := NewEventBus(mq, eblogger)
	if err != nil {
		return nil, err
	}

	rpcConfig := &RpcServerOptions{
		Addr:            cfg.RPCAddr,
		NodeRPCAddr:     cfg.NodeRPCAddr,
		ContractAddress: cfg.ContractAddress,
		Logger:          cfg.Logger,
		EB:              eb,
		Accounts:        accounts,
		Secret:          cfg.Secret,
		MKey:            cfg.MKey,
		MSecret:         cfg.MSecret,
	}

	rpc, err := NewRpcServer(rpcConfig)
	if err != nil {
		return nil, err
	}

	svc := &Service{
		cfg: cfg,
		rpc: rpc,
		eb:  eb,
	}

	return svc, nil
}

func (s *Service) Start() error {
	go s.rpc.Start()
	go s.eb.Start()
	return nil
}

func (s *Service) Stop() error {
	s.eb.Stop()
	return nil
}
