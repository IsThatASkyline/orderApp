package order

import (
	"errors"
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/dto"
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/exceptions"
	appRepo "github.com/IsThatASkyline/fiberGo/order/internal/application/order/interfaces/repo"
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/db/models"
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/db/repo"
	"gorm.io/gorm"
)

type RepoImpl struct {
	repo.BaseGormRepo
	appRepo.OrderRepo
}

// TODO: Конвертировать все в ДТО

func (repo *RepoImpl) GetOrder(uid string) (dto.Order, error) {
	var order models.Order
	result := repo.Session.Preload("Tickets").Where("id = ?", uid).Find(&order)
	err := exceptions.OrderIDNotExist{}.Exception(uid)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return dto.Order{}, &err
	}
	if result.RowsAffected == 0 {
		return dto.Order{}, &err
	}
	if result.Error != nil {
		return dto.Order{}, result.Error
	}
	return ConvertOrderModelToDTO(order), nil
}

func (repo *RepoImpl) AddOrder(order *dto.Order) error {
	model := ConvertOrderDTOToModel(order)
	result := repo.Session.Omit("Tickets.*").Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *RepoImpl) GetAllOrders() ([]dto.Order, error) {
	var ordersModel []models.Order
	err := repo.Session.Model(&models.Order{}).Preload("Tickets").Find(&ordersModel).Error
	if err != nil {
		return []dto.Order{}, errors.New("the result set is empty, Orders")
	}
	orders := make([]dto.Order, len(ordersModel))
	for index, orderModel := range ordersModel {
		order := ConvertOrderModelToDTO(orderModel)
		orders[index] = order
	}
	return orders, nil
}
