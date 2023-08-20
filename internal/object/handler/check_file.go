package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) CheckFile(w http.ResponseWriter, r *http.Request) {
	filename, ok := mux.Vars(r)["filename"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.ObjectRepo.CheckFile(r.Context(), filename); err != nil {
		log.Println(err)
		HandleError(err, w)
	}
}
