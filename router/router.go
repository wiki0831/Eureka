package router

import (
	"eureka/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.Welcome)

	api := app.Group("/api", logger.New())
	api.Get("/", handler.Welcome)
	api.Get("/ping", handler.Pong)
	api.Get("/health", handler.HealthCheck)

	advisory := api.Group("/advisory")
	advisory.Post("/", handler.GetAdvisory)
}
