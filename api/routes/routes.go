package routes

import (
	"github.com/IsThatASkyline/fiberGo/api/handlers"
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	app.Get("/", handlers.IndexHandler)
	app.Get("/concerts", handlers.GetConcertsHandler)
	app.Post("/concerts", handlers.CreateConcertHandler)
	app.Get("/concerts/:id", handlers.GetConcertHandler)
}
