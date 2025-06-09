// main.go
package main

import (
	"log"

	"github.com/dulsaragimhana/restaurant-api/config"
	"github.com/dulsaragimhana/restaurant-api/database"
	"github.com/dulsaragimhana/restaurant-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to database
	database.ConnectDB()

	// Set up Gin
	r := gin.Default()
	routes.SetupRoutes(r)

	// Run server
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
