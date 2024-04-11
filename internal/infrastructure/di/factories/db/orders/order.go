package orders

import (
	appRepo "github.com/IsThatASkyline/fiberGo/internal/application/order/interfaces/repo"
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/db/repo"
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/db/repo/order"
	"go.uber.org/fx"
)

func BuildOrderRepo(base repo.BaseGormRepo) appRepo.OrderRepo {
	return &order.RepoImpl{
		BaseGormRepo: base,
	}
}

var Module = fx.Provide(
	BuildOrderRepo,
)
