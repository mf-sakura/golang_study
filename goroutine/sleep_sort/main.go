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
		wg.Add(1)

		go func(n int) {
			time.Sleep(time.Duration(n) * time.Second)
			mu.Lock()

			defer func() {
				mu.Unlock()
				wg.Done()
			}()
			sortedNums = append(sortedNums, n)
		}(num) // (num)が無名関数の引数
	}

	wg.Wait()
	fmt.Println(sortedNums)

}
