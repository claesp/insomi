package main

import (
	"github.com/gofiber/fiber/v2"
)

func getAuth(c *fiber.Ctx) error {
	return c.Render("auth", fiber.Map{
		"Title": "Sign in",
	})
}

func getAuthApple(c *fiber.Ctx) error {
	return c.Render("auth_apple", fiber.Map{
		"Title": "Sign in with Apple",
	})
}
