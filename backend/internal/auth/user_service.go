package auth

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userService struct {
	repo UserRepository
}

func newUserService(userRepo UserRepository) *userService {
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

func (u *userService) signIn(ctx context.Context, signIn *signIn) (jsonWebToken, error) {
	user, err := u.repo.GetUserByEmail(ctx, signIn.email)
	if err != nil {
		return "", ErrUserNotFound
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signIn.password)); err != nil {
		return "", ErrInvalidPassword
	}

	jwtPayload := &jwtPayload{
		Id:        user.Id,
		FirstName: user.FirstName,
		Email:     user.Email,
	}

	jwt, err := newJwt(jwtPayload, 12*time.Hour, signIn.hmacSecretKey)
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

	user := &User{
		Id:        uuid.New(),
		FirstName: signUp.fistName,
		Email:     signUp.email,
		Password:  string(passHash),
	}

	if _, err = u.repo.CreateUser(ctx, user); err != nil {
		return uuid.Nil, err
	}

	return user.Id, nil
}
