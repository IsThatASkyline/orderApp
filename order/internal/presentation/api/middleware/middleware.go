package middleware

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/middleware/logging"
	"github.com/gofiber/fiber/v3"
)

type Middlewares []fiber.Handler

func NewMiddlewares(
	loggingMiddleware logging.LoggerMiddleware,
) Middlewares {
	return Middlewares{
		loggingMiddleware.Handle,
	}
}
