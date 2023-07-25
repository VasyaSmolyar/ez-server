package adaptor

import (
	"ex-server/internal/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository struct {
}

func (repo TaskRepository) GetTasksList(db *gorm.DB) ([]*entity.Task, error) {
	tasks := make([]*entity.Task, 0)

	err := db.Model(&entity.Task{}).Find(&tasks).Error

	return tasks, err
}

func (repo TaskRepository) GetTask(db *gorm.DB, taskID string) (*entity.Task, error) {
	var task entity.Task
	result := db.Model(&entity.Task{}).First(&task, "id = ?", taskID)
	return &task, result.Error
}

func (repo TaskRepository) CreateTask(db *gorm.DB, task *entity.Task) error {
	result := db.Create(task)
	return result.Error
}

func (repo TaskRepository) UpdateTask(db *gorm.DB, taskID string, task *entity.Task) (*entity.Task, error) {
	var item entity.Task

	query := db.Model(&item).Where("id = ?", taskID).Clauses(clause.Returning{})
	result := query.Updates(map[string]interface{}{"ID": taskID, "Title": task.Title, "Desc": task.Desc})
	return &item, result.Error
}

func (repo TaskRepository) DeleteTask(db *gorm.DB, taskID string) error {
	result := db.Model(&entity.Task{}).Delete("id = ?", taskID)
	return result.Error
}
