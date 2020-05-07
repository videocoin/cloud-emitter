package contract

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/opentracing/opentracing-go"
	"github.com/videocoin/cloud-pkg/ethutils"
	"github.com/videocoin/go-protocol/streams"
)

func (c *Client) RequestStream(ctx context.Context, userID string, streamID *big.Int, profileNames []string) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RequestStream")
	defer span.Finish()

	transactOpts, err := c.getClientTransactOpts(ctx, userID)
	if err != nil {
		return nil, err
	}

	tx, err := c.streamManager.RequestStream(transactOpts, streamID, profileNames)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) GetStreamAddress(ctx context.Context, streamID *big.Int) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetStreamAddress")
	defer span.Finish()

	s, err := c.streamManager.Requests(new(bind.CallOpts), streamID)
	if err != nil {
		return "", err
	}

	return s.Stream.Hex(), nil
}

func (c *Client) ApproveStream(ctx context.Context, streamID *big.Int) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ApproveStream")
	defer span.Finish()

	k, s := GetManagerKS()
	opts, err := c.getManagerTransactOpts(ctx, c.ethClient, k, s)
	if err != nil {
		return nil, err
	}

	tx, err := c.streamManager.ApproveStreamCreation(opts, streamID)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) CreateStream(ctx context.Context, userID string, streamID *big.Int) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateStream")
	defer span.Finish()

	opts, err := c.getClientTransactOpts(ctx, userID)
	if err != nil {
		return nil, err
	}

	opts.Value = ethutils.EthToWei(5)

	tx, err := c.streamManager.CreateStream(opts, streamID)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) EndStream(ctx context.Context, userID string, streamID *big.Int) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EndStream")
	defer span.Finish()

	opts, err := c.getClientTransactOpts(ctx, userID)
	if err != nil {
		return nil, err
	}

	tx, err := c.streamManager.EndStream(opts, streamID)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) AllowRefund(ctx context.Context, streamID *big.Int) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AllowRefund")
	defer span.Finish()

	k, s := GetManagerKS()
	opts, err := c.getManagerTransactOpts(ctx, c.ethClient, k, s)
	if err != nil {
		return nil, err
	}

	tx, err := c.streamManager.AllowRefund(opts, streamID)
	if err != nil {
		return nil, nil
	}

	return tx, nil
}

func (c *Client) EscrowRefund(ctx context.Context, streamContractAddress string) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EscrowRefund")
	defer span.Finish()

	k, s := GetManagerKS()
	opts, err := c.getManagerTransactOpts(ctx, c.ethClient, k, s)
	if err != nil {
		return nil, err
	}

	stream, err := streams.NewStream(common.HexToAddress(streamContractAddress), c.ethClient)
	if err != nil {
		return nil, err
	}

	tx, err := stream.Refund(opts)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) AddInputChunkID(ctx context.Context, streamID, chunkID *big.Int, rewards []*big.Int) (*types.Transaction, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "addInputchunkId")
	defer span.Finish()

	k, s := GetManagerKS()
	opts, err := c.getManagerTransactOpts(ctx, c.ethClient, k, s)
	if err != nil {
		return nil, err
	}

	tx, err := c.streamManager.AddInputChunkId(opts, streamID, chunkID, rewards)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) Deposit(ctx context.Context, userID string, to, value *big.Int) (*types.Transaction, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Deposit")
	defer span.Finish()

	opts, err := c.getClientTransactOpts(ctx, userID)
	if err != nil {
		return nil, err
	}
	opts.Value = value

	s, err := streams.NewStream(common.BigToAddress(to), c.ethClient)
	if err != nil {
		return nil, err
	}

	tx, err := s.Deposit(opts)
	if err != nil {
		return tx, err
	}

	return tx, nil
}

func (c *Client) ValidateProof(ctx context.Context, streamContractAddress string, profileID, chunkID *big.Int) (*types.Transaction, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ValidateProof")
	defer span.Finish()

	stream, err := streams.NewStream(common.HexToAddress(streamContractAddress), c.ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new stream: %s", err.Error())
	}

	k, s := GetValidatorKS()
	opts, err := getValidatorTransactOpts(ctx, c.ethClient, k, s)
	if err != nil {
		return nil, fmt.Errorf("failed to get transact opts: %s", err.Error())
	}

	tx, err := stream.ValidateProof(opts, profileID, chunkID)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) ScrapProof(ctx context.Context, streamContractAddress string, profileID, chunkID *big.Int) (*types.Transaction, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ScrapProof")
	defer span.Finish()

	stream, err := streams.NewStream(common.HexToAddress(streamContractAddress), c.ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new stream: %s", err.Error())
	}

	k, s := GetValidatorKS()
	opts, err := getValidatorTransactOpts(ctx, c.ethClient, k, s)
	if err != nil {
		return nil, fmt.Errorf("failed to get transact opts: %s", err.Error())
	}

	tx, err := stream.ScrapProof(opts, profileID, chunkID)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
