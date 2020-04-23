package main

import (
	"fmt"
)

func main() {

	greetCh := helloWorldViaChan()
	// channelがCloseされるとforを抜ける
	for g := range greetCh {
		fmt.Println(g)
	}
	// // 受信専用チャンネルに送信しようとするとコンパイルが通らない。
	// greetCh <- "abc"
}

// `<-chan`は受信専用チャンネル
func helloWorldViaChan() <-chan string {
	// channelの初期化
	ch := make(chan string)
	go func() {
		// deferでchannelをCloseする
		defer close(ch)
		ch <- "hello"
		ch <- "world"
	}()
	// 1. channelを初期化
	// 2. 処理はgoroutine内で並行に行う
	// 3. 読み取り専用Channelを返す
	// というよく見られるケース
	return ch
}
