package userHandler

import (
	"sithil/database"
	"sithil/internals/model"
	"sithil/internals/utils"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	db := *database.DB
	user := new(model.User)
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err.Error()})
	}
	user.Password, _ = utils.HashPassword(user.Password)
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err.Error()})
	}
	return c.Status(200).JSON(&user)
}
