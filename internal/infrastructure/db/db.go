package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BuildConnection() *gorm.DB {
	// TODO: Убрать в env параметры
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: "host=127.0.0.1 port=5432 user=postgres dbname=orders password=postgres sslmode=disable"}))
	if err == nil {
		migrate(db)
		return db
	}
	panic(err.Error())
}
