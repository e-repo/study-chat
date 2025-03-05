package api

import (
	echojwt "github.com/labstack/echo-jwt/v4"
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

var jwtSkippedRoute = []string{
	"/ping",
	"/sign-in",
	"/sign-up",
}

type HttpServer struct {
	*auth.Auth
	*chat.Chat
}

func SetupRESTServer(locator locator.ServiceLocator) *echo.Echo {
	e := createDefaultEcho()

	server := HttpServer{}
	server.Auth = auth.CreateAuth(locator)
	server.Chat = chat.CreateChat(locator)

	e.Use(createJwtMiddleware(server.HmacSecretKey))

	openapi.RegisterHandlers(e, server)

	return e
}

func SetupRESTTestServer(t *testing.T) (*echo.Echo, HttpServer) {
	e := createDefaultEcho()

	server := HttpServer{}
	server.Auth = auth.CreateTestAuth(t)

	openapi.RegisterHandlers(e, server)

	return e, server
}

func createDefaultEcho() *echo.Echo {
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

func createJwtMiddleware(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secretKey),
		Skipper: func(c echo.Context) bool {
			for _, path := range jwtSkippedRoute {
				if path == c.Path() {
					return true
				}
			}

			return false
		},
	})
}

func setupPingEndpoint(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})
}
