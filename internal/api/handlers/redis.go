package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Get(c *fiber.Ctx) error {
	val, err := h.DBLayer.Get(c.Params("key"))
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": true,
		"value":  val,
	})
}

func (h *Handler) Put(c *fiber.Ctx) error {
	if err := h.DBLayer.Put(c.Params("key"), c.Params("value")); err != nil {
		return c.JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "ok",
	})
}
