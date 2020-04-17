package main

import (
	"fmt"
	"sync"
)

func sayworld(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("world")
	}
}

func sayhelllo() {
	for i := 0; i < 5; i++ {
		fmt.Println("hello")
	}
}

func main() {
	var wg sync.WaitGroup
	// wg.Add(n)でwgにn個の並列処理があるということを伝える
	wg.Add(1)
	// 実行したいgoroutineにwgのアドレスを渡す
	go sayworld(&wg)
	sayhelllo()
	// wg.Doneまで待つ
	wg.Wait()
}
