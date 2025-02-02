package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/raeandres/golang-rest-product.git/database"
	"github.com/raeandres/golang-rest-product.git/model"
)

/* create unique randomizer function for productId */
func generateUniqueId() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetProducts(c *fiber.Ctx) error {
	// return
	products := database.GetAllProducts(database.DB.Db)
	c.Set("Content-type", "application/json; charset=utf-8")
	c.Accepts("application/json")
	return c.Status(200).SendString(products)
}

func AddProduct(c *fiber.Ctx) error {
	product := new(model.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	c.Set("Content-type", "application/json; charset=utf-8")
	c.Accepts("application/json")
	addProduct := database.InsertProduct(database.DB.Db, product)
	return c.Status(200).JSON(addProduct)
}

func EditProduct(c *fiber.Ctx) error {
	product := new(model.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	c.Set("Content-type", "application/json; charset=utf-8")
	c.Accepts("application/json")
	updateProduct := database.UpdateProduct(database.DB.Db, product)
	return c.Status(200).SendString(updateProduct)
}
