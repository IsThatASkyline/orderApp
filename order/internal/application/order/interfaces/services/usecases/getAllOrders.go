package usecases

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/dto"
)

type GetAllOrders interface {
	GetAllOrders() ([]dto.Order, error)
}
