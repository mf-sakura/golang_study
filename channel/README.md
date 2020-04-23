# basic
基本的なChannelの扱い方  
Goroutine間で値を受け渡しする為にChannelを用いる。  
`make(chan {型})`の形で初期化をする。  
`ch <- "Hello World"`で値の送信  
`value := <- ch`で受信が出来る。  

# buffer
Channelにはbufferという概念がある。  
bufferによりchannelに値を蓄積する事が可能になる。  
| 種類  | bufferが満杯 | bufferが空 | それ以外 |
| :---: | :----------: | :--------: | :------: |
| 送信  |   ブロック   |    可能    |   可能   |
| 受信  |     可能     |  ブロック  |   可能   |

bufferを指定していないChannelはbufferが0となる。  
送信と受信の両方が揃うまで、処理が進まない状態になる。  

# for 
`for v := range := ch`の形式でchannelから受信した値に対してループが回せる。  
channelから値を受信するまでは、各ループがブロックされる。  
channelがcloseされるとforを抜ける。  
## close
チャンネルは使い終わったらCloseする必要がある。  
Goroutineによる並行処理の完了はChannelのCloseによって判断する事が多い。  
メモリリーク防止の目的もある。  
deferで行い最終的にChannelが閉じられる事を担保するのがお作法  
Closeしたチャンネルへの送信はPanicする。  
Closeしたチャンネルからの受信はデフォルト値が入る。  
## 単一方向のchannel
`<-chan int`という型は受信専用channel  
`chan<- int`という型は送信専用channel  
chanの左右どちらかにあるかで受信/送信を判断する。  
単一方向のchannelに逆方向の操作をしようとするとコンパイルエラーになる。  
# select
Channelに対するSwitchのようなもの。  
selectのcase評価はdefault以外をランダムな順番で評価する。  
どのchannelも受信できない場合はdefaultが実行される。  
## ChannelのClose情報
`v, ok := <-ch`の様な形式でChannelがCloseしているかどうかの情報が取得できる。  
channelからの受信の第2引数が相当する。  
Closeしている場合はokがfalseになる。  
Selectでたまに用いる。  

# 課題
`./task`配下にフォルダと`main.go`だけ作っています。
## goroutineとChannelを使ってみる
`./task/basic/main.go`にコードを記載してください。  
`channel/basic`でやった事をIntでやってみる課題  
Goroutine内から好きなIntをchannelを経由でMain Goroutineに渡し標準出力にPrintしてみてください。  

## Sleep SortをChannelで実装してみる
`./task/sleep_sort/main.go`にコードを記載してください。  
以下を使ってGoroutineで扱ったSleep Sortを実装してみてください。  
- 受信専用channel
- for range
- close

## time.Tickerを用いた定期実行
`./task/ticker/main.go`にコードを記載してください。  
time.Tickerを用いて定期的にChannelから値を受信する事が可能です。  
Goにおける定期実行は、time.Tickerをよく用います。  
time.Tickerのレファレンス  
https://golang.org/pkg/time/#NewTicker
以下の要件を満たす様にコードを書いてみてください。  

- Selectを使う
- 何かしらの処理を定期的に行う
- 定期的な処理を終了させてmain関数を終了させる。(終了条件は好き好き)
  - ただしticker.Stop()を呼んでもチャンネルはCloseされないので、終了処理は工夫する必要があります。
