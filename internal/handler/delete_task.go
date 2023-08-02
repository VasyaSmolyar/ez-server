package handler

import (
	"errors"
	"ex-server/internal/action"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.TaskRepo.Delete(id)

	switch {
	case errors.Is(err, &action.NotFoundError{}):
		w.WriteHeader(http.StatusNotFound)
		return
	case err != nil:
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
