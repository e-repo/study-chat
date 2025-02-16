package user_app

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	userdmn "study-chat/internal/domain/user_dmn"
)

func CreateUser(
	command *CreateUserCommand,
	ctx context.Context,
	r userdmn.UserRepository,
) (*userdmn.UserDto, error) {
	password := []byte(command.Password())

	passHash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userDto := userdmn.UserDto{
		FistName: command.FirstName(),
		Email:    command.Email(),
		PassHash: string(passHash),
	}

	user, err := userdmn.CreateUser(
		&userDto,
		r,
		ctx,
	)
	if err != nil {
		return nil, err
	}

	userDto.Id = user.ID()
	return &userDto, nil
}

//func (h UserServer) CreateUser(
//	ctx context.Context,
//	req *protobuf.CreateUserRequest,
//) (*protobuf.CreateUserResponse, error) {
//	email := req.GetEmail()
//	user, err := h.repo.CreateUser(ctx, email, func() (*userdmn.User, error) {
//		return userdmn.CreateUser(req.GetName(), email)
//	})
//	if err != nil {
//		if errors.Is(err, userdmn.ErrInvalidUser) || errors.Is(err, userdmn.ErrUserValidation) {
//			return nil, status.Error(codes.InvalidArgument, err.Error())
//		}
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &protobuf.CreateUserResponse{
//		Id: user.ID().String(),
//	}, nil
//}
