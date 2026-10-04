package routes

import (
	"net/http"

	"example.com/event-mgt/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	// Bind JSON and validate using Gin
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	err = (&user).Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}
