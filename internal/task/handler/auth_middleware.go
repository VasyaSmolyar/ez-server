package handler

import (
	"net/http"
)

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := r.Header["Token"]
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		_, err := h.AuthAgent.Check(token[0])
		if err != nil {
			HandleError(err, w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
