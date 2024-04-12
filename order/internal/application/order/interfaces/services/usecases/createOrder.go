package usecases

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/dto"
)

type CreateOrder interface {
	CreateOrder(order *dto.Order) error
}
