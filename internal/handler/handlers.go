package handler

import (
	"ex-server/internal/adaptor"
)

func Init(taskRepo adaptor.TaskRepository) *Handler {
	return &Handler{TaskRepo: taskRepo}
}

type Handler struct {
	TaskRepo adaptor.TaskRepository
}
