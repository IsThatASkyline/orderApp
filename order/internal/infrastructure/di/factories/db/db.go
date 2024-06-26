package db

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/db"
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/db/repo"
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/di/factories/db/orders"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func NewBaseRepo(conn *gorm.DB) repo.BaseGormRepo {
	return repo.BaseGormRepo{Session: conn}
}

var Module = fx.Options(
	orders.Module,
	fx.Provide(
		db.BuildConnection,
		NewBaseRepo,
	),
)
