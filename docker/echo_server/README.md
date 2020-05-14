Dockerを扱う。  
簡単なHTTP ServerをDockerする。  

# 環境変数の読み込み
Dockerを扱う前に環境変数の扱い方に触れる。  
メジャーな方法を3つ記載する。  
## 直接読み込む
標準パッケージの`os.Getenv`でKeyを指定して読み込みが可能  
https://golang.org/pkg/os/#Getenv  
[コード例](https://play.golang.org/p/VmMfpztuO6R)  
環境変数の数が増えてくると、この方法では大変になる。  

## Structタグ
Structにタグを付与して、Unmarshalする事が出来る。  
https://github.com/Netflix/go-env  
https://github.com/kelseyhightower/envconfig  
envconfigをコードでは扱う。  

## dotenv
`.env`を読み込む事が出来る。  
https://github.com/joho/godotenv  


# 検証方法
## ローカル
```
$ make build_local
$ make up_local
$ make hello
```

## docker

```
$ make build_docker
$ make up_docker
# デーモンで起動したい場合
# $ make up_local_deamon
$ make hello
# プロセスの確認
$ docker ps
CONTAINER ID        IMAGE                           COMMAND                  CREATED              STATUS              PORTS                                NAMES
c466f536affe        myapp:latest                    "./myapp"                About a minute ago   Up About a minute   0.0.0.0:8000->8000/tcp               keen_sammet
# コンテナの停止
$ docker stop {container_id}
# コンテナの削除
$ docker rm {container_id}
# イメージの削除
$ docker rmi myapp:latest
```

# 課題
自分でDockerfileを書いてみる。  
CLIよりはサーバーとして起動するプロセスの方が分かりやすいので以下のどれかを推奨  
その際、環境変数経由で設定情報を一つは読み込む様にする。(例) ポートの情報)  
- http_server/echo
- grpc/server
