package handler

import (
	"encoding/json"
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
	getTaskAct := action.NewGetTask(h.TaskRepo)
	item, err := getTaskAct.Do(h.db, id)

	if err != nil {
		/// TODO: нормальная обработка ошибок
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}
