package rpc

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	protoempty "github.com/gogo/protobuf/types"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	v1 "github.com/videocoin/cloud-api/emitter/v1"
	"github.com/videocoin/cloud-api/rpc"
	streamsv1 "github.com/videocoin/cloud-api/streams/v1"
)

func (s *Server) InitStream(ctx context.Context, req *v1.InitStreamRequest) (*protoempty.Empty, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "InitStream")
	defer span.Finish()

	span.SetTag("user_id", req.UserId)
	span.SetTag("stream_id", req.StreamContractId)

	actx := opentracing.ContextWithSpan(context.Background(), span)

	go func(ctx context.Context, req *v1.InitStreamRequest) {
		_, ctx = opentracing.StartSpanFromContext(ctx, "AsyncInitStream")

		streamID := new(big.Int).SetUint64(req.StreamContractId)
		streamStatus := streamsv1.StreamStatusFailed

		defer func() {
			_, err := s.streams.UpdateStatus(ctx, &streamsv1.UpdateStreamRequest{
				Id:     req.StreamId,
				Status: streamStatus,
			})

			if err != nil {
				s.logger.WithError(err).Errorf("failed to update stream status to %d on defer", streamStatus)
				return
			}
		}()

		logger := s.logger.WithFields(logrus.Fields{
			"user_id":   req.UserId,
			"stream_id": streamID.Uint64(),
		})

		logger.Info("request stream")

		tx, err := s.contract.RequestStream(context.Background(), req.UserId, streamID, req.ProfilesIds)
		if err != nil {
			logger.WithError(err).Error("failed to request stream")
			return
		}

		if tx != nil {
			logger.Infof("request stream tx %s", tx.Hash().String())
		}

		err = s.contract.WaitMinedAndCheck(tx)
		if err != nil {
			logger.WithError(err).Error("failed to wait mined")
			return
		}

		logger.Info("approve stream")

		tx, err = s.contract.ApproveStream(ctx, streamID)
		if err != nil {
			logger.WithError(err).Error("failed to approve stream")
			return
		}

		if tx != nil {
			logger.Infof("approve stream tx %s", tx.Hash().String())
		}

		err = s.contract.WaitMinedAndCheck(tx)
		if err != nil {
			logger.WithError(err).Error("failed to wait mined")
			return
		}

		logger.Info("create stream")

		tx, err = s.contract.CreateStream(ctx, req.UserId, streamID)
		if err != nil {
			logger.WithError(err).Error("failed to create stream")
			return
		}

		if tx != nil {
			logger.Infof("create stream tx %s", tx.Hash().String())
		}

		err = s.contract.WaitMinedAndCheck(tx)
		if err != nil {
			logger.WithError(err).Error("failed to wait mined")
			return
		}

		logger.Info("get stream address")

		streamAddress, err := s.contract.GetStreamAddress(ctx, streamID)
		if err != nil {
			logger.WithError(err).Error("failed to get requests")
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
	}(actx, req)

	return &protoempty.Empty{}, nil
}

func (s *Server) EndStream(ctx context.Context, req *v1.EndStreamRequest) (*protoempty.Empty, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "EndStream")
	defer span.Finish()

	span.SetTag("user_id", req.UserId)
	span.SetTag("stream_id", req.StreamContractId)

	actx := opentracing.ContextWithSpan(context.Background(), span)

	go func(ctx context.Context, req *v1.EndStreamRequest) {
		_, ctx = opentracing.StartSpanFromContext(ctx, "EndStreamAsync")

		streamID := new(big.Int).SetUint64(req.StreamContractId)

		s.logger.Infof("end stream %s of user %s", streamID.String(), req.UserId)

		tx, err := s.contract.EndStream(ctx, req.UserId, streamID)
		if err != nil {
			s.logger.WithError(err).Error("failed to end stream")
			return
		}

		err = s.contract.WaitMinedAndCheck(tx)
		if err != nil {
			s.logger.WithError(err).Error("failed to wait mined")
			return
		}

		s.logger.Info("end stream completed")

		_, err = s.contract.EscrowRefund(ctx, req.StreamContractAddress)
		if err != nil {
			s.logger.WithError(err).Error("failed to escrow refund")
		}

		s.logger.Info("escrow refund completed")
	}(actx, req)

	return &protoempty.Empty{}, nil
}

