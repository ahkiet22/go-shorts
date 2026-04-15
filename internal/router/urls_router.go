package router

import (
	"fmt"
	"go-shorts/internal/handler"
	"go-shorts/internal/repository"
	"go-shorts/internal/service"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterUrlRoutes(r fiber.Router, db *pgxpool.Pool) {
	repo := repository.NewUrlRepository(db)
	s := service.NewUrlService(repo)
	h := handler.NewUrlHandler(s)

	url := r.Group("/urls")

	url.Post("/shorten", h.Create)

	url.Get("/:code", func(c fiber.Ctx) error {
		fmt.Println("/:code")
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})
}
