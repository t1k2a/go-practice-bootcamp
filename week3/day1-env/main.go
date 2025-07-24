// 🎯 ゴール
// .env ファイルに設定を定義（PORTなど）
// os.Getenv() で値を取得
// godotenv パッケージを使って .env を読み込む
// JSONレスポンスに設定値を出力して動作確認

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
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
	if appEnv == "dev" {
		if _, err := os.Stat(".env.local"); err == nil {
			godotenv.Overload(".env.local")
		}
	}

	// 環境変数を取得（.env.localの読み込み後）
	appName := os.Getenv("APP_NAME")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		config := map[string]string{
			"app_name": appName,
			"app_env": appEnv,
			"port": port,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(config)
	})

	fmt.Println("サーバー起動中 : http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}