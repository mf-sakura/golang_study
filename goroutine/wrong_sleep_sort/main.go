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

	for _, num := range nums {
		// goroutineでは無名関数も使える
		// numをgoroutineに渡さない悪い例
		// numの値がforが進むにつれて変わるので、上手くソートが出来ない。
		go func() {
			// n秒スリープする
			time.Sleep(time.Duration(num) * time.Second)
			// mutexのLock
			// 他のGoroutineからのアクセスがブロックされる
			mu.Lock()
			// deferで必ずUnlockする
			defer mu.Unlock()
			sortedNums = append(sortedNums, num)
		}()
	}

	time.Sleep(6 * time.Second)
	fmt.Println(sortedNums)

}
