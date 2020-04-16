# Interface/example
Interfaceは実装すべきメソッドを定義した型  
型なので関数の引数や戻り値として使用できる。  
Interfaceで定義されたメソッドを実装した型は、Interfaceとして渡す事が可能になる。  

## interface実装の例
`Error() string`というメソッドを実装すればerrorとして扱える。  
自分で定義したerrorを使う事が可能になる。  
```
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}

type MyError struct{}
func (e MyError) Error() string {
	return "My Error Occurred"
}
```

## Interfaceの良い点
Interfaceにより実装を遮蔽できるので、コードが疎結合になる。  
JavaやC#といった静的言語でよくあるInterfaceに近いがImplement宣言が不要。  
コンパイル時に型チェックをしつつ、動的言語の様にduck typingが出来る。  

## テーマ
AWS S3へファイルを読み書きするCLIを考え  
Interfaceを用いると、S3以外にも書き込みが出来る様に出来、AWS環境以外でもCLIを利用できる。  
今回は以下の3種類の実装を作った。  
- S3に保存(AWS環境)
- ローカルファイルシステムに保存(ローカル環境)
- モック(テスト環境)
S3へのアクセスは、S3のモックのlocalstackを使用する。  
テストのカバレッジから、エラーケースもモックにより簡単に通せている事が確認できる。  

## Setup
```sh
# localstackの起動
$ make setup
# Bucketの作成
$ make create_bucket
# Bucketが作成された事の確認
$ list_buckets:
$ make build
```

## Command
```sh
Usage of interface_example
    -f アップロードするファイルのパス
    -s 保存時のファイル名
    -a アクション名(write or read)
    -p 保存先を指定する為のprovider(aws or local)
```

### S3に保存
```sh
make s3_write
```

### S3から読み取り
```sh
make s3_read
```

### 実際にS3に保存されているかの確認
```sh
# AWS CLIが必要
make list_objects
```

### ローカルに保存
```sh
# output配下に保存される
make local_write
```

### ローカルから読み取り
```sh
make local_read
```

### テストカバレッジの確認
```sh
make coverage
```

### Localstackを落とす
```sh
docker-compose down
```
