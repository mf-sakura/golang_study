Table Of Content

# sqlxを使ったCLIツール
sqlxとは。。

## Table Of Content
- [Setup](#Setup)
- [Command](Command)
  - [ユーザー一覧](#ユーザー一覧)
  - [ユーザー詳細](#ユーザー詳細)
  - [ユーザー作成](#ユーザー作成)
- [DB](#DB)
- [課題](#課題)

## Setup
```sh
$ make docker-compose/up
$ make setup
```

## Command
```sh
Usage of main.go
  -a アクション名の指定 (e.g. index, show, create)
  -f ユーザーのfirst nameの指定 (actionがcreateの時のみ有効)
  -i ユーザーのidの指定　(actionがshowの時のみ有効)
  -l ユーザーのlast nameの指定 (actionがcreateの時のみ有効)
```

### ユーザー一覧
```sh
$ ./main -a index
```

### ユーザー詳細
```sh
$ ./main -a show -i 1
```

### ユーザー作成
```sh
$ ./main -a create -f Alan -l Turing
```

### DB
```sh
# mysqlのコンソールに入る
$ make mysql # passwordはそのままpassword

# migrate/up
# railsでいう rails db:migrate
$ make migrate/up

# migrate/down
# railsでいう rails db:rollback
$ make migrate/down
```

# 課題
1. 下記のコマンドを実行した時にユーザーの情報を編集できるようにしてください。
```sh
$ ./main -a update -i [ID] -f [firstName] -l [lastName]
```
2. トランザクションかけてみる。
3. 下記のようなAPIサーバーを実装してください。
```sh
GET /users ユーザー一覧
GET /users/:id ユーザー詳細
POST /users ユーザー作成
PUT /users/:id ユーザー情報編集
```
