// 🎯 ゴール
// curl/Postmanを使ってAPIを手動テスト
// GoでAPIハンドラの自動テストを書く
// httptestパッケージに慣れる
// 将来的なE2EやCIへの布石にする

package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"day6-testing/model"
	"day6-testing/handler"

	"fmt"
)

func TestCreateTodo(t *testing.T) {
	payload := model.Todo{Title: "タスクテスト", Done: false}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateTodo(rr, req)

	// 条件分岐は失敗した場合にCLIへ出力される
	if rr.Code != http.StatusOK {
		t.Errorf("期待されたステータス200だが、実際は %d", rr.Code)
	}

	var todo model.Todo
	if err := json.NewDecoder(rr.Body).Decode(&todo); err != nil {
		t.Errorf("レスポンスのJSONデコードに失敗: %v", err)
	}

	if todo.Title != "タスクテスト" {
		t.Errorf("期待されたタイトルと異なる: %s", todo.Title)
	}
}

func TestCreateTodoTitleEmpty(t *testing.T) {
	// 🧠 ミニチャレンジ（任意）
	//  バリデーションエラー時のテストケースを追加する（例：titleが空文字など）

	payload := model.Todo{Title: "", Done: false}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateTodo(rr, req)
	
	// バリデーションエラーの場合は400が期待される
	if rr.Code != http.StatusBadRequest {
		t.Errorf("期待ステータス400だが、実際は %d", rr.Code)
	}

	// エラーレスポンスの確認
	var response map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("レスポンスのJSONデコードに失敗: %v", err)
	}

	// errorsフィールドが存在することを確認
	fmt.Println(response)
	if _, ok := response["errors"]; !ok {
		t.Errorf("エラーレスポンスにerrorsフィールドが存在しない")
	}
}