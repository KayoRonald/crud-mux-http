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
	v1.HandleFunc("/done/", router.GetTasksDone).Methods("GET")
	v1.HandleFunc("/{id}", router.GetTasksByID).Methods("GET")
	v1.HandleFunc("/", router.PostTasks).Methods("POST")
	v1.HandleFunc("/{id}", router.PutTasks).Methods("PUT")
	v1.HandleFunc("/{id}", router.DeleteTasks).Methods("DELETE")

	log.Print("Rodando na porta 3000")
	http.ListenAndServe(":3000", r)
	return nil
}
