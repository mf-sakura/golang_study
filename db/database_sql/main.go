// メイン関数(実行時に呼ばれる関数)を含むpackageはmainにする
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mf-sakura/golang_study/db/database_sql/infrastructure"
	"github.com/mf-sakura/golang_study/db/database_sql/interfaces/controllers"
)

func main() {
	e := echo.New()
	sqlHandler := infrastructure.NewSqlHandler()
	userController := controllers.NewUserController(sqlHandler.Conn)

	e.GET("/user", helloHandler)
	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/users", func(c echo.Context) error { return userController.Create(c) })

	e.Logger.Fatal(e.Start(":8080"))
}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World from Go.")
}
