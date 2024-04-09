package handlers

import (
	"encoding/json"
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/db/handlers"
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/db/models"
	"github.com/IsThatASkyline/fiberGo/internal/infrastructure/db/postgres"
	"github.com/gofiber/fiber/v3"
	"log"
)

func IndexHandler(c fiber.Ctx) error {
	// Send a string response to the client
	return c.SendString("Hello, Leatherman 👋!")
}

func GetConcertHandler(c fiber.Ctx) error {
	// TODO: Сделать di
	db, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatalf("qq: %s", err.Error())
	}
	concert, err := handlers.GetConcert(db, c.Params("id"))
	if err != nil {
		log.Fatalf("hzhz: %s", err.Error())
	}
	return c.JSON(concert)
}

func CreateConcertHandler(c fiber.Ctx) error {
	// TODO: Сделать di
	db, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatalf("qq: %s", err.Error())
	}

	var concert models.Concert
	err = json.Unmarshal(c.Body(), &concert)
	err = handlers.SetConcert(concert, db)
	if err != nil {
		log.Fatalf("hzhz: %s", err.Error())
	}

	return nil
}

func GetConcertsHandler(c fiber.Ctx) error {
	db, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatalf("qq: %s", err.Error())
	}

	concerts, err := handlers.GetConcerts(db)
	if err != nil {
		log.Fatalf("hzhz: %s", err.Error())
	}

	return c.JSON(concerts)
}
