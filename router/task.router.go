package router

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gtihub.com/KayoRonald/crud-mux-http/database"
	"gtihub.com/KayoRonald/crud-mux-http/models"
	"gtihub.com/KayoRonald/crud-mux-http/response"
)



func GetTasks(w http.ResponseWriter, rq *http.Request) {
	tasks := []models.Tasks{}
	result := database.Database.Db.Where("done = ?", false).Find(&tasks)
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Nenhuma tasks criada :/",
			"status":  "error",
		})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTasksDone(w http.ResponseWriter, rq *http.Request) {
	tasks := []models.Tasks{}
	result := database.Database.Db.Where("done = ?", true).Find(&tasks)
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Nenhuma tasks finalizada.",
			"status":  "error",
		})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTasksByID(w http.ResponseWriter, rq *http.Request) {
	w.Header().Add("Content-Type", "application/json")
  tasks := []models.Tasks{}
  id := mux.Vars(rq)["id"]
	result := database.Database.Db.Where("id = ?", id).First(&tasks)
  if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Nenhuma tasks finalizada.",
			"status":  "error",
		})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tasks)
}

func PostTasks(w http.ResponseWriter, rq *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var tasks response.CreateTasks
	err := json.NewDecoder(rq.Body).Decode(&tasks)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
	}
	create := models.Tasks{
		ID:           uuid.New().String(),
		Name:         tasks.Name,
		Descripition: tasks.Descripition,
		Done:         tasks.Done,
	}
	w.WriteHeader(http.StatusCreated)
  database.Database.Db.Create(&create)
	json.NewEncoder(w).Encode(create)
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
