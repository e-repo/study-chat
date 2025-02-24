package api

import (
	"google.golang.org/grpc"
	"study-chat/generated/protobuf"
	"study-chat/internal/auth"
	"study-chat/pkg/locator"
)

type gRPCServer struct {
	auth.Auth
}

func SetupGRPCServer(locator locator.ServiceLocator) *grpc.Server {
	s := grpc.NewServer()

	server := gRPCServer{}
	server.Auth = auth.CreateAuth(locator)

	protobuf.RegisterUserServiceServer(s, server)

	return s
}
