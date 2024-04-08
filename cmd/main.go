package main

import (
	"github.com/IsThatASkyline/fiberGo/api/routes"
	models2 "github.com/IsThatASkyline/fiberGo/internal/infrastructure/db/models"
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/db/postgres"
	"github.com/gofiber/fiber/v3"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	err = db.AutoMigrate(&models2.Concert{}, &models2.Ticket{})
	if err != nil {
		log.Fatalf("failed to migrate: %s", err.Error())
	}

	// Initialize a new Fiber app
	app := fiber.New()
	// Define a route for the GET method on the root path '/'
	routes.Setup(app)
	// Start the server on port 8000
	log.Fatal(app.Listen(":8000"))
}
