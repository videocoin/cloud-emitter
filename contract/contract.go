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
	sm "github.com/videocoin/go-protocol/streams"
)

type ClientOpts struct {
	EthClient    *ethclient.Client
	ContractAddr string
	Logger       *logrus.Entry
	Accounts     accountsv1.AccountServiceClient
	ClientSecret string
}

type Client struct {
	logger        *logrus.Entry
	clientSecret  string
	ethClient     *ethclient.Client
	streamManager *sm.StreamManager
	accounts      accountsv1.AccountServiceClient
}

func NewContractClient(opts *ClientOpts) (*Client, error) {
	address := common.HexToAddress(opts.ContractAddr)
	manager, err := sm.NewStreamManager(address, opts.EthClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create smart contract stream manager: %s", err.Error())
	}

	return &Client{
		ethClient:     opts.EthClient,
		logger:        opts.Logger,
		accounts:      opts.Accounts,
		clientSecret:  opts.ClientSecret,
		streamManager: manager,
	}, nil
}

func (c *Client) getClientTransactOpts(ctx context.Context, userID string) (*bind.TransactOpts, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getClientTransactOpts")
	defer span.Finish()

	keyReq := &accountsv1.AccountRequest{OwnerId: userID}
	key, err := c.accounts.Key(ctx, keyReq)
	if err != nil {
		return nil, err
	}

	decrypted, err := keystore.DecryptKey([]byte(key.Key), c.clientSecret)
	if err != nil {
		return nil, err
	}

	transactOpts, err := bcops.GetBCAuth(c.ethClient, decrypted)
	if err != nil {
		return nil, err
	}

	from := common.HexToAddress(key.Address)
	transactOpts.From = from
	transactOpts.GasPrice = big.NewInt(10000000000)

	return transactOpts, nil
}

func (c *Client) getManagerTransactOpts(ctx context.Context, client *ethclient.Client, key []byte, secret string) (*bind.TransactOpts, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "getManagerTransactOpts")
	defer span.Finish()

	decrypted, err := keystore.DecryptKey(key, secret)
	if err != nil {
		return nil, err
	}

	opts, err := bcops.GetBCAuth(client, decrypted)
	if err != nil {
		return nil, err
	}

	return opts, nil
}

func getValidatorTransactOpts(ctx context.Context, client *ethclient.Client, key []byte, secret string) (*bind.TransactOpts, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "getValidatorTransactOpts")
	defer span.Finish()

	decrypted, err := keystore.DecryptKey(key, secret)
	if err != nil {
		return nil, err
	}

	opts, err := bcops.GetBCAuth(client, decrypted)
	if err != nil {
		return nil, err
	}

	return opts, nil
}

func (c *Client) WaitMined(tx *types.Transaction) (*types.Receipt, error) {
	cancelCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	receipt, err := bind.WaitMined(cancelCtx, c.ethClient, tx)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

func (c *Client) WaitMinedAndCheck(tx *types.Transaction) error {
	receipt, err := c.WaitMined(tx)
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
