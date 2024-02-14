package router

import (
	"sithil/internals/middleware"
	productHandler "sithil/internals/service/product"
	userHandler "sithil/internals/service/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// User Api Routes
	user := api.Group("/user")
	user.Get("/", middleware.Protected(), userHandler.TestJWT)
	user.Post("/register", userHandler.Create)
	user.Post("/login", userHandler.Login)

	// Product Api Routes
	product := api.Group("/product")
	product.Get("/", productHandler.GetAll)
}
