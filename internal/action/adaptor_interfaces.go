package action

import (
	"ex-server/internal/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasksList(db *gorm.DB) ([]*entity.Task, error)
}
