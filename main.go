package main

import (
	"sithil/database"
	"sithil/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	database.SeedDatabase(25)
	router.SetupRoutes(app)
	app.Listen(":8000")
}
