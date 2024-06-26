package repo

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/dto"
)

type OrderRepo interface {
	GetAllOrders() ([]dto.Order, error)
	AddOrder(order *dto.Order) error
	GetOrder(uid string) (dto.Order, error)
}
