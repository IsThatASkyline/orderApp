package factories

import (
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/db/repo"
	"gorm.io/gorm"
)

func NewBaseRepo(conn *gorm.DB) repo.BaseGormRepo {
	return repo.BaseGormRepo{Session: conn}
}
