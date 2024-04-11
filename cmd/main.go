package main

import (
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/di"
	"github.com/IsThatASkyline/fiberGo/internal/presentation/api"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

/*
func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	orderRepository := orderRepo.BuildOrderRepo(factories.NewBaseRepo(postgresDB.BuildConnection()))
	orderService := services.NewOrderService(orderRepository)

	// Define a route for the GET method on the root path '/'
	handler := handlers.NewHandler(orderService)
	newRoutes := order.NewOrderRoutes(handler)
	newRoutes.Setup(app)

	// Start the server on port 8000
	log.Fatal(app.Listen(":8080"))
}
*/

/*
func registerHooks(
	lifecycle fx.Lifecycle,
	app *fiber.App,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go app.Listen(":8080")
				return nil
			},
			OnStop: func(context.Context) error {
				return nil
			},
		},
	)
}

func main() {
	fx.New(
		fx.Provide(postgresDB.BuildConnection),
		fx.Provide(fiber.New),
		fx.Invoke(order.New),
		fx.Invoke(registerHooks),
	).Run()
	// Initialize a new Fiber app
}
*/

func main() {
	fx.New(
		di.Module,
		api.Module,
	).Run()
}
