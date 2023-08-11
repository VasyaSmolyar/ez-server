package handler

import (
	"context"
	"net/http"
)

type userKey struct{}

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := r.Header["Token"]
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		user, err := h.AuthAgent.Check(token[0])
		if err != nil {
			HandleError(err, w)
			return
		}

		newCtx := context.WithValue(r.Context(), userKey{}, user.Id)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}
