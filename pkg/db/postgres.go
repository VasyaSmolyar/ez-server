package db

import (
	"context"
	"ex-server/pkg/env"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
)

func Init(cfg *viper.Viper) (*DBConnect, error) {
	db := DBConnect{config: cfg}

	db.newConnection()

	return &db, nil
}

type DBConnect struct {
	config     *viper.Viper
	Connection *pgx.Conn
}

func (db *DBConnect) newConnection() error {
	conn, err := db.connect()
	if err != nil {
		return err
	}

	db.Connection = conn

	return nil
}

func (db *DBConnect) connect() (*pgx.Conn, error) {
	host := env.GetHost(db.config.GetString("DB.Host"), "db")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
		host,
		db.config.GetString("DB.User"),
		db.config.GetString("DB.Name"),
		db.config.GetString("DB.SSLMode"),
		db.config.GetString("DB.Pass"),
	)

	return pgx.Connect(context.Background(), dsn)
}
