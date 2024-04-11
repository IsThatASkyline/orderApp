package logger

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/logger"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	logger.NewLogger,
)
