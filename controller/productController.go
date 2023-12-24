package controller

import (
	"buyleaf/initializer"
	"buyleaf/models"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetAllProducts docs
//	@Summary		List Products
//	@Description	Get List of Products
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	models.Product
//	@Router			/products [get]
func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product

	//err := initializer.DB.Find(&products).Association("product_col")
	err := initializer.DB.Preload("Category").Preload("ProductImages").Order("product_name").Find(&products).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	// caching with redis
	redisproducts, _ := json.Marshal(products)

	rediserr := initializer.RedisClient.Set(c.Path(), redisproducts, 10*time.Second).Err()
	if rediserr != nil {
		log.Fatal(rediserr)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Cacheing error"})
	}

	return c.JSON(products)
}

// GetProduct docs
//	@Summary		Get Product
//	@Description	Get Product with id
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	query		integer	true	"product id"
//	@Success		200	{object}	models.Product
//	@Router			/product [get]
func GetProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is required",
		})
	}
	// TODO add validation

	var product models.Product
	err := initializer.DB.Preload("Category").First(&product, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}
	return c.JSON(product)
}
func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// TODO add validation
	err := initializer.DB.Create(&product).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {

	productUpdate := new(models.Product)

	if err := c.BodyParser(productUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// TODO add validation
	err := initializer.DB.Model(&models.Product{}).Select("*").Omit("CreatedAt", "UpdatedAt", "DeletedAt").Where("id", productUpdate.ID).Updates(&productUpdate).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(productUpdate)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is required",
		})
	}
	//TODO add validation
	err := initializer.DB.Unscoped().Delete(&models.Product{}, id).Error
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
