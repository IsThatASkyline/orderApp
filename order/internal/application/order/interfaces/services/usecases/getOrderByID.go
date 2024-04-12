package usecases

import "github.com/IsThatASkyline/fiberGo/order/internal/application/order/dto"

type GetOrderByID interface {
	GetOrderByID(uid string) (dto.Order, error)
}
