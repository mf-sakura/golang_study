package main

import (
	"github.com/labstack/echo/v4"

	"github.com/mf-sakura/golang_study/db/api/infrastructure"
	"github.com/mf-sakura/golang_study/db/api/interfaces/controllers"
)

func main() {
	e := echo.New()
	sqlHandler := infrastructure.NewSQLHandler()
	userController := controllers.NewUserController(sqlHandler.Conn)

	e.GET("/users", userController.Index)
	e.POST("/users", userController.Create)

	e.Logger.Fatal(e.Start(":8080"))
}
