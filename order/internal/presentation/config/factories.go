package config

import (
	load "github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/config"
	db "github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/db/config"
	api "github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/config"
)

func NewConfig() Config {
	var conf Config
	load.LoadConfig(&conf, "", "")
	return conf
}
func NewDBConfig(config Config) db.DBConfig {
	return config.DBConfig
}
func NewAppConfig(config Config) AppConfig {
	return config.AppConfig
}
func NewAPIConfig(config Config) api.APIConfig {
	config.APIConfig.Mode = config.AppConfig.Mode
	return config.APIConfig
}
