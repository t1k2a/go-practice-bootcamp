package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}

var (
	todos []Todo
	mutex sync.Mutex
)

// GET /todos - タスク一覧を取得
func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}


// PUT /todos/{id} - タスクのDoneステータスを更新
func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Path[len("/todos/"):]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "IDが無効です", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for i, task := range todos {
		if task.ID == id {
			var updatedTodo Todo
			if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
				http.Error(w, "JSONデコードエラー", http.StatusBadRequest)
				return
			}
			todos[i].Done = updatedTodo.Done
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}
	http.Error(w, "タスクが見つかりません", http.StatusNotFound)
}

// DELETE /todos/{id} - タスクを削除
func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Path[len("/todos/"):]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "IDが無効です", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for i, task := range todos {
		if task.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // 204 noContent
			return 
		}
	}
	http.Error(w, "タスクが見つかりません", http.StatusNotFound)
}

// POST /todos -新しいタスクを追加
func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "JSONデコードに失敗しました", http.StatusBadRequest)
		return
	}

	if todo.Title == "" {
		http.Error(w, "タイトルを入力してください", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	todo.ID = len(todos) + 1 // IDはタスク数に基づいて決定
	todos = append(todos, todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func main() {
	r := mux.NewRouter()

	// CRUD操作のルーティングを分ける
	r.HandleFunc("/todos", createTodoHandler).Methods("POST") // POST
	r.HandleFunc("/todos", getTodosHandler).Methods("GET") // GET
	r.HandleFunc("/todos/{id:[0-9]+}", updateTodoHandler).Methods("PUT") // PUT
	r.HandleFunc("/todos/{id:[0-9]+}", deleteTodoHandler).Methods("DELETE") // DELETE

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("サーバー起動エラー:", err)
	}
}


//  1. POST - 新しいタスクを作成
//   curl -X POST http://localhost:8080/todos \
//     -H "Content-Type: application/json" \
//     -d '{"title": "買い物に行く", "done": false}'

//   2. GET - タスク一覧を取得
//   curl http://localhost:8080/todos

//   3. PUT - タスクのステータスを更新
//   curl -X PUT http://localhost:8080/todos/1 \
//     -H "Content-Type: application/json" \
//     -d '{"done": true}'

//   4. DELETE - タスクを削除
//   curl -X DELETE http://localhost:8080/todos/1