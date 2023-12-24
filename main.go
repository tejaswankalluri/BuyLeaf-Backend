package main

import (
	_ "buyleaf/docs"
	"buyleaf/initializer"
	"buyleaf/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializer.InitializeEnv()
	initializer.InitDB()
	initializer.InitRedis()
	initializer.InitCLD()
}

// @title			BuyLeaf
// @version		1.0
// @description	It is a ecommerce backend built with golang
// @termsOfService	http://swagger.io/terms/
// @contact.name	Tejaswan Kalluri
// @contact.email	tejaswan@proton.me
// @license.name	MIT
// @license.url	https://mit-license.org/
// @host			localhost:8000
// @BasePath		/
func main() {

	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Access-Control-Allow-Headers",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

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
