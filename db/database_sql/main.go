package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mf-sakura/golang_study/db/database_sql/infrastructure"
	"github.com/mf-sakura/golang_study/db/database_sql/interfaces/controllers"
)

func main() {
	e := echo.New()
	sqlHandler := infrastructure.NewSqlHandler()
	userController := controllers.NewUserController(sqlHandler.Conn)

	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/users", func(c echo.Context) error { return userController.Create(c) })

	e.Logger.Fatal(e.Start(":8080"))
}
