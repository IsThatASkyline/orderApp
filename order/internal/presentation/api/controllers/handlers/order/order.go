package order

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/interfaces/services"
)

func NewOrderHandler(
	service services.OrderService,
) Handler {
	return Handler{service: service}
}
