package orderService

import (
	"sithil/database"
	"sithil/internals/model"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Checkout(c *fiber.Ctx) error {
	db := database.DB
	t := c.Locals("user").(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)
	userID := claims["id"].(float64)
	amount := c.Query("amount", "0")
	println(amount)
	cart := new(model.Cart)
	if err := db.First(&cart, "user_id = ?", userID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err.Error()})
	}

	var cartProducts []model.CartProduct
	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err.Error()})
	}

	order := model.Order{
		UserID:      uint(userID),
		OrderDate:   time.Now(),
		TotalAmount: amountInt,
	}
	if err := db.Create(&order).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err.Error()})
	}
	var orderProducts []model.OrderProduct
	for _, cartProduct := range cartProducts {
		orderProduct := model.OrderProduct{
			OrderID:   order.ID,
			ProductID: cartProduct.ProductID,
			Quantity:  cartProduct.Quantity,
		}
		orderProducts = append(orderProducts, orderProduct)
	}
	if err := db.Create(&orderProducts).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err.Error()})
	}
	err = clearCart(db, cart.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success"})

}

func clearCart(db *gorm.DB, cartID uint) error {
	if err := db.Exec("DELETE FROM cart_products WHERE cart_id = ?", cartID).Error; err != nil {
		return err
	}
	return nil
}
