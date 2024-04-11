package order

import (
	"github.com/IsThatASkyline/fiberGo/internal/presentation/api/controllers/handlers/order"
	"github.com/IsThatASkyline/fiberGo/internal/presentation/api/engine"
)

type Routes struct {
	engine.GroupRoutes
	controller order.Handler
}

func (r Routes) Setup() {
	r.Get("/", r.controller.IndexHandler)
	r.Get("/orders", r.controller.GetAllOrders)
	r.Post("/orders", r.controller.CreateOrder)
	r.Get("/orders/:id", r.controller.GetOrder)
}

func NewOrderRoutes(
	group engine.GroupRoutes,
	controller order.Handler,
) Routes {
	return Routes{controller: controller, GroupRoutes: group}
}
