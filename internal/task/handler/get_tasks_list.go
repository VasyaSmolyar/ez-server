package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetTasksList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks, err := h.TaskRepo.GetList()
	if err != nil {
		HandleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
