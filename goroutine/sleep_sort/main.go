package main

import (
	"fmt"
	"sync"
	"time"
)

// mutex
var mu sync.Mutex

func main() {
	nums := []int{1, 3, 2, 5, 4}

	// mutexを使わずにChannelを使う方が楽
	// Channelは別の回でやるので、Goroutineだけで頑張る例
	sortedNums := make([]int, 0, len(nums))
	var wg sync.WaitGroup

	for _, num := range nums {
		// goroutineでは無名関数も使える
		// ここでnumを渡す事で、forが進んでも各Goroutineのスコープ内でnは変化しない。
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			// n秒スリープする
			time.Sleep(time.Duration(n) * time.Second)
			// mutexのLock
			// 他のGoroutineからのアクセスがブロックされる
			mu.Lock()
			// deferで必ずUnlockする
			defer mu.Unlock()
			sortedNums = append(sortedNums, n)
		}(num) // (num)が無名関数の引数
	}

	wg.Wait()
	fmt.Println(sortedNums)

}
