package api

import (
	"github.com/Splucheviy/codewithteeEchoLesson/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

type Application struct {
	Logger  echo.Logger
	Server  *echo.Echo
	Handler handlers.Handler
}

func (app *Application) Routes(handler handlers.Handler) {
	app.Server.GET("/", handler.HealthCheck)
	app.Server.POST("/register", handler.RegisterHandler)
}
