package action

import (
	"ex-server/internal/entity"
)

type UpdateTask struct {
	taskRepo TaskRepository
}

func NewUpdateTask(taskRepo TaskRepository) UpdateTask {
	return UpdateTask{taskRepo: taskRepo}
}

func (act UpdateTask) Do(taskID string, task *entity.Task) (*entity.Task, error) {
	return act.taskRepo.UpdateTask(taskID, task)
}
