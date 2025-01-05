package routes

import (
	"events-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("api/v1/events", getEvents)
	server.GET("api/v1/events/:id", getEventsById)

	authenticated := server.Group("api/v1")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", middlewares.Authenticate, createEvent)
	authenticated.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	authenticated.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("api/v1/signup", signup)
	server.POST("api/v1/login", login)
}
