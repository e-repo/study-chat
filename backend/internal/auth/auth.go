package auth

import (
	hasql "golang.yandex/hasql/sqlx"
	"study-chat/generated/protobuf"
	"study-chat/internal/config"
	"study-chat/pkg/locator"
	"study-chat/pkg/validator"
)

type Auth struct {
	protobuf.UnimplementedUserServiceServer
	repo          userRepository
	service       *userService
	validator     validator.Validator
	hmacSecretKey string
}

func SetupEndpoints(locator locator.ServiceLocator) Auth {
	cfg := locator.Get(config.ConfigServiceKey).(config.Config)

	cluster := locator.Get(config.ClusterServiceKey).(*hasql.Cluster)
	userRepo := newUserRepository(cluster)
	userService := newUserService(userRepo)
	validate := locator.Get(config.ValidatorServiceKey).(validator.Validator)

	hmacSecretKey := cfg.Server.HmacSecretKey

	return Auth{
		repo:          userRepo,
		service:       userService,
		validator:     validate,
		hmacSecretKey: hmacSecretKey,
	}
}
