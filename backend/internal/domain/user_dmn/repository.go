package user_dmn

import (
	"context"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, updateFn func(*User) (bool, error)) (*User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CheckUserExist(ctx context.Context, email string) (bool, error)
}
