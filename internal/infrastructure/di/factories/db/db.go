package factories

import (
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/db"
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/db/repo"
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/di/factories/db/orders"
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
