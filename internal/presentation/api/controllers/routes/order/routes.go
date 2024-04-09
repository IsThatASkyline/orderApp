package routes

import (
	"github.com/IsThatASkyline/fiberGo/internal/presentation/api/controllers/handlers"
	"github.com/gofiber/fiber/v3"
)

type Routes struct {
	fiber.Group
	controller handlers.Handler
}

func (r Routes) Setup() {
	router := fiber.New()
	router.Get("/", r.controller.IndexHandler)
	router.Get("/orders", r.controller.GetAllOrders)
	router.Post("/orders", r.controller.CreateOrder)
	router.Get("/orders/:id", r.controller.GetOrder)
}
