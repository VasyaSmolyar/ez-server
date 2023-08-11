package handler

import (
	"encoding/json"
	"ex-server/internal/auth/request"
	"net/http"
)

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds request.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.AuthRepo.Signup(r.Context(), creds.Login, creds.Password)
	if err != nil {
		HandleError(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
