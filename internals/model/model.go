package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	UserName string
	Password string
	Cart     Cart
	Orders   []Order
}

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	CreatedAt time.Time
	Products  []Product `gorm:"many2many:cart_products;"`
}

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	CategoryID  uint
}

type Category struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Order struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	OrderDate   time.Time
	TotalAmount int
	Products    []Product `gorm:"many2many:order_products;"`
}

// join tables

type CartProduct struct {
	CartID    uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Quantity  uint
}

type OrderProduct struct {
	OrderID   uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Quantity  uint
}
