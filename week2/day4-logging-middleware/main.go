package main

import (
	"fmt"
	"log"
	"net/http"
	"day4-logging-middleware/handlers"
	"day4-logging-middleware/middleware"
	"sync"
	"os"

	"github.com/gorilla/mux"
)

var (
	mutex sync.Mutex
)

func main() {
	r := mux.NewRouter()

	expectedToken := os.Getenv("API_TOKEN")

	if expectedToken == "" {
		expectedToken = "default-secret-token"
	}


	// ハンドラーの登録
	r.HandleFunc("/todos", handlers.GetTodosHandler).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", handlers.UpdateTodoHandler).Methods("PUT")
	r.HandleFunc("/todos/{id:[0-9]+}", handlers.DeleteTodoHandler).Methods("DELETE")

	// ミドルウェアをチェーンで適用

	//💡 チャレンジ課題（1つだけ）
	// ミドルウェアに 認証処理 を追加してみよう。
	// 例えば、Authorization ヘッダーをチェックし、認証トークンが無ければリクエストを拒否する処理を実装。
	authRouter := middleware.Auth(r)
	loggedRouter := middleware.Logging(middleware.Recovery(authRouter))


	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}

