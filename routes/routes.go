package routes

import (
	"net/http"

	"bojana.dev/api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	{
		authenticated.POST("/events", createEvent)
		authenticated.PUT("/events/:id", updateEvent)
		authenticated.DELETE("/events/:id", deleteEvent)
		authenticated.POST("/events/:id/register", registerForEvent)
		authenticated.DELETE("/events/:id/register", cancelRegistration)
	}

	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/test", dummyEndpoint)
}

func dummyEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
