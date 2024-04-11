package api

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/di/api/controllers"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/di/api/engine"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"presentation.api",
	fx.Options(
		engine.Module,
		controllers.Module,
	),
)
