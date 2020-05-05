package rpc

import (
	"context"
	"fmt"
	"math/big"
	"time"

	ethcore "github.com/ethereum/go-ethereum/core"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	v1 "github.com/videocoin/cloud-api/emitter/v1"
)

func (s *Server) initStream(ctx context.Context, req *v1.InitStreamRequest) (*v1.InitStreamResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "core.initStream")
	defer span.Finish()

	logger := s.logger.WithFields(logrus.Fields{
		"user_id":            req.UserId,
		"stream_id":          req.StreamId,
		"stream_contract_id": req.StreamContractId,
	})

	streamID := new(big.Int).SetUint64(req.StreamContractId)

	resp := &v1.InitStreamResponse{}

	logger.Info("request stream")

	tx, err := s.contract.RequestStream(context.Background(), req.UserId, streamID, req.ProfilesIds)
	if err != nil {
		return resp, fmt.Errorf("failed to RequestStream: %s", err)
	}

	resp.RequestStreamTx = tx.Hash().String()

	logger.WithField("tx", resp.RequestStreamTx).Info("request stream tx")

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		resp.RequestStreamTxStatus = v1.ReceiptStatusFailed
		return resp, fmt.Errorf("failed to wait RequestStream: %s", err)
	}

	resp.RequestStreamTxStatus = v1.ReceiptStatusSuccess

	logger.Info("approve stream")

	tx, err = s.contract.ApproveStream(ctx, streamID)
	if err != nil {
		return resp, fmt.Errorf("failed to ApproveStream: %s", err)
	}

	resp.ApproveStreamTx = tx.Hash().String()

	logger.WithField("tx", resp.ApproveStreamTx).Info("approve stream tx")

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		resp.ApproveStreamTxStatus = v1.ReceiptStatusFailed
		return resp, fmt.Errorf("failed to wait ApproveStream: %s", err)
	}

	resp.ApproveStreamTxStatus = v1.ReceiptStatusSuccess

	logger.Info("create stream")

	tx, err = s.contract.CreateStream(ctx, req.UserId, streamID)
	if err != nil {
		return resp, fmt.Errorf("failed to CreateStream: %s", err)
	}

	resp.CreateStreamTx = tx.Hash().String()

	logger.WithField("tx", resp.CreateStreamTx).Info("create stream tx")

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		resp.CreateStreamTxStatus = v1.ReceiptStatusFailed
		return resp, fmt.Errorf("failed to wait CreateStreamTxStatus: %s", err)
	}

	logger.Info("getting stream address")

	streamAddress, err := s.contract.GetStreamAddress(ctx, streamID)
	if err != nil {
		return resp, fmt.Errorf("failed to GetStreamAddress: %s", err)
	}

	resp.StreamContractAddress = streamAddress

	logger.Info("allow refund")

	tx, err = s.contract.AllowRefund(ctx, streamID)
	if err != nil {
		return resp, fmt.Errorf("failed to AllowRefund: %s", err)
	}

	resp.AllowRefundTx = tx.Hash().String()

	logger.WithField("tx", resp.AllowRefundTx).Info("allow refund tx")

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		resp.AllowRefundTxStatus = v1.ReceiptStatusFailed
		return resp, fmt.Errorf("failed to wait AllowRefundTxStatus: %s", err)
	}

	resp.AllowRefundTxStatus = v1.ReceiptStatusSuccess

	return resp, nil
}

func (s *Server) endStream(ctx context.Context, req *v1.EndStreamRequest) (*v1.EndStreamResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "core.endStream")
	defer span.Finish()

	logger := s.logger.WithFields(logrus.Fields{
		"user_id":                 req.UserId,
		"stream_id":               req.StreamId,
		"stream_contract_id":      req.StreamContractId,
		"stream_contract_address": req.StreamContractAddress,
	})

	streamID := new(big.Int).SetUint64(req.StreamContractId)

	resp := &v1.EndStreamResponse{}

	logger.Info("end stream")

	tx, err := s.contract.EndStream(ctx, req.UserId, streamID)
	if err != nil {
		return resp, fmt.Errorf("failed to end stream: %s", err)
	}

	resp.EndStreamTx = tx.Hash().String()

	logger.WithField("tx", resp.EndStreamTx).Info("end stream tx")

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		resp.EndStreamTxStatus = v1.ReceiptStatusFailed
		return resp, fmt.Errorf("failed to wait EndStream: %s", err)
	}

	resp.EndStreamTxStatus = v1.ReceiptStatusSuccess

	logger.Info("escrow refund")

	tx, err = s.contract.EscrowRefund(ctx, req.StreamContractAddress)
	if err != nil {
		return resp, fmt.Errorf("failed to escrow refund: %s", err)
	}

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		resp.EscrowRefundTxStatus = v1.ReceiptStatusFailed
		return resp, fmt.Errorf("failed to wait EscrowRefundTxStatus: %s", err)
	}

	resp.EscrowRefundTxStatus = v1.ReceiptStatusSuccess

	return resp, nil
}

func (s *Server) addInputChunk(ctx context.Context, req *v1.AddInputChunkRequest) (*v1.AddInputChunkResponse, error) {
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

	resp := &v1.AddInputChunkResponse{}

	logger.Info("add input chunk")

	tx, err := s.contract.AddInputChunkID(ctx, streamID, chunkID, rewards)
	if err != nil {
		if err.Error() == ethcore.ErrNonceTooLow.Error() ||
			err.Error() == ethcore.ErrReplaceUnderpriced.Error() {
			logger.Info("add input chunk (retry)")
			time.Sleep(time.Millisecond * 500)
			var retryErr error
			tx, retryErr = s.contract.AddInputChunkID(ctx, streamID, chunkID, rewards)
			if retryErr != nil {
				logger.Errorf("failed to retry contract.AddInputChunkID: %s", retryErr)
				return resp, fmt.Errorf("failed to retry contract.AddInputChunkID: %s", retryErr)
			}
		} else {
			logger.Errorf("failed to contract.AddInputChunkID: %s", err)
			return resp, fmt.Errorf("failed to add input chunk: %s", err)
		}
	}

	resp.Tx = tx.Hash().String()

	logger = logger.WithField("tx", resp.Tx)
	logger.Info("add input chunk tx")

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		resp.Status = v1.ReceiptStatusFailed
		logger.Errorf("failed to wait contract.AddInputChunkID: %s", err)
		return resp, fmt.Errorf("failed to wait AddInputChunkID")
	}

	resp.Status = v1.ReceiptStatusSuccess

	return resp, nil
}

func (s *Server) deposit(ctx context.Context, userID, streamID string, to *big.Int) (*v1.DepositResponse, error) {
	logger := s.logger.WithFields(logrus.Fields{
		"user_id":   userID,
		"stream_id": streamID,
		"to":        to.String(),
	})

	resp := &v1.DepositResponse{}

	logger.Info("deposit")

	tx, err := s.contract.Deposit(ctx, userID, to, big.NewInt(1000000000000000000))
	if err != nil {
		return resp, fmt.Errorf("failed to Deposit: %s", err)
	}

	resp.Tx = tx.Hash().String()

	logger.WithField("tx", resp.Tx).Info("deposit tx")

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		resp.Status = v1.ReceiptStatusFailed
		return resp, fmt.Errorf("failed to wait Deposit: %s", err)
	}

	resp.Status = v1.ReceiptStatusSuccess

	return resp, nil
}
