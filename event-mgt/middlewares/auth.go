package middlewares

import (
	"net/http"

	"example.com/event-mgt/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")
	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	userId, err := utils.VerifyToken(authHeader)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
