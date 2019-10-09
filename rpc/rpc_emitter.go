package rpc

import (
	"context"
	"math/big"

	protoempty "github.com/gogo/protobuf/types"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	v1 "github.com/videocoin/cloud-api/emitter/v1"
	streamsv1 "github.com/videocoin/cloud-api/streams/v1"
)

func (s *RpcServer) InitStream(ctx context.Context, req *v1.InitStreamRequest) (*protoempty.Empty, error) {
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

		s.logger.Infof("before request stream: %s %d %v", req.UserId, streamID.Uint64(), req.ProfilesIds)

		tx, err := s.contract.RequestStream(context.Background(), req.UserId, streamID, req.ProfilesIds)
		if err != nil {
			s.logger.WithError(err).Error("failed to request stream")
			return
		}

		err = s.contract.WaitMinedAndCheck(tx)
		if err != nil {
			s.logger.WithError(err).Error("failed to wait mined")
			return
		}

		tx, err = s.contract.ApproveStream(ctx, streamID)
		if err != nil {
			s.logger.WithError(err).Error("failed to approve stream")
			return
		}

		err = s.contract.WaitMinedAndCheck(tx)
		if err != nil {
			s.logger.WithError(err).Error("failed to wait mined")
			return
		}

		tx, err = s.contract.CreateStream(ctx, req.UserId, streamID)
		if err != nil {
			s.logger.WithError(err).Error("failed to create stream")
			return
		}

		err = s.contract.WaitMinedAndCheck(tx)
		if err != nil {
			s.logger.WithError(err).Error("failed to wait mined")
			return
		}

		streamAddress, err := s.contract.GetStreamAddress(ctx, streamID)
		if err != nil {
			s.logger.WithError(err).Error("failed to get requests")
			return
		}

		_, err = s.streams.UpdateStatus(ctx, &streamsv1.UpdateStreamRequest{
			Id:                    req.StreamId,
			StreamContractAddress: streamAddress,
			StreamContractId:      streamID.Uint64(),
		})
		if err != nil {
			s.logger.WithError(err).Error("failed to update stream")
			return
		}

		// will be updated on defer
		streamStatus = streamsv1.StreamStatusPrepared

		_, err = s.contract.AllowRefund(ctx, streamID)
		if err != nil {
			s.logger.WithError(err).Error("failed to allow refund")
			return
		}
	}(actx, req)

	return &protoempty.Empty{}, nil
}

func (s *RpcServer) EndStream(ctx context.Context, req *v1.EndStreamRequest) (*protoempty.Empty, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "EndStream")
	defer span.Finish()

	span.SetTag("user_id", req.UserId)
	span.SetTag("stream_id", req.StreamContractId)

	actx := opentracing.ContextWithSpan(context.Background(), span)

	go func(ctx context.Context, req *v1.EndStreamRequest) {
		_, ctx = opentracing.StartSpanFromContext(ctx, "EndStreamAsync")

		streamID := new(big.Int).SetUint64(req.StreamContractId)

		tx, err := s.contract.EndStream(ctx, req.UserId, streamID)
		if err != nil {
			s.logger.WithError(err).Error("failed to request stream")
			return
		}

		err = s.contract.WaitMinedAndCheck(tx)
		if err != nil {
			s.logger.WithError(err).Error("failed to wait mined")
			return
		}

		_, err = s.contract.EscrowRefund(ctx, req.StreamContractAddress)
		if err != nil {
			s.logger.WithError(err).Error("failed to request stream")
		}
	}(actx, req)

	return &protoempty.Empty{}, nil
}

func (s *RpcServer) AddInputChunkId(ctx context.Context, req *v1.AddInputChunkIdRequest) (*protoempty.Empty, error) {
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
