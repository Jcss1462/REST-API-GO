package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"restapi.com/m/models"
)

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

	event.UserId = context.GetInt64("userId")

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo crear el evento", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Evento creado!", "event": event})

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo parsear el id del evento", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo obtener el evento", "error": err.Error()})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "no autorizado a actualizar este evento"})
		return
	}

	var upadterdEvent models.Event
	err = context.ShouldBindBodyWithJSON(&upadterdEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo convertir la data del body", "error": err.Error()})
		return
	}

	upadterdEvent.ID = eventId

	err = upadterdEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo actualizar el evento", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Evento actualizado con exito!"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo parsear el id del evento", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo obtener el evento", "error": err.Error()})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "no autorizado a eliminar este evento"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo eliminar el evento", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Evento eliminado con exito!"})
}
