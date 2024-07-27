package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raeandres/golang-rest-product.git/database"
	"github.com/raeandres/golang-rest-product.git/model"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetProducts(c *fiber.Ctx) error {
	// return
	products := database.GetAllProducts(database.DB.Db)
	return c.Status(200).JSON(products)
}

func AddProduct(c *fiber.Ctx) error {
	product := new(model.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	addProduct := database.InsertProduct(database.DB.Db, product)
	return c.Status(200).JSON(addProduct)
}
