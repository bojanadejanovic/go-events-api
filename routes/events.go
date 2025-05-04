package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"bojana.dev/api/models"
	"bojana.dev/api/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	models, err := models.GetAllEvents()
	if err != nil {
		fmt.Printf("Error getting events: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}
	context.JSON(http.StatusOK, models)
}

func createEvent(context *gin.Context) {
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		utils.HandleValidationError(err, context)
		return
	}

	event.UserID = context.GetInt64("userID")
	err := event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save event"})
		return
	}
	context.JSON(http.StatusCreated, event)
}

func getEventByID(context *gin.Context) {
	id := context.Param("id")
	eventID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Could not find event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {

	id := context.Param("id")
	eventID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	loggedInUserID := context.GetInt64("userID")

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Could not find event"})
		return
	}

	if event.UserID != loggedInUserID {
		context.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		utils.HandleValidationError(err, context)
		return
	}

	updatedEvent.ID = eventID // Set the ID of the updated event
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	id := context.Param("id")
	loggedInUserID := context.GetInt64("userID")

	eventID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Could not find event"})
		return
	}

	if event.UserID != loggedInUserID {
		context.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this event"})
		return
	}
	err = event.Delete(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
