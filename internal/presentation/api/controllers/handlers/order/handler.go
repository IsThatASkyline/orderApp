package order

import (
	"github.com/gofiber/fiber/v3"
)

/*
	type Handler struct {
		service *services.Service
	}

	func NewHandler(service *services.Service) *Handler {
		return &Handler{service: service}
	}

	func (h *Handler) GetOrder(c fiber.Ctx) error {
		getOrder, err := h.service.Repo.GetOrder(c.Params("id"))
		if err != nil {
			log.Fatalf("hzhz: %s", err.Error())
		}
		return c.JSON(getOrder)
	}

	func (h *Handler) CreateOrder(c fiber.Ctx) error {
		var newOrder models.Order
		err := json.Unmarshal(c.Body(), &newOrder)
		err = h.service.Repo.AddOrder(&newOrder)
		if err != nil {
			log.Fatalf("hzhz: %s", err.Error())
		}

		return nil
	}

	func (h *Handler) GetAllOrders(c fiber.Ctx) error {
		orders, err := h.service.Repo.GetAllOrders()
		if err != nil {
			log.Fatalf("hzhz: %s", err.Error())
		}

		return c.JSON(orders)
	}
*/
type Handler struct {
}

func (h *Handler) IndexHandler(c fiber.Ctx) error {
	// Send a string response to the client
	return c.SendString("Hello, Leatherman 👋!")
}
