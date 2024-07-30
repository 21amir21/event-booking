package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/21amir21/event-booking/models"
	"github.com/gin-gonic/gin"
)

func registerEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId := c.Param("id")

	i, err := strconv.Atoi(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Enter an id in the correct form"})
		return
	}

	event, err := models.GetEventByID(i)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Could not register user with id %v .", userId)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId := c.Param("id")

	i, err := strconv.Atoi(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Enter an id in the correct form"})
		return
	}

	event, err := models.GetEventByID(i)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.CancelRegistration(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Could not delete user with id %v .", userId)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration Cancelled Successfully!"})
}
