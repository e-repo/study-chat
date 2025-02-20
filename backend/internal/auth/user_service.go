package auth

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userService struct {
	repo userRepository
}

func newUserService(userRepo userRepository) *userService {
	return &userService{repo: userRepo}
}

type signUp struct {
	fistName string
	email    string
	password string
}

type signIn struct {
	email         string
	password      string
	hmacSecretKey string
}

func (u *userService) signIn(ctx context.Context, command *signIn) (jsonWebToken, error) {
	user, err := u.repo.getUserByEmail(ctx, command.email)
	if err != nil {
		return "", ErrUserNotFound
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(command.password)); err != nil {
		return "", ErrInvalidPassword
	}

	jwtPayload := &jwtPayload{
		Id:        user.Id,
		FirstName: user.FirstName,
		Email:     user.Email,
	}

	jwt, err := newJwt(jwtPayload, 12*time.Hour, command.hmacSecretKey)
	if err != nil {
		return "", err
	}
	return jwt, nil
}

func (u *userService) signUp(ctx context.Context, signUp *signUp) (uuid.UUID, error) {
	password := []byte(signUp.password)

	passHash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, err
	}

	user := &user{
		Id:        uuid.New(),
		FirstName: signUp.fistName,
		Email:     signUp.email,
		Password:  string(passHash),
	}

	if _, err = u.repo.createUser(ctx, user); err != nil {
		return uuid.Nil, err
	}

	return user.Id, nil
}
