# Goroutine
Goの軽量スレッド  

Goroutineは裏ではOSのスレッドを利用するが、実装者がOS毎の違いを意識する必要はない。  
シンプルに`go `と呼ぶだけで済む。  
## Goroutineの中身
Goroutineはコルーチンでもある  
コルーチンはプログラミングの構造の種類なので、他の言語に依存するものではない([Wiki](https://ja.wikipedia.org/wiki/%E3%82%B3%E3%83%AB%E3%83%BC%E3%83%81%E3%83%B3))  
GoのランタイムがGoroutineとOSスレッドの紐付け(スケジューリング, マルチプレキシング)を行い、Goroutineを軽量にしてくれている。  

## 並列と並行
Rob Pikeの資料を元に解説する
https://talks.golang.org/2012/waza.slide#6
### 並列
英語ではParallel  
> Programming as the simultaneous execution of (possibly related) computations.

計算の同時実行としてのプログラミング  
ランタイムの性質  
## 並行
英語ではConcurrency  
> Programming as the composition of independently executing processes.

ここで言うプロセスはlinuxのプロセスではなく、一般的なプロセス  


### Goにおける並行性
Goは同時実行である並列性を担保している訳ではなく、独立に実行できるコンポーネントを組み合わせる並行性をサポートする。  
Goroutineが並列(同時)に実行される保証はない  
同時に実行される前提でコードを書くのは非推奨  

# hello_world
一番シンプルな例  
Goroutineにより並行で関数が呼び出される。  
world helloの順で表示される。  

# sleep_sort
Goroutineとmutexを使って、intのSortをする。  
配列の要素の数字だけGoroutine内でSleepをするとSortをする事が出来る  
Goroutineはメモリ空間を全Goroutineで共有する。  
並列性は担保されないので少し微妙だけれど、分かりやすい例なので記載する  

## wrong_sleep_sort
Goroutineで指定する関数のスコープで変数を定義していないので、forが進む毎に値が変わる  
## sync.Mutex
複数Goroutine間で同一のメモリにアクセスする時に使用する。  
変数アクセスをThread Safeにする為の機構  
例えばカウンタのインクリメントはMutexを用いる必要がある。  
※ Mutexを用いずに関数やメソッドを用いる場合はThread Safeか確認する必要がある  

### Mutex.Lock()
mutexをLockする。  
UnLockされるまで、他のGoroutineからのLock()呼び出しの完了はブロックされる。  
このブロックにより、特定リソースへのアクセスが排他制御出来る。  
### Mutex.UnLock()
mutexをUnLockする。  
これにより他のGoroutineによるLockが可能になる。  
deferでUnLockするのがお作法  

# wait_group
sync.WaitGroupを用いると、Goroutineの完了を待つ事が出来る  
複数のGoroutine(mainも含む)を同期させるのに使われる。  
WaitGroupにはカウンタがあり、このカウンタをメソッド経由で操作してGoroutineの実行を待ち合わせする。  
## WaitGroup.Wait()
カウンタが0になるまでブロックする
## WaitGroup.Add(num)
カウンタをnumだけ増やす
## WaitGroup.Done()
カウンタを1減らす


# error_group
準標準パッケージのerrgroup.Groupを使うと、Goroutineの完了同期だけでなく、Goroutine内で起こったエラーをハンドリング出来る。  
SleepSortで100秒以上待つ場合は、長すぎるとしてエラーにする。  
## eg.Go(func() error)
`func() error`を満たす関数を受け取って、Goroutineを作る
## eg.Wait()
全てのeg.Go()の完了を待つ。  
いずれかのGoroutineでerrorが返された場合は、そのerrorを受け取る  
# 課題
Goroutineを使ってみる。  
Sleep SortでWaitGroupを使って確実にGoroutineの実行完了を待てる様にしてください。  
`sleep_sort/main.go`を変更して、この課題を実装してください。  



# Furhter Reading
## Goの並行処理の書籍
有名なGoの並行処理に関してよく纏まった書籍
https://www.oreilly.co.jp/books/9784873118468/
## Concurrency is not Parallelism
並行性と並列性の違いとGoについてのスライド
https://talks.golang.org/2012/waza.slide#1
