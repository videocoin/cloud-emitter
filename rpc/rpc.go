package rpc

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/ethereum/go-ethereum/common"
	ethcore "github.com/ethereum/go-ethereum/core"
	protoempty "github.com/gogo/protobuf/types"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	accountsv1 "github.com/videocoin/cloud-api/accounts/v1"
	v1 "github.com/videocoin/cloud-api/emitter/v1"
	"github.com/videocoin/cloud-api/rpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) InitStream(ctx context.Context, req *v1.InitStreamRequest) (*v1.InitStreamResponse, error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("user_id", req.UserId)
	span.SetTag("stream_id", req.StreamId)
	span.SetTag("stream_contract_id", req.StreamContractId)

	resp, err := s.initStream(ctx, req)
	if err != nil {
		return resp, rpc.NewRpcInternalError(err)
	}

	return resp, nil
}

func (s *Server) EndStream(ctx context.Context, req *v1.EndStreamRequest) (*v1.EndStreamResponse, error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("user_id", req.UserId)
	span.SetTag("stream_id", req.StreamId)
	span.SetTag("stream_contract_id", req.StreamContractId)
	span.SetTag("stream_contract_address", req.StreamContractAddress)

	resp, err := s.endStream(ctx, req)
	if err != nil {
		return resp, rpc.NewRpcInternalError(err)
	}

	return resp, nil
}

func (s *Server) AddInputChunk(ctx context.Context, req *v1.AddInputChunkRequest) (*v1.AddInputChunkResponse, error) {
	span := opentracing.SpanFromContext(ctx)

	if req.StreamContractId == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid stream contract address")
	}

	if req.ChunkId == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid chunk id")
	}

	if req.Reward == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid reward")
	}

	span.SetTag("stream_contract_id", req.StreamContractId)
	span.SetTag("chunk_id", req.ChunkId)
	span.SetTag("reward", req.Reward)

	resp, err := s.addInputChunk(ctx, req)
	if err != nil {
		return resp, rpc.NewRpcInternalError(err)
	}

	return resp, nil
}

func (s *Server) Deposit(ctx context.Context, req *v1.DepositRequest) (*v1.DepositResponse, error) {
	span := opentracing.SpanFromContext(ctx)

	to := new(big.Int).SetBytes(req.To)
	toStr := common.BytesToAddress(req.To).String()

	span.SetTag("user_id", req.UserId)
	span.SetTag("to", toStr)

	resp, err := s.deposit(ctx, req.UserId, req.StreamId, to)
	if err != nil {
		return resp, rpc.NewRpcInternalError(err)
	}

	return resp, nil
}

func (s *Server) GetBalance(ctx context.Context, req *v1.BalanceRequest) (*v1.BalanceResponse, error) {
	span := opentracing.SpanFromContext(ctx)
	addr := new(big.Int).SetBytes(req.Address)
	span.SetTag("address", addr.String())

	logger := s.logger.WithField("address", addr.String())
	logger.Info("get balance")

	value, err := s.contract.EthClient().BalanceAt(ctx, common.BytesToAddress(req.Address), nil)
	if err != nil {
		logger.WithError(err).Error("failed to get balance")
		return nil, rpc.ErrRpcInternal
	}

	return &v1.BalanceResponse{
		Address: req.Address,
		Value:   value.Bytes(),
	}, nil
}

