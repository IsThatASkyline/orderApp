package config

import (
	db "github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/db/config"
	logger "github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/logger/config"
	api "github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/config"
)

type AppConfig struct {
	Mode string
}
type Config struct {
	AppConfig           `toml:"app"`
	db.DBConfig         `toml:"db"`
	logger.LoggerConfig `toml:"logging"`
	api.APIConfig       `toml:"api"`
}
