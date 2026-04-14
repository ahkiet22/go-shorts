package main

import (
	"go-shorts/internal/database"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	database.GetDatabasePool()

	app.Listen(":3000")
}
