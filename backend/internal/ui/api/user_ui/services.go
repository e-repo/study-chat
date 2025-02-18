package user_ui

import (
	hasql "golang.yandex/hasql/sqlx"
	userdmn "study-chat/internal/domain/user_dmn"
	"study-chat/internal/infra/conf"
	"study-chat/internal/infra/service"
	userinfra "study-chat/internal/infra/user_infra"
	"study-chat/pkg/locator"
)

type Services struct {
	//protobuf.UnimplementedUserServiceServer
	locator locator.ServiceLocator
}

func SetupEndpoints(locator locator.ServiceLocator) Services {
	return Services{
		locator: locator,
	}
}

func (s Services) HmacSecretKey() string {
	cfg := s.locator.Get(conf.ConfigServiceKey).(conf.ServerConfig)
	return cfg.Server.HmacSecretKey
}

func (s Services) Validator() service.Validator {
	validatorSrv := s.locator.Get(conf.ValidatorServiceKey).(service.Validator)
	return validatorSrv
}

func (s Services) UserRepo() userdmn.UserRepository {
	cluster := s.locator.Get(conf.ClusterServiceKey).(*hasql.Cluster)
	return userinfra.NewPostgres(cluster)
}
