package routes

import (
	"net/http"
	"strconv"

	"example.com/eventer-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch events"})
		return
	}

	c.JSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	event, err := models.GetEventByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	var event models.Event

	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	userId := c.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event"})
		return
	}

	if event.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"message": "Not authorized to update this event"})
		return
	}

	var updatedEvent models.Event

	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	updatedEvent.ID = id

	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event"})
		return
	}

	userId := c.GetInt64("userId")
	if event.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"message": "Not authorized to delete this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
