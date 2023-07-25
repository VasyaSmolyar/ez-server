package action

import "gorm.io/gorm"

type DeleteTask struct {
	taskRepo TaskRepository
}

func NewDeleteTask(taskRepo TaskRepository) DeleteTask {
	return DeleteTask{taskRepo: taskRepo}
}

func (act DeleteTask) Do(db *gorm.DB, taskID string) error {
	return act.taskRepo.DeleteTask(db, taskID)
}
