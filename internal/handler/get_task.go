package handler

import (
	"encoding/json"
	"errors"
	"ex-server/internal/action"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	item, err := h.TaskRepo.Get(id)

	switch {
	case errors.Is(err, &action.NotFoundError{}):
		w.WriteHeader(http.StatusNotFound)
		return
	case err != nil:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}
