package handler

import (
	"encoding/json"
	"ex-server/internal/action"
	"ex-server/internal/entity"
	"log"
	"net/http"
)

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	createTaskAct := action.NewCreateTask(h.TaskRepo)
	err = createTaskAct.Do(&task)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
