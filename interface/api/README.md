# Interface/api
APIサーバーのデータ保存をInterfaceを用いて遮蔽する。  
以下の3つの実装を作った。
- DBに保存
- オンメモリに保存
- モック(テスト環境)

## テスト
Interfaceを使う事でAPIのロジックに対して単体テストを行う事が可能になる。  
実装例は非常にシンプルだが、難しいロジックを含んでいる本番稼働中のコードでは、interfaceを経由させる事でテストが非常に楽になる。  

## Setup
```sh
$ make docker-compose/up
$ make setup
$ make build
```

## Command
起動時にAPIサーバーが保存するストレージを切り替えられる様にしている。  
```sh
Usage of api_example
    -p 保存先を指定する為のprovider(aws or on_memory)
```

### ユーザー作成
```sh
make create_users
```

### ユーザー表示
```sh
make show_users
```

### テストカバレッジの確認
```sh
make coverage
```
### Localstackを落とす
```sh
docker-compose down
```

## 課題
1. `interfaces/database/user_repository.go`の`onMemoryUserRepository.FindAll`の実装をする。
全ユーザーを返す実装をする。
可能であれば、ユーザーID昇順でSliceを返す。
2. `TestUserController_Create`を修正して、`UserController.Create`でユーザー名が空でエラーになる箇所のCoverageを通す
3. `TestUserController_Index`にテストを追加して、`UserController.Index`のカバレッジを100%にする。
