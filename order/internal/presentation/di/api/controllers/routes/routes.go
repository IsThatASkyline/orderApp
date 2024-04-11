package routes

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/controllers/routes"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/controllers/routes/order"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	routes.NewRoutes,
	routes.NewGroupRoutes,

	order.NewOrderRoutes,
)
