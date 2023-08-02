package adaptor

import (
	"errors"
	"ex-server/internal/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrNotFound error = errors.New("resource was not found")

func Init(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

type TaskRepository struct {
	db *gorm.DB
}

func (repo TaskRepository) GetList() ([]*entity.Task, error) {
	tasks := make([]*entity.Task, 0)

	err := repo.db.Model(&entity.Task{}).Find(&tasks).Error

	return tasks, err
}

func (repo TaskRepository) Get(taskID string) (*entity.Task, error) {
	var task entity.Task
	result := repo.db.Model(&entity.Task{}).First(&task, "id = ?", taskID)
	if result.RowsAffected == 0 {
		return nil, ErrNotFound
	}

	return &task, result.Error
}

func (repo TaskRepository) Create(task *entity.Task) error {
	result := repo.db.Create(task)
	return result.Error
}

func (repo TaskRepository) Update(taskID string, task *entity.Task) (*entity.Task, error) {
	var item entity.Task

	query := repo.db.Model(&item).Where("id = ?", taskID).Clauses(clause.Returning{})
	result := query.Updates(map[string]interface{}{"ID": taskID, "Title": task.Title, "Desc": task.Desc})
	if result.RowsAffected == 0 {
		return nil, ErrNotFound
	}

	return &item, result.Error
}

func (repo TaskRepository) Delete(taskID string) error {
	result := repo.db.Model(&entity.Task{}).Delete("id = ?", taskID)
	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return result.Error
}
