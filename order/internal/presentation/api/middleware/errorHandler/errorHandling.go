package errorHandler

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/logger"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/controllers/response"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/middleware/interfaces"
)

type ErrorMiddleware struct {
	interfaces.Middleware
	logger.Logger
}
type ErrorCatching struct {
	status    *int
	err       error
	exception *response.ExceptionResponse
}
type ErrorStatus struct {
	status    int
	exception any
}

func NewErrorMiddleware(logger logger.Logger) ErrorMiddleware {
	return ErrorMiddleware{
		Logger: logger,
	}
}
