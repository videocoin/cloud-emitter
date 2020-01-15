package contract

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/opentracing/opentracing-go"
	"github.com/videocoin/cloud-pkg/stream"
)

func (c *ContractClient) RequestStream(
	ctx context.Context, userId string, streamId *big.Int, profileNames []string) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RequestStream")
	defer span.Finish()

	span.SetTag("user_id", userId)
	span.SetTag("stream_id", streamId.Uint64())

	transactOpts, err := c.getClientTransactOpts(ctx, userId)
	if err != nil {
		c.logger.WithError(err).Error("failed to get client transact opts")
		return nil, err
	}

	tx, err := c.streamManager.RequestStream(
		transactOpts,
		streamId,
		profileNames,
	)

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *ContractClient) GetStreamAddress(ctx context.Context, streamId *big.Int) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetStreamAddress")
	defer span.Finish()

	span.SetTag("stream_id", streamId.Uint64())

	s, err := c.streamManager.Requests(new(bind.CallOpts), streamId)
	if err != nil {
		return "", err
	}

	return s.Stream.Hex(), nil
}

func (c *ContractClient) ApproveStream(ctx context.Context, streamId *big.Int) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ApproveStream")
	defer span.Finish()

	span.SetTag("stream_id", streamId.Uint64())

	transactOpts, err := c.getManagerTransactOpts(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := c.streamManager.ApproveStreamCreation(
		transactOpts,
		streamId,
	)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *ContractClient) CreateStream(ctx context.Context, userId string, streamId, deposit *big.Int) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateStream")
	defer span.Finish()

	span.SetTag("user_id", userId)
	span.SetTag("stream_id", streamId.Uint64())

	transactOpts, err := c.getClientTransactOpts(ctx, userId)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}

	transactOpts.Value = deposit

	tx, err := c.streamManager.CreateStream(
		transactOpts,
		streamId,
	)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}

	return tx, nil
}

func (c *ContractClient) EndStream(ctx context.Context, userId string, streamId *big.Int) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EndStream")
	defer span.Finish()

	span.SetTag("user_id", userId)
	span.SetTag("stream_id", streamId.Uint64())

	transactOpts, err := c.getClientTransactOpts(ctx, userId)
	if err != nil {
		return nil, err
	}

	tx, err := c.streamManager.EndStream(
		transactOpts,
		streamId,
	)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *ContractClient) AllowRefund(ctx context.Context, streamId *big.Int) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AllowRefund")
	defer span.Finish()

	transactOpts, err := c.getManagerTransactOpts(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := c.streamManager.AllowRefund(transactOpts, streamId)
	if err != nil {
		return nil, nil
	}

	return tx, nil
}

func (c *ContractClient) EscrowRefund(ctx context.Context, streamContractAddress string) (*types.Transaction, error) {
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

func (c *ContractClient) AddInputChunkID(
	ctx context.Context, streamId, chunkId *big.Int, rewards []*big.Int) (*types.Transaction, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "addInputChunkId")
	defer span.Finish()

	transactOpts, err := c.getManagerTransactOpts(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := c.streamManager.AddInputChunkId(transactOpts, streamId, chunkId, rewards)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *ContractClient) Deposit(ctx context.Context, userID string, to, value *big.Int) (*types.Transaction, error) {
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
