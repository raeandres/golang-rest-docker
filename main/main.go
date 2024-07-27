package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raeandres/golang-rest-product.git/database"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
