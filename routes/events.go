package routes

import (
	"events-api/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := model.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"events": events})
}

func getEventsById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := model.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

func createEvent(c *gin.Context) {
	var event model.Event
	err := c.BindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not parse request data": err.Error()})
	}

	userId := c.GetInt64("user_id")
	event.UserID = userId
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var event model.Event
	err = c.BindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not parse request data": err.Error()})
	}

	eventUpdate, err := model.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetInt64("user_id")
	if eventUpdate.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to update this event"})
		return
	}

	eventUpdate.Name = event.Name
	eventUpdate.Description = event.Description
	eventUpdate.DateTime = event.DateTime
	eventUpdate.Location = event.Location

	err = eventUpdate.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event updated!", "event": eventUpdate})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := model.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	userId := c.GetInt64("user_id")
	if event.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to delete this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted!"})
}
