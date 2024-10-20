package routes

import (
	"example.com/eventer-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(server *gin.Engine) {
	server.GET("/readiness-check", ReadinessCheck)
	server.GET("/liveness-check", LivenessCheck)
}

func RegisterApplicationRoutes(server *gin.Engine) {
	// Event routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// Protected Event routes
	authenticatedRoute := server.Group("/")
	authenticatedRoute.Use(middlewares.Authenticate)

	authenticatedRoute.POST("/events", createEvent)
	authenticatedRoute.PUT("/events/:id", updateEvent)
	authenticatedRoute.DELETE("/events/:id", deleteEvent)
	authenticatedRoute.POST("/events/:id/register", registerForEvent)
	authenticatedRoute.DELETE("/events/:id/register", cancelRegistrationForEvent)

	// User routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
