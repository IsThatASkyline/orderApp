package db

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/db/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	_ = db.AutoMigrate(&models.Ticket{}, &models.Order{})
}
