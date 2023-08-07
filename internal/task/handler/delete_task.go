package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.TaskRepo.Delete(r.Context(), id); err != nil {
		HandleError(err, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
