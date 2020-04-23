package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	nums := []int{3, 2, 4, 5, 1}
	var sortedNums []int
	sortedCh := sleepSort(nums)
	for n := range sortedCh {
		fmt.Println(n)
		sortedNums = append(sortedNums, n)
	}
	fmt.Println(sortedNums)
}

// `<-ch int`は受信専用チャンネル
func sleepSort(nums []int) <-chan int {
	// channelの初期化
	ch := make(chan int)
	go func(nums []int) {
		// deferでchannelをCloseする
		defer close(ch)
		var wg sync.WaitGroup
		// sleep sort
		for _, num := range nums {
			wg.Add(1)
			go func(n int) {
				defer wg.Done()
				time.Sleep(time.Duration(n) * time.Second)
				ch <- n
			}(num)
		}
		wg.Wait()
	}(nums)
	return ch
}
