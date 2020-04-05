# Goのエラー

Goのerrorは以下の通り、Error()というメソッドを持ったInterfaceです。([ref](https://github.com/golang/go/blob/20a838ab94178c55bc4dc23ddc332fce8545a493/src/builtin/builtin.go#L258-L262))  
他の言語の例外とは異なり、関数の戻り値として利用されます。  
```
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}
```

## よくあるエラーハンドリング
よくある例。  
`(error以外の型, error)`が関数の戻り値として返ってくるので、ifでハンドリングする。  
```
if i, err := strconv.Atoi("a"); err != nil {
	fmt.Printf("Atoi of i failed. %v\n", err)
} 
```

## 標準パッケージ
`errors.New`でエラーの生成が可能になります。  

## github.com/pkg/errors
よく使われるライブラリ  
エラーのスタックトレースを出す事が出来ます。  

## 自分で定義したError型
上記の通り、errorはInterfaceなので、errorを実装した型を自分で作れます。  
Interfaceはまだ扱っていないので、雰囲気が分かれば大丈夫です。  
type Assertionによりエラー型の判別が可能です。  
