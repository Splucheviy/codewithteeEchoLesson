package main

import (
	"github.com/Splucheviy/codewithteeEchoLesson/cmd/api/handlers"
)

func (app *Application) routes(handler handlers.Handler) {
	app.server.GET("/", handler.HealthCheck)
}
