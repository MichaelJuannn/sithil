package userHandler

import (
	"errors"
	"sithil/config"
	"sithil/database"
	"sithil/internals/model"
	"sithil/internals/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func TestJWT(c *fiber.Ctx) error {
	test := c.Locals("user").(*jwt.Token)
	claims := test.Claims.(jwt.MapClaims)
	name := claims["username"].(string)
	return c.SendString("welcome " + name)
}

func Create(c *fiber.Ctx) error {
	db := database.DB
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

	// create cart associated with user after successfull write to user DB
	cart := &model.Cart{UserID: user.ID}
	err = db.Create(&cart).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err.Error()})
	}

	return c.Status(200).JSON(user)
}

func Login(c *fiber.Ctx) error {

	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	input := new(LoginInput)
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error"})
	}
	user, err := new(model.User), *new(error)

	user, err = getUserByEmail(input.Email)
	if err != nil {
		return c.Status(303).JSON(err)
	}

	if !utils.ComparePassword(input.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	// JWT claims
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.UserName
	claims["email"] = user.Email
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(200).JSON(fiber.Map{"status": "authorized", "token": t})

}

func getUserByEmail(e string) (*model.User, error) {
	db := database.DB
	user := new(model.User)
	if err := db.Select("email", "id", "user_name", "password").Where(&model.User{Email: e}).Find(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
