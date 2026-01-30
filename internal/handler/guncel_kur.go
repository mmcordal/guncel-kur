package handler

import (
	"guncel-kur/internal/service"

	"github.com/gofiber/fiber/v2"
)

type KurHandler struct {
	service service.KurService
}

func NewKurHandler(s service.KurService) *KurHandler {
	return &KurHandler{service: s}
}

func (h *KurHandler) GetKur(c *fiber.Ctx) error {
	resp, err := h.service.FetchFromTDV()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(resp)
}
