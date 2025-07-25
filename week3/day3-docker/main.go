// 🎯 ゴール
// Goアプリを無料のクラウドサービスにデプロイ
// .env や PORT などの環境変数に対応
// 最小の構成で「外部からアクセスできるAPI」を動かす

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// 💡 チャレンジ課題（任意）
	// APP_ENV=production に切り替えて .env.local を読まないようにしてみる
	// .env 読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env ファイルの読み込みに失敗しました")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 🧩 チャレンジ課題（任意）
	// APP_ENV を入れて dev / prod の切り替え判定をしてみる
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "prod" {
		if _, err := os.Stat(".env.local"); err == nil {
			godotenv.Overload(".env.local")
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from %s on port %s\n", os.Getenv("APP_NAME"), port)
	})
	
	fmt.Println("サーバー起動中 : http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}