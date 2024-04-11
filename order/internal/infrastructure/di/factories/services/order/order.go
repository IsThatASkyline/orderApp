package order

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/interfaces/repo"
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/interfaces/services"
	"go.uber.org/fx"
)

type orderServiceImpl struct {
	repo.OrderRepo
}

func NewOrderService(repo repo.OrderRepo) services.OrderService {
	return &orderServiceImpl{
		OrderRepo: repo,
	}
}

var Module = fx.Provide(
	NewOrderService,
)
