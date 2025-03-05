package auth

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"study-chat/generated/openapi"
	"study-chat/generated/protobuf"
	"study-chat/pkg/validator/vlutils"
)

type RequestSignIn struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=4,max=100"`
}

func (a Auth) PostSignIn(c echo.Context) error {
	var request RequestSignIn
	var msg string

	if err := c.Bind(&request); err != nil {
		return err
	}
	if err := a.validator.Validate.Struct(request); err != nil {
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			msg = vlutils.ErrTranslationsToStr(errs.Translate(a.validator.Trans))

			return echo.NewHTTPError(
				http.StatusBadRequest,
				openapi.ErrorResponse{Message: &msg},
			)
		}
		msg = err.Error()
		return echo.NewHTTPError(
			http.StatusUnprocessableEntity,
			openapi.ErrorResponse{Message: &msg},
		)
	}

	signIn := &signIn{
		email:         request.Email,
		password:      request.Password,
		hmacSecretKey: a.HmacSecretKey,
	}

	ctx := c.Request().Context()
	jwt, err := a.service.signIn(ctx, signIn)
	if err != nil {
		msg = err.Error()
		if errors.Is(err, ErrUserNotFound) || errors.Is(err, ErrInvalidPassword) {
			return c.JSON(http.StatusBadRequest, openapi.ErrorResponse{Message: &msg})
		}
		return c.JSON(http.StatusInternalServerError, openapi.ErrorResponse{Message: &msg})
	}

	return c.JSON(http.StatusCreated, openapi.AuthUserResponse{Token: string(jwt)})
}

func (a Auth) SignIn(
	ctx context.Context,
	r *protobuf.AuthUserRequest,
) (*protobuf.AuthUserResponse, error) {
	request := &RequestSignIn{r.Email, r.Password}

	if err := a.validator.Validate.Struct(request); err != nil {
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			translatedErr := errs.Translate(a.validator.Trans)
			return nil, status.Errorf(codes.InvalidArgument, "%v", translatedErr)
		}
		return nil, status.Errorf(codes.Unknown, "%v", err.Error())
	}

	signIn := &signIn{
		email:         request.Email,
		password:      request.Password,
		hmacSecretKey: a.HmacSecretKey,
	}

	jwt, err := a.service.signIn(ctx, signIn)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "%v", err.Error())
	}

	return &protobuf.AuthUserResponse{Token: string(jwt)}, nil
}
