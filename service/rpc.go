package service

import (
	"context"
	"fmt"
	"math/big"
	"net"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	protoempty "github.com/gogo/protobuf/types"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	accountsv1 "github.com/videocoin/cloud-api/accounts/v1"
	v1 "github.com/videocoin/cloud-api/emitter/v1"
	"github.com/videocoin/cloud-api/rpc"
	"github.com/videocoin/cloud-pkg/bcops"
	"github.com/videocoin/cloud-pkg/grpcutil"
	sm "github.com/videocoin/cloud-pkg/streamManager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RpcServerOptions struct {
	Addr            string
	RPCNodeHTTPAddr string
	ContractAddr    string
	Logger          *logrus.Entry
	Accounts        accountsv1.AccountServiceClient
	ClientSecret    string
	ManagerKey      string
	ManagerSecret   string
}

type RpcServer struct {
	addr          string
	grpc          *grpc.Server
	listen        net.Listener
	logger        *logrus.Entry
	accounts      accountsv1.AccountServiceClient
	ethClient     *ethclient.Client
	streamManager *sm.Manager

	clientSecret  string
	managerKey    string
	managerSecret string
}

func NewRpcServer(opts *RpcServerOptions) (*RpcServer, error) {
	grpcOpts := grpcutil.DefaultServerOpts(opts.Logger)
	grpcServer := grpc.NewServer(grpcOpts...)

	listen, err := net.Listen("tcp", opts.Addr)
	if err != nil {
		return nil, err
	}

	ethClient, err := ethclient.Dial(opts.RPCNodeHTTPAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial eth client: %s", err.Error())
	}

	managerAddress := common.HexToAddress(opts.ContractAddr)
	manager, err := sm.NewManager(managerAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create smart contract stream manager: %s", err.Error())
	}

	rpcServer := &RpcServer{
		addr:          opts.Addr,
		grpc:          grpcServer,
		listen:        listen,
		logger:        opts.Logger,
		accounts:      opts.Accounts,
		ethClient:     ethClient,
		streamManager: manager,
		clientSecret:  opts.ClientSecret,
		managerKey:    opts.ManagerKey,
		managerSecret: opts.ManagerSecret,
	}

	v1.RegisterEmitterServiceServer(grpcServer, rpcServer)
	reflection.Register(grpcServer)

	return rpcServer, nil
}

func (s *RpcServer) Start() error {
	s.logger.Infof("starting rpc server on %s", s.addr)
	return s.grpc.Serve(s.listen)
}

func (s *RpcServer) Health(ctx context.Context, req *protoempty.Empty) (*rpc.HealthStatus, error) {
	return &rpc.HealthStatus{Status: "OK"}, nil
}

func (s *RpcServer) RequestStream(ctx context.Context, req *v1.StreamRequest) (*v1.Tx, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RequestStream")
	defer span.Finish()

	span.SetTag("user_id", req.UserId)
	span.SetTag("stream_id", fmt.Sprintf("%d", req.StreamId))

	transactOpts, err := s.getClientTransactOpts(ctx, req.UserId)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	streamID := big.NewInt(int64(req.StreamId))

	s.logger.Infof("request stream on stream id %d", streamID.Uint64())
	tx, err := s.streamManager.RequestStream(
		transactOpts,
		streamID,
		req.ProfileNames,
	)

	if err != nil {
		s.logger.Errorf("failed to request stream: %s", err.Error())
		return nil, err
	}

	return &v1.Tx{Hash: tx.Hash().Bytes()}, nil
}

func (s *RpcServer) ApproveStream(ctx context.Context, req *v1.StreamRequest) (*v1.Tx, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ApproveStream")
	defer span.Finish()

	span.SetTag("user_id", req.UserId)
	span.SetTag("stream_id", fmt.Sprintf("%d", req.StreamId))

	transactOpts, err := s.getManagerTransactOpts(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	streamID := new(big.Int).SetUint64(req.StreamId)

	s.logger.Infof("allow refund on stream id %d", streamID.Uint64())
	_, err = s.streamManager.AllowRefund(transactOpts, streamID)
	if err != nil {
		s.logger.Errorf("failed to allow refund: %s", err)
	}

	s.logger.Infof("approve stream creation on stream id %d", streamID.Uint64())
	tx, err := s.streamManager.ApproveStreamCreation(
		transactOpts,
		streamID,
	)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return &v1.Tx{Hash: tx.Hash().Bytes()}, nil
}

func (s *RpcServer) CreateStream(ctx context.Context, req *v1.StreamRequest) (*v1.Tx, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateStream")
	defer span.Finish()

	span.SetTag("user_id", req.UserId)
	span.SetTag("stream_id", fmt.Sprintf("%d", req.StreamId))

	transactOpts, err := s.getClientTransactOpts(ctx, req.UserId)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	streamID := new(big.Int).SetUint64(req.StreamId)

	// todo: constant value ???
	i, e := big.NewInt(10), big.NewInt(19)
	transactOpts.Value = i.Exp(i, e, nil)

	s.logger.Infof("create stream on stream id %d", streamID.Uint64())
	tx, err := s.streamManager.CreateStream(
		transactOpts,
		streamID,
	)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return &v1.Tx{Hash: tx.Hash().Bytes()}, nil
}

func (s *RpcServer) EndStream(ctx context.Context, req *v1.StreamRequest) (*v1.Tx, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EndStream")
	defer span.Finish()

	span.SetTag("user_id", req.UserId)
	span.SetTag("stream_id", fmt.Sprintf("%d", req.StreamId))

	transactOpts, err := s.getClientTransactOpts(ctx, req.UserId)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	streamID := new(big.Int).SetUint64(req.StreamId)

	s.logger.Infof("end stream on stream id %d", streamID.Uint64())
	tx, err := s.streamManager.EndStream(
		transactOpts,
		streamID,
	)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return &v1.Tx{Hash: tx.Hash().Bytes()}, nil
}

func (s *RpcServer) getClientTransactOpts(ctx context.Context, userID string) (*bind.TransactOpts, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getClientTransactOpts")
	defer span.Finish()

	keyReq := &accountsv1.AccountRequest{OwnerId: userID}
	key, err := s.accounts.Key(ctx, keyReq)
	if err != nil {
		return nil, err
	}

	decrypted, err := keystore.DecryptKey([]byte(key.Key), s.clientSecret)
	if err != nil {
		return nil, err
	}

	transactOpts, err := bcops.GetBCAuth(s.ethClient, decrypted)
	if err != nil {
		return nil, err
	}

	from := common.HexToAddress(key.Address)
	transactOpts.From = from
	transactOpts.GasPrice = big.NewInt(10000000000)

	return transactOpts, nil
}

func (s *RpcServer) getManagerTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "getManagerTransactOpts")
	defer span.Finish()

	decrypted, err := keystore.DecryptKey([]byte(s.managerKey), s.managerSecret)
	if err != nil {
		return nil, err
	}

	transactOpts, err := bcops.GetBCAuth(s.ethClient, decrypted)
	if err != nil {
		return nil, err
	}

	return transactOpts, nil
}
