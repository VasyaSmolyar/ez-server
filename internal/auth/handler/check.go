package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) Check(w http.ResponseWriter, r *http.Request) {
	token, ok := r.Header["Token"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.JwtHelper.ReadToken(token[0])
	if err != nil {
		HandleError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)
}
