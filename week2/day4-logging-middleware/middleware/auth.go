package middleware

import (
	"net/http"

)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Authorizationヘッダーを取得
		token := r.Header.Get("Authorization")
		expectedToken := "Bearer default-secret-token"
		// トークンの検証
		if token != expectedToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// 認証成功時は次のハンドラーへ
		next.ServeHTTP(w, r)
	})
}