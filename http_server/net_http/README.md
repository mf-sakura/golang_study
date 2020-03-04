# `net/http`を用いたHTTPサーバー
`net/http`はGoの標準パッケージ  
`net/http`をそのまま扱う事は少ないので、簡単に4例だけ紹介します。  
- GET
- GET 200以外のStatus
- GET Headerの読み込み
- POST Bodyの読み込み

# 課題
1. `squareHandler`でnumが100以上の場合はバリデーションエラーとして、400エラーを返してください。
2. `incrementHandler`でPOST以外のMethodの場合に、405エラーを返してください。
3. 自分で1つAPIを追加してください。(何でも良いです)
