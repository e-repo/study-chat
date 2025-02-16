package user_ui

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"study-chat/generated/openapi"
	userapp "study-chat/internal/application/user_app"
	userdmn "study-chat/internal/domain/user_dmn"
)

func (s UserServer) PostUsers(c echo.Context) error {
	ctx := c.Request().Context()
	var userReq openapi.CreateUserRequest
	if err := c.Bind(&userReq); err != nil {
		return err
	}

	command, err := userapp.NewCreateUserCommand(
		userReq.FirstName,
		string(userReq.Email),
		userReq.Password,
	)
	if err != nil {
		return err
	}

	user, err := userapp.CreateUser(command, ctx, s.repo)

	if err != nil {
		msg := err.Error()
		if errors.Is(err, userdmn.ErrInvalidUser) || errors.Is(err, userdmn.ErrUserValidation) {
			return c.JSON(http.StatusBadRequest, openapi.ErrorResponse{Message: &msg})
		}
		if errors.Is(err, userdmn.ErrUserAlreadyExist) {
			return c.JSON(http.StatusConflict, openapi.ErrorResponse{Message: &msg})
		}
		return c.JSON(http.StatusInternalServerError, openapi.ErrorResponse{Message: &msg})
	}

	id := user.Id
	return c.JSON(http.StatusCreated, openapi.CreateUserResponse{Id: &id})
}
