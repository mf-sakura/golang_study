package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	"github.com/kelseyhightower/envconfig"
	"github.com/mf-sakura/golang_study/db/api/infrastructure"
	"github.com/mf-sakura/golang_study/db/api/interfaces/controllers"
)

func main() {

	var config APIConfig
	if err := envconfig.Process("", &config); err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	sqlHandler := infrastructure.NewSQLHandler()
	userController := controllers.NewUserController(sqlHandler.Conn)
	e.GET("/users", userController.Index)
	e.POST("/users", userController.Create)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}

type APIConfig struct {
	Port int64 `envconfig:"API_PORT" default:"8080"`
}
