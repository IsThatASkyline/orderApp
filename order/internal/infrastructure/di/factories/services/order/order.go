package order

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/interfaces/repo"
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/interfaces/services"
	order "github.com/IsThatASkyline/fiberGo/order/internal/application/order/services"
	"go.uber.org/fx"
)

func NewOrderService(repo repo.OrderRepo) services.OrderService {
	return &order.OrderServiceImpl{
		OrderRepo: repo,
	}
}

var Module = fx.Provide(
	NewOrderService,
)
