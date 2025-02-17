package user_ui

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"study-chat/generated/openapi"
	userapp "study-chat/internal/application/user_app"
	userdmn "study-chat/internal/domain/user_dmn"
)

type Request struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=4,max=100"`
}

func (s UserEndpoints) PostAuth(c echo.Context) error {
	var userReq Request
	var errs validator.ValidationErrors

	if err := c.Bind(&userReq); err != nil {
		return err
	}

	validatorSrv := getValidatorService(s.locator)
	if err := validatorSrv.Validate.Struct(userReq); err != nil {
		errors.As(err, &errs)
		return echo.NewHTTPError(http.StatusBadRequest, errs.Translate(validatorSrv.Trans))
	}

	cfg := getConfig(s.locator)
	command := userapp.NewAuthUserCommand(
		userReq.Email,
		userReq.Password,
		cfg.Server.HmacSecretKey,
	)

	ctx := c.Request().Context()
	userRepo := getUserRepo(s.locator)
	jwt, err := userapp.AuthUser(command, ctx, userRepo)
	if err != nil {
		msg := err.Error()
		if errors.Is(err, userdmn.ErrUserNotFound) || errors.Is(err, userdmn.ErrInvalidPassword) {
			return c.JSON(http.StatusBadRequest, openapi.ErrorResponse{Message: &msg})
		}
		return c.JSON(http.StatusInternalServerError, openapi.ErrorResponse{Message: &msg})
	}

	return c.JSON(http.StatusCreated, openapi.AuthUserResponse{Token: string(jwt)})
}
