package main

import (
	"log"
	"os"

	"rs-bandung-api/config"
	"rs-bandung-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		_ = godotenv.Load()
	}

	config.ConnectMongo()

	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running on port", port)
	r.Run(":" + port)
}
