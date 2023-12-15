package router

import (
    "github.com/zakpaw/gotex/pkg/handler"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes (app *fiber.App) { 
    app.Get("/", handler.IndexPage)

    api := app.Group("/api")
    api.Get("/", handler.GetAPI)
}
