Channelを使ったチャットシステム
GoroutineとChannelの例として、gRPCを利用した複数クライアントによるチャットシステムを作った。  
今回作成したコードで扱っている内容を説明します。  

# Goroutineのよくある用途
並行で扱える重い処理を別Goroutineに切り出す。  
重い処理同士を独立に処理出来る様にする。  
今回の例では、以下の3パートをGoroutineで分離している。  
- メッセージ送信リクエストの受理
- 受け付けたメッセージの処理
- 処理したメッセージの配信

# エラーハンドリング
エラーも含んだ型をchannelで処理する。  
エラーを別Channelにすると処理が難しくなる。  
```go
type MyStruct struct{
    Field string
}
// この構造体をChannelで使う
type MyStructOrError{
    MyStruct
    Err error
}
```

# Context
Goroutineを終了させるのに用いられる。  
contextのキャンセルとタイムアウトが可能になる。  

```go
for {
    select {
    case r, ok := <-ch:
        // 処理
    // Contextのキャンセルorタイムアウトが呼ばれるとチャンネルに値が入る
    case <-ctx.Done():
        // ContextがDoneされた理由がErr()に入る。(キャンセル or タイムアウト)
        fmt.Printf("Conetxt Done. %v\n", s.ctx.Err())
        return
    }
}
```

## Cancel
[context.WithCancel](https://golang.org/pkg/context/#WithCancel)  
Contextを引数として渡すと、新しいContextとキャンセル用の関数が返される。  
キャンセル関数を呼ぶと、ctx.Done()に値が入る。  
このContextから新たにContextを作る事が出来るが、親がキャンセルされると子Contextもキャンセルされる。  

## Timeout
[context.WithTimeout](https://golang.org/pkg/context/#WithTimeout)  
ContextとTimeoutを引数として渡すと、新しいContextが返される。  
Timeoutだけ経過すると、ctx.Done()に値が入る。  
このContextから新たにContextを作る事が出来るが、親がTimeoutすると子ContextもTimeoutする。  

# gracefull shutdown
## シグナルハンドリング
よくある例  
```go
signals := make(chan os.Signal, 1)
// 受け取りたいシグナルを受信した際に、第一引数のチャンネルに
signal.Notify(signals, syscall.SIGTERM)
go func() {
    s := <-signals
    // 停止処理
}()
```
## gRPC
[server.GracefullStop](https://godoc.org/google.golang.org/grpc#Server.GracefulStop)を呼ぶ  
リクエストの受付をやめ、全リクエストの完了を待つ。  
こういったAPIサーバーでは、gracefull shutdownの仕組みが用意されている事が多い。  
基本的にはそういった停止関数を呼べば良い。  

## Goroutineの完了待ち
WaitGroupを活用する等。
Goroutineの作成前後でWaitGroupを用いる。  
終了処理の中で、WaitGroup.Wait()を用いて全リクエストが完了する様にする。  
処理が完了しない可能性があるなら、contextによるキャンセルも検討する。  

## チャンネルのClose
deferでGoroutineが完了する際にCloseをする。  
受信専用チャンネルはClose出来ないので注意([ref](https://play.golang.org/p/uZd3_CQXVQ9))  

## CloseしたChannelに対する送受信
panicするので、行わない  

### 送信
CloseしたChannelに送信すると、panicする。  
チャンネルをcloseする前に送信を停止する必要がある。  
今回の例では、StopSendingの様なboolを用意し、それを用いてCloseする前に送信を停止した。  

### 受信
`for range`や第二引数のokを活用する  
