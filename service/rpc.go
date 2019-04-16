package service

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"net"

	accountsv1 "github.com/VideoCoin/cloud-api/accounts/v1"
	v1 "github.com/VideoCoin/cloud-api/emitter/v1"
	"github.com/VideoCoin/cloud-api/rpc"
	streamsv1 "github.com/VideoCoin/cloud-api/streams/v1"
	"github.com/VideoCoin/cloud-pkg/auth"
	"github.com/VideoCoin/cloud-pkg/grpcutil"
	"github.com/VideoCoin/common/bcops"
	sm "github.com/VideoCoin/common/streamManager"
	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
	"github.com/VideoCoin/go-videocoin/accounts/keystore"
	"github.com/VideoCoin/go-videocoin/common"
	"github.com/VideoCoin/go-videocoin/ethclient"
	protoempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RpcServerOptions struct {
	Addr            string
	NodeRPCAddr     string
	ContractAddress string
	Logger          *logrus.Entry
	Accounts        accountsv1.AccountServiceClient
	EB              *EventBus

	KeyManager       string
	SecretKeyManager string
	SecretKeyClient  string
	Secret           string
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

	keyManager       string
	secretKeyManager string
	secretKeyClient  string
	secret           string
}

func NewRpcServer(opts *RpcServerOptions) (*RpcServer, error) {
	grpcOpts := grpcutil.DefaultServerOpts(opts.Logger)
	grpcServer := grpc.NewServer(grpcOpts...)

	listen, err := net.Listen("tcp", opts.Addr)
	if err != nil {
		return nil, err
	}

	client, err := ethclient.Dial(opts.NodeRPCAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial eth client: %s", err.Error())
	}

	managerAddress := common.HexToAddress(opts.ContractAddress)
	manager, err := sm.NewManager(managerAddress, client)
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
		streamManager: manager,
		eventListener: eventListener,

		keyManager:       opts.KeyManager,
		secretKeyManager: opts.SecretKeyManager,
		secretKeyClient:  opts.SecretKeyClient,
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

func (s *RpcServer) RequestStream(ctx context.Context, req *protoempty.Empty) (*v1.RequestStreamResponse, error) {
	transactOpts, _, err := s.authenticate(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	streamID := big.NewInt(int64(rand.Intn(math.MaxInt64)))

	_, err = s.streamManager.RequestStream(
		transactOpts,
		streamID,
		"videocoin",
		[]*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2)},
	)
	if err != nil {
		s.logger.Error(err)
		return nil, rpc.ErrRpcInternal
	}

	return &v1.RequestStreamResponse{
		Address:  transactOpts.From.Hex(),
		StreamId: streamID.Uint64(),
	}, nil
}

func (s *RpcServer) CreateStream(ctx context.Context, req *v1.CreateStreamRequest) (*protoempty.Empty, error) {
	err := req.Validate()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	transactOpts, ctx, err := s.authenticate(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	id, ok := auth.StreamIDFromContext(ctx)
	if !ok {
		s.logger.Error(err)
		return nil, err
	}

	userID, ok := auth.UserIDFromContext(ctx)
	if !ok {
		s.logger.Error(err)
		return nil, err
	}

	streamID := new(big.Int).SetUint64(req.StreamId)

	// todo: constant value ???
	i, e := big.NewInt(10), big.NewInt(19)
	transactOpts.Value = i.Exp(i, e, nil)

	_, err = s.streamManager.CreateStream(
		transactOpts,
		streamID,
	)
	if err != nil {
		s.logger.Error(err)
		return nil, rpc.ErrRpcInternal
	}

	go func() {
		resultCh, errCh := s.eventListener.LogStreamCreateEvent(streamID)

		select {
		case err := <-errCh:
			s.logger.Error(err)
			err = s.eb.UpdateStreamStatus(
				&streamsv1.UpdateStreamRequest{
					Id:     id,
					UserId: userID,
					Status: streamsv1.StreamStatusFailed,
				})
			if err != nil {
				s.logger.Error(err)
			}
			return
		case e := <-resultCh:
			err := s.eb.UpdateStreamAddress(
				&streamsv1.UpdateStreamRequest{
					Id:            id,
					UserId:        userID,
					StreamAddress: e.StreamAddress.Hex(),
				})
			if err != nil {
				s.logger.Error(err)
			}
			return
		}
	}()

	return new(protoempty.Empty), nil
}

func (s *RpcServer) authenticate(ctx context.Context) (*bind.TransactOpts, context.Context, error) {
	ctx = auth.NewContextWithSecretKey(ctx, s.secret)
	ctx, err := auth.AuthFromContext(ctx)
	if err != nil {
		return nil, ctx, rpc.ErrRpcUnauthenticated
	}

	userID, ok := auth.UserIDFromContext(ctx)
	if !ok {
		return nil, ctx, rpc.ErrRpcUnauthenticated
	}

	keyReq := &accountsv1.AccountRequest{OwnerID: userID}
	key, err := s.accounts.Key(ctx, keyReq)
	if err != nil {
		return nil, ctx, rpc.ErrRpcUnauthenticated
	}

	decrypted, err := keystore.DecryptKey([]byte(key.Key), s.secretKeyClient)
	if err != nil {
		return nil, ctx, err
	}

	transactOpts, err := bcops.GetBCAuth(s.ethClient, decrypted)
	if err != nil {
		return nil, ctx, rpc.ErrRpcUnauthenticated
	}

	from := common.HexToAddress(key.Address)
	transactOpts.From = from

	return transactOpts, ctx, nil
}
