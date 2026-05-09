package router

import (
	"go-shorts/internal/database"
	"log"

	"github.com/gofiber/fiber/v3"
)

func RouterApp(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	db := database.GetDatabasePool()

	log.Println("Registering URL routes...", db)
	RegisterUrlRoutes(v1, db)
}
