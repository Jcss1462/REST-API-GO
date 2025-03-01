package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"restapi.com/m/models"
)

func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
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

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo registrar el usuario para el evento", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registrado!"})
}

func cancelRegistration(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo parsear el id del evento", "error": err.Error()})
		return
	}

	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo cancelar registro", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelado!"})
}
