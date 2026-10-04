package routes

import (
	"example.com/event-mgt/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvents)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// User Onboarding
	server.POST("/signup", signup)
	server.POST("/login", login)
}
