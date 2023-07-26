package handler

import (
	"errors"
	"ex-server/internal/action"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	deleteTaskAct := action.NewDeleteTask(h.TaskRepo)
	err := deleteTaskAct.Do(id)

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		w.WriteHeader(http.StatusNotFound)
		return
	case err != nil:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
