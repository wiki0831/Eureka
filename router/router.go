package router

import (
	"eureka/handler"
	"eureka/handler/constraint"

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

	// Constrainst Group
	constraints := api.Group("/constraints")
	constraints.Get("/", constraint.GetConstraints)

	advisory := api.Group("/advisory")
	advisory.Post("/", handler.GetAdvisory)
}
