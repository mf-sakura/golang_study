package main

import (
	"fmt"
	"time"
)

// channel/basicのInt版
// Goroutine内で好きなIntをchannelを経由でMain Goroutineに渡し標準出力にPrintしてみる。
func main() {
	ch := make(chan int, 0)
	go func(n int) {
		// いる？
		time.Sleep(1 * time.Second)
		// channelへの値の送信(`ch <-`の形式)
		ch <- n
	}(13)

	recieved := <-ch
	fmt.Println("recieved")
	fmt.Println(recieved)
}