func (s *Server) AddInputChunkId(ctx context.Context, req *v1.AddInputChunkIdRequest) (*protoempty.Empty, error) {  // nolint
	span, ctx := opentracing.StartSpanFromContext(ctx, "AddInputChunkId")
	defer span.Finish()

	span.SetTag("stream_id", req.StreamContractId)

	streamID := new(big.Int).SetUint64(req.StreamContractId)
	chunkID := new(big.Int).SetUint64(req.ChunkId)

	// 0.01 const reward
	reward, _ := big.NewFloat(10000000000000000 * req.ChunkDuration).Int64()
	rewards := []*big.Int{big.NewInt(reward)}

	s.logger.WithFields(logrus.Fields{
		"stream_id": streamID,
		"chunk_id":  chunkID,
	}).Debugf("calling AddInputChunkID")

	go func() {
		tx, err := s.contract.AddInputChunkID(ctx, streamID, chunkID, rewards)
		if err != nil {
			s.logger.WithError(err).Error("failed to add input chunk id")
			return
		}

		err = s.contract.WaitMinedAndCheck(tx)
		if err != nil {
			s.logger.WithError(err).Error("failed to wait mined")
			return
		}
	}()

	return &protoempty.Empty{}, nil
}

func (s *Server) GetBalance(ctx context.Context, req *v1.BalanceRequest) (*v1.BalanceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetBalance")
	defer span.Finish()

	addr := new(big.Int).SetBytes(req.Address)
	span.SetTag("address", addr.String())

	s.logger.WithFields(logrus.Fields{
		"address": addr.String(),
	}).Debugf("calling BalanceAt")

	value, err := s.contract.EthClient().BalanceAt(ctx, common.BytesToAddress(req.Address), nil)
	if err != nil {
		s.logger.WithError(err).Error("failed to get balance")
		return nil, rpc.ErrRpcInternal
	}

	return &v1.BalanceResponse{
		Address: req.Address,
		Value:   value.Bytes(),
	}, nil
}

func (s *Server) Deposit(ctx context.Context, req *v1.DepositRequest) (*v1.DepositResponse, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Deposit")
	defer span.Finish()

	to := new(big.Int).SetBytes(req.To)
	value := new(big.Int).SetBytes(req.Value)

	toStr := common.BytesToAddress(req.To).String()

	span.SetTag("user_id", req.UserId)
	span.SetTag("to", toStr)
	span.SetTag("value", value.String())

	logger := s.logger.WithFields(logrus.Fields{
		"user_id": req.UserId,
		"to":      toStr,
		"value":   value.String(),
	})

	logger.Info("deposit")

	go func(userID string, streamID string, to *big.Int, logger *logrus.Entry) {
		emptyCtx := context.Background()
		tx, err := s.contract.Deposit(emptyCtx, userID, to, big.NewInt(1000000000000000000))
		if err != nil {
			logger.WithError(err).Error("failed to deposit")
			err = s.markStreamAsFailed(streamID)
			if err != nil {
				logger.WithError(err).Error("failed to mark stream as failed")
			}
			return
		}

		logger.Infof("deposit tx %s", tx.Hash().String())

		err = s.contract.WaitMinedAndCheck(tx)
		if err != nil {
			logger.WithError(err).Error("failed to wait deposit tx")
			err = s.markStreamAsFailed(streamID)
			if err != nil {
				logger.WithError(err).Error("failed to mark stream as failed")
			}
			return
		}
	}(req.UserId, req.StreamId, to, logger)

	return &v1.DepositResponse{}, nil
}

func (s *Server) markStreamAsFailed(streamID string) error {
	_, err := s.streams.UpdateStatus(context.Background(), &streamsv1.UpdateStreamRequest{
		Id:     streamID,
		Status: streamsv1.StreamStatusFailed,
	})

	return err
}
