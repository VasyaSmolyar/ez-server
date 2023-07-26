package action

import (
	"ex-server/internal/entity"
)

type TaskRepository interface {
	GetTasksList() ([]*entity.Task, error)
	GetTask(taskID string) (*entity.Task, error)
	CreateTask(task *entity.Task) error
	UpdateTask(taskID string, task *entity.Task) (*entity.Task, error)
	DeleteTask(taskID string) error
}
