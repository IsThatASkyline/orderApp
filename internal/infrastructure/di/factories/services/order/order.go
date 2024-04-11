package order

import (
	"github.com/IsThatASkyline/fiberGo/internal/application/order/interfaces/repo"
	"github.com/IsThatASkyline/fiberGo/internal/application/order/interfaces/services"
	"go.uber.org/fx"
)

type OrderServiceImpl struct {
	repo.OrderRepo
}

func NewOrderService(repo repo.OrderRepo) services.OrderService {
	return &OrderServiceImpl{
		OrderRepo: repo,
	}
}

var Module = fx.Provide(
	NewOrderService,
)
