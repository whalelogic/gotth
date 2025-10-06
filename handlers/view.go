package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Type("html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
