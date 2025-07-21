package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"day4-logging-middleware/models"

	// "github.com/gorilla/mux"
)




var (
	todos []models.Todo
	mutex sync.Mutex
)


func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// POST /todos -新しいタスクを追加
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
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

// PUT /todos/{id} - タスクのDoneステータスを更新
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
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
			var updatedTodo models.Todo
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
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
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