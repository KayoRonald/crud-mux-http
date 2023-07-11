package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gtihub.com/KayoRonald/crud-mux-http/router"
)

func Run(r *mux.Router) error {
  v1 := r.PathPrefix("/tasks/").Subrouter()
  v1.HandleFunc("/", router.GetTasks).Methods("GET")
	v1.HandleFunc("/{id}", router.GetTasksByID).Methods("GET")
	v1.HandleFunc("/", router.PostTasks).Methods("POST")
	v1.HandleFunc("/{id}", router.PutTasks).Methods("PUT")
	v1.HandleFunc("/{id}", router.DeleteTasks).Methods("DELETE")

  log.Print("Rodando na porta 8080")
	http.ListenAndServe(":8080", r)
  return nil
}