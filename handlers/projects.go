package handlers

import (
	"gotth/templates"

	"github.com/gofiber/fiber/v2"
)

func Projects(c *fiber.Ctx) error {
	return Render(c, templates.ProjectsPage())
}
