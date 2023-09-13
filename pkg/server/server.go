package server

import (
	"context"
	"grpc-echo-server/pkg/pb"
	"log/slog"
)

type Server struct{}

func (s *Server) Do(ctx context.Context, request *pb.Request) (response *pb.Response, err error) {
	slog.Info("Received", "msg", request.Message)
	return &pb.Response{
		Message: request.Message,
	}, nil
}
