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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) InitStream(ctx context.Context, req *v1.InitStreamRequest) (*protoempty.Empty, error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("user_id", req.UserId)
	span.SetTag("stream_id", req.StreamId)
	span.SetTag("stream_contract_id", req.StreamContractId)

	s.initStreamAsync(opentracing.ContextWithSpan(context.Background(), span), req)

	return &protoempty.Empty{}, nil
}

func (s *Server) EndStream(ctx context.Context, req *v1.EndStreamRequest) (*protoempty.Empty, error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("user_id", req.UserId)
	span.SetTag("stream_id", req.StreamId)
	span.SetTag("stream_contract_id", req.StreamContractId)
	span.SetTag("stream_contract_address", req.StreamContractAddress)

	s.endStreamAsync(opentracing.ContextWithSpan(context.Background(), span), req)

	return &protoempty.Empty{}, nil
}

func (s *Server) AddInputChunk(ctx context.Context, req *v1.AddInputChunkRequest) (*protoempty.Empty, error) {
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

	s.addInputChunkAsync(opentracing.ContextWithSpan(context.Background(), span), req)

	return &protoempty.Empty{}, nil
}

func (s *Server) Deposit(ctx context.Context, req *v1.DepositRequest) (*v1.DepositResponse, error) {
	span := opentracing.SpanFromContext(ctx)

	to := new(big.Int).SetBytes(req.To)
	toStr := common.BytesToAddress(req.To).String()

	span.SetTag("user_id", req.UserId)
	span.SetTag("to", toStr)

	s.depositAsync(opentracing.ContextWithSpan(context.Background(), span), req.UserId, req.StreamId, to)

	return &v1.DepositResponse{}, nil
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

	logger.Info("validate proof")

	tx, err := s.contract.ValidateProof(ctx, req.StreamContractAddress, profileID, chunkID)
	if err != nil {
		logger.
			WithError(err).
			Error("failed to validate proof")
		return nil, rpc.ErrRpcInternal
	}

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		logger.
			WithError(err).
			WithField("call", "ValidateProof").
			WithField("tx", tx.Hash().String()).
			Error("failed to wait mined")
		return nil, rpc.ErrRpcInternal
	}

	return &v1.ValidateProofResponse{
		TxId: []byte(tx.Hash().String()),
	}, nil
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

	logger.Info("scrap proof")

	tx, err := s.contract.ScrapProof(ctx, req.StreamContractAddress, profileID, chunkID)
	if err != nil {
		logger.
			WithError(err).
			Error("failed to validate proof")
		return nil, rpc.ErrRpcInternal
	}

	err = s.contract.WaitMinedAndCheck(tx)
	if err != nil {
		logger.
			WithError(err).
			WithField("call", "ValidateProof").
			WithField("tx", tx.Hash().String()).
			Error("failed to wait mined")
		return nil, rpc.ErrRpcInternal
	}

	return &v1.ScrapProofResponse{
		TxId: []byte(tx.Hash().String()),
	}, nil
}

func (s *Server) ListWorkers(ctx context.Context, req *protoempty.Empty) (*v1.ListWorkersResponse, error) {
	_ = opentracing.SpanFromContext(ctx)

	s.logger.Info("list workers")

	workers, err := s.staking.GetAllTranscoders(context.Background())
	if err != nil {
		return nil, rpc.ErrRpcInternal
	}

	resp := &v1.ListWorkersResponse{
		Items: []*v1.WorkerResponse{},
	}

	for _, worker := range workers {
		resp.Items = append(resp.Items, &v1.WorkerResponse{
			Address:        worker.Address.Hex(),
			State:          uint32(worker.State),
			TotalStake:     worker.TotalStake.Bytes(),
			SelfStake:      worker.SelfStake.Bytes(),
			DelegatedStake: worker.DelegatedStake.Bytes(),
			RegisteredAt:   worker.Timestamp,
		})
	}

	return resp, nil
}
