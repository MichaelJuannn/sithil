package productHandler

import (
	"sithil/database"
	"sithil/internals/model"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	db := *database.DB
	var products []model.Product
	if err := db.Find(&products).Error; err != nil {
		c.Status(400).SendString("data not found")
	}
	return c.Status(200).JSON(products)
}
