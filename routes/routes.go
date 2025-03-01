package routes

import (
	"github.com/gin-gonic/gin"
	"restapi.com/m/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	server.POST("/signup", signUp)
	server.POST("/login", logIn)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", middlewares.Authenticate, createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

}
