package main

import (
	"fmt"
	"net/http"

	"day5-validator/handler"
)

func main() {
	http.HandleFunc("/todos", handler.CreateTodo)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("サーバー起動エラー:", err)
	}
}