package services

import (
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/di/factories/services/order"
	"go.uber.org/fx"
)

var Module = fx.Options(
	order.Module,
)
