package chat

import (
	hasql "golang.yandex/hasql/sqlx"
	"study-chat/internal/config"
	"study-chat/pkg/locator"
	"study-chat/pkg/validator"
)

type Chat struct {
	Repo      MessageOutboxRepository
	service   *chatService
	validator *validator.Validator
}

func CreateChat(locator locator.ServiceLocator) *Chat {
	cluster := locator.Get(config.ClusterServiceKey).(*hasql.Cluster)
	validate := locator.Get(config.ValidatorServiceKey).(*validator.Validator)
	messageRepo := newMessageOutboxRepository(cluster)

	return &Chat{
		Repo:      messageRepo,
		service:   newChatService(messageRepo),
		validator: validate,
	}
}
