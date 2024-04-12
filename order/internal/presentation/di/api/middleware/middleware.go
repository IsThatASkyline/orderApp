package middleware

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/middleware"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/middleware/errorHandler"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/middleware/logging"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	errorHandler.NewErrorMiddleware,
	logging.NewLoggerMiddleware,
	middleware.NewMiddlewares,
)
