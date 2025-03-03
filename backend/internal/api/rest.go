package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"study-chat/internal/auth"
	"study-chat/internal/chat"
	"study-chat/pkg/locator"
	"testing"

	"study-chat/generated/openapi"
	"study-chat/pkg/echomiddleware"

	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
	*auth.Auth
	*chat.Chat
}

func SetupRESTServer(locator locator.ServiceLocator) *echo.Echo {
	e := createEcho()

	server := HttpServer{}
	server.Auth = auth.CreateAuth(locator)
	server.Chat = chat.CreateChat(locator)

	openapi.RegisterHandlers(e, server)

	return e
}

func SetupRESTTestServer(t *testing.T) (*echo.Echo, HttpServer) {
	e := createEcho()

	server := HttpServer{}
	server.Auth = auth.CreateTestAuth(t)

	openapi.RegisterHandlers(e, server)

	return e, server
}

func createEcho() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(echomiddleware.PutRequestIDContext)
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	//e.Use(echomiddleware.SlogLoggerMiddleware(slog.Default()))
	//e.Use(sentryecho.New(sentryecho.Options{Repanic: true}))
	//e.Use(echomiddleware.PutSentryContext)

	setupPingEndpoint(e)

	return e
}

func setupPingEndpoint(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})
}
