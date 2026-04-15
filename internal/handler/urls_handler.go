package handler

import (
	"go-shorts/internal/dto"
	"go-shorts/internal/service"

	"github.com/gofiber/fiber/v3"
)

type UrlHandler struct {
	service *service.UrlService
}

func NewUrlHandler(s *service.UrlService) *UrlHandler {
	return &UrlHandler{service: s}
}

func (h *UrlHandler) Create(c fiber.Ctx) error {
	var req dto.URLCreateDTO
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	data, err := h.service.Create(req.URL, req.ExpiresAt)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(fiber.Map{
		"short_code": data,
		"message":    "URL shortened successfully",
	})
}

// func (h *UrlHandler) Redirect(c fiber.Ctx) error {
// 	code := c.Params("code")

// 	url, err := h.service.GetOriginalURL(code)
// 	if err != nil {
// 		return c.Status(404).SendString("Not found")
// 	}

// 	// return c.Redirect(url)
// }
