# gRPC
googleが開発したRPC(Remote Procedure Call)のフレームワーク  
RPCは日本語では遠隔手続き呼び出しと訳される。  
簡単に書くと、リモートにあるサーバーや関数(手続き)の呼び出しを、ローカルにあるメソッド, 関数呼び出しと同じ様に見せる技術のこと  
通信の詳細を遮蔽してくれるとい利点がある。  
詳しく知りたい人は[参考書](https://www.oreilly.co.jp/books/9784873118703/)を読んでみてください。  

## gRPCの特徴
### HTTP/2
gRPCはHTTP/2を使って通信する。  
gRPCに関連したHTTP/2の良い点を上げます。  
Wantedlyの[記事](https://www.wantedly.com/companies/wantedly/post_articles/220495)が参考になると思います。  
- ストリーム
  - コネクションを貼りっぱなしにする事ができる。
    - gRPCでもストリーミング処理が可能になる。
    - TCPハンドシェイクのオーバーヘッドがなくなる。
- ストリームの多重化
  - ストリームを複数リクエストで使い回す事が出来る
    - 複数スレッドから通信する場合でも、TCPハンドシェイクのオーバーヘッドがなくなる。
- Header圧縮
  - ペイロードが少なくなり通信と処理の速度が上がる

### [Protocol Buffers](https://developers.google.com/protocol-buffers)
Googleが作成した、言語やプラットフォームに中立な、構造化データをシリアライズする為の仕組み  
`.proto`で構造化データを定義し、それをシリアライズする事が出来る。  
独自のエンコーディング方法により、JSONやXMLと比べてシリアライズのバイト数が少なくなる。  
gRPCでは、これをRPCサービスの定義として使用し、シリアライズの方法としても利用している。  

### コード生成
gRPCではRPC部分のコードは`.proto`を元に自動生成して作る。  
具体的な生成方法は各言語のドキュメントを参照([Goの例](https://grpc.io/docs/tutorials/basic/go/))  
RPCが遮蔽する呼び出し部分は実装する必要がない。  
クライアントではリクエスト前後の処理、サーバーではリクエストを受け取った後の処理を実装すれば良い。  
### Stream
gRPCではStreamを使って、1リクエスト内でクライアントとサーバー間でメッセージを複数回やり取りできる。  
#### Server-Side Streaming
ServerからClientへストリームを送る。  
例) サーバーからクライアントへのPush通知, 情報の定期取得(板情報とか)  
#### Client-Side Streaming
ClientからServerへストリームを送る。  
後の章で扱うので、今回は扱わない  
例) ファイルのマルチパートアップロード  
#### Bi-Directional Streaming
双方向でストリームを送り合う。  
後の章で扱うので、今回は扱わない  
例) サーバーを介したチャットルーム  

## gRPCの長所
- gRPCがサポートされている言語であれば、どの言語とでも通信が出来る。
- protobufによりAPI定義が決まるので、クライアントとサーバーの並行開発が非常に行いやすい
  - OpenAPIやSwaggerでは、その仕様通りにAPIが実装されるとは限らない(開発者に依存する)
  - gRPCではprotobufの通りにAPIが自動生成されるので、protoが変更されない限りは定義から外れる事がない。
- HTTP/1.1に比べれば速い
## gRPCの短所
- デバッグが少し面倒
  - grpcurlでそれっぽくデバッグは出来る([ref](https://github.com/fullstorydev/grpcurl))
    - ただし、サーバーがreflectionを有効化しprotobufの情報を公開している必要がある  
  - gRPCのレスポンスをJSONに変換してくれるプロキシ([ref](https://github.com/grpc-ecosystem/grpc-gateway))
    - プロキシがprotobufを把握している必要があるのでproto変更の度にデプロイが必要
- HTTP/2なので負荷分散の技術がHTTP/1.1に比べて難しい
  - [gRPC公式の資料](https://grpc.io/blog/grpc-load-balancing/)
  - 結構長い話なので別の資料で後日書く。
## 課題
### RubyでClientを作ってみる
gRPCは言語に中立なフレームワークなので、rubyのgRPCクライアントからGoのgRPCサーバーにリクエスト出来ます。  
以下の自動生成コマンドでRubyのクライアントは自動生成しています。  

```
grpc_tools_ruby_protoc --ruby_out=ruby_client --grpc_out=ruby_client proto/*.proto
```
生成されたクライアントを使って、以下の実装をしてください。  
- Greeter.SayHelloの結果をプリントする。
- Notifier.PereodicHelloでストリームを受け取り、受け取ったメッセージをプリントする。
