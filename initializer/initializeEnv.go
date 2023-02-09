package initializer

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func InitializeEnv() {
	var err error

	if !fiber.IsChild() {
		err = godotenv.Load()
		log.Println("Loaded .env")
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
