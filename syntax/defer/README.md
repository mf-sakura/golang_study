# defer

Tour of Goではdeferの例が実用的でなかったので、実用的な例を記載します。  
サンプルでは、開いたファイルをdeferで閉じています。  
`defer func`の形で無名関数の使用も可能です。  
他にもDBコネクションを閉じたり、DBロールバックする時に使ったりします。  
## 課題
defer内で呼んでいる`file.Close()`はエラーを戻り値に持ちます。  
このエラーのハンドリングをしてください。  
Named Return Valueを上手く活用してください。  
`func CatFile(path string) (err error)`と定義すると、 `err`に代入された値が戻り値として返ります。  
