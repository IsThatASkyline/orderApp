package order

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/interfaces/repo"
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/interfaces/services"
	"go.uber.org/fx"
)

func NewOrderService(repo repo.OrderRepo) services.OrderService {
	return &services.OrderServiceImpl{
		OrderRepo: repo,
	}
}

var Module = fx.Provide(
	NewOrderService,
)
