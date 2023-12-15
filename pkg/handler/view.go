package handler

import (
    "github.com/zakpaw/gotex/pkg/components"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/adaptor"
    "github.com/a-h/templ"
)

func IndexPage(c *fiber.Ctx) error {
    return adaptor.HTTPHandler(templ.Handler(components.IndexPage()))(c)
}
