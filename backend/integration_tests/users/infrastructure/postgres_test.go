package infrastructure_test

import (
	"context"
	"study-chat/internal/infra/service"
	"testing"

	"study-chat/pkg/postgres"

	infra "study-chat/internal/infra/user_infra"

	"github.com/stretchr/testify/suite"
)

type PostgresRepoSuite struct {
	suite.Suite
	repo *infra.PostgresRepo
}

func (suite *PostgresRepoSuite) SetupSuite() {
	cfg, err := service.LoadConfig()
	if err != nil {
		suite.Fail("Failed to load config", err)
	}
	connData, err := postgres.NewConnectionData(
		cfg.Postgres.Hosts,
		cfg.Postgres.Database,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Port,
		cfg.Postgres.SSL,
	)
	if err != nil {
		suite.Fail("Failed to init postgres connection data", err)
	}
	cluster, err := postgres.InitCluster(context.Background(), connData)
	if err != nil {
		suite.Fail("Failed to init postgres cluster", err)
	}
	suite.repo = infra.NewPostgres(cluster)
}

//func (suite *PostgresRepoSuite) TestUserCRUD() {
//	email := "test@test.com"
//	created, err := suite.repo.CreateUser(context.Background(), email, func() (*domain.User, error) {
//		return domain.CreateUser("test", email)
//	})
//	suite.Require().NoError(err)
//
//	gotten, err := suite.repo.GetUserById(context.Background(), created.Id())
//	suite.Require().NoError(err)
//	suite.Require().Equal(created, gotten)
//
//	updated, err := suite.repo.UpdateUser(context.Background(), created.Id(), func(u *domain.User) (bool, error) {
//		err := u.ChangeEmail("test@test2.com")
//		return true, err
//	})
//	suite.Require().NoError(err)
//
//	gotten, err = suite.repo.GetUserById(context.Background(), created.Id())
//	suite.Require().NoError(err)
//	suite.Require().Equal(updated, gotten)
//
//	err = suite.repo.DeleteUser(context.Background(), created.Id())
//	suite.Require().NoError(err)
//}

func TestPostgresRepoSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(PostgresRepoSuite))
}
