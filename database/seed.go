package database

import "sithil/internals/model"

func SeedDatabase() {
	db := DB
	var count int64
	db.Table("categories").Count(&count)
	if count > 0 {
		return
	} else {
		// seeding category table
		categories := []model.Category{
			{ID: 1, Name: "Electronics"},
			{ID: 2, Name: "Clothing"},
			{ID: 3, Name: "Books"},
		}
		db.Create(&categories)
	}
}
