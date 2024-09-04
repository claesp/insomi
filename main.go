package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var (
	Version string = "0.0.0"
	Port    string = "8080"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Static("/lib", "./lib")

	app.Get("/", getRoot)
	app.Get("/auth", getAuth)
	app.Get("/auth/apple", getAuthApple)

	app.Listen(fmt.Sprintf(":%s", Port))
}
