package main

import (
	"events-api/db"
	"events-api/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("api/v1/events", getEvents)
	server.POST("api/v1/events", createEvent)

	server.Run("localhost:8080")
}

func getEvents(c *gin.Context) {
	events, err := model.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvent(c *gin.Context) {
	var event model.Event
	err := c.BindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not parse request data": err.Error()})
	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
