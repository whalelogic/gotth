package handlers

import (
	"gotth/templates"

	"github.com/gofiber/fiber/v2"
)

func About(c *fiber.Ctx) error {
	return Render(c, templates.AboutPage())
}
