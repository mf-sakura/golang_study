package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Goroutine内で好きなIntをchannelを経由でMain Goroutineに渡し標準出力にPrintしてみる。
func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)

		rand.Seed(time.Now().Unix())
		max := rand.Intn(10)

		for i := 0; i < max; i++ {
			time.Sleep(1 * time.Second)
			ch <- rand.Intn(100)
		}
	}()

	for v := range ch {
		fmt.Printf("revieved: %d\n", v)
	}
}
