package service

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"net"
	"time"

	accountsv1 "github.com/VideoCoin/cloud-api/accounts/v1"
	v1 "github.com/VideoCoin/cloud-api/emitter/v1"
	pipelinesv1 "github.com/VideoCoin/cloud-api/pipelines/v1"
	"github.com/VideoCoin/cloud-api/rpc"
	"github.com/VideoCoin/cloud-pkg/bcops"
	"github.com/VideoCoin/cloud-pkg/grpcutil"
	sm "github.com/VideoCoin/cloud-pkg/streamManager"
	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
	"github.com/VideoCoin/go-videocoin/accounts/keystore"
	"github.com/VideoCoin/go-videocoin/common"
	"github.com/VideoCoin/go-videocoin/ethclient"
	protoempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type RpcServerOptions struct {
	Addr         string
	NodeHTTPAddr string
	ContractAddr string
	Logger       *logrus.Entry
	Accounts     accountsv1.AccountServiceClient
	EB           *EventBus

	Secret  string
	MKey    string
	MSecret string
}

type RpcServer struct {
	addr string

	grpc          *grpc.Server
	listen        net.Listener
	logger        *logrus.Entry
	eb            *EventBus
	accounts      accountsv1.AccountServiceClient
	ethClient     *ethclient.Client
	streamManager *sm.Manager
	eventListener *EventListener

	secret  string
	mKey    string
	mSecret string
}

func NewRpcServer(opts *RpcServerOptions) (*RpcServer, error) {
	grpcOpts := grpcutil.DefaultServerOpts(opts.Logger)
	grpcServer := grpc.NewServer(grpcOpts...)

	listen, err := net.Listen("tcp", opts.Addr)
	if err != nil {
		return nil, err
	}

	ethClient, err := ethclient.Dial(opts.NodeHTTPAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial eth client: %s", err.Error())
	}

	managerAddress := common.HexToAddress(opts.ContractAddr)
	manager, err := sm.NewManager(managerAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create smart contract stream manager: %s", err.Error())
	}

	eventListenerConfig := &EventListenerConfig{
		StreamManager: manager,
		Timeout:       60,
		Logger:        opts.Logger,
	}
	eventListener := NewEventListener(eventListenerConfig)

	rpcServer := &RpcServer{
		addr:          opts.Addr,
		grpc:          grpcServer,
		listen:        listen,
		logger:        opts.Logger,
		eb:            opts.EB,
		accounts:      opts.Accounts,
		ethClient:     ethClient,
		streamManager: manager,
		eventListener: eventListener,
		secret:        opts.Secret,
		mKey:          opts.MKey,
		mSecret:       opts.MSecret,
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

func (s *RpcServer) RequestStream(ctx context.Context, req *v1.StreamRequest) (*protoempty.Empty, error) {
	err := req.Validate()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	transactOpts, err := s.getClientTransactOpts(ctx, req.UserId)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	userId := req.UserId
	pipelineId := req.PipelineId
	streamId := big.NewInt(int64(rand.Intn(math.MaxInt64)))
	clientAddress := transactOpts.From

	go func() {
		s.logger.Infof("request stream on stream id %d", streamId.Uint64())
		_, err = s.streamManager.RequestStream(
			transactOpts,
			streamId,
			"videocoin",
			[]*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2)},
		)
		if err != nil {
			s.logger.Errorf("failed to request stream: %s", err)
		}

		resultCh, errCh := s.eventListener.LogStreamRequestEvent(
			streamId, clientAddress)

		s.logger.Infof("log stream request event on stream id %d", streamId.Uint64())
		select {
		case err := <-errCh:
			s.logger.Error(err)
			err = s.eb.UpdatePipelineStatus(
				&pipelinesv1.UpdatePipelineRequest{
					Id:     pipelineId,
					UserId: userId,
					Status: pipelinesv1.PipelineStatusFailed,
				})
			if err != nil {
				s.logger.Error(err)
			}
			return
		case e := <-resultCh:
			err := s.eb.UpdatePipelineStatus(
				&pipelinesv1.UpdatePipelineRequest{
					Id:            pipelineId,
					UserId:        userId,
					StreamId:      e.StreamID.Uint64(),
					StreamAddress: e.StreamAddress.Hex(),
					Status:        pipelinesv1.PipelineStatusApprovePending,
				})
			if err != nil {
				s.logger.Error(err)
			}
			return
		}
	}()

	return &protoempty.Empty{}, nil
}

