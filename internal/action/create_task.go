package action

import (
	"ex-server/internal/entity"

	"gorm.io/gorm"
)

type CreateTask struct {
	taskRepo TaskRepository
}

func NewCreateTask(taskRepo TaskRepository) CreateTask {
	return CreateTask{taskRepo: taskRepo}
}

func (act CreateTask) Do(db *gorm.DB, task *entity.Task) error {
	return act.taskRepo.CreateTask(db, task)
}
