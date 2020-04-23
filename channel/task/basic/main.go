package main

import (
	"fmt"
	"time"
)

func main() {
	// stringを送るchannelの初期化
	// channelを使う事でgoroutineとデータをやり取り出来る
	// 今回はHello WorldのGoroutineとMain Goroutineでのデータのやり取り
	ch := make(chan int)

	go func() {
		// channelからの受信で処理が止まる事を確認する為のSleep
		time.Sleep(1 * time.Second)
		// channelへの値の送信(`ch <-`の形式)
		ch <- 2
	}()

	// channelから値の受信(`<- ch`の形式)
	// chanelから値を受け取るまでブロックされる
	// Goroutine内のSleepにより受信が1秒遅れる
	num := <-ch
	fmt.Println("recieved")
	fmt.Println(num)
}
