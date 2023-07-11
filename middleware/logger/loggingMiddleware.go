package logger

import (
	"log"
	"net/http"
)

func New(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(r.Method, r.RequestURI, r.Method, r.Host)
		next.ServeHTTP(w, r)
	})
}
