package service

import (
	"fmt"

	"github.com/videocoin/cloud-pkg/paymentmanager"

	"github.com/ethereum/go-ethereum/common"
	accountsv1 "github.com/videocoin/cloud-api/accounts/v1"
	"github.com/videocoin/cloud-emitter/contract"
	"github.com/videocoin/cloud-emitter/eventbus"
	"github.com/videocoin/cloud-emitter/rpc"
	faucetcli "github.com/videocoin/cloud-pkg/faucet"
	"github.com/videocoin/cloud-pkg/grpcutil"
	"github.com/videocoin/cloud-pkg/rpcutils"
	"github.com/videocoin/go-staking"
)

type Service struct {
	cfg *Config
	rpc *rpc.Server
	eb  *eventbus.EventBus
}

func NewService(cfg *Config) (*Service, error) {
	conn, err := grpcutil.Connect(cfg.AccountsRPCAddr, cfg.Logger.WithField("system", "accountcli"))
	if err != nil {
		return nil, err
	}
	accounts := accountsv1.NewAccountServiceClient(conn)

	ethClient, err := rpcutils.SymphonyRPCClient(cfg.SymphonyAddr, cfg.SymphonyOauthClientID, cfg.SymphonyRPCKey)
	if err != nil {
		return nil, fmt.Errorf("failed to dial eth client: %s", err.Error())
	}

	contractOpts := &contract.ClientOpts{
		EthClient:    ethClient,
		ContractAddr: cfg.StreamManagerContractAddr,
		ClientSecret: cfg.ClientSecret,
		Accounts:     accounts,
	}

	err = contract.LoadKSFromFiles(cfg.ManagersKSPath, cfg.ValidatorsKSPath)
	if err != nil {
		return nil, err
	}

	contract, err := contract.NewContractClient(contractOpts)
	if err != nil {
		return nil, err
	}

	stakingClient, err := staking.NewClient(ethClient, common.HexToAddress(cfg.StakingManagerContractAddr))
	if err != nil {
		return nil, err
	}

	faucet := faucetcli.NewClient(
		fmt.Sprintf("%s/v1/faucet", cfg.SymphonyAddr),
		faucetcli.WithTokenSource(cfg.SymphonyOauthClientID, cfg.SymphonyFaucetKey),
	)

	pm := paymentmanager.NewClient(cfg.PaymentManagerHost)

	rpcConfig := &rpc.ServerOpts{
		Addr:     cfg.RPCAddr,
		Contract: contract,
		Staking:  stakingClient,
		Logger:   cfg.Logger.WithField("system", "rpc"),
		Accounts: accounts,
		Faucet:   faucet,
		PM:       pm,
	}

	rpc, err := rpc.NewRPCServer(rpcConfig)
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
		cfg: cfg,
		rpc: rpc,
		eb:  eb,
	}

	return svc, nil
}

func (s *Service) Start(errCh chan error) {
	go func() {
		s.cfg.Logger.WithField("addr", s.cfg.RPCAddr).Info("starting rpc server")
		errCh <- s.rpc.Start()
	}()

	go func() {
		s.cfg.Logger.Info("starting eventbus")
		errCh <- s.eb.Start()
	}()
}

func (s *Service) Stop() error {
	s.rpc.Stop()
	return s.eb.Stop()
}
