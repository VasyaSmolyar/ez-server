package handler

import (
	"encoding/json"
	"errors"
	"ex-server/internal/adaptor"
	"ex-server/internal/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	switch {
	case errors.Is(err, adaptor.ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
		return
	case err != nil:
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(task)
}
