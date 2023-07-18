package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func globalMiddlewares(router fiber.Router) {
	router.Use(recover.New())
	router.Use(logger.New())
}
