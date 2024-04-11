package db

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/db/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BuildConnection(config config.DBConfig) *gorm.DB {
	// TODO: Убрать в env параметры
	// db, err := gorm.Open(postgres.New(postgres.Config{DSN: "host=127.0.0.1 port=5432 user=postgres dbname=orders password=postgres sslmode=disable"}))
	db, err := gorm.Open(postgres.Open(config.FullDNS()))
	sqlDB, errSQL := db.DB()
	if errSQL != nil {
		panic(errSQL)
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConnection)
	if err == nil {
		if config.Migration {
			migrate(db)
		}

		return db
	}
	panic(err.Error())
}
