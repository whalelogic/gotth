package handlers

import (
	"gotth/templates"

	"github.com/gofiber/fiber/v2"
)

func HandleHtmxButton(c *fiber.Ctx) error {
	return Render(c, templates.HtmxButtonContent())
}

func HandleHtmxForm(c *fiber.Ctx) error {
	name := c.FormValue("name")
	return Render(c, templates.HtmxFormResult(name))
}
