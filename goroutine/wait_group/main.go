package main

import (
	"fmt"
	"sync"
)

func main() {
	// goroutineの同期待ちをsync.WaitGroupで行う
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// waitGroupのcounterを引数だけ増やす
		wg.Add(1)
		// goroutineでは無名関数も使える
		go func(n int) {
			fmt.Println(n)
			// waitGroupのcounterを1減らす
			defer wg.Done()
		}(i)
	}

	// waitGroupのcounterが0になるまでブロックする
	// 正しく使えば全Goroutineの完了を待つ事が出来る。
	// 非同期処理の同期が出来る。
	wg.Wait()
}
