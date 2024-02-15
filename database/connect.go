package database

import (
	"fmt"
	"log"
	"os"
	"sithil/internals/model"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Println("db something wrong in port ", port)
	}

	// connection url
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), port, os.Getenv("DB_NAME"))
	// connect to db and create connection var
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("connection to db failed")
	}
	//Migrate DB care about the order

	if err := DB.SetupJoinTable(&model.Cart{}, "Products", &model.CartProduct{}); err != nil {
		println(err.Error())
		panic("failed to setup join tables")
	}
	if err := DB.SetupJoinTable(&model.Order{}, "Products", &model.OrderProduct{}); err != nil {
		println(err.Error())
		panic("failed to setup join tables")
	}
	DB.AutoMigrate(&model.User{}, &model.Cart{}, &model.Category{}, &model.Product{}, &model.Order{}, &model.CartProduct{}, &model.OrderProduct{})

	fmt.Println("connection to db established")

}
