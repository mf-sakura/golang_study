// メイン関数(実行時に呼ばれる関数)を含むpackageはmainにする
package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	counter = 0
)

func main() {
	e := echo.New()
	// アクセスログを出す
	e.Use(middleware.Logger())

	// GET
	e.GET("/hello", helloHandler)
	// GET 200以外のStatus
	e.GET("/401", unAuthorizedHandler)
	// GET Headerの読み込み
	e.GET("/square", squareHandler)
	// POST Bodyの読み込み
	e.POST("/incr", incrementHandler)
	// GET Counterの読み込み
	e.GET("/counter", counterHandler)
	// 8080ポートで起動
	e.Logger.Fatal(e.Start(":8080"))
}

// レスポンスに`Hello World`を書き込むハンドラー
// 引数をこの形にするのはechoの仕様から決まっている
func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World from Go.")
}

// 200以外のHTTP Statusを返すハンドラー
func unAuthorizedHandler(c echo.Context) error {
	return echo.NewHTTPError(http.StatusUnauthorized, "UnAuthorized")
}

// Headerから数字を取得して、その二乗を返すハンドラー
func squareHandler(c echo.Context) error {
	// Headerの読み込み
	numStr := c.Request().Header.Get("num")
	// String -> Intの変換
	num, err := strconv.Atoi(numStr)
	if err != nil {
		// 他のエラーの可能性もあるがサンプルとして纏める
		return echo.NewHTTPError(http.StatusBadRequest, "num is not integer")
	}
	if num >= 100 {
		return echo.NewHTTPError(http.StatusBadRequest, "num must be less than 100")
	}
	// fmt.Sprintfでフォーマットに沿った文字列を生成できる。
	return c.String(http.StatusOK, fmt.Sprintf("Square of %d is equal to %d", num, num*num))
}

// Bodyから数字を取得してその数字だけCounterをIncrementするハンドラー
// DBがまだないので簡易的なもの
func incrementHandler(c echo.Context) error {
	incrRequest := incrRequest{}
	if err := c.Bind(&incrRequest); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	counter += incrRequest.Num

	return c.JSON(http.StatusOK, incrResponse{Counter: counter})
	// 以下の様にmap形式でも返却可能
	// return c.JSON(http.StatusOK, map[string]interface{}{"counter": counter})
}

type incrRequest struct {
	// jsonタグをつける事でjsonのunmarshalが出来る
	// jsonパッケージに渡すので、Publicである必要がある
	Num int `json:"num"`
}

type incrResponse struct {
	// jsonタグをつける事でjsonのunmarshalが出来る
	// jsonパッケージに渡すので、Publicである必要がある
	Counter int `json:"counter"`
}

// counterの値を返す
func counterHandler(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("Value of Counter is %d \n", counter))
}
