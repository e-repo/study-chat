package chat

import (
	"context"
	"encoding/json"
	"errors"
	hasql "golang.yandex/hasql/sqlx"
	"time"
)

var ErrSaveMessage = errors.New("ошибка добавления сообщения")

type Message struct {
	Id        int             `db:"id"`
	Method    string          `db:"method"`
	Payload   json.RawMessage `db:"payload"`
	Partition int8            `db:"partition"`
	CreatedAt time.Time       `db:"created_at"`
}

type MessageOutboxRepository interface {
	AddMessage(ctx context.Context, message *Message) (*Message, error)
}

type PostgresRepo struct {
	cluster *hasql.Cluster
}

func newMessageOutboxRepository(cluster *hasql.Cluster) MessageOutboxRepository {
	return &PostgresRepo{
		cluster: cluster,
	}
}

func (r *PostgresRepo) AddMessage(
	ctx context.Context,
	message *Message,
) (*Message, error) {
	var messageId int
	db := r.cluster.Primary().DBx()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errors.Join(ErrSaveMessage, err)
	}

	query := `
		INSERT INTO centrifugo_outbox (method, payload, partition) 
		VALUES ($1, $2, $3) 
		RETURNING id
	`

	err = tx.
		QueryRowxContext(ctx, query, message.Method, message.Payload, message.Partition).
		Scan(&messageId)

	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return nil, errors.Join(ErrSaveMessage, rollbackErr)
		}

		return nil, errors.Join(ErrSaveMessage, err)
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		return nil, errors.Join(ErrSaveMessage, commitErr)
	}

	message.Id = int(messageId)
	return message, nil
}
