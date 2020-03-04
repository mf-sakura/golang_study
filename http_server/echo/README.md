# `echo`を用いたHTTPサーバー
`github.com/labstack/echo`は軽量なWeb Application Frameworkです。  
この様なFrameworkを使う事が多いです。  
`net/http`の例と同様に以下の4パターンを扱います。  
- GET
- GET 200以外のStatus
- GET Headerの読み込み
- POST Bodyの読み込み

# 課題
1. `squareHandler`でnumが100以上の場合はバリデーションエラーとして、400エラーを返してください。
2. `incrementHandler`にPOST以外のリクエストを行った場合、405エラーが返る事を確認してください。  
追加実装しなくてもechoの機能でリクエストが弾かれます。
3. 自分で1つAPIを追加してください。(何でも良いです)
4. `incrementHandler`で以下のJSONの形式でCounterの値を返してください。
```
{"counter": 9}
```
参考: https://echo.labstack.com/guide/response
