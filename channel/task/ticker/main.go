package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("+++++ Metronome is start +++++")

	bpm := 120

	// metronomeCh := metronomeChan(120)
	doneCh := make(chan interface{})

	waitSpan := time.Duration(60 * 1000 / bpm)
	flag := 0
	ticker := time.NewTicker(waitSpan * time.Millisecond)
	defer ticker.Stop()

	go func() {
		time.Sleep(5 * time.Second)
		doneCh <- "5 second timeout"
	}()

loop:
	for {
		select {
		case <-ticker.C:
			if flag == 0 {
				fmt.Println("tic")
				flag = 1
			} else {
				fmt.Println("tac")
				flag = 0
			}
		case v := <-doneCh:
			fmt.Println(v)
			break loop
		}
	}
	fmt.Println("+++++ Metronome is end +++++")
}
