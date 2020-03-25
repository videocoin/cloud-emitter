package contract

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	accountsv1 "github.com/videocoin/cloud-api/accounts/v1"
	"github.com/videocoin/cloud-pkg/bcops"
	sm "github.com/videocoin/cloud-pkg/streamManager"
)

type ClientOpts struct {
	RPCNodeHTTPAddr string
	ContractAddr    string
	Logger          *logrus.Entry
	Accounts        accountsv1.AccountServiceClient
	ClientSecret    string
	ManagerKey      string
	ManagerSecret   string
}

type Client struct {
	accounts      accountsv1.AccountServiceClient
	ethClient     *ethclient.Client
	streamManager *sm.Manager

	clientSecret  string
	managerKey    string
	managerSecret string

	logger *logrus.Entry
}

func NewContractClient(opts *ClientOpts) (*Client, error) {
	ethClient, err := ethclient.Dial(opts.RPCNodeHTTPAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial eth client: %s", err.Error())
	}

	address := common.HexToAddress(opts.ContractAddr)
	manager, err := sm.NewManager(address, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create smart contract stream manager: %s", err.Error())
	}

	return &Client{
		accounts:      opts.Accounts,
		ethClient:     ethClient,
		streamManager: manager,
		clientSecret:  opts.ClientSecret,
		managerKey:    opts.ManagerKey,
		managerSecret: opts.ManagerSecret,
		logger:        opts.Logger,
	}, nil
}

func (c *Client) getClientTransactOpts(ctx context.Context, userID string) (*bind.TransactOpts, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getClientTransactOpts")
	defer span.Finish()

	keyReq := &accountsv1.AccountRequest{OwnerId: userID}
	key, err := c.accounts.Key(ctx, keyReq)
	if err != nil {
		c.logger.WithError(err).WithField("user_id", userID).Error("failed to get account key")
		return nil, err
	}

	decrypted, err := keystore.DecryptKey([]byte(key.Key), c.clientSecret)
	if err != nil {
		c.logger.WithError(err).WithField("key", []byte(key.Key)).Error("failed to decrypt key")
		return nil, err
	}

	transactOpts, err := bcops.GetBCAuth(c.ethClient, decrypted)
	if err != nil {
		c.logger.WithError(err).Error("failed to get bc auth")
		return nil, err
	}

	from := common.HexToAddress(key.Address)
	transactOpts.From = from
	transactOpts.GasPrice = big.NewInt(10000000000)

	return transactOpts, nil
}

func (c *Client) getManagerTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "getManagerTransactOpts")
	defer span.Finish()

	decrypted, err := keystore.DecryptKey([]byte(c.managerKey), c.managerSecret)
	if err != nil {
		c.logger.WithError(err).WithField("key", []byte(c.managerKey)).Error("failed to decrypt key")
		return nil, err
	}

	transactOpts, err := bcops.GetBCAuth(c.ethClient, decrypted)
	if err != nil {
		c.logger.WithError(err).Error("failed to get bc auth")
		return nil, err
	}

	return transactOpts, nil
}

func (c *Client) WaitMinedAndCheck(tx *types.Transaction) error {
	cancelCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	receipt, err := bind.WaitMined(cancelCtx, c.ethClient, tx)
	if err != nil {
		return err
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("transaction %s failed", tx.Hash().String())
	}

	return nil
}

func (c *Client) EthClient() *ethclient.Client {
	return c.ethClient
}
