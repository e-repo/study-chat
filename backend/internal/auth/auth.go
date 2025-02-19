package auth

import (
	hasql "golang.yandex/hasql/sqlx"
	"study-chat/internal/config"
	"study-chat/pkg/locator"
	"study-chat/pkg/validator"
)

type Auth struct {
	repo          UserRepository
	service       *UserService
	validator     validator.Validator
	hmacSecretKey string
}

func SetupEndpoints(locator locator.ServiceLocator) Auth {
	cfg := locator.Get(config.ConfigServiceKey).(config.Config)

	cluster := locator.Get(config.ClusterServiceKey).(*hasql.Cluster)
	userRepo := NewUserRepository(cluster)
	userService := NewUserService(userRepo)
	validate := locator.Get(config.ValidatorServiceKey).(validator.Validator)

	hmacSecretKey := cfg.Server.HmacSecretKey

	return Auth{
		repo:          userRepo,
		service:       userService,
		validator:     validate,
		hmacSecretKey: hmacSecretKey,
	}
}
