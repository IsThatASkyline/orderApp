package handlers

import (
	"github.com/IsThatASkyline/fiberGo/internal/presentation/api/controllers/handlers/order"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	order.NewOrderHandler,
)
