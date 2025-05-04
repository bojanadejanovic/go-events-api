package main

import (
	"os"
	"strings"

	"bojana.dev/api/db"
	"bojana.dev/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	db.InitDB() // Initialize the database connection and create tables if they don't exist
	server := gin.Default()

	// Add CORS middleware
	server.Use(func(c *gin.Context) {
		allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
		if allowedOrigins == "" {
			allowedOrigins = "http://localhost:5173" // Default to SvelteKit dev server
		}

		origin := c.Request.Header.Get("Origin")
		if strings.Contains(allowedOrigins, origin) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	routes.RegisterRoutes(server) // Register the routes with the server
	server.Run(":8080")
}

func loadEnv() {
	// Load base config first
	_ = godotenv.Load(".env")

	// Then override with local
	_ = godotenv.Overload(".env.local")
}
