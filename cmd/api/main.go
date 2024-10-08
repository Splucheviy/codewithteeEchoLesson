package main

import (
	"fmt"
	"net/http"

	"github.com/Splucheviy/codewithteeEchoLesson/cmd/api/handlers"
	"github.com/Splucheviy/codewithteeEchoLesson/common"
	"github.com/Splucheviy/codewithteeEchoLesson/configs"
	"github.com/labstack/echo/v4"
)

type Application struct {
	logger  echo.Logger
	server  *echo.Echo
	handler handlers.Hadler
}

func main() {
	env := configs.NewEnv()
	e := echo.New()
	db, err := common.NewMySQL()
	if err != nil {
		panic(err.Error())
	}
	h := handlers.Hadler{
		DB: db,
	}
	app := Application{
		logger:  e.Logger,
		server:  e,
		handler: h,
	}
	fmt.Println(app)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(env.AppPort))
}
