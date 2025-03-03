package auth

import (
	"errors"
)

var (
	ErrUserNotFound     = errors.New("пользователь не найден")
	ErrUserAlreadyExist = errors.New("данный пользователь уже существует")
	ErrInvalidPassword  = errors.New("пароль неверный")
)
