package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)

	// 終了処理の為にdoneを用意する
	done := make(chan interface{})

	go func() {
		time.Sleep(2 * time.Second)
		// ticker.Stop()すると、チャンネルからの送信は止まるがCloseはされない
		ticker.Stop()
		fmt.Println("ticker stop")
	}()
	go func() {
		time.Sleep(5 * time.Second)
		// 終了処理
		done <- "done"
	}()
loop:
	for {
		select {
		case t, ok := <-ticker.C:
			// Stop()でチャンネルがCloseされない事の確認
			if !ok {
				fmt.Println("ticker.C is closed")
			}
			fmt.Printf("Periodic: current time is %v\n", t)
		case <-done:
			fmt.Printf("Done: current time is %v\n", time.Now())
			break loop
		}
	}
	fmt.Println("Successfully Shutdown")
}
