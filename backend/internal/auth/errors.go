package auth

import "errors"

var (
	ErrUserNotFound     = errors.New("пользователь не найден")
	ErrInvalidUser      = errors.New("не валидный пользователь")
	ErrUserInit         = errors.New("ошибка инициализации пользователя")
	ErrUserValidation   = errors.New("ошибка валидации пользователя")
	ErrUserAlreadyExist = errors.New("данный пользователь уже существует")
	ErrInvalidPassword  = errors.New("пароль неверный")
)
