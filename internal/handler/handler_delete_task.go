package handler

import (
	"ex-server/internal/action"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	deleteTaskAct := action.NewDeleteTask(h.TaskRepo)
	err := deleteTaskAct.Do(h.db, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
