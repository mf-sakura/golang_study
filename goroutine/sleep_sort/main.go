package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

// mutex
var mu sync.Mutex
var sortedNums []int

func main() {
	var eg errgroup.Group
	nums := []int{1, 3, 2, 5, 4}

	// mutexを使わずにChannelを使う方が楽
	// Channelは別の回でやるので、Goroutineだけで頑張る例
	for _, num := range nums {
		// goroutineでは無名関数も使える
		// ここでnumを渡す事で、forが進んでも各Goroutineのスコープ内でnは変化しない。
		eg.Go(SleepLessThan100(num))
	}

	// 全てのGoroutineの実行完了を待ち、1つ目のエラーを起こす
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sortedNums)
}

func init() {
	sortedNums = []int{}
}

// SleepLessThan100 using SleepSort and return error when variable more than 100.
func SleepLessThan100(n int) func() error {
	return func() error {
		if n >= 100 {
			return errors.New("n must be less than 100")
		}
		// n秒スリープする
		time.Sleep(time.Duration(n) * time.Second)
		// mutexのLock
		// 他のGoroutineからのアクセスがブロックされる
		mu.Lock()
		// deferで必ずUnlockする
		defer mu.Unlock()
		sortedNums = append(sortedNums, n)
		return nil
	}
}
