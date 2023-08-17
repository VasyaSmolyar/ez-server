package handler

import (
	"encoding/json"
	"ex-server/internal/object/response"
	"net/http"
)

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	path, contentType, err := FileUploadToLocal("filename", r)
	if err != nil {
		HandleError(err, w)
		return
	}

	err = h.ObjectRepo.UploadFile(r.Context(), path, path, contentType)
	if err != nil {
		HandleError(err, w)
		return
	}

	json.NewEncoder(w).Encode(response.File{FileName: path})
}
