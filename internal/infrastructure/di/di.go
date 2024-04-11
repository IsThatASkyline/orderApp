package di

import (
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/di/factories/db"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"infrastructure.di",
	fx.Options(
		db.Module,
	),
)
