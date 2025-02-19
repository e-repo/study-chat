package auth

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"study-chat/generated/openapi"
)

func (a Auth) PostSignUp(c echo.Context) error {
	ctx := c.Request().Context()
	var userReq openapi.CreateUserRequest
	if err := c.Bind(&userReq); err != nil {
		return err
	}

	signUp := SignUp{
		fistName: userReq.FirstName,
		email:    string(userReq.Email),
		password: userReq.Password,
	}

	userId, err := a.service.SignUp(ctx, &signUp)

	if err != nil {
		msg := err.Error()
		if errors.Is(err, ErrInvalidUser) || errors.Is(err, ErrUserValidation) {
			return c.JSON(http.StatusBadRequest, openapi.ErrorResponse{Message: &msg})
		}
		if errors.Is(err, ErrUserAlreadyExist) {
			return c.JSON(http.StatusConflict, openapi.ErrorResponse{Message: &msg})
		}
		return c.JSON(http.StatusInternalServerError, openapi.ErrorResponse{Message: &msg})
	}

	return c.JSON(http.StatusCreated, openapi.CreateUserResponse{Id: &userId})
}
