package routes

import (
	"fiber-api/controller"
	"fiber-api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func Routes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Message": "Ecommerce Backend with golang",
		})
	})
	app.Get("/books", controller.GetBooks)
	app.Get("/book", controller.GetBookById)
	app.Post("/book", controller.PostBooks)
	app.Delete("/book", controller.DeleteBooks)

	// limter
	app.Use(limiter.New(limiter.Config{
		Max:        5,
		Expiration: 10 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "To many request plz try later",
			})
		},
	}))

	//	user
	app.Post("/signup", controller.RegisterUser)
	app.Post("/signin", controller.LoginUser)
	app.Post("/refreshtoken", middleware.AuthMiddleware, controller.RefreshToken)
	app.Post("/logout", middleware.AuthMiddleware, controller.LogoutUser)
	app.Get("/whoami", middleware.AuthMiddleware, controller.Whoami)

	// products
	app.Get("/products", controller.GetAllProducts)
	app.Get("/product", controller.GetProduct)
	app.Post("/product", middleware.AuthAdmin, middleware.AuthMiddleware, controller.CreateProduct)
	app.Put("/product", middleware.AuthAdmin, middleware.AuthMiddleware, controller.UpdateProduct)
	app.Delete("/product", middleware.AuthAdmin, middleware.AuthMiddleware, controller.DeleteProduct)

	// Categories
	app.Get("/categories", controller.GetAllCategory)
	app.Get("/category", controller.GetCategory)
	app.Post("/category", middleware.AuthAdmin, middleware.AuthMiddleware, controller.CreateCategory)
	app.Put("/category", middleware.AuthAdmin, middleware.AuthMiddleware, controller.UpdateCategory)
	app.Delete("/category", middleware.AuthAdmin, middleware.AuthMiddleware, controller.DeleteCategory)

	// Product Inventory
	app.Get("/productsinventory", controller.GetAllProductInventory)
	app.Get("/productinventory", controller.GetProductInventory)
	app.Post("/productinventory", middleware.AuthAdmin, middleware.AuthMiddleware, controller.CreateProductInventory)
	app.Put("/productinventory", middleware.AuthAdmin, middleware.AuthMiddleware, controller.UpdateProductInventory)
	app.Delete("/productinventory", middleware.AuthAdmin, middleware.AuthMiddleware, controller.DeleteProductInventory)

	// upload images
	app.Post("/uploadimage", middleware.AuthAdmin, middleware.AuthMiddleware, controller.CreateImage)
	app.Delete("/uploadimage", middleware.AuthAdmin, middleware.AuthMiddleware, controller.DeleteImage)
}
