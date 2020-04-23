package main

import (
	"fmt"
	"time"
)

func main() {

	worldCh := helloWorldViaChan()
	doneCh := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		// timeoutになってselectのループが終了する。
		// 実際のケースとしてありそうな例
		doneCh <- "5 second timeout."
	}()

	// GOではLoopに対してLabelを付ける事が出来、break,
loop:
	for {
		// selectのcase評価はdefault以外をランダムな順番で評価する
		// どのchannelも受信できない場合はdefaultが実行される
		select {
		// channelからの受信の第2引数でChannelがOpenしているかどうかの情報が取得できる。
		// Openしている場合はTrue
		case v, ok := <-worldCh:
			if !ok {
				fmt.Println("world chan is closed")
				break loop
			}
			fmt.Println(v)
		//　selectの無限ループを外部から抜けさせる事を可能にする為に、doneChを用意している。
		case v := <-doneCh:
			fmt.Printf("done. Reason is %s\n", v)
			break loop
		// どのChannelも準備出来ていない場合はdefaultを通る
		default:
			fmt.Println("default case.Sleep 1s")
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("select loop exit")
}

// `<-chan`は受信専用チャンネル
func helloWorldViaChan() <-chan string {
	ch := make(chan string)
	go func() {
		// // ChannelのCloseをし忘れるケース
		// // このケースでdoneChが無ければ、mainの無限ループが終了しない
		// defer close(ch)
		ch <- "hello"
		time.Sleep(2 * time.Second)
		ch <- "world"
	}()
	return ch
}
