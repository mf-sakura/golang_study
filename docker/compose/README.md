docker-composeを扱う。  
DBに接続するAPI Serverをdocker-composeで構築する。  
マイクロサービスが増えてもこの要領で、docker-composeのサービスを増やしていけば良い  

# サービス
## db
mysqlのサーバー  
mysqlイメージでは、`docker-entrypoint-initdb.d`配下のSQLが起動時に実行される。  
data dirをボリュームマウントする。  
これによりdocker-compose downをしてもデータが残る様になる。  

## web
APIサーバー  
環境変数経由で、DBの接続情報を読み込む。  
環境変数はdocker-compose.ymlで設定する。  
DBのホスト名はサービス名のdbで名前解決が出来る  

# 検証方法
```
$ make up
$ make create_user
$ make show_user
$ make down
```

# 課題
以前のdbの課題で作ったAPIサーバーをdocker-composeで起動してみる。  
docker-composeでDBも別サービスとして立ててください。  
恐らく今回のコードとほぼ同じになると思いますが、docker-compose.ymlを書く練習をしてみてください。  
