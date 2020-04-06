package main

import (
	"flag"
	"log"

	"github.com/labstack/echo/v4"

	"github.com/mf-sakura/golang_study/interface/api/infrastructure"
	"github.com/mf-sakura/golang_study/interface/api/interfaces/controllers"
	"github.com/mf-sakura/golang_study/interface/api/interfaces/database"
)

func main() {
	// 起動オプションでmysql or on_memoryを指定する
	var provider string
	flag.StringVar(&provider, "p", "mysql", "provider name")
	flag.Parse()

	e := echo.New()
	sqlHandler := infrastructure.NewSQLHandler()
	userRepository, err := database.NewUserRepository(database.Provider(provider))
	if err != nil {
		log.Fatal(err)
	}
	userController := controllers.NewUserController(sqlHandler.Conn, userRepository)

	e.GET("/users", userController.Index)
	e.POST("/users", userController.Create)

	e.Logger.Fatal(e.Start(":8080"))
}
