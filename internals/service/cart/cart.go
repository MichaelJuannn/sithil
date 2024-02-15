package cartService

import (
	"sithil/database"
	"sithil/internals/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ProductWithQuantity struct {
	model.Product
	Quantity uint
}

func Add(c *fiber.Ctx) error {
	db := database.DB
	t := c.Locals("user").(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)
	userID := claims["id"].(float64)
	reqProduct := c.Query("product", "empty")

	if reqProduct == "empty" {
		return c.Status(400).JSON(fiber.Map{"status": "no product id"})
	}

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

func GetCart(c *fiber.Ctx) error {

	db := database.DB
	t := c.Locals("user").(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)
	userID := claims["id"].(float64)
	var user model.User
	if err := db.Preload("Cart").First(&user, "id = ?", userID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err.Error()})
	}
	var products []ProductWithQuantity
	db.Raw(`SELECT p.*, cp.quantity 
	FROM products p 
	INNER JOIN cart_products cp ON cp.product_id = p.id 
	WHERE cp.cart_id = ?;`, user.Cart.ID).Scan(&products)
	if len(products) == 0 {
		return c.Status(200).JSON(fiber.Map{"status": "cart is empty"})

	}

	return c.Status(200).JSON(products)
}

func DeleteProduct(c *fiber.Ctx) error {
	db := database.DB
	t := c.Locals("user").(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)
	userID := claims["id"].(float64)
	reqProduct := c.Query("product", "empty")
	if reqProduct == "empty" {
		return c.Status(400).JSON(fiber.Map{"status": "no product id"})
	}

	var user model.User
	if err := db.Preload("Cart").First(&user, "id = ?", userID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err.Error()})
	}
	var cartProduct model.CartProduct
	if err := db.Delete(&cartProduct, "cart_id = ? AND product_id = ?", user.Cart.ID, reqProduct).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success"})

}
