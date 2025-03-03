package chat

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"study-chat/generated/openapi"
	"study-chat/pkg/validator/vlutils"
)

type MessageData struct {
	UserId   uuid.UUID `json:"userId" validate:"required,uuid"`
	UserName string    `json:"userName" validate:"required,min=4,max=100"`
	Message  string    `json:"message" validate:"required,min=4,max=500"`
}

type RequestAddMessage struct {
	Channel string      `json:"channel" validate:"required,min=2,max=100"`
	Data    MessageData `json:"data"`
}

func (ch *Chat) PostMessage(c echo.Context) error {
	var request RequestAddMessage

	if err := c.Bind(&request); err != nil {
		msg := err.Error()

		return c.JSON(
			http.StatusUnprocessableEntity,
			openapi.ErrorResponse{Message: &msg},
		)
	}
	if err := ch.validator.Validate.Struct(request); err != nil {
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			msg := vlutils.ErrTranslationsToStr(errs.Translate(ch.validator.Trans))

			return c.JSON(
				http.StatusBadRequest,
				openapi.ErrorResponse{Message: &msg},
			)
		}
		msg := err.Error()
		return c.JSON(
			http.StatusUnprocessableEntity,
			openapi.ErrorResponse{Message: &msg},
		)
	}

	ctx := c.Request().Context()
	rawMessage, _ := json.Marshal(request)
	_, err := ch.service.addMessage(
		ctx,
		&newMessage{Payload: rawMessage},
	)
	if err != nil {
		if errors.Is(err, ErrSaveMessage) {
			msg := err.Error()
			return c.JSON(
				http.StatusUnprocessableEntity,
				openapi.ErrorResponse{Message: &msg},
			)
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, "Сообшение успешно добавленно")
}
