package routes

import (
	"net/http"

	"example.com/eventer-api/models"
	"example.com/eventer-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		return
	}

	userResponse := models.UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created", "user": userResponse})
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to login, invalid request"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to login, could not authenticate user"})
		return
	}

	token, err := utils.GenerateJWTToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to login, could not authenticate user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
