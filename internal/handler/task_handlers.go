package handler

import (
	"encoding/json"
	"ex-server/internal/entity"
	"math/rand"
	"net/http"
	"strconv"

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

	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range tasks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&entity.Task{})
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task entity.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = strconv.Itoa(rand.Intn(1000000))
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
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
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(tasks)
}
