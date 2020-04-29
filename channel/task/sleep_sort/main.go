package main

import (
	"fmt"
	"sync"
	"time"
)

// 受信専用channel, for range, close←を使ってSleepSortを実装する
func main() {
	nums := []int{1, 3, 2, 5, 4}
	sortedNums := []int{}

	for v := range receivedChan(nums) {
		fmt.Println(v)
		sortedNums = append(sortedNums, v)
	}
	fmt.Println(sortedNums)
}

func receivedChan(nums []int) <-chan int {
	count := len(nums)
	ch := make(chan int, count)
	go func(nums []int) {
		defer close(ch)
		sleepSortCh(ch, nums)
	}(nums)

	return ch
}

func sleepSortCh(ch chan int, nums []int) chan int {
	var wg sync.WaitGroup
	for _, num := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			time.Sleep(time.Duration(num) * time.Second)
			ch <- num
		}(num)
	}
	wg.Wait()

	return ch
}
