package action

import (
	"ex-server/internal/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasksList(db *gorm.DB) ([]*entity.Task, error)
	GetTask(db *gorm.DB, taskID string) (*entity.Task, error)
	CreateTask(db *gorm.DB, task *entity.Task) error
	UpdateTask(db *gorm.DB, taskID string, task *entity.Task) (*entity.Task, error)
	DeleteTask(db *gorm.DB, taskID string) error
}
