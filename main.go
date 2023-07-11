package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gtihub.com/KayoRonald/crud-mux-http/middleware/logger"
)

func main() {
	r := mux.NewRouter()
	r.Use(logger.New)
  r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", HandleHello).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func HandleHello(w http.ResponseWriter, rq *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "ahhhhhhhhhh",
	})
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Not Found",
	})
}
