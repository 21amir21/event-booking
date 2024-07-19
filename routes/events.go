package routes

import (
	"net/http"
	"strconv"

	"github.com/21amir21/event-booking/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEevents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
		return
	}
	c.JSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Enter an id in the correct form"})
		return
	}

	event, err := models.GetEventByID(i)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Could not retrive the event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse this request"})
		return
	}

	event.ID = 1
	event.UserID = 1

	if err = event.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
