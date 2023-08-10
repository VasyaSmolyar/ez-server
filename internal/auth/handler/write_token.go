package handler

import (
	"encoding/json"
	"ex-server/internal/auth/entity"
	"ex-server/internal/auth/response"
	"net/http"
)

func (h *Handler) writeToken(w http.ResponseWriter, user *entity.User) {
	access, err := h.JwtHelper.GenerateJWT(user.Id, user.Login)
	if err != nil {
		HandleError(err, w)
		return
	}

	refresh, err := h.JwtHelper.GenerateRefresh(user.Id)
	if err != nil {
		HandleError(err, w)
		return
	}

	json.NewEncoder(w).Encode(response.Token{
		Access:  access,
		Refresh: refresh,
	})
}
