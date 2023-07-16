package bootstrap

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/ViniciusDJM/jusbrasil-teste/api"
	"github.com/ViniciusDJM/jusbrasil-teste/pkg/api"
)

var initAt = time.Now()

func registerRoutes(router fiber.Router) {
	ctrl := api.NewController()
	router.Get("/first-instance", ctrl.FirstInstanceHandler)
}

func SetupApp() *fiber.App {
	app := fiber.New()
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("Server at since " + initAt.Format(time.DateTime))
	})

	app.Use("/api/swagger/*", swagger.HandlerDefault)
	registerRoutes(app.Group("api/v1"))

	return app
}
