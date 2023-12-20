package handler

import (
	"github.com/zakpaw/gotex/pkg/view"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func IndexPage(c *fiber.Ctx) error {
    return adaptor.HTTPHandler(templ.Handler(view.IndexPage()))(c)
}
