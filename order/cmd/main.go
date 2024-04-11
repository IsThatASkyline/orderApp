package main

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/di"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/di/config"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.Module,
		config.Module,
		api.Module,
	).Run()
}
