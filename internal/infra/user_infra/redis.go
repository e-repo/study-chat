// Description: The file contains just the example of the Redis repository implementation and isn't used in the project.

package user_infra

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"

	"study-chat/internal/domain/user_dmn"
)

type redisUser struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
}

type redisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
}

type RedisRepo struct {
	client     redisClient
	expiration time.Duration
}

func NewRedis(clusterMode, tlsEnabled bool, addr, username, password string, expiration time.Duration) *RedisRepo {
	var (
		tlsConfig *tls.Config
		client    redisClient
	)
	if tlsEnabled {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true, //nolint:gosec // It's okay in intranet
		}
	}

	if clusterMode {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:     []string{addr},
			Username:  username,
			Password:  password,
			TLSConfig: tlsConfig,
		})
	} else {
		client = redis.NewClient(&redis.Options{
			Addr:      addr,
			Username:  username,
			Password:  password,
			TLSConfig: tlsConfig,
		})
	}

	return &RedisRepo{
		client:     client,
		expiration: expiration,
	}
}

func (r *RedisRepo) SaveUser(ctx context.Context, user user_dmn.User) error {
	ru := redisUser{
		FirstName: user.FirstName(),
		Email:     user.Email(),
	}
	val, err := json.Marshal(ru)
	if err != nil {
		return fmt.Errorf("failed to serialize users: %w", err)
	}

	err = r.client.Set(ctx, user.ID().String(), val, r.expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to save users to redis: %w", err)
	}
	return nil
}

//func (r *RedisRepo) GetUserById(ctx context.Context, id uuid.UUID) (*users.UserRepository, error) {
//	val, err := r.client.Get(ctx, id.String()).Result()
//	if err != nil {
//		if errors.Is(err, redis.Nil) {
//			return nil, users.ErrUserNotFound
//		}
//		return nil, fmt.Errorf("failed to get users from redis: %w", err)
//	}
//
//	var ru redisUser
//	err = json.Unmarshal([]byte(val), &ru)
//	if err != nil {
//		return nil, fmt.Errorf("failed to deserialize users: %w", err)
//	}
//
//	return users.NewUser(id, ru.FirstName(), ru.Email())
//}

func (r *RedisRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := r.client.Del(ctx, id.String()).Err()
	if err != nil {
		return fmt.Errorf("failed to delete users from redis: %w", err)
	}
	return nil
}
