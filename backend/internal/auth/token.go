package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type jsonWebToken string

type jwtPayload struct {
	Id        uuid.UUID
	FirstName string
	Email     string
}

func newJwt(
	payload *jwtPayload,
	duration time.Duration,
	hmacSecretKey string,
) (jsonWebToken, error) {
	claims := jwt.MapClaims{
		"sub":       payload.Id,
		"firstName": payload.FirstName,
		"email":     payload.Email,
		"exp":       time.Now().Add(duration).Unix(),
	}

	token, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(hmacSecretKey))
	if err != nil {
		return "", err
	}
	return jsonWebToken(token), nil
}
