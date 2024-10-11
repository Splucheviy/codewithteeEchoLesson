package main

import (
	"fmt"

	"github.com/Splucheviy/codewithteeEchoLesson/cmd/api"
	"github.com/Splucheviy/codewithteeEchoLesson/cmd/api/handlers"
	"github.com/Splucheviy/codewithteeEchoLesson/cmd/api/middlewares"
	"github.com/Splucheviy/codewithteeEchoLesson/common"
	"github.com/Splucheviy/codewithteeEchoLesson/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


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
	app := api.Application{
		Logger:  e.Logger,
		Server:  e,
		Handler: h,
	}
	e.Use(middlewares.CustomMiddleware, middleware.Logger())
	app.Routes(h)
	fmt.Println(app)
	e.Logger.Fatal(e.Start(env.AppPort))
}

