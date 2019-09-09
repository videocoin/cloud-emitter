package rpc

import (
	"context"
	"math/big"

	protoempty "github.com/gogo/protobuf/types"
	"github.com/opentracing/opentracing-go"
	v1 "github.com/videocoin/cloud-api/emitter/v1"
	streamsv1 "github.com/videocoin/cloud-api/streams/v1"
)

func (s *RpcServer) InitStream(ctx context.Context, req *v1.InitStreamRequest) (*protoempty.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InitStream")
	defer span.Finish()

	span.SetTag("user_id", req.UserId)
	span.SetTag("stream_id", req.StreamContractId)

	streamId := new(big.Int).SetUint64(req.StreamContractId)

	go func() {
		tx, err := s.contract.RequestStream(ctx, req.UserId, streamId, req.ProfileNames)
		if err != nil {
			s.logger.WithError(err).Error("failed to request stream")
			return
		}

		_, err = s.contract.WaitMined(ctx, tx)
		if err != nil {
			s.logger.WithError(err).Error("failed to wait mined")
			return
		}

		tx, err = s.contract.ApproveStream(ctx, streamId)
		if err != nil {
			s.logger.WithError(err).Error("failed to approve stream")
			return
		}

		_, err = s.contract.WaitMined(ctx, tx)
		if err != nil {
			s.logger.WithError(err).Error("failed to wait mined")
			return
		}

		tx, err = s.contract.CreateStream(ctx, req.UserId, streamId)
		if err != nil {
			s.logger.WithError(err).Error("failed to create stream")
			return
		}

		receipt, err := s.contract.WaitMined(ctx, tx)
		if err != nil {
			s.logger.WithError(err).Error("failed to wait mined")
			return
		}

		_, err = s.streams.Update(ctx, &streamsv1.UpdateStreamRequest{
			StreamContractAddress: receipt.ContractAddress.String(),
			StreamContractId:      streamId.Uint64(),
		})

		if err != nil {
			s.logger.WithError(err).Error("failed to update stream")
			return
		}

		_, err = s.contract.AllowRefund(ctx, streamId)
		if err != nil {
			s.logger.WithError(err).Error("failed to allow refund")
			return
		}
	}()

	return &protoempty.Empty{}, nil
}

func (s *RpcServer) EndStream(ctx context.Context, req *v1.EndStreamRequest) (*protoempty.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EndStream")
	defer span.Finish()

	span.SetTag("user_id", req.UserId)
	span.SetTag("stream_id", req.StreamContractId)

	streamId := new(big.Int).SetUint64(req.StreamContractId)

	go func() {
		tx, err := s.contract.EndStream(ctx, req.UserId, streamId)
		if err != nil {
			s.logger.WithError(err).Error("failed to request stream")
		}

		_, err = s.contract.WaitMined(ctx, tx)
		if err != nil {
			s.logger.WithError(err).Error("failed to wait mined")
			// return
		}

		_, err = s.contract.EscrowRefund(ctx, req.StreamContractAddress)
		if err != nil {
			s.logger.WithError(err).Error("failed to request stream")
		}
	}()

	return &protoempty.Empty{}, nil
}

func (s *RpcServer) AddInputChunkId(ctx context.Context, req *v1.AddInputChunkIdRequest) (*v1.Tx, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AddInputChunkId")
	defer span.Finish()

	span.SetTag("stream_id", req.StreamContractId)

	streamId := new(big.Int).SetUint64(req.StreamContractId)
	chunkId := new(big.Int).SetUint64(req.ChunkId)

	// 0.01 const reward
	reward, _ := big.NewFloat(10000000000000000 * req.ChunkDuration).Int64()
	rewards := []*big.Int{big.NewInt(reward)}

	tx, err := s.contract.AddInputChunkID(ctx, streamId, chunkId, rewards)
	if err != nil {
		s.logger.WithError(err).Error("failed to add input chunk id")
		return nil, err
	}

	return &v1.Tx{
		Hash: tx.Hash().Bytes(),
	}, nil
}
