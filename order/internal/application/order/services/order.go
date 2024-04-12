package services

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/dto"
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/interfaces/repo"
)

type OrderServiceImpl struct {
	OrderRepo repo.OrderRepo
}

func (interactor *OrderServiceImpl) GetOrderByID(uid string) (dto.Order, error) {
	return interactor.OrderRepo.GetOrder(uid)
}

func (interactor *OrderServiceImpl) GetAllOrders() ([]dto.Order, error) {
	return interactor.OrderRepo.GetAllOrders()
}

func (interactor *OrderServiceImpl) CreateOrder(order *dto.Order) error {
	return interactor.OrderRepo.AddOrder(order)
}
