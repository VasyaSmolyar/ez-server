package handler

import (
	"ex-server/internal/auth/adaptor"
)

func Init(authRepo adaptor.AuthRepository) *Handler {
	return &Handler{AuthRepo: authRepo}
}

type Handler struct {
	AuthRepo adaptor.AuthRepository
}
