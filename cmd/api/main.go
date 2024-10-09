package main

import (
	"fmt"

	"github.com/Splucheviy/codewithteeEchoLesson/cmd/api/handlers"
	"github.com/Splucheviy/codewithteeEchoLesson/common"
	"github.com/Splucheviy/codewithteeEchoLesson/configs"
	"github.com/labstack/echo/v4"
)

type Application struct {
	logger  echo.Logger
	server  *echo.Echo
	handler handlers.Handler
}

func main() {
	env := configs.NewEnv()
	e := echo.New()
	db, err := common.NewMySQL()
	if err != nil {
		panic(err.Error())
	}
	h := handlers.Handler{
		DB: db,
	}
	app := Application{
		logger:  e.Logger,
		server:  e,
		handler: h,
	}
	app.routes(h)
	fmt.Println(app)
	e.Logger.Fatal(e.Start(env.AppPort))
}
