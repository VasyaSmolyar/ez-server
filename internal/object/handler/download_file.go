package handler

import (
	"net/http"
	"strconv"
)

func (h *Handler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
	w.Header().Set("Content-Type", "application/octet-stream")

	err := h.ObjectRepo.DownloadFile(r.Context(), filename, filename)
	if err != nil {
		HandleError(err, w)
		return
	}
	http.ServeFile(w, r, filename)
}
