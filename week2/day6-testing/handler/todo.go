package handler

import (
	"encoding/json"
	"net/http"

	"day6-testing/model"
	"day6-testing/validator"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POSTメソッドのみ許可されています", http.StatusMethodNotAllowed)
		return
	}

	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "JSONデコード失敗", http.StatusBadRequest)
		return
	}

	// バリデーション
	if errs  := validator.ValidateStruct(todo); errs != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": errs,
		})
		return
	}

	// 成功時レスポンス
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}