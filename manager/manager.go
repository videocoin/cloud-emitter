package manager

import (
	"context"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	protoempty "github.com/gogo/protobuf/types"
	"github.com/sirupsen/logrus"
	accountsv1 "github.com/videocoin/cloud-api/accounts/v1"
	"github.com/videocoin/cloud-emitter/contract"
	"github.com/videocoin/cloud-pkg/ethutils"
	faucetcli "github.com/videocoin/cloud-pkg/faucet"
)

type Opts struct {
	Logger   *logrus.Entry
	Faucet   *faucetcli.Client
	Accounts accountsv1.AccountServiceClient
	Contract *contract.Client
}

type Manager struct {
	logger   *logrus.Entry
	bTicker  *time.Ticker
	bTimeout time.Duration
	faucet   *faucetcli.Client
	accounts accountsv1.AccountServiceClient
	contract *contract.Client
}

func NewManager(opts *Opts) (*Manager, error) {
	bTimeout := 60 * time.Second

	return &Manager{
		logger:   opts.Logger,
		bTimeout: bTimeout,
		bTicker:  time.NewTicker(bTimeout),
		faucet:   opts.Faucet,
		accounts: opts.Accounts,
		contract: opts.Contract,
	}, nil
}

func (m *Manager) StartBackgroundTasks() {
	m.logger.Info("starting background tasks")
	go m.checkBalanceTask()
}

func (m *Manager) StopBackgroundTasks() {
	m.bTicker.Stop()
}

func (m *Manager) checkBalanceTask() {
	for range m.bTicker.C {
		ctx := context.Background()
		accounts, err := m.accounts.List(ctx, &protoempty.Empty{})
		if err != nil {
			m.logger.Errorf("failed to get accounts list: %s", err)
			continue
		}

		for _, account := range accounts.Items {
			logger := m.logger.WithField("address", account.Address)
			logger.Info("checking balance")

			balanceWei, err := m.contract.EthClient().BalanceAt(ctx, ethcommon.HexToAddress(account.Address), nil)
			if err != nil {
				logger.Errorf("failed to balance at: %s", err)
				continue
			}

			balanceVID, err := ethutils.WeiToEth(balanceWei)
			if err != nil {
				logger.Errorf("failed to wei to eth: %s", err)
				continue
			}

			balance, _ := balanceVID.Float64()
			logger.Infof("balance is %f", balance)

			if balance >= 10 {
				continue
			}

			err = m.faucet.Do(account.Address, 1)
			if err != nil {
				m.logger.Errorf("failed to faucet: %s", err)
				continue
			}
		}
	}
}
