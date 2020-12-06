package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mf-sakura/golang_study/db/api/infrastructure"
	"github.com/mf-sakura/golang_study/db/api/interfaces/controllers"
)

func main() {
	e := echo.New()
	sqlHandler := infrastructure.NewSQLHandler()
	defer sqlHandler.Conn.Close()

	userController := controllers.NewUserController(sqlHandler.Conn)

	e.GET("/users", userController.Index)
	e.GET("/users/:id", userController.Show)
	e.POST("/users", userController.Create)
	e.PUT("/users/:id", userController.Update)

	// 8080ポートで起動
	e.Logger.Fatal(e.Start(":8080"))
}
