package main

import (
	"buyleaf/initializer"
	"buyleaf/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializer.InitializeEnv()
	initializer.InitDB()
	initializer.InitRedis()
	initializer.InitCLD()
}

func main() {

	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// routes
	routes.Routes(app)
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Ecommerce Metrics"}))

	// server port
	err := app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalln(err)
		return
	}
}
