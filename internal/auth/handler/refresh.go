package handler

import (
	"log"
	"net/http"
)

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	token, ok := r.Header["Token"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.JwtHelper.ValidateRefresh(token[0])
	if err != nil {
		HandleError(err, w)
		return
	}
	log.Println(id)

	user, err := h.AuthRepo.Get(r.Context(), id)
	if err != nil {
		HandleError(err, w)
		return
	}

	h.writeToken(w, user)
}
