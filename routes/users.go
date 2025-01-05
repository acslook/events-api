package routes

import (
	"events-api/internal/model"
	"events-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not parse request data": err.Error()})
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created!"})
}

func login(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not parse request data": err.Error()})
	}
	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User authenticated!", "token": token})
}
