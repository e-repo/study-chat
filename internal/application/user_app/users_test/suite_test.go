package users_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"study-chat/internal/application"
)

type UsersSuite struct {
	suite.Suite
	application.ServiceLocator
	GRPCHandlers users_app.UserHandlers
}

func (s *UsersSuite) SetupTest() {
	s.ServiceLocator.SetupTest()
	s.GRPCHandlers = users_app.SetupHandlers(s.UsersRepo)
}

func TestUsersSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UsersSuite))
}
