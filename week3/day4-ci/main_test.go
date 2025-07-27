package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTodosHandler(t *testing.T) {
	todos = []Todo{
		{ID: 1, Title: "テストタスク1", Done: false},
		{ID: 2, Title: "テストタスク2", Done: true},
	}

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rr := httptest.NewRecorder()

	getTodosHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("期待されたステータス200だが、実際は %d", rr.Code)
	}

	var result []Todo
	if err := json.NewDecoder(rr.Body).Decode(&result); err != nil {
		t.Errorf("レスポンスのJSONデコードに失敗: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("期待されたタスク数は2だが、実際は %d", len(result))
	}
}

func TestCreateTodoHandler(t *testing.T) {
	todos = []Todo{}

	payload := Todo{Title: "新しいタスク", Done: false}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	createTodoHandler(rr, req)


	if rr.Code != http.StatusOK {
		t.Errorf("期待されたステータス200だが、実際は %d", rr.Code)
	}

	var result Todo
	if err := json.NewDecoder(rr.Body).Decode(&result); err != nil {
		t.Errorf("レスポンスのJSONデコードに失敗: %v", err)
	}

	if result.ID != 1 {
		t.Errorf("期待されたIDは1だが、実際は %d", result.ID)
	}

	if result.Title != "新しいタスク" {
		t.Errorf("期待されたタイトルは'新しいタスク'だが、実際は'%s'", result.Title)
	}
}

func TestCreateTodoHandlerEmptyTitle(t *testing.T) {
	todos = []Todo{}

	payload := Todo{Title: "", Done: false}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	createTodoHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
	    t.Errorf("期待されたステータス400だが、実際は %d", rr.Code)
	}    
}