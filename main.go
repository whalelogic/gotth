package main

import (
	"gotth/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/static", "./static")

	app.Get("/", handlers.Home) 

	app.Get("/projects", handlers.Projects)
	app.Get("/about", handlers.About)

	app.Get("/projects/htmx/button", handlers.HandleHtmxButton)
	app.Post("/projects/htmx/form", handlers.HandleHtmxForm)

	app.Get("/projects/htmx/page", handlers.HandlePagination)

	app.Listen(":3000")
}
