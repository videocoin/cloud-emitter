package rpc

import (
	"context"
	"math/big"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	v1 "github.com/videocoin/cloud-api/emitter/v1"
	streamsv1 "github.com/videocoin/cloud-api/streams/v1"
)

func (s *Server) initStream(ctx context.Context, req *v1.InitStreamRequest) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "core.initStream")
	defer span.Finish()

	logger := s.logger.WithFields(logrus.Fields{
		"user_id":            req.UserId,
		"stream_id":          req.StreamId,
		"stream_contract_id": req.StreamContractId,
	})

	streamID := new(big.Int).SetUint64(req.StreamContractId)
	streamStatus := streamsv1.StreamStatusFailed

	defer func() {
		_, err := s.streams.UpdateStatus(ctx, &streamsv1.UpdateStreamRequest{
			Id:     req.StreamId,
			Status: streamStatus,
		})

		if err != nil {
			logger.
				WithError(err).
				WithField("new_status", streamStatus).
				Errorf("failed to update stream status on defer")
			return
		}
	}()

	logger.Info("request stream")

	tx, err := s.contract.RequestStream(context.Background(), req.UserId, streamID, req.ProfilesIds)
	if err != nil {
		logger.WithError(err).Error("failed to request stream")
		return
	}

	logger.WithField("tx", tx.Hash().String()).Info("request stream tx")

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		logger.
			WithError(err).
			WithField("call", "RequestStream").
			WithField("tx", tx.Hash().String()).
			Error("failed to wait mined")
		return
	}

	logger.Info("approve stream")

	tx, err = s.contract.ApproveStream(ctx, streamID)
	if err != nil {
		logger.WithError(err).Error("failed to approve stream")
		return
	}

	logger.WithField("tx", tx.Hash().String()).Info("approve stream tx")

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		logger.WithError(err).
			WithField("call", "ApproveStream").
			WithField("tx", tx.Hash().String()).
			Error("failed to wait mined")
		return
	}

	logger.Info("create stream")

	tx, err = s.contract.CreateStream(ctx, req.UserId, streamID)
	if err != nil {
		logger.WithError(err).Error("failed to create stream")
		return
	}

	logger.WithField("tx", tx.Hash().String()).Info("create stream tx")

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		logger.
			WithError(err).
			WithField("call", "CreateStream").
			WithField("tx", tx.Hash().String()).
			Error("failed to wait mined")
		return
	}

	logger.Info("getting stream address")

	streamAddress, err := s.contract.GetStreamAddress(ctx, streamID)
	if err != nil {
		logger.WithError(err).Error("failed to get stream address")
		return
	}

	_, err = s.streams.UpdateStatus(ctx, &streamsv1.UpdateStreamRequest{
		Id:                    req.StreamId,
		StreamContractAddress: streamAddress,
		StreamContractId:      streamID.Uint64(),
	})
	if err != nil {
		logger.WithError(err).Error("failed to update stream")
		return
	}

	// will be updated on defer
	streamStatus = streamsv1.StreamStatusPrepared

	logger.Info("allow refund")

	_, err = s.contract.AllowRefund(ctx, streamID)
	if err != nil {
		logger.WithError(err).Error("failed to allow refund")
		return
	}
}

func (s *Server) endStream(ctx context.Context, req *v1.EndStreamRequest) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "core.endStream")
	defer span.Finish()

	logger := s.logger.WithFields(logrus.Fields{
		"user_id":                 req.UserId,
		"stream_id":               req.StreamId,
		"stream_contract_id":      req.StreamContractId,
		"stream_contract_address": req.StreamContractAddress,
	})

	streamID := new(big.Int).SetUint64(req.StreamContractId)

	logger.Info("end stream")

	tx, err := s.contract.EndStream(ctx, req.UserId, streamID)
	if err != nil {
		logger.WithError(err).Error("failed to end stream")
		return
	}

	logger.WithField("tx", tx.Hash().String()).Info("end stream tx")

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		logger.
			WithError(err).
			WithField("call", "EndStream").
			WithField("tx", tx.Hash().String()).
			Error("failed to wait mined")
		return
	}

	logger.Info("escrow refund")

	_, err = s.contract.EscrowRefund(ctx, req.StreamContractAddress)
	if err != nil {
		logger.WithError(err).Error("failed to escrow refund")
	}
}

func (s *Server) addInputChunk(ctx context.Context, req *v1.AddInputChunkRequest) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "core.addInputChunk")
	defer span.Finish()

	logger := s.logger.WithFields(logrus.Fields{
		"stream_contract_id": req.StreamContractId,
		"chunk_id":           req.ChunkId,
		"reward":             req.Reward,
	})

	streamID := new(big.Int).SetUint64(req.StreamContractId)
	chunkID := new(big.Int).SetUint64(req.ChunkId)

	reward, _ := new(big.Float).Mul(big.NewFloat(req.Reward), big.NewFloat(1000000000000000000)).Int64()
	rewards := []*big.Int{big.NewInt(reward)}

	logger.Info("add input chunk")

	tx, err := s.contract.AddInputChunkID(ctx, streamID, chunkID, rewards)
	if err != nil {
		logger.WithError(err).Error("failed to add input chunk")
		return
	}

	logger.WithField("tx", tx.Hash().String()).Info("add input chunk tx")

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		logger.
			WithError(err).
			WithField("call", "addInputChunkId").
			WithField("tx", tx.Hash().String()).
			Error("failed to wait mined")
		return
	}
}

func (s *Server) deposit(ctx context.Context, userID, streamID string, to *big.Int) {
	logger := s.logger.WithFields(logrus.Fields{
		"user_id":   userID,
		"stream_id": streamID,
		"to":        to.String(),
	})

	logger.Info("deposit")

	tx, err := s.contract.Deposit(ctx, userID, to, big.NewInt(1000000000000000000))
	if err != nil {
		logger.WithError(err).Error("failed to deposit")

		err = s.markStreamAsFailed(streamID)
		if err != nil {
			logger.WithError(err).Error("failed to mark stream as failed")
		}

		return
	}

	logger.WithField("tx", tx.Hash().String()).Info("deposit tx")

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		logger.
			WithError(err).
			WithField("tx", tx.Hash().String()).
			Error("failed to wait mined")

		err = s.markStreamAsFailed(streamID)
		if err != nil {
			logger.WithError(err).Error("failed to mark stream as failed")
		}

		return
	}
}
