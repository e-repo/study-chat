package user_dmn

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

var (
	ErrUserNotFound     = errors.New("пользователь не найден")
	ErrInvalidUser      = errors.New("не валидный пользователь")
	ErrUserInit         = errors.New("ошибка инициализации пользователя")
	ErrUserValidation   = errors.New("ошибка валидации пользователя")
	ErrUserAlreadyExist = errors.New("данный пользователь уже существует")
	ErrInvalidPassword  = errors.New("пароль неверный")
)

type User struct {
	id        uuid.UUID
	firstName string
	email     string
	passHash  string
}

func NewUser(user *UserDto) (*User, error) {
	return &User{
		id:        uuid.New(),
		firstName: user.FistName,
		email:     user.Email,
		passHash:  user.PassHash,
	}, nil
}

func CreateUser(user *UserDto, userRepo UserRepository, ctx context.Context) (*User, error) {
	err := checkEmail(user.Email, userRepo, ctx)
	if err != nil {
		return nil, ErrUserAlreadyExist
	}

	userEntity, err := NewUser(user)
	if err != nil {
		return nil, ErrUserInit
	}

	_, err = userRepo.CreateUser(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	return userEntity, nil
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) Email() string {
	return u.email
}

func (u *User) PassHash() string { return u.passHash }

func (u *User) SendToEmail(_ string) error {
	return errors.New("not implemented")
}

func (u *User) ChangeEmail(email string) error {
	u.email = email
	return nil
}

func checkEmail(email string, r UserRepository, ctx context.Context) error {
	exist, err := r.CheckUserExist(ctx, email)
	if err != nil {
		return fmt.Errorf("failed to check users exist: %w", err)
	}
	if exist {
		return ErrUserAlreadyExist
	}

	return nil
}

//func validateUsername(username string) error {
//	if username == "" {
//		return fmt.Errorf("%w: firstName is required", ErrUserValidation)
//	}
//	return nil
//}
//
//func validateEmail(email string) error {
//	if email == "" {
//		return fmt.Errorf("%w: email is required", ErrUserValidation)
//	}
//	return nil
//}
