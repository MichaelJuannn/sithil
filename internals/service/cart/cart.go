package cartService

import (
	"sithil/database"
	"sithil/internals/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Add(c *fiber.Ctx) error {
	db := database.DB
	t := c.Locals("user").(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)
	userID := claims["id"].(float64)
	reqProduct := c.Query("product", "empty")
	ProductID, _ := strconv.Atoi(reqProduct)
	var user model.User
	var cartProduct model.CartProduct

	// get user and cart
	if err := db.Preload("Cart").First(&user, int(userID)).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err.Error()})
	}

	// check if product is already in cart
	if err := db.First(&cartProduct, "cart_id = ? AND product_id = ?", user.Cart.ID, reqProduct).Error; err != nil {
		cartProduct = model.CartProduct{CartID: user.Cart.ID, ProductID: uint(ProductID), Quantity: 1}
		db.Create(&cartProduct)
	} else {
		cartProduct.Quantity++
		db.Save(&cartProduct)
	}

	return c.Status(200).JSON(fiber.Map{"status": "success"})
}
