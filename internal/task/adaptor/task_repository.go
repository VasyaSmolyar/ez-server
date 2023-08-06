package adaptor

import (
	"context"
	"errors"
	"ex-server/internal/task/entity"

	"github.com/jackc/pgx/v5"
)

var ErrNotFound error = errors.New("resource was not found")

func Init(db *pgx.Conn) *TaskRepository {
	return &TaskRepository{db: db}
}

type TaskRepository struct {
	db *pgx.Conn
}

func (repo TaskRepository) GetList(ctx context.Context) ([]*entity.Task, error) {
	tasks := make([]*entity.Task, 0)

	var query = "select id, title, description from tasks"
	rows, err := repo.db.Query(ctx, query)
	for rows.Next() {
		var id, title, description string
		if err := rows.Scan(&id, &title, &description); err != nil {
			return tasks, err
		}
		tasks = append(tasks, &entity.Task{
			Id:          id,
			Title:       title,
			Description: description,
		})
	}

	return tasks, err
}

func (repo TaskRepository) Get(ctx context.Context, taskID string) (*entity.Task, error) {
	var id, title, description string
	var query = "select * from tasks where id=$1"
	err := repo.db.QueryRow(ctx, query, taskID).
		Scan(&id, &title, &description)

	return &entity.Task{
		Id:          id,
		Title:       title,
		Description: description,
	}, err
}

func (repo TaskRepository) Create(ctx context.Context, task *entity.Task) error {
	var query = "insert into tasks(title, description) values($1, $2)"
	_, err := repo.db.Exec(ctx, query, task.Title, task.Description)
	return err
}

func (repo TaskRepository) Update(ctx context.Context, taskID string, task *entity.Task) (*entity.Task, error) {
	var id, title, description string
	var query = "update tasks set title=$1, description=$2 where id=$3 returning id, title, description"
	err := repo.db.QueryRow(ctx, query, task.Title, task.Description, taskID).Scan(&id, &title, &description)
	return &entity.Task{
		Id:          id,
		Title:       title,
		Description: description,
	}, err
}

func (repo TaskRepository) Delete(ctx context.Context, taskID string) error {
	var query = "delete from tasks where id=$1"
	_, err := repo.db.Exec(ctx, query, taskID)
	return err
}
