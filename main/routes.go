package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raeandres/golang-rest-product.git/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/products", handlers.GetProducts)
}
