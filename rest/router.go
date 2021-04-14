package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/plopezlpz/fiber-starter/handler"
)

func mapUrls(app *fiber.App) {
	app.Get("/health", handler.Health)
}
