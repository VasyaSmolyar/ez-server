package action

import (
	"ex-server/internal/entity"

	"gorm.io/gorm"
)

type GetTasksList struct {
	taskRepo TaskRepository
}

func NewGetTasksList(taskRepo TaskRepository) GetTasksList {
	return GetTasksList{taskRepo: taskRepo}
}

func (act GetTasksList) Do(db *gorm.DB) ([]*entity.Task, error) {
	return act.taskRepo.GetTasksList(db)
}