func (s *Server) ValidateProof(ctx context.Context, req *v1.ValidateProofRequest) (*v1.ValidateProofResponse, error) {
	span := opentracing.SpanFromContext(ctx)

	profileID := new(big.Int).SetBytes(req.ProfileId)
	chunkID := new(big.Int).SetBytes(req.ChunkId)

	span.SetTag("stream_contract_address", req.StreamContractAddress)
	span.SetTag("profile_id", profileID.String())
	span.SetTag("chunk_id", chunkID.String())

	logger := s.logger.WithFields(logrus.Fields{
		"stream_contract_address": req.StreamContractAddress,
		"profile_id":              profileID.String(),
		"chunk_id":                chunkID.String(),
	})

	resp := &v1.ValidateProofResponse{}

	logger.Info("validate proof")

	tx, err := s.contract.ValidateProof(ctx, req.StreamContractAddress, profileID, chunkID)
	if err != nil {
		if err.Error() == ethcore.ErrNonceTooLow.Error() ||
			err.Error() == ethcore.ErrReplaceUnderpriced.Error() {
			logger.Info("validate proof (retry)")
			time.Sleep(time.Millisecond * 500)
			var retryErr error
			tx, retryErr = s.contract.ValidateProof(ctx, req.StreamContractAddress, profileID, chunkID)
			if retryErr != nil {
				logger.Errorf("failed to retry contract.ValidateProof: %s", retryErr)
				return resp, fmt.Errorf("failed to retry ValidateProof: %s", retryErr)
			}
		} else {
			logger.Errorf("failed to contract.ValidateProof: %s", err)
			fmtErr := fmt.Errorf("failed to ValidateProof: %s", err)
			return resp, rpc.NewRpcInternalError(fmtErr)
		}
	}

	resp.Tx = tx.Hash().String()

	logger = logger.WithField("tx", resp.Tx)

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		resp.Status = v1.ReceiptStatusFailed
		logger.Errorf("failed to wait contract.ValidateProof: %s", err)
		fmtErr := fmt.Errorf("failed to wait ValidateProof: %s", err)
		return resp, rpc.NewRpcInternalError(fmtErr)
	}

	resp.Status = v1.ReceiptStatusSuccess

	return resp, nil
}

func (s *Server) ScrapProof(ctx context.Context, req *v1.ScrapProofRequest) (*v1.ScrapProofResponse, error) {
	span := opentracing.SpanFromContext(ctx)

	profileID := new(big.Int).SetBytes(req.ProfileId)
	chunkID := new(big.Int).SetBytes(req.ChunkId)

	span.SetTag("stream_contract_address", req.StreamContractAddress)
	span.SetTag("profile_id", profileID.String())
	span.SetTag("chunk_id", chunkID.String())

	logger := s.logger.WithFields(logrus.Fields{
		"stream_contract_address": req.StreamContractAddress,
		"profile_id":              profileID.String(),
		"chunk_id":                chunkID.String(),
	})

	resp := &v1.ScrapProofResponse{}

	logger.Info("scrap proof")

	tx, err := s.contract.ScrapProof(ctx, req.StreamContractAddress, profileID, chunkID)
	if err != nil {
		fmtErr := fmt.Errorf("failed to ScrapProof: %s", err)
		return resp, rpc.NewRpcInternalError(fmtErr)
	}

	resp.Tx = tx.Hash().String()

	logger = logger.WithField("tx", resp.Tx)

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		resp.Status = v1.ReceiptStatusFailed
		logger.Errorf("failed to wait contract.ScrapProof: %s", err)
		fmtErr := fmt.Errorf("failed to wait ScrapProof: %s", err)
		return resp, rpc.NewRpcInternalError(fmtErr)
	}

	resp.Status = v1.ReceiptStatusSuccess

	return resp, nil
}

func (s *Server) ListWorkers(ctx context.Context, req *protoempty.Empty) (*v1.ListWorkersResponse, error) {
	workers, err := s.staking.GetAllTranscoders(ctx)
	if err != nil {
		return nil, rpc.NewRpcInternalError(err)
	}

	resp := &v1.ListWorkersResponse{
		Items: []*v1.WorkerResponse{},
	}

	for _, worker := range workers {
		resp.Items = append(resp.Items, &v1.WorkerResponse{
			Address:        worker.Address.Hex(),
			State:          v1.WorkerState(worker.State),
			TotalStake:     worker.TotalStake.String(),
			SelfStake:      worker.SelfStake.String(),
			DelegatedStake: worker.DelegatedStake.String(),
			RegisteredAt:   pointer.ToTime(time.Unix(int64(worker.Timestamp), 0)),
		})
	}

	return resp, nil
}

func (s *Server) AddFunds(ctx context.Context, req *v1.AddFundsRequest) (*v1.AddFundsResponse, error) {
	resp := new(v1.AddFundsResponse)
	accountReq := &accountsv1.AccountRequest{OwnerId: req.UserID}
	accountResp, err := s.accounts.GetByOwner(ctx, accountReq)
	if err != nil {
		return resp, err
	}

	err = s.faucet.Do(accountResp.Address, uint64(req.AmountUsd)+1)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
