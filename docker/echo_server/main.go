package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
)

// Environment is struct to unmarshal environment variables
type Environment struct {
	Port int64 `envconfig:"PORT" default:"8080"`
}

func main() {
	var env Environment
	if err := envconfig.Process("", &env); err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	// GET
	e.GET("/hello", helloHandler)

	// 8080ポートで起動
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", env.Port)))
}

// レスポンスに`Hello World`を書き込むハンドラー
// 引数をこの形にするのはechoの仕様から決まっている
func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World from Go.")
}
