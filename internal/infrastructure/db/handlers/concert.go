package handlers

import (
	"errors"
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/db/models"
	"gorm.io/gorm"
)

func GetConcert(db *gorm.DB, uid string) (models.Concert, error) {
	var concert models.Concert
	err := db.Model(&models.Concert{}).Preload("Tickets").Find(&concert, uid).Error
	if err != nil {
		return concert, errors.New("the result set is empty, Concerts")
	}

	return concert, nil
}

func SetConcert(concert models.Concert, db *gorm.DB) error {
	err := db.Create(&models.Concert{Title: concert.Title, Tickets: concert.Tickets}).Error
	if err != nil {
		return errors.New("Ошибка при создании концерта")
	}
	return nil
}

func GetConcerts(db *gorm.DB) ([]models.Concert, error) {
	var concerts []models.Concert
	err := db.Model(&models.Concert{}).Preload("Tickets").Find(&concerts).Error
	if err != nil {
		return concerts, errors.New("the result set is empty, Concerts")
	}
	return concerts, nil
}
