package handler

import (
	"ex-server/internal/adaptor"

	"gorm.io/gorm"
)

func Init(db *gorm.DB, taskRepo adaptor.TaskRepository) *Handler {
	return &Handler{db: db, TaskRepo: taskRepo}
}

type Handler struct {
	TaskRepo adaptor.TaskRepository
	db       *gorm.DB
}
