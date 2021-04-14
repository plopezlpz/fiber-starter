// Package rest has the setup and route mapping for this microservice's Rest API
package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/plopezlpz/fiber-starter/config/vars"
)

func Setup() *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  vars.AppSetting.ReadTimeout,
		WriteTimeout: vars.AppSetting.WriteTimeout,
		IdleTimeout:  vars.AppSetting.IdleTimeout,
	})
	// middlewares
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(requestid.New())
	app.Use(logger.New())

	// routes
	mapUrls(app)

	return app
}
