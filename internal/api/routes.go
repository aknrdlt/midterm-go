package api

import (
	"github.com/gofiber/fiber/v2"
	"rd/internal/api/handlers"
)

func Routes(app *fiber.App, h *handlers.Handler) {
	app.Get("/store/:key", h.Get)
	app.Put("/store/:key/:value", h.Put)
}
