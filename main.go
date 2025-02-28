package main

import (
	"github.com/gin-gonic/gin"
	"restapi.com/m/db"
	"restapi.com/m/routes"
)

func main() {

	//inicializo la db
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}
