// 🎯 ゴール
// Goアプリを無料のクラウドサービスにデプロイ
// .env や PORT などの環境変数に対応
// 最小の構成で「外部からアクセスできるAPI」を動かす

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from %s on port %s\n", os.Getenv("APP_NAME"), port)
	})

	fmt.Println("Starting server on port:", port)
	http.ListenAndServe(":" + port, nil)
}