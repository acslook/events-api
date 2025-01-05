package main

import (
	"events-api/db"
	"events-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run("localhost:8080")
}
