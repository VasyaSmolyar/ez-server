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

	var query = "select id, title, description, user_id from tasks"
	rows, err := repo.db.Query(ctx, query)
	for rows.Next() {
		var id, title, description, userId string
		if err := rows.Scan(&id, &title, &description, &userId); err != nil {
			return tasks, err
		}
		tasks = append(tasks, &entity.Task{
			Id:          id,
			Title:       title,
			Description: description,
			UserId:      userId,
		})
	}

	return tasks, err
}

func (repo TaskRepository) Get(ctx context.Context, taskID string) (*entity.Task, error) {
	var id, title, description string
	var query = "select id, title, description, user_id from tasks where id=$1"
	if err := repo.db.QueryRow(ctx, query, taskID).Scan(&id, &title, &description); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}

	return &entity.Task{
		Id:          id,
		Title:       title,
		Description: description,
	}, nil
}

func (repo TaskRepository) Create(ctx context.Context, task *entity.Task) error {
	var query = "insert into tasks(title, description, user_id) values($1, $2, $3)"
	if _, err := repo.db.Exec(ctx, query, task.Title, task.Description, task.UserId); err == pgx.ErrNoRows {
		return ErrNotFound
	} else if err != nil {
		return err
	}

	return nil
}

func (repo TaskRepository) Update(ctx context.Context, taskID string, task *entity.Task) (*entity.Task, error) {
	var id, title, description, userId string
	var query = "update tasks set title=$1, description=$2 where id=$3 returning id, title, description, user_id"
	if err := repo.db.QueryRow(ctx, query, task.Title, task.Description, taskID).Scan(&id, &title, &description, &userId); err == pgx.ErrNoRows {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &entity.Task{
		Id:          id,
		Title:       title,
		Description: description,
		UserId:      userId,
	}, nil
}

func (repo TaskRepository) Delete(ctx context.Context, taskID string) error {
	var count int
	var query = "WITH deleted AS (delete from tasks where id=$1 IS TRUE RETURNING *) SELECT count(*) FROM deleted"
	if err := repo.db.QueryRow(ctx, query, taskID).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		return ErrNotFound
	}
	return nil
}
