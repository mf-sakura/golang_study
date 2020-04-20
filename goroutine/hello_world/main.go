package main

import (
	"fmt"
	"time"
)

func sayWorld() {
	fmt.Println("world")
}

func sayHelllo() {
	fmt.Println("hello")
}

func main() {

	// sayHellloをGoroutineで呼び出す
	// mainとは並行で実行されるので、world helloの順で表示される。
	// ちなみにmain関数はmain goroutineと呼ばれる
	go sayHelllo()
	sayWorld()

	// main関数が終了するとgoroutineは終了させられる
	// sayHelloを完了させる為に、少しのSleepを挟む
	// goroutineの完了をちゃんと待つ場合は、後述のwaitGroupを使用する。
	time.Sleep(1 * time.Millisecond)
}
