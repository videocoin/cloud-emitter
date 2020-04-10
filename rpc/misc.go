package rpc

import (
	"context"

	streamsv1 "github.com/videocoin/cloud-api/streams/v1"
)

func (s *Server) markStreamAsFailed(streamID string) error {
	_, err := s.streams.UpdateStatus(context.Background(), &streamsv1.UpdateStreamRequest{
		Id:     streamID,
		Status: streamsv1.StreamStatusFailed,
	})

	return err
}
