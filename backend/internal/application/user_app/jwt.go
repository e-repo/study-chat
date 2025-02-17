package user_app

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Jwt string

type JwtPayload struct {
	Id    uuid.UUID
	Email string
}

func NewJwt(
	payload *JwtPayload,
	duration time.Duration,
	hmacSecretKey string,
) (Jwt, error) {
	claims := jwt.MapClaims{
		"sub":   payload.Id,
		"email": payload.Email,
		"exp":   time.Now().Add(duration).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(hmacSecretKey))
	if err != nil {
		return "", err
	}
	return Jwt(token), nil
}
