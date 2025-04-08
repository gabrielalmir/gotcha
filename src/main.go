package main

import (
	"gotcha/src/db"
	"gotcha/src/routes"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	database := db.NewDatabase(os.Getenv("MONGO_URI"))
	if err := database.Connect(); err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Initialize router
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Setup routes
	routes.SetupRoutes(router)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
