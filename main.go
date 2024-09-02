package main

import (
	events "api/Events"
	"api/db"
	"time"

	"github.com/gin-gonic/gin"
)
func main() {
	server:=gin.Default()
	db.InitDB()
	//endpoints
	server.GET("/",index)
	server.POST("events",createEvent)


	//run server
	server.Run(":8080")
}

func index(context *gin.Context) {
  
	body:=map[string]string{"message":"Hello World!"}
	context.JSON(200,body);
}
func createEvent(context *gin.Context) {
	var event events.Event;
	err:=context.BindJSON(&event);

	if err!=nil{
		context.JSON(400,gin.H{"message": err.Error()})
		return;
	}

	event.Id=1;
	event.DateTime=time.Now();
	event.Save()
	context.JSON(200,event);
}