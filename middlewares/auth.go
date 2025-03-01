package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"restapi.com/m/utils"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No autorizado"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No autorizado", "error": err.Error()})
		return
	}

	context.Set("userId", userId)
	context.Next()

}
