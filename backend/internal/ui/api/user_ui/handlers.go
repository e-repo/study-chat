package user_ui

import (
	userdmn "study-chat/internal/domain/user_dmn"
	"study-chat/internal/infra/service"
)

type UserLocator interface {
	Get(key string) interface{}
}

type UserEndpoints struct {
	//protobuf.UnimplementedUserServiceServer
	locator UserLocator
}

func SetupEndpoints(locator UserLocator) UserEndpoints {
	return UserEndpoints{
		locator: locator,
	}
}

func getConfig(locator UserLocator) service.Config {
	cfg := locator.Get(service.ConfigServiceKey).(service.Config)
	return cfg
}

func getValidatorService(locator UserLocator) service.Validator {
	validatorSrv := locator.Get(service.ValidatorServiceKey).(service.Validator)
	return validatorSrv
}

func getUserRepo(locator UserLocator) userdmn.UserRepository {
	userRepo := locator.Get(service.UserRepositoryServiceKey).(userdmn.UserRepository)
	return userRepo
}
