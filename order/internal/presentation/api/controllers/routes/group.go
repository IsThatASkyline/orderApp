package routes

import "github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/engine"

func NewGroupRoutes(handler engine.RequestHandler) engine.GroupRoutes {
	return engine.GroupRoutes{handler.Fiber.Group("/")}
}
