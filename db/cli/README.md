# sqlxを使ったCLIツール
https://mf.esa.io/posts/129444

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
$ ./bin/cli -a index
```

### ユーザー詳細
```sh
$ ./bin/cli -a show -i 1
```

### ユーザー作成
```sh
$ ./bin/cli -a create -f Alan -l Turing
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
$ ./bin/cli -a update -i [ID] -f [firstName] -l [lastName]
```
2. トランザクションかけてみる。（ここどういう表現にするかちょっと迷い中）
3. 下記のようなAPIサーバーを実装してください。

## GET /users ユーザー一覧
res body
```json
[
  {
    id: 1,
    first_name: "first_name",
    last_name: "last_name"
  },
  {
    id: 2,
    first_name: "first_name_2",
    last_name: "last_name_2"
  }
]
```

## GET /users/:id ユーザー詳細
res body
```json
{
  id: 1,
  first_name: "first_name",
  last_name: "last_name"
}
```

## POST /users ユーザー作成
req body
```json
{
  first_name: "first_name",
  last_name: "last_name"
}
```

## PUT /users/:id ユーザー情報編集
req body
```json
{
  id: 1,
  first_name: "first_name",
  last_name: "last_name"
}
```
