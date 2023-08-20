package handler

import (
	"ex-server/internal/object/adaptor"
)

func Init(objectRepo adaptor.ObjectRepository) *Handler {
	return &Handler{ObjectRepo: objectRepo}
}

type Handler struct {
	ObjectRepo adaptor.ObjectRepository
}
