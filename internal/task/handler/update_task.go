package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"ex-server/internal/task/entity"
)

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var item entity.Task
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task, err := h.TaskRepo.Update(id, &item)
	if err != nil {
		HandleError(err, w)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(task)
}
