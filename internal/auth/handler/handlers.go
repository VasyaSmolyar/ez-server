package handler

import (
	"ex-server/internal/auth/adaptor"
	"ex-server/internal/auth/jwt"
)

func Init(authRepo adaptor.AuthRepository, jwtHelper jwt.JWTHelper) *Handler {
	return &Handler{AuthRepo: authRepo, JwtHelper: jwtHelper}
}

type Handler struct {
	AuthRepo  adaptor.AuthRepository
	JwtHelper jwt.JWTHelper
}
