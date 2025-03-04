package config

import (
	"fmt"
	"go-pos/helpers"
	"go-pos/models"
	"log"
)

func UserSeeder() {
	user := models.User{Name: "Super Admin", Username: "admin", Password: helpers.HashPass("admin123")}

	var count int64
	Connection().Model(user).Count(&count)

	if count == 0 {
		if err := Connection().Create(&user).Error; err != nil {
			log.Fatal("User seeder failed")
		}

		fmt.Println("User seeder succeed")
	} else {
		fmt.Println("User seeder skipped")
	}
}

func ProductsSeeder() {
	product := []models.Product{
		{Name: "Holandmie", Price: 3500},
		{Name: "Teh Wangi Sari", Price: 8000},
		{Name: "Sosis SoWell", Price: 12000},
	}

	var count int64
	Connection().Model(product).Count(&count)

	if count == 0 {
		if err := Connection().Create(&product).Error; err != nil {
			log.Fatal("Products seeder failed")
		}

		fmt.Println("Products seeder succeed")
	} else {
		fmt.Println("Products seeder skipped")
	}
}

func RunSeeder() {
	UserSeeder()
	ProductsSeeder()
}
