// File: main.go
package main

import (
	"go-big-internal/config"
	"go-big-internal/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init DB & Migrate
	db := config.InitDB()
	if err := config.Migrate(db); err != nil {
		log.Fatal("Migration error:", err)
	}

	// Setup router & run
	r := routes.SetupRouter(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
