package middleware

import (
	"buyleaf/initializer"
	"buyleaf/models"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func VerifyCache(c *fiber.Ctx) error {
	id := c.Path()
	val, err := initializer.RedisClient.Get(id).Bytes()
	if err != nil {
		return c.Next()
	}
	var dat []models.Product
	if err := json.Unmarshal(val, &dat); err != nil {
		panic(err)
	}
	return c.JSON(dat)
}