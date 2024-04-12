package order

import (
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/dto"
	"github.com/IsThatASkyline/fiberGo/order/internal/infrastructure/db/models"
)

func ConvertOrderModelToDTO(model models.Order) dto.Order {
	return dto.Order{
		OrderID:       int(model.ID),
		OrderStatus:   model.OrderStatus,
		PaymentMethod: model.PaymentMethod,
		ClientID:      model.ClientID,
		Tickets:       nil, // Потом сделать
	}
}
func ConvertOrderDTOToModel(order *dto.Order) models.Order {
	return models.Order{
		ID:            uint(order.OrderID),
		OrderStatus:   order.OrderStatus,
		ClientID:      order.ClientID,
		PaymentMethod: order.PaymentMethod,
		Tickets:       nil, // Потом сделать
	}
}
