package controller

import (
	"fiber-api/initializer"
	"fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllProductInventory(c *fiber.Ctx) error {
	var inventory []models.Product
	err := initializer.DB.Preload("ProductInventory").Find(&inventory).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}
	return c.JSON(inventory)
}

func GetProductInventory(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is required",
		})
	}
	var inventory models.Product
	err := initializer.DB.Preload("ProductInventory").First(&inventory, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "ProductInventory not found",
		})
	}
	return c.JSON(inventory)
}

func CreateProductInventory(c *fiber.Ctx) error {
	inventory := new(models.ProductInventory)
	if err := c.BodyParser(inventory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// TODO add validation
	err := initializer.DB.Create(&inventory).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(inventory)
}

func UpdateProductInventory(c *fiber.Ctx) error {
	inventoryUpdate := new(models.ProductInventory)
	if err := c.BodyParser(inventoryUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err := initializer.DB.Model(&models.ProductInventory{}).Select("*").Omit("CreatedAt", "UpdatedAt", "DeletedAt").Where("id", inventoryUpdate.ID).Updates(&inventoryUpdate).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(inventoryUpdate)
}

func DeleteProductInventory(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is required",
		})
	}
	//TODO add validation
	err := initializer.DB.Delete(&models.ProductInventory{}, id).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server err",
			"Error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"Message": "Deleted",
		"ID":      id,
	})
}
