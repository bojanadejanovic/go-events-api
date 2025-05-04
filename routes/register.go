package routes

import (
	"net/http"
	"strconv"

	"bojana.dev/api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	userID := context.GetInt64("userID")

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	err = event.RegisterForEvent(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered user for event"})
}

func cancelRegistration(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	loggedInUserID := context.GetInt64("userID")

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	err = event.CancelRegistration(loggedInUserID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled"})
}

func getRegisteredUsers(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	users, err := models.GetRegisteredUsersForEvent(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get registered users"})
		return
	}

	if users == nil {
		context.JSON(http.StatusOK, []string{})
		return
	}

	// Map users to just emails
	emails := make([]string, len(*users))
	for i, user := range *users {
		emails[i] = user.Email
	}

	context.JSON(http.StatusOK, emails)
}
