# Day3 Todo API 開発日誌

## 開発概要
REST APIのCRUD操作（GET, PUT, DELETE）を実装したTodoリストAPI

## 実装内容

### エンドポイント
- `GET /todos` - タスク一覧を取得
- `POST /todos` - 新しいタスクを作成
- `PUT /todos/{id}` - タスクのステータスを更新
- `DELETE /todos/{id}` - タスクを削除

### 主な変更点

#### 1. ルーティングの問題解決
最初は`http.HandleFunc`を使用していたが、同じパスに複数のハンドラーを登録できない問題に遭遇。
```go
// エラーが発生したコード
http.HandleFunc("/todos/", getTodosHandler)   // GET
http.HandleFunc("/todos/", updateTodoHandler) // PUT
http.HandleFunc("/todos/", deleteTodoHandler) // DELETE
```

解決策として`gorilla/mux`パッケージを導入し、HTTPメソッドごとにルーティングを分けた。

#### 2. バリデーション追加
チャレンジ課題として、タスクのタイトルが空の場合のバリデーションを実装。
```go
if todo.Title == "" {
    http.Error(w, "タイトルを入力してください", http.StatusBadRequest)
    return
}
```

#### 3. 開発環境改善
`air`を導入してホットリロード環境を構築。ファイル変更時に自動的にサーバーが再起動される。

### テストコマンド
```bash
# POST - 新しいタスクを作成
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "買い物に行く", "done": false}'

# GET - タスク一覧を取得
curl http://localhost:8080/todos

# PUT - タスクのステータスを更新
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"done": true}'

# DELETE - タスクを削除
curl -X DELETE http://localhost:8080/todos/1
```

## 学んだこと
- Go標準の`net/http`の制限とサードパーティルーターの必要性
- Mutexを使った並行安全性の確保
- RESTful APIの設計原則
- ホットリロード環境の構築方法

## 次の改善点
- データベース連携
- エラーハンドリングの改善
- テストの追加
- 認証機能の実装