package handler

import (
	"ex-server/internal/task/adaptor"
	"ex-server/internal/task/agent"
)

func Init(taskRepo adaptor.TaskRepository, authAgent agent.AuthAgent) *Handler {
	return &Handler{TaskRepo: taskRepo, AuthAgent: authAgent}
}

type Handler struct {
	TaskRepo  adaptor.TaskRepository
	AuthAgent agent.AuthAgent
}
