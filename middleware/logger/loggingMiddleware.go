package logger

import (
	"fmt"
	"net/http"
)

func New(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("| %v | %v | %v |", r.Method, r.RequestURI, r.Host)
		next.ServeHTTP(w, r)
	})
}
