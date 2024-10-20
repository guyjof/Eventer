package middlewares

import (
	"net/http"

	"example.com/eventer-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	// Get the token from the request header
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	// Validate the token
	userId, err := utils.ValidateJWTToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	// Set the user ID in the context
	c.Set("userId", userId)
	c.Next()
}
