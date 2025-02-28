package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"restapi.com/m/db"
	"restapi.com/m/models"
)

func main() {

	//inicializo la db
	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo obtener los eventos", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo parsear el id del evento", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo obtener el evento", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	//bindea la entradaal tipo de objeto de la estructura
	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo convertir la data del body", "error": err.Error()})
		return
	}

	event.ID = 1
	event.UserId = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo crear el evento", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Evento creado!", "event": event})

}
