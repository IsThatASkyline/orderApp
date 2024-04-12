package interfaces

import (
	"github.com/gofiber/fiber/v3"
)

type Middleware interface {
	Handle(c *fiber.Ctx)
}
