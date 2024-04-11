package engine

import "github.com/gofiber/fiber/v3"

type RequestHandler struct {
	Fiber *fiber.App
}
type GroupRoutes struct {
	fiber.Router
}

func NewRequestHandler() RequestHandler {
	return RequestHandler{Fiber: fiber.New()}
}
