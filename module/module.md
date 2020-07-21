Module周りの質問が思ったより多かったので、Moduleについて纏めてみる。  

# Moduleとは
Go言語公式の外部ライブラリを管理する為の機構([Wiki](https://github.com/golang/go/wiki/Modules))  
昔はいくつか3rd Party製のバージョン管理ツールがあったが、Moduleが出た事によりModuleがデファクトスタンダードになった。  
> In Go 1.14, module support is considered ready for production use  

Go1.14からModuleがProduction Readyになった。  

# Moduleの使い方
## go mod init
一番初めに実行するコマンド  
これにより、以下の様な`go.mod`というバージョン管理の為のファイルが作成される。  
```
module github.com/mf-sakura/golang_study/module

go 1.14
```
GOPATH配下の場合は自動でModule名が推測されるが、GOPATH外の場合は`go mod init {module名}`の形で初期化する必要がある。([ref](https://github.com/golang/go/wiki/Modules#why-does-go-mod-init-give-the-error-cannot-determine-module-path-for-source-directory))  
module名を指定しないと以下のエラーが出る。  
```
go: cannot determine module path for source directory /Users/sakura.yuto/go/src/github.com/mf-sakura/golang_study/module (outside GOPATH, module path must be specified)
```

Module名は何でも良い。  
他のライブラリのImportと形式を合わせるなら、`github.com/{user|organaization}/{レポジトリ名}/{go.modを作るサブディレクトリ}`にする。(GitHubの場合)  
今回は、`go mod init github.com/mf-sakura/golang_study/module`とする。  

## go mod vendor
依存ライブラリの取得をする。これにより`./vendor`配下にライブラリのソースコードがDLされる。  
go.modにはコード内でImportされている標準パッケージと自Module以外のライブラリが追加される。  
```go.mod
module github.com/mf-sakura/golang_study/module

go 1.14

# 最新版が入る。バージョンを変えたい場合はバージョンを変える
require github.com/pkg/errors v0.9.1
```
go.sumというバージョンとハッシュを含んだロックファイルが生成される。  
```go.sum
github.com/pkg/errors v0.9.1 h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=
github.com/pkg/errors v0.9.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
```
## go mod tidy
使用されていないライブラリを消す。  
go.modを以下の様にバージョンだけ変える。  
`go mod vendor`を実行すると、go.sumから元のバージョンの情報が消えない。  
これにより、go.sumが無駄に大きくなっていく。  
```go.mod
module github.com/mf-sakura/golang_study/module

go 1.14

require github.com/pkg/errors v0.9.0
```
```go.sum
github.com/pkg/errors v0.9.0 h1:J8lpUdobwIeCI7OiSxHqEwJUKvJwicL5+3v1oe2Yb4k=
github.com/pkg/errors v0.9.0/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pkg/errors v0.9.1 h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=
github.com/pkg/errors v0.9.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
```
こういうケースでは`go mod tidy`を実行して、go.sumを綺麗にするのが望ましい。
```tidy実行後のgo.sum
github.com/pkg/errors v0.9.0 h1:J8lpUdobwIeCI7OiSxHqEwJUKvJwicL5+3v1oe2Yb4k=
github.com/pkg/errors v0.9.0/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
```


## 自ModuleのImport
`{go.mod に記載されているmodule名}/subdir`の形でImport出来ます。([ref](https://github.com/golang/go/wiki/Modules#do-modules-work-with-relative-imports-like-import-subdir))  

## GO111MODULE
Moduleに関する環境変数([ref](https://github.com/golang/go/wiki/Modules#when-do-i-get-old-behavior-vs-new-module-based-behavior))  
### auto
GOPATH配下かつgo.modが見つからない場合は、Moduleを利用しない。(Go 1.10以前の挙動)  
それ以外の場合では、Moduleを利用する  
### on
ソースコードの場所に関わらずModuleを有効化する
### off
ソースコードの場所に関わらずModuleを無効化する

## Moduleでよくあるエラー
### go mod initをしていない
以下の様にライブラリが見つからないエラーが出る。  
go.modがない事が原因なので、`go mod init`を実行する  
```
main.go:5:2: cannot find package "github.com/pkg/errors" in any of:
    /Users/sakura.yuto/.goenv/versions/1.14.0/src/github.com/pkg/errors (from $GOROOT)
    /Users/sakura.yuto/go/src/github.com/pkg/errors (from $GOPATH)
```
### Moduleがないエラーが出る
`GO111Module=on`の時に出る。`go mod init`を実行する。  
```
go: cannot find main module, but found .git/config in /Users/sakura.yuto/go/src/github.com/mf-sakura/golang_study
    to create a module there, run:
    cd .. && go mod init
```
