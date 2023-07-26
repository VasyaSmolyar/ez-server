package adaptor

import (
	"ex-server/internal/action"
	"ex-server/internal/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (repo TaskRepository) GetTasksList() ([]*entity.Task, error) {
	tasks := make([]*entity.Task, 0)

	err := repo.db.Model(&entity.Task{}).Find(&tasks).Error

	return tasks, err
}

func (repo TaskRepository) GetTask(taskID string) (*entity.Task, error) {
	var task entity.Task
	result := repo.db.Model(&entity.Task{}).First(&task, "id = ?", taskID)
	if result.RowsAffected == 0 {
		return nil, &action.NotFoundError{}
	}

	return &task, result.Error
}

func (repo TaskRepository) CreateTask(task *entity.Task) error {
	result := repo.db.Create(task)
	return result.Error
}

func (repo TaskRepository) UpdateTask(taskID string, task *entity.Task) (*entity.Task, error) {
	var item entity.Task

	query := repo.db.Model(&item).Where("id = ?", taskID).Clauses(clause.Returning{})
	result := query.Updates(map[string]interface{}{"ID": taskID, "Title": task.Title, "Desc": task.Desc})
	if result.RowsAffected == 0 {
		return nil, &action.NotFoundError{}
	}

	return &item, result.Error
}

func (repo TaskRepository) DeleteTask(taskID string) error {
	result := repo.db.Model(&entity.Task{}).Delete("id = ?", taskID)
	if result.RowsAffected == 0 {
		return &action.NotFoundError{}
	}

	return result.Error
}
