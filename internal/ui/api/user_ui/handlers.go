package user_ui

import (
	"study-chat/generated/protobuf"
	userdmn "study-chat/internal/domain/user_dmn"
)

type UserServer struct {
	protobuf.UnimplementedUserServiceServer
	repo userdmn.UserRepository
}

func SetupServer(repo userdmn.UserRepository) UserServer {
	return UserServer{
		repo: repo,
	}
}
