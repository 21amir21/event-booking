package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/21amir21/event-booking/models"
	"github.com/21amir21/event-booking/utils"
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
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"mesage": "Not Authorized"})
		return
	}

	err := utils.VerfiyToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"mesage": "Not Authorized"})
		return
	}

	var event models.Event
	err = c.ShouldBindJSON(&event)

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

func updateEvent(c *gin.Context) {
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Enter an id in the correct form"})
		return
	}

	_, err = models.GetEventByID(i)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Could not retrive the event"})
		return
	}

	var updatedEvent models.Event

	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Faild to bind json with event data id = %v", i)})
		return
	}

	updatedEvent.ID = int64(i)

	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Faild to update the event with id = %v", i)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated the event!"})
}

func deleteEvent(c *gin.Context) {
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Enter an id in the correct form"})
		return
	}

	event, err := models.GetEventByID(i)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Event with id %v doesn't exist", i)})
		return
	}

	if err = event.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Faild while deleting the event!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted the event!"})
}
