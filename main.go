package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gtihub.com/KayoRonald/crud-mux-http/handle"
	"gtihub.com/KayoRonald/crud-mux-http/middleware"
	"gtihub.com/KayoRonald/crud-mux-http/middleware/logger"
)

func main() {
	r := mux.NewRouter()
	r.Use(logger.New)
	r.Use(middleware.CorsMiddleware)
	r.NotFoundHandler = http.HandlerFunc(handle.NotFound)
	r.HandleFunc("/", GetTasks).Methods("GET")
	r.HandleFunc("/{id}", GetTasksByID).Methods("GET")
	r.HandleFunc("/", PostTasks).Methods("POST")
	r.HandleFunc("/{id}", PutTasks).Methods("PUT")
	r.HandleFunc("/{id}", DeleteTasks).Methods("DELETE")
	log.Print("Rodando na porta 8080")
	http.ListenAndServe(":8080", r)
}

func GetTasks(w http.ResponseWriter, rq *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello method GET",
	})
}

func GetTasksByID(w http.ResponseWriter, rq *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello method GET",
	})
}

func PostTasks(w http.ResponseWriter, rq *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello method Post",
	})
}

func PutTasks(w http.ResponseWriter, rq *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello method Put",
	})
}
func DeleteTasks(w http.ResponseWriter, rq *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello method Delete",
	})
}