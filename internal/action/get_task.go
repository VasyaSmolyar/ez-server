package action

import (
	"ex-server/internal/entity"

	"gorm.io/gorm"
)

type GetTask struct {
	taskRepo TaskRepository
}

func NewGetTask(taskRepo TaskRepository) GetTask {
	return GetTask{taskRepo: taskRepo}
}

func (act GetTask) Do(db *gorm.DB, taskID string) (*entity.Task, error) {
	return act.taskRepo.GetTask(db, taskID)
}
