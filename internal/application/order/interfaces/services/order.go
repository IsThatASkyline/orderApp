package services

import "github.com/IsThatASkyline/fiberGo/internal/application/order/interfaces/repo"

// Не хотел тратить время на новые методы для сервиса, все равно это круд и методы совпадают с репо
type OrderService interface {
	repo.OrderRepo
}
