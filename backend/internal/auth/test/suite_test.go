package test

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"study-chat/internal/api"
	"study-chat/internal/auth"
	"testing"
)

type UserSuite struct {
	suite.Suite
	Echo *echo.Echo
}

func (s *UserSuite) SetupTest() {
	echoInst, server := api.SetupRESTTestServer(s.T())

	initUserRepositoryMock(server.Auth.Repo)
	s.Echo = echoInst
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func initUserRepositoryMock(userRepo auth.UserRepository) {
	m, ok := userRepo.(*auth.MockUserRepository)
	if !ok {
		panic("MockUserRepository не имплементирует auth.UserRepository")
	}

	m.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, user any) (auth.User, error) {
			return auth.User{
				Id:        uuid.New(),
				FirstName: user.(*auth.User).FirstName,
				Email:     user.(*auth.User).Email,
				Password:  user.(*auth.User).Password,
			}, nil
		}).
		AnyTimes()
}
