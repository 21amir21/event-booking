package middlewares

import (
	"net/http"

	"github.com/21amir21/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"mesage": "Not Authorized"})
			return
		}

		userId, err := utils.VerfiyToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"mesage": "Not Authorized", "error": err.Error()})
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}
