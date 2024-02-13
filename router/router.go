package router

import (
	productHandler "sithil/internals/service/product"
	userHandler "sithil/internals/service/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// User Api Routes
	user := api.Group("/user")
	user.Post("/register", userHandler.Create)

	// Product Api Routes
	product := api.Group("/product")
	product.Get("/", productHandler.GetAll)
}
