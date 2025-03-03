package chat

import (
	"context"
	"encoding/json"
)

const (
	DefaultPublish   string = "publish"
	DefaultPartition int8   = 1
)

type newMessage struct {
	Method    string
	Partition int8
	Payload   json.RawMessage
}

type chatService struct {
	outboxRepo MessageOutboxRepository
}

func newChatService(outboxRepo MessageOutboxRepository) *chatService {
	return &chatService{outboxRepo: outboxRepo}
}

func (s *chatService) addMessage(
	ctx context.Context,
	message *newMessage,
) (*int, error) {
	msg, err := s.outboxRepo.AddMessage(
		ctx,
		&Message{
			Method:    DefaultPublish,
			Partition: DefaultPartition,
			Payload:   message.Payload,
		},
	)
	if err != nil {
		return nil, err
	}

	return &msg.Id, nil
}
