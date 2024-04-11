package config

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/config"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"presentation.configs",
	fx.Provide(
		config.NewConfig,
		config.NewDBConfig,
		config.NewAPIConfig,
		config.NewAppConfig,
		config.NewLoggerConfig,
	),
)
