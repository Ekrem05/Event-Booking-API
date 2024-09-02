package main

import (
	"api/db"
	"api/routes"

	"github.com/gin-gonic/gin"
)
func main() {
	server:=gin.Default()
	db.InitDB()
	//endpoints
	routes.AddRoutes(server);

	//run server
	server.Run(":8080")
}
