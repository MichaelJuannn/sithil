package categoriesHandler

import (
	"sithil/database"
	"sithil/internals/model"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	db := database.DB
	categories := new(model.Category)
	if err := c.BodyParser(&categories); err != nil {
		c.Status(400).SendString("invalid data")
	}
	if err := db.Create(&categories).Error; err != nil {
		c.Status(400).SendString("failed to create categories")
	}
	return c.Status(200).JSON(categories)
}

func GetAll(c *fiber.Ctx) error {
	db := database.DB
	var categories []model.Category
	if err := db.Select("name", "id").Find(&categories).Error; err != nil {
		c.Status(400).SendString("Data not found or empty")
	}
	return c.Status(200).JSON(categories)
}
