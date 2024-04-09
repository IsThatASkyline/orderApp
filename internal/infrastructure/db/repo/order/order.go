package handlers

import (
	"errors"
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/db/models"
	"gorm.io/gorm"
)

func GetOrder(db *gorm.DB, uid string) (models.Order, error) {
	var order models.Order
	err := db.Model(&models.Order{}).Preload("Tickets").Find(&order, uid).Error
	if err != nil {
		return order, errors.New("the result set is empty, Orders")
	}

	return order, nil
}

// TODO: Понять что нужно для создания ордера
/*
func SetOrder(order models.Order, db *gorm.DB) error {
	err := db.Create(&models.Order{E: order.E, Tickets: order.Tickets}).Error
	if err != nil {
		return errors.New("Ошибка при создании концерта")
	}
	return nil
}
*/

func GetOrders(db *gorm.DB) ([]models.Order, error) {
	var orders []models.Order
	err := db.Model(&models.Order{}).Preload("Tickets").Find(&orders).Error
	if err != nil {
		return orders, errors.New("the result set is empty, Orders")
	}
	return orders, nil
}
