package adaptor

import "github.com/jackc/pgx/v5"

func Init(db *pgx.Conn) *ObjectRepository {
	return &ObjectRepository{db: db}
}

type ObjectRepository struct {
	db *pgx.Conn
}
