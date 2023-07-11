package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gtihub.com/KayoRonald/crud-mux-http/database"
	"gtihub.com/KayoRonald/crud-mux-http/handle"
	"gtihub.com/KayoRonald/crud-mux-http/middleware"
	"gtihub.com/KayoRonald/crud-mux-http/middleware/logger"
	"gtihub.com/KayoRonald/crud-mux-http/router"
)

func main() {
	r := mux.NewRouter()
	r.Use(logger.New)
	r.Use(middleware.CorsMiddleware)
	r.NotFoundHandler = http.HandlerFunc(handle.NotFound)
	// connect database
	database.ConnectDB()
	r.HandleFunc("/", router.GetTasks).Methods("GET")
	r.HandleFunc("/{id}", router.GetTasksByID).Methods("GET")
	r.HandleFunc("/", router.PostTasks).Methods("POST")
	r.HandleFunc("/{id}", router.PutTasks).Methods("PUT")
	r.HandleFunc("/{id}", router.DeleteTasks).Methods("DELETE")
	
  log.Print("Rodando na porta 8080")
	http.ListenAndServe(":8080", r)
}