package auth

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"study-chat/generated/openapi"
)

type RequestSignUp struct {
	FirstName string `json:"first_name" validate:"required,min=4,max=100"`
	Email     string `json:"email" validate:"required,email,max=100"`
	Password  string `json:"password" validate:"required,min=4,max=100"`
}

func (a Auth) PostSignUp(c echo.Context) error {
	var request RequestSignUp

	ctx := c.Request().Context()
	if err := c.Bind(&request); err != nil {
		return err
	}
	if err := a.validator.Validate.Struct(request); err != nil {
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			return echo.NewHTTPError(http.StatusBadRequest, errs.Translate(a.validator.Trans))
		}
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	signUp := signUp{
		fistName: request.FirstName,
		email:    request.Email,
		password: request.Password,
	}

	userId, err := a.service.signUp(ctx, &signUp)

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
