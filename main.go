package main

import (
	"go-shorts/internal/database"
	"go-shorts/internal/middleware"
	"go-shorts/internal/router"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	defer func() {
		database.GetDatabasePool().Close()
	}()

	app.Use(middleware.Logging())

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	router.RouterApp(app)

	database.GetDatabasePool()

	app.Listen(":3000")
}
