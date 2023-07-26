package action

import (
	"ex-server/internal/entity"
)

type GetTask struct {
	taskRepo TaskRepository
}

func NewGetTask(taskRepo TaskRepository) GetTask {
	return GetTask{taskRepo: taskRepo}
}

func (act GetTask) Do(taskID string) (*entity.Task, error) {
	return act.taskRepo.GetTask(taskID)
}
