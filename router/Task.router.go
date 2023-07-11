package router

import (
	"encoding/json"
	"net/http"
)

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