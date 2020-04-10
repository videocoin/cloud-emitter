package rpc

import (
	"context"
	"math/big"

	v1 "github.com/videocoin/cloud-api/emitter/v1"
)

func (s *Server) initStreamAsync(ctx context.Context, req *v1.InitStreamRequest) {
	go func(ctx context.Context, req *v1.InitStreamRequest) {
		s.initStream(ctx, req)
	}(ctx, req)
}

func (s *Server) endStreamAsync(ctx context.Context, req *v1.EndStreamRequest) {
	go func(ctx context.Context, req *v1.EndStreamRequest) {
		s.endStream(ctx, req)
	}(ctx, req)
}

func (s *Server) addInputChunkAsync(ctx context.Context, req *v1.AddInputChunkRequest) {
	go func(ctx context.Context, req *v1.AddInputChunkRequest) {
		s.addInputChunk(ctx, req)
	}(ctx, req)
}

func (s *Server) depositAsync(ctx context.Context, userID, streamID string, to *big.Int) {
	go func(ctx context.Context, userID, streamID string, to *big.Int) {
		s.deposit(ctx, userID, streamID, to)
	}(ctx, userID, streamID, to)
}
