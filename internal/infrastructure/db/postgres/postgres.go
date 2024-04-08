package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {
	// TODO: Убрать в env параметры
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: "host=127.0.0.1 port=5432 user=postgres dbname=concert password=postgres sslmode=disable"}))
	if err != nil {
		return nil, err
	}
	return db, nil
}
