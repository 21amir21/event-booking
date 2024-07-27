package routes

import (
	"net/http"

	"github.com/21amir21/event-booking/models"
	"github.com/21amir21/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data."})
		return
	}

	if err := user.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not save user."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created Successfully"})
}

func login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data."})
		return
	}

	if err := user.ValidateCredentials(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login Successfull!", "token": token})
}
