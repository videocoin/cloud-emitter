package contract

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/opentracing/opentracing-go"
	"github.com/videocoin/cloud-pkg/stream"
)

func (c *Client) RequestStream(ctx context.Context, userID string, streamID *big.Int, profileNames []string) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RequestStream")
	defer span.Finish()

	span.SetTag("user_id", userID)
	span.SetTag("stream_id", streamID.Uint64())

	transactOpts, err := c.getClientTransactOpts(ctx, userID)
	if err != nil {
		c.logger.WithError(err).Error("failed to get client transact opts")
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

	span.SetTag("stream_id", streamID.Uint64())

	s, err := c.streamManager.Requests(new(bind.CallOpts), streamID)
	if err != nil {
		return "", err
	}

	return s.Stream.Hex(), nil
}

func (c *Client) ApproveStream(ctx context.Context, streamID *big.Int) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ApproveStream")
	defer span.Finish()

	span.SetTag("stream_id", streamID.Uint64())

	transactOpts, err := c.getManagerTransactOpts(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := c.streamManager.ApproveStreamCreation(
		transactOpts,
		streamID,
	)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) CreateStream(ctx context.Context, userID string, streamID *big.Int) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateStream")
	defer span.Finish()

	span.SetTag("user_id", userID)
	span.SetTag("stream_id", streamID.Uint64())

	transactOpts, err := c.getClientTransactOpts(ctx, userID)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}

	// todo: constant value ???
	i, e := big.NewInt(10), big.NewInt(19)
	transactOpts.Value = i.Exp(i, e, nil)

	tx, err := c.streamManager.CreateStream(
		transactOpts,
		streamID,
	)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}

	return tx, nil
}

func (c *Client) EndStream(ctx context.Context, userID string, streamID *big.Int) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EndStream")
	defer span.Finish()

	span.SetTag("user_id", userID)
	span.SetTag("stream_id", streamID.Uint64())

	transactOpts, err := c.getClientTransactOpts(ctx, userID)
	if err != nil {
		return nil, err
	}

	tx, err := c.streamManager.EndStream(
		transactOpts,
		streamID,
	)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) AllowRefund(ctx context.Context, streamID *big.Int) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AllowRefund")
	defer span.Finish()

	transactOpts, err := c.getManagerTransactOpts(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := c.streamManager.AllowRefund(transactOpts, streamID)
	if err != nil {
		return nil, nil
	}

	return tx, nil
}

func (c *Client) EscrowRefund(ctx context.Context, streamContractAddress string) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EscrowRefund")
	defer span.Finish()

	transactOpts, err := c.getManagerTransactOpts(ctx)
	if err != nil {
		return nil, err
	}

	stream, err := stream.NewStream(common.HexToAddress(streamContractAddress), c.ethClient)
	if err != nil {
		return nil, err
	}

	tx, err := stream.Refund(transactOpts)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) AddInputChunkID(ctx context.Context, streamID, chunkID *big.Int, rewards []*big.Int) (*types.Transaction, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "addInputchunkId")
	defer span.Finish()

	transactOpts, err := c.getManagerTransactOpts(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := c.streamManager.AddInputChunkId(transactOpts, streamID, chunkID, rewards)
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

	s, err := stream.NewStream(common.BigToAddress(to), c.ethClient)
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

	stream, err := stream.NewStream(common.HexToAddress(streamContractAddress), c.ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new stream: %s", err.Error())
	}

	transactOpts, err := getTransactOpts(context.Background(), c.ethClient, c.validatorKey, c.validatorSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to get transact opts: %s", err.Error())
	}

	tx, err := stream.ValidateProof(transactOpts, profileID, chunkID)
	if err != nil {
		return nil, fmt.Errorf("failed to validate proof: %s", err.Error())
	}

	return tx, nil
}

func (c *Client) ScrapProof(ctx context.Context, streamContractAddress string, profileID, chunkID *big.Int) (*types.Transaction, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ScrapProof")
	defer span.Finish()

	stream, err := stream.NewStream(common.HexToAddress(streamContractAddress), c.ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new stream: %s", err.Error())
	}

	transactOpts, err := getTransactOpts(context.Background(), c.ethClient, c.validatorKey, c.validatorSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to get transact opts: %s", err.Error())
	}

	tx, err := stream.ScrapProof(transactOpts, profileID, chunkID)
	if err != nil {
		return nil, fmt.Errorf("failed to scrap proof: %s", err.Error())
	}

	return tx, nil
}
