package user_dmn

import "github.com/google/uuid"

type UserDto struct {
	Id       uuid.UUID
	FistName string
	Email    string
	PassHash string
}
