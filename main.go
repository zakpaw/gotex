package main

import (
    "github.com/zakpaw/gotex/pkg/router"

	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
    app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":3000")
}
