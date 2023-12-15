package handler

import "github.com/gofiber/fiber/v2"

func GetAPI(c *fiber.Ctx) error {
    return c.SendString("Hello, API!")
}
