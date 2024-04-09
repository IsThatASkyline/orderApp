package main

import (
	"github.com/IsThatASkyline/fiberGo/internal/application/services"
	postgresDB "github.com/IsThatASkyline/fiberGo/internal/infrastructure/db"
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/factories"
	orderRepo "github.com/IsThatASkyline/fiberGo/internal/infrastructure/factories/db"
	"github.com/IsThatASkyline/fiberGo/internal/presentation/api/controllers/handlers"
	"github.com/IsThatASkyline/fiberGo/internal/presentation/api/controllers/routes/order"
	"github.com/gofiber/fiber/v3"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// Initialize a new Fiber app
	orderRepository := orderRepo.BuildOrderRepo(factories.NewBaseRepo(postgresDB.BuildConnection()))
	orderService := services.NewOrderService(orderRepository)

	// Define a route for the GET method on the root path '/'
	handler := handlers.NewHandler(orderService)
	newRoutes := order.NewOrderRoutes(handler)
	app := fiber.New()
	newRoutes.Setup(app)

	// Start the server on port 8000
	log.Fatal(app.Listen(":8080"))
}
