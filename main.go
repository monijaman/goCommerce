package main

import (
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.SetupEnv()
	if err != nil {
		log.Printf("config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)
}
