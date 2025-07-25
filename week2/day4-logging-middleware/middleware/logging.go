package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[LOG] %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("[DONE] %s %s - %s", r.Method, r.URL.Path, time.Since(start))
	})
}