package main

import (
	"github.com/gofiber/fiber/v2"
)

func getRoot(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Home",
	})
}
