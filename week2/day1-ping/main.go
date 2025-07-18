// 🎯 ゴール
// Goで net/http を使って、最小のGET API（/ping）を作成
// ブラウザやcurl、Postmanで動作確認できるようにする
// Goでのルーティング・ハンドラー関数の基本構造を理解する

package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

//  チャレンジ課題
// 🔹 /hello?name=Taro のようなリクエストに対応し、「Hello, Taro!」と返すAPIを追加してみよう
func helloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("name")
	if query == "" {
		fmt.Fprintln(w, "名前が指定されていません")
		return
	}
	fmt.Fprintf(w, "Hello,%s!\n", query)
}


func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("サーバー起動エラー:", err)
	}

}