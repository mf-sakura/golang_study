package main

import (
	"fmt"
	"time"
)

// 課題:
func main() {
	// stringを送るchannelの初期化
	// channelを使う事でgoroutineとデータをやり取り出来る
	// 今回はHello WorldのGoroutineとMain Goroutineでのデータのやり取り
	ch := make(chan string)

	go func() {
		// channelからの受信で処理が止まる事を確認する為のSleep
		time.Sleep(1 * time.Second)
		// channelへの値の送信(`ch <-`の形式)
		ch <- "Hello World"
	}()

	// channelから値の受信(`<- ch`の形式)
	// chanelから値を受け取るまでブロックされる
	// Goroutine内のSleepにより受信が1秒遅れる
	str := <-ch
	fmt.Println("recieved")
	fmt.Println(str)
}
