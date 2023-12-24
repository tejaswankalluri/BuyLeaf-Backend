package controller

import (
	"buyleaf/initializer"
	"buyleaf/models"
	"buyleaf/validator"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetBooks docs
//
//	@Summary		Get Books
//	@Description	Get List of Books
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	models.Book
//	@Router			/books [get]
func GetBooks(c *fiber.Ctx) error {
	var Books []models.Book
	initializer.DB.Find(&Books)
	return c.JSON(Books)
}

func GetBookById(c *fiber.Ctx) error {
	id := c.Query("id")
	// check for user input
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "id is needed",
		})
	}
	// validate user input
	if _, err := strconv.Atoi(id); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "id doesnt looks like a number",
		})
	}
	// fetch the data from DB
	var Book models.Book
	initializer.DB.First(&Book, id)

	// if user not found
	if Book.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
		})
	}
	return c.JSON(Book)
}

func PostBooks(c *fiber.Ctx) error {
	Book := new(models.Book)
	// var Book models.Book
	// parsing json

	if err := c.BodyParser(Book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// validation
	errors := validator.ValidateBook(*Book)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	// store data
	initializer.DB.Create(&Book)
	return c.JSON(Book)
}

func DeleteBooks(c *fiber.Ctx) error {
	id := c.Query("id")
	// check for user input
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "id is needed",
		})
	}
	// validate user input
	if _, err := strconv.Atoi(id); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "id doesnt looks like a number",
		})
	}

	// delete the user
	initializer.DB.Delete(models.Book{}, id)

	return c.JSON(fiber.Map{
		"id":      id,
		"deleted": true,
	})
}
