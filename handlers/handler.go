package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raeandres/golang-rest-product.git/database"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetProducts(c *fiber.Ctx) error {
	// return
	products := database.GetAllProducts(database.DB.Db)
	return c.Status(200).JSON(products)
}
