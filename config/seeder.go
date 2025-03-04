package config

import (
	"fmt"
	"go-pos/helpers"
	"go-pos/models"
	"log"
)

func UserSeeder() {
	data := models.User{
		Name: "Super Admin", Username: "admin", Password: helpers.HashPass("admin123"),
	}

	var count int64
	Connection().Model(&data).Count(&count)

	if count == 0 {
		if err := Connection().Create(&data).Error; err != nil {
			log.Fatal("User seeder failed")
		}

		fmt.Print("User seeder succeed")
	} else {
		fmt.Print("User seeder skipped")
	}
}

func ProductsSeeder() {
	var product models.Product

	data := []models.Product{
		{Name: "Teh Wangi Sari", Price: 8000},
		{Name: "Sosis SoWell", Price: 12000},
		{Name: "Holandmie", Price: 3500},
	}

	var count int64
	Connection().Model(&product).Count(&count)

	if count == 0 {
		if err := Connection().Create(&data).Error; err != nil {
			log.Fatal("Product seeder failed")
		}

		fmt.Print("Product  seeder succeed")
	} else {
		fmt.Print("Product  seeder skipped")
	}
}

func RunSeeder() {
	UserSeeder()
	ProductsSeeder()
}
