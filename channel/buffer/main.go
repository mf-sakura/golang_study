package main

import (
	"fmt"
	"time"
)

// 課題:
func main() {
	// buffer付きchannelの初期化
	// 第二引数がbufferになる
	// channelはbufferだけ値をストックする事が可能
	bufferedCh := make(chan string, 2)

	go func() {

		bufferedCh <- "a"
		fmt.Println("Send a")
		bufferedCh <- "b"
		fmt.Println("Send b")
		// bufferが2しかないのでここでブロックされる
		bufferedCh <- "c"
		fmt.Println("Send c")
	}()

	// Bufferが埋まるのを待つ為に1秒Sleep
	time.Sleep(1 * time.Second)
	str1 := <-bufferedCh
	fmt.Printf("recieve %s\n", str1)
	str2 := <-bufferedCh
	fmt.Printf("recieve %s\n", str2)
	str3 := <-bufferedCh
	fmt.Printf("recieve %s\n", str3)

	// bufferを指定しない場合は0になる
	ch := make(chan string)
	go func() {

		// bufferが0の場合は、送信と受信が揃うまでどちらもブロックされる
		ch <- "d"
		// 受信が準備出来る約1秒間、送信がブロックされる事が標準出力から分かる
		fmt.Println("Send d")
	}()

	time.Sleep(1 * time.Second)
	str := <-ch
	fmt.Printf("recieve %s\n", str)
}
