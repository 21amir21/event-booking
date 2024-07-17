package main

import (
	"net/http"
	"strconv"

	"github.com/21amir21/event-booking/db"
	"github.com/21amir21/event-booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()

	r.GET("/events", func(ctx *gin.Context) {
		events, err := models.GetAllEevents()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
			return
		}
		ctx.JSON(http.StatusOK, events)
	})

	r.POST("/events", func(ctx *gin.Context) {
		var event models.Event
		err := ctx.ShouldBindJSON(&event)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse this request"})
			return
		}

		event.ID = 1
		event.UserID = 1

		if err = event.Save(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event."})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
	})

	r.GET("/events/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		i, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Enter an id in the correct form"})
			return
		}

		event, err := models.GetEventByID(i)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Could not retrive the event"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"event": event})
	})

	r.Run(":8080") // localhost:8080
}
