package routes

import "github.com/IsThatASkyline/fiberGo/internal/presentation/api/controllers/routes/order"

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	orderRoutes order.Routes,
) Routes {
	return Routes{
		orderRoutes,
	}
}
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
