起動手順  
go mod tidy  
でパッケージをインストール
  
go run main.go  
でサーバーを起動

動作確認  
curlコマンドを使って確認する。  

POST  
curl -i -X POST -H "Content-Type: application/json" -d '{"title":"タイトル"}' http://localhost:8080/todos
  
GET  
curl -i http://localhost:8080/todos
  
PUT  
curl -i -X PUT -H "Content-Type: application/json" -d '{"title":"タイトル"}' http://localhost:8080/todos/1
  
DELETE  
curl -i -X DELETE http://localhost:8080/todos/1# TODO_API

技術の選定理由  
フレームワークとしてGinを利用した。メジャーなフレームワークであり、情報が多くある為使用。
データベースはSQLiteを使用した。これにより、別途データベースをインストールせずともAPIを動かせる為使用。

工夫点  
404ステータスはresult.RowsAffected()を使って影響が0行の場合に返すようにした。