package main

import (
	events "api/Events"
	"api/db"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)
func main() {
	server:=gin.Default()
	db.InitDB()
	//endpoints
	server.GET("/",index)
	server.POST("events",createEvent)
	server.GET("events/:id",getEvent)


	//run server
	server.Run(":8080")
}

func index(context *gin.Context) {
	events,err:=events.GetAll();
	if err!=nil{
		context.JSON(500,gin.H{"error":"Failed to fetch events"})
		return;

	}

	
	context.JSON(200,events);
}
func createEvent(context *gin.Context) {
	var event events.Event;
	err:=context.BindJSON(&event);

	if err!=nil{
		context.JSON(400,gin.H{"error": err.Error()})
		return;
	}

	event.Id=1;
	event.DateTime=time.Now();
	event.Save()
	context.JSON(200,event);
}

func getEvent(context *gin.Context){
	id,err:=strconv.ParseInt(context.Param("id"),10,64);

	if err!=nil{
		context.JSON(400,gin.H{"error":"Invalid parameter"})
		return;
	}
	event,err:=events.GetById(id)
	if err!=nil{
		context.JSON(500,gin.H{"error":"Failed to get an event by the specified id"})
		return;

	}
	context.JSON(200,event)
}