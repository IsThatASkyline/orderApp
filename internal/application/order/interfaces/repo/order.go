package repo

import (
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/db/models"
)

type OrderRepo interface {
	GetAllOrders() ([]models.Order, error)
	AddOrder(order *models.Order) error
	GetOrder(uid string) (models.Order, error)
}
