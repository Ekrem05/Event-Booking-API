package main

import (
	"github.com/gin-gonic/gin"
)
func main() {
	server:=gin.Default()
	server.GET("/",index)
	server.Run(":8080")
}

func index(context *gin.Context) {
  
	body:=map[string]string{"message":"Hello World!"}
	context.JSON(200,body);
}