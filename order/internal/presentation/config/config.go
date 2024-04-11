package config

import (
	db "github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/db/config"
	api "github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/config"
)

type AppConfig struct {
	Mode string
}
type Config struct {
	AppConfig     `toml:"app"`
	db.DBConfig   `toml:"db"`
	api.APIConfig `toml:"api"`
}
