package controller

import (
	"fiber-api/config"
	"fiber-api/initializer"
	"fiber-api/models"
	"fiber-api/service"
	"fiber-api/util"
	"fiber-api/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
	"time"
)

func RegisterUser(c *fiber.Ctx) error {
	user := new(models.User)

	//	parsing json
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// validation
	errors := validator.ValidateUser(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	//check user exist
	isUserEmailExist := new(models.User)
	initializer.DB.Where("email = ?", user.Email).First(&isUserEmailExist)
	if isUserEmailExist.Email != "" {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "user already exist",
		})
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "unable to hash the password",
		})
	}
	user.Password = string([]byte(hash))

	//store user in db
	prodEnv := os.Getenv("ENV")
	var result *gorm.DB

	if prodEnv == "production" {
		result = initializer.DB.Omit("role").Create(&user)
	} else {
		result = initializer.DB.Create(&user)
	}

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "unable to create user",
			"Error":   result.Error,
		})
	}

	util.SanitizeUserModel(*&user)
	return c.Status(http.StatusOK).JSON(user)
}
func LoginUser(c *fiber.Ctx) error {
	userBody := new(validator.Logininput)

	//	parsing json
	if err := c.BodyParser(userBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// validation
	errors := validator.ValidateLogin(*userBody)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Get the user from DB
	user := new(models.User)
	initializer.DB.First(&user, "email = ?", userBody.Email)
	if user.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userBody.Password))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	//	generate Token
	var JwtError error
	var tokenString string
	tokenString, JwtError = service.JwtSignIn(user.ID, config.JwtTokenExpiration)

	if JwtError != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "unable to generate the token",
		})
	}
	//	generate Token
	var JwtRefreshError error
	var RefreshTokenString string
	RefreshTokenString, JwtRefreshError = service.JwtSignIn(user.ID, config.JwtRefreshTokenExpiration)

	if JwtRefreshError != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "unable to generate the token",
		})
	}
	//	store into the redis db
	var rediserr error
	rediserr = initializer.RedisClient.Set(
		"refresh_"+strconv.Itoa(int(user.ID)),
		RefreshTokenString,
		config.JwtRefreshTokenExpiration*time.Hour,
	).Err()
	if rediserr != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "unable to generate the token",
			"err":     rediserr,
		})
	}
	// set cookie token
	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.HTTPOnly = true
	c.Cookie(cookie)

	//	return token
	return c.JSON(fiber.Map{
		"token":         tokenString,
		"refresh_token": RefreshTokenString,
	})
}

func RefreshToken(c *fiber.Ctx) error {
	userBody := new(validator.RefreshToken)

	//	parsing json
	if err := c.BodyParser(userBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// Validate the input
	errors := validator.ValidateRefreshToken(*userBody)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// validate the token
	_, err, claims := service.JwtValid(userBody.Refreshtoken)

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Token!",
		})
	}
	// validate from db
	val, err := initializer.RedisClient.Get("refresh_" + strconv.Itoa(int(claims["sub"].(float64)))).Result()

	if err != nil || val != userBody.Refreshtoken {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Token!",
		})
	}

	//	generate Token
	var JwtError error
	var tokenString string
	tokenString, JwtError = service.JwtSignIn(claims["sub"], config.JwtTokenExpiration)

	if JwtError != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "unable to generate the token",
		})
	}
	//	generate RefreshToken
	var JwtRefreshError error
	var RefreshTokenString string
	RefreshTokenString, JwtRefreshError = service.JwtSignIn(claims["sub"], config.JwtRefreshTokenExpiration)

	if JwtRefreshError != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "unable to generate the token",
		})
	}
	//	store into the redis db
	var rediserr error
	rediserr = initializer.RedisClient.Set(
		"refresh_"+strconv.Itoa(int(claims["sub"].(float64))),
		RefreshTokenString,
		config.JwtRefreshTokenExpiration*time.Hour,
	).Err()

	if rediserr != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "unable to generate the token",
			"err":     rediserr,
		})
	}
	// set cookie token
	//cookie := new(fiber.Cookie)
	//cookie.Name = "Authorization"
	//cookie.Value = tokenString
	//cookie.Expires = time.Now().Add(time.Hour * 24)
	//cookie.HTTPOnly = true
	//c.Cookie(cookie)

	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})
	//	return token
	return c.JSON(fiber.Map{
		"token":         tokenString,
		"refresh_token": RefreshTokenString,
	})
}
func LogoutUser(c *fiber.Ctx) error {
	user := c.Locals("user")
	id := strconv.Itoa(int(user.(models.User).ID))
	initializer.RedisClient.Del("refresh_" + id)
	c.ClearCookie()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logout",
	})
}

func Whoami(c *fiber.Ctx) error {
	user := c.Locals("user")
	return c.JSON(user)
}
