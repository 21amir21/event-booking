package main

import (
	"net/http"

	"github.com/21amir21/event-booking/db"
	"github.com/21amir21/event-booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()

	r.GET("/events", func(ctx *gin.Context) {
		events := models.GetAllEevents()
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

		event.Save()

		ctx.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
	})

	r.Run(":8080") // localhost:8080
}
