package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadinessCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func LivenessCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
