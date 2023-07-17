package router

import (
	"encoding/json"
	"net/http"

	// "gtihub.com/KayoRonald/crud-mux-http/utils"
)

type Teste struct {
	Name         string `json:"name"`
	Descripition string `json:"descripition"`
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
  var teste Teste

  err := json.NewDecoder(rq.Body).Decode(&teste)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(map[string]string{
      "message": err.Error(),
    })
  }
  w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(teste)
}

func PutTasks(w http.ResponseWriter, rq *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello method Put",
	})
}
func DeleteTasks(w http.ResponseWriter, rq *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello method Delete",
	})
}