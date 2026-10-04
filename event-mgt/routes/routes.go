package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvents)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	// User Onboarding
	server.POST("/signup", signup)
	server.POST("/login", login)
}
