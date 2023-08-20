package handler

import (
	"ex-server/internal/task/adaptor"
	"ex-server/internal/task/agent"
)

func Init(taskRepo adaptor.TaskRepository, authAgent agent.AuthAgent, objectAgent agent.ObjectAgent) *Handler {
	return &Handler{TaskRepo: taskRepo, AuthAgent: authAgent, ObjectAgent: objectAgent}
}

type Handler struct {
	TaskRepo    adaptor.TaskRepository
	AuthAgent   agent.AuthAgent
	ObjectAgent agent.ObjectAgent
}
