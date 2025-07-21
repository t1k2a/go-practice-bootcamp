package middleware

import (
	"fmt"
	"net/http"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, fmt.Sprintf("内部サーバーエラー: %v", err), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}