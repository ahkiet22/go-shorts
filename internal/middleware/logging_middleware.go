package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func Logging() fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		log.Printf(
			"%s %s | %dms | %d",
			c.Method(),
			c.Path(),
			time.Since(start).Milliseconds(),
			c.Response().StatusCode(),
		)

		return err
	}
}
