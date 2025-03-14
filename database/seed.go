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
	var count int64
	db := DB
	db.Model(&model.Product{}).Count(&count)
	if count < 75 {
		db.Create(&data)
	}
}
