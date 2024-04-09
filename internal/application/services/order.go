package services

import (
	"github.com/IsThatASkyline/fiberGo/internal/application/order/interfaces/repo"
)

type Service struct {
	Repo repo.OrderRepo
}

func NewOrderService(repo repo.OrderRepo) *Service {
	return &Service{Repo: repo}
}
