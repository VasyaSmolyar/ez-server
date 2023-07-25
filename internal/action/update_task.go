package action

import (
	"ex-server/internal/entity"

	"gorm.io/gorm"
)

type UpdateTask struct {
	taskRepo TaskRepository
}

func NewUpdateTask(taskRepo TaskRepository) UpdateTask {
	return UpdateTask{taskRepo: taskRepo}
}

func (act UpdateTask) Do(db *gorm.DB, taskID string, task *entity.Task) (*entity.Task, error) {
	return act.taskRepo.UpdateTask(db, taskID, task)
}
