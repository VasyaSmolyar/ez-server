package handler

import (
	"encoding/json"
	"ex-server/internal/auth/request"
	"ex-server/internal/auth/response"
	"net/http"
)

func (h *Handler) Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds request.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.AuthRepo.Signin(r.Context(), creds.Login, creds.Password)
	if err != nil {
		HandleError(err, w)
		return
	}

	access, err := h.JwtHelper.GenerateJWT(user.Id, user.Login)
	if err != nil {
		HandleError(err, w)
		return
	}

	refresh, err := h.JwtHelper.GenerateRefresh()
	if err != nil {
		HandleError(err, w)
		return
	}

	json.NewEncoder(w).Encode(response.Token{
		Access:  access,
		Refresh: refresh,
	})
}
