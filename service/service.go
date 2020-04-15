package service

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	accountsv1 "github.com/videocoin/cloud-api/accounts/v1"
	streamsv1 "github.com/videocoin/cloud-api/streams/v1"
	"github.com/videocoin/cloud-emitter/contract"
	"github.com/videocoin/cloud-emitter/eventbus"
	"github.com/videocoin/cloud-emitter/manager"
	"github.com/videocoin/cloud-emitter/rpc"
	faucetcli "github.com/videocoin/cloud-pkg/faucet"
	"github.com/videocoin/cloud-pkg/grpcutil"
	"github.com/videocoin/go-staking"
)

type Service struct {
	cfg     *Config
	rpc     *rpc.Server
	manager *manager.Manager
	eb      *eventbus.EventBus
}

func NewService(cfg *Config) (*Service, error) {
	conn, err := grpcutil.Connect(cfg.AccountsRPCAddr, cfg.Logger.WithField("system", "accountcli"))
	if err != nil {
		return nil, err
	}
	accounts := accountsv1.NewAccountServiceClient(conn)

	conn, err = grpcutil.Connect(cfg.StreamsRPCAddr, cfg.Logger.WithField("system", "streamscli"))
	if err != nil {
		return nil, err
	}
	streams := streamsv1.NewStreamServiceClient(conn)

	ethClient, err := ethclient.Dial(cfg.RPCNodeHTTPAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial eth client: %s", err.Error())
	}

	contractOpts := &contract.ClientOpts{
		EthClient:       ethClient,
		ContractAddr:    cfg.StreamManagerContractAddr,
		Accounts:        accounts,
		ClientSecret:    cfg.ClientSecret,
		ManagerKey:      cfg.ManagerKey,
		ManagerSecret:   cfg.ManagerSecret,
		ValidatorKey:    cfg.ValidatorKey,
		ValidatorSecret: cfg.ValidatorSecret,
		Logger:          cfg.Logger.WithField("system", "contract"),
	}

	contract, err := contract.NewContractClient(contractOpts)
	if err != nil {
		return nil, err
	}

	stakingClient, err := staking.NewClient(ethClient, common.HexToAddress(cfg.StakingManagerContractAddr))
	if err != nil {
		return nil, err
	}

	rpcConfig := &rpc.ServerOpts{
		Addr:     cfg.RPCAddr,
		Streams:  streams,
		Contract: contract,
		Staking:  stakingClient,
		Logger:   cfg.Logger.WithField("system", "rpc"),
	}

	rpc, err := rpc.NewRPCServer(rpcConfig)
	if err != nil {
		return nil, err
	}

	faucet := faucetcli.NewClient(cfg.FaucetURL)
	managerOpts := &manager.Opts{
		Logger:   cfg.Logger.WithField("system", "manager"),
		Faucet:   faucet,
		Accounts: accounts,
		Contract: contract,
	}
	manager, err := manager.NewManager(managerOpts)
	if err != nil {
		return nil, err
	}

	ebConfig := &eventbus.Config{
		URI:    cfg.MQURI,
		Name:   cfg.Name,
		Logger: cfg.Logger.WithField("system", "eventbus"),
		Faucet: faucet,
	}
	eb, err := eventbus.New(ebConfig)
	if err != nil {
		return nil, err
	}

	svc := &Service{
		cfg:     cfg,
		rpc:     rpc,
		manager: manager,
		eb:      eb,
	}

	return svc, nil
}

func (s *Service) Start(errCh chan error) {
	go func() {
		s.cfg.Logger.Info("starting rpc server")
		errCh <- s.rpc.Start()
	}()

	go func() {
		s.cfg.Logger.Info("starting eventbus")
		errCh <- s.eb.Start()
	}()

	s.manager.StartBackgroundTasks()
}

func (s *Service) Stop() error {
	s.manager.StopBackgroundTasks()
	return s.eb.Stop()
}
