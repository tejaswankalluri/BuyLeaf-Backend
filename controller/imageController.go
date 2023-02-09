package controller

import (
	"fiber-api/initializer"
	"fiber-api/models"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"regexp"
	"strconv"
)

func CreateImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	buffer, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	productID := c.Query("id")
	if productID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is required",
		})
	}

	// Get Buffer from file
	defer buffer.Close()
	resp, err := initializer.Cld.Upload.Upload(initializer.CldCtx, buffer, uploader.UploadParams{Folder: "/ByLeaf"})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}

	// store in db
	var image models.Image
	productIDUint, _ := strconv.ParseUint(productID, 10, 64)
	image.ProductID = uint(productIDUint)
	image.ImageUrl = resp.SecureURL
	initializer.DB.Create(&image)
	return c.JSON(resp)
}

func DeleteImage(c *fiber.Ctx) error {
	var err error
	ImageId := c.Query("id")
	if ImageId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is required",
		})
	}
	// check the public ID
	var DeleteImage models.Image
	err = initializer.DB.First(&DeleteImage, ImageId).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	re := regexp.MustCompile(`upload\/(?:v\d+\/)?([^\.]+)`)
	match := re.FindStringSubmatch(DeleteImage.ImageUrl)
	_, err = initializer.Cld.Upload.Destroy(initializer.CldCtx, uploader.DestroyParams{PublicID: match[1]})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server err",
			"Error":   err,
		})
	}

	// delete the image
	err = initializer.DB.Unscoped().Delete(&models.Image{}, ImageId).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server err",
			"Error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"Message": "Deleted",
		"ID":      ImageId,
	})

}
