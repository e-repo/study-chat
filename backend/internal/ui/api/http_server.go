package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	userui "study-chat/internal/ui/api/user_ui"

	"study-chat/generated/openapi"
	"study-chat/pkg/echomiddleware"

	"github.com/labstack/echo/v4/middleware"
)

type httpServer struct {
	userui.UserEndpoints
}

func SetupHTTPServer(locator userui.UserLocator) *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	//e.Use(echomiddleware.SlogLoggerMiddleware(slog.Default()))
	e.Use(echomiddleware.PutRequestIDContext)
	e.Use(middleware.Recover())
	//e.Use(sentryecho.New(sentryecho.Options{Repanic: true}))
	//e.Use(echomiddleware.PutSentryContext)

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	server := httpServer{}
	server.UserEndpoints = userui.SetupEndpoints(locator)

	openapi.RegisterHandlers(e, server)

	return e
}
