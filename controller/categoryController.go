package controller

import (
	"buyleaf/initializer"
	"buyleaf/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllCategory(c *fiber.Ctx) error {
	var category []models.Category
	err := initializer.DB.Find(&category).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}
	return c.JSON(category)
}

func GetCategory(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is required",
		})
	}
	var category models.Category
	err := initializer.DB.First(&category, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Category not found",
		})
	}
	return c.JSON(category)
}

func CreateCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// TODO add validation
	err := initializer.DB.Create(&category).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(category)
}

func UpdateCategory(c *fiber.Ctx) error {
	categoryUpdate := new(models.Category)
	if err := c.BodyParser(categoryUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err := initializer.DB.Model(&models.Category{}).Select("*").Omit("CreatedAt", "UpdatedAt", "DeletedAt").Where("id", categoryUpdate.ID).Updates(&categoryUpdate).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(categoryUpdate)
}
func DeleteCategory(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is required",
		})
	}
	//TODO add validation
	err := initializer.DB.Delete(&models.Category{}, id).Error
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
