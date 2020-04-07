# for

forで定義した変数は初回のみしかAllocateされません。  
なので、forで定義した変数のポインタを使うと想定していない副作用が発生します。  
Let's Encryptのブログに同じ問題でハマった事が書かれています。  
https://jovi0608.hatenablog.com/entry/2020/03/09/094737  

## 対策
linterを変えれば、検知してくれる。  
[golangci-lint](https://github.com/golangci/golangci-lint)に含まれる[Scopelint](https://github.com/kyoh86/scopelint)を使う  

検証  
```
$ make install-lint
$ make golint
# Errorにならない
$ make scopelint
# Errorになる
```

## 課題
numbersの型を変えずに、以下を満たす様にfor文中の処理を修正してください。  
- Scopelintが通る
- numbersの各要素が別々のアドレスを参照する
