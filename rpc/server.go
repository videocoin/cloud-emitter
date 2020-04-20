package rpc

import (
	"net"

	"github.com/videocoin/go-staking"

	"github.com/sirupsen/logrus"
	v1 "github.com/videocoin/cloud-api/emitter/v1"
	"github.com/videocoin/cloud-emitter/contract"
	"github.com/videocoin/cloud-pkg/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type ServerOpts struct {
	Logger   *logrus.Entry
	Addr     string
	Contract *contract.Client
	Staking  *staking.Client
}

type Server struct {
	logger   *logrus.Entry
	addr     string
	grpc     *grpc.Server
	listen   net.Listener
	contract *contract.Client
	staking  *staking.Client
}

func NewRPCServer(opts *ServerOpts) (*Server, error) {
	grpcOpts := grpcutil.DefaultServerOpts(opts.Logger)
	grpcServer := grpc.NewServer(grpcOpts...)

	healthService := health.NewServer()
	healthv1.RegisterHealthServer(grpcServer, healthService)

	listen, err := net.Listen("tcp", opts.Addr)
	if err != nil {
		return nil, err
	}

	rpcServer := &Server{
		logger:   opts.Logger,
		addr:     opts.Addr,
		grpc:     grpcServer,
		listen:   listen,
		contract: opts.Contract,
		staking:  opts.Staking,
	}

	v1.RegisterEmitterServiceServer(grpcServer, rpcServer)
	reflection.Register(grpcServer)

	return rpcServer, nil
}

func (s *Server) Start() error {
	return s.grpc.Serve(s.listen)
}

func (s *Server) Stop() {
	s.grpc.GracefulStop()
}
