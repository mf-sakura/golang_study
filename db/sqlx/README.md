Table Of Content

# database_sqlを使ったCLIツール
database/sqlは

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
$ go run main.go -a index
```

### ユーザー詳細
```sh
$ go run main.go -a show -i 1
```

### ユーザー作成
```sh
$ go run main.go -a create -f Alan -l Turing
```

### DB
```sh
# mysqlのコンソールに入る
$ make mysql # passwordはそのままpassword

# migrate/up
# railsでいう rails db:migrate
$ make migrate/up

# migrate/down
# railsでいう rails db:rollback step=1
$ make migrate/down
```

# 課題
1. ユーザーの情報を編集(PUT)するアクションを追加し
2. トランザクションかけてみる。
