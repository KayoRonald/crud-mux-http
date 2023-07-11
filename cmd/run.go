package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gtihub.com/KayoRonald/crud-mux-http/router"
)

func Run(r *mux.Router) error {
  r.HandleFunc("/", router.GetTasks).Methods("GET")
	r.HandleFunc("/{id}", router.GetTasksByID).Methods("GET")
	r.HandleFunc("/", router.PostTasks).Methods("POST")
	r.HandleFunc("/{id}", router.PutTasks).Methods("PUT")
	r.HandleFunc("/{id}", router.DeleteTasks).Methods("DELETE")

  log.Print("Rodando na porta 8080")
	http.ListenAndServe(":8080", r)
  return nil
}