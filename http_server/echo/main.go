// メイン関数(実行時に呼ばれる関数)を含むpackageはmainにする
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"math/rand"

	"github.com/labstack/echo/v4"
)

var (
	counter = 0
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func main() {
	e := echo.New()

	// GET
	e.GET("/hello", helloHandler)
	// GET 200以外のStatus
	e.GET("/401", unAuthorizedHandler)
	// GET Headerの読み込み
	e.GET("/square", squareHandler)
	// POST Bodyの読み込み
	e.POST("/incr", incrementHandler)
	// ランダムなポーカーのハンドを返す
	e.GET("/random_hand", randomHandHandler)

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
	if num*num > 100 {
		return echo.NewHTTPError(http.StatusBadRequest, "square is larger than 100")
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
	response := &incrResponse{
		Counter: counter,
	}
	return c.JSON(http.StatusOK, response)
}

type incrRequest struct {
	// jsonタグをつける事でjsonのunmarshalが出来る
	// jsonパッケージに渡すので、Publicである必要がある
	Num int `json:"num"`
}

type incrResponse struct {
	Counter int `json:counter`
}

// ランダムなポーカーのハンドを返す
func randomHandHandler(c echo.Context) error {
	// cards を抽出するロジック部分を `shuffledDeck` みたいな関数に切り出したい
	suits := []string{"s", "h", "d", "c"}
	numbers := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	var cards []string
	for _, s := range suits {
		for _, n := range numbers {
			cards = append(cards, n+s)
		}
	}

	card1, restCards := pickup(cards)
	card2, _ := pickup(restCards)
	hand := card1 + card2

	return c.String(http.StatusOK, fmt.Sprintf("your hand is %s. raise or fold?", hand))
}

func pickup(slice []string) (string, []string) {
	i := rand.Intn(len(slice))
	var restSlice []string
	// 0から(i-1)番目の要素までを追加する
	restSlice = append(restSlice, slice[:i]...)
	// 一番後ろの要素がピックアップされていなければ、i+1番目から最後の要素までを追加する
	if len(restSlice) != len(slice)-1 {
		restSlice = append(restSlice, slice[i+1:]...)
	}

	return slice[i], restSlice
}
