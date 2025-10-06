package handlers

import (
	"gotth/templates"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// In a real app, this would be a database or other data source.
var projectItems = []string{
	"Project Alpha", "Project Beta", "Project Gamma", "Project Delta",
	"Project Epsilon", "Project Zeta", "Project Eta", "Project Theta",
	"Project Iota", "Project Kappa",
}
const pageSize = 3

func HandlePagination(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "0"))
	if err != nil {
		page = 0
	}

	start := page * pageSize
	end := start + pageSize

	if start >= len(projectItems) {
		// If the start is out of bounds, it means no more items.
		// We render an empty response and an OOB swap to remove the button.
		return Render(c, templates.PaginationPage(nil, -1))
	}

	if end > len(projectItems) {
		end = len(projectItems)
	}

	nextPage := -1
	if end < len(projectItems) {
		nextPage = page + 1
	}

	return Render(c, templates.PaginationPage(projectItems[start:end], nextPage))
}
