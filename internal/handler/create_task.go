package handler

import (
	"encoding/json"
	"ex-server/internal/entity"
	"net/http"
)

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.TaskRepo.Create(&task); err != nil {
		HandleError(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
