package action

type DeleteTask struct {
	taskRepo TaskRepository
}

func NewDeleteTask(taskRepo TaskRepository) DeleteTask {
	return DeleteTask{taskRepo: taskRepo}
}

func (act DeleteTask) Do(taskID string) error {
	return act.taskRepo.DeleteTask(taskID)
}
