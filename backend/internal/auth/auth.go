package auth

import (
	"go.uber.org/mock/gomock"
	hasql "golang.yandex/hasql/sqlx"
	"study-chat/generated/protobuf"
	"study-chat/internal/config"
	"study-chat/pkg/locator"
	"study-chat/pkg/validator"
	"testing"
)

type Auth struct {
	protobuf.UnimplementedUserServiceServer
	Repo          UserRepository
	service       *userService
	validator     validator.Validator
	hmacSecretKey string
}

func CreateAuth(locator locator.ServiceLocator) Auth {
	cfg := locator.Get(config.ConfigServiceKey).(config.Config)

	cluster := locator.Get(config.ClusterServiceKey).(*hasql.Cluster)
	userRepo := newUserRepository(cluster)
	userService := newUserService(userRepo)
	validate := validator.NewRuValidator()

	hmacSecretKey := cfg.Server.HmacSecretKey

	return Auth{
		Repo:          userRepo,
		service:       userService,
		validator:     validate,
		hmacSecretKey: hmacSecretKey,
	}
}

func CreateTestAuth(t *testing.T) Auth {
	ctrl := gomock.NewController(t)
	userRepo := NewMockUserRepository(ctrl)

	return Auth{
		Repo:          userRepo,
		service:       newUserService(userRepo),
		validator:     validator.NewRuValidator(),
		hmacSecretKey: "hmac-secret-key",
	}
}
