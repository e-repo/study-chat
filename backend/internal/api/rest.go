package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"study-chat/internal/auth"
	"study-chat/pkg/locator"

	"study-chat/generated/openapi"
	"study-chat/pkg/echomiddleware"

	"github.com/labstack/echo/v4/middleware"
)

type httpServer struct {
	auth.Auth
}

func SetupHTTPServer(locator locator.ServiceLocator) *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(echomiddleware.PutRequestIDContext)
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	//e.Use(echomiddleware.SlogLoggerMiddleware(slog.Default()))
	//e.Use(sentryecho.New(sentryecho.Options{Repanic: true}))
	//e.Use(echomiddleware.PutSentryContext)

	setupPingEndpoint(e)

	server := httpServer{}
	server.Auth = auth.SetupEndpoints(locator)

	openapi.RegisterHandlers(e, server)

	return e
}

func setupPingEndpoint(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})
}
