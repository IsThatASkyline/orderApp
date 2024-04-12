package services

import "github.com/IsThatASkyline/fiberGo/order/internal/application/order/interfaces/services/usecases"

type OrderService interface {
	usecases.GetOrderByID
	usecases.GetAllOrders
	usecases.CreateOrder
}
