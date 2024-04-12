package api

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/di/api/controllers"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/di/api/engine"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/di/api/middleware"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"presentation.api",
	fx.Options(
		middleware.Module,
		engine.Module,
		controllers.Module,
	),
)
