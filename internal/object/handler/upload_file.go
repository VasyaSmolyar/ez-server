package handler

import (
	"encoding/json"
	"ex-server/internal/object/response"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	path, contentType, err := FileUploadToLocal("filename", r)
	if err != nil {
		HandleError(err, w)
		return
	}

	objectName := uuid.New().String()
	if err = h.ObjectRepo.UploadFile(r.Context(), objectName, path, contentType); err != nil {
		HandleError(err, w)
		return
	}

	json.NewEncoder(w).Encode(response.File{FileName: objectName})
}
