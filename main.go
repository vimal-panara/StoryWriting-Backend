package main

import (
	"log"
	"os"
	"story-plateform/config"
	"story-plateform/controllers"
	"story-plateform/routes"

	"github.com/joho/godotenv"
)

func main() {
	go controllers.RunHub()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file -", err)
	}

	config.ConnectToDatabase()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := routes.SetupRouter()
	router.Run(":" + port)
}
