package routes

import (
	c "github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/config"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/engine"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/middleware"
)

func NewGroupRoutes(config c.APIConfig, handler engine.RequestHandler, middlewares middleware.Middlewares) engine.GroupRoutes {
	return engine.GroupRoutes{Router: handler.Fiber.Group(config.BaseURLPrefix, middlewares...)}
}
