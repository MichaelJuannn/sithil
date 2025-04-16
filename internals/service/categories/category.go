package categoriesHandler

import (
	"sithil/database"
	"sithil/internals/model"

	"github.com/gofiber/fiber/v2"
)

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PostResponse struct {
	Status string `json:"status"`
}

func Create(c *fiber.Ctx) error {
	db := database.DB
	categories := new(model.Category)
	if err := c.BodyParser(&categories); err != nil {
		c.Status(400).SendString("invalid data")
	}
	if err := db.Create(&categories).Error; err != nil {
		c.Status(400).SendString("failed to create categories")
	}
	response := PostResponse{Status: "success"}
	return c.Status(200).JSON(response)
}

func GetAll(c *fiber.Ctx) error {
	db := database.DB
	var categories []CategoryResponse
	if err := db.Model(&model.Category{}).Select("id", "name").Find(&categories).Error; err != nil {
		c.Status(400).SendString("Data not found or empty")
	}
	return c.Status(200).JSON(categories)
}
