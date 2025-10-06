package handlers

import (
	"gotth/templates"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return Render(c, templates.HomePage())
}

