package main

import (
	"bojana.dev/api/db"
	"bojana.dev/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	db.InitDB() // Initialize the database connection and create tables if they don't exist
	server := gin.Default()
	routes.RegisterRoutes(server) // Register the routes with the server
	server.Run(":8080")
}

func loadEnv() {
	// Load base config first
	_ = godotenv.Load(".env")

	// Then override with local
	_ = godotenv.Overload(".env.local")
}
