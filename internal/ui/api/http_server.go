package api

import (
	"net/http"
	userapp "study-chat/internal/ui/api/user_ui"

	"study-chat/generated/openapi"
	userdmn "study-chat/internal/domain/user_dmn"
	"study-chat/pkg/echomiddleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type httpServer struct {
	userapp.UserServer
}

func SetupHTTPServer(userRepo userdmn.UserRepository) *echo.Echo {
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
	server.UserServer = userapp.SetupServer(userRepo)

	openapi.RegisterHandlers(e, server)

	return e
}
