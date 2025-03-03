package main

import (
	"github.com/joho/godotenv"
	"go-pos/config"
	"go-pos/routes"
	"log"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("env not found")
	}

	config.Connection()
	config.RunSeeder()

	err = routes.SetupRoutes().Run(":8080")
	if err != nil {
		log.Fatal("Service is failed to run")
	}
}
