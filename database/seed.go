package database

import (
	"sithil/internals/model"

	"github.com/go-faker/faker/v4"
)

func SeedDatabase(row int) {
	// seeding products table
	data := make([]model.Product, 0)
	single := model.Product{}
	for i := 1; i <= row; i++ {
		err := faker.FakeData(&single)
		single.CategoryID = uint(i%3 + 1)
		single.ID = uint(i)
		if err != nil {
			println(err)
		}
		data = append(data, single)
	}
	var productCount int64
	var categoryCount int64
	db := DB
	db.Model(&model.Category{}).Count(&categoryCount)
	db.Model(&model.Product{}).Count(&productCount)
	if productCount < 25 && categoryCount == 0 {
		categories := []model.Category{
			{ID: 1, Name: "Electronics"},
			{ID: 2, Name: "Clothings"},
			{ID: 3, Name: "Furnitures"},
		}
		db.Create(&categories)
		db.Create(&data)
	}
}
