package order

import (
	"errors"
	appRepo "github.com/IsThatASkyline/fiberGo/order/internal/application/order/interfaces/repo"
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/db/models"
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/db/repo"
)

type RepoImpl struct {
	repo.BaseGormRepo
	appRepo.OrderRepo
}

// TODO: Конвертировать все в ДТО

func (repo *RepoImpl) GetOrder(uid string) (models.Order, error) {
	var order models.Order
	err := repo.Session.Model(&models.Order{}).Preload("Tickets").Find(&order, uid).Error
	if err != nil {
		return order, errors.New("the result set is empty, Orders")
	}

	return order, nil
}

func (repo *RepoImpl) AddOrder(order *models.Order) error {
	result := repo.Session.Omit("Tickets.*").Create(&order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *RepoImpl) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := repo.Session.Model(&models.Order{}).Preload("Tickets").Find(&orders).Error
	if err != nil {
		return orders, errors.New("the result set is empty, Orders")
	}
	return orders, nil
}
