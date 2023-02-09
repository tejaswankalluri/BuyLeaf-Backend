package middleware

import (
	"fiber-api/initializer"
	"fiber-api/models"
	"fiber-api/service"
	"fiber-api/util"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func AuthAdmin(c *fiber.Ctx) error {
	c.Locals("isAdmin", true)
	return c.Next()
}
func AuthMiddleware(c *fiber.Ctx) error {

	//	Get the token
	tokenString := c.Cookies("Authorization")
	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "No token provided!",
		})
	}

	// validate the token
	_, err, claims := service.JwtValid(tokenString)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Token!",
		})
	}

	// find the user
	var user models.User
	initializer.DB.First(&user, int(claims["sub"].(float64)))
	util.SanitizeUserModel(&user)
	if user.ID == 0 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Token!",
		})
	}
	//	check if admin
	admin := c.Locals("isAdmin")
	if admin != nil {
		if user.Role != "ADMIN" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Only admin route!",
			})
		}
	}

	// attach to req
	c.Locals("user", user)

	//	continue
	return c.Next()

}
