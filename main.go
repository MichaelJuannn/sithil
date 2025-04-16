package main

import (
	"github.com/gofiber/fiber/v2"
	"sithil/database"
	"sithil/router"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	database.SeedDatabase(25)
	router.SetupRoutes(app)
	app.Listen(":8000")
}
