package routes

import (
	"net/http"

	"github.com/21amir21/event-booking/components"
	"github.com/21amir21/event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/events", getEvents)
	router.GET("/events/:id", getEvent)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", components.Page(0, 0))
	})

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
