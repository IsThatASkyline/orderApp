package order

import (
	"encoding/json"
	"github.com/IsThatASkyline/fiberGo/internal/application/order/interfaces/services"
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/db/models"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type Handler struct {
	service services.OrderService
}

func (h *Handler) GetOrder(c fiber.Ctx) error {
	getOrder, err := h.service.GetOrder(c.Params("id"))
	if err != nil {
		log.Fatalf("hzhz: %s", err.Error())
	}
	return c.JSON(getOrder)
}

func (h *Handler) CreateOrder(c fiber.Ctx) error {
	var newOrder models.Order
	err := json.Unmarshal(c.Body(), &newOrder)
	err = h.service.AddOrder(&newOrder)
	if err != nil {
		log.Fatalf("hzhz: %s", err.Error())
	}

	return nil
}

func (h *Handler) GetAllOrders(c fiber.Ctx) error {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		log.Fatalf("hzhz: %s", err.Error())
	}

	return c.JSON(orders)
}

func (h *Handler) IndexHandler(c fiber.Ctx) error {
	// Send a string response to the client
	return c.SendString("Hello, Leatherman 👋!")
}
