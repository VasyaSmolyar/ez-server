package adaptor

import (
	"ex-server/internal/entity"

	"gorm.io/gorm"
)

type TaskRepository struct {
}

func (repo TaskRepository) GetTasksList(db *gorm.DB) ([]*entity.Task, error) {
	tasks := make([]*entity.Task, 0)

	err := db.Model(&entity.Task{}).Find(&tasks).Error

	return tasks, err
}
