package routes

import (
	"github.com/21amir21/event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/events", getEvents)
	router.GET("/events/:id", getEvent)

	authenticated := router.Group("/")
	authenticated.Use(middlewares.Authenticate())
	{
		authenticated.POST("/events", createEvent)
		authenticated.PUT("/events/:id", updateEvent)
		authenticated.DELETE("/events/:id", deleteEvent)
		authenticated.POST("/events/:id/register", registerEvent)
		authenticated.DELETE("/events/:id/register", cancelRegistration)
	}

	router.POST("/signup", signup)
	router.POST("/login", login)
}
