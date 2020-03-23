# database_sql

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
$ ./main main.go -a index
```

### ユーザー詳細
```sh
$ ./main main.go -a show -i 1
```

### ユーザー作成
```sh
$ ./main main.go -a create -f Alan -l Turing
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
