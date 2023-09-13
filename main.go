package main

import (
	"fmt"
	"grpc-echo-server/pkg/pb"
	"grpc-echo-server/pkg/server"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	port := "8080"
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEchoServer(grpcServer, &server.Server{})
	slog.Info("GRPC server started on", "port", port)
	grpcServer.Serve(listener)
}
