// 🎯 ゴール
// ・POST/todosでJSONデータを受信
// ・Goの構造体にデコード
// ・レスポンスとしてJSON形式で受け取ったデータをそのまま返す

package main 

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	Title string `json:"title"`
	Done bool `json:"done"`
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POSTメソッドのみ許可されています", http.StatusMethodNotAllowed)
		return
	}

	var todo Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "JSONのでコードに失敗しました", http.StatusBadRequest)
		return 
	}

	// 💡 チャレンジ課題
	// 🔹 title が空だったら 400 Bad Request を返すようにしてみよう！
	if todo.Title == "" {
		http.Error(w, "タイトルを入力してください", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func main() {
	http.HandleFunc("/todos", createTodoHandler)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("サーバー起動エラー:", err)
	}
}