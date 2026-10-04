package main

import (
	"example.com/event-mgt/db"
	"example.com/event-mgt/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
