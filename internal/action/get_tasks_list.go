package action

import (
	"ex-server/internal/entity"
)

type GetTasksList struct {
	taskRepo TaskRepository
}

func NewGetTasksList(taskRepo TaskRepository) GetTasksList {
	return GetTasksList{taskRepo: taskRepo}
}

func (act GetTasksList) Do() ([]*entity.Task, error) {
	return act.taskRepo.GetTasksList()
}
