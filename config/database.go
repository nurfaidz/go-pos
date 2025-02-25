package config

import (
	"fmt"
	"go-pos/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connection() *gorm.DB {
	HOST := os.Getenv("HOST")
	USERNAME := os.Getenv("USERNAME")
	PASSWORD := os.Getenv("PASSWORD")
	DBNAME := os.Getenv("DBNAME")
	PORT := os.Getenv("PORT")
	TIMEZONE := os.Getenv("TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", HOST, USERNAME, PASSWORD, DBNAME, PORT, TIMEZONE)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database is failed to connect")
	}

	fmt.Println("Database connected")

	err = db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Transaction{},
		&models.TransactionList{},
	)

	if err != nil {
		log.Fatal("Migration if failed")
	}

	fmt.Println("Migration succeed")

	return db
}
