package controllers

import (
	"github.com/IsThatASkyline/fiberGo/internal/presentation/di/api/controllers/handlers"
	"github.com/IsThatASkyline/fiberGo/internal/presentation/di/api/controllers/routes"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handlers.Module,
	routes.Module,
)
