package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	filename, ok := mux.Vars(r)["filename"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
	w.Header().Set("Content-Type", "application/octet-stream")

	if err := h.ObjectRepo.DownloadFile(r.Context(), filename, filename); err != nil {
		HandleError(err, w)
		return
	}
	http.ServeFile(w, r, filename)
}
