package routes

import (
	"github.com/gin-gonic/gin"
	"restapi.com/m/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	server.POST("/signup", signUp)
	server.POST("/login", logIn)

	//protejo las rutas con jwt que esten dentro del grupo authenticated
	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
}
