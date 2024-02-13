package productHandler

import (
	"sithil/database"
	"sithil/internals/model"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	category := c.Query("category")
	db := *database.DB
	var products []model.Product
	// fmt.Printf(" \n %s \n", category)
	if category == "" {
		print("NO QUERY")
		if err := db.Find(&products).Error; err != nil {
			c.Status(400).SendString("data not found")
		}
		return c.Status(200).JSON(products)
	}
	if err := db.Joins("INNER JOIN categories on products.category_id = categories.id").Where("categories.name = ?", category).Find(&products).Error; err != nil {
		c.Status(400).SendString("data not found")
	}
	return c.Status(200).JSON(products)
}
