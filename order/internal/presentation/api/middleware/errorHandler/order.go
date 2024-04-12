package errorHandler

import (
	"errors"
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/exceptions"
	"github.com/IsThatASkyline/fiberGo/order/internal/presentation/api/controllers/response"
	"net/http"
)

func handleOrderError(e ErrorCatching) {
	var orderIDError *exceptions.OrderIDNotExist

	if errors.As(e.err, &orderIDError) {
		*e.status = http.StatusNotFound
		response.SetExceptionPayload(e.exception, orderIDError.CustomException)
	}
}
