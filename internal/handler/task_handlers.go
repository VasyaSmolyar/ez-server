package handler

import (
	"encoding/json"
	"ex-server/internal/entity"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var user = entity.User{
	Firstname: "John",
	Lastname:  "Doe",
}

var tasks = []entity.Task{
	{ID: "1", Title: "Gym", Desc: "Go To Gym", Assigned: &user},
	{ID: "2", Title: "Learn Go", Desc: "Write a REST server", Assigned: &user},
}

func (h *Handler) GetTasksList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range tasks {
		if item.ID == params["id"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	task.ID = uuid.New().String()
	tasks = append(tasks, task)
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			var task entity.Task
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = params["id"]
			tasks = append(tasks, task)
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
