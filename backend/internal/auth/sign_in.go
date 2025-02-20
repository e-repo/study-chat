package auth

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"study-chat/generated/openapi"
)

type RequestSignIn struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=4,max=100"`
}

func (a Auth) PostSignIn(c echo.Context) error {
	var request RequestSignIn
	var errs validator.ValidationErrors

	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := a.validator.Validate.Struct(request); err != nil {
		errors.As(err, &errs)
		return echo.NewHTTPError(http.StatusBadRequest, errs.Translate(a.validator.Trans))
	}

	signIn := &signIn{
		email:         request.Email,
		password:      request.Password,
		hmacSecretKey: a.hmacSecretKey,
	}

	ctx := c.Request().Context()
	jwt, err := a.service.signIn(ctx, signIn)
	if err != nil {
		msg := err.Error()
		if errors.Is(err, ErrUserNotFound) || errors.Is(err, ErrInvalidPassword) {
			return c.JSON(http.StatusBadRequest, openapi.ErrorResponse{Message: &msg})
		}
		return c.JSON(http.StatusInternalServerError, openapi.ErrorResponse{Message: &msg})
	}

	return c.JSON(http.StatusCreated, openapi.AuthUserResponse{Token: string(jwt)})
}
