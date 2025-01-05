package middlewares

import (
	"events-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	c.Set("user_id", userId)
	c.Next()

}
