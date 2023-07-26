package action

import (
	"ex-server/internal/entity"
)

type CreateTask struct {
	taskRepo TaskRepository
}

func NewCreateTask(taskRepo TaskRepository) CreateTask {
	return CreateTask{taskRepo: taskRepo}
}

func (act CreateTask) Do(task *entity.Task) error {
	return act.taskRepo.CreateTask(task)
}
