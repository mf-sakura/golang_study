package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

var mu sync.Mutex
var sortedNums []int

func init() {
	sortedNums = []int{}
}
func main() {

	var eg errgroup.Group
	nums := []int{4, 2, 3, 101}
	// nums := []int{4, 2, 3}
	for _, n := range nums {
		eg.Go(SleepLessThan100(n))
	}
	// 全てのGoroutineの実行完了を待ち、1つ目のエラーを起こす
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sortedNums)
}
func SleepLessThan100(n int) func() error {
	return func() error {
		if n >= 100 {
			return errors.New("n must be less than 100")
		}
		time.Sleep(time.Duration(n) * time.Second)
		mu.Lock()
		// deferで必ずUnlockする
		defer mu.Unlock()
		sortedNums = append(sortedNums, n)

		return nil
	}
}
