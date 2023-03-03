package middleware

import (
	"buyleaf/initializer"
	"buyleaf/models"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

func VerifyCache(c *fiber.Ctx) error {
	id := c.Path()
	val, err := initializer.RedisClient.Get(id).Bytes()
	if err != nil {
		return c.Next()
	}
	var cachedProducts []models.Product
	if err := json.Unmarshal(val, &cachedProducts); err != nil {
		log.Println(err)
	}
	return c.JSON(cachedProducts)
}
