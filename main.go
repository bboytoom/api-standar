package main

import (
	"api_standard/www"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvironmentFile()

	app_port := os.Getenv("APP_PORT")
	get_server := www.ServerConfig(app_port)

	get_server.Listener()
}

func loadEnvironmentFile() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}
}