func (s *RpcServer) ApproveStream(ctx context.Context, req *v1.StreamRequest) (*protoempty.Empty, error) {
	err := req.Validate()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	transactOpts, err := s.getManagerTransactOpts(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	userId := req.UserId
	pipelineId := req.PipelineId
	streamId := new(big.Int).SetUint64(req.StreamId)

	go func() {
		s.logger.Infof("allow refund on stream id %d", streamId.Uint64())
		_, err := s.streamManager.AllowRefund(transactOpts, streamId)
		if err != nil {
			s.logger.Errorf("failed to allow refund: %s", err)
		}

		s.logger.Infof("approve stream creation on stream id %d", streamId.Uint64())
		_, err = s.streamManager.ApproveStreamCreation(
			transactOpts,
			streamId,
			nil,
		)
		if err != nil {
			s.logger.Errorf("failed to approve stream: %s", err)
		}

		s.logger.Infof("log stream approve event on stream id %d", streamId.Uint64())
		resultCh, errCh := s.eventListener.LogStreamApproveEvent(streamId)
		select {
		case err := <-errCh:
			s.logger.Error(err)
			err = s.eb.UpdatePipelineStatus(
				&pipelinesv1.UpdatePipelineRequest{
					Id:     pipelineId,
					UserId: userId,
					Status: pipelinesv1.PipelineStatusFailed,
				})
			if err != nil {
				s.logger.Error(err)
			}
			return
		case e := <-resultCh:
			err := s.eb.UpdatePipelineStatus(
				&pipelinesv1.UpdatePipelineRequest{
					Id:       pipelineId,
					UserId:   userId,
					StreamId: e.StreamID.Uint64(),
					Status:   pipelinesv1.PipelineStatusCreatePending,
				})
			if err != nil {
				s.logger.Error(err)
			}
			return
		}
	}()

	return &protoempty.Empty{}, nil
}

func (s *RpcServer) CreateStream(ctx context.Context, req *v1.StreamRequest) (*protoempty.Empty, error) {
	err := req.Validate()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	transactOpts, err := s.getClientTransactOpts(ctx, req.UserId)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	userId := req.UserId
	pipelineId := req.PipelineId
	streamId := new(big.Int).SetUint64(req.StreamId)

	// todo: constant value ???
	i, e := big.NewInt(10), big.NewInt(19)
	transactOpts.Value = i.Exp(i, e, nil)

	go func() {
		s.logger.Infof("create stream on stream id %d", streamId.Uint64())
		_, err = s.streamManager.CreateStream(
			transactOpts,
			streamId,
		)
		if err != nil {
			s.logger.Errorf("failed to create stream: %s", err)
		}

		resultCh, errCh := s.eventListener.LogStreamCreateEvent(streamId)

		select {
		case err := <-errCh:
			s.logger.Error(err)
			err = s.eb.UpdatePipelineStatus(
				&pipelinesv1.UpdatePipelineRequest{
					Id:     pipelineId,
					UserId: userId,
					Status: pipelinesv1.PipelineStatusFailed,
				})
			if err != nil {
				s.logger.Error(err)
			}
			return
		case e := <-resultCh:
			err := s.eb.UpdatePipelineStatus(
				&pipelinesv1.UpdatePipelineRequest{
					Id:            pipelineId,
					UserId:        userId,
					StreamId:      e.StreamID.Uint64(),
					StreamAddress: e.StreamAddress.Hex(),
					Status:        pipelinesv1.PipelineStatusJobPending,
				})
			if err != nil {
				s.logger.Error(err)
			}
			return
		}
	}()

	return &protoempty.Empty{}, nil
}

func (s *RpcServer) getClientTransactOpts(ctx context.Context, userID string) (*bind.TransactOpts, error) {
	keyReq := &accountsv1.AccountRequest{OwnerID: userID}
	key, err := s.accounts.Key(ctx, keyReq)
	if err != nil {
		return nil, err
	}

	decrypted, err := keystore.DecryptKey([]byte(key.Key), s.secret)
	if err != nil {
		return nil, err
	}

	transactOpts, err := bcops.GetBCAuth(s.ethClient, decrypted)
	if err != nil {
		return nil, err
	}

	from := common.HexToAddress(key.Address)
	transactOpts.From = from

	return transactOpts, nil
}

func (s *RpcServer) getManagerTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	decrypted, err := keystore.DecryptKey([]byte(s.mKey), s.mSecret)
	if err != nil {
		return nil, err
	}

	transactOpts, err := bcops.GetBCAuth(s.ethClient, decrypted)
	if err != nil {
		return nil, err
	}

	return transactOpts, nil
}
