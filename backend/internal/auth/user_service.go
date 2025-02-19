package auth

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{repo: userRepo}
}

type SignUp struct {
	fistName string
	email    string
	password string
}

type SignIn struct {
	email         string
	password      string
	hmacSecretKey string
}

func (u *UserService) SignIn(ctx context.Context, command *SignIn) (Jwt, error) {
	user, err := u.repo.GetUserByEmail(ctx, command.email)
	if err != nil {
		return "", ErrUserNotFound
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(command.password)); err != nil {
		return "", ErrInvalidPassword
	}

	jwtPayload := &JwtPayload{
		Id:        user.Id,
		FirstName: user.FirstName,
		Email:     user.Email,
	}

	jwt, err := NewJwt(jwtPayload, 12*time.Hour, command.hmacSecretKey)
	if err != nil {
		return "", err
	}
	return jwt, nil
}

func (u *UserService) SignUp(ctx context.Context, signUp *SignUp) (uuid.UUID, error) {
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
