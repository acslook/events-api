package routes

import (
	"events-api/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("user_id")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := model.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register for event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registered for event", "event": event})
}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("user_id")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := model.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve event"})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not cancel registration for event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cancelled registration for event", "event": event})
}
