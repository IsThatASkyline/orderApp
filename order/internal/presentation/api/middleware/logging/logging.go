package logging

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/logger"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/middleware/interfaces"
	"github.com/gofiber/fiber/v3"
)

type LoggerMiddleware struct {
	logger.Logger
	interfaces.Middleware
}

func NewLoggerMiddleware(logger logger.Logger) LoggerMiddleware {
	return LoggerMiddleware{
		Logger: logger,
	}
}
func (m LoggerMiddleware) Handle(c fiber.Ctx) error {
	c.Next()
	m.Logger.Infow(
		"Api good answer Request",
		"Request type", c.Method(),
		"Request status", c.Response().StatusCode(),
		"Request path", c.Path(),
	)
	return nil
}
