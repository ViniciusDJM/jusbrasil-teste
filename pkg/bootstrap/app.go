package bootstrap

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/ViniciusDJM/jusbrasil-teste/api"
)

// initAt stores the time when the application is initialized.
var initAt = time.Now()

// SetupApp sets up the Fiber app with routes, middlewares, and Swagger documentation.
func SetupApp() *fiber.App {
	app := fiber.New()

	// Endpoint for health check, returning the server's start time.
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("Server at since " + initAt.Format(time.DateTime))
	})

	// Global middlewares applied to all routes.
	globalMiddlewares(app)

	// Swagger documentation for the API at "/api/swagger/*".
	app.Use("/api/swagger/*", swagger.HandlerDefault)

	// Register routes under the "api/v1" group.
	registerRoutes(app.Group("api/v1"))

	return app
}
