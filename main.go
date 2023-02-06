package main

import (
	"github.com/joho/godotenv"
	"log"
	"vettel-backend-app/src/infrastructure/web_server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	web_server.Start()
}
