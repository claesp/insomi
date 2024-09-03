package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

var (
	Version string = "0.0.0"
	Port    string = "8080"
)

func main() {
	fmt.Fprintf(os.Stdout, "%s\n", Version)

	app := fiber.New()

	app.Get("/", root)
	app.Listen(fmt.Sprintf(":%s", Port))
}

func root(c *fiber.Ctx) error {
	return c.SendString("hello world")
}
