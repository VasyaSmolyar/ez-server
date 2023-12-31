package handler

import (
	"encoding/json"
	"net/http"

	"ex-server/internal/task/entity"
)

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	task.UserId = r.Context().Value(userKey{}).(string)

	if err := h.TaskRepo.Create(r.Context(), &task); err != nil {
		HandleError(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
