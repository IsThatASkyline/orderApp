package order

import (
	"encoding/json"
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/dto"
	"github.com/IsThatASkyline/fiberGo/order/internal/application/order/interfaces/services"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	service services.OrderService
}

func (h *Handler) GetOrder(c fiber.Ctx) error {
	orderID := c.Params("id")
	if orderID == "" {
		return fiber.ErrInternalServerError
	}
	order, err := h.service.GetOrderByID(orderID)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON("Page not found")
	}

	return c.JSON(order)
}

func (h *Handler) CreateOrder(c fiber.Ctx) error {
	var newOrder dto.Order
	err := json.Unmarshal(c.Body(), &newOrder)
	if err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return c.JSON("Проверьте отправленные данные")
	}
	err = h.service.CreateOrder(&newOrder)
	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) GetAllOrders(c fiber.Ctx) error {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(err.Error())
	}

	return c.JSON(orders)
}

func (h *Handler) IndexHandler(c fiber.Ctx) error {
	// Send a string response to the client
	return c.SendString("Hello, Leatherman 👋!")
}
