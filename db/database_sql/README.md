# database_sql
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
$ make build
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
$ ./bin/database_sql -a index
```

### ユーザー作成
```sh
$ ./bin/database_sql -a create -f Alan -l Turing
```

### ユーザー詳細
```sh
$ ./bin/database_sql -a show -i 1
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
