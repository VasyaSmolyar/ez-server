package db

import (
	"ex-server/internal/entity"
	"ex-server/pkg/env"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func Init(cfg *viper.Viper) (*DBConnect, error) {
	db := DBConnect{config: cfg}

	db.newConnection(cfg,
		&entity.Task{},
	)

	return &db, nil
}

type DBConnect struct {
	config     *viper.Viper
	Connection *gorm.DB
}

func (db *DBConnect) newConnection(entities ...interface{}) error {
	conn, err := db.connect()
	if err != nil {
		return err
	}

	if err = migrate(conn, entities); err != nil {
		return err
	}

	db.Connection = conn

	return nil
}

func (db *DBConnect) connect() (*gorm.DB, error) {
	host := env.GetHost(db.config.GetString("DB.Host"), "db")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
		host,
		db.config.GetString("DB.User"),
		db.config.GetString("DB.Name"),
		db.config.GetString("DB.SSLMode"),
		db.config.GetString("DB.Pass"),
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func migrate(conn *gorm.DB, entities []interface{}) error {
	return conn.AutoMigrate(entities...)
}
