package user_app

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	userdmn "study-chat/internal/domain/user_dmn"
	"time"
)

const JwtDuration = time.Hour * 12

func AuthUser(
	command *AuthUserCommand,
	ctx context.Context,
	r userdmn.UserRepository,
) (Jwt, error) {
	user, err := r.GetUserByEmail(ctx, command.Email())
	if err != nil {
		return "", userdmn.ErrUserNotFound
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash()), []byte(command.Password())); err != nil {
		return "", userdmn.ErrInvalidPassword
	}

	jwtPayload := &JwtPayload{
		Id:    user.ID(),
		Email: user.Email(),
	}

	jwt, err := NewJwt(jwtPayload, JwtDuration, command.HmacSecretKey())
	if err != nil {
		return "", err
	}
	return jwt, nil
}
