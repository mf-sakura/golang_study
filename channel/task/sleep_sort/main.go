package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{1, 3, 2, 5, 4}

	sortedNumbersChan := sortedNumbersChan(nums)

	for n := range sortedNumbersChan {
		fmt.Println(n)
	}
}

func sortedNumbersChan(nums []int) <-chan int {
	ch := make(chan int)

	for _, num := range nums {
		isLast := num == len(nums)
		
		go func(n int, isLast bool) {
			if isLast {
				defer close(ch)
			}
	
			time.Sleep(time.Duration(n) * time.Second)
			ch <- n
		}(num, isLast)					
	}

	return ch
}
