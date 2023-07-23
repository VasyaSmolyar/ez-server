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

var tasks = map[string]entity.Task{
	"1111-1111-1111-1111": {ID: "1111-1111-1111-1111", Title: "Gym", Desc: "Go To Gym", Assigned: &user},
	"1111-1111-1111-1112": {ID: "1111-1111-1111-1112", Title: "Learn Go", Desc: "Write a REST server", Assigned: &user},
}

func (h *Handler) GetTasksList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	item, ok := tasks[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	task.ID = uuid.New().String()
	tasks[task.ID] = task
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, ok = tasks[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var item entity.Task

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item.ID = id

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(item)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, ok = tasks[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	delete(tasks, id)
	w.WriteHeader(http.StatusNoContent)
}
