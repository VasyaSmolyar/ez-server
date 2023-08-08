package adaptor

import (
	"errors"

	"github.com/jackc/pgx/v5"
)

var ErrNotFound error = errors.New("resource was not found")

func Init(db *pgx.Conn) *AuthRepository {
	return &AuthRepository{db: db}
}

type AuthRepository struct {
	db *pgx.Conn
}
