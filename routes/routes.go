package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.GET("/events", getEvents)
	router.GET("/events/:id", getEvent)
	router.POST("/events", createEvent)
}