package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"restapi.com/m/models"
)

func signUp(context *gin.Context) {

	var user models.User
	//bindea la entradaal tipo de objeto de la estructura
	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo convertir la data del body", "error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo registrar el usuario", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Usuario registrado!"})

}
